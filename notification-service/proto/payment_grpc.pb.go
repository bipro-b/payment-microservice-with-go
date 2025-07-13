package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
)

type NotifierClient interface {
	Send(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*Notification, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) Send(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/payment.Notifier/Send", in, out, opts...)
	return out, err
}

type NotifierServer interface {
	Send(context.Context, *PaymentRequest) (*Notification, error)
}
