package main

import (
	"log"
	"time"

	"github.com/iskandardevan/book-library/app/middlewares"
	"github.com/iskandardevan/book-library/app/routes"
	userUseCase "github.com/iskandardevan/book-library/business/users"
	userController "github.com/iskandardevan/book-library/controllers/users"
	"github.com/iskandardevan/book-library/driver/mysql"
	userRepo "github.com/iskandardevan/book-library/driver/repository/users"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)



func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	if viper.GetBool(`debug`){
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
}

func main(){
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}
	// JWT := middlewares.ConfigJWT{
	// 	SecretJWT:       viper.GetString("jwt.secret"),
	// 	ExpiresDuration: viper.GetInt("jwt.expired"),
	// }
	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &middlewares.ConfigJWT{})
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}

