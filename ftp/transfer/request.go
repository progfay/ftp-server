package transfer

import (
	"fmt"
	"strings"
)

type Request struct {
	Command string
	Message string
}

func ParseRequest(text string) Request {
	s := strings.SplitN(text, " ", 2)

	switch len(s) {
	case 0:
		return Request{}

	case 1:
		return Request{
			Command: strings.ToUpper(s[0]),
		}

	default:
		return Request{
			Command: strings.ToUpper(s[0]),
			Message: s[1],
		}
	}
}

func (req *Request) String() string {
	if req.Message == "" {
		return req.Command
	}
	return fmt.Sprintf("%s %s", req.Command, req.Message)
}
