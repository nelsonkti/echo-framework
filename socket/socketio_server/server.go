package socketio_server

import (
	"echo-framework/lib/logger"
	"github.com/googollee/go-socket.io"
	engineio "github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	_ "github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const NameSpace = "/"

type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var ClientNode sync.Map
var UserClientNodeKey sync.Map

func Start(port int) {

	wt := websocket.Default
	pt := polling.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			wt,
			pt,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect(NameSpace, func(s socketio.Conn) error {
		//logger.Sugar.Warn("/:open: id: "+s.ID())
		return nil
	})

	server.OnError(NameSpace, func(s socketio.Conn, e error) {
		//logger.Sugar.Error("/:error: id: " + s.ID() + "; message:" + e.Error())
	})

	//断开连接
	server.OnDisconnect(NameSpace, func(s socketio.Conn, reason string) {

	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:8000...")
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), nil))
}

func StopDevice() {

	ClientNode.Range(func(key, value interface{}) bool {
		ClientNode.Delete(key)
		return true
	})

	UserClientNodeKey.Range(func(key, value interface{}) bool {
		UserClientNodeKey.Delete(key)
		return true
	})

	logger.Sugar.Info("stop device")
}
