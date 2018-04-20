export function equalArrays(a1: any[], a2: any[]) {
  if (!Array.isArray(a1) || !Array.isArray(a2) || a1.length !== a2.length) {
    return false;
  }
  return JSON.stringify(a1) === JSON.stringify(a2);
}

export function generateRandomString(length: number) {
  let d = new Date().getTime();
  function randomChar() {
    const r = (d + Math.random() * 16) % 16 | 0;
    d = Math.floor(d / 16);
    return r.toString(16);
  }
  let str = '';
  for (let i = 0; i < length; ++i) str += randomChar();
  return str;
};
