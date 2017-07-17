package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var duration = 0
var allRes = ""
var lastPer = -1

func durToSec(dur string) (sec int) {
	durAry := strings.Split(dur, ":")
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.Atoi(durAry[0])
	sec = hr * (60 * 60)
	min, _ := strconv.Atoi(durAry[1])
	sec += min * (60)
	second, _ := strconv.Atoi(durAry[2])
	sec += second
	return

}
func getRatio(res string) {
	i := strings.Index(res, "Duration")
	if i >= 0 {
		dur := res[i+10:]
		if len(dur) > 8 {
			dur = dur[0:8]
			duration = durToSec(dur)
			fmt.Println("duration:", duration)
			allRes = ""
		}

	}
	if duration == 0 {
		return
	}
	i = strings.Index(res, "time=")
	if i >= 0 {
		time := res[i+5:]
		if len(time) > 8 {
			time = time[0:8]
			sec := durToSec(time)
			per := (sec * 100) / duration
			if lastPer != per {
				lastPer = per
				fmt.Println("Percentage:", per)
			}

			allRes = ""
		}
	}

}

func main() {
	os.Remove("cmd1.mp4")
	cmdName := "ffmpeg -i 1.mp4  -acodec aac -vcodec libx264  cmd1.mp4 2>&1"
	cmd := exec.Command("sh", "-c", cmdName)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte, 8)
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break

		}
		allRes += string(oneByte)
		getRatio(allRes)

	}

}
