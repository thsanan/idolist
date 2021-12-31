package services

import (
	"strconv"
	"strings"

	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/utils"
)

type NameService interface {
	GetAlphabetNumber() []utils.Shaphabet
	GetAlphabetKalakini() []string
	GetAlphabetShadows() []utils.Shaphabet
	SumAlphabetNumber() int
	SumAlShadowNumber() int
}

type Name struct {
	repo  repositories.AlphabetRepo
	title string
	day   string
}

func NewName(repo repositories.AlphabetRepo, title string, day string) NameService {
	return Name{repo: repo, title: title, day: day}
}

func (service Name) GetAlphabetNumber() []utils.Shaphabet {
	var pairs []utils.Shaphabet
	ss := strings.Split(service.title, "")
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

func (service Name) GetAlphabetKalakini() []string {
	var ks []string

	for day, kalakiniList := range service.repo.GetAlphabetKalakini() {

		if service.day == day {

			for _, s := range strings.Split(service.title, "") {

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

func (service Name) GetAlphabetShadows() []utils.Shaphabet {

	var pairs []utils.Shaphabet

	ss := strings.Split(service.title, "")
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

func (service Name) SumAlphabetNumber() int {
	var sum int
	ss := strings.Split(service.title, "")
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

func (service Name) SumAlShadowNumber() int {
	var sum int
	ss := strings.Split(service.title, "")
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
