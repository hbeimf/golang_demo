package controller

// https://www.ctolib.com/gods.html



// hub maintains the set of active clients and broadcasts messages to the
// clients.
// hub结构作为一个聊天室结构，用于给在clients　哈希列表内的客户端广播信息。
type Hub struct {
    // hub　名称, 用于区分不同的 hub
    hub_name string

    // Registered clients.
    clients map[*Client]bool

    // Inbound messages from the clients.
    broadcast chan []byte

    // Register requests from the clients.
    register chan *Client

    // Unregister requests from clients.
    unregister chan *Client
}

// 每个场景对应一个　k/v ，　k指hub的hub_name，　v对应hub的指针.
// 客户端切换场景啥的，后面再慢慢实现，所有的聊天室哈希列表
var hubList map[string]*Hub

func init() {
    hubList = make(map[string]*Hub)

    // 初始化第一个默认的聊天室
    var hub_name = "default_chat_hub"
    newHub(hub_name)
}


func newHub(name string) {
    // 判断hub是否已经存在，如不存在则创建新hub

    if _, ok := hubList[name]; !ok {
        hub := &Hub{
            hub_name: name,
            broadcast:  make(chan []byte),
            register:   make(chan *Client),
            unregister: make(chan *Client),
            clients:    make(map[*Client]bool),
        }

        hubList[name] = hub

        // 开启消息收发, 这里得做成一个异步的协程，不然死循环会一直阻塞，
        go hub.run()
    }
}

// 每个聊天室用来异步接收和广播消息的进程　
func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}


