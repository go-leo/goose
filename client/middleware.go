package client

import (
	"context"
	"net/http"
)

type Invoker func(ctx context.Context, cli *http.Client, request *http.Request) (*http.Response, error)

type Middleware func(ctx context.Context, cli *http.Client, request *http.Request, invoker Invoker) (*http.Response, error)

func Chain(middlewares ...Middleware) Middleware {
	var mdw Middleware
	if len(middlewares) == 0 {
		mdw = nil
	} else if len(middlewares) == 1 {
		mdw = middlewares[0]
	} else {
		mdw = func(ctx context.Context, cli *http.Client, request *http.Request, invoker Invoker) (*http.Response, error) {
			return middlewares[0](ctx, cli, request, getInvoker(middlewares, 0, invoker))
		}
	}
	return mdw
}

func getInvoker(interceptors []Middleware, curr int, finalInvoker Invoker) Invoker {
	if curr == len(interceptors)-1 {
		return finalInvoker
	}
	return func(ctx context.Context, cli *http.Client, request *http.Request) (*http.Response, error) {
		return interceptors[curr+1](ctx, cli, request, getInvoker(interceptors, curr+1, finalInvoker))
	}
}

func Invoke(ctx context.Context, middleware Middleware, cli *http.Client, request *http.Request) (*http.Response, error) {
	invoke := func(ctx context.Context, cli *http.Client, request *http.Request) (*http.Response, error) {
		return cli.Do(request)
	}
	if middleware == nil {
		return invoke(ctx, cli, request)
	}
	return middleware(ctx, cli, request, invoke)
}
