package blogposts

import (
	"io/fs"
	"log"
)

type post struct {
}

func NewPostsFromFs(fileSystem fs.FS) []post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []post
	for range dir {
		log.Print(dir)
		posts = append(posts, post{})
	}
	return posts
}
