package services

import (
	"github.com/thsanan/idolist/repositories"
)

type MiracleService interface {
	GetTypeName(sumAlpha string) (string, error)
	GetTypeShaName(sumSha string) (string, error)
}

type Miracle struct {
	repo repositories.MiracleRepository
}

func NewMiracle(repo repositories.MiracleRepository) MiracleService {
	return Miracle{repo: repo}
}

func (miracle Miracle) GetTypeName(sumAlpha string) (string, error) {
	var numType string
	numberMiras, err := miracle.repo.GetNumberMiracle()
	if err != nil {
		return "", nil
	}

	for _, numObj := range numberMiras {
		if numObj.PairNumber == sumAlpha {
			numType = numObj.PairType
			break
		}
	}

	return numType, nil
}

func (miracle Miracle) GetTypeShaName(sumSha string) (string, error) {
	var numType string
	numberMiras, err := miracle.repo.GetNumberMiracle()
	if err != nil {
		return "", err
	}

	for _, numObj := range numberMiras {
		if numObj.PairNumber == sumSha {
			numType = numObj.PairType
			break
		}
	}

	return numType, nil
}
