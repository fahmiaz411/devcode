package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/fahmiaz411/devcode/config/database"
	"github.com/fahmiaz411/devcode/helper/constant"
	_activityHandler "github.com/fahmiaz411/devcode/modules/activity/delivery"
	_activityRepo "github.com/fahmiaz411/devcode/modules/activity/repository"
	_activityUsecase "github.com/fahmiaz411/devcode/modules/activity/usecase"

	_todoHandler "github.com/fahmiaz411/devcode/modules/todo/delivery"
	_todoRepo "github.com/fahmiaz411/devcode/modules/todo/repository"
	_todoUsecase "github.com/fahmiaz411/devcode/modules/todo/usecase"
)

func main() {
	app := fiber.New()

	mysqlPort, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if mysqlPort == constant.ZeroValue {
		mysqlPort = 3306
	}
	
	db := database.NewMysqlDB(database.MysqlConfig{
		DatabaseName: os.Getenv("MYSQL_DBNAME"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host: os.Getenv("MYSQL_HOST"),
		Port: mysqlPort,
	})

	// Dev

	// db := database.NewMysqlDB(database.MysqlConfig{
	// 	DatabaseName: "devcode",
	// 	Username: "root",
	// 	Password: "1234",
	// 	Host: "localhost",
	// 	Port: 3306,
	// })

	timeout := time.Duration(1 * time.Minute)

	activityRepo := _activityRepo.NewRepository(db)
	activityUsecase := _activityUsecase.NewUsecase(activityRepo, timeout)
	_activityHandler.NewRESTHandler(app, activityUsecase)

	todoRepo := _todoRepo.NewRepository(db)
	todoUsecase := _todoUsecase.NewUsecase(todoRepo, timeout)
	_todoHandler.NewRESTHandler(app, todoUsecase)

	app.Listen(":3030")
}