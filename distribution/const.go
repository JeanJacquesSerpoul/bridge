package distribution

//FIRSTHANDSEPARATOR ...
const FIRSTHANDSEPARATOR = ":"

//BOARDMINUSHAND ...
const BOARDMINUSHAND = 3

//LoadingData ...
var LoadingData []string

//DATAFILE ...
const DATAFILE = "data.txt"

//TAB ...
const TAB = "\t"

//NEWLINE ...
const NEWLINE = "\n"

//MAXPOINTSINHAND ...
const MAXPOINTSINHAND = 37

//NBPC 8 points constraints
var NBPC = 8

//MaxPbnGeneration Arbitrary limit
var MaxPbnGeneration = 10000

// COMMA separator for suits distribution input
const COMMA = ","

// NOPOINT in distribution point
const NOPOINT = -1

// NOCARD in distribution array
const NOCARD = -1

// NOVALUESUIT in distribution array
const NOVALUESUIT = -1

//DC Desck Count ...
const DC = 52

//HCC Hand Cards count
const HCC = 13

//HC Hand count
const HC = 4

//DP Desk Point
const DP = 40

// SUITCOUNT ...
const SUITCOUNT = 16

// NbDist Number of different distribution CNP(52,13)*CNP(39,13)*CNP(26,13)
const NbDist = "53644737765488792839237440000"

//POINT in pbn
var POINT = "."

//SPACE in pbn
var SPACE = " "

//EMPTYHAND in pbn
var EMPTYHAND = "..."

//EMPTYDESK in pbn
var EMPTYDESK = "... ... ... ..."

//EMPTYSUIT ...
var EMPTYSUIT = "-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1"

//EMPTYPOINTS ...
var EMPTYPOINTS = "-1,-1,-1,-1,-1,-1,-1,-1"

//MAXTRY ...
var MAXTRY = 50

var faceCards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

var initHand = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

var firstDefaultHand = "N"

//Position ...
var Position = []string{"N", "E", "S", "W"}

//Vulnerable ...
var Vulnerable = []string{"ALL", "NS", "EW", "NONE"}

//POINTSVALUE to be replace by real values
var POINTSVALUE = "<POINTS>"

//INDEXVALUE to be replace by index value
var INDEXVALUE = "<INDEX>"

//ErrMsg ...
var ErrMsg = map[string]string{
	"no_40_points":               "Error in mask-err 1",
	"4_hands":                    " should have 4 hands-err 2",
	"4_suits":                    " should have 4 suits-err 3",
	"13_cards":                   " should have less than 13 cards-err 4",
	"err_board":                  "Error in board-err 5",
	"err_dealer":                 "Dealer should be N or E or W or S-err 6",
	"err_vuln":                   "Vulnerable should be ALL or NS or EW or NONE-err 7",
	"more_13_in_hand":            "You can't have more 13 cards in one hand-err 8",
	"more_13_in_suit":            "You can't have more 13 cards in one suit-err 9",
	"16_positions":               "You should have 16 positions in suit mask-err 10",
	"less_-1":                    "Value of the suit mask should be more than -1-err 11",
	"more_13":                    "Value of the suit mask should be less than 13-err 12",
	"Error_shuffle_remain_hands": "Error in shuffleRemainHands-err 13",
	"points_beetween_0_and_37":   "Points should be beetween 0 and 37-err 14",
	"pmin_more_than_pmax":        "Min points should be less than max points-err 15",
	"error_in_getRandomFromData": "Error in getRandomFromData-err 16",
	"Hands_with_more_40":         "Total of min points can't be more than 40 -err 17",
	"not_8_constraints":          "You should have 8 constraints in point -err 18",
	"bad_pbn":                    "Bad Pbn -err 19",
}
