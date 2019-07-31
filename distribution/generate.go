package distribution

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* Copyright 2018 Jean-Jacques Serpoul
under APACHE 2.0 license*
https://github.com/JeanJacquesSerpoul/bridge
*/

/* Cards coding
Cards	Suit	Card Value(D)	Height Value(D)	Suit Value(D)	Card Value(B)	Height Value(B)	Suit Value(B)
2	C	0	0	0	000000	0000	00
2	D	1	0	1	000001	0000	01
2	H	2	0	2	000010	0000	10
2	S	3	0	3	000011	0000	11
3	C	4	1	0	000100	0001	00
3	D	5	1	1	000101	0001	01
3	H	6	1	2	000110	0001	10
3	S	7	1	3	000111	0001	11
4	C	8	2	0	001000	0010	00
4	D	9	2	1	001001	0010	01
4	H	10	2	2	001010	0010	10
4	S	11	2	3	001011	0010	11
5	C	12	3	0	001100	0011	00
5	D	13	3	1	001101	0011	01
5	H	14	3	2	001110	0011	10
5	S	15	3	3	001111	0011	11
6	C	16	4	0	010000	0100	00
6	D	17	4	1	010001	0100	01
6	H	18	4	2	010010	0100	10
6	S	19	4	3	010011	0100	11
7	C	20	5	0	010100	0101	00
7	D	21	5	1	010101	0101	01
7	H	22	5	2	010110	0101	10
7	S	23	5	3	010111	0101	11
8	C	24	6	0	011000	0110	00
8	D	25	6	1	011001	0110	01
8	H	26	6	2	011010	0110	10
8	S	27	6	3	011011	0110	11
9	C	28	7	0	011100	0111	00
9	D	29	7	1	011101	0111	01
9	H	30	7	2	011110	0111	10
9	S	31	7	3	011111	0111	11
10	C	32	8	0	100000	1000	00
10	D	33	8	1	100001	1000	01
10	H	34	8	2	100010	1000	10
10	S	35	8	3	100011	1000	11
J	C	36	9	0	100100	1001	00
J	D	37	9	1	100101	1001	01
J	H	38	9	2	100110	1001	10
J	S	39	9	3	100111	1001	11
Q	C	40	10	0	101000	1010	00
Q	D	41	10	1	101001	1010	01
Q	H	42	10	2	101010	1010	10
Q	S	43	10	3	101011	1010	11
K	C	44	11	0	101100	1011	00
K	D	45	11	1	101101	1011	01
K	H	46	11	2	101110	1011	10
K	S	47	11	3	101111	1011	11
A	C	48	12	0	110000	1100	00
A	D	49	12	1	110001	1100	01
A	H	50	12	2	110010	1100	10
A	S	51	12	3	110011	1100	11

*/
func init() {
	rand.Seed(time.Now().UnixNano())
	LoadingData = genDistWithPointToString(1, 13)
}

func cardsWithPoints() []int {
	return []int{
		36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}

func genereList(seq []int, k int) []int {
	var p, s []int
	var ii, jj uint
	fmax := math.Exp2((float64)(len(seq))) - 1
	imax := int(fmax)
	for i := 0; i <= imax; i++ {
		s = nil
		jmax := len(seq) - 1
		for j := 0; j <= jmax; j++ {
			ii = uint(i)
			jj = uint(j)
			if (ii>>jj)&1 == 1 {
				s = append(s, seq[j])
			}
		}
		if len(s) == k {
			p = append(p, s...)
		}
	}
	return p
}

func genereArrayFromList(seq []int, k int) [][]int {
	var m int
	v := genereList(seq, k)
	m = len(v) / k
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			t[i] = append(t[i], v[i*k+j])
		}
	}
	return t
}

func genDistWithPoint(k int) ([]int, []string) {
	var mask, tpoints []int
	var points int
	var stemp string
	var result []string
	tmask := genereArrayFromList(cardsWithPoints(), k)
	m := len(tmask)
	for i := 0; i < m; i++ {
		mask = tmask[i]
		points = 0
		for k := 0; k < len(mask); k++ {
			points += cardLevel(mask[k])
		}
		tpoints = append(tpoints, points)
		stemp = ""
		for k := 0; k < len(mask); k++ {
			stemp = stemp + strconv.Itoa(mask[k]) + " "
		}
		stemp = strings.TrimSpace(stemp)
		result = append(result, stemp)
	}
	return tpoints, result
}

func saveGenDistWithPointToString(
	points []int,
	list []string,
) []string {
	var result []string
	for i, po := range points {

		result = append(result, strconv.Itoa(po)+"\t"+list[i])
	}
	return result
}

