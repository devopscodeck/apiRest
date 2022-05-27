package main

import (
	"fmt"

	"github.com/devopscodeck/apiRest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
