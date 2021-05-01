package app

import (
	"github.com/dembygenesis/go-rest-industry-standard/mvc/controllers"
)

/**
Somehow the purpose of this file is to enabling quickly swapping Gin with another
http framework if we so choose.

Also, you just have a single endpoint to look at your routes.
 */

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}