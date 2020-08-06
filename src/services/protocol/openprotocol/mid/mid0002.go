package mid

const (
	MID0002 = "0002"
)

type Mid0002 struct {
	BaseMid

	CellID                    int32
	ChannelID                 int32
	ControllerName            string
	SupplierCode              string
	OpenProtocolVersion       string
	ControllerSoftwareVersion string
	ToolSoftwareVersion       string
	RBUType                   string
	ControllerSerialNumber    string
}

func (s *Mid0002) DeserializeData(buf []byte) error {
	if err := s.BaseMid.DeserializeData(buf); err != nil {
		return err
	}

	data := string(buf)

	// rev0/1
	if s.Rev >= 0 {
		if err := extractMidDataValue(data, DataTypeInt32, &s.CellID, 2, 6); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeInt32, &s.ChannelID, 8, 10); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeString, &s.ControllerName, 12, 37); err != nil {
			return err
		}
	}

	// rev2
	if s.Rev >= 2 {
		if err := extractMidDataValue(data, DataTypeString, &s.SupplierCode, 39, 42); err != nil {
			return err
		}
	}

	// rev3
	if s.Rev >= 3 {
		if err := extractMidDataValue(data, DataTypeString, &s.OpenProtocolVersion, 44, 63); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeString, &s.ControllerSoftwareVersion, 65, 84); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeString, &s.ToolSoftwareVersion, 86, 105); err != nil {
			return err
		}
	}

	// rev4
	if s.Rev >= 4 {
		if err := extractMidDataValue(data, DataTypeString, &s.RBUType, 107, 131); err != nil {
			return err
		}

		if err := extractMidDataValue(data, DataTypeString, &s.ControllerSerialNumber, 133, 143); err != nil {
			return err
		}
	}

	return nil
}
