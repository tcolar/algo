function containsDuplicates(a) {
  all = {};
  for (e of a) {
    if (e in all) return true;
    all[e] = true;
  }
  return false;
}
/*
Given an array of integers, write a function that determines whether the array contains any duplicates. Your function should return true if any element appears at least twice in the array, and it should return false if every element is distinct.
*/
