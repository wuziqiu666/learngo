package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/wuziqiu666/learngo/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte("hi")},
		"hello-world.md": {Data: []byte("hola")},
	}
	posts := blogposts.NewPostsFromFs(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
