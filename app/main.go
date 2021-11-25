package main

import (
	"log"
	"time"

	"github.com/iskandardevan/book-library/app/routes"
	_userusecase "github.com/iskandardevan/book-library/business/users"
	_usercontroller "github.com/iskandardevan/book-library/controllers/users"
	"github.com/iskandardevan/book-library/driver/mysql"
	_userrepo "github.com/iskandardevan/book-library/driver/repository/users"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	// _userusecase "github.com/iskandardevan/book-library/business/users"
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
	DB.AutoMigrate(&_userrepo.User{})
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

	
	UserRepoInterface := _userrepo.NewUserRepo(DB)
	userUseCaseInterface := _userusecase.NewUseCase(UserRepoInterface, timeoutContext)
	userUseControllerInterface := _usercontroller.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}

