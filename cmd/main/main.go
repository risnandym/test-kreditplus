package main

import (
	"test-kreditplus/core/config"
	app "test-kreditplus/src"
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

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	dep := app.Dependencies()

	r := app.SetupRouter(dep)
	r.Run()
}
