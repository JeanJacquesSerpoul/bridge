package distribution

import (
	"reflect"
	"testing"
)

var mockHand = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44, 45, 46}
var mockTest = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
var mockTestError = []int{50, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
var mockTestOne = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
var mockDelta = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 47, 48, 49, 50, 51}
var mockIndexToArray = []int{11, 7, 3, 10, 6, 2, 9, 5, 1, 12, 8, 4,
	0, 23, 19, 15, 22, 18, 14, 25, 21, 17, 13, 24, 20, 16, 35, 31,
	27, 38, 34, 30, 26, 37, 33, 29, 36, 32, 28, 51, 47, 43, 39, 50, 46,
	42, 49, 45, 41, 48, 44, 40}

//var maskDeskOk = "QT942.75.3.KQJ76 AJ.JT82.98642.98 865.AQ93.KQJ5.A4 K73.K64.AT7.T532"
const maskDeskOk = ".75.3.KQJ76 AJ.JT82.98.98 865..KQJ5.A4 K73.K64.AT7.T"

const maskDeskNot40 = "AKQJ.A.. .A.A. A..A. .AAA.."

const maskDesk4HandsNotOk = "QT942.75.3.KQJ76 865.AQ93.KQJ5.A4 K73.K64.AT7.T532"

const maskHand4SuitsNotOk = "QT942.75.3.KQJ76 AJ.JT82.98642.98 AQ93.KQJ5.A4 K73.K64.AT7.T532"

const maskDesk13CardsNotOk = "QT942.75.3.KQJ76 AJ.JT82.98642.98 865.AQ93.KQJ5.A4 K73.K64.AT7.T532T"

const handSuitOk = "QT942"

var mockDeskArrayOk = []int{22, 14, 5, 44, 40, 36, 20, 16,
	0, 1, 3, 4, 6, 51, 39, 38, 34, 26, 2, 29, 25, 28, 24, 9,
	11, 12, 27, 19, 15, 45, 41, 37, 13, 48, 8, 17, 30, 31, 35,
	47, 23, 7, 46, 18, 10, 49, 33, 21, 32, 42, 43, 50}

var suitTest = []int{42, 34, 30, 10, 2}

const maskTest = "... ... 865.AQ93.KQJ5.A4 ..."

