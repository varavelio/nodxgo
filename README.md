# NodX for Go

<p align="center">
  <a href="https://github.com/varavelio/nodxgo/actions">
    <img src="https://github.com/varavelio/nodxgo/actions/workflows/ci.yaml/badge.svg" alt="CI status"/>
  </a>
  <a href="https://pkg.go.dev/github.com/varavelio/nodxgo">
    <img src="https://pkg.go.dev/badge/github.com/varavelio/nodxgo" alt="Go Reference"/>
  </a>
  <a href="https://github.com/varavelio/nodxgo/releases/latest">
    <img src="https://img.shields.io/github/release/varavelio/nodxgo.svg" alt="Release Version"/>
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/varavelio/nodxgo.svg" alt="License"/>
  </a>
  <a href="https://github.com/varavelio/nodxgo">
    <img src="https://img.shields.io/github/stars/varavelio/nodxgo?style=flat&label=github+stars"/>
  </a>
</p>

<p align="center">
  <a href="https://varavel.com">
    <img src="https://cdn.jsdelivr.net/gh/varavelio/brand@1.0.0/dist/badges/project.svg" alt="A Varavel project"/>
  </a>
</p>

---

NodX is a modern and developer-friendly Go template engine for generating
**safe**, **clean**, and **maintainable** HTML. Designed for maximum
productivity and easy maintenance, it combines **simplicity**, **type safety**
and **robust formatting**, making it the perfect fit for your Go-based web
projects.

## Key Features

- **Type Safety 🛡️**: Fully typed APIs ensure you write error-free HTML, even at
  scale.
- **Robust Formatting 🧹**: Works seamlessly with `go fmt` to keep your code
  clean and consistent.
- **Zero Dependencies 📦**: Lightweight and fast, with no external dependencies.
- **DX in mind 🧠**: If you can write HTML and Go, you can write NodX.
- **Security by Default 🔒**: Automatically escapes unsafe text to prevent XSS
  vulnerabilities.

## Quick Start

Install the library:

```bash
# Tested on Go 1.22 and later
go get github.com/varavelio/nodxgo
```

Start building your HTML with intuitive, type-safe functions:

```go
package main

import (
  "os"
  nodx "github.com/varavelio/nodxgo"
)

func main() {
  // you can fetch this from a database or some other source
  happiness := 100
  hideContainer := false

  myTemplate := nodx.Group(
    nodx.DocType(),
    nodx.Html(
      nodx.Head(
        nodx.TitleEl(nodx.Text("My NodX Page")),
      ),
      nodx.Body(
        nodx.Div(
          nodx.ClassMap{
            "container": true,
            "hidden":    hideContainer,
          },
          nodx.H1(
            nodx.Class("title"),
            nodx.Textf("Welcome to %s!", "NodX"),
          ),
          nodx.P(nodx.Text("This is a type-safe HTML generator for Go.")),
          nodx.If(happiness > 90, nodx.P(nodx.Textf("With NodX, you will be %d%% happy!", happiness))),
        ),
      ),
    ),
  )

  _ = myTemplate.Render(os.Stdout)
  // or
  // str, err := myTemplate.RenderString()
  // byt, err := myTemplate.RenderBytes()
}
```

### Output:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My NodX Page</title>
  </head>
  <body>
    <div class="container">
      <h1 class="title">Welcome to NodX!</h1>
      <p>This is a type-safe HTML generator for Go.</p>
      <p>With NodX, you will be 100% happy!</p>
    </div>
  </body>
</html>
```

## Key Concepts

### 1. **Elements made easy**

Every HTML tag is a function! Just call the function with child elements,
attributes, groups or text.

```go
nodx.Div(
  nodx.Class("container"),
  nodx.H1(nodx.Text("Hello, NodX!")),
  nodx.P(nodx.Text("Build clean and safe HTML effortlessly.")),
)
```

### 2. **Attributes with helpers**

Attributes like `class`, `src`, and `alt` have their own functions for
simplicity.

```go
nodx.Img(
  nodx.Src("image.jpg"),
  nodx.Alt("A beautiful image"),
)
```

**Boolean attributes** (like `checked`, `disabled` and `required`) take a `bool`
value: when `true` they render as just the attribute name, and when `false` they
are omitted entirely.

```go
isChecked := true
isDisabled := false

