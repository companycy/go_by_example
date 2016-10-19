package main

// https://mp.weixin.qq.com/s?__biz=MzAxMzc4Mzk1Mw==&mid=2649836694&idx=4&sn=6886c47866ff85b9e32dea02bd1bf356&chksm=8398aa68b4ef237edf3743103b7265f806ece622252ea1b4cebff29c188b2ddd2214e2ba5a01&scene=0&pass_ticket=FxmWugrv%2BSbJKKew5SIOwjEg2cSekgSRglYh1D%2B9VNNHxzAA70YmlYpf8d2%2B2hcr

import (
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(wsHandler))
	http.ListenAndServe(":7224", nil)

}

func wsHandler(ws *websocket.Conn) {
	for {
		s := ""
		if e := websocket.Message.Receive(ws, &s); e != nil {
			break

		}
		c := exec.Command("go", "run", s)
		if s, e := c.StderrPipe(); e == nil {
			go io.Copy(ws, s)

		}
		if s, e := c.StdoutPipe(); e == nil {
			if e := c.Start(); e == nil {
				io.Copy(ws, s)

			}

		}

	}
	ws.Close()

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTpl.Execute(w, r.Host)

}

var homeTpl = template.Must(template.New("ws").Parse(`<html>
<textarea id=idout rows=24 cols=72></textarea><hr>
<input id=idin type=search placeholder='Enter a unix command '
onchange='send(this.value)'></input>
<script>
        var vout=document.getElementById('idout')
        var vin =  document.getElementById('idin')
        var wscon = new WebSocket("ws://{{.}}/ws")
        wscon.onclose = function(e) {vout.value = 'websocket closed'}
        wscon.onmessage = function(e) { vout.value += e.data }
        function send(s) {
                vout.value = ""
                wscon.send(s)
        }
</script>
`))
