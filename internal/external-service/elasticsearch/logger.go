// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import "context"

type Logger interface {
	Info(ctx context.Context, message string)
	Trace(ctx context.Context, message string)
	Error(ctx context.Context, message string)
}
