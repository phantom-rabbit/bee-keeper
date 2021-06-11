package api

import (
	"bee-keeper/module"
	"bee-keeper/service"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsSsh(c *gin.Context) {
	id := c.Query("id")
	sshConfig, err := module.GetNodeById(id)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if handleError(c, err) {
		return
	}
	defer wsConn.Close()

	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))

	if wshandleError(wsConn, err) {
		return
	}

	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if wshandleError(wsConn, err) {
		return
	}

	client, err := service.NewSSHClient(sshConfig.User, sshConfig.Password, sshConfig.Ip, sshConfig.Port)
	if wshandleError(wsConn, err) {
		return
	}
	defer client.Close()

	conn, err := service.NewSSHConn(cols, rows, client)
	if wshandleError(wsConn, err) {
		return
	}
	defer conn.Close()

	quitChan := make(chan bool, 3)
	var logBuff = new(bytes.Buffer)

	go conn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go conn.SendComboOutput(wsConn, quitChan)
	go conn.SessionWait(quitChan)

	<-quitChan
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func jsonError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": msg})
}

func wshandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		logrus.WithError(err).Error("handler ws ERROR:")
		dt := time.Now().Add(time.Second)
		if err := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); err != nil {
			logrus.WithError(err).Error("websocket writes control message failed:")
		}
		return true
	}
	return false
}