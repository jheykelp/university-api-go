package main

import (
	"log"
	"os"

	_ "github.com/p2-graded-challenge-2-jspheykel/docs" // swagger generated files

	"github.com/labstack/echo/v4"
	"github.com/p2-graded-challenge-2-jspheykel/config"
	"github.com/p2-graded-challenge-2-jspheykel/internal/handler"
	"github.com/p2-graded-challenge-2-jspheykel/internal/middleware"
	"github.com/p2-graded-challenge-2-jspheykel/internal/repository"
	"github.com/p2-graded-challenge-2-jspheykel/internal/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title University API
// @version 1.0
// @description Academic program API (students, courses, enrollments).
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	db := config.OpenDB()

	// repos & services
	stRepo := repository.NewStudentRepo(db)
	crRepo := repository.NewCourseRepo(db)
	enRepo := repository.NewEnrollmentRepo(db)

	authSvc := service.NewAuthService(stRepo, os.Getenv("JWT_SECRET"))
	crSvc := service.NewCourseService(crRepo)
	enSvc := service.NewEnrollmentService(crRepo, enRepo)

	// handlers
	stHandler := handler.NewStudentHandler(authSvc, stRepo)
	crHandler := handler.NewCourseHandler(crSvc)
	enHandler := handler.NewEnrollmentHandler(enSvc)

	e := echo.New()

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// public
	e.POST("/students/register", stHandler.Register)
	e.POST("/students/login", stHandler.Login)

	// protected
	api := e.Group("")
	api.Use(middleware.RequireJWT)
	api.GET("/students/me", stHandler.Me)
	api.GET("/courses", crHandler.List)
	api.POST("/enrollments", enHandler.Enroll)
	api.DELETE("/enrollments/:id", enHandler.Delete)

	addr := ":" + config.Port()
	log.Printf("listening on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
