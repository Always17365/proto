// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"context"

	"github.com/teamgram/proto/v2/rpc/examples/echo/internal/svc"
)

type EchoCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *EchoCore {
	return &EchoCore{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
