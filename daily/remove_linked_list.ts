// Singly-linked lists are already defined with this interface:
// class ListNode<T> {
//   value: T;
//   next: ListNode<T>;
//
//   constructor(value: T) {
//     this.value = value;
//     this.next = null;
//   }
// }
//
function removeKFromList(l: ListNode<number>, k: number): ListNode<number> {
  let node = l;
  let root = l;
  let prev = null;
  while (node !== null) {
    if (node.value === k) {
      if (prev === null) {
        root = node.next;
      } else {
        prev.next = node.next;
      }
    } else {
      prev = node;
    }
    node = node.next;
  }
  return root;
}
