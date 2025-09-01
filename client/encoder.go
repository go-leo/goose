package client

import (
	"context"
	"io"
	"net/http"

	"github.com/go-leo/goose"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func EncodeMessage(ctx context.Context, req proto.Message, header http.Header, body io.Writer, marshalOptions protojson.MarshalOptions) error {
	data, err := marshalOptions.Marshal(req)
	if err != nil {
		return err
	}
	if _, err = body.Write(data); err != nil {
		return err
	}
	header.Set(goose.ContentTypeKey, goose.JsonContentType)
	return nil
}

func EncodeHttpBody(ctx context.Context, req *httpbody.HttpBody, header http.Header, body io.Writer) error {
	if _, err := body.Write(req.GetData()); err != nil {
		return err
	}
	header.Set(goose.ContentTypeKey, req.GetContentType())
	return nil
}

func EncodeHttpRequest(ctx context.Context, req *rpchttp.HttpRequest, header http.Header, body io.Writer) error {
	if _, err := body.Write(req.GetBody()); err != nil {
		return err
	}
	for _, httpHeader := range req.GetHeaders() {
		header.Add(httpHeader.GetKey(), httpHeader.GetValue())
	}
	return nil
}
