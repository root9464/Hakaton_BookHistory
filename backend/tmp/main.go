package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ClientID     = "33"                                                                                                                            // Ваш ID в ЕЛК
	ClientSecret = "J6fs2LlWwslyTq56jeuXfefXeoG3Mrh92PcSYe2z"                                                                                      // Ваш секретный ключ
	RedirectURI  = "http://hackathon-5.orb.ru/profile/rsaag"                                                                                       // Адрес редиректа
	AuthURL      = "https://lk.orb.ru/oauth/authorize"                                                                                             // URL авторизации
	TokenURL     = "https://lk.orb.ru/oauth/token"                                                                                                 // URL для получения токенов
	UserInfoURL  = "https://lk.orb.ru/api/get_user?scope=rsaag_id+personal_data+esia_data+email+phone+esia_user_id+organizations_user+auth_method" // URL для получения информации о пользователе
)

// Структура для ответа с токенами
type TokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Структура для информации о пользователе
type UserInfo struct {
	ID    string `json:"rsaag_id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	// Добавьте другие поля, которые вы хотите получить
}

func main() {
	// Запуск HTTP-сервера
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/profile/rsaag", handleCallback)

	fmt.Println("Сервер запущен на http://hackathon-5.orb.ru")
	http.ListenAndServe(":8080", nil)
}

// Обработчик для главной страницы
func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
}

// Обработчик для начала авторизации
func handleLogin(w http.ResponseWriter, r *http.Request) {
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=email+auth_method&state=http://hackathon-5.orb.ru/", AuthURL, ClientID, RedirectURI)
	http.Redirect(w, r, authURL, http.StatusFound)
}

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

	// Выводим информацию о пользователе
	fmt.Fprintf(w, "Информация о пользователе: %+v", userInfo)
}

// Функция для получения токенов
func getTokens(code string) (*TokenResponse, error) {
	payload := url.Values{}
	payload.Set("client_id", ClientID)
	payload.Set("client_secret", ClientSecret)
	payload.Set("redirect_uri", RedirectURI)
	payload.Set("code", code)
	payload.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(TokenURL, payload)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе токенов: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
	}

	return &tokenResp, nil
}

// Функция для получения информации о пользователе
func getUserInfo(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", UserInfoURL, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании запроса: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении тела ответа: %w", err)
	}

	// Вывод тела ответа в консоль
	fmt.Println("Тело ответа:", string(bodyBytes))

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", fmt.Errorf("ошибка при декодировании информации о пользователе: %w", err)
	}

	return string(bodyBytes), nil
}