func genDistWithPointToString(kmin, kmax int) []string {
	var tp, p []int
	var td, d []string
	var list []listData
	for i := kmin; i <= kmax; i++ {
		p, d = genDistWithPoint(i)
		tp = append(tp, p...)
		td = append(td, d...)
	}
	list = make([]listData, len(tp))
	for i := 0; i < len(tp); i++ {
		list[i].points = tp[i]
		list[i].dist = td[i]
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].points < list[j].points
	})
	for i := 0; i < len(tp); i++ {
		tp[i] = list[i].points
		td[i] = list[i].dist
	}
	result := saveGenDistWithPointToString(tp, td)
	return result
}

//GetPbnHandsFromPoints ...
func GetPbnHandsFromPoints(sh ShuffleInterface, input string) (string, error) {
	var ar, c []int
	var err error

	c, err = inputPointsDistToStruct(input)
	if err != nil {
		return "", err
	}
	if c[1] == NOPOINT {
		c[1] = MAXPOINTSINHAND
	}
	if c[3] == NOPOINT {
		c[3] = MAXPOINTSINHAND
	}
	if c[5] == NOPOINT {
		c[5] = MAXPOINTSINHAND
	}
	if c[7] == NOPOINT {
		c[7] = MAXPOINTSINHAND
	}

	if c[0] == NOPOINT {
		c[0] = MINPOINTSINHAND
	}
	if c[2] == NOPOINT {
		c[2] = MINPOINTSINHAND
	}
	if c[4] == NOPOINT {
		c[4] = MINPOINTSINHAND
	}
	if c[6] == NOPOINT {
		c[6] = MINPOINTSINHAND
	}

	for i := 0; i < MAXTRY; i++ {
		ar, err = GetHandsFromPoints(sh, c)
		if err == nil {
			break
		}
	}
	if err != nil {
		return "", err
	}
	ar = sortCards(ar)
	s := toPbn(ar)
	v := boardValueToArray(ar)
	result := ResultHTTP{Index: ArrayToIndex(ar), Pbn: s, NorthPoint: v[0], EastPoint: v[1], SouthPoint: v[2], WestPoint: v[3]}
	r, _ := json.Marshal(result)
	return string(r), nil
}

//GetHandsFromPoints ...
func GetHandsFromPoints(sh ShuffleInterface, c []int) ([]int, error) {
	var v, r []int
	var hand [HC][]int
	var err error
	mask := make([]int, DC)
	ps := make([]pointStruct, HC)
	ps[0].Orientation = 0
	ps[1].Orientation = 1
	ps[2].Orientation = 2
	ps[3].Orientation = 3
	ps[0].MinPoints = c[0]
	ps[1].MinPoints = c[2]
	ps[2].MinPoints = c[4]
	ps[3].MinPoints = c[6]

	ps[0].MaxPoints = c[1]
	ps[1].MaxPoints = c[3]
	ps[2].MaxPoints = c[5]
	ps[3].MaxPoints = c[7]

	sort.Slice(ps, func(i, j int) bool { return ps[i].MinPoints > ps[j].MinPoints })
	if DP-ps[0].MinPoints < ps[1].MaxPoints {
		ps[1].MaxPoints = ps[1].MaxPoints - ps[0].MinPoints + BOARDMINUSHAND
	}
	if DP-ps[1].MinPoints < ps[2].MaxPoints {
		ps[2].MaxPoints = ps[2].MaxPoints - ps[1].MinPoints - ps[0].MinPoints + BOARDMINUSHAND
	}
	if DP-ps[2].MinPoints-ps[1].MinPoints < ps[3].MaxPoints {
		ps[3].MaxPoints = ps[3].MaxPoints - ps[2].MinPoints - ps[1].MinPoints - ps[0].MinPoints + BOARDMINUSHAND
	}
	hand[ps[0].Orientation], err = GetRandomFromData(sh, ps[0].MinPoints, ps[0].MaxPoints, nil)
	if err != nil {
		return nil, err
	}
	for i := 1; i < HC; i++ {
		v = append(v, hand[ps[i-1].Orientation]...)
		hand[ps[i].Orientation], err = GetRandomFromData(sh, ps[i].MinPoints, ps[i].MaxPoints, v)
		if err != nil {
			return nil, err
		}
	}
	for i := 0; i < DC; i++ {
		mask[i] = NOCARD
	}
	for i := 0; i < HC; i++ {
		for j := 0; j < len(hand[i]); j++ {
			mask[HCC*i+j] = hand[i][j]
		}
	}
	r, err = shuffleRemainHands(sh, nil, c, mask)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func loadData(input []string, pMin, pMax int) ([]string, error) {
	var d, stemp []string
	var k int
	var err error
	if pMin < 0 || pMax < 0 || pMin > MAXPOINTSINHAND || pMax > MAXPOINTSINHAND {
		return nil, fmt.Errorf(ErrMsg["points_beetween_0_and_37"])
	}
	if pMin > pMax {
		return nil, fmt.Errorf(ErrMsg["pmin_more_than_pmax"])
	}
	// Select Random value beetween pmin and pmax
	m := shuffleInterval(pMin, pMax)
	pMin = m
	pMax = m
	//
	for i := 0; i < len(input); i++ {
		stemp = strings.Split(input[i], TAB)
		k, err = strconv.Atoi(stemp[0])
		if err != nil {
			return nil, err
		}
		if k > pMax {
			break
		}
		if k >= pMin {
			d = append(d, stemp[1])
		}
	}

	return d, err
}
func checkList(list, notInList []int) bool {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(notInList); j++ {
			if list[i] == notInList[j] {
				return false
			}
		}
	}
	return true
}
func atoiArray(s []string) ([]int, error) {
	var result []int
	var itemp int
	var err error
	for i := 0; i < len(s); i++ {
		itemp, err = strconv.Atoi(s[i])
		if err != nil {
			return nil, err
		}
		result = append(result, itemp)
	}
	return result, nil
}

