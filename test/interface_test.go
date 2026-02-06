package test

// 已实现接口
var _ Pluginer = (*Plugin1)(nil)

// 未实现接口
// var _ Pluginer = (*Plugin2)(nil)

type Pluginer interface {
	Run()
	Info() string
	Stop() error
}

type Plugin1 struct {
	Name string
}

func (bp *Plugin1) Info() string {
	return bp.Name
}

func (bp *Plugin1) Stop() error {
	return nil
}

func (bp *Plugin1) Run() {
	// do nothing
}

type Plugin2 struct {
	Name string
}

func (bp *Plugin2) Info() string {
	return bp.Name
}
func (bp *Plugin2) Stop() error {
	return nil
}
