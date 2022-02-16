package service

import (
	"fmt"
	"livechat/model"
	"livechat/utils/color"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type chatRequest struct {
	Room  string `json:"room" form:"room"`
	Title string `json:"title" form:"title"`
	Nick  string `json:"nick" form:"nick"`
	Text  string `json:"text" form:"text"`
}

type Chat struct {
	db *gorm.DB
}

func NewChatService(db *gorm.DB) *Chat {
	return &Chat{db: db}
}

func (s *Chat) Write(platform string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var chat chatRequest
		_ = ctx.ShouldBind(&chat)
		s.writeToFile(platform, chat)
		go s.writeToDB(platform, chat)
	}
}

func (s *Chat) writeToFile(platform string, chat chatRequest) {
	curDir, _ := os.Getwd()
	saveDir := curDir + "/storage/" + platform
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.MkdirAll(saveDir, os.ModePerm)
	}

	room := chat.Room
	if room == "" {
		room = platform
	}

	var f *os.File
	logFile := fmt.Sprintf("%s/%s_%s.log", saveDir, room, time.Now().Format("2006-01-02"))

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ = os.Create(logFile)
	} else {
		f, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	}
	defer f.Close()

	dt := time.Now().Format("15:04:05")

	fmt.Printf("%s %s %s %s %s %s\n",
		color.PrintWith("["+dt+"]", color.TEXT_BLUE),
		color.PrintWith(platform, color.TEXT_GREEN),
		color.PrintWith(room, color.TEXT_CYAN),
		color.PrintWith(chat.Nick, color.TEXT_YELLOW),
		color.PrintWith("->", color.TEXT_GREEN),
		pureText(platform, chat.Text))

	_, _ = f.WriteString(fmt.Sprintf("[%s] %s -> %s\n", dt, chat.Nick, pureText(platform, chat.Text)))
}

func pureText(platform, msg string) string {
	switch platform {
	case "douyin":
		msg = douyinPureText(msg)
	case "afreecatv":
		msg = afreecatvPureText(msg)
	}
	return msg
}

func douyinPureText(msg string) string {
	var gift strings.Builder
	isGift := false

	ptn := regexp.MustCompile(`>(送出了)<img.+?src="(.+?)".+?>.+?(×)&nbsp;(\d+)`)
	matches := ptn.FindAllStringSubmatch(msg, -1)
	if len(matches) > 0 {
		isGift = true
		for _, match := range matches[0][1:] {
			gift.WriteString(match + " ")
		}
	}
	if isGift {
		return gift.String()
	}

	ptn = regexp.MustCompile(`<div.+?webcast-chatroom___content-with-emoji-emoji.+?<img.+?alt="(.+?)".*?></div>`)
	msg = ptn.ReplaceAllString(msg, "$1")

	return removeHTML(msg)
}

func afreecatvPureText(msg string) string {
	msg = replaceSpace(msg)
	msg = strings.Trim(msg, " ")

	// replace emoji text
	ptn := regexp.MustCompile(`<img class="emoticon".+?<div style="display:none;">(.+?)</div>`)
	msg = ptn.ReplaceAllString(msg, "$1")

	// replace sticker image url
	msg = regexp.MustCompile("^\n").ReplaceAllString(msg, "")
	msg = strings.Trim(msg, "\t")
	msg = regexp.MustCompile("\n").ReplaceAllString(msg, "")
	ptn = regexp.MustCompile(`<div class="ogq_img[\s\S]+<img style="cursor:pointer" src="(.+?)".+?>[\s\S]*</div>`)
	msg = ptn.ReplaceAllString(msg, "http:$1 ")

	msg = strings.Replace(msg, `&lt;`, "<", -1)
	msg = strings.Replace(msg, `&gt;`, ">", -1)
	msg = strings.Replace(msg, `&amp;`, "&", -1)

	msg = regexp.MustCompile(`\s{2,}`).ReplaceAllString(msg, " ")

	return msg
}

func replaceSpace(msg string) string {
	return strings.Replace(msg, `&nbsp;`, " ", -1)
}

func removeHTML(msg string) string {
	ptn := regexp.MustCompile(`<[^>]*>`)
	return ptn.ReplaceAllString(msg, "")
}

func (s *Chat) writeToDB(platform string, req chatRequest) {
	now := time.Now()

	chat := model.LiveChat{
		Platform:  platform,
		Room:      req.Room,
		Title:     req.Title,
		Date:      now.Format("2006-01-02"),
		Nickname:  req.Nick,
		Message:   pureText(platform, req.Text),
		CreatedAt: now,
	}

	switch platform {
	case "douyin":
		liveChat := model.DouyinLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "kuaishou":
		liveChat := model.KuaishouLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "douyu":
		liveChat := model.DouyuLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "migu":
		liveChat := model.MiguLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "afreecatv":
		liveChat := model.AfreecatvLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "pandatv":
		liveChat := model.PandatvLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	case "flextv":
		liveChat := model.FlextvLiveChat{
			LiveChat: chat,
		}
		_ = s.db.Create(&liveChat)
	}
}
