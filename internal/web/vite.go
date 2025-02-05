//go:build develop

package web

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gorilla/websocket"
)

type ViteProxy struct {
	origin string
}

func NewViteProxy(host string) *ViteProxy {
	return &ViteProxy{
		origin: host,
	}
}

func (p *ViteProxy) WebSocketHandShake(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.Next()
	}
	url := p.origin + "/ws" + "?token=" + token
	viteConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer viteConn.Close()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clientConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	go p.transferWebSocketMessage(clientConn, viteConn)
	p.transferWebSocketMessage(viteConn, clientConn)

	c.Abort()
}

func (p *ViteProxy) transferWebSocketMessage(src, dest *websocket.Conn) {
	for {
		msgType, msg, err := src.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		err = dest.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}

func (p *ViteProxy) ServeProxyData(c *gin.Context) {
	url := url.URL{
		Scheme:   "http",
		Host:     p.origin,
		Path:     c.Request.URL.Path,
		RawQuery: c.Request.URL.RawQuery,
	}
	resp, err := http.Get(url.String())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}

type ViteHTMLRender struct {
	origin string
}

func NewViteHTMLRender(origin string) *ViteHTMLRender {
	return &ViteHTMLRender{
		origin: origin,
	}
}

func (r *ViteHTMLRender) Instance(name string, data any) render.Render {
	tmpl, err := r.fetchTemplate(name)
	if err != nil {
		return render.HTML{
			Name: "error",
			Data: data,
		}
	}
	return render.HTML{
		Template: tmpl,
		Name:     name,
		Data:     data,
	}
}

func (r *ViteHTMLRender) fetchTemplate(path string) (*template.Template, error) {
	url := url.URL{
		Scheme: "http",
		Host:   r.origin,
		Path:   path + ".html",
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(path).Parse(string(body))
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
