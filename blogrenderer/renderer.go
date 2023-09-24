package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

const ()

type PostRenderer struct {
	templ      *template.Template
	extensions parser.Extensions
	htmlFlags  html.Flags
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	return &PostRenderer{
		templ:      templ,
		extensions: extensions,
		htmlFlags:  htmlFlags,
	}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(post, r)); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}
	if err := templ.Execute(w, posts); err != nil {
		return err
	}
	return nil
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostViewModel(post Post, r *PostRenderer) *postViewModel {
	var vm postViewModel
	vm.Post = post
	vm.HTMLBody = template.HTML(ConfigMarkDownRender(r.extensions, r.htmlFlags)([]byte(post.Body)))
	return &vm
}

func ConfigMarkDownRender(extensions parser.Extensions, htmlFlag html.Flags) func([]byte) []byte {
	return func(md []byte) []byte {
		p := parser.NewWithExtensions(extensions)
		doc := p.Parse(md)
		opts := html.RendererOptions{Flags: htmlFlag}
		render := html.NewRenderer(opts)
		return markdown.Render(doc, render)
	}

}
