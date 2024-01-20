package config

import (
	"fmt"
	"log"
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

	host := fmt.Sprintf("\x1b[1;34mhttp://%s/kredit-plus/customer/swagger/index.html\x1b[0m\n", cfg.ServiceHost)
	log.Println(host)

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
