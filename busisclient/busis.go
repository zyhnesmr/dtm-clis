// Code generated by goctl. DO NOT EDIT.
// Source: busis.proto

package busisclient

import (
	"context"

	"dtm-clis/busis"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReq = busis.AddReq
	DelReq = busis.DelReq
	Empty  = busis.Empty

	Busis interface {
		AddMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error)
		DelMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultBusis struct {
		cli zrpc.Client
	}
)

func NewBusis(cli zrpc.Client) Busis {
	return &defaultBusis{
		cli: cli,
	}
}

func (m *defaultBusis) AddMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error) {
	client := busis.NewBusisClient(m.cli.Conn())
	return client.AddMoney(ctx, in, opts...)
}

func (m *defaultBusis) DelMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error) {
	client := busis.NewBusisClient(m.cli.Conn())
	return client.DelMoney(ctx, in, opts...)
}
