/*
 * Copyright 2020 The CCID Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http: //www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package container

type CheckList struct {
	DockerEngineVersion string  //docker 版本
	IsAvailable         bool    //docker 是否可用
	FreeDiskSpace       float64 //存储位置的磁盘大小
	DiskSpaceUnit       string  //磁盘大小的单位
	ImageStorePath      string  //磁盘存储位置
	ContainerStorePath  string  //容器存储位置
}

//func (cl *CheckList) Check() (*CheckList, error) {
//	return GetDockerEngineInfo()
//}
