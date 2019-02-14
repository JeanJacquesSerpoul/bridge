package tools

import "sort"

// GenDistWithPointToFile will generate distribution
func GenDistWithPointToFile(filepath string, kmin, kmax int) error {
	var tp, p []int
	var td, d []string
	var list []listData
	for i := kmin; i <= kmax; i++ {
		p, d = genDistWithPoint(i)
		tp = appendArray(tp, p)
		td = appendStrArray(td, d)
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
	err := saveGenDistWithPointToTXT(filepath, tp, td)
	return err
}
