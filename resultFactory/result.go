package resultFactory

import (
	"time"

	"github.com/supnice/common-utils/common/logger"
)

type ResultBean struct {
	Timestamp int64       `json:"timestamp"`
	Status    int         `json:"status"`
	Exception interface{} `json:"exception`
	Error     interface{} `json:"error"`
	Message   interface{} `json:"message"`
	Data      interface{} `json:"data"`
}

func Success(data interface{}) ResultBean {
	timeStamp := time.Now().UnixNano() / 1e6
	result := ResultBean{
		Timestamp: timeStamp,
		Status:    0,
		Data:      data,
	}

	logger.Info("result:%v", result)
	return result
}

func Error(args ...interface{}) ResultBean {
	var err interface{}
	var msg interface{}
	var status int

	switch len(args) {
	case 0:
		status = -1
		err = ""
		msg = ""
	case 1:
		status = -1
		err = args[0]
		msg = args[0]
	case 2:
		status = -1
		err = args[0]
		msg = args[1]
	case 3:
		defer func() {
			if err := recover(); err != nil {
				logger.Error("status参数类型错误, %v", err)
			}
		}()
		status = args[0].(int)
		err = args[1]
		msg = args[2]
	default:
		logger.Warn("传入参数错误")
	}
	// logger.Info("type:%v", reflect.TypeOf(args[0]))

	timeStamp := time.Now().UnixNano() / 1e6
	result := ResultBean{
		Timestamp: timeStamp,
		Status:    status,
		Error:     err,
		Message:   msg,
	}

	logger.Info("result:%v", result)
	return result
}
