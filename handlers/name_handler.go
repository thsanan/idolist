package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thsanan/idolist/domains"
	"github.com/thsanan/idolist/services"
)

type nameHandler struct {
	nameService services.NameService
	miracle     services.MiracleService
}

func NewNameHandler(nameSer services.NameService, miracleSer services.MiracleService) nameHandler {
	return nameHandler{nameSer, miracleSer}
}

func (handler nameHandler) NameMasterHandler(c *fiber.Ctx) error {
	title := "อณัญญา"
	dayBirth := "sunday"

	sumName := handler.nameService.SumAlphabetNumber(title)
	typeName, err := handler.miracle.GetTypeName(strconv.Itoa(sumName))
	if err != nil {
		panic(err.Error())
	}

	sumSha := handler.nameService.SumAlShadowNumber(title)
	typeSha, err := handler.miracle.GetTypeShaName(strconv.Itoa(sumSha))
	if err != nil {
		panic(err.Error())
	}

	nameMaster := domains.NameMaster{
		Name:             "อณัญญา",
		Day:              "sunday",
		AlphabetKalakini: handler.nameService.GetAlphabetKalakini(dayBirth, title),
		AlphabetNumbers:  handler.nameService.GetAlphabetNumber(title),
		SumnumAlphabet:   handler.nameService.SumAlphabetNumber(title),
		TypeNumAlphabet:  typeName,
		AlphabetShadows:  handler.nameService.GetAlphabetShadows(title),
		TypeShaAlphabet:  typeSha,
		SumshaAlphabet:   handler.nameService.SumAlShadowNumber(title),
	}

	return c.JSON(
		fiber.Map{
			"nameMaster": nameMaster,
		},
	)

}
