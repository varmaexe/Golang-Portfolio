package handlers

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var tmpl *template.Template

// func TemplateParser() {
// 	fileserver := http.FileServer(http.Dir("assets"))
// 	("/assets/", http.StripPrefix("/assets", fileserver))
// 	tmpl, _ = tmpl.ParseGlob("templates/*.html")

// }

func Loginhandlers(c *gin.Context) {
	tmpl.ExecuteTemplate(c.Writer, "loginform.html", nil)

}
