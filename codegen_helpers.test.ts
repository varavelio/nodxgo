import { assertEquals } from "jsr:@std/assert@1.0.11";
import {
  createFuncName,
  decapitalize,
  hasConflict,
} from "./codegen_helpers.ts";
import type { Attr, El } from "./codegen_helpers.ts";

Deno.test("test hasConflict", async (t) => {
  await t.step("no conflict", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "c", description: "" },
      { name: "d", description: "" },
    ];

    assertEquals(hasConflict("z", els, attrs, []), false);
  });

  await t.step("one conflict", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "c", description: "" },
      { name: "d", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs, []), false);
  });

  await t.step("two conflicts", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "a", description: "" },
      { name: "c", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs, []), true);
  });

  await t.step("three conflicts", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "a", description: "" },
      { name: "b", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs, []), true);
  });

  await t.step("keyword conflict", () => {
    assertEquals(hasConflict("map", [], [], ["map"]), true);
    assertEquals(hasConflict("Map", [], [], ["map"]), true);
    assertEquals(hasConflict("select", [], [], ["select"]), true);
    assertEquals(hasConflict("type", [], [], ["type"]), true);
    assertEquals(hasConflict("var", [], [], ["var"]), true);
    assertEquals(hasConflict("default", [], [], ["default"]), true);
    assertEquals(hasConflict("div", [], [], ["map"]), false);
    assertEquals(hasConflict("map", [], [], []), false);
  });
});

Deno.test("test createFuncName", async (t) => {
  await t.step("normal element name", () => {
    assertEquals(createFuncName("div", "El", false), {
      name: "Div",
      isGlob: false,
    });

    assertEquals(createFuncName("DiV", "El", false), {
      name: "Div",
      isGlob: false,
    });

    assertEquals(createFuncName("dIv", "El", false), {
      name: "Div",
      isGlob: false,
    });
  });

  await t.step("normal attribute name", () => {
    assertEquals(createFuncName("class", "Attr", false), {
      name: "Class",
      isGlob: false,
    });

    assertEquals(createFuncName("ClAsS", "Attr", false), {
      name: "Class",
      isGlob: false,
    });

    assertEquals(createFuncName("cLaSs", "Attr", false), {
      name: "Class",
      isGlob: false,
    });
  });

  await t.step("element with conflict", () => {
    assertEquals(createFuncName("div", "El", true), {
      name: "DivEl",
      isGlob: false,
    });

    assertEquals(createFuncName("DiV", "El", true), {
      name: "DivEl",
      isGlob: false,
    });

    assertEquals(createFuncName("dIv", "El", true), {
      name: "DivEl",
      isGlob: false,
    });
  });

  await t.step("attribute with conflict", () => {
    assertEquals(createFuncName("class", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });

    assertEquals(createFuncName("ClAsS", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });

    assertEquals(createFuncName("cLaSs", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });
  });

  await t.step("glob attribute", () => {
    assertEquals(createFuncName("data-*", "Attr", false), {
      name: "Data",
      isGlob: true,
    });

    assertEquals(createFuncName("data-*", "Attr", true), {
      name: "DataAttr",
      isGlob: true,
    });
  });
});

Deno.test("test decapitalize", () => {
  assertEquals(decapitalize("Hello"), "hello");
  assertEquals(decapitalize("HELLO"), "hELLO");
  assertEquals(decapitalize("HELLO WORLD"), "hELLO WORLD");
  assertEquals(decapitalize("HELLO-WORLD"), "hELLO-WORLD");
});
