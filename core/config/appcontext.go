package config

import (
	"test-kreditplus/swagger"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type appContext struct {
	db               *gorm.DB
	cfg              *Configuration
	requestValidator *validator.Validate
	// tracer           trace.Tracer
}

var appCtx appContext

func Init() error {
	// logger.Init(ctx)

	cfg, err := InitConfig()
	if err != nil {
		return err
	}

	db, err := ConnectDataBase(cfg.Postgres)
	if err != nil {
		return err
	}

	swagger.SwaggerInfo.Host = cfg.Swagger.SwaggerHost
	swagger.SwaggerInfo.Schemes = []string{"http", "https"}

	// _, err = tracer.Init(ctx, cfg.ServiceName, "0.0.1")
	// if err != nil {
	// 	return err
	// }

	appCtx = appContext{
		db:  db,
		cfg: cfg,
		// tracer: otel.Tracer(cfg.ServiceName),
	}

	return nil
}

func RequestValidator() *validator.Validate {
	return appCtx.requestValidator
}

func DB() *gorm.DB {
	return appCtx.db
}

func Config() Configuration {
	return *appCtx.cfg
}

// func Tracer() trace.Tracer {
// 	return appCtx.tracer
// }
