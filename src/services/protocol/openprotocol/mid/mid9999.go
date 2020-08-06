package mid

const (
	MID9999 = "9999"
)

type Mid9999 struct {
	BaseMid
}

func (s *Mid9999) Serialize(rev int32) []byte {
	s.Mid = "9999"
	s.Len = HeaderLen

	return s.BaseMid.Serialize(rev)
}
