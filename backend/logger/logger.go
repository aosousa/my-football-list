package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// FileLog has all the important information to store in a log file
type FileLog struct {
	EventID string
	Text    string
	Data    map[string]interface{}
	Time    time.Time
	Message []interface{}
	Level   string
	Trace   map[string]string
}

// Debug logs a debugging message
func Debug(eventID interface{}, data map[string]interface{}, param ...interface{}) {
	var fileLog FileLog
	fileLog.Level = "Debug"
	processLog(eventID, data, param, &fileLog)
}

// Info logs an informative message
func Info(eventID interface{}, data map[string]interface{}, param ...interface{}) {
	var fileLog FileLog
	fileLog.Level = "Info"
	processLog(eventID, data, param, &fileLog)
}

// Error logs an error message
func Error(eventID interface{}, data map[string]interface{}, param ...interface{}) {
	var fileLog FileLog
	fileLog.Level = "Error"
	processLog(eventID, data, param, &fileLog)
}

func processLog(eventID interface{}, data map[string]interface{}, param []interface{}, fileLog *FileLog) {
	for _, v := range param {
		fileLog.Message = append(fileLog.Message, v)
		fileLog.Text += fmt.Sprint(v)
	}

	fileLog.EventID = fmt.Sprint(eventID)
	fileLog.Time = time.Now()
	if data != nil {
		fileLog.Data = data
	}

	trace(fileLog)
	fileWriter(fileLog)
}

func trace(fileLog *FileLog) {
	var packageName string

	fileLog.Trace = make(map[string]string)
	pc, file, line, _ := runtime.Caller(3)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	partsLen := len(parts)
	funcName := parts[partsLen-1]
	if len(parts[partsLen-2]) > 0 {
		if parts[partsLen-2][0] == '(' {
			funcName = parts[partsLen-2] + "." + funcName
			packageName = strings.Join(parts[0:partsLen-2], ".")
		} else {
			packageName = strings.Join(parts[0:partsLen-1], ".")
		}
	}

	parts = strings.Split(packageName, "/")
	partsLen = len(parts)

	fileLog.Trace["pack"] = parts[partsLen-1]
	fileLog.Trace["function"] = funcName
	fileLog.Trace["line"] = strconv.Itoa(line)
	fileLog.Trace["filename"] = fileName
}

// writes to a standard text file
func fileWriter(fileLog *FileLog) {
	var d, y, m, msgText, folderPath, filePath string

	currentTime := time.Now()
	d, y, m = currentTime.Format("02"), currentTime.Format("2006"), currentTime.Format("01")

	for _, v := range fileLog.Message {
		msgText += fmt.Sprint(v) + " "
	}

	folderPath = fmt.Sprintf("../logs/footballtracker/%s/%s", y, m)
	filePath = fmt.Sprintf("../logs/footballtracker/%s/%s/%s.log", y, m, d)

	// check if dir exists, otherwise create it
	// if error report it in the command line, otherwise open the file for writing
	if err := os.MkdirAll(folderPath, 0777); err != nil {
		Error(nil, nil, err)
	} else {
		if z, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			Error(nil, nil, err)
		} else {
			defer z.Close()
			log.SetOutput(z)
			log.SetFlags(log.LstdFlags | log.Lmicroseconds)

			log.Printf("%v\t[%v]\t%v\t%v\t%v\t%v\n",
				fileLog.Level,
				fileLog.Trace["pack"],
				msgText,
				fileLog.Trace["fileName"],
				fileLog.Trace["function"],
				fileLog.Trace["line"],
			)
		}
	}
}

func SetData(name string, data interface{}) map[string]interface{} {
	var d = make(map[string]interface{})
	temp, _ := json.Marshal(data)
	d[name] = string(temp)
	return d
}
