/*
Given an array strings, determine whether it follows the sequence given in the patterns array. In other words, there should be no i and j for which strings[i] = strings[j] and patterns[i] ≠ patterns[j] or for which strings[i] ≠ strings[j] and patterns[i] = patterns[j].

Example

For strings = ["cat", "dog", "dog"] and patterns = ["a", "b", "b"], the output should be
areFollowingPatterns(strings, patterns) = true;
For strings = ["cat", "dog", "doggy"] and patterns = ["a", "b", "b"], the output should be
areFollowingPatterns(strings, patterns) = false.
*/
function areFollowingPatterns(strings: string[], patterns: string[]): boolean {
  const map1 = new Map<string, string>();
  const map2 = new Map<string, string>();
  if (strings.length !== patterns.length) return false;

  for (let i = 0; i != strings.length; i++) {
    const str = strings[i];
    const pattern = patterns[i];
    if (map1.has(pattern)) {
      if (map1.get(pattern) !== str) {
        return false;
      }
    } else {
      map1.set(pattern, str);
    }
    if (map2.has(str)) {
      if (map2.get(str) !== pattern) {
        return false;
      }
    } else {
      map2.set(str, pattern);
    }
    //console.log(map1 , map2)
  }
  return true;
}
