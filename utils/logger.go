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
	"io"
	"os"
	"reflect"
	"sync"
)

type Log struct {
	logger       *zap.Logger
	wg           sync.WaitGroup
	rev          string
	logPath      string
	logInfoPath  string
	logErrorPath string
}

func (l *Log) InitLog() *Log {
	_, err := os.Stat(l.logPath)
	if err != nil {
		// 创建文件夹
		err := os.Mkdir(l.logPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "",
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
		Level:            atom,                               // 日志级别
		Development:      true,                               // 开发模式，堆栈跟踪
		Encoding:         "console",                          // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                      // 编码器配置
		OutputPaths:      []string{"stdout", l.logInfoPath},  // 输出到指定文件 stdout（标准输出，正常颜色）
		ErrorOutputPaths: []string{"stderr", l.logErrorPath}, // stderr（错误输出，红色）
	}
	// 构建日志
	l.logger, _ = config.Build()
	return l
}

func (l *Log) LogStream(reader io.Reader) *chan string {
	r := make(chan string)
	go func() {

		//out, _ := os.Create(l.logPath)
		//defer out.Close()
		for {
			buf := make([]byte, 1024)
			// 循环读取文件
			n, err2 := reader.Read(buf)
			if err2 != nil { // io.EOF表示文件末尾
				break
			}
			r <- string(buf[:n])
		}
		close(r)
	}()
	return &r
}

func (l *Log) LogMsg(msg string) {
	go l.logger.Info(msg)
}

func (l *Log) LogError(err error) {
	go l.logger.Error(err.Error())
}

func (l *Log) LogOne(desc string, u interface{}) {
	go l.logObj(desc, u)
}

func (l *Log) LogList(list map[string]interface{}) {
	for k, v := range list {
		go l.logObj(k, v)
	}
}

//log interface 类型
func (l *Log) logObj(desc string, u interface{}) {
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
}
