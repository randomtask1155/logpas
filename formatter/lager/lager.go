package lager

import (
	"io"
	"time"
	"encoding/json"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"github.com/randomtask1155/logpas/formatter"
)

type Lager struct {
	Reader io.Reader
	Writer io.Writer
}

type Record struct { //https://github.com/cloudfoundry/lager/blob/master/models.go#L19
	Timestamp string `json:"timestamp"`
	Source string `json:"source"`
	Message string `json:"message"`
	LogLevel int `json:"log_level"`
	Data interface{} `json:"data"`
}

func (lager Lager) Reformat(l []byte) ([]byte, error) {
	r := new(Record)
	err := json.Unmarshal(l, &r)
	if err != nil {
		return []byte{}, err
	}

	splitTS := strings.Split(r.Timestamp, ".")
	if len(splitTS)< 2 {
		return []byte{}, fmt.Errorf("invalid timestamp format.  Should be seconds.nanoseconds: %s", r.Timestamp)
	}

	converted, err := strconv.Atoi(splitTS[0])
	if err != nil {
		return []byte{}, err
	}
	seconds := int64(converted)
	converted, err = strconv.Atoi(splitTS[1])
	if err != nil {
		return []byte{}, err
	}

	nano := int64(converted)
	t := time.Unix(seconds, nano)

	data, err := json.Marshal(r.Data)
	if err != nil {
		return []byte{}, err
	}

	newLine := []byte(fmt.Sprintf("%s%s\033[0m:%s:%s:%d: %s\n", formatter.Colors["magenta"],t,
					r.Source,
					r.Message,
					r.LogLevel,
					data))
	return newLine, nil
}
func (lager Lager) Read() error {
	r := bufio.NewReader(lager.Reader)

	for {
		b, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return  err
		} else if err == io.EOF {
			break
		}
		l, err := lager.Reformat(b)
		if err != nil {
			return err
		}
		lager.Write(l)
	}
	return nil
}

func (lager Lager) Write(l []byte) error {
	_, err  := fmt.Fprint(lager.Writer, fmt.Sprintf("%s",l))
	if err != nil {
		return err
	}
	return nil
}
