package gen

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	ProtoService *protogen.Service
	Endpoints    []*Endpoint
}

func (s *Service) Unexported(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func (s *Service) FullName() string {
	return string(s.ProtoService.Desc.FullName())
}

func (s *Service) Name() string {
	return s.ProtoService.GoName
}

func (s *Service) GonicName() string {
	return s.Name() + "Gonic"
}

func (s *Service) ServiceName() string {
	return s.GonicName() + "Service"
}

func (s *Service) AppendRouteName() string {
	return "Append" + s.GonicName() + "Route"
}

func (s *Service) HandlerName() string {
	return s.GonicName() + "Handler"
}

func (s *Service) RequestDecoderName() string {
	return s.GonicName() + "RequestDecoder"
}

func (s *Service) ResponseEncoderName() string {
	return s.GonicName() + "ResponseEncoder"
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			ProtoService: pbService,
		}
		var endpoints []*Endpoint
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("gonic: unsupport stream method, %s", endpoint.FullName())
			}
			endpoint.SetHttpRule()
			if err := checkRoute(router, endpoint); err != nil {
				return nil, fmt.Errorf("gonic: %s", err)
			}
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}

func checkRoute(router gin.IRoutes, endpoint *Endpoint) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	router.Match([]string{endpoint.Method()}, endpoint.Path(), func(ctx *gin.Context) {})
	return err
}
