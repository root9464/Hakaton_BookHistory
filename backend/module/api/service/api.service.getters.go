package api_service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
)

type UserResponse struct {
	UserID string `json:"userID"`
	Coords struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"cords"`
}

func (s *apiService) GetRecords(ctx context.Context) ([]api_dto.FeatureCord, error) {
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

	var results []api_dto.FeatureCord

	for _, feature := range features {
		if feature.Geom != "" {
			convertedGeom, err := convertEPSG3857to4326(feature.Geom)
			if err != nil {
				s.logger.Infof("Ошибка при конвертации geom: %s", err)
				return nil, err
			}

			var lon, lat float64
			convertedGeom = strings.TrimPrefix(convertedGeom, "POINT(")
			convertedGeom = strings.TrimSuffix(convertedGeom, ")")
			parts := strings.Split(convertedGeom, " ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("неверный формат координат: %s", convertedGeom)
			}

			lon, err = strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка при парсинге долготы: %s", err)
			}

			lat, err = strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка при парсинге широты: %s", err)
			}

			featureCord := api_dto.FeatureCord{
				ID: feature.ID,
				Coords: struct {
					Lat float64 `json:"lat"`
					Lon float64 `json:"lon"`
				}{
					Lat: lat,
					Lon: lon,
				},
				Fields:     feature.Fields,
				Extensions: feature.Extensions,
			}
			featureCord.ID = 0
			//uuid
			results = append(results, featureCord)
		}
	}

	return results, nil
}

func (s *apiService) GetRecordByID(ctx context.Context, id string) (map[string]interface{}, error) {
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

	for _, record := range features {
		if strconv.Itoa(record.ID) == id {
			// Конвертация координат
			coords, err := convertEPSG3857to4326(record.Geom)
			if err != nil {
				return nil, err
			}

			// Разбираем координаты
			var lon, lat float64
			coords = strings.TrimPrefix(coords, "POINT(")
			coords = strings.TrimSuffix(coords, ")")
			parts := strings.Split(coords, " ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("неверный формат координат: %s", coords)
			}
			lon, err = strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка при парсинге долготы: %s", err)
			}
			lat, err = strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка при парсинге широты: %s", err)
			}

			
			record.ID=0
			//record.Fields.UUID=id
			return map[string]interface{}{
				"uuid": record.Fields.UUID,
				"cords": map[string]float64{
					"lat": lat,
					"lon": lon,
				},
			}, nil
		}
	}

	return nil, &fiber.Error{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf("Запись с ID %s не найдена", id),
	}
}


func (s *apiService) GetFullRecordByID(ctx context.Context, id string) (*api_dto.Feature, error) {
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
			record.ID = 0
			//record.Fields.UUID = id
			return &record, nil
		}
	}
	// Если запись не найдена
	return nil, &fiber.Error{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf("Запись с ID %s не найдена", id),
	}

	
}

func convertEPSG3857to4326(geom string) (string, error) {
	// Удаляем "POINT(" и ")" из строки
	geom = strings.TrimPrefix(geom, "POINT(")
	geom = strings.TrimSuffix(geom, ")")

	// Разделяем координаты по пробелу
	coords := strings.Split(geom, " ")
	if len(coords) != 2 {
		return "", fmt.Errorf("неверный формат координат: %s", geom)
	}

	// Парсим координаты
	x, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга координаты X: %v", err)
	}
	y, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга координаты Y: %v", err)
	}

	// Преобразуем координаты из EPSG:3857 в EPSG:4326
	const earthRadius = 6378137.0 // Радиус Земли в метрах
	lon := (x / earthRadius) * (180 / math.Pi)
	lat := (math.Atan(math.Exp(y/earthRadius))*2 - math.Pi/2) * (180 / math.Pi)

	// Возвращаем координаты в формате "POINT(lon lat)"
	return fmt.Sprintf("POINT(%f %f)", lon, lat), nil
}

func (s *apiService) GetArtInfo(ctx context.Context, id string) (*api_dto.Feature, error) {
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