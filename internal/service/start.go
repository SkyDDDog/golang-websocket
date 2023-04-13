package service

import (
	"demo04/pkg/errno"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func (manager *ClientManager) Start() {
	for {
		log.Println("-----监听管道通信-----")
		select {
		case conn := <-Manager.Register:
			log.Println("有新连接:", conn.ID)
			Manager.Clients[conn.ID] = conn // 把连接放到用户管理
			replyMsg := ReplyMsg{
				Code:    errno.WebsocketSuccess,
				Content: "已连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			log.Println("连接失败,", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    errno.WebsocketEnd,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false // 默认对方不在线
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				replyMsg := &ReplyMsg{
					Code:    errno.WebsocketOnlineReply,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := InsertMsg(id, string(message), 1, int64(3*month))
				if err != nil {
					log.Println("InsertOne Err: {}", err)
				}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    errno.WebsocketOfflineReply,
					Content: "对方不在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(id, string(message), 0, int64(3*month))
				if err != nil {
					log.Println("InsertOneMsg Err: {}", err)
				}
			}
		}
	}
}
