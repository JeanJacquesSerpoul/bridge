package tools

//ShuffleInterface ...
type ShuffleInterface interface {
	fYShuffle(int) []int
}

//Random ...
type Random struct {
}

type listData struct {
	points int
	dist   string
}
