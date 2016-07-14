package dbclient

import (
	"fmt"
	"regexp"
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
	CallInfo  *callInfo
	rdate     *regexp.Regexp
	rduration *regexp.Regexp
	rspace    *regexp.Regexp
	rletter   *regexp.Regexp
	year      int
)

func init() {
	CallInfo = new(callInfo)
	rdate = regexp.MustCompile(`(\d{2})(\d{2})(\d{2})(\d{2})`)
	rduration = regexp.MustCompile(`(\d{1})(\d{2})(\d{2})`)
	rspace = regexp.MustCompile(`\s+`)
	rletter = regexp.MustCompile(`[[:alpha:]]`)
	year = time.Now().Year()
}

func (h *callInfo) Write(p []byte) (n int, err error) {
	p = rspace.ReplaceAll(p, []byte(" "))
	sp := strings.TrimSpace(string(p))
	parse(strings.Split(sp, " "), h)

	fmt.Println(h.datetime + " " +
		h.duration + "\t" +
		h.seg + "\t" +
		h.sop + "\t" +
		h.dest + "\t" +
		h.numin + "\t\t" +
		h.numout + "\t\t" +
		h.str1 + " " +
		h.str2)

	return len(p), nil
}

func parse(line []string, h *callInfo) {
	if len(line) < 6 {
		return
	}

	temp := rdate.FindStringSubmatch(line[0])
	h.datetime = fmt.Sprintf("%d-%s-%s %s:%s", year, temp[1], temp[2], temp[3], temp[4])

	temp = rduration.FindStringSubmatch(line[1])
	h.duration = fmt.Sprintf("%d%s:%s:%s", 0, temp[1], temp[2], temp[3])

	if !rletter.MatchString(line[2]) {
		h.seg = ""
	} else {
		h.seg = line[2]
	}

	if len(line) < 9 {
		h.sop = "NULL"
	} else {
		h.sop = line[3]
	}

	if len(line) < 7 || rletter.MatchString(line[2]) && len(line) != 9 {
		h.dest = ""
	} else {
		h.dest = line[len(line)-5]
	}

	h.numin = line[len(line)-4]
	h.numout = line[len(line)-3]
	h.str1 = line[len(line)-2]
	h.str2 = line[len(line)-1]
}
