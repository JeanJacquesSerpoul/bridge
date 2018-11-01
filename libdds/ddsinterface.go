package libdds

/* USE dds library
https://github.com/dds-bridge/dds
*/
/* Very minimalist use */
import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
)

var sep = ";"
var lRes = 20

//JSONtableRes ...
type JSONtableRes struct {
	Vulnerable   int    `json:"vul"`
	ParNSScore   string `json:"nss,omitempty"`
	ParEWScore   string `json:"ews,omitempty"`
	ParNSList    string `json:"nsl,omitempty"`
	ParEWList    string `json:"ewl,omitempty"`
	NorthNT      int    `json:"nnt"`
	NorthSpades  int    `json:"ns"`
	NorthHeart   int    `json:"nh"`
	NorthDiamond int    `json:"nd"`
	NorthClub    int    `json:"nc"`

	SouthNT      int `json:"snt"`
	SouthSpades  int `json:"ss"`
	SouthHeart   int `json:"sh"`
	SouthDiamond int `json:"sd"`
	SouthClub    int `json:"sc"`

	EastNT      int `json:"ent"`
	EastSpades  int `json:"es"`
	EastHeart   int `json:"eh"`
	EastDiamond int `json:"ed"`
	EastClub    int `json:"ec"`

	WestNT      int `json:"wnt"`
	WestSpades  int `json:"ws"`
	WestHeart   int `json:"wh"`
	WestDiamond int `json:"wd"`
	WestClub    int `json:"wc"`
}

var tableRes JSONtableRes

func convertVul(s string) int {
	if s == "ALL" {
		return 1
	}
	if s == "NS" {
		return 2
	}
	if s == "EW" {
		return 3
	}
	return 0
}

//CallParDDS ...
func CallParDDS(pbn string, sVul string) (string, error) {
	var err error
	pbn = strings.ToUpper(pbn)
	vul := convertVul(sVul)
	data, err := exec.Command("calcPar", pbn, strconv.Itoa(vul)).Output()
	if err != nil {
		return "", err
	}
	sData := string(data)
	r := strings.Split(sData, "\n")
	sPar := r[0]
	sp := strings.Split(sPar, sep)
	tab := r[1]
	v := strings.Split(tab, sep)
	for i := 0; i < len(v); i++ {
		v[i] = strings.TrimSpace(v[i])
	}
	tableRes.Vulnerable = vul
	tableRes.ParNSScore = sp[0]
	tableRes.ParEWScore = sp[1]
	tableRes.ParNSList = sp[2]
	tableRes.ParEWList = sp[3]

	tableRes.NorthNT, _ = strconv.Atoi(v[0])
	tableRes.SouthNT, _ = strconv.Atoi(v[1])
	tableRes.EastNT, _ = strconv.Atoi(v[2])
	tableRes.WestNT, _ = strconv.Atoi(v[3])

	tableRes.NorthSpades, _ = strconv.Atoi(v[4])
	tableRes.SouthSpades, _ = strconv.Atoi(v[5])
	tableRes.EastSpades, _ = strconv.Atoi(v[6])
	tableRes.WestSpades, _ = strconv.Atoi(v[7])

	tableRes.NorthHeart, _ = strconv.Atoi(v[8])
	tableRes.SouthHeart, _ = strconv.Atoi(v[9])
	tableRes.EastHeart, _ = strconv.Atoi(v[10])
	tableRes.WestHeart, _ = strconv.Atoi(v[11])

	tableRes.NorthDiamond, _ = strconv.Atoi(v[12])
	tableRes.SouthDiamond, _ = strconv.Atoi(v[13])
	tableRes.EastDiamond, _ = strconv.Atoi(v[14])
	tableRes.WestDiamond, _ = strconv.Atoi(v[15])

	tableRes.NorthClub, _ = strconv.Atoi(v[16])
	tableRes.SouthClub, _ = strconv.Atoi(v[17])
	tableRes.EastClub, _ = strconv.Atoi(v[18])
	tableRes.WestClub, _ = strconv.Atoi(v[19])
	json, err := json.Marshal(tableRes)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
