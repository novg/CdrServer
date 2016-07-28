package dbclient

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Parser write message to stdout
type callInfo struct {
	datetime string
	duration string
	seg      string
	sop      string
	dest     string
	numin    string
	numout   string
	str1     string
	str2     string
}

var (
	// CallInfo incapsulates info about calls
	CallInfo *callInfo
	year     int
)

func init() {
	CallInfo = new(callInfo)
	year = time.Now().Year()
}

func (h *callInfo) Write(p []byte) (n int, err error) {
	sp := strings.Fields(string(p))
	parse(sp, h)
	sendToDB(h)
	return len(p), nil
}

func parse(line []string, h *callInfo) {
	if len(line) < 6 {
		return
	}

	var month, day, hours, minutes string
	_, err := fmt.Sscanf(line[0], "%2s%2s%2s%2s", &month, &day, &hours, &minutes)
	if err != nil {
		log.Printf("%s is wrong datetime string", line[0])
	}
	h.datetime = fmt.Sprintf("%d-%s-%s %s:%s", year, month, day, hours, minutes)

	var seconds string
	_, err = fmt.Sscanf(line[1], "%1s%2s%2s", &hours, &minutes, &seconds)
	if err != nil {
		log.Printf("%s is wrong duration string", line[1])
	}
	h.duration = fmt.Sprintf("%d%s:%s:%s", 0, hours, minutes, seconds)

	f := func(r rune) bool {
		return r < 'A' || r > 'z'
	}
	if strings.IndexFunc(line[2], f) != -1 {
		h.seg = ""
	} else {
		h.seg = line[2]
	}

	if len(line) < 9 {
		h.sop = ""
	} else {
		h.sop = line[3]
	}

	if len(line) < 7 || strings.IndexFunc(line[2], f) == -1 && len(line) != 9 {
		h.dest = ""
	} else {
		h.dest = line[len(line)-5]
	}

	h.numin = line[len(line)-4]
	h.numout = line[len(line)-3]
	h.str1 = line[len(line)-2]
	h.str2 = line[len(line)-1]
}
