package main

import (
	"fmt"

	"github.com/thsanan/idolist/dbs"
	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/services"
)

func main() {
	db := dbs.NewAlphabetNumbers()
	repo := repositories.NewAlphabetDb(*db)
	service := services.NewName(repo, "อณัญญา", "sunday")

	fmt.Println(service.GetAlphabetNumber())
	fmt.Println(service.GetAlphabetKalakini())
	fmt.Println(service.GetAlphabetShadows())
	fmt.Println(service.SumAlphabetNumber())

}
