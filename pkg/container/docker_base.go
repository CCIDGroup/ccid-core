package container

type DI interface {
	Check() (*CheckList, error) //查看当前env是否支持操作docker
	PullImage(image string)
}

type D struct {
	CheckList
	ConOpr
}
