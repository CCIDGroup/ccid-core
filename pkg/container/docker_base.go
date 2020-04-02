package container

type DI interface {
	Check() *CheckList //查看当前env是否支持操作docker
	PullImage()
}

type D struct {
	CheckList
	C
}
