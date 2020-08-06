package mid

const (
	MID0005 = "0005"
)

type Mid0005 struct {
	BaseMid

	TargetMid string
}

func (s *Mid0005) DeserializeData(buf []byte) error {
	if err := s.BaseMid.DeserializeData(buf); err != nil {
		return err
	}

	data := string(buf)
	// rev 0/1
	if s.Rev >= 0 {
		if err := extractMidDataValue(data, DataTypeString, &s.TargetMid, 0, 4); err != nil {
			return err
		}
	}

	return nil
}
