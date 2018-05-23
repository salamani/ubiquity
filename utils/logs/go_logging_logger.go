/**
 * Copyright 2017 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logs

import (
    "github.com/op/go-logging"
    "io"
     "github.com/IBM/ubiquity/resources"
     "runtime"
     "bytes"
     "strconv"
     "sync"
    
)

const (
    traceEnter = "ENTER"
    traceExit = "EXIT"
)

type goLoggingLogger struct {
    logger *logging.Logger
}

func newGoLoggingLogger(level Level, writer io.Writer) *goLoggingLogger {
    newLogger := logging.MustGetLogger("")
    newLogger.ExtraCalldepth = 1
    format := logging.MustStringFormatter("%{time:2006-01-02 15:04:05.999} %{level:.5s} %{pid} %{shortfile} %{shortpkg}::%{shortfunc} %{message}")
    backend := logging.NewLogBackend(writer, "", 0)
    backendFormatter := logging.NewBackendFormatter(backend, format)
    backendLeveled := logging.AddModuleLevel(backendFormatter)
    backendLeveled.SetLevel(getLevel(level), "")
    newLogger.SetBackend(backendLeveled)
    return &goLoggingLogger{newLogger}
}


func getContextValues(context resources.RequestContext) Args{
	return Args {{Name: "request_id", Value:context.Id}}

}

func getIdOfContext(args []Args) (int, int){
	for args_index, a := range args {
		for value_index , v := range a {
			if v.Name == "context" && v.Value != nil{
				return args_index, value_index
			}
		}
	}
	return -1, -1
}


func getArgsWithContext(args []Args) []Args{
	args_index, value_index := getIdOfContext(args)
	if args_index == -1 || value_index == -1 {
		return args
	}
	
	context := args[args_index][value_index].Value.(resources.RequestContext)
	context_parsed := getContextValues(context)
	context_args_with_parsed_context := append(append(args[args_index][:value_index], args[args_index][value_index+1:]...), context_parsed...)
	args_without_context := append(args[:args_index],args[args_index+1:]...)
	new_args := append(args_without_context, context_args_with_parsed_context)
	return new_args
}

func GetGoID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}

var GoIdToRequestIdMap = new(sync.Map)

func (l *goLoggingLogger) Debug(str string, args ...Args) {
	args = getArgsWithContext(args)
	go_id :=  GetGoID()
	request_id, exists := GoIdToRequestIdMap.Load(go_id)
	if !exists {
		request_id = "XXXXXXX"
	}
	
	new_args := append(args, Args{{Name: "go Id ", Value: GetGoID()}, {Name: "request-id", Value: request_id}})
    l.logger.Debugf(str + " %v", new_args)
}

func (l *goLoggingLogger) Info(str string, args ...Args) {
	args = getArgsWithContext(args)
	go_id :=  GetGoID()
	request_id, exists := GoIdToRequestIdMap.Load(go_id)
	if !exists {
		request_id = "XXXXXXX"
	}
	
	new_args := append(args, Args{{Name: "go Id ", Value: GetGoID()}, {Name: "request-id", Value: request_id}})
    l.logger.Infof(str + " %v", new_args)
}

func (l *goLoggingLogger) Error(str string, args ...Args) {
	args = getArgsWithContext(args)
	go_id :=  GetGoID()
	request_id, exists := GoIdToRequestIdMap.Load(go_id)
	if !exists {
		request_id = "XXXXXXX"
	}
	
	new_args := append(args, Args{{Name: "go Id ", Value: GetGoID()}, {Name: "request-id", Value: request_id}})
    l.logger.Errorf(str + " %v", new_args)
}

func (l *goLoggingLogger) ErrorRet(err error, str string, args ...Args) error {
	args = getArgsWithContext(args)
    l.logger.Errorf(str + " %v ", append(args, Args{{"error", err}}))
    return err
}

func (l *goLoggingLogger) Trace(level Level, args ...Args) func() {
	args = getArgsWithContext(args)
	go_id :=  GetGoID()
	request_id, exists := GoIdToRequestIdMap.Load(go_id)
	if !exists {
		request_id = "XXXXXXX"
	}
	
	new_args := append(args, Args{{Name: "go Id ", Value: GetGoID()}, {Name: "request-id", Value: request_id}})
    switch level {
    case DEBUG:
        l.logger.Debug(traceEnter, new_args)
        return func() { l.logger.Debug(traceExit, new_args) }
    case INFO:
        l.logger.Info(traceEnter, new_args)
        return func() { l.logger.Info(traceExit, new_args) }
    case ERROR:
        l.logger.Error(traceEnter, new_args)
        return func() { l.logger.Error(traceExit, new_args) }
    default:
        panic("unknown level")
    }
}

func getLevel(level Level) logging.Level {
    switch level {
    case DEBUG:
        return logging.DEBUG
    case INFO:
        return logging.INFO
    case ERROR:
        return logging.ERROR
    default:
        panic("unknown level")
    }
}
