package distribution

// Board ...
type Board struct {
	Content    []int  `json:"content"`
	Dealer     string `json:"dealer,omitempty"`
	Vulnerable string `json:"vulnerable,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type listHand struct {
	value     int
	sortOrder int
}

//handSuit Array of number of suit
type handSuit [HC][HCC]int

// JSONPbnToString ...
type JSONPbnToString struct {
	Count      int    `json:"count,omitempty"`
	Mask       string `json:"mask,omitempty"`
	Vulnerable string `json:"vulnerable,omitempty"`
	Dealer     string `json:"dealer,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type jsonIndexWeb struct {
	Index string `json:"index,omitempty"`
}

// ResultHTTP ...
type ResultHTTP struct {
	Index      string `json:"index,omitempty"`
	Pbn        string `json:"pbn,omitempty"`
	NorthPoint int    `json:"np"`
	EastPoint  int    `json:"ep"`
	SouthPoint int    `json:"sp"`
	WestPoint  int    `json:"wp"`
}

//ShuffleInterface ...
type ShuffleInterface interface {
	fYShuffle(int) []int
}

//Random ...
type Random struct {
}

type pointStruct struct {
	Orientation int
	MinPoints   int
	MaxPoints   int
}

type listData struct {
	points int
	dist   string
}
