import { toPascalCase } from "jsr:@std/text@1.0.10";

export interface El {
  name: string;
  isVoid: boolean;
  description: string;
}

export interface Attr {
  name: string;
  description: string;
  isBoolean?: boolean;
  isList?: boolean;
}

export function hasConflict(
  name: string,
  els: El[],
  attrs: Attr[],
  keywords: string[],
): boolean {
  // Keyword collisions: if the name collides with a reserved keyword of the
  // target language (see data/keywords.json), it must be escaped.
  for (const keyword of keywords) {
    if (keyword.toLowerCase() === name.toLowerCase()) return true;
  }

  // Element/attribute collisions: if more than one element or attribute share
  // the same name, they must be escaped (e.g. the title element and the title
  // attribute).
  let matches = 0;

  for (const el of els) {
    if (el.name.toLowerCase() === name.toLowerCase()) matches++;
  }
  for (const attr of attrs) {
    if (attr.name.toLowerCase() === name.toLowerCase()) matches++;
  }

  return matches > 1;
}

export interface FuncName {
  name: string;
  isGlob: boolean;
}

export function createFuncName(
  name: string,
  type: "El" | "Attr",
  hasConflict: boolean,
): FuncName {
  const isGlob = name.endsWith("*");
  if (isGlob) {
    name = name.slice(0, -1);
  }

  const word = name.split("-").join(" ");
  name = toPascalCase(word.toLowerCase());

  if (hasConflict) {
    name = `${name}${type}`;
  }

  return { name, isGlob };
}

export function decapitalize(str: string): string {
  return str.charAt(0).toLowerCase() + str.slice(1);
}
