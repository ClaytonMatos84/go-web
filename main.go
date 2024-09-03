package main

import (
	"net/http"
	"web/routes"
)

func main()  {
	routes.LoadingRoutes()
	http.ListenAndServe(":8000", nil)
}
