package main

import (
	config "github.com/lucioPintanel/personal_finance_project/configs"
	"github.com/lucioPintanel/personal_finance_project/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		//Adicionar log
		return
	}
	router.Initialize()
}
