package server

import (
	"fmt"

	"github.com/go-leo/gonic/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (generator *Generator) GenerateEncodeResponse(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.ResponseEncoderName()), " struct {")
	g.P("marshalOptions ", gen.ProtoJsonMarshalOptionsIdent)
	g.P("unmarshalOptions ", gen.ProtoJsonUnmarshalOptionsIdent)
	g.P("responseTransformer ", gen.ResponseTransformerIdent)
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (encoder ", service.Unexported(service.ResponseEncoderName()), ")", endpoint.Name(), "(ctx *", gen.GinContextIdent, ", resp *", endpoint.OutputGoIdent(), ") error {")
		bodyParameter := endpoint.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				generator.PrintHttpBodyEncodeBlock(g, srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				generator.PrintHttpResponseEncodeBlock(g, srcValue)
			default:
				srcValue := []any{"encoder.responseTransformer(ctx, resp)"}
				generator.PrintResponseEncodeBlock(g, srcValue)
			}
		default:
			bodyField := gen.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					srcValue := []any{"resp.Get", bodyField.GoName, "()"}
					generator.PrintHttpBodyEncodeBlock(g, srcValue)
				default:
					srcValue := []any{"encoder.responseTransformer(ctx, resp.Get", bodyField.GoName, "())"}
					generator.PrintResponseEncodeBlock(g, srcValue)
				}
			}
		}
		g.P("}")
	}
	g.P()
	return nil
}

func (generator *Generator) PrintHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", gen.EncodeHttpBodyIdent, "(ctx, ctx.Writer, "}, srcValue...), ")")...)
}

func (generator *Generator) PrintHttpResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", gen.EncodeHttpResponseIdent, "(ctx, ctx.Writer, "}, srcValue...), ")")...)
}

func (generator *Generator) PrintResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", gen.EncodeResponseIdent, "(ctx, ctx.Writer, "}, srcValue...), ", encoder.marshalOptions)")...)
}
