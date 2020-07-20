package openprotocol

import (
	"fmt"
	"strconv"
)

const (
	MidHeaderLen = 20
	MidTerminal  = 0x00
)

func getMidContext(mid string) (IMid, error) {
	m, exist := MidContext[mid]
	if !exist {
		return nil, fmt.Errorf("Mid Context Not Found: %s ", mid)
	}

	return m, nil
}

var MidContext = map[string]IMid{
	"0001": &Mid0001{},
	"0002": &Mid0002{},
}

type MidHeader struct {
	LEN      int
	MID      string
	Revision int
	NoAck    string
	Station  string
	Spindle  string
	Spare    string
}

func (s *MidHeader) Serialize() string {
	return fmt.Sprintf("%04d%04s%03d%-1s%-2s%-2s%-4s", s.LEN, s.MID, s.Revision, s.NoAck, s.Station, s.Spindle, s.Spare)
}

func (s *MidHeader) Deserialize(str string) error {
	if len(str) != MidHeaderLen {
		return fmt.Errorf("Mid Header Len Error: %s ", str)
	}

	n, _ := strconv.ParseInt(str[0:4], 10, 32)
	s.LEN = int(n) - MidHeaderLen
	s.MID = str[4:8]

	rev, err := strconv.ParseInt(str[8:10], 10, 32)
	if err != nil {
		return err
	}
	s.Revision = int(rev)

	s.NoAck = str[10:11]
	s.Station = str[11:13]
	s.Spindle = str[13:15]
	s.Spare = str[15:19]

	return nil
}

type IMid interface {
	Serialize(rev int) string
	DeserializeData(data string) error
	GetRespChannel() chan Mid
}

type Mid struct {
	IMid
	MidHeader
	chanResp chan Mid
}

func (s *Mid) Serialize(rev int) string {
	s.Revision = rev
	return s.MidHeader.Serialize()
}

func (s *Mid) DeserializeData(data string) error {
	return nil
}

func (s *Mid) GetRespChannel() chan Mid {
	if s.chanResp == nil {
		s.chanResp = make(chan Mid, 2)
	}

	return s.chanResp
}

type Mid9999 struct {
	Mid
}

func (s *Mid9999) Serialize(rev int) string {
	s.MID = "9999"
	s.LEN = MidHeaderLen

	return s.Mid.Serialize(rev)
}

type Mid0001 struct {
	Mid
}

func (s *Mid0001) Serialize(rev int) string {
	s.MID = "0001"
	s.LEN = 0
	return ""
}

func (s *Mid0001) DeserializeData(data string) error {
	return nil
}

type Mid0002 struct {
	Mid
}

type Mid0004 struct {
	Mid
}

type Mid0005 struct {
	Mid
}

type Mid0060 struct {
	Mid
}

type Mid0061 struct {
	Mid
}

type Mid0064 struct {
	Mid
}

type Mid0065 struct {
	Mid
}

type Mid7408 struct {
	Mid
}

type Mid7410 struct {
	Mid
}
