package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB

func init() {
	//open a db connection
	//mysql = database driver
	//root = database user name
	//12345 = database password
	//demo = database name
	var err error
	//db, err = gorm.Open("mysql", "root:12345@/demo?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "go_user_name:go_password@/go_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	//Migrate the schema
	db.AutoMigrate(&registerModel{}, &getProfile{})

}

func main() {
	handleRequests()
}

func handleRequests() {
	router := gin.Default()
	auth := router.Group("/api/v1/auth/")
	{
		auth.GET("login", login)
		auth.POST("register", register)
		auth.GET("profile", profile)
	}
	//router.GET("/profile/:name", profile)
	router.Run(":8081")
}

type (
	registerModel struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	getProfile struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func profile(context *gin.Context) {

	var userProfile []getProfile
	//var _userProfile [] transformedUserProfile

	db.Find(userProfile)

	if len(userProfile) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Profile Not found!!"})
	}

}

func login(context *gin.Context) {
	context.String(http.StatusOK, "we'll unmarshal the json")
}

func register(context *gin.Context) {
	doRegister := registerModel{Name: context.PostForm("name"), Email: context.PostForm("email")}
	db.Save(&doRegister)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Registered Successfully!", "resourceId": doRegister.ID})

}
