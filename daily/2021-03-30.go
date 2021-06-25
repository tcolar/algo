/*
An XOR linked list is a more memory efficient doubly linked list. Instead of each node holding next and prev fields, 
it holds a field named both, which is an XOR of the next node and the previous node. Implement an XOR linked list; 
it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

If using a language that has no pointers (such as Python), you can assume you have access to get_pointer and 
dereference_pointer functions that converts between nodes and memory addresses.
*/

package main

import "fmt"

type XorLinkedList struct {
  root *XorNode;
  end *XorNode;
}

func(l *XorLinkedList) add(data string) {
    if (l.end==nil) {
      l.end = &XorNode{
        both: 0,
        data: data,
      };
    } else {
      ptr := l.end.ptr();
      l.end = &XorNode{
        both: ptr,
        data: data,
      };
    }
  }

func(l *XorLinkedList) get(index int) *XorNode {
  return l.end
}

type XorNode struct {
  both int;
  data string; 
}

func (n *XorNode) ptr() int {
  return 1233// (uint64)(unsafe.Pointer(n));
}

func main(){
  l:=XorLinkedList{}
  l.add("foo")
  l.add("baz")
  l.add("bar")
  fmt.Println(l.get(0))
  fmt.Println(l.get(1))
  fmt.Println(l.get(2))
}
