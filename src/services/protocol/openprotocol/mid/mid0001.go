package mid

const (
	MID0001 = "0001"
)

type Mid0001 struct {
	BaseMid
}

func (s *Mid0001) Serialize(rev int32) []byte {
	s.Mid = MID0001
	s.Len = HeaderLen
	return s.BaseMid.Serialize(rev)
}

func (s *Mid0001) DeserializeData(data string) error {
	return nil
}
