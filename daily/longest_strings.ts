function allLongestStrings(inputArray: string[]): string[] {
  let longest = 0;
  let strings = [];
  for (const str of inputArray) {
    if (str.length > longest) {
      strings = [str];
      longest = str.length;
    } else if (str.length === longest) {
      strings.push(str);
    }
  }
  return strings;
}
