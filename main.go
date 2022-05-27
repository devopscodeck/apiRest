package main

import (
	"fmt"

	"github.com/devopscodeck/apiRest/models"
	"github.com/devopscodeck/apiRest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
