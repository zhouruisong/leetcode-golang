package main

import (
	"github.com/gorilla/websocket"
	"github.com/gofrs/uuid"
	"time"
	"log"
	"sync"
	"encoding/json"
	"net/http"
)

type Client struct {
	id string

	conn *websocket.Conn

	sendLock sync.Mutex
	// sendCallback is callback based on packetID
	sendCallback     map[string]func(req WSPacket)
	sendCallbackLock sync.Mutex
	// recvCallback is callback when receive based on ID of the packet
	recvCallback map[string]func(req WSPacket)

	Done chan struct{}
}

type WSPacket struct {
	ID           string `json:"id"`
	Data         string `json:"data"`
	RoomID       string `json:"room_id"`
	PlayerIndex  int    `json:"player_index"`
	TargetHostID string `json:"target_id"`
	PacketID     string `json:"packet_id"`
	// Globally ID of a session
	SessionID string `json:"session_id"`
}

func (c *Client) Send(request WSPacket, callback func(response WSPacket)) {
	request.PacketID = uuid.Must(uuid.NewV4()).String()
	data, err := json.Marshal(request)
	if err != nil {
		return
	}

	// TODO: Consider using lock free
	// Wrap callback with sessionID and packetID
	if callback != nil {
		wrapperCallback := func(resp WSPacket) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("Recovered from err in client callback ", err)
				}
			}()

			resp.PacketID = request.PacketID
			resp.SessionID = request.SessionID
			callback(resp)
		}
		c.sendCallbackLock.Lock()
		c.sendCallback[request.PacketID] = wrapperCallback
		c.sendCallbackLock.Unlock()
	}

	c.sendLock.Lock()
	c.conn.SetWriteDeadline(time.Now().Add(time.Duration(1)))
	c.conn.WriteMessage(websocket.TextMessage, data)
	c.sendLock.Unlock()
}

// Heartbeat maintains connection to server
func (c *Client) Heartbeat() {
	// send heartbeat every 1s
	timer := time.Tick(time.Second)

	for range timer {
		select {
		case <-c.Done:
			log.Println("Close heartbeat")
			return
		default:
		}
		c.Send(WSPacket{ID: "heartbeat"}, nil)
	}
}

func (o *Client) WS(w http.ResponseWriter, r *http.Request) {
	log.Println("Browser connected to overlord")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Warn: Something wrong. Recovered in ", r)
		}
	}()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("[!] WS upgrade:", err)
		return
	}
	defer c.Close()

	client := NewBrowserClient(c)
	go client.Cl.Listen()
}

func NewClient(conn *websocket.Conn) *Client {
	id := uuid.Must(uuid.NewV4()).String()
	sendCallback := map[string]func(WSPacket){}
	recvCallback := map[string]func(WSPacket){}

	return &Client{
		id:   id,
		conn: conn,

		sendCallback: sendCallback,
		recvCallback: recvCallback,

		Done: make(chan struct{}),
	}
}

// NewOverlordClient returns a client connecting to browser. This connection exchanges information between clients and server
func NewBrowserClient(c *websocket.Conn) *BrowserClient {
	return &BrowserClient{
		Cl: NewClient(c),
	}
}

type BrowserClient struct {
	Cl *Client
}

func (c *Client) Listen() {
	for {
		c.conn.SetReadDeadline(time.Now().Add(time.Duration(1)))
		_, rawMsg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("[!] read:", err)
			// TODO: Check explicit disconnect error to break
			close(c.Done)
			break
		}
		wspacket := WSPacket{}
		err = json.Unmarshal(rawMsg, &wspacket)

		if err != nil {
			log.Println("Warn: error decoding", rawMsg)
			continue
		}

		// Check if some async send is waiting for the response based on packetID
		// TODO: Change to read lock.
		//c.sendCallbackLock.Lock()
		callback, ok := c.sendCallback[wspacket.PacketID]
		//c.sendCallbackLock.Unlock()
		if ok {
			go callback(wspacket)
			//c.sendCallbackLock.Lock()
			delete(c.sendCallback, wspacket.PacketID)
			//c.sendCallbackLock.Unlock()
			// Skip receiveCallback to avoid duplication
			continue
		}
		// Check if some receiver with the ID is registered
		if callback, ok := c.recvCallback[wspacket.ID]; ok {
			go callback(wspacket)
		}
	}
}

func main() {
	c := &Client{
		sendCallback: make(map[string]func(req WSPacket)),
		recvCallback: make(map[string]func(req WSPacket)),
		Done:         make(chan struct{}),
	}

	// browser facing port
	go func() {
		http.HandleFunc("/ws", c.WS)
	}()

	http.ListenAndServe(":53000", nil)

	con, _, err := websocket.DefaultDialer.Dial("ws://localhost:9000/wso", nil)
	if err != nil {
		panic(err)
	}
	c.conn = con

}
