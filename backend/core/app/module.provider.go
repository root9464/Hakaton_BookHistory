package app

import (
	//jwt_module "github.com/root9464/Hakaton_Zalupa/module/jwt"
)

type moduleProvider struct {
	// userModule          *user_module.UserModule

	//jwtModule *jwt_module.JwtModule
	app       *App
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
		// p.UserModule,
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

// func (p *moduleProvider) UserModule() error {
// 	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator, p.app.db, *p.jwtModule, p.app.config.JwtPublicKey)
// 	return nil
// }
