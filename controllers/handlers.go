package controllers

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func Loginpage(c *gin.Context) {
	c.HTML(200, "loginpage.html", gin.H{})
}
func Signinpage(c *gin.Context) {
	c.HTML(200, "SignUppage.html", gin.H{})
}

func Contactpage(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{})
}

func Portfolio(c *gin.Context) {
	c.HTML(200, "portfolio.html", gin.H{})
}
func Indexpage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func Skillspage(c *gin.Context) {
	c.HTML(200, "skills.html", gin.H{})
}
func WeatherHome(c *gin.Context) {
	c.HTML(200, "weatherhome.html", gin.H{})
}
func PostPage(c *gin.Context) {
	c.HTML(200, "posthomepage.html", gin.H{})
}
func PostsCreatePage(c *gin.Context) {
	c.HTML(200, "createposts.html", gin.H{})
}
func PostUpdatePage(c *gin.Context) {
	c.HTML(200, "updatepost.html", gin.H{})
}
func PostsDeletePage(c *gin.Context) {
	c.HTML(200, "deletepost.html", gin.H{})
}

type Msg struct {
	Name    string
	Email   string
	Message string
}

func Contactform(c *gin.Context) {

	var Bbody struct {
		Name    string `form:"name"`
		Email   string `form:"email"`
		Message string `form:"message"`
	}
	c.Bind(&Bbody)
	userdetails := Msg{
		Name:    Bbody.Name,
		Email:   Bbody.Email,
		Message: Bbody.Message,
	}

	c.HTML(200, "thankyou.html", userdetails.Name)
	from := os.Getenv("EMAIL_ADD")
	password := os.Getenv("APP_PASSWORD")
	toEmail := os.Getenv("EMAIL_ADD")
	to := []string{toEmail}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	//message
	subject := "Subject: Email from my website\n"
	name := userdetails.Name
	messagee := userdetails.Message
	email := userdetails.Email
	value := []byte(subject + name + "\n" + email + "\n" + messagee)
	message := value
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		fmt.Println("err:", err)
	}
}

// var tmpl *template.Template

// func processGetHandler(w http.ResponseWriter, r *http.Request) {
// 	var user Msg
// 	user.Name = r.FormValue("usernameame")
// 	user.Email = r.FormValue("useremail")
// 	user.Message = r.FormValue("usermessage")
// 	tmpl.ExecuteTemplate(w, "thankyou.html", user)
// }
