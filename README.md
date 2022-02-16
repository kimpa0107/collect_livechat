# 采集直播平台聊天内容

## 启动服务器
```bash
go run main.go
```

## 客户端

> 在浏览器打开直播间，打开开发者工具，在`console`下执行下面的`JS`采集聊天内容

#### 抖音直播
```js
document
  .querySelector('.webcast-chatroom___items > div')
  .addEventListener('DOMNodeInserted', function(evt) {
    const node = evt.srcElement;
    const spans = node.querySelectorAll(':scope > div > span');
    const nick = spans[1].innerText.replace(/：\s*$/, '');
    let text = spans[2].innerHTML;

    const roomInfo = document.querySelector('#_douyin_live_scroll_container_ > div > div > div > div > div');
    const infos = roomInfo.querySelectorAll(':scope > div');
    if (!infos || infos.length != 2) {
      return;
    }
    const title = infos[0].querySelector('h1').innerText;
    const room = infos[1].querySelectorAll(':scope > div')[0].innerText;

    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://localhost:8888/chat/douyin', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({ room, title, nick, text }));
}, false);
```

#### 斗鱼直播
```js
const chat = document.querySelector('#js-barrage-list');

chat.addEventListener('DOMNodeInserted', function(evt) {
  const li = evt.srcElement;
  const nick = li.querySelector('.Barrage-nickName').title;
  let textNode = li.querySelector('.Barrage-content');
  if (!textNode) {
  	textNode = li.querySelector('.Barrage-text');
  }
  const text = textNode.innerText;
  
  const room = '';
  const title = '';
  
  var xhr = new XMLHttpRequest();
  xhr.open('POST', 'http://localhost:8888/chat/douyu', true);
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.send(JSON.stringify({ room, title, nick, text }));
}, false);
```

#### 快手直播
```js
document
  .querySelector('.chat-container .chat-history')
  .addEventListener('DOMNodeInserted', function (evt) {
    const node = evt.srcElement;
    const nick = node.querySelector('.username') && node.querySelector('.username').innerText.replace(/：\s*$/, '');
    let text = node.querySelector('.comment').innerText.replace(/\s+$/, '');
    if (node.querySelector('.chat-comment-cell .like')) {
      text = `${text} 红心`;
    } else {
      const giftNode = node.querySelector('.chat-comment-cell.gift-comment');
      if (giftNode) {
        const img = giftNode.querySelector('.gift-img');
        if (img) {
          text = `${text} ${img.outerHTML}`;
        }
      }
    }
  
  	const room = document.querySelector('.live-user .live-user-name').innerText.trim();
  	const title = '';
  
    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://localhost:8888/chat/kuaishou', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({ room, title, nick, text }));
  }, false);
```

#### 咪咕视频
```js
document
  .querySelector('.chat-room .list-wrapper')
  .addEventListener('DOMNodeInserted', function (evt) {
		const item = evt.srcElement;
  	const nick = item.querySelector('.name').innerText;
  	const text = item.querySelector('.text').innerText;
  	const title = document.querySelector('.competition-info .title').innerText.trim();
  	const room = '';
  
  	var xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://localhost:8888/chat/migu', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({ room, title, nick, text }));
	}, false);
```

#### afreecatv
```js
const chat = document.querySelector('#chat_area');

chat.addEventListener('DOMNodeInserted', function(event) {
  const dl = event.srcElement;
  const dt = dl.querySelector('dt');
  const dd = dl.querySelector('dd');
  const nick = dt.innerText.replace(/\s?\s*$/, '');
  const text = dd.innerHTML;

  const bc = document.querySelector('.broadcast_information');

  const room = bc.querySelector('.text_information .nickname').title;
  const title = bc.querySelector('.text_information .broadcast_title').innerText;

  var xhr = new XMLHttpRequest();
  xhr.open('POST', 'http://localhost:8888/chat/afreecatv', true);
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.send(JSON.stringify({ room, title, nick, text }));
}, false)
```
