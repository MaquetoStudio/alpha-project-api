package main

import (
	"fmt"

	"alpha-project-api/api/v1"
)

func main() {
	r := api.SetupRouter()
	fmt.Println("toto 3")
	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
