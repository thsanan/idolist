package repositories

import (
	"github.com/jmoiron/sqlx"
	models "github.com/thsanan/idolist/domains"
)

type MiracleRepository interface {
	GetNumberMiracle() ([]models.NumberMiracle, error)
}

type MiracleNumberDb struct {
	db *sqlx.DB
}

func NewMiracleNumber(db *sqlx.DB) MiracleRepository {
	return MiracleNumberDb{db: db}
}

func (repo MiracleNumberDb) GetNumberMiracle() ([]models.NumberMiracle, error) {
	numbers := []models.NumberMiracle{}
	err := repo.db.Select(&numbers, "SELECT pairnumber, pairtype, pairpoint FROM numbers ORDER BY pairnumber ASC")
	if err != nil {
		return nil, err
	}
	return numbers, nil
}
