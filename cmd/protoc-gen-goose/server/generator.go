package server

import (
	"strconv"

	"github.com/go-leo/goose/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct{}

func (generator *Generator) GenerateServices(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.ServiceName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", gen.ContextIdent, ", req *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (generator *Generator) GenerateAppendServerFunc(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("func ", service.AppendRouteName(), "(router *", gen.RouterIdent, ", service ", service.ServiceName(), ", opts ...", gen.ServerOptionIdent, ") ", "*", gen.RouterIdent, " {")
	g.P("options := ", gen.ServerNewOptionsIdent, "(opts...)")
	g.P("handler :=  ", service.Unexported(service.HandlerName()), "{")
	g.P("service: service,")
	g.P("decoder: ", service.Unexported(service.RequestDecoderName()), "{")
	g.P("unmarshalOptions: options.UnmarshalOptions(),")
	g.P("},")
	g.P("encoder: ", service.Unexported(service.ResponseEncoderName()), "{")
	g.P("marshalOptions: options.MarshalOptions(),")
	g.P("unmarshalOptions: options.UnmarshalOptions(),")
	g.P("},")
	g.P("errorEncoder: options.ErrorEncoder(),")
	g.P("shouldFailFast: options.ShouldFailFast(),")
	g.P("onValidationErrCallback: options.OnValidationErrCallback(),")
	g.P("middleware: ", gen.ServerChainIdent, "(options.Middlewares()...),")
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("router.Handle(", strconv.Quote(endpoint.Method()+" "+endpoint.Path()), ", ", gen.HttpHandlerFuncIdent, "(handler.", endpoint.Name(), "))")
	}
	g.P("return router")
	g.P("}")
	g.P()
	return nil
}

func (generator *Generator) GenerateHandlers(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HandlerName()), " struct {")
	g.P("service ", service.ServiceName())
	g.P("decoder ", service.Unexported(service.RequestDecoderName()))
	g.P("encoder ", service.Unexported(service.ResponseEncoderName()))
	g.P("errorEncoder ", gen.ErrorEncoderIdent)
	g.P("shouldFailFast bool")
	g.P("onValidationErrCallback ", gen.OnErrCallbackIdent)
	g.P("middleware ", gen.ServerMiddlewareIdent)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (h ", service.Unexported(service.HandlerName()), ")", endpoint.Name(), "(response ", gen.ResponseWriterIdent, ", request *", gen.RequestIdent, ") {")
		g.P("invoke := func(response ", gen.ResponseWriterIdent, ", request *", gen.RequestIdent, ") {")
		g.P("ctx := request.Context()")
		g.P("req, err := h.decoder.", endpoint.Name(), "(ctx, request)")
		g.P("if err != nil {")
		g.P("h.errorEncoder(ctx, err, response)")
		g.P("return")
		g.P("}")
		g.P("if err := ", gen.ValidateRequestIdent, "(ctx, req, h.shouldFailFast, h.onValidationErrCallback)", "; err != nil {")
		g.P("h.errorEncoder(ctx, err, response)")
		g.P("return")
		g.P("}")
		g.P("resp, err := h.service.", endpoint.Name(), "(ctx, req)")
		g.P("if err != nil {")
		g.P("h.errorEncoder(ctx, err, response)")
		g.P("return")
		g.P("}")
		g.P("if err := h.encoder.", endpoint.Name(), "(ctx, response, resp); err != nil {")
		g.P("h.errorEncoder(ctx, err, response)")
		g.P("return")
		g.P("}")
		g.P("}")
		g.P(gen.ServerInvokeIdent, "(h.middleware, response, request, invoke)")
		g.P("}")
		g.P()
	}
	return nil
}
