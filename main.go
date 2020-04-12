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
package main

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	buf := new(bytes.Buffer)
	_, err := git.PlainClone("C:\\Users\\xuren\\Downloads\\t1", false, &git.CloneOptions{
		URL:      "https://github.com/CCIDGroup/ccid-core.git",
		Progress: buf,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(buf.String())
}



