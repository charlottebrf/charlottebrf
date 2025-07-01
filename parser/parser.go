package parser

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MarkdownToHTML(markdownText []byte) []byte {
	// Create a new parser with extensions (optional)
	extensions := parser.CommonExtensions
	parser := parser.NewWithExtensions(extensions)

	// Parse the Markdown into an AST
	doc := parser.Parse(markdownText)

	// Create a new HTML renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// Render the AST to HTML
	htmlOutput := markdown.Render(doc, renderer)
	return htmlOutput
}