//GetRandomFromData ...
func GetRandomFromData(sh ShuffleInterface, pMin, pMax int, notInList []int) ([]int, error) {
	var list, atemp []int
	var r string
	var v []string
	s, err := loadData(LoadingData, pMin, pMax)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	k := len(s)
	for i := 0; i < k; i++ {
		list = append(list, i)
	}
	random := shuffle(sh, list)
	for i := 0; i < len(random); i++ {
		r = s[random[i]]
		v = strings.Split(r, SPACE)
		atemp, err = atoiArray(v)
		if err != nil {
			return nil, err
		}
		if checkList(atemp, notInList) {
			return atemp, nil
		}
	}
	err = fmt.Errorf(ErrMsg["error_in_getRandomFromData"])
	return nil, err
}

func intInSlice(a int, list []int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == a {
			return a
		}
	}
	return -1
}
func delta(slice []int, ToRemove []int) []int {
	var diff []int
	var n int
	for i := 0; i < len(slice); i++ {
		n = intInSlice(slice[i], ToRemove)
		if n < 0 {
			diff = append(diff, slice[i])
		}
	}
	return diff
}

func getSuitFromArrayOfSuits(v int) int {
	return v % HC
}
func getHandFromArrayOfSuits(v int) int {
	return v / HC
}

func checkInputSuit(t []int) error {
	var err error
	for i := 0; i < SUITCOUNT; i++ {
		if t[i] > HCC {
			err = fmt.Errorf(ErrMsg["more_13"])
			return err
		}
		if t[i] < NOCARD {
			err = fmt.Errorf(ErrMsg["less_-1"])
			return err
		}
	}

	return err
}

