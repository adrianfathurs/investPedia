package main

import (
	"investPedia/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/investpedia2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.FullName = "setyawan"
	userInput.Email = "coba@gmail.com"
	userInput.Occupation = "Backend"
	userInput.Password = "hell"
	userService.RegisterUser(userInput)

}

// func handler(c *gin.Context) {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "root:@tcp(127.0.0.1:3306)/investpedia2?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var users []user.User
// 	db.Find(&users)
// 	fmt.Println(users)
// 	c.JSON(http.StatusOK, users)
// }
