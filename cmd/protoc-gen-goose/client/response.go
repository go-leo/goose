package client

import (
	"fmt"

	"github.com/go-leo/goose/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (f *Generator) GenerateResponseDecoder(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.ResponseDecoderName()), " struct {")
	g.P("unmarshalOptions ", gen.ProtoJsonUnmarshalOptionsIdent)
	g.P("errorDecoder ", gen.ErrorDecoderIdent)
	g.P("errorFactory ", gen.ErrorFactoryIdent)
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (decoder *", service.Unexported(service.ResponseDecoderName()), ") ", endpoint.Name(), "(ctx ", gen.ContextIdent, ", response *", gen.ResponseIdent, ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("if respErr, ok := decoder.errorDecoder(ctx, response, decoder.errorFactory); ok {")
		g.P("return nil, respErr")
		g.P("}")
		g.P("resp := &", endpoint.Output().GoIdent, "{}")
		bodyParameter := endpoint.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				f.DecodeHttpBody(g, srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				f.DecodeHttpResponse(g, srcValue)
			default:
				srcValue := []any{"resp"}
				f.PrintDecodeMessage(g, srcValue)
			}
		default:
			bodyField := gen.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.", bodyField.GoName}
			g.P(append(append([]any{"if "}, srcValue...), " == nil {")...)
			g.P(append(srcValue, " = &", bodyField.Message.GoIdent, "{}")...)
			g.P("}")
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.DecodeHttpBody(g, srcValue)
				default:
					f.PrintDecodeMessage(g, srcValue)
				}
			}
		}
		g.P("return resp, nil")
		g.P("}")
		g.P()
	}
	g.P()
	return nil
}

func (f *Generator) PrintDecodeMessage(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"if err := ", gen.DecodeMessageIdent, "(ctx, response, "}, srcValue...), ", decoder.unmarshalOptions); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *Generator) DecodeHttpBody(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"if err := ", gen.DecodeHttpBodyFromResponseIdent, "(ctx, response, "}, srcValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *Generator) DecodeHttpResponse(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"if err := ", gen.DecodeHttpResponseIdent, "(ctx, response, "}, srcValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}
