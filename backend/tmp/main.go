package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ... [остальные константы без изменений] ...

// Структура для информации о пользователе
type UserInfo struct {
	ID    string `json:"rsaag_id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	// Добавьте другие поля, которые вы хотите получить
}

// Хранилище сессий
type SessionStore struct {
	sessions map[string]Session
	mu       sync.RWMutex
}

type Session struct {
	UserInfo UserInfo
	Expires  time.Time
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]Session),
	}
}

func (s *SessionStore) Set(sessionID string, userInfo UserInfo) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[sessionID] = Session{
		UserInfo: userInfo,
		Expires:  time.Now().Add(1 * time.Hour), // Сессия на 1 час
	}
}

func (s *SessionStore) Get(sessionID string) (UserInfo, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[sessionID]
	if !exists || time.Now().After(session.Expires) {
		return UserInfo{}, false
	}
	return session.UserInfo, true
}

var sessionStore = NewSessionStore()

func main() {
	// ... [остальная часть main без изменений] ...

	// Новый эндпоинт для получения данных пользователя
	http.HandleFunc("/userinfo", handleUserInfo)

	fmt.Println("Сервер запущен на http://hackathon-5.orb.ru")
	http.ListenAndServe(":8080", nil)
}

// ... [остальные обработчики без изменений до handleCallback] ...

// Обработчик для редиректа после авторизации
func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Код авторизации не получен", http.StatusBadRequest)
		return
	}

	// Получаем токены
	tokenResp, err := getTokens(code)
	if err != nil {
		http.Error(w, "Ошибка при получении токенов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем информацию о пользователе
	userInfo, err := getUserInfo(tokenResp.AccessToken)
	if err != nil {
		http.Error(w, "Ошибка при получении информации о пользователе: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем сессию
	sessionID := generateSessionID()
	sessionStore.Set(sessionID, userInfo)

	// Устанавливаем куку
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true, // Для HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	// Перенаправляем на страницу профиля
	http.Redirect(w, r, "/userinfo", http.StatusFound)
}

// Новый обработчик для получения данных пользователя
func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
		return
	}

	userInfo, exists := sessionStore.Get(cookie.Value)
	if !exists {
		http.Error(w, "Сессия устарела или не существует", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

// Генератор ID сессии
func generateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Исправленная функция получения информации о пользователе
func getUserInfo(accessToken string) (UserInfo, error) {
	req, err := http.NewRequest("GET", UserInfoURL, nil)
	if err != nil {
		return UserInfo{}, fmt.Errorf("ошибка при создании запроса: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return UserInfo{}, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, fmt.Errorf("ошибка при чтении тела ответа: %w", err)
	}

	fmt.Println("Тело ответа:", string(bodyBytes))

	var userInfo UserInfo
	if err := json.Unmarshal(bodyBytes, &userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("ошибка при декодировании информации о пользователе: %w", err)
	}

	return userInfo, nil
}
