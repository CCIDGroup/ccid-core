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
package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"reflect"
	"sync"
)

const (
	LOGPATH = "./logs"
)

type L struct {
	logger *zap.Logger
	wg     sync.WaitGroup
}

func init() {
	_, err := os.Stat(LOGPATH)
	if err != nil {
		// 创建文件夹
		err := os.Mkdir(LOGPATH, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
}

//初始化日志模块
func (l *L) InitL() (*L, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                       // 日志级别
		Development:      true,                                       // 开发模式，堆栈跟踪
		Encoding:         "json",                                     // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                              // 编码器配置
		OutputPaths:      []string{"stdout", LOGPATH + "/info.log"},  // 输出到指定文件 stdout（标准输出，正常颜色）
		ErrorOutputPaths: []string{"stderr", LOGPATH + "/error.log"}, // stderr（错误输出，红色）
	}

	// 构建日志
	logger, err := config.Build()

	l.logger = logger
	return l, err
}

func (l *L) LogMsg(desc string, msg string) {
	l.wg.Add(1)
	go func() {
		l.logger.Info(desc, zap.String("message", msg))
		l.wg.Done()
	}()
	l.wg.Wait()

}

func (l *L) LogOne(desc string, u interface{}) {
	l.wg.Add(1)
	go l.logObj(desc, u)
	l.wg.Wait()

}

func (l *L) LogList(list map[string]interface{}) {
	l.wg.Add(len(list))

	for k, v := range list {
		go l.logObj(k, v)
	}
	l.wg.Wait()

}

//log interface 类型
func (l *L) logObj(desc string, u interface{}) {
	keys := reflect.TypeOf(u)
	values := reflect.ValueOf(u)
	m := &[]zap.Field{}
	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	fmt.Println(keys.NumField())
	for i := 0; i < keys.NumField(); i++ {
		field := keys.Field(i)
		value := values.Field(i)
		*m = append(*m, zap.String(field.Name, fmt.Sprintf("%v", value)))
	}
	fmt.Println(m)
	l.logger.Info(desc, *m...)
	l.wg.Done()
}
