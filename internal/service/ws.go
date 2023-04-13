package service

import (
	"demo04/internal/cache"
	"demo04/pkg/errno"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

const month = 60 * 60 * 24 * 30 // 一个月30天

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	RoomId string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreateId(uid, toUid string) string {
	return fmt.Sprintf("%s->%s", uid, toUid)
}

func PrivateChatHandler(ginCtx *gin.Context) {
	uid := ginCtx.Query("uid")
	toUid := ginCtx.Query("toUid")
	// 升级为ws协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if err != nil {
		http.NotFound(ginCtx.Writer, ginCtx.Request)
	}
	// 创建用户实例
	client := &Client{
		ID:     CreateId(uid, toUid),
		SendID: CreateId(toUid, uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 将用户注册到manager
	Manager.Register <- client
	go client.Read()
	go client.Write()

}

func RoomChatHandler(ginCtx *gin.Context) {
	uid := ginCtx.Query("uid")
	roomId := ginCtx.Query("roomId")
	// 升级为ws协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if err != nil {
		http.NotFound(ginCtx.Writer, ginCtx.Request)
	}
	// 创建用户实例
	client := &Client{
		ID:     uid,
		RoomId: roomId,
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	go client.Write()

}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			log.Println("数据格式不正确", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}

		// 群聊
		if c.SendID == "" || c.RoomId != "" {
			replyMsg := ReplyMsg{
				Code:    errno.WebsocketSuccess,
				Content: sendMsg.Content,
			}
			msg, _ := json.Marshal(replyMsg) // 序列化
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			continue
		}

		// 单聊
		if sendMsg.Type == 1 { // 发送消息
			r1, _ := cache.RedisClient.Get(c.ID).Result()     // 1->2
			r2, _ := cache.RedisClient.Get(c.SendID).Result() // 2->1
			if r1 > "3" && r2 == "" {                         // 1给2发了三条消息,2无回应 -> 停止1发送
				replyMsg := ReplyMsg{
					Code:    errno.WebsocketLimit,
					Content: "达到限制",
				}
				msg, _ := json.Marshal(replyMsg) // 序列化
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				cache.RedisClient.Incr(c.ID)
				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
				// 防止过快"分手", 建立链接三个月过期
			}
			// 广播
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content), // 发送过来的信息
			}
		} else if sendMsg.Type == 2 {
			// 获取历史消息
			timeT, err := strconv.Atoi(sendMsg.Content) // string to int
			if err != nil {
				timeT = 999999
			}
			results, _ := FindMany(c.SendID, c.ID, int64(timeT), 10)
			if len(results) > 10 {
				results = results[:10]
			} else if len(results) == 0 {
				replyMsg := ReplyMsg{
					Code:    errno.WebsocketEnd,
					Content: "到底了",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}
			for _, result := range results {
				replyMsg := ReplyMsg{
					From:    result.From,
					Content: result.Msg,
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}

}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				Code:    errno.WebsocketSuccess,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}
