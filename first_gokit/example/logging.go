package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type LogginMiddleware struct {
	logger log.Logger
	next OperationService
}

func (mw LogginMiddleware)Add(left int, right int) (output int) {
	defer func(begin time.Time) {
		_=mw.logger.Log(
			"method", "Add",
			"params", left, right,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Add(left, right)
	return
}