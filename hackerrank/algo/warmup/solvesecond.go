package main

import "fmt"

func main() {
	var n int
	var a int
	var b int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var sum int
		fmt.Scan(&a, &b)
		sum = a + b
		fmt.Println(uint(sum))
	}
}

/*
Problem Statement

You learnt about STDIN and STDOUT in Solve me first.

This is the second challenge in the introduction series. The purpose of this challenge is to give you a working I/O template in your preferred language. It includes scanning two space-separated integers from STDIN in a loop over T lines, calling a function, returning a value, and printing it to STDOUT.

A pseudo code looks like the following:

read T
loop from 1 to T
    read A and B
        compute the sum
	    print value in a newline
	    end loop

	    The task is to scan two numbers from STDIN, and print the sum A+B on STDOUT. The code has already been provided for most of the popular languages. This is primarily for you to read and inspect how the IO is handled.

	    Note: The code has been saved in a template, which you can submit if you want. Or, you may try rewriting it and building it up from scratch.

	    Input Format
	    (This section specifies the Input Format.)
	    The first line contains T (number of test cases) followed by T lines
	    Each line contains A and B, separated by a space.

	    As you can see that we have provided in advance the number of lines, we discourage the use of scanning till EOF as not every language has an easy way to handle that. In fact, every HackerRank challenge is designed in such a way that multitests begin with a T line to indicate the number of lines.

	    Output Format
	    (This section specifies the Output Format.)
	    An integer that denotes Sum (A+B) printed on new line for every testcase.

	    Constraints
	    (This section tells what input you can expect. You can freely assume that the input will remain within the boundaries specified.)
	    1≤T,A,B≤1000

	    Sample Input

	    2
	    2 3
	    3 7

	    Sample Output

	    5
	    10

	    The above sample should be taken seriously. 2 in the first line describes how many lines will follow, and your test cases are 2, 3 and 3, 7 in two separate lines. Your output should be 5 and 10 printed on two separate lines. If you print extra lines or "The answer is: 5", any such extra characters in output will result in a Wrong Answer, as the judging is done using diff checker.

*/
