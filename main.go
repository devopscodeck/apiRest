package main

import (
	"fmt"

	"github.com/devopscodeck/apiRest/models"
	"github.com/devopscodeck/apiRest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome1", Historia: "Historia1"},
		{Nome: "Nome2", Historia: "Historia2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
