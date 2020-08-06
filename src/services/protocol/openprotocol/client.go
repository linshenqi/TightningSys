package openprotocol

import (
	"fmt"
	"github.com/linshenqi/asbt/src/services/protocol/openprotocol/mid"
	"net"
	"time"
)

const (
	DailTimeOut = 1 * time.Second
	BufferSize  = 65535
)

func CreateOPClient(addr string) *OPClient {
	return &OPClient{
		bufReceive:         make(chan byte, BufferSize),
		bufWrite:           make(chan []byte, 4096),
		addr:               addr,
		asyncConnectDone:   make(chan bool),
		asyncHandleDone:    make(chan bool),
		asyncWriteDone:     make(chan bool),
		midDefines:         DefaultMidDefine,
		param:              DefaultConnParam,
		midRequestChannels: MidRequestChannels,
		handlers:           Handlers{},
		subscribeMids:      []string{},
	}
}

type OPClient struct {
	sockClient net.Conn
	bufReceive chan byte
	bufWrite   chan []byte
	addr       string
	param      ConnParams
	midDefines map[string]int

	asyncConnectDone chan bool
	asyncHandleDone  chan bool
	asyncWriteDone   chan bool

	handlers Handlers

	midRequestChannels map[string]chan mid.IMid
	subscribeMids      []string
}

func (s *OPClient) Start() error {
	go s.asyncConnect()
	return nil
}

func (s *OPClient) Stop() error {
	return nil
}

func (s *OPClient) SetParam(param ConnParams) {
	s.param = param
}

func (s *OPClient) SetMidDefine() {

}

func (s *OPClient) DoMidRequest(m mid.IMid) (mid.IMid, error) {
	rev, err := s.getRevByMidDefine(m.GetHeader().Mid)
	if err != nil {
		return nil, err
	}

	ch, err := s.getMidRequestChannel(m.GetHeader().Mid)
	if err != nil {
		return nil, err
	}

	s.doPostBuf(m.Serialize(int32(rev)))

	select {
	case <-time.After(s.param.MidReqTimeout):
		{
			return nil, fmt.Errorf("Mid Request Timeout: %s ", m.GetHeader().Mid)
		}

	case midResp := <-ch:
		{
			return midResp, nil
		}

	default:
		return nil, fmt.Errorf("DoMidRequest: Error Unknown ")
	}
}

func (s *OPClient) getRevByMidDefine(mid string) (int, error) {
	rev, exist := s.midDefines[mid]
	if !exist {
		return -1, fmt.Errorf("Mid %s Not Supported ", mid)
	}

	return rev, nil
}

func (s *OPClient) onRecvSubscribedMid() {

}

func (s *OPClient) onStatus() {

}

func (s *OPClient) onLog() {

}

func (s *OPClient) doSubscribeMids() {
	//for _, v := range s.subscribeMids {
	//
	//}
}

func (s *OPClient) asyncConnect() {

	for {
		select {
		case <-s.asyncConnectDone:
			return

		default:
			s.doHandleConnect()
		}
	}
}

func (s *OPClient) asyncHandle() {
	for {
		select {
		case <-s.asyncHandleDone:
			return

		default:
			if err := s.doHandleRecv(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func (s *OPClient) asyncWrite() {
	for {
		select {
		case <-s.asyncWriteDone:
			return

		case buf := <-s.bufWrite:
			_, err := s.sockClient.Write(append(buf, mid.MidTerminal))
			if err != nil {

			}

			time.Sleep(s.param.WriteItv)
		}
	}
}

func (s *OPClient) doHandleConnect() {
	go s.asyncHandle()
	go s.asyncWrite()

	for {
		conn, err := net.DialTimeout("tcp", s.addr, DailTimeOut)
		if err != nil {
			time.Sleep(DailTimeOut)
			continue
		}

		s.sockClient = conn
		break
	}

	if err := s.doRead(); err != nil {

	}
}

func (s *OPClient) doRead() error {
	buf := make([]byte, BufferSize)
	for {
		n, err := s.sockClient.Read(buf)
		if err != nil {
			return err
		}

		targetBuf := buf[0:n]
		for _, v := range targetBuf {
			s.bufReceive <- v
		}
	}
}

func (s *OPClient) doPostBuf(buf []byte) {
	s.bufWrite <- buf
}

func (s *OPClient) doHandleRecv() error {
	// handle header
	var bufHeader []byte
	for i := 0; i < mid.HeaderLen; i++ {
		bufHeader = append(bufHeader, <-s.bufReceive)
	}

	header := mid.Header{}
	if err := header.DeserializeHeader(string(bufHeader)); err != nil {
		return err
	}

	// handle data
	dataLen := header.Len - mid.HeaderLen + 1
	var bufData []byte
	for i := 0; i < int(dataLen); i++ {
		bufData = append(bufData, <-s.bufReceive)
	}

	if err := s.doHandleMid(&header, bufData); err != nil {
		return err
	}

	return nil
}

func (s *OPClient) doHandleMid(header *mid.Header, data []byte) error {

	switch header.Mid {
	case mid.MID0002:
		mid0002 := mid.Mid0002{}
		if err := mid0002.DeserializeData(data); err != nil {
			return err
		}

		ch, err := s.getMidRequestChannel(mid.MID0001)
		if err != nil {
			return err
		}

		ch <- &mid0002

	case mid.MID0004:
		mid0004 := mid.Mid0004{}
		if err := mid0004.DeserializeData(data); err != nil {
			return err
		}

		ch, err := s.getMidRequestChannel(mid0004.TargetMid)
		if err != nil {
			return err
		}

		ch <- &mid0004

	case mid.MID0005:
		mid0005 := mid.Mid0005{}
		if err := mid0005.DeserializeData(data); err != nil {
			return err
		}

		ch, err := s.getMidRequestChannel(mid0005.TargetMid)
		if err != nil {
			return err
		}

		ch <- &mid0005
	}

	return nil
}

func (s *OPClient) getMidRequestChannel(mid string) (chan mid.IMid, error) {
	ch, exist := s.midRequestChannels[mid]
	if !exist {
		return nil, fmt.Errorf("Mid Channel Not Found: %s ", mid)
	}

	return ch, nil
}
