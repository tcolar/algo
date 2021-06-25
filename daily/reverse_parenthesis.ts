function reverseInParentheses(inputString: string): string {
  let starts: number[] = [];
  // make an array so we can edit in place
  let s = inputString.split('');
  for (let i = 0; i < s.length; i++) {
    const c = s[i];
    if (c === '(') {
      // keep track of starting brackets (revert start)
      starts.push(i);
    } else if (c === ')') {
      // when finding a closing bracket pop the opening location
      const from = starts.pop();
      // revert from ( to ) in place
      for (let j = 0; j < (i - from) / 2; j++) {
        const tmp = s[from + j];
        s[from + j] = s[i - j];
        s[i - j] = tmp;
      }
    }
  }
  // Remove the brackets at the end, makes index tracking simpler
  return s.join('').replace(/[()]/g, '');
}
