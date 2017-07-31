package controller

// https://www.ctolib.com/gods.html


// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.



// hub maintains the set of active clients and broadcasts messages to the
// clients.
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


// var hub *Hub

// 每个场景对应一个　k/v ，　k指hub的hub_name，　v对应hub的指针.
// 客户端切换场景啥的，后面再慢慢实现
var hubList map[string]*Hub

func init() {
    hubList = make(map[string]*Hub)

    var hub_name = "default_hub"
    newHub(hub_name)
    // hubList[hub_name] = hub
    // hub.run()
}


func newHub(name string) {
    hub := &Hub{
        hub_name: name,
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
        clients:    make(map[*Client]bool),
    }

    hubList[name] = hub

    // 开启消息收发
    hub.run()
}

// func newHub(name string) *Hub {
//     return &Hub{
//         hub_name: name,
//         broadcast:  make(chan []byte),
//         register:   make(chan *Client),
//         unregister: make(chan *Client),
//         clients:    make(map[*Client]bool),
//     }
// }

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


