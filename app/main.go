package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id 			uint 		`json:"id" form:"id" `
	Email		string 		`json:"email" form:"email" `
	Name 		string 		`json:"name" form:"name" `
	Age			int			`json:"age" form:"age" `
	Phone		string 		`json:"phonenum" form:"phonenum" `
	CreatedAt 	time.Time   `json:"createdAt"`
	UpdatedAt 	time.Time   `json:"updatedAt"`
}

func initDB() *gorm.DB {

	dsn := "root:190901@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
	panic(err.Error())
	}
	DB.AutoMigrate(&User{})
	return DB
}
type Controller struct {
	DB *gorm.DB
}

func main() {
	DB := initDB()
	c := Controller{DB}
	e := echo.New()

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUsecase(userRepoInterface, timeoutContext)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userControllerInterface,

	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}




