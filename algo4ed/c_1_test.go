// History: Dec 27 13 tcolar Creation
// Chapter 1 Exercises

package algo4ed

import (
  . "github.com/smartystreets/goconvey/convey"
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

// Return true ony if x and y strictly between 0 and 1
// Note : Is "strictly between" inclusive or not ?
func Test_1_1_5(t *testing.T){
  Convey("Strictly betwen 0 and 1", t, func() {
    So(xyBetween0And1(0.0,  1.0), ShouldEqual, true)
    So(xyBetween0And1(-0.1, 0.5), ShouldEqual, false)
    So(xyBetween0And1(0.3, 0.9), ShouldEqual, true)
    So(xyBetween0And1(0.7, 999), ShouldEqual, false)
  })
}

func Test_1_1_13(t *testing.T){
  array := [][]int{
    []int{1,2,3,4},
    []int{11,12,13,14},
    []int{21,22,23,24},
  }
  transpose(array)
}

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

func threeInt(a int, b int, c int) bool {
  return a == b && b == c
}

func xyBetween0And1(x float64, y float64) bool {
  return x >= 0.0 && x <= 1.0 && y >= 0.0 && y <= 1.0
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

// Find largest int <= base2 log n
func lg(n int) int {
  i := 0
  for n > 1 {
    n = n >> 1
    i++
  }
  return i
}