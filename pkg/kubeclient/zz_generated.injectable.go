/*
Package kubeclient GENERATED BY gengo:injectable 
DON'T EDIT THIS FILE
*/
package kubeclient

import (
	context "context"
)

type contextClient struct{}

func ClientFromContext(ctx context.Context) (Client, bool) {
	if v, ok := ctx.Value(contextClient{}).(Client); ok {
		return v, true
	}
	return nil, false
}

func ClientInjectContext(ctx context.Context, tpe Client) context.Context {
	return context.WithValue(ctx, contextClient{}, tpe)
}
func (p *KubeClient) InjectContext(ctx context.Context) context.Context {
	ctx = ClientInjectContext(ctx, p.c)

	return ctx
}
func (v *KubeClient) Init(ctx context.Context) error {
	if err := v.beforeInit(ctx); err != nil {
		return err
	}

	return nil
}
