// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"context"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
)

type customLogger struct{}

func NewDefaultLogger() Logger {
	return new(customLogger)
}

func (c *customLogger) Error(ctx context.Context, message string) {
	log.Error(log_traceable.GetMessage(ctx, message))
}

func (c *customLogger) Info(ctx context.Context, message string) {
	log.Error(log_traceable.GetMessage(ctx, message))
}

func (c *customLogger) Trace(ctx context.Context, message string) {
	log.Error(log_traceable.GetMessage(ctx, message))
}
