package gen

import (
	"fmt"
	"strings"

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

func (s *Service) GooseName() string {
	return s.Name() + "Goose"
}

func (s *Service) ServiceName() string {
	return s.GooseName() + "Service"
}

func (s *Service) AppendRouteName() string {
	return "Append" + s.GooseName() + "Route"
}

func (s *Service) HandlerName() string {
	return s.GooseName() + "Handler"
}

func (s *Service) RequestDecoderName() string {
	return s.GooseName() + "RequestDecoder"
}

func (s *Service) ResponseEncoderName() string {
	return s.GooseName() + "ResponseEncoder"
}

func (s *Service) NewClientName() string {
	return "New" + s.GooseName() + "Client"
}

func (s *Service) ClientName() string {
	return s.GooseName() + "Client"
}

func (s *Service) RequestEncoderName() string {
	return s.GooseName() + "RequestEncoder"
}

func (s *Service) ResponseDecoderName() string {
	return s.GooseName() + "ResponseDecoder"
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			ProtoService: pbService,
		}
		var endpoints []*Endpoint
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("goose: unsupport stream method, %s", endpoint.FullName())
			}
			endpoint.SetHttpRule()
			pattern, err := ParsePattern(endpoint.Path())
			if err != nil {
				return nil, fmt.Errorf("goose: %s", err)
			}
			endpoint.SetPattern(pattern)
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}
