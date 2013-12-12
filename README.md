algo
====

This is nothing useful but myself doing some algo exercises for practice since this
is the kind of stuff asked in interviews but not necessarely used everyday in the field.

* List
* Stack
* Heap
- Tree
- Graph
  - Dijkstra
- Hash

- Sorts / Big O
  - Bubble sort -> On2 / On2 -> swap with neihboor again and again -> useless
  - Insertion sort -> On2 / On2 from 0 to n take lement and move to proper place in beginnging of sorted part
  - Selection sort -> On2 / On2. from o to n: find smallest element and move it to n
  * Quicksort -> OnLogn / On2. Pivot, create low, high lists, recurse and then merge back (low+pivot+high)
  - Merge sort -> Onlogn / Onlogn can parallel. Good for linked list, slow access ()
  * Heap sort -> OnLogn / OnLogn. Build heap and then traverse
  - Radix / Bucket sort -> Onk. Best for known & limited data ranges (ex: Ascii)
  - Introsort -> Quicksort / heapsort combo

- Search
  * Binary search

- Concurrency
  * Go mutex
  * Go routines / channels

* Design Patterns
  * Strategy -> Multiple implemtations behing interface decide which to use at runtime
  * Adater -> Basically a wrapper to adapt a model/API to another
  * Factory -> Create specific instances of an interface depending of params
  * Decorator -> Wrap model and "decorate" it with extra or modified features
  * Singleton -> Single instance
  * Iterator -> ... easy :)
  * Builder ->  Construct object in several steps (add behavior). DSL's
  * Memento -> Typically do / undo ()~ stack

