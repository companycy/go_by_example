package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	// "strings"
)

const (
	DataContextPath = "/data"
	Datapath        = DataContextPath + "/ansible"
	Logpath         = Datapath + "/log"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ws := new(restful.WebService)
	// ws.Route(ws.POST("/api/").Consumes("application/x-www-form-urlencoded").To(apiHandle))
	// ws.Route(ws.POST("/agent/").Consumes("application/x-www-form-urlencoded").To(agentHandle))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

// write shell file to run background: ( (cmd) > log ) &
// then check process: pgrep -fl cmd
func execScriptWithFileName(script, shDirName, fileName string) ([]byte, error) {
	logFile := fmt.Sprintf("%s/%s.log", Logpath, fileName)
	backgroudCmd := fmt.Sprintf("( (%s) > %s ) &", script, logFile)
	shFile := fmt.Sprintf("%s/%s.sh", shDirName, fileName)
	glog.Infof("write %s to file %s", backgroudCmd, shFile)
	if err := ioutil.WriteFile(shFile, []byte(backgroudCmd), 0644); err != nil {
		glog.Errorf("fail to write file %s", shFile)
		return nil, err
	}

	var cmd = exec.Command("/bin/bash", shFile)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Start(); err != nil {
		glog.Errorf("fail to exec cmd %s", cmd)
		return nil, err
	}

	err := cmd.Wait()
	if err != nil {
		glog.Errorf("fail to exec cmd %s", cmd)
	}
	return output.Bytes(), err
}

func execScript(script string) ([]byte, error) {
	file, err := ioutil.TempFile("", "tmp-script")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if _, err = file.WriteString(script); err != nil {
		return nil, err
	}
	file.Close()

	var cmd = exec.Command("/bin/bash", file.Name())
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	err = cmd.Wait()
	return output.Bytes(), err
}
