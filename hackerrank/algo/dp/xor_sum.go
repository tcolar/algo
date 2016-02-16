package main

import (
	"fmt"
	"strconv"
)

// 11101001011000001101000110011

func main() {
	var sa, sb string
	var a, b int64
	//var result int
	fmt.Scanln(&sa)
	fmt.Scanln(&sb)
	fmt.Printf("%s %s\n", sa, sb)
	a, err := strconv.ParseInt(sa, 2, 32)
	if err != nil {
		panic(err)
	}
	b, err = strconv.ParseInt(sb, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %d\n", a, b)
	for i := 0; i != 10; i++ {
		mask := int64(1) << uint64(i)
		bit := (a & mask) ^ (b & mask)
		bit = 0 ^ 1
		fmt.Println(bit)
	}

}

/*
Problem Statement

You are given two positive integers a and b in binary representation. You should find the following sum modulo 109+7:

∑i=0314159(a xor (b shl i))

where operation xor means exclusive OR operation, operation shl means binary shift to the left.

Please note, that we consider ideal model of binary integers. That is there is infinite number of bits in each number, and there are no disappearings (or cyclic shifts) of bits.

Input Format

The first line contains number a (1≤a<2105) in binary representation. The second line contains number b (1≤b<2105) in the same format. All the numbers do not contain leading zeros.

Output Format

Output a single integer − the required sum modulo 109+7.

Sample Input

10
1010

Sample Output

489429555
*/
