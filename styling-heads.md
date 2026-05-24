Just use a writing.css file for consistent styling. ie.

```css
@theme {
  /* Tailwind v4 theme configurations go here */
}

@layer base {
  h2 {
    @apply text-3xl font-bold text-gray-900 tracking-tight mb-4;
  }
  
  p {
    @apply text-base text-gray-600 leading-relaxed;
  }
}
```

It saves doing this below...
Style headings consistency:
```go
package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// 1. Define the transformer
type headingTransformer struct{}

func (g *headingTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && n.Kind() == ast.KindHeading {
			heading := n.(*ast.Heading)
			
			// Style based on heading level (h1, h2, h3...)
			switch heading.Level {
			case 1:
				heading.SetAttributeString("class", []byte("text-4xl font-bold my-4 text-gray-900"))
			case 2:
				heading.SetAttributeString("class", []byte("text-2xl font-semibold my-3 text-gray-800"))
			default:
				heading.SetAttributeString("class", []byte("text-xl font-medium my-2"))
			}
		}
		return ast.WalkContinue, nil
	})
}

func main() {
	// 2. Register the transformer with Goldmark
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&headingTransformer{}, 100),
			),
		),
	)

	src := []byte("# Hello World\n## Subheading")
	var buf bytes.Buffer
	if err := md.Convert(src, &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
```
