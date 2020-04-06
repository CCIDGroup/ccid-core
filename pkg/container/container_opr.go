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

//container的基本数据类型
type ConOpr struct {
	ID       string   // container ID(when started, It has an ID)
	Name     string   // container name
	Image    string   // container image name, format: image_name:tag
	Endpoint string   // Image endpoint, such as: docker.com/nginx
	Env      []string //运行docker所需要的环境变量
	Cmd      []string //创建container时候需要传递的参数
	Options  string   //创建container时的可选参数
	Ports    []string //端口映射
	Volumes  []string //磁盘映射
	Result   *chan string
}

func (c *ConOpr) PullImage() {

}
