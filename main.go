package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id 		uint 	`json:"id" form:"id" `
	Email	string 	`json:"email" form:"email" `
	FirstName 	string 	`json:firstname" form:"firstname" `
	LastName 	string 	`json:"lastname" form:"lastname" `
	Age		int	`json:"age" form:"age" `
	Phone	string 	`json:"phonenum" form:"phonenum" `

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
	e.GET("/users", c.GetUserController)
	e.POST("/user", c.CreateUser)

	e.Start(":8080")
}

func(ctrl *Controller) GetUserController(c echo.Context) error {
	var users []User

	err := ctrl.DB.Find(&users).Error
	fmt.Println(err)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" 	: "success",
		"data"		: users,
	})
}


func (ctrl *Controller) CreateUser(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	

	err := ctrl.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : err.Error(),
		})
	}

	return c.JSON(http.StatusOK,map[string]interface{}{
		"message": "success create user",
		"user" : user,
	})
}





