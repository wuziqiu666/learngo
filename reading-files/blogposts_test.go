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
	posts, err := blogposts.NewPostsFromFs(fs)
	assertNoError(t, err)
	assertPostsLength(t, posts, fs)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS){
	t.Helper()
	if len(posts) != len(fs){
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}