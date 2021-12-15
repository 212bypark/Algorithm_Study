package twoNumSum_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Alogorithm_Study/category/Arrays/twoNumSum"
)

type TestData struct {
	testArray []int
	targetNum int
}

var testData = []TestData{
	{testArray1[:], 17},
	{testArray2[:], 18},
	{testArray3[:], 10},
	{testArray4[:], 10},
	{testArray5[:], 5},
	{testArray6[:], 3},
	{testArray7[:], 17},
	{testArray8[:], 141},
	{testArray9[:], 1234},
}

var testArray1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var testArray2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
var testArray3 = [...]int{3, 5, -4, 8, 11, 1, -1, 6}
var testArray4 = [...]int{4, 6}
var testArray5 = [...]int{4, 6, 1}
var testArray6 = [...]int{4, 6, 1, -3}
var testArray7 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var testArray8 = rand.Perm(100)
var testArray9 = rand.Perm(1000)

func TestTwoNumSum(t *testing.T) {
	i := 1
	for _, d := range testData {
		fmt.Printf("Test: %d\n", i)
		sT1 := time.Now()
		twoNumSum.QuardraticTwoNumberSum(d.testArray, d.targetNum)
		eT1 := time.Since(sT1)
		fmt.Printf("Quardratic: %s", eT1)
		fmt.Print("  ")
		sT2 := time.Now()
		twoNumSum.LinearTwoNumberSum(d.testArray, d.targetNum)
		eT2 := time.Since(sT2)
		fmt.Printf("Linear: %s", eT2)
		fmt.Print("  ")
		sT3 := time.Now()
		twoNumSum.LogLinearTwoNumberSum(d.testArray, d.targetNum)
		eT3 := time.Since(sT3)
		fmt.Printf("LogLinear: %s", eT3)
		fmt.Println()
		i++
	}
}