nodx.Input(
  nodx.Checked(isChecked),
  nodx.Disabled(isDisabled),
)
// Output: <input checked>
```

### 3. **Dynamic class management**

Use `ClassMap` to conditionally apply classes based on your logic.

```go
nodx.Div(
  nodx.ClassMap{
    "visible": true,
    "hidden":  false,
  },
  nodx.Text("Conditional classes made simple!"),
)
```

### 4. **Fully typed server rendered components**

You can create your own components to avoid code duplication and keep your code
clean.

```go
func button(text string) nodx.Node {
  return nodx.Button(
    nodx.Class("btn"),
    nodx.Text(text),
  )
}

func main() {
  myTemplate := nodx.Div(
    button("Click me 1!"),
    button("Click me 2!"),
    button("Click me 3!"),
  )

  _ = myTemplate.Render(os.Stdout)

  /*
    Output:
    <div>
      <button class="btn">Click me 1!</button>
      <button class="btn">Click me 2!</button>
      <button class="btn">Click me 3!</button>
    </div>
  */
}
```

### 5. **Advanced features**

Please refer to the
[Full Documentation](https://pkg.go.dev/github.com/varavelio/nodxgo) to read
more about all the included features.

- **Custom components**: You can create your own components to avoid code
  duplication and keep your code clean.
- **Conditional rendering**: You can use the `nodx.If` and `nodx.IfFunc`
  function to conditionally render elements based on your logic.
- **Map**: You can use the `nodx.Map` function to loop over a list of items and
  render them as a list of Nodes.
- **Custom elements and attributes**: You can use `nodx.El` and `nodx.Attr` to
  create your own elements and attributes if they are not included in the
  library.
- **Component library**: You can write your own component library and publish it
  for others to use.
- **Dynamic classes and styles**: You can use the `nodx.ClassMap` and
  `nodx.StyleMap` to conditionally apply classes and styles based on your own
  logic.

### 6. **Naming collisions**

Following the
[NodX specification](https://github.com/varavelio/nodx/blob/main/SPEC.md), the
generated functions never collide with Go keywords or with each other. When a
name would collide, it gets a deterministic suffix: `El` for elements and `Attr`
for attributes.

- **Go keyword collisions** (see
  [`data/keywords.json`](https://github.com/varavelio/nodx/blob/main/data/keywords.json)):
  `defer` → `DeferAttr`, `default` → `DefaultAttr`, `for` → `ForAttr`, `map` →
  `MapEl`, `select` → `SelectEl`, `type` → `TypeAttr`, `var` → `VarEl`.
- **Element/attribute collisions**: `data` → `DataEl`/`DataAttr`, `dir` →
  `DirEl`/`DirAttr`, `form` → `FormEl`/`FormAttr`, `title` →
  `TitleEl`/`TitleAttr`.

## Why Choose NodX?

- **Expressive and Clear Code**: Great helpers (`Div`, `H1`, `P`, etc.) that
  mirror HTML semantics, making your Go code as readable as HTML.
- **Battle-Tested for Safety**: Text is escaped automatically to protect your
  app from XSS vulnerabilities.
- **Lightweight and Efficient**: With no dependencies, you can focus on building
  without bloat.
- **Perfect for Modern Go Developers**: Designed with Go's simplicity and
  elegance in mind.

## Ecosystem

NodX can be extended with additional libraries that provide ready-to-use
components, attributes, and server utilities. Here are some projects built on
top of NodX:

- [**nodxgo-lucide**](https://github.com/varavelio/nodxgo-lucide) — Beautiful &
  consistent icons for NodX Go.
- [**nodxgo-alpine**](https://github.com/varavelio/nodxgo-alpine) — Alpine.js
  attributes for NodX Go.
- [**nodxgo-htmx**](https://github.com/varavelio/nodxgo-htmx) — HTMX attributes
  and server utilities for NodX Go.

> **Building your own NodX library?** We'd love to feature it here! Open a pull
> request adding your project to this list, and help grow the NodX ecosystem
> together.

## License

NodX is open-source and available under the [MIT License](LICENSE). Feel free to
use it in your personal or commercial projects.
