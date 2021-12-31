package domains

type NumberMiracle struct {
	PairNumber string `db:"pairnumber" json:"pairnumber"`
	PairType   string `db:"pairtype" json:"pairtype"`
	PairPoint  int    `db:"pairpoint" json:"pairpoint"`
}
