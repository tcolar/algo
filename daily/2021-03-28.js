/*
Given an array of integers, find the first missing positive integer in linear time and constant space. 
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1,-2] should give 2. The input [1, 2, 0] should give 3.
 */

function findMissing(array) {
  // swap values to their index positions
  for (i = 0; i < array.length; i++) {
    swap(array, i);
    // If we swapped in a value that is a positive number and,
    // it is not already in place, we need to process it, so don't move index forward yet
    if (array[i] > 0 && array[i] !== i + 1) i--;
  }
  // array should now have the numbers in order
  // as soon we find a value which is out of place, we are done
  for (i = 0; i < array.length; i++) {
    if (array[i] !== i + 1) {
      return i + 1;
    }
  }
  // the array had all expected values, so return the next possible value
  return array.length;
}

// swap the value at index i to where (index == value -1) so 1 will be at index 0
function swap(array, i) {
  if (i < 0) return;
  target = array[i] - 1;
  // we are not interested in negative numbers so keep going
  if (target < 0) return;
  // if the value would be more than the array length it cannot be a value we
  // are interested in since the value gotta be between 1 and array length,
  // so just gonna ignore it
  if (target >= array.length) {
    array[i] = -1; // dummy value, since there is nothing to swap from target index
    return;
  }
  // Do the swap
  next = array[target];
  array[target] = array[i];
  array[i] = next;
}

function debug(array) {
  orig = [...array];
  val = findMissing(array);
  console.log(`missing value for ${orig} is ${val}`);
}

debug([3, 4, -1, 1, -2]);
debug([1, 2, 0]);
debug([2, 3, -7, 6, 8, 1, -10, 15]);
