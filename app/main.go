package main

import (
	"log"
	"time"

	"github.com/iskandardevan/book-library/app/middlewares"
	"github.com/iskandardevan/book-library/app/routes"

	authorUseCase "github.com/iskandardevan/book-library/business/authors"
	bookUseCase "github.com/iskandardevan/book-library/business/books"
	publisherUseCase "github.com/iskandardevan/book-library/business/publishers"
	reservationUseCase "github.com/iskandardevan/book-library/business/reservations"
	userUseCase "github.com/iskandardevan/book-library/business/users"

	authorController "github.com/iskandardevan/book-library/controllers/authors"
	bookController "github.com/iskandardevan/book-library/controllers/books"
	publisherController "github.com/iskandardevan/book-library/controllers/publishers"
	reservationController "github.com/iskandardevan/book-library/controllers/reservations"
	userController "github.com/iskandardevan/book-library/controllers/users"

	authorRepo "github.com/iskandardevan/book-library/driver/repository/authors"
	bookRepo "github.com/iskandardevan/book-library/driver/repository/books"
	publisherRepo "github.com/iskandardevan/book-library/driver/repository/publishers"
	reservationRepo "github.com/iskandardevan/book-library/driver/repository/reservations"
	userRepo "github.com/iskandardevan/book-library/driver/repository/users"

	"github.com/iskandardevan/book-library/driver/mysql"
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
	DB.AutoMigrate(&authorRepo.Author{})
	DB.AutoMigrate(&bookRepo.Book{})
	DB.AutoMigrate(&publisherRepo.Publisher{})
	DB.AutoMigrate(&reservationRepo.Reservation{})
}

func main(){
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}
	
	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &middlewares.ConfigJWT{})
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	authorRepoInterface := authorRepo.NewAuthorRepo(DB)
	authorUseCaseInterface := authorUseCase.NewUseCase(authorRepoInterface, timeoutContext)
	authorUseControllerInterface := authorController.NewAuthorController(authorUseCaseInterface)

	publisherRepoInterface := publisherRepo.NewPublisherRepo(DB)
	publisherUseCaseInterface := publisherUseCase.NewUseCase(publisherRepoInterface , timeoutContext)
	publisherUseControllerInterface := publisherController.NewPublisherController(publisherUseCaseInterface)

	bookRepoInterface := bookRepo.NewBookRepo(DB)
	bookUseCaseInterface := bookUseCase.NewUseCase(bookRepoInterface , timeoutContext )
	bookUseControllerInterface := bookController.NewBookController(bookUseCaseInterface)

	reservationRepoInterface := reservationRepo.NewReservationRepo(DB)
	reservationUseCaseInterface := reservationUseCase.NewUseCase(reservationRepoInterface , timeoutContext )
	reservationUseControllerInterface := reservationController.NewReservationController(reservationUseCaseInterface)


	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
		AuthorController: *authorUseControllerInterface,
		PublisherController: *publisherUseControllerInterface,
		BookController: *bookUseControllerInterface,
		ReservationController: *reservationUseControllerInterface,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}