func checkArraySuit(t2 [HC][HC]int) error {
	var err error
	var valHand, valSuit int
	err = nil
	for i := 0; i < HC; i++ {
		valHand = 0
		for j := 0; j < HC; j++ {
			valHand = valHand + t2[i][j]
		}
		if valHand > HCC {
			err = fmt.Errorf(ErrMsg["more_13_in_suit"])
			return err
		}
	}
	for j := 0; j < HC; j++ {
		valSuit = 0
		for i := 0; i < HC; i++ {
			valSuit = valSuit + t2[i][j]
		}
		if valSuit > HCC {
			err = fmt.Errorf(ErrMsg["more_13_in_hand"])
			return err

		}
	}

	return err
}
func inputPointsDistToStruct(input string) ([]int, error) {
	var n int
	var err error
	v := make([]int, 8)
	r := strings.Split(input, COMMA)
	if len(r) != NBPC {
		err = fmt.Errorf(ErrMsg["not_8_constraints"])
		return v, err
	}
	n, err = strconv.Atoi(r[0])
	if err != nil {
		return v, err
	}
	v[0] = n
	n, err = strconv.Atoi(r[1])
	if err != nil {
		return v, err
	}
	v[1] = n
	n, err = strconv.Atoi(r[2])
	if err != nil {
		return v, err
	}
	v[2] = n
	n, err = strconv.Atoi(r[3])
	if err != nil {
		return v, err
	}
	v[3] = n

	n, err = strconv.Atoi(r[4])
	if err != nil {
		return v, err
	}
	v[4] = n
	n, err = strconv.Atoi(r[5])
	if err != nil {
		return v, err
	}
	v[5] = n
	n, err = strconv.Atoi(r[6])
	if err != nil {
		return v, err
	}
	v[6] = n
	n, err = strconv.Atoi(r[7])
	if err != nil {
		return v, err
	}
	v[7] = n

	return v, nil
}
func inputSuitDistToArray(input string) ([]int, error) {
	var err error
	var t2 [HC][HC]int
	r := strings.Split(input, COMMA)
	if len(r) != SUITCOUNT {
		err = fmt.Errorf(ErrMsg["16_positions"])
		return nil, err
	}
	t := make([]int, SUITCOUNT)
	for i := 0; i < SUITCOUNT; i++ {
		t[i], err = strconv.Atoi(r[i])
		if err != nil {
			err = fmt.Errorf(r[i] + ErrMsg["not_a number"])
			return nil, err
		}
	}
	err = checkInputSuit(t)
	if err != nil {
		return nil, err
	}
	for i := 0; i < SUITCOUNT; i++ {
		if t[i] > 0 {
			t2[getSuitFromArrayOfSuits(i)][getHandFromArrayOfSuits(i)] += t[i]
		}
	}
	err = checkArraySuit(t2)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// cardValueInt ...
func cardValueInt(cardValue int) int { return cardValue >> 2 }

// cardSuitInt ...
func cardSuitInt(cardValue int) int { return cardValue & 3 }

func cardLevel(cardValue int) int {
	if cardValue < 36 {
		return 0
	}
	v := cardValueInt(cardValue)
	return (v - 8)
}

func boardValue(board []int) int {
	v := 0
	for i := 0; i < len(board); i++ {
		v += cardLevel(board[i])
	}
	return v
}

func handArrayToString(hand []int) string {
	var suitHand handSuit
	var c [HC]int
	l := len(hand)
	r := ""
	for i := 0; i < l; i++ {
		suitHand[cardSuitInt(hand[i])][c[cardSuitInt(hand[i])]] = hand[i]
		c[cardSuitInt(hand[i])]++
	}
	for i := HC - 1; i >= 0; i-- {
		for j := 0; j < c[i]; j++ {
			r += faceCards[cardValueInt(suitHand[i][j])]
		}
		if i > 0 {
			r += POINT
		}
	}
	return r
}

func faceCardsValue(s string) int {
	for i := 0; i < HCC; i++ {
		if faceCards[i] == s {
			return i
		}
	}
	return -1
}

func sortCards(content []int) []int {
	var handList []listHand
	var cardValue, cardSuit, CardSort int
	l := len(content)
	tabCopy := make([]int, l)
	tabHand := make([]int, HCC)
	handList = make([]listHand, HCC)
	for i := 0; i < HC; i++ {
		for j := 0; j < HCC; j++ {
			tabHand[j] = content[j+i*HCC]
			cardValue = cardValueInt(tabHand[j])
			cardSuit = cardSuitInt(tabHand[j])
			CardSort = (cardSuit << 4) + cardValue
			handList[j].value = tabHand[j]
			handList[j].sortOrder = CardSort
		}
		sort.Slice(handList, func(i, j int) bool { return handList[i].sortOrder > handList[j].sortOrder })
		for j := 0; j < HCC; j++ {
			tabCopy[j+i*HCC] = handList[j].value
		}
	}
	return tabCopy
}

//toPbn ...
func toPbn(content []int) string {
	r := ""
	tabHand := make([]int, HCC)
	tabCopy := sortCards(content)
	for i := 0; i < HC; i++ {
		for j := 0; j < HCC; j++ {
			tabHand[j] = tabCopy[j+i*HCC]
		}
		r += handArrayToString(tabHand) + SPACE
	}
	r = strings.TrimSpace(r)
	return r
}

//PbnAndIndexGenerateFromSuits ...
func PbnAndIndexGenerateFromSuits(sh ShuffleInterface, input string) (string, error) {
	var hand, suit, itemp int
	var k [HC]int
	var c [HC][HC]string
	var v []int
	var err error
	var result string
	v, err = inputSuitDistToArray(input)
	if err != nil {
		return "", err
	}
	r := randomSuitsToArraySuits(sh)
	for i := 0; i < SUITCOUNT; i++ {
		suit = getSuitFromArrayOfSuits(i)
		hand = getHandFromArrayOfSuits(i)
		if v[i] >= 0 {
			if suit > 0 {
				c[hand][suit] += POINT
			}
			for j := 0; j < v[i]; j++ {
				itemp = r[suit][k[suit]]
				k[suit]++
				c[hand][suit] += faceCards[itemp]
			}
		} else {
			if suit > 0 {
				c[hand][suit] += POINT
			}
		}
	}
	mask := ""
	for i := 0; i < HC; i++ {
		for j := 0; j < HC; j++ {
			mask += c[i][j]
		}
		mask += " "
	}
	mask = strings.TrimSpace(mask)

	result, err = PbnAndIndexGenerateFromMask(sh, v, mask)
	return result, err
}

// fYShuffle Implement Fisher and Yates shuffle method
func (rd *Random) fYShuffle(n int) []int {
	var random, temp int
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}
	for i := len(t) - 1; i >= 0; i-- {
		temp = t[i]
		random = rand.Intn(i + 1)
		t[i] = t[random]
		t[random] = temp
	}
	return t
}

