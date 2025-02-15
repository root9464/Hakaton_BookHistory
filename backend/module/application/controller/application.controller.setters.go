package application_controller

import (
	"github.com/gofiber/fiber/v2"
	application_dto "github.com/root9464/Hakaton_Zalupa/module/application/dto"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (c *ApplicationController) Create(ctx *fiber.Ctx) error {
	dto := new(application_dto.CreateApplicationDto)
	if err := ctx.BodyParser(dto); err != nil {
		c.logger.Errorf("error parsing body: %v", err)
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to parse multipart form"})
	}

	files := file_dto.CreateManyFileDto{
		Files: form.File["files"],
	}

	if err := c.applicationServ.Create(ctx.Context(), dto, &files); err != nil {
		c.logger.Errorf("error creating application: %v", err)
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application created successfully",
	})
}

func (c *ApplicationController) UpdateStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	status := new(application_dto.UpdateStatus)
	if err := ctx.BodyParser(status); err != nil {
		c.logger.Errorf("error parsing body: %v", err)
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.applicationServ.UpdateStatus(ctx.Context(), id, status); err != nil {
		c.logger.Errorf("error updating application: %v", err)
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application updated successfully",
	})
}


func (c *ApplicationController) SendEmail(ctx *fiber.Ctx) error {
	email := new(application_dto.Email)
	if err := ctx.BodyParser(email); err != nil {
		c.logger.Errorf("error parsing body: %v", err)
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.applicationServ.SendEmail(ctx.Context(), email); err != nil {
		c.logger.Errorf("error sending email: %v", err)
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)	
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Email sent successfully",
	})
}