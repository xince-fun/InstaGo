package main

import (
	post "github.com/xince-fun/InstaGo/server/shared/kitex_gen/post/postservice"
	"log"
)

func main() {
	svr := post.NewServer(new(PostServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