func randomSuitsToArraySuits(s ShuffleInterface) handSuit {
	var v handSuit
	var sh []int
	for i := 0; i < HC; i++ {
		sh = s.fYShuffle(HCC)
		for j := 0; j < HCC; j++ {
			v[i][j] = sh[j]
		}
	}
	return v
}
func shuffleInterval(min, max int) int {
	var t []int
	for i := min; i <= max; i++ {
		t = append(t, i)
	}
	l := max - min + 1
	r := rand.Intn(l)
	return (t[r])
}

//Shuffle ...
func shuffle(sh ShuffleInterface, content []int) []int {
	l := len(content)
	tabCopy := make([]int, l)
	for i := 0; i < l; i++ {
		tabCopy[i] = content[i]
	}
	tab := sh.fYShuffle(l)
	for i := 0; i < l; i++ {
		tabCopy[tab[i]] = content[i]
	}
	return tabCopy
}
func suitHandPbnToArray(suitString string, suit int) []int {
	var r []int
	var v int
	for i := 0; i < len(suitString); i++ {
		v = faceCardsValue(suitString[i : i+1])
		v = (v << 2) + suit
		r = append(r, v)
	}
	return r
}

func shuffleHand(sh ShuffleInterface, arrayOfSuit []int, mask string) ([]int, error) {
	mask = strings.ToUpper(mask)
	err := maskCheck(mask)
	if err != nil {
		return nil, err
	}
	v := MaskToArray(mask)
	r, err := shuffleRemainHands(sh, arrayOfSuit, nil, v)
	if err != nil {
		return nil, err
	}
	if boardValue(r) != DP {
		err = fmt.Errorf(ErrMsg["no_40_points"])
	}
	if err != nil {
		return nil, err
	}
	return r, err
}

func fillHand(board []int) []int {
	result := make([]int, HCC)
	l := len(board)
	for i := 0; i < l; i++ {
		result[i] = board[i]
	}
	for i := l; i < HCC; i++ {
		result[i] = NOCARD
	}
	return result
}

//MaskToArray ...
func MaskToArray(mask string) []int {
	var hand string
	var v, r []int
	var tabHand []string
	hcol := make([][]int, HC)
	s := strings.Split(mask, SPACE)
	for i := 0; i < HC; i++ {
		hand = strings.TrimSpace(s[i])
		tabHand = strings.Split(hand, POINT)
		for j := 0; j < HC; j++ {
			v = suitHandPbnToArray(tabHand[j], HC-j-1)
			for k := 0; k < len(v); k++ {
				hcol[i] = append(hcol[i], v[k])
			}
		}
	}
	for i := 0; i < HC; i++ {
		hcol[i] = fillHand(hcol[i])
		r = append(r, hcol[i]...)
	}
	return r
}

func getASuit(arrayOfSuit []int, handFromPosition, suitFromRandom int) int {
	suit := cardSuitInt(suitFromRandom)
	hand := getHandFromDist(handFromPosition)
	aSuit := arrayOfSuit[hand*HC+(HC-suit-1)]
	return aSuit
}

func inNewPos(v int, np []int) bool {
	result := false
	for i := 0; i < len(np); i++ {
		if v == np[i] {
			return true
		}
	}
	return result
}

func checkInMask(tabMask []int, v int) bool {
	for i := 0; i < len(tabMask); i++ {
		if v == tabMask[i] {
			return true
		}
	}
	return false
}

