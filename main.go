package main

import (
	"funding/handler"
	"funding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1 dbname=gostartup port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// newuserRepository, err := userRepository.FindByEmail("tikoh@gmail.com")

	// if newuserRepository.ID == 0 {
	// 	fmt.Println("User Tidak di temukan")
	// 	return
	// }
	// fmt.Println(newuserRepository.Name)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	router.Run()
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test Simpang dari Service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "CEO"
	// userInput.Password = "passwordajah"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "Test Simpan",
	// }

	// userRepository.Save(user)
}
