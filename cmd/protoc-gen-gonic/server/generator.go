package server

import (
	"strconv"

	"github.com/go-leo/gonic/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct{}

func (generator *Generator) GenerateServices(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.ServiceName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", gen.ContextIdent, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (generator *Generator) GenerateAppendServerFunc(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("func ", service.AppendRouteName(), "[Router gin.IRoutes](router Router, service ", service.ServiceName(), ", opts ...", gen.OptionIdent, ") Router {")
	g.P("options := ", gen.NewOptionsIdent, "(opts...)")
	g.P("handler :=  ", service.Unexported(service.HandlerName()), "{")
	g.P("service: service,")
	g.P("decoder: ", service.Unexported(service.RequestDecoderName()), "{")
	g.P("unmarshalOptions: options.UnmarshalOptions(),")

	g.P("},")
	g.P("encoder: ", service.Unexported(service.ResponseEncoderName()), "{")
	g.P("marshalOptions: options.MarshalOptions(),")
	g.P("unmarshalOptions: options.UnmarshalOptions(),")
	g.P("responseTransformer: options.ResponseTransformer(),")
	g.P("},")
	g.P("errorEncoder: ", gen.DefaultEncodeErrorIdent, ",")
	g.P("shouldFailFast: options.ShouldFailFast(),")
	g.P("onValidationErrCallback: options.OnValidationErrCallback(),")
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("router.Match([]string{", strconv.Quote(endpoint.Method()), "}, ", strconv.Quote(endpoint.Path()), ", ", gen.ChainIdent, "(", "handler.", endpoint.Name(), "(), options.Middlewares()...)...)")
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
	g.P("onValidationErrCallback ", gen.OnValidationErrCallbackIdent)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (h ", service.Unexported(service.HandlerName()), ")", endpoint.Name(), "()", gen.HandlerFuncIdent, " {")
		g.P("return ", gen.HandlerFuncIdent, "(func(ctx *", gen.GinContextIdent, ") {")
		g.P("in, err := h.decoder.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("h.errorEncoder(ctx, err, ctx.Writer)")
		g.P("return")
		g.P("}")
		g.P("if err := ", gen.ValidateRequestIdent, "(ctx, in, h.shouldFailFast, h.onValidationErrCallback)", "; err != nil {")
		g.P("h.errorEncoder(ctx, err, ctx.Writer)")
		g.P("return")
		g.P("}")
		g.P("out, err := h.service.", endpoint.Name(), "(ctx, in)")
		g.P("if err != nil {")
		g.P("h.errorEncoder(ctx, err, ctx.Writer)")
		g.P("return")
		g.P("}")
		g.P("if err := h.encoder.", endpoint.Name(), "(ctx, out); err != nil {")
		g.P("h.errorEncoder(ctx, err, ctx.Writer)")
		g.P("return")
		g.P("}")
		g.P("})")
		g.P("}")
		g.P()
	}
	return nil
}
