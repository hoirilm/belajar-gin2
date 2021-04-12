package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

type User struct {
	ID string `json: "id"`
	Name string `json: "name"'`
	Age string `json: "age"`
}

// slice
var Users[]User

func main() {
	r := gin.Default()

	// grouping routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUsers)
		userRoutes.PUT("/:id", EditUsers) // users/[id disini]
		userRoutes.DELETE("/:id", DeleteUsers)
	}

	// penulisan singkat untuk error sekaligus runnya
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(200, Users) // []
}

func CreateUsers(c *gin.Context) {
	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"status": "fail",
			"detail": "invalid request body",
		})

		return // untuk menghentikan proses
	}

	// set id dinamis menggunakan package UUID
	reqBody.ID = uuid.New().String()

	// append ke data Users
	Users = append(Users, reqBody)

	c.JSON(200, gin.H{
		"status": "success",
	})
}

func EditUsers(c *gin.Context) {
	id := c.Param("id") // disesuaikan dengan parameter di EditUser

	var reqBody User

	// cek data yang dikirim
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"status":"fail",
			"info":"invalid request body",
		})
		return
	}

	// compare data untuk diganti dengan kiriman sebelumnya
	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"status":"success",
			})

			return
		}
	}

	// kondisi jika id yang dicari tidak ada
	c.JSON(404, gin.H{
		"status":"error",
		"info":"invalid user id",
	})
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			// t  := [1, 431, 44, 546, 61]
			// t[:2] == [1,431]
			// t[2 + 1:] == [546,61]
			// [1,431,546,61]
			Users = append(Users[:i], Users[i + 1:]...)

			c.JSON(200, gin.H{
				"status":"success",
			})

			return
		}
	}

	c.JSON(404, gin.H{
		"status": "fail",
		"info": "invalid user id",
	})
}




//NOTE: ada istilahnya pass by reference dan pass by value
// digolang normalnya menggunakan pass by value, jadi semua data di duplikasi