package main

import (
	"fmt"
)

func main() {
	var n, j, k int
	var v, prd, max int64
	var ar []int64
	fmt.Scanln(&n)
	for i := 0; i != n; i++ {
		fmt.Scan(&v)
		ar = append(ar, v)
	}


	ln := len(ar)
	var jdone, kdone bool
	for i := 0; i != n; i++ {
		jdone = false
		kdone = false
		if i > 0 {
           		// if ar[i] >= ar[i-1] , we can optimize
			if ar[i] >= ar[i-1] {
				if ar[i] < ar[j] {
					// a[i] >= a[i-1] < a[j] -> same j
					jdone = true
				}
				// else : start looking before j
				if ar[i] < ar[k] {
					// a[i] < a[i-1] < a[k] -> same k
					kdone = true
				}
				// else : start looking after k
			} else {
				j = i - 1
				k = i + 1
			}
		}
		for !jdone && j >= 0 && ar[j] <= ar[i] {
			j--
		}
		if j < 0 { // j not found -> 0
			j = 0
            continue // Optimization : x*0 always 0
		}
		for !kdone && k < ln && ar[k] <= ar[i] {
			k++
		}
		if k >= ln { // k not found -> 0
			k = ln -1
			continue // Optimization : x*0 always 0
		}
		//fmt.Printf("%d : %d x %d\n", i+1, j+1, k+1)
		prd = int64(j+1) * int64(k+1)
		if prd > max {
			max = prd
		}
	}
	fmt.Println(max)
}

/*


Problem Statement

You are given a list of N numbers a1,a2,…,an. For each i (1≤i≤N), define LEFT(i) to be the nearest index j before i such that aj>ai. Note that if j does not exist, we should consider LEFT(i) to be 0. Similarly, define RIGHT(i) to be the nearest index k after i such that ak>ai. Note that if k does not exist, we should consider RIGHT(i) to be 0.

Define INDEXPRODUCT(i) as the product of LEFT(i) and RIGHT(i). Find the maximum INDEXPRODUCT(i) among all 1≤i≤N.

Input Format

The first line contains an integer N, the number of integers. The next line contains the N integers.

Constraints:

1≤N≤105

Each of the N integers will range from 0 to 109

Output Format

Output the maximum INDEXPRODUCT among all indices from 1 to N.

Sample Input

5
5 4 3 4 5

Sample Output

8

Explanation

We can compute the following:

INDEXPRODUCT(1)=0
INDEXPRODUCT(2)=1×5=5
INDEXPRODUCT(3)=2×4=8
INDEXPRODUCT(4)=1×5=5
INDEXPRODUCT(5)=0

The largest of these is 8, so it is the answer.

*/
