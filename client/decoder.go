package client

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/go-leo/goose"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func DecodeMessage(ctx context.Context, response *http.Response, resp proto.Message, unmarshalOptions protojson.UnmarshalOptions) error {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Join(err, response.Body.Close())
	}
	if err := unmarshalOptions.Unmarshal(data, resp); err != nil {
		return errors.Join(err, response.Body.Close())
	}
	return response.Body.Close()
}

func DecodeHttpBody(ctx context.Context, response *http.Response, resp *httpbody.HttpBody) error {
	resp.ContentType = response.Header.Get(goose.ContentTypeKey)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	resp.Data = body
	return response.Body.Close()
}

func DecodeHttpResponse(ctx context.Context, response *http.Response, resp *rpchttp.HttpResponse) error {
	resp.Status = int32(response.StatusCode)
	resp.Reason = http.StatusText(response.StatusCode)
	resp.Headers = make([]*rpchttp.HttpHeader, 0, len(response.Header))
	for key, values := range response.Header {
		for _, value := range values {
			elems := &rpchttp.HttpHeader{
				Key:   key,
				Value: value,
			}
			resp.Headers = append(resp.Headers, elems)
		}
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Join(err, response.Body.Close())
	}
	resp.Body = data
	return response.Body.Close()
}
