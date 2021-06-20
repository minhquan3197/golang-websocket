package cmd

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ws "github.com/minhquan3197/golang-chat/internal/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func joinRoom(c echo.Context) error {
	ws.ServeWs(c.Response(), c.Request(), c.Param("roomId"))
	return nil
}

// Excute func
func Excute() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", hello)
	e.GET("/room/:roomId", joinRoom)
	e.Logger.Fatal(e.Start(":1323"))
}
