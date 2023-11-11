package main

import (
	app "ginchat/router"
)

func main() {
	application := app.Router()
	err := application.Run()
	if err != nil {
		return
	}
}
