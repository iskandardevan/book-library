package controllers

import

type UserController struct {
	usecase users.UserUsecaseInterface
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

func NewUserController(uc users.UserUsecaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (ctrl *Controller) CreateUserController(c echo.Context) error {
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