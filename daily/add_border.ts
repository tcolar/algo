function addBorder(picture: string[]): string[] {
  const ln = picture[0].length;
  const result: string[] = [];
  result.push('*'.repeat(ln + 2));
  for (const s of picture) {
    result.push(`*${s}*`);
  }
  result.push('*'.repeat(ln + 2));
  return result;
}
