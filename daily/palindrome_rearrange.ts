function palindrome(s: string): boolean {
  const map = {};
  for (const c of s) {
    map[c] = (map[c] || 0) + 1;
  }
  let oddCount = 0;
  console.log(map);
  for (const key in map) {
    const count = map[key];
    if (count % 2 === 1) {
      oddCount++;
    }
  }
  return oddCount <= 1;
}

console.log(palindrome('racecar'));
console.log(palindrome('raczecar'));
console.log(palindrome('baba'));
console.log(palindrome('ae'));
console.log(palindrome('a'));
console.log(palindrome(''));
console.log(palindrome('abebaa'));
/*
"abcntnnnn"
dog -> dog
baba -> abba baab+
ababa

racecar

a 2 3
r 2 
c 2
e 1
*/
