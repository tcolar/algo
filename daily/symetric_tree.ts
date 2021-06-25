//
// Binary trees are already defined with this interface:
// class Tree<T> {
//   value: T;
//   left: Tree<T>;
//   right: Tree<T>;
//
//   constructor(value: T) {
//     this.value = value;
//     this.left = null;
//     this.right = null;
//   }
// }
function isTreeSymmetric(t: Tree<number>): boolean {
  return compareNodes(t?.right, t?.left);
}

function compareNodes(left: Tree<number>, right: Tree<number>): boolean {
  //console.log(left, right)
  if (!left && !right) return true;
  if (left && !right) return false;
  if (!left && right) return false;
  if (left.value !== right.value) return false;
  return compareNodes(left.left, right.right) && compareNodes(left.right, right.left);
}
