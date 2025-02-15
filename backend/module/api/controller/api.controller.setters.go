package api_controller

import (
	"github.com/gofiber/fiber/v2"
	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
)

func (c *ApiController) CreateRecord(ctx *fiber.Ctx) error {
    dto := new(api_dto.CreateFeatureRequest)
    if err := ctx.BodyParser(dto); err != nil {
        c.logger.Info("01")
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    c.logger.Infof("CreateRecord: %+v", dto)

    id, err := c.apiService.CreateRecord(ctx.Context(), dto)
    if err != nil {
        c.logger.Info("02")
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status":  "success",
        "message": "record created successfully",
        "data": fiber.Map{
            "id": id,
        },
    })
}
func (c *ApiController) UpadateRecord(ctx *fiber.Ctx) error {
    // Получаем ID записи из параметров запроса
    id := ctx.Params("id")
    if id == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": "ID записи не указан",
        })
    }

    // Парсим тело запроса в DTO
    dto := new(api_dto.CreateFeatureRequest)
    if err := ctx.BodyParser(dto); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    c.logger.Infof("ChangeRecord: ID=%s, DTO=%+v", id, dto)

    // Вызываем метод сервиса для изменения записи
    err := c.apiService.UpdateRecord(ctx.Context(), id, dto)
    if err != nil {
        c.logger.Info("03")
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    // Возвращаем успешный ответ
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "status":  "success",
        "message": "record updated successfully",
    })
}
func (c *ApiController) DeleteRecord(ctx *fiber.Ctx) error {
    // Парсим тело запроса в массив ID
    var ids []api_dto.CreateFeatureResponse
    if err := ctx.BodyParser(&ids); err != nil {
        c.logger.Info("01")
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    c.logger.Infof("DeleteRecord: IDs=%+v", ids)

    // Вызываем метод сервиса для удаления записей
    err := c.apiService.DeleteRecord(ctx.Context(), ids)
    if err != nil {
        c.logger.Info("02")
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    // Возвращаем успешный ответ
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "status":  "success",
        "message": "records deleted successfully",
    })
}

func (c *ApiController) UploadAttachment(ctx *fiber.Ctx) error {
    // Получаем файл из запроса
    file, err := ctx.FormFile("file")
    if err != nil {
        c.logger.Info("01")
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    // Получаем имя файла (если передано)
    fileName := ctx.FormValue("name")
    if fileName == "" {
        fileName = file.Filename // Используем имя файла по умолчанию
    }

    c.logger.Infof("UploadAttachment: FileName=%s, Size=%d", fileName, file.Size)

    // Вызываем метод сервиса для загрузки файла
    uploadMeta, err := c.apiService.UploadAttachment(ctx.Context(), file, fileName)
    if err != nil {
        c.logger.Info("02")
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "failed",
            "message": err.Error(),
        })
    }

    // Возвращаем успешный ответ
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "status":  "success",
        "message": "file uploaded successfully",
        "data":    uploadMeta,
    })
}