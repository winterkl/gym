package app

import (
	"awesomeProject/config"
	auth_usecase "awesomeProject/internal/domain/auth/usecase"
	member_usecase "awesomeProject/internal/domain/member/usecase"
	service_usecase "awesomeProject/internal/domain/service/usecase"
	subscription_usecase "awesomeProject/internal/domain/subscription/usecase"
	trainer_usecase "awesomeProject/internal/domain/trainer/usecase"
	v1 "awesomeProject/internal/infrastructure/controller/http/v1"
	"awesomeProject/internal/infrastructure/repository/member"
	service_repository "awesomeProject/internal/infrastructure/repository/service"
	subscription_repository "awesomeProject/internal/infrastructure/repository/subscription"
	"awesomeProject/internal/infrastructure/repository/trainer"
	"awesomeProject/pkg/http_server"
	"awesomeProject/pkg/jwt_auth"
	"awesomeProject/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	db         *postgres.Postgres
	handler    *gin.Engine
	httpServer *http_server.Server
	jwtAuthKey string
}

func NewApp(cfg *config.Config) App {
	psql, err := postgres.New(cfg.Pg.User, cfg.Pg.Password, cfg.Pg.Host, cfg.Pg.Port, cfg.Pg.DbName, cfg.Pg.SslMode)
	if err != nil {
		log.Fatal(err)
	}

	handler := gin.Default()
	server := http_server.NewHttpServer(cfg.Http.Host, cfg.Http.Port, handler)

	return App{
		db:         psql,
		handler:    handler,
		httpServer: server,
		jwtAuthKey: cfg.Auth.Key,
	}
}

func (app *App) Run() {

	// Инициализация Repository
	memberRepo := member_repository.NewMemberRepository(app.db)
	trainerRepo := trainer_repository.NewTrainerRepository(app.db)
	subscriptionRepo := subscription_repository.NewSubscriptionRepository(app.db)
	serviceRepo := service_repository.NewServiceRepository(app.db)

	// Инициализация пакета JWT
	jwtAuth := jwt_auth.NewJwtAuth(app.jwtAuthKey)

	// Инициализация UseCase
	memberUC := member_usecase.NewMemberUseCase(memberRepo)

	useCaseList := v1.UC{
		MemberUC:       memberUC,
		TrainerUC:      trainer_usecase.NewTrainerUseCase(trainerRepo, memberRepo),
		AuthUC:         auth_usecase.NewAuthUseCase(memberUC, jwtAuth),
		SubscriptionUC: subscription_usecase.NewSubscriptionUseCase(subscriptionRepo),
		ServiceUC:      service_usecase.NewServiceUseCase(serviceRepo),
	}

	// Инициализация Router
	v1.NewRouter(app.handler, useCaseList, jwtAuth)
	// Запуск сервера
	if err := app.httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
