package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"livechat/config"
	"livechat/global"
	"livechat/initial"
	"livechat/middleware"
	"livechat/service"
)

var (
	port  int
	debug bool
)

func init() {
	flag.IntVar(&port, "port", 8888, "listen port")
	flag.BoolVar(&debug, "debug", false, "debug mode")

	initial.DB(config.MySQL{
		Host: "localhost",
		Port: "3306",
		User: "root",
		Pass: "root",
		Name: "livechat",
	})
}

func main() {
	flag.Parse()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.Cors())

	chat := r.Group("/chat")
	{
		srv := service.NewChatService(global.DB)

		platforms := []string{"douyin", "kuaishou", "douyu", "migu", "afreecatv", "pandatv", "flextv"}
		for _, platform := range platforms {
			chat.POST(fmt.Sprintf("/%s", platform), srv.Write(platform))
		}
	}

	r.Run(fmt.Sprintf(":%d", port))
}
