function isIPv4Address(s: string): boolean {
  const parts = s.split('.');
  if (parts.length !== 4) return false;
  for (const part of parts) {
    const n = Number(part);
    if (part.length === 0 || isNaN(n)) return false;
    if (n < 0 || n > 255) return false;
    if (n.toString() !== part) return false;
  }
  return true;
}
