package repositories

import "github.com/thsanan/idolist/dbs"

type AlphabetRepo interface {
	GetAlphabets() map[string]string
	GetAlphabetKalakini() map[string]string
	GetAlphaShadow() map[string]string
}

type alphabetDb struct {
	db dbs.AlphabetDb
}

func NewAlphabetDb(db dbs.AlphabetDb) AlphabetRepo {
	return alphabetDb{db}
}

func (repo alphabetDb) GetAlphabets() map[string]string {
	return repo.db.GetAlphabetNumbers()
}

func (repo alphabetDb) GetAlphabetKalakini() map[string]string {
	return repo.db.GetAlphabetKalakini()

}

func (repo alphabetDb) GetAlphaShadow() map[string]string {
	return repo.db.GetAlphabetShadow()

}
