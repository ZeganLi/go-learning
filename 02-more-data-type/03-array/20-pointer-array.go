package main

import "fmt"

type ArrayTest struct {
	ATA []string
	ATB *[]string
	ATC [][]string
	ATD *[][]string
	ATE []*[][]string
	ATF *[][][]string
}

func main() {
	arrayTest := ArrayTest{
		ATA: []string{"A", "T", "A"},
		ATB: &[]string{"A", "T", "B"},
		ATC: [][]string{{"A"}, {"T"}, {"C"}},
		ATD: &[][]string{{"A"}, {"T"}, {"D"}},
		ATE: []*[][]string{{{"A"}}, {{"T"}}, {{"E"}}},
		ATF: &[][][]string{{{"A"}}, {{"T"}}, {{"F"}}},
	}

	fmt.Println(arrayTest)
}
