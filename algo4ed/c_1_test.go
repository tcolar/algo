// History: Dec 27 13 tcolar Creation
// Chapter 1 Exercises

package algo4ed

import (
  . "github.com/smartystreets/goconvey/convey"
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

func threeInt(a int, b int, c int) bool {
  return a == b && b == c
}

func xyBetween0And1(x float64, y float64) bool {
  return x >= 0.0 && x <= 1.0 && y >= 0.0 && y <= 1.0
}