package api_service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
)

func (s *apiService) GetRecord(ctx context.Context) ([]api_dto.Feature, error) {
	req, err := http.NewRequest("GET", s.config.EXTERNAL_API, nil)
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

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

	var features []api_dto.Feature
	err = json.Unmarshal(body, &features)
	if err != nil {
		s.logger.Infof("Ошибка при десериализации ответа: %s", err)
		return nil, err
	}

	return features, nil
}

func (s *apiService) GetRecordByID(ctx context.Context, id string) (*api_dto.Feature, error) {
	req, err := http.NewRequest("GET", s.config.EXTERNAL_API, nil)
	if err != nil {
		s.logger.Infof("Ошибка при создании запроса: %s", err)
		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

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

	// Десериализация JSON-ответа в массив объектов
	var features []api_dto.Feature
	err = json.Unmarshal(body, &features)
	if err != nil {
		s.logger.Infof("Ошибка при десериализации ответа: %s", err)
		return nil, err
	}

	// Поиск нужной записи по ID
	for _, record := range features {
		if strconv.Itoa(record.ID) == id {
			return &record, nil
		}
	}

	// Если запись не найдена
	return nil, &fiber.Error{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf("Запись с ID %s не найдена", id),
	}
}
