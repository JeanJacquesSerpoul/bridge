package tools

import (
	"io"
	"math"
	"strconv"
	"strings"
)

func appendArray(a, b []int) []int {
	r := a
	for i := 0; i < len(b); i++ {
		r = append(r, b[i])
	}
	return r
}
func cardValueInt(cardValue int) int { return cardValue >> 2 }

func cardLevel(cardValue int) int {
	if cardValue < 36 {
		return 0
	}
	v := cardValueInt(cardValue)
	return (v - 8)
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
			p = appendArray(p, s)
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

func saveGenDistWithPointToTXT(
	w io.Writer,
	points []int,
	list []string,
) error {
	for i, po := range points {
		if i > 0 {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if _, err := w.Write(
			[]byte(strconv.Itoa(po) + "\t" + list[i]),
		); err != nil {
			return err
		}
	}
	return nil
}