func advanceCheck(tabMask, remain, arrayOfSuit []int, r, newpos *[]int) bool {
	var check bool
	success := false
	if len(remain) == 0 {
		return true
	}
	for k := 0; k < len(remain); k++ {
		success = false
		for i := 0; i < DC; i++ {
			if success {
				break
			}
			if (*r)[i] == NOCARD {
				for jj := 0; jj < DC; jj++ {
					check = (*r)[jj] != NOCARD && !inNewPos((*r)[jj], *newpos) && !checkInMask(tabMask, (*r)[jj]) && getASuit(arrayOfSuit, i, (*r)[jj]) == NOVALUESUIT && getASuit(arrayOfSuit, jj, remain[k]) == NOVALUESUIT
					if check {
						(*r)[i] = (*r)[jj]
						(*r)[jj] = remain[k]
						*newpos = append(*newpos, remain[k])
						success = true
						break
					}
				}
			}
		}
		if !success {
			break
		}
	}

	return success
}
func checkHandPoint(content, c []int) bool {
	result := true
	ch := boardValueToArray(content)
	if ch[0] < c[0] || ch[0] > c[1] || ch[1] < c[2] || ch[1] > c[3] || ch[2] < c[4] || ch[2] > c[5] || ch[3] < c[6] || ch[3] > c[7] {
		return false
	}
	return result
}

func putDataInDist(sh ShuffleInterface, content, tabMask []int, r *[]int) {
	j := 0
	tab := shuffle(sh, content)
	for i := 0; i < DC; i++ {
		if tabMask[i] == NOCARD {
			(*r)[i] = tab[j]
			j++
		} else {
			(*r)[i] = tabMask[i]
		}

	}

}

//shuffleRemainHands ...
func shuffleRemainHands(sh ShuffleInterface, arrayOfSuit, arrayOfPoint []int, tabMask []int) ([]int, error) {
	var remain, newpos []int
	var aSuit, j int
	var err error

	r := make([]int, DC)
	content := delta(initHand, tabMask)
	// Simple case for fixed cards
	if arrayOfSuit == nil {
		if arrayOfPoint == nil {
			putDataInDist(sh, content, tabMask, &r)
		} else {
			successPoint := false
			for i := 0; i < MAXTRY; i++ {
				putDataInDist(sh, content, tabMask, &r)
				if checkHandPoint(r, arrayOfPoint) {
					successPoint = true
					break
				}
			}
			if !successPoint {
				err = fmt.Errorf("Point-" + ErrMsg["Error_shuffle_remain_hands"])
				return nil, err
			}
		}
	} else { // More complicated for suits fixed
		var success bool
		var tab []int
		for m := 0; m < MAXTRY; m++ {
			tab = shuffle(sh, content)
			j = 0
			remain = nil
			newpos = nil
			for i := 0; i < DC; i++ {
				r[i] = NOCARD
			}
			for i := 0; i < DC; i++ {
				if tabMask[i] == NOCARD {
					aSuit = getASuit(arrayOfSuit, i, tab[j])
					if aSuit == NOVALUESUIT {
						r[i] = tab[j]
						j++
					} else {
						remain = append(remain, tab[j])
						j++
					}
				} else {
					r[i] = tabMask[i]
				}
			}
			success = advanceCheck(tabMask, remain, arrayOfSuit, &r, &newpos)
			if success {
				break
			}
		}
		if !success {
			err = fmt.Errorf(ErrMsg["Error_shuffle_remain_hands"])
			return nil, err
		}
	}
	return r, nil
}

func boardValueToArray(board []int) []int {
	var r []int
	v := 0
	for i := 0; i < HC; i++ {
		v = 0
		for j := 0; j < HCC; j++ {
			v += cardLevel(board[i*HCC+j])
		}
		r = append(r, v)
	}
	return r
}

func rotateStrArray(r *[]string) {
	m := len(*r)
	itemp := (*r)[m-1]
	for i := m - 2; i >= 0; i-- {
		(*r)[i+1] = (*r)[i]
	}
	(*r)[0] = itemp
}

func getIntFirstHand(s string) int {
	if s == "E" {
		return 1
	}
	if s == "S" {
		return 2
	}
	if s == "W" {
		return 3
	}
	return 0
}
func getFirstHandFromPBN(pbn string) (int, string) {
	r := strings.Split(pbn, FIRSTHANDSEPARATOR)
	if len(r) != 2 {
		return 0, pbn
	}
	return getIntFirstHand(r[0]), r[1]
}

func rotateMask(pbn string) string {
	k, mask := getFirstHandFromPBN(pbn)
	tmask := strings.Split(mask, SPACE)
	for i := 0; i < k; i++ {
		rotateStrArray(&tmask)
	}
	mask = ""
	for i := 0; i < HC; i++ {
		mask += tmask[i] + SPACE
	}
	mask = strings.TrimSpace(mask)
	return mask
}

