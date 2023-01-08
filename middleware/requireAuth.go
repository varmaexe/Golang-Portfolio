package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vermaexe/go-jwt/initializers"
	"github.com/vermaexe/go-jwt/models"
)

func RequireAuth(c *gin.Context) {
	// fmt.Println("in middleware")
	//get the cookie from req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.HTML(200, "loginpage.html", gin.H{})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode and validate
	// Parse takes the token string and a function for looking up the key. The latter is especially
	//token, err :=
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find the user with sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(200, gin.H{
				"message": "user doesnt exit",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		c.Set("user", user)
		//
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

}
