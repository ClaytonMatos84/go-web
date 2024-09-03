package routes

import (
	"net/http"
	"web/controllers"
)

func LoadingRoutes()  {
	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)

	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Remove)
}
