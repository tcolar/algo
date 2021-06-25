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
function isListPalindrome(l: ListNode<number>): boolean {
  let node = l;
  if (!node) return true;
  let count = 0;
  // find length of list, 0(n) time
  while (node) {
    count++;
    node = node.next;
  }
  node = l;
  let prevNode = l;
  // reverse second part of list in place, O(n) time
  // seems bad to mutate, but only way I can think off to use constant space
  for (var i = 0; i < count; i++) {
    let curNode = node;
    node = node.next;
    if (i > Math.ceil(count / 2)) {
      curNode.next = prevNode;
    }
    if (curNode) prevNode = curNode;
  }
  let tail = prevNode;
  node = l;
  var i = 0;
  // compare head pointer and tail pointer until we reach middle, 0(n) time
  while (i < Math.floor(count / 2)) {
    if (node.value !== tail.value) return false;
    node = node.next;
    tail = tail.next;
    i++;
  }
  return true;
}

// recursion would use space, so not o(1) space
/*function compare(head: ListNode<number>, tail:ListNode<number>) :boolean { 
    if(!tail){
        return true
    }
    if(!compare(head, tail.next)) return false
    console.log(head.value, tail.value)
    if (head.value !== tail.value) return false
    head = head.next
    return true
}*/
