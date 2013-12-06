// History: Dec 01 13 tcolar Creation

package algo

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	var mutex = sync.Mutex{}
	// this will be incremented using mutex
	var total = struct {
		Total int64
	}{Total: 7}
	// this will be updated using atomic.AddInt64
	var total2 int64 = 7

	// Start a hundred routines
	for i := 0; i < 50; i++ {
		go func() {
			// Each updating the total values a 100 times
			for j := 0; j != 50; j++ {
				toAdd := int64(i)
				// adding or removing  'j' to total if eve or odd
				if j%2 == 0 {
					toAdd *= -1
				}
				// using the mutex
				mutex.Lock()

				// Doing the addition in such a way that it makes it much more likely
				// that their would be a race condition (when not using a mutex)
				// Not so easy to make happen somehow.
				t := total.Total
				time.Sleep(time.Duration(rand.Int()%1000) * time.Microsecond)
				t += toAdd
				total.Total = t

				mutex.Unlock()

				// using atomic
				atomic.AddInt64(&total2, toAdd)

				// Keep scheduler happy
				// Note: Possibly no longer needed as of Go 1.2
				runtime.Gosched()
			}
		}()
	}

	// Give some time to compute
	time.Sleep(5 * time.Second)

	// Both totals should still have the start value by the time it ends
	if total.Total != 7 || total2 != 7 {
		log.Printf("Total: %d", total)
		log.Printf("Total2: %d", total2)
		log.Panic("Unexpected values")
	}
}
