export function makeSlug(prefix: string) {
  const randomPart =
    typeof crypto !== "undefined" && crypto.getRandomValues
      ? (() => {
          const bytes = new Uint32Array(1);
          crypto.getRandomValues(bytes);
          return bytes[0].toString(36).slice(0, 6);
        })()
      : Math.random().toString(36).slice(2, 8);

  return `${prefix}-${Date.now().toString(36)}-${randomPart}`;
}
