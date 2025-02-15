package app

import (
	application_module "github.com/root9464/Hakaton_BookHistory/module/application"
	auth_module "github.com/root9464/Hakaton_BookHistory/module/auth"
	file_module "github.com/root9464/Hakaton_BookHistory/module/file"
	jwt_module "github.com/root9464/Hakaton_BookHistory/module/jwt"
	reward_module "github.com/root9464/Hakaton_BookHistory/module/reward"
	user_module "github.com/root9464/Hakaton_BookHistory/module/user"
)

// jwt_module "github.com/root9464/Ton-students/module/jwt"

type moduleProvider struct {
	userModule        *user_module.UserModule
	authModule        *auth_module.AuthModule
	jwtModule         *jwt_module.JwtModule
	fileModule        *file_module.FileModule
	applicationModule *application_module.ApplicationModule
	rewardModule      *reward_module.RewardModule

	app *App
}

func NewModuleProvider(app *App) (*moduleProvider, error) {
	provider := &moduleProvider{
		app: app,
	}

	err := provider.initDeps()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *moduleProvider) initDeps() error {
	inits := []func() error{
		p.JwtModule,
		p.UserModule,
		p.AuthModule,
		p.FileModule,
		p.ApplicationModule,
		p.RewardModule,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			p.app.logger.Errorf("%s", "✖ Failed to initialize module: "+err.Error())
			return err
		}
	}
	return nil
}

func (p *moduleProvider) JwtModule() error {
	p.jwtModule = jwt_module.NewJwtModule(p.app.logger, p.app.validator, p.app.db, p.app.config.JwtPrivateKey, p.app.config.JwtPublicKey)
	return nil
}

func (p *moduleProvider) UserModule() error {
	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator, p.app.db, *p.jwtModule, p.app.config.JwtPublicKey)
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.config, p.userModule.UserService(), *p.jwtModule)
	return nil
}

func (p *moduleProvider) FileModule() error {
	p.fileModule = file_module.NewFileModule(p.app.logger, p.app.db)
	return nil
}

func (p *moduleProvider) ApplicationModule() error {
	p.applicationModule = application_module.NewApplicationModule(p.app.logger, p.app.db, p.fileModule.FileService())
	return nil
}

func (p *moduleProvider) RewardModule() error {
	p.rewardModule = reward_module.NewRewardModule(p.app.logger, p.app.validator, p.app.db, p.fileModule.FileService())
	return nil
}
