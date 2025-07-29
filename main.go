package main

import (
	"fmt"
	"os"

	"github.com/vknow360/shareIO/routes"
	"github.com/vknow360/shareIO/utils"
)

func main() {

	ip := utils.GetLocalIP()
	r := routes.RegisterRoutes()

	err := r.Run("0.0.0.0:8000")
	fmt.Printf("Access at: http://%s:8000\n", ip)

	if err != nil {
		os.Exit(-1)
	}
}