var maskArrayTest = []int{-1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, 27, 19, 15, 50, 42, 30, 6, 45, 41, 37, 13, 48, 8,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

const strDeskOk = `{"index":"53607499485911673692344531919","pbn":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ","np":0,"ep":0,"sp":3,"wp":37}`

const jsHTTPOk = `{
	"count":1,
	"mask":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPNotMaskOk = `{
	"count":1,
	"mask":"AAA.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPMaxNotOk = `{
	"count":11,
	"mask":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPSuitOk = `{
	"count":1,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPSuitCommentOk = `{
	"count":1,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"OK",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPSuitNotOk2 = `{
	"count":1,
	"mask":"-1,-1,-1,A,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPSuitNotOk = `{
	"count":1,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPSuitMaxNotOk = `{
	"count":11,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPNotOk = `{
	"count":1
	"mask":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const maskHTTPOk = "432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"
const maskHTTPToComplete = "43.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"
const maskHTTPNotOk = "432A.A32.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"
const jsResultHTTPOk = `% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

`

const jsResultHTTPMaxOk = `% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

`

const jsResultCommentHTTPOk = `% OK
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

`

const jsonIndexWebOk = `{
	"index":"212287561920562011067931171"
}`

const jsonIndexWebNotOk = `{
	"index":"",
}`

const inputSuitDistok = "4,4,3,2,4,4,3,2,-1,4,-1,2,1,1,4,7"
const inputSuitDistNotOK1 = "10,4,3,2,4,4,3,2,-1,4,-1,2,1,1,4,7"
const inputSuitDistNotOK2 = "10,4,3,2,4,8,3,2,-8,4,-1,2,1,1,4,7"
const inputSuitDistNotOK3 = "4,4,3,20,4,4,3,2,-1,4,-1,2,1,1,4,7"
const inputSuitDistNotOK4 = "8,8,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1"
const inputSuitDistNotokCount = "4,4,3,2,4,4,3,2,4,4,3,2,1,4,7"
const inputSuitDistNotokNotaNumber = "4,4,3,2,A,4,3,2,4,4,3,2,1,1,4,7"

var arraySuitDistok = []int{4, 4, 3, 2, 4, 4, 3, 2, -1, 4, -1, 2, 1, 1, 4, 7}
var mockArraySuitDistok = []int{4, 2, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1}

const DeskIndexWebOk = `{\"index\":\"212287561920562011067931171\",\"pbn\":
\"AKQ.82.T9753.J72 JT432.AKQ753..T8 975.96.AKQJ42.65 86.JT4.86.AKQ943\",\"np\":10,
\"ep\":10,\"sp\":10,\"wp\":10}`

var mockSuitSorted = handSuit{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}}

const maskSuitOk = "4,2,-1,-1,2,-1,-1,-1,-1,-1,-1,-1,-1,-1,3,-1"
const maskSuitEmptyOK = "-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1"

const maskSuitOk2 = "7,4,2,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,8"
const maskSuitOk3 = "11,-1,-1,-1,-1,8,-1,-1,-1,-1,9,-1,-1,-1,-1,7"
const maskSuitNotOk4 = "11,-1,-1,-1,-1,8,-1,-1,-1,-1,9,-1,-1,-1"
const maskSuitNotOK = "7,-1,-1,5,-1,6,-1,2,1,-1,3,-1,-1,4,3,2"
const resultSuitOk = `{\"index\":\"53451083042398609111071938586\",\"pbn\":
\"5432.32.AK65.654 76.87654.987.987 JT98.JT9.QJT.QJT AKQ.AKQ.432.AK32\",
\"np\":7,\"ep\":0,\"sp\":8,\"wp\":25}`
const resultSuitOk2 = `{\"index\":\"52686974512864668390811270149\",\"pbn\":
\"8765432.5432.32. 9.9876.T98765.AT QJT.KQJT.KQJ.KQJ AK.A.A4.98765432\",\"np\":0,
\"ep\":4,\"sp\":21,\"wp\":15}`
const resultSuitOk3 = `{\"index\":\"50489164962052643089238898166\",
\"pbn\":\"QJT98765432.JT.. .98765432.KQ.AQT .Q.T98765432.KJ9 AK.AK.AJ.8765432\",
\"np\":4,\"ep\":11,\"sp\":6,\"wp\":19}`

const resultSuitOk4 = `{"index":"53607499485911673692344531919","pbn":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ","np":0,"ep":0,"sp":3,"wp":37}`

var mockNewpos = []int{10, 14, 27}
var mockDataInt = []int{40, 41, 43, 50}
var mockDataIntNotInList = []int{42, 49}

//var cardsWithPoints = []int{36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

const mockPointsOk = "-1,-1,-1,-1,-1,-1,-1,-1"
const mockPointResult = `{"index":"53607499485911673692344531919","pbn":"432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ","np":0,"ep":0,"sp":3,"wp":37}`
const mockPointsNotOk = "0,0,0,0,0,0"

const mockPointsNotOk2 = "32,8,0,0,0,0,0,0"

var MockArrayOfPointNotOK = []int{32, 8, 0, 0, 0, 0, 0, 0, 8}
var MockConstraintsPointOK = []int{0, 0, 0, 0, 0, 37, 0, 37}
var MockConstraintsPointNotOK = []int{0, 0, 0, 0, 8, 37, 0, 37}

const jsHTTPPointOk = `{
	"count":1,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const jsHTTPPointOkLimit = `{
	"count":11,
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`
const jsHTTPPointNotOk = `{
	"count":1
	"mask":"-1,-1,-1,-1,-1,-1,-1,-1",
	"comment":"",
	"dealer":"N",
	"vulnerable":"ALL"
}`

const mockJSONPointResult = `% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

`

const mockJSONPointLimitResult = `% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

% Index: 53607499485911673692344531919
[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"]

`

const dealNorthOk = "N:432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"
const dealEastOk = "E:765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ 432.432.432.5432"
const dealSouthOk = "S:T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ 432.432.432.5432 765.765.8765.876"
const dealWestOk = "W:AKQJ.AKQ.AKQ.AKQ 432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9"
const pbnFromDealOk = "432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"

var fakePointStruct = []int{0, 0, 0, 0, 0, 0, 0, 0}
var fakeLoadData = []string{}
var fake2LoadData = []string{"A\tB"}

//fakeRandom ...
type fakeRandom struct {
}

func (test *fakeRandom) fYShuffle(n int) []int {
	var random, temp int
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}
	for i := len(t) - 1; i >= 0; i-- {
		temp = t[i]
		random = i
		t[i] = t[random]
		t[random] = temp
	}
	return t
}

func Test_cardValueInt(t *testing.T) {
	type args struct {
		cardValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{28}, 7},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := cardValueInt(tt.args.cardValue); got != tt.want {
				t.Errorf("cardValueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardSuitInt(t *testing.T) {
	type args struct {
		cardValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{15}, 3},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := cardSuitInt(tt.args.cardValue); got != tt.want {
				t.Errorf("cardSuitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardLevel(t *testing.T) {
	type args struct {
		cardValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{17}, 0},
		{"Test1", args{46}, 3},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := cardLevel(tt.args.cardValue); got != tt.want {
				t.Errorf("cardLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CardsValue ...
func cardsValue(hand []int) int {
	hv := 0
	for _, v := range hand {
		hv += cardLevel(v)
	}
	return hv
}

func TestHandValue(t *testing.T) {
	type args struct {
		hand []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{mockHand}, 9},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := cardsValue(tt.args.hand); got != tt.want {
				t.Errorf("cardsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_faceCardsValue(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{"J"}, 9},
		{"Test2", args{"X"}, -1},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := faceCardsValue(tt.args.s); got != tt.want {
				t.Errorf("faceCardsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkVulnerable(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"X"}, false},
		{"Test2", args{"NONE"}, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := checkVulnerable(tt.args.s); got != tt.want {
				t.Errorf("checkVulnerable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDealer(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"X"}, false},
		{"Test2", args{"N"}, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := checkDealer(tt.args.s); got != tt.want {
				t.Errorf("checkDealer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskCheck(t *testing.T) {
	type args struct {
		mask string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test3", args{mask: ""}, true},
		{"Test4", args{mask: maskDeskOk}, false},
		{"Test5", args{mask: maskDesk4HandsNotOk}, true},
		{"Test6", args{mask: maskHand4SuitsNotOk}, true},
		{"Test7", args{mask: maskDesk13CardsNotOk}, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := maskCheck(tt.args.mask); (err != nil) != tt.wantErr {
				t.Errorf(" %v maskCheck() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func Test_deskCheck(t *testing.T) {
	type args struct {
		board Board
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test1", args{board: Board{
			Content:    []int{},
			Dealer:     "X",
			Vulnerable: "",
			Comment:    jsHTTPSuitCommentOk,
		}}, true},
		{"Test2", args{board: Board{
			Content:    []int{},
			Dealer:     "N",
			Vulnerable: "X",
			Comment:    jsHTTPSuitCommentOk,
		}}, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := boardCheck(tt.args.board); (err != nil) != tt.wantErr {
				t.Errorf(" %v deskCheck() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
func Test_handToArray(t *testing.T) {
	type args struct {
		hand []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockHand}, "23.234K.234K.234K"},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := handArrayToString(tt.args.hand); got != tt.want {
				t.Errorf("handToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_suitHandPbnToArray(t *testing.T) {
	type args struct {
		suitString string
		suit       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{suitString: handSuitOk, suit: 2}, suitTest},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := suitHandPbnToArray(tt.args.suitString, tt.args.suit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suitHandPbnToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPbn(t *testing.T) {
	type args struct {
		content []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockTest}, "432.432.432.5432 765.765.8765.876 T98.JT98.JT9.JT9 AKQJ.AKQ.AKQ.AKQ"},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := toPbn(tt.args.content); got != tt.want {
				t.Errorf("ToPbn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskToArray(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{maskTest}, maskArrayTest},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := MaskToArray(tt.args.mask); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maskToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intInSlice(t *testing.T) {
	type args struct {
		a    int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{a: 12, list: mockTest}, 12},
		{"Test2", args{a: 99, list: mockTest}, -1},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := intInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("intInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_delta(t *testing.T) {
	type args struct {
		slice    []int
		ToRemove []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{slice: mockTest, ToRemove: mockHand}, mockDelta},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := delta(tt.args.slice, tt.args.ToRemove); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToIndex(t *testing.T) {
	type args struct {
		content []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockTest}, "53607499485911673692344531919"},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ArrayToIndex(tt.args.content); got != tt.want {
				t.Errorf("ArrayToIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexToArray(t *testing.T) {
	type args struct {
		sIndex string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{"53607499485911673692344531919"}, mockIndexToArray},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IndexToArray(tt.args.sIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fYshuffle(t *testing.T) {
	t.Parallel()
	sh := new(Random)
	t1 := sh.fYShuffle(DC)
	t2 := sh.fYShuffle(DC)
	if reflect.DeepEqual(t1, t2) {
		t.Errorf("fYshuffle is not working")
	}

}

func Test_shuffleRemainHands(t *testing.T) {
	type args struct {
		tabMask      []int
		arrayOfPoint []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test1", args{mockTest, nil}, mockTest, false},
		{"Test2", args{mockTestOne, nil}, mockTest, false},
		{"Test2", args{mockTestError, MockArrayOfPointNotOK}, mockTest, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got, err := shuffleRemainHands(sh, nil,
				tt.args.arrayOfPoint,
				tt.args.tabMask); !reflect.DeepEqual(got, tt.want) {
				if tt.wantErr {
					if err == nil {
						t.Errorf("shuffleRemainHands() error = %v, wantErr %v", err, tt.wantErr)
					}
				} else {
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("shuffleRemainHands() = %v, want %v", got, tt.want)

					}
				}
			}
		})
	}
}

func Test_shuffleHand(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test1", args{mask: maskDeskOk}, mockDeskArrayOk, false},
		{"Test2", args{mask: ""}, mockDeskArrayOk, true},
		{"Test3", args{mask: maskDeskNot40}, mockDeskArrayOk, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := shuffleHand(sh, nil, tt.args.mask)
			if tt.wantErr {
				if err == nil {
					t.Errorf("shuffleHand() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("shuffleHand() = %v, want %v", got, tt.want)

				}
			}
		})
	}
}

func TestPbnGenerateFromJSON(t *testing.T) {
	type args struct {
		js string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{jsHTTPOk}, jsResultHTTPOk, false},
		{"Test2", args{jsHTTPNotOk}, jsResultHTTPOk, true},
		{"Test3", args{jsHTTPMaxNotOk}, jsResultHTTPMaxOk, false},
		{"Test4", args{jsHTTPNotMaskOk}, jsResultHTTPOk, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	MaxPbnGeneration = 2 // Less time for running test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnDataGenerateFromJSON(sh, tt.args.js)
			if tt.wantErr {
				if err == nil {
					t.Errorf("PbnGenerateFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("PbnGenerateFromJSON() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPbnAndIndexGenerateFromMask(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{maskHTTPOk}, strDeskOk, false},
		{"Test2", args{maskHTTPToComplete}, strDeskOk, false},
		{"Test3", args{maskHTTPNotOk}, strDeskOk, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnAndIndexGenerateFromMask(sh, nil, tt.args.mask)
			if tt.wantErr {
				if err == nil {
					t.Errorf("PbnAndIndexGenerateFromMask() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("PbnAndIndexGenerateFromMask() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPbnGenerateFromJSONIndex(t *testing.T) {
	type args struct {
		js string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{jsonIndexWebOk}, DeskIndexWebOk, false},
		{"Test2", args{jsonIndexWebNotOk}, DeskIndexWebOk, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnGenerateFromJSONIndex(tt.args.js)
			if tt.wantErr {
				if err == nil {
					t.Errorf("PbnGenerateFromJSONIndex() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("PbnGenerateFromJSONIndex() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_inputSuitDistToArray(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test1", args{inputSuitDistok}, arraySuitDistok, false},
		{"Test2", args{inputSuitDistNotokCount}, arraySuitDistok, true},
		{"Test3", args{inputSuitDistNotokNotaNumber}, arraySuitDistok, true},
		{"Test4", args{inputSuitDistNotOK1}, arraySuitDistok, true},
		{"Test5", args{inputSuitDistNotOK2}, arraySuitDistok, true},
		{"Test6", args{inputSuitDistNotOK3}, arraySuitDistok, true},
		{"Test7", args{inputSuitDistNotOK4}, arraySuitDistok, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := inputSuitDistToArray(tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("inputSuitDistToArray() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("inputSuitDistToArray() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_getSuitFromArrayOfSuits(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{7}, 3},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := getSuitFromArrayOfSuits(tt.args.v); got != tt.want {
				t.Errorf("getSuitFromArrayOfSuits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHandFromArrayOfSuits(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{5}, 1},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := getHandFromArrayOfSuits(tt.args.v); got != tt.want {
				t.Errorf("getHandFromArrayOfSuits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomSuitsToArraySuits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want handSuit
	}{
		{"Test1", args{maskSuitOk}, mockSuitSorted},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := randomSuitsToArraySuits(sh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomSuitsToArraySuits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbnAndIndexGenerateFromSuits(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{maskSuitOk}, resultSuitOk, false},
		{"Test2", args{maskSuitOk2}, resultSuitOk2, false},
		{"Test3", args{maskSuitOk3}, resultSuitOk3, false},
		{"Test4", args{maskSuitNotOk4}, resultSuitOk3, true},
		{"Test5", args{maskSuitNotOK}, resultSuitOk3, true},
		{"Test6", args{maskSuitEmptyOK}, resultSuitOk4, false},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnAndIndexGenerateFromSuits(sh, tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("TestPbnAndIndexGenerateFromSuits() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("TestPbnAndIndexGenerateFromSuits() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_getASuit(t *testing.T) {
	type args struct {
		arrayOfSuit      []int
		suitFromRandom   int
		handFromPosition int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{mockArraySuitDistok, 0, 3}, 4},
		{"Test2", args{mockArraySuitDistok, 0, 8}, -1},
		{"Test3", args{mockArraySuitDistok, 27, 39}, -1},
		{"Test4", args{mockArraySuitDistok, 12, 10}, 2},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := getASuit(tt.args.arrayOfSuit, tt.args.suitFromRandom, tt.args.handFromPosition); got != tt.want {
				t.Errorf("getASuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inNewPos(t *testing.T) {
	type args struct {
		v  int
		np []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{40, mockNewpos}, false},
		{"Test1", args{10, mockNewpos}, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := inNewPos(tt.args.v, tt.args.np); got != tt.want {
				t.Errorf("inNewPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbnSuitDataGenerateFromJSON(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{jsHTTPSuitOk}, jsResultHTTPOk, false},
		{"Test2", args{jsHTTPSuitNotOk}, jsResultHTTPOk, true},
		{"Test3", args{jsHTTPSuitMaxNotOk}, jsResultHTTPMaxOk, false},
		{"Test4", args{jsHTTPSuitNotOk2}, jsResultHTTPOk, true},
		{"Test5", args{jsHTTPSuitCommentOk}, jsResultCommentHTTPOk, false},
	}

	t.Parallel()
	sh := new(fakeRandom)
	MaxPbnGeneration = 2 // Less time for running test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnSuitDataGenerateFromJSON(sh, tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("TestPbnSuitDataGenerateFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("TestPbnSuitDataGenerateFromJSON() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_getRandomFromData(t *testing.T) {
	type args struct {
		pMin      int
		pMax      int
		notInList []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test1", args{10, 11, mockDataIntNotInList}, mockDataInt, false},
		{"Test2", args{11, 10, mockDataIntNotInList}, mockDataInt, true},
		{"Test3", args{-1, -1, mockDataIntNotInList}, mockDataInt, true},
		{"Test4", args{10, 11, cardsWithPoints()}, mockDataInt, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	loadingData := genDistWithPointToString(1, HCC)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetRandomFromData(sh, loadingData, tt.args.pMin, tt.args.pMax, tt.args.notInList)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Test_getRandomFromData() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Test_getRandomFromData() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGetPbnHandsFromPoints(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{mockPointsOk}, mockPointResult, false},
		{"Test2", args{mockPointsNotOk}, mockPointResult, true},
		{"Test3", args{mockPointsNotOk2}, mockPointResult, true},
	}

	t.Parallel()
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetPbnHandsFromPoints(sh, tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("TestGetPbnHandsFromPoints() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("TestGetPbnHandsFromPoints() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPbnPointDataGenerateFromJSON(t *testing.T) {
	type args struct {
		js string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{jsHTTPPointOk}, mockJSONPointResult, false},
		{"Test2", args{jsHTTPPointNotOk}, mockJSONPointResult, true},
		{"Test3", args{jsHTTPPointOkLimit}, mockJSONPointLimitResult, false},
	}

	t.Parallel()
	sh := new(fakeRandom)
	MaxPbnGeneration = 2 // Less time for running test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PbnPointDataGenerateFromJSON(sh, tt.args.js)
			if tt.wantErr {
				if err == nil {
					t.Errorf("TestPbnPointDataGenerateFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("TestPbnPointDataGenerateFromJSON() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_checkHandPoint(t *testing.T) {
	type args struct {
		content []int
		c       []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{mockTest, MockConstraintsPointOK}, true},
		{"Test2", args{mockTest, MockConstraintsPointNotOK}, false},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := checkHandPoint(tt.args.content, tt.args.c); got != tt.want {
				t.Errorf("checkHandPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateMask(t *testing.T) {
	type args struct {
		pbn string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{dealNorthOk}, pbnFromDealOk},
		{"Test2", args{dealEastOk}, pbnFromDealOk},
		{"Test3", args{dealSouthOk}, pbnFromDealOk},
		{"Test4", args{dealWestOk}, pbnFromDealOk},
		{"Test5", args{pbnFromDealOk}, pbnFromDealOk},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := rotateMask(tt.args.pbn); got != tt.want {
				t.Errorf("rotateMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPointsDistToStruct(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test1", args{"A,0,0,0,0,0,0,0"}, fakePointStruct, true},
		{"Test2", args{"0,A,0,0,0,0,0,0"}, fakePointStruct, true},
		{"Test3", args{"0,0,A,0,0,0,0,0"}, fakePointStruct, true},
		{"Test4", args{"0,0,0,A,0,0,0,0"}, fakePointStruct, true},
		{"Test5", args{"0,0,0,0,A,0,0,0"}, fakePointStruct, true},
		{"Test6", args{"0,0,0,0,0,A,0,0"}, fakePointStruct, true},
		{"Test7", args{"0,0,0,0,0,0,A,0"}, fakePointStruct, true},
		{"Test8", args{"0,0,0,0,0,0,0,A"}, fakePointStruct, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := inputPointsDistToStruct(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("inputPointsDistToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPointsDistToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadData(t *testing.T) {
	type args struct {
		input []string
		pMin  int
		pMax  int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Test1", args{fake2LoadData, 1, 13}, fakeLoadData, true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := loadData(tt.args.input, tt.args.pMin, tt.args.pMax)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