//PbnAndIndexGenerateFromMask API Web...
func PbnAndIndexGenerateFromMask(sh ShuffleInterface, arrayOfSuit []int, pbn string) (string, error) {
	pbn = strings.ToUpper(pbn)
	mask := rotateMask(pbn)
	ar, err := shuffleHand(sh, arrayOfSuit, mask)
	if err != nil {
		return "", err
	}
	ar = sortCards(ar)
	s := toPbn(ar)
	v := boardValueToArray(ar)
	result := ResultHTTP{Index: ArrayToIndex(ar), Pbn: s, NorthPoint: v[0], EastPoint: v[1], SouthPoint: v[2], WestPoint: v[3]}
	r, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(r), err
}

//PbnGenerateFromJSONIndex ...
func PbnGenerateFromJSONIndex(js string) (string, error) {
	var parameters jsonIndexWeb
	var content []int
	err := json.Unmarshal([]byte(js), &parameters)
	if err != nil {
		return "", err
	}
	content = IndexToArray(parameters.Index)
	v := boardValueToArray(content)
	r := toPbn(content)
	result := ResultHTTP{Index: parameters.Index, Pbn: r, NorthPoint: v[0], EastPoint: v[1], SouthPoint: v[2], WestPoint: v[3]}
	s, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(s), err
}

func getPbnData(t string, d Board) (string, error) {
	var index, pbn, r string
	var indexPbn ResultHTTP
	err := boardCheck(d)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal([]byte(t), &indexPbn)
	if err != nil {
		return "", err
	}
	index = indexPbn.Index
	pbn = indexPbn.Pbn
	if d.Comment == "" {
		r = "% Index: " + index + "\n"
	} else {
		d.Comment = strings.Replace(d.Comment, INDEXVALUE, index, -1)
		r = "% " + d.Comment + "\n"
	}
	r += "[Dealer \"" + d.Dealer + "\"]\n"
	r += "[Vulnerable \"" + d.Vulnerable + "\"]\n"
	r += "[Deal \"" + firstDefaultHand + ":" + pbn + "\"]"

	return r, nil
}

//PbnSuitDataGenerateFromJSON ...
func PbnSuitDataGenerateFromJSON(sh ShuffleInterface, js string) (string, error) {
	var err error
	var parameters JSONPbnToString
	var d Board
	var sx, r, t string
	err = json.Unmarshal([]byte(js), &parameters)
	if err != nil {
		return "", err
	}
	d.Dealer = parameters.Dealer
	d.Vulnerable = parameters.Vulnerable
	d.Comment = parameters.Comment
	mask := parameters.Mask
	if parameters.Count > MaxPbnGeneration {
		parameters.Count = MaxPbnGeneration
	}
	for i := 0; i < parameters.Count; i++ {
		t, err = PbnAndIndexGenerateFromSuits(sh, mask)
		if err != nil {
			return "", err
		}
		r, err = getPbnData(t, d)
		if err != nil {
			return "", err
		}
		sx += r + "\n\n"
	}
	return sx, err

}

//PbnPointDataGenerateFromJSON ...
func PbnPointDataGenerateFromJSON(sh ShuffleInterface, js string) (string, error) {
	var err error
	var parameters JSONPbnToString
	var d Board
	var sx, r, t string
	err = json.Unmarshal([]byte(js), &parameters)
	if err != nil {
		return "", err
	}
	d.Dealer = parameters.Dealer
	d.Vulnerable = parameters.Vulnerable
	d.Comment = parameters.Comment
	mask := parameters.Mask
	if parameters.Count > MaxPbnGeneration {
		parameters.Count = MaxPbnGeneration
	}
	for i := 0; i < parameters.Count; i++ {
		t, err = GetPbnHandsFromPoints(sh, mask)
		if err != nil {
			return "", err
		}
		r, err = getPbnData(t, d)
		if err != nil {
			return "", err
		}
		sx += r + "\n\n"
	}
	return sx, err

}

//PbnDataGenerateFromJSON ...
func PbnDataGenerateFromJSON(sh ShuffleInterface, js string) (string, error) {
	var err error
	var parameters JSONPbnToString
	var d Board
	var s, t, r string
	err = json.Unmarshal([]byte(js), &parameters)
	if err != nil {
		return "", err
	}
	d.Dealer = parameters.Dealer
	d.Vulnerable = parameters.Vulnerable
	d.Comment = parameters.Comment
	mask := parameters.Mask
	if parameters.Count > MaxPbnGeneration {
		parameters.Count = MaxPbnGeneration
	}
	for i := 0; i < parameters.Count; i++ {
		t, err = PbnAndIndexGenerateFromMask(sh, nil, mask)
		if err != nil {
			return "", err
		}
		r, err = getPbnData(t, d)
		if err != nil {
			return "", err
		}
		s += r + "\n\n"
	}
	return s, err
}

