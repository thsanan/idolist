package services

import (
	"strconv"
	"strings"

	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/utils"
)

type NameService interface {
	GetAlphabetNumber(title string) []utils.Shaphabet
	GetAlphabetKalakini(dayBirth string, title string) []string
	GetAlphabetShadows(title string) []utils.Shaphabet
	SumAlphabetNumber(title string) int
	SumAlShadowNumber(title string) int
}

type Name struct {
	repo repositories.AlphabetRepo
}

func NewName(repo repositories.AlphabetRepo) NameService {
	return Name{repo: repo}
}

func (service Name) GetAlphabetNumber(title string) []utils.Shaphabet {
	var pairs []utils.Shaphabet
	ss := strings.Split(title, "")
	for _, s := range ss {
		for key, value := range service.repo.GetAlphabets() {
			if key == s {
				sha := new(utils.Shaphabet)
				sha.CharName = s
				sha.PairNumber = value
				pairs = append(pairs, *sha)
				break
			}
		}
	}
	return pairs

}

func (service Name) GetAlphabetKalakini(dayBirth string, title string) []string {
	var ks []string

	for day, kalakiniList := range service.repo.GetAlphabetKalakini() {

		if dayBirth == day {

			for _, s := range strings.Split(title, "") {

				for _, k := range strings.Split(kalakiniList, "") {

					if s == k {

						ks = append(ks, k)
						break
					}
				}
			}

			break
		}
	}

	return ks

}

func (service Name) GetAlphabetShadows(title string) []utils.Shaphabet {

	var pairs []utils.Shaphabet

	ss := strings.Split(title, "")
	for _, s := range ss {
		for key, value := range service.repo.GetAlphaShadow() {
			if key == s {
				sha := new(utils.Shaphabet)
				sha.CharName = s
				sha.PairNumber = value
				pairs = append(pairs, *sha)
				break
			}
		}
	}

	return pairs

}

func (service Name) SumAlphabetNumber(title string) int {
	var sum int
	ss := strings.Split(title, "")
	for _, s := range ss {
		for key, value := range service.repo.GetAlphabets() {
			if key == s {
				intConv, _ := strconv.Atoi(value)
				sum = sum + intConv
				break
			}
		}
	}
	return sum

}

func (service Name) SumAlShadowNumber(title string) int {
	var sum int
	ss := strings.Split(title, "")
	for _, s := range ss {
		for key, value := range service.repo.GetAlphaShadow() {
			if key == s {
				intConv, _ := strconv.Atoi(value)
				sum = sum + intConv
				break
			}
		}
	}
	return sum

}
