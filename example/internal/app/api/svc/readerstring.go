package svc

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-leo/gors"
)

type ReaderStringService struct {
}

func (svc *ReaderStringService) GetReaderString(ctx context.Context, reader io.Reader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", gors.NewHttpError(http.StatusBadRequest, err)
	}
	return "hello " + string(data), nil
}

func (svc *ReaderStringService) PostReaderString(ctx context.Context, reader io.Reader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", gors.NewHttpError(http.StatusBadRequest, err)
	}
	return "hello " + string(data), nil
}
