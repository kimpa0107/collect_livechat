package main

import (
	"fmt"
	"livechat/config"
	"livechat/global"
	"livechat/initial"
	"livechat/middleware"
	"livechat/service"

	"github.com/gin-gonic/gin"
)

func main() {
	initial.DB(config.MySQL{
		Host: "localhost",
		Port: "3306",
		User: "root",
		Pass: "root",
		Name: "livechat",
	})

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

	r.Run(":8888")
}
