// History: Dec 27 13 tcolar Creation
// Chapter 1 Exercises

package algo4ed

import (
  . "github.com/smartystreets/goconvey/convey"
  "fmt"
  "log"
  "testing"
)

// 1.1.3 : Method that cheks if 3 ints are equal
func Test_1_1_3(t *testing.T){
  Convey("3 Equal Ints", t, func() {
    So(threeInt(2,3,4), ShouldEqual, false)
    So(threeInt(2,2,4), ShouldEqual, false)
    So(threeInt(2,2,2), ShouldEqual, true)
    So(threeInt(2,2,-2), ShouldEqual, false)
  })
}

func threeInt(a int, b int, c int) bool {
  return a == b && b == c
}

// 1.1.5 : Return true ony if x and y strictly between 0 and 1
// Note : Is "strictly between" inclusive or not ?
func Test_1_1_5(t *testing.T){
  Convey("Strictly betwen 0 and 1", t, func() {
    So(xyBetween0And1(0.0,  1.0), ShouldEqual, true)
    So(xyBetween0And1(-0.1, 0.5), ShouldEqual, false)
    So(xyBetween0And1(0.3, 0.9), ShouldEqual, true)
    So(xyBetween0And1(0.7, 999), ShouldEqual, false)
  })
}

func xyBetween0And1(x float64, y float64) bool {
  return x >= 0.0 && x <= 1.0 && y >= 0.0 && y <= 1.0
}

// 1.1.13 : Transpose a matrix
func Test_1_1_13(t *testing.T){
  array := [][]int{
    []int{1,2,3,4},
    []int{11,12,13,14},
    []int{21,22,23,24},
  }
  transpose(array)
}

// Print matrix transposition
func transpose(array [][]int) {
  for i := 0; i != len(array[0]); i++{
    log.Print("--")
    for j := 0; j != len(array); j++{
      log.Print(array[j][i])
    }
  }
}

// 11.14 : Find largest int <= base2 log n
func Test_1_1_14(t *testing.T){
  Convey("Lg", t, func() {
    So(lg(1), ShouldEqual, 0)
    So(lg(2), ShouldEqual, 1)
    So(lg(3), ShouldEqual, 1)
    So(lg(4), ShouldEqual, 2)
    So(lg(7), ShouldEqual, 2)
    So(lg(9), ShouldEqual, 3)
  })
}

func lg(n int) int {
  i := 0
  for n > 1 {
    n = n >> 1
    i++
  }
  return i
}

// 1.1.15 : return array of length M where ith entry is i count in a
// if all values  in a are 0<x<M-1 then sum of values in a should = len(a)
func Test_1_1_15(t *testing.T){
  a := []int{0,3,3,4,4,4,5}
  h := histogram(a)
  Convey("Histogram", t, func() {
    So(len(h), ShouldEqual, 7)
    So(h[0], ShouldEqual, 1)
    So(h[1], ShouldEqual, 0)
    So(h[2], ShouldEqual, 0)
    So(h[3], ShouldEqual, 2)
    So(h[4], ShouldEqual, 3)
    So(h[5], ShouldEqual, 1)
    So(h[6], ShouldEqual, 0)
    So(arraySum(h), ShouldEqual, len(a))
  })
}

func histogram(a []int) []int {
  M := len(a)
  result := make([]int, M)
  for _, n := range a {
    result[n] += 1
  }
  return result
}

// Sum of elements in array
func arraySum(a []int) (sum int) {
  for _, n := range a {
    sum += n
  }
  return sum
}

// 11.16 : Recursion exercise 1
func exR1(n int) string {
  if n <= 0 { return "" }
  return fmt.Sprintf("%s%d%s%d", exR1(n - 3), n ,exR1(n - 2), n)
}


