package tools

import (
	"io"
	"sort"
)

// GenDistWithPointToFile will generate distribution
func GenDistWithPointToFile(w io.Writer, kmin, kmax int) error {
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
	err := saveGenDistWithPointToTXT(w, tp, td)
	return err
}
