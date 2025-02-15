package file_service

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	file_dto "github.com/root9464/Hakaton_BookHistory/module/file/dto"
)

func (s *FileService) CreateMany(ctx context.Context, dto *file_dto.CreateManyFileDto) ([]string, error) {
	var files []string

	// Обработка каждого файла
	for _, fileHeader := range dto.Files {
		// Генерация уникального имени файла
		lastDot := strings.LastIndex(fileHeader.Filename, ".")
		if lastDot == -1 {
			s.logger.Warnf("file %s has no extension", fileHeader.Filename)
			return nil, &fiber.Error{
				Code:    400,
				Message: "file has no extension",
			}
		}
		name := uuid.New().String() + fileHeader.Filename[lastDot:]
		fullPath := "../image/" + name

		// Сохранение файла на диск
		file, err := fileHeader.Open()
		if err != nil {
			s.logger.Errorf("failed to open file: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: "failed to process file",
			}
		}
		defer file.Close()

		outFile, err := os.Create(fullPath)
		if err != nil {
			s.logger.Errorf("failed to create file on disk: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: "failed to save file",
			}
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			s.logger.Errorf("failed to copy file content: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: "failed to save file",
			}
		}

		// Создание записи о файле
		files = append(files, name)
	}

	// Сохранение записей в репозитории
	// if err := s.repo.CreateMany(ctx, files); err != nil {
	// 	s.logger.Errorf("failed to save files in repository: %s", err.Error())
	// 	return &fiber.Error{
	// 		Code:    500,
	// 		Message: "failed to save files",
	// 	}
	// }

	s.logger.Info("files successfully saved")
	return files, nil
}
