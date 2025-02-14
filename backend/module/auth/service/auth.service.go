package auth_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Hakaton_Zalupa/config"
	auth_dto "github.com/root9464/Hakaton_Zalupa/module/auth/dto"
	jwt_funcs "github.com/root9464/Hakaton_Zalupa/module/jwt/functions"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	user_service "github.com/root9464/Hakaton_Zalupa/module/user/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IAuthService = (*authService)(nil)

type IAuthService interface {
	Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*user_model.User, error)
	Register(ctx context.Context, dto *auth_dto.RegisterDto) error
}

type authService struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	userService user_service.IUserService
	jwtFuncs    jwt_funcs.IJwtFuncs
}

func NewAuthService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	userService user_service.IUserService,
	jwtFuncs jwt_funcs.IJwtFuncs,
) *authService {
	return &authService{
		logger:      logger,
		validator:   validator,
		config:      config,
		userService: userService,
		jwtFuncs:    jwtFuncs,
	}
}
