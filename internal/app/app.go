package app

import (
	"awesomeProject/config"
	auth_usecase "awesomeProject/internal/domain/auth/usecase"
	"awesomeProject/internal/domain/member/usecase"
	v1 "awesomeProject/internal/infrastructure/controller/http/v1"
	member_repository "awesomeProject/internal/infrastructure/repository/member"
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

	// Инициализация пакета JWT
	jwtAuth := jwt_auth.NewJwtAuth(app.jwtAuthKey)

	// Инициализация UseCase
	memberUC := member_usecase.NewMemberUseCase(memberRepo)

	useCaseList := v1.UC{
		MemberUC: memberUC,
		AuthUC:   auth_usecase.NewAuthUseCase(memberUC, jwtAuth),
	}

	// Инициализация Router
	v1.NewRouter(app.handler, useCaseList, jwtAuth)
	// Запуск сервера
	if err := app.httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
