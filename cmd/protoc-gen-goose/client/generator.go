package client

import (
	"github.com/go-leo/goose/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct{}

func (f *Generator) GenerateNewClient(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("func ", service.NewClientName(), "(target string, opts ...", gen.ClientOptionIdent, ") ", service.ServiceName(), " {")
	g.P("options := ", gen.ClientNewOptionsIdent, "(opts...)")
	g.P("client :=  &", service.Unexported(service.ClientName()), "{")
	g.P("client: options.Client(),")
	g.P("encoder: ", service.Unexported(service.RequestEncoderName()), "{")
	g.P("target: target,")
	g.P("marshalOptions: options.MarshalOptions(),")
	g.P("},")
	g.P("decoder: ", service.Unexported(service.ResponseDecoderName()), "{")
	g.P("unmarshalOptions: options.UnmarshalOptions(),")
	g.P("errorDecoder: options.ErrorDecoder(),")
	g.P("errorFactory: options.ErrorFactory(),")
	g.P("},")
	g.P("shouldFailFast: options.ShouldFailFast(),")
	g.P("onValidationErrCallback: options.OnValidationErrCallback(),")
	g.P("middleware: ", gen.ClientChainIdent, "(options.Middlewares()...),")
	g.P("}")
	g.P("return client")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClient(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.ClientName()), " struct {")
	g.P("client *", gen.ClientIdent)
	g.P("encoder ", service.Unexported(service.RequestEncoderName()))
	g.P("decoder ", service.Unexported(service.ResponseDecoderName()))
	g.P("shouldFailFast bool")
	g.P("onValidationErrCallback ", gen.OnErrCallbackIdent)
	g.P("middleware ", gen.ClientMiddlewareIdent)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.Unexported(service.ClientName()), ") ", endpoint.Name(), "(ctx ", gen.ContextIdent, ", req *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("if err := ", gen.ValidateRequestIdent, "(ctx, req, c.shouldFailFast, c.onValidationErrCallback); err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("request, err := c.encoder.", endpoint.Name(), "(ctx, req)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("response, err := ", gen.ClientInvokeIdent, "(ctx, c.middleware, c.client, request)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("resp, err := c.decoder.", endpoint.Name(), "(ctx, response)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return resp, nil")
		g.P("}")
		g.P()
	}
	return nil
}
