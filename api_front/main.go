package main

import "github.com/Eitol/document_color_meter/api_front/routes"

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
