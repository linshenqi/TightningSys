package openprotocol

import (
	"github.com/linshenqi/asbt/src/services/protocol/openprotocol/mid"
	"time"
)

type ConnParams struct {
	WriteItv          time.Duration
	KeepAliveItv      time.Duration
	MidReqTimeout     time.Duration
	MaxKeepAliveCheck int
}

var DefaultConnParam = ConnParams{
	WriteItv:          300 * time.Millisecond,
	KeepAliveItv:      8 * time.Second,
	MidReqTimeout:     500 * time.Millisecond,
	MaxKeepAliveCheck: 3,
}

var DefaultMidDefine = map[string]int{
	mid.MID0001: 1,
	mid.MID0002: 1,
	mid.MID0004: 1,
	mid.MID0005: 1,
	mid.MID9999: 1,
}

var MidRequestChannels = map[string]chan mid.IMid{
	mid.MID0001: make(chan mid.IMid),
}

type handlerMidNotify func(data interface{})
type handlerStatusNotify func(data interface{})
type handlerLogNotify func(logType string, logContent string)

type Handlers struct {
	handlerMidNotify    handlerMidNotify
	handlerStatusNotify handlerStatusNotify
	handlerLogNotify    handlerLogNotify
}
