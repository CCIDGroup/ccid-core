package docker

type Model struct {
	ID       string
	Image    string   // docker image name, format: image_name:tag
	Endpoint string   // Image endpoint, such as: docker.com/nginx
	Env      []string //运行docker所需要的环境变量
	Cmd      []string //创建container时候需要传递的参数
	Options  string   //创建container时的可选参数
	Ports    []string //端口映射
	Volumes  []string //磁盘映射
	CodePath string   //code path
}