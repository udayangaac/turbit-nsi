// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"fmt"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

type errorLogger struct{}
type infoLogger struct{}
type traceLogger struct{}

func NewErrorLogger() elastic.Logger {
	return &errorLogger{}
}
func NewInfoLogger() elastic.Logger {
	return &infoLogger{}
}
func NewTraceLogger() elastic.Logger {
	return &traceLogger{}
}

func (e errorLogger) Printf(format string, v ...interface{}) {
	log.Error(fmt.Sprintf(format, v...))
}

func (e traceLogger) Printf(format string, v ...interface{}) {
	log.Trace(fmt.Sprintf(format, v...))
}

func (e infoLogger) Printf(format string, v ...interface{}) {
	log.Info(fmt.Sprintf(format, v...))
}
