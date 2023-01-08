package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vermaexe/go-jwt/initializers"
	"github.com/vermaexe/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//gets the email/pass frm req body
	var body struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	// body.Email = c.PostForm("email")
	// body.Password = c.PostForm("password")
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
			// "error": "failed to create user",
		})

		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//get email/pass from body
	// c.HTML(200, "loginpage.html", gin.H{})
	var body struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//look up reqested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//compare the password with hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//Generate jwt-token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	//sign and get the encoded token as string using secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}
	//send token back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true) //set secure bool false if app is running in localhost

	// c.JSON(http.StatusOK, gin.H{})
	c.HTML(200, "portfolio.html", gin.H{})

}
func LogOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged out",
	})
}

func Validate(c *gin.Context) {
	// user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}
