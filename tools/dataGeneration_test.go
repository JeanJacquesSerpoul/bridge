package tools

import (
	"reflect"
	"testing"
)

var mockGenerelist = []int{36, 37, 36, 38, 37, 38, 36, 39, 37, 39, 38, 39, 36, 40, 37, 40, 38, 40, 39, 40, 36, 41, 37, 41, 38, 41, 39, 41, 40, 41, 36, 42, 37, 42, 38, 42, 39, 42, 40, 42, 41, 42, 36, 43, 37, 43, 38, 43, 39, 43, 40, 43, 41, 43, 42, 43, 36, 44, 37, 44, 38, 44, 39, 44, 40, 44, 41, 44, 42, 44, 43, 44, 36, 45, 37, 45, 38, 45, 39, 45, 40, 45, 41, 45, 42, 45, 43, 45, 44, 45, 36, 46, 37, 46, 38, 46, 39, 46, 40, 46, 41, 46, 42, 46, 43, 46, 44, 46, 45, 46, 36, 47, 37, 47, 38, 47, 39, 47, 40, 47, 41, 47, 42, 47, 43, 47, 44, 47, 45, 47, 46, 47, 36, 48, 37, 48, 38, 48, 39, 48, 40, 48, 41, 48, 42, 48, 43, 48, 44, 48, 45, 48, 46, 48, 47, 48, 36, 49, 37, 49, 38, 49, 39, 49, 40, 49, 41, 49, 42, 49, 43, 49, 44, 49, 45, 49, 46, 49, 47, 49, 48, 49, 36, 50, 37, 50, 38, 50, 39, 50, 40, 50, 41, 50, 42, 50, 43, 50, 44, 50, 45, 50, 46, 50, 47, 50, 48, 50, 49, 50, 36, 51, 37, 51, 38, 51, 39, 51, 40, 51, 41, 51, 42, 51, 43, 51, 44, 51, 45, 51, 46, 51, 47, 51, 48, 51, 49, 51, 50, 51}
var mockArrayFromGenereList = [][]int{{36, 37}, {36, 38}, {37, 38}, {36, 39}, {37, 39}, {38, 39}, {36, 40}, {37, 40}, {38, 40}, {39, 40}, {36, 41}, {37, 41}, {38, 41}, {39, 41}, {40, 41}, {36, 42}, {37, 42}, {38, 42}, {39, 42}, {40, 42}, {41, 42}, {36, 43}, {37, 43}, {38, 43}, {39, 43}, {40, 43}, {41, 43}, {42, 43}, {36, 44}, {37, 44}, {38, 44}, {39, 44}, {40, 44}, {41, 44}, {42, 44}, {43, 44}, {36, 45}, {37, 45}, {38, 45}, {39, 45}, {40, 45}, {41, 45}, {42, 45}, {43, 45}, {44, 45}, {36, 46}, {37, 46}, {38, 46}, {39, 46}, {40, 46}, {41, 46}, {42, 46}, {43, 46}, {44, 46}, {45, 46}, {36, 47}, {37, 47}, {38, 47}, {39, 47}, {40, 47}, {41, 47}, {42, 47}, {43, 47}, {44, 47}, {45, 47}, {46, 47}, {36, 48}, {37, 48}, {38, 48}, {39, 48}, {40, 48}, {41, 48}, {42, 48}, {43, 48}, {44, 48}, {45, 48}, {46, 48}, {47, 48}, {36, 49}, {37, 49}, {38, 49}, {39, 49}, {40, 49}, {41, 49}, {42, 49}, {43, 49}, {44, 49}, {45, 49}, {46, 49}, {47, 49}, {48, 49}, {36, 50}, {37, 50}, {38, 50}, {39, 50}, {40, 50}, {41, 50}, {42, 50}, {43, 50}, {44, 50}, {45, 50}, {46, 50}, {47, 50}, {48, 50}, {49, 50}, {36, 51}, {37, 51}, {38, 51}, {39, 51}, {40, 51}, {41, 51}, {42, 51}, {43, 51}, {44, 51}, {45, 51}, {46, 51}, {47, 51}, {48, 51}, {49, 51}, {50, 51}}
var mockMaskGenereList = []int{36, 37}

var mockMemoryGenereList = []string{"36 37", "36 38", "37 38", "36 39", "37 39", "38 39", "36 40", "37 40", "38 40", "39 40", "36 41", "37 41", "38 41", "39 41", "40 41", "36 42", "37 42", "38 42", "39 42", "40 42", "41 42", "36 43", "37 43", "38 43", "39 43", "40 43", "41 43", "42 43", "36 44", "37 44", "38 44", "39 44", "40 44", "41 44", "42 44", "43 44", "36 45", "37 45", "38 45", "39 45", "40 45", "41 45", "42 45", "43 45", "44 45", "36 46", "37 46", "38 46", "39 46", "40 46", "41 46", "42 46", "43 46", "44 46", "45 46", "36 47", "37 47", "38 47", "39 47", "40 47", "41 47", "42 47", "43 47", "44 47", "45 47", "46 47", "36 48", "37 48", "38 48", "39 48", "40 48", "41 48", "42 48", "43 48", "44 48", "45 48", "46 48", "47 48", "36 49", "37 49", "38 49", "39 49", "40 49", "41 49", "42 49", "43 49", "44 49", "45 49", "46 49", "47 49", "48 49", "36 50", "37 50", "38 50", "39 50", "40 50", "41 50", "42 50", "43 50", "44 50", "45 50", "46 50", "47 50", "48 50", "49 50", "36 51", "37 51", "38 51", "39 51", "40 51", "41 51", "42 51", "43 51", "44 51", "45 51", "46 51", "47 51", "48 51", "49 51", "50 51"}
var mockListpoint = []int{2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 4, 4, 4, 4, 5, 5, 5, 5, 6, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8}

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
func Test_genereList(t *testing.T) {
	type args struct {
		seq []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{cardsWithPoints, 2}, mockGenerelist},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genereList(tt.args.seq, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genereList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenereArrayFromList(t *testing.T) {
	type args struct {
		seq []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test1", args{cardsWithPoints, 2}, mockArrayFromGenereList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genereArrayFromList(tt.args.seq, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenereArrayFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genDistWithPoint(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name  string
		args  args
		want1 []int
		want2 []string
	}{
		{"Test1", args{2}, mockListpoint, mockMemoryGenereList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := genDistWithPoint(tt.args.k)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GenDistWithPoint() = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GenDistWithPoint() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

// use test to generate data.txt
func Test_genDistWithPointToFile(t *testing.T) {
	type args struct {
		filepath string
		kmin     int
		kmax     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test1", args{"data.txt", 1, 13}, false},
	}
	sh := new(fakeRandom)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := genDistWithPointToFile(sh, tt.args.filepath, tt.args.kmin, tt.args.kmax); (err != nil) != tt.wantErr {
				t.Errorf("genDistWithPointToFile() error = %v, wantErr %v", err, tt.wantErr)
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
		{"Test1", args{3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardLevel(tt.args.cardValue); got != tt.want {
				t.Errorf("cardValueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeStringToFile(t *testing.T) {
	type args struct {
		filepath string
		s        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test1", args{"t/t.txt", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeStringToFile(tt.args.filepath, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("writeStringToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
