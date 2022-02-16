package service

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestGetwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("cannot get working directory: %v", err)
		return
	}
	fmt.Printf("working directory: %s\n", dir)
}

func TestParseDouyinEmoji(t *testing.T) {
	msg := `abc<div class="webcast-chatroom___content-with-emoji-emoji"><img draggable="false" src="https://p3-pc-sign.douyinpic.com/obj/tos-cn-i-tsj2vxp0zn/b869ffdb59324b5bb4956274c608e4eb?x-expires=1960034400&amp;x-signature=%2F%2Fl38ktJXtgxfcm1vykKYwEr77g%3D&amp;from=876277922" alt="[鼓掌]"></div><div class="webcast-chatroom___content-with-emoji-emoji"><img draggable="false" src="https://p3-pc-sign.douyinpic.com/obj/tos-cn-i-tsj2vxp0zn/b869ffdb59324b5bb4956274c608e4eb?x-expires=1960034400&amp;x-signature=%2F%2Fl38ktJXtgxfcm1vykKYwEr77g%3D&amp;from=876277922" alt="[鼓掌]"></div><div class="webcast-chatroom___content-with-emoji-emoji"><img draggable="false" src="https://p3-pc-sign.douyinpic.com/obj/tos-cn-i-tsj2vxp0zn/b869ffdb59324b5bb4956274c608e4eb?x-expires=1960034400&amp;x-signature=%2F%2Fl38ktJXtgxfcm1vykKYwEr77g%3D&amp;from=876277922" alt="[鼓掌]"></div>def`
	ptn := regexp.MustCompile(`<div.+?webcast-chatroom___content-with-emoji-emoji".+?<img.+?alt="(.+?)".*?></div>`)
	msg = ptn.ReplaceAllString(msg, "$1")
	fmt.Printf("%s\n", msg)
}

func TestParseAfreecatvEmoji(t *testing.T) {
	msg := `<img class="emoticon" src="https://res.afreecatv.com/images/chat/emoticon/small/24.png"> <div style="display:none;">/더럽/</div>`
	ptn := regexp.MustCompile(`<img class="emoticon".+?<div style="display:none;">(.+?)</div>`)
	msg = ptn.ReplaceAllString(msg, "$1")
	fmt.Printf("%s\n", msg)
}

func TestParseAfreecatvEmojiImg(t *testing.T) {
	msg := `
	<div class="ogq_img " data-id="17c5a057c54378f" data-subid="21">
	<img style="cursor:pointer" src="//ogq-sticker-global-cdn-z01.afreecatv.com/sticker/17c5a057c54378f/21_80.png?ver=1" onerror="this.src='//res.afreecatv.com/images/chat/ogq_default.png'">
</div>
`
	msg = strings.Trim(msg, " ")
	msg = regexp.MustCompile("^\n").ReplaceAllString(msg, "")
	msg = strings.Trim(msg, "\t")
	msg = regexp.MustCompile("\n$").ReplaceAllString(msg, "")
	ptn := regexp.MustCompile(`<div class="ogq_img[\s\S]+<img style="cursor:pointer" src="(.+?)".+?>[\s\S]*</div>`)
	msg = ptn.ReplaceAllString(msg, "$1 ")

	fmt.Printf("%s\n", msg)
}
