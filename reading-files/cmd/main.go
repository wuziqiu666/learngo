package main

import (
	"log"
	"os"

	blogposts "github.com/wuziqiu666/learngo/reading-files"
)

func main() {
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Printf("%+v", post)
	}
}
