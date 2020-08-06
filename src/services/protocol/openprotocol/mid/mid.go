package mid

import (
	"fmt"
	"strconv"
)

const (
	HeaderLen   = 20
	MidTerminal = 0x00

	DataTypeString  = "DataTypeString"
	DataTypeInt32   = "DataTypeInt32"
	DataTypeFloat32 = "DataTypeFloat32"
	DataTypeBool    = "DataTypeBool"
)

func extractMidDataValue(rawData string, dataType string, data interface{}, begin int, end int) error {
	if len(rawData) == 0 {
		return fmt.Errorf("extractMidDataValue Err: Raw Data Is Empty ")
	}

	if begin < 0 {
		return fmt.Errorf("extractMidDataValue Err: Begin Index Should Be Greater Than 0: %d ", begin)
	}

	if end > len(rawData) {
		return fmt.Errorf("extractMidDataValue Err: End Index Should Be Less Than Data Length: %d ", end)
	}

	valStr := rawData[begin:end]
	switch dataType {
	case DataTypeString:
		data = valStr

	case DataTypeInt32:
		val, err := strconv.ParseInt(valStr, 10, 32)
		if err != nil {
			return err
		}

		data = val

	case DataTypeFloat32:
		val, err := strconv.ParseFloat(valStr, 32)
		if err != nil {
			return err
		}

		data = val

	case DataTypeBool:
		val, err := strconv.ParseBool(valStr)
		if err != nil {
			return err
		}

		data = val

	default:
		return fmt.Errorf("extractMidDataValue Err: Mid Data Type Error: %s ", dataType)
	}

	return nil
}

type Header struct {
	Len     int32
	Mid     string
	Rev     int32
	NoAck   string
	Station string
	Spindle string
	Spare   string
}

func (s *Header) Serialize() string {
	return fmt.Sprintf("%04d%04s%03d%-1s%-2s%-2s%-4s", s.Len, s.Mid, s.Rev, s.NoAck, s.Station, s.Spindle, s.Spare)
}

func (s *Header) DeserializeHeader(str string) error {
	if len(str) != HeaderLen {
		return fmt.Errorf("Mid Header Len Error: %s ", str)
	}

	if err := extractMidDataValue(str, DataTypeInt32, &s.Len, 0, 4); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeString, &s.Mid, 4, 8); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeInt32, &s.Rev, 8, 10); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeString, &s.NoAck, 10, 11); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeString, &s.Station, 11, 13); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeString, &s.Spindle, 13, 15); err != nil {
		return err
	}

	if err := extractMidDataValue(str, DataTypeString, &s.Spare, 15, 19); err != nil {
		return err
	}

	return nil
}

type IMid interface {
	GetHeader() Header
	Serialize(rev int32) []byte
	DeserializeData(buf []byte) error
}

type BaseMid struct {
	IMid
	Header
}

func (s *BaseMid) GetHeader() Header {
	return s.Header
}

func (s *BaseMid) Serialize(rev int32) []byte {
	s.Rev = rev
	return []byte(s.Header.Serialize())
}

func (s *BaseMid) DeserializeData(buf []byte) error {
	dataLen := s.Len - HeaderLen
	if dataLen != int32(len(buf)) {
		return fmt.Errorf("Data Len Error: %s ", buf)
	}

	return nil
}
