package main

import (
	"net/http"

	"github.com/KelvinWanyama/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
