package main

import (
	"alpha-project-api/api"
)

func main() {
	r := api.SetupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
