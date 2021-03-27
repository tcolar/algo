/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

function productArray(array) {
  var totalProduct = 1;
  for (i of array) {
    totalProduct *= i;
  }
  console.log(totalProduct);
  results = [];
  for (i of array) {
    results.push(totalProduct / i);
  }
  return results;
}

function productArray2(array) {
  console.log(array);
  results = [array.length];
  // product of stuff before index
  fromStart = [array.length];
  // product of stuff after index
  fromEnd = [array.length];
  total1 = 1;
  total2 = 1;
  for (i = 0; i < array.length; i++) {
    total1 *= array[i];
    total2 *= array[array.length - 1 - i];
    fromStart[i] = total1;
    fromEnd[array.length - 1 - i] = total2;
  }
  for (i = 0; i < array.length; i++) {
    // multiply the product of stuff before index, with the product of stuff after
    // should leave us with the total product without value at i
    results[i] = (fromStart[i - 1] || 1) * (fromEnd[i + 1] || 1);
  }

  return results;
}

console.log(productArray2([1, 2, 3, 4, 5]));
console.log(productArray2([3, 2, 1]));
