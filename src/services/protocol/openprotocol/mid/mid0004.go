package mid

const (
	MID0004 = "0004"
)

type Mid0004 struct {
	BaseMid
	TargetMid string
	ErrCode   string
}

func (s *Mid0004) DeserializeData(buf []byte) error {
	if err := s.BaseMid.DeserializeData(buf); err != nil {
		return err
	}

	data := string(buf)
	// rev 0/1
	if s.Rev >= 0 {
		if err := extractMidDataValue(data, DataTypeString, &s.TargetMid, 0, 4); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeString, &s.ErrCode, 4, 6); err != nil {
			return err
		}
	}

	return nil
}
