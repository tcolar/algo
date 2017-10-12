/*

A third-party API that we're using has a paginated API. It returns results in chunks of 10. Assume that we've got the following API created for us by that third party API:

def fetch_page(placeholder: str=None):
   '''Return the next page of results.  If placeholder == None, starts from the
      beginning.  Otherwise, fetches the next 10 records after the last
      placeholder.

      returns:
        {
          "results": [...],
          "placeholder": "abc123"
        }
    '''
We don't think that API is very useful, and would prefer the following:

class ResultFetcher(object):
    def fetch(n):
Your task will be to create a new class that encapsulates the calls to fetch_page into a new function that presents an API that allows fetching N results at a time, keeping it's place.

*/

// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	r := ResultFetcher{}
	fmt.Println(r.Fetch(10))
	fmt.Println(r.Fetch(6))
	fmt.Println(r.Fetch(4))
	fmt.Println(r.Fetch(10))
}

type ResultFetcher struct {
	placeHolder *string
	pending     []int
}

func (r *ResultFetcher) Fetch(n int) []int {
	for {
		if len(r.pending) >= n { // enough results
			results := r.pending[:n]
			r.pending = r.pending[n:]
			return results
		}
		// get more results
		newResults := []int{}
		r.placeHolder, newResults = r.fetchPage(r.placeHolder)
		r.pending = append(r.pending, newResults...)
		if r.placeHolder == nil && len(r.pending) < n {
			// no more, return all we got
			results := r.pending
			r.pending = []int{}
			return results
		}
	}
}

var state int
var meh = ""

func (r *ResultFetcher) fetchPage(placeHolder *string) (*string, []int) {
	state += 10
	if state >= 30 {
		return nil, []int{1, 2, 3}
	}
	return &meh, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}
