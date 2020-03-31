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
 *
 * Project: ccid-core
 * File Created: China Standard Time UT+8:00 - Sunday, 22nd March 2020 4:44:12 pm
 * Author: Lucas Ren (nichokiki@hotmail.com)
 * Last Modified: China Standard Time UT+8:00 - Sunday, 22nd March 2020 4:44:18 pm
 * Modified By: Lucas Ren (nichokiki@hotmail.com>)
 * Copyright 2020 - 2020 CCID Group
 */
package main

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		zap.L().Info("hello world")

		fmt.Println("hello world")
	}

	fmt.Println(cli)
	defer cli.Close()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}
