package base

type TightingController struct {
	BaseDevice
}

type TightingTool struct {
	BaseDevice
}

// 工具使能控制
func (s *TightingTool) ToolControl(enable bool) error {
	return nil
}

// 设置pset
func (s *TightingTool) SetPSet(pset int) error {
	return nil
}

// 设置job
func (s *TightingTool) SetJob(job int) error {
	return nil
}

// 模式选择: job/pset
func (s *TightingTool) ModeSelect(mode string) error {
	return nil
}

// 取消job
func (s *TightingTool) AbortJob() error {
	return nil
}

// 设置pset次数
func (s *TightingTool) SetPSetBatch(pset int, batch int) error {
	return nil
}

// pset列表
func (s *TightingTool) GetPSetList() ([]int, error) {
	return nil, nil
}

// job列表
func (s *TightingTool) GetJobList() ([]int, error) {
	return nil, nil
}

// 追溯信息设置
func (s *TightingTool) TraceSet(str string) error {
	return nil
}
