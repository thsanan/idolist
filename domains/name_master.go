package domains

import "github.com/thsanan/idolist/utils"

type NameMaster struct {
	Name             string            `json:"name"`
	Day              string            `json:"day"`
	AlphabetKalakini []string          `json:"alphabet_kalakini"`
	AlphabetNumbers  []utils.Shaphabet `json:"alphabet_numbers"`
	SumnumAlphabet   int               `json:"sumnum_alphabet"`
	TypeNumAlphabet  string            `json:"type_num_alphabet"`
	AlphabetShadows  []utils.Shaphabet `json:"alphabet_shadows"`
	SumshaAlphabet   int               `json:"sumsha_alphabet"`
	TypeShaAlphabet  string            `json:"type_sha_alphabet"`
}
