package main

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
