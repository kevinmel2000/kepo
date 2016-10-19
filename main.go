package main

import (
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

func main() {
	res, err := fb.Get("/me", fb.Params{
		"access_token": accessToken,
		"fields":       "id,name",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("ID: ", res["id"])
	fmt.Println("Name: ", res["name"])
}
