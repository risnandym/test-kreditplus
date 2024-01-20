package main

import (
	"test-kreditplus/core/config"
	"test-kreditplus/src"
)

//	@title			Kredit+
//	@description	This is a Kredit+ test.
//	@version		1.0
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@termsOfService	http://swagger.io/terms/

//	@securitydefinitions.apikey	kreditplus-token
//	@description				Value is: "Bearer {access_token}", where access_token is retrieved from cms-service/v1/login
//	@in							header
//	@name						Authorization

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	dep := src.Dependencies()

	r := src.SetupRouter(dep)
	r.Run()
}
