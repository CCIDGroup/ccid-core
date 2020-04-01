package container

type D interface {
	Check() CheckList //查看当前env是否支持操作docker
}

type CheckList struct {
	DockerEngineVersion string //docker 版本
	IsAvailable         bool   //docker 是否可用
	DiskSpace           uint   //存储位置的磁盘大小
	DiskSpaceUnit       string //磁盘大小的单位
	ImageStorePath      string //磁盘存储位置
	ContainerStorePath  string //容器存储位置
}
