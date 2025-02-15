package api_service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
)

func (s *apiService) CreateRecord(ctx context.Context, dto *api_dto.CreateFeatureRequest) (*api_dto.CreateFeatureResponse, error) {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		s.logger.Infof("Ошибка при сериализации JSON: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Отправляемый JSON: %s", string(jsonData))

	req, err := http.NewRequest("POST", s.config.EXTERNAL_API, bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем заголовки
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", authString())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Infof("Ошибка при выполнении запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
		return nil, err
	}
	s.logger.Infof("Статус ответа: %s", resp.Status)
	s.logger.Infof("Тело ответа: %s", string(body))

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	response := new(api_dto.CreateFeatureResponse)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&response)
	if err != nil {
		s.logger.Infof("Ошибка при десериализации ответа: %s", err)
		return nil, err
	}

	return response, nil
}	
func (s *apiService) DeleteRecord(ctx context.Context, ids []api_dto.CreateFeatureResponse) error {
	// Сериализуем массив ID в JSON
	jsonData, err := json.Marshal(ids)
	if err != nil {
		s.logger.Infof("Ошибка при сериализации JSON: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Отправляемый JSON: %s", string(jsonData))

	// Формируем URL для удаления записей
	url := s.config.EXTERNAL_API // Например, https://geois2.orb.ru/api/resource/8563/feature/

	// Создаем HTTP-запрос
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем заголовки
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", authString())

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Infof("Ошибка при выполнении запроса: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
		return err
	}

	s.logger.Infof("Статус ответа: %s", resp.Status)
	s.logger.Infof("Тело ответа: %s", string(body))

	// Проверяем статус ответа
	if resp.StatusCode >= 400 {
		return fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	return nil
}
func (s *apiService) UpdateRecord(ctx context.Context, id string, dto *api_dto.CreateFeatureRequest) error {
	// Сериализуем DTO в JSON
	jsonData, err := json.Marshal(dto)
	if err != nil {
		s.logger.Infof("Ошибка при сериализации JSON: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Отправляемый JSON: %s", string(jsonData))

	// Формируем URL для изменения записи
	url := s.config.EXTERNAL_API + id
	s.logger.Infof("URL: %s", url)

	// Создаем HTTP-запрос
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем заголовки
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", authString())

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Infof("Ошибка при выполнении запроса: %s", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
		return err
	}

	s.logger.Infof("Статус ответа: %s", resp.Status)
	s.logger.Infof("Тело ответа: %s", string(body))

	// Проверяем статус ответа
	if resp.StatusCode >= 400 {
		return fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	return nil
}

func (s *apiService) UploadAttachment(ctx context.Context, file *multipart.FileHeader, fileName string) ([]map[string]interface{}, error) {
	// Открываем файл
	src, err := file.Open()
	if err != nil {
		s.logger.Infof("Ошибка при открытии файла: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer src.Close()

	// Создаем буфер для хранения данных файла
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл в форму
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		s.logger.Infof("Ошибка при создании формы для файла: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Копируем содержимое файла в часть формы
	if _, err := io.Copy(part, src); err != nil {
		s.logger.Infof("Ошибка при копировании файла: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем поле "name" (если нужно)
	if err := writer.WriteField("name", fileName); err != nil {
		s.logger.Infof("Ошибка при добавлении поля name: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Закрываем writer
	if err := writer.Close(); err != nil {
		s.logger.Infof("Ошибка при закрытии writer: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Формируем URL для загрузки файла
	url := "https://geois2.orb.ru/api/component/file_upload/"

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем заголовки
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", authString())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Infof("Ошибка при выполнении запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
		return nil, err
	}

	s.logger.Infof("Статус ответа: %s", resp.Status)
	s.logger.Infof("Тело ответа: %s", string(responseBody))

	// Проверяем статус ответа
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(responseBody))
	}

	// Десериализуем ответ
	var result map[string][]map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		s.logger.Infof("Ошибка при десериализации ответа: %s", err)
		return nil, err
	}

	return result["upload_meta"], nil
}
func (s *apiService) AttachingFile(ctx context.Context, recordID string, attachmentRequest api_dto.AttachmentRequest) (int, error) {
	// Сериализуем запрос в JSON
	jsonData, err := json.Marshal(attachmentRequest)
	if err != nil {
		s.logger.Infof("Ошибка при сериализации JSON: %s", err)
		return 0, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Отправляемый JSON: %s", string(jsonData))

	// Формируем URL для прикрепления файла
	url := s.config.EXTERNAL_API + recordID + "/attachment/"
	s.logger.Infof("URL: %s", url)

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return 0, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Добавляем заголовки
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", authString())
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Infof("Ошибка при выполнении запроса: %s", err)
		return 0, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
		return 0, err
	}

	s.logger.Infof("Статус ответа: %s", resp.Status)
	s.logger.Infof("Тело ответа: %s", string(body))

	// Проверяем статус ответа
	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
	}

	// Десериализуем ответ
	var response struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		s.logger.Infof("Ошибка при десериализации ответа: %s", err)
		return 0, err
	}

	return response.ID, nil
}
func (s *apiService) DeleteAttachment(ctx context.Context, recordID string, attachmentID string) error {
    // Формируем URL для удаления вложения
    url := s.config.EXTERNAL_API + recordID +"/attachment/"+ attachmentID

    // Создаем HTTP-запрос
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        s.logger.Infof("Ошибка при создании запроса: %s", err)
        return &fiber.Error{
            Code:    fiber.StatusInternalServerError,
            Message: err.Error(),
        }
    }

    // Добавляем заголовки
    req.Header.Set("Accept", "*/*")
    req.Header.Set("Authorization", authString())

    // Выполняем запрос
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        s.logger.Infof("Ошибка при выполнении запроса: %s", err)
        return &fiber.Error{
            Code:    fiber.StatusInternalServerError,
            Message: err.Error(),
        }
    }
    defer resp.Body.Close()

    // Читаем тело ответа
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        s.logger.Infof("Ошибка при чтении тела ответа: %s", err)
        return err
    }

    s.logger.Infof("Статус ответа: %s", resp.Status)
    s.logger.Infof("Тело ответа: %s", string(body))

    // Проверяем статус ответа
    if resp.StatusCode >= 400 {
        return fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(body))
    }

    return nil
}

func authString() string {
	auth := "hackathon_15:hackathon_15_25"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	return basicAuth
}