func checkVulnerable(s string) bool {
	result := false
	for _, v := range Vulnerable {
		if v == s {
			return true
		}
	}
	return result
}
func checkDealer(s string) bool {
	result := false
	for _, v := range Position {
		if v == s {
			return true
		}
	}
	return result
}

func boardCheck(d Board) error {
	var err error
	if !checkDealer(d.Dealer) {
		err = fmt.Errorf((ErrMsg["err_dealer"]))
		return err
	}
	if !checkVulnerable(d.Vulnerable) {
		err = fmt.Errorf(ErrMsg["err_vuln"])
		return err
	}
	return nil
}

func maskCheck(mask string) error {
	var err error
	var str string
	var tab []string
	s := strings.Split(mask, SPACE)
	if len(s) != HC {
		err = fmt.Errorf(ErrMsg["4_hands"])
		return err
	}
	for i := 0; i < HC; i++ {
		tab = strings.Split(s[i], POINT)
		if len(tab) != HC {
			err = fmt.Errorf(Position[i] + ErrMsg["4_suits"])
			return err
		}
		if s[i] != EMPTYHAND {
			str = strings.Replace(s[i], POINT, "", -1)
			if len(str) > HCC {
				err = fmt.Errorf(Position[i] + ErrMsg["13_cards"])
				return err
			}
		}
	}
	return err
}

func getHandFromDist(index int) int {
	return index / HCC
}

func valueAndSuitToCard(value, suit int) int {
	return (value << 2) + suit
}

//convertIndexArrayToDist ...
func convertIndexArrayToDist(index []int) []int {
	var suit, hand, value, cardValue int
	r := make([]int, DC)
	tabHand := make([]int, HC)
	reverse := make([]int, DC)
	for i := 0; i < DC; i++ {
		hand = index[i]
		value = i % HCC
		suit = i / HCC
		cardValue = valueAndSuitToCard(value, suit)
		r[tabHand[hand]+hand*HCC] = cardValue
		tabHand[hand]++
	}
	for i := 0; i < HC; i++ {
		for j := 0; j < HCC; j++ {
			reverse[HCC*i+HCC-j-1] = r[HCC*i+j]
		}
	}
	return reverse
}

//convertDistToIndexArray ...
func convertDistToIndexArray(content []int, index *[DC]int) {
	var suit, height, v int
	for i := 0; i < len(content); i++ {
		v = content[i]
		suit = cardSuitInt(v)
		height = cardValueInt(v)
		index[suit*HCC+height] = getHandFromDist(i)
	}
}

// From java code Copyright (@)1999, Thomas Andrews
//http://bridge.thomasoandrews.com/impossible/
// Free for non-commercial use

func fraction(val *big.Int, num, den int) *big.Int {
	numer := int64(num)
	denom := int64(den)
	v := new(big.Int)
	n := big.NewInt(numer)
	d := big.NewInt(denom)
	v.Mul(val, n)
	v.Div(v, d)
	return v
}

func pages() *big.Int {
	v := new(big.Int)
	v.SetString((NbDist), 10)
	return v
}

//ArrayToIndex Get Richard Pavileck index from distribution
func ArrayToIndex(content []int) string {
	var r string
	var cardsNeeded = [HC]int{HCC, HCC, HCC, HCC}
	var hand, skipped, goesTo int
	var index [DC]int
	convertDistToIndexArray(content, &index)
	width := pages()
	minimum := big.NewInt(0)
	for c := DC; c > 0; c-- {
		hand = 0
		skipped = 0
		goesTo = index[c-1]
		for hand < goesTo {
			skipped = skipped + cardsNeeded[hand]
			hand++
		}
		minimum.Add(minimum, fraction(width, skipped, c))
		width = fraction(width, cardsNeeded[goesTo], c)
		cardsNeeded[goesTo]--
	}
	r = fmt.Sprintf("%v", minimum)
	return r
}

//IndexToArray Get distribution from Richard Pavileck index
func IndexToArray(sIndex string) []int {
	result := make([]int, DC)
	var cardsNeeded = [HC]int{HCC, HCC, HCC, HCC}
	var x *big.Int
	index := new(big.Int)
	index.SetString((sIndex), 10)
	k := pages()
	for c := DC; c > 0; c-- {
		for hand := 0; hand < 4; hand++ {
			x = fraction(k, cardsNeeded[hand], c)
			if index.Cmp(x) < 0 {
				result[c-1] = hand
				k = x
				cardsNeeded[hand]--
				break
			} else {
				index.Sub(index, x)
			}
		}
	}
	return convertIndexArrayToDist(result)
}

// End of Copyright (@)1999, Thomas Andrews
