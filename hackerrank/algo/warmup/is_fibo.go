package main

import (
	"fmt"
)

var fibos []int64 = []int64{0, 1, 2}
var highest int64 = 2

func isFibo(n int64) bool {
	if n < 0 {
		return false
	}
	for highest <= n {
		highest = fibos[len(fibos)-1] + fibos[len(fibos)-2]
		fibos = append(fibos, highest)
	}
	if highest == n {
		return true
	}
	for _, v := range fibos {
		if v == n {
			return true
		}
	}
	return false
}

func main() {
	var t int
	var n int64
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		if isFibo(n) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}
	}
}

/*
Problem Statement

You are given an integer, N. Write a program to determine if N is an element of the Fibonacci sequence.

The first few elements of the Fibonacci sequence are 0,1,1,2,3,5,8,13,⋯. A Fibonacci sequence is one where every element is a sum of the previous two elements in the sequence. The first two elements are 0 and 1.

Formally:
fib0fib1⋮fibn=0=1=fibn−1+fibn−2∀n>1

Input Format
The first line contains T, number of test cases.
T lines follow. Each line contains an integer N.

Output Format
Display IsFibo if N is a Fibonacci number and IsNotFibo if it is not. The output for each test case should be displayed in a new line.

Constraints
1≤T≤105
1≤N≤1010

Sample Input

3
5
7
8

Sample Output

IsFibo
IsNotFibo
IsFibo

Explanation
5 is a Fibonacci number given by fib5=3+2
7 is not a Fibonacci number
8 is a Fibonacci number given by fib6=5+3

Time Limit
Time limit for this challenge is given here.
*/
