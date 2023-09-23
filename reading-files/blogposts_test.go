package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/wuziqiu666/learngo/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
world!`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
E
L
M`
	)
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello-world.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFs(fs)
	assertNoError(t, err)
	assertPostsLength(t, posts, fs)
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
world!`,
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
