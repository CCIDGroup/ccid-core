package utils

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
//func DiskUsage(path string) (disk DiskStatus) {
//	fs := syscall.Statfs_t{}
//	err := syscall.Statfs(path, &fs)
//	if err != nil {
//		return
//	}
//	disk.All = fs.Blocks * uint64(fs.Bsize)
//	disk.Free = fs.Bfree * uint64(fs.Bsize)
//	disk.Used = disk.All - disk.Free
//	return
//}
//
//const (
//	B  = 1
//	KB = 1024 * B
//	MB = 1024 * KB
//	GB = 1024 * MB
//)
//
//func GetFreeDiskSpace(path string) float64 {
//	disk := DiskUsage(path)
//	return float64(disk.Free) / float64(MB)
//}


//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址



