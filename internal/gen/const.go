package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	ProtoJsonPackage               = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoJsonMarshalOptionsIdent   = ProtoJsonPackage.Ident("MarshalOptions")
	ProtoJsonUnmarshalOptionsIdent = ProtoJsonPackage.Ident("UnmarshalOptions")
)

var (
	ContextPackage = protogen.GoImportPath("context")
	ContextIdent   = ContextPackage.Ident("Context")
)

var (
	HttpPackage          = protogen.GoImportPath("net/http")
	HttpHandlerIdent     = HttpPackage.Ident("Handler")
	HttpHandlerFuncIdent = HttpPackage.Ident("HandlerFunc")
	ResponseWriterIdent  = HttpPackage.Ident("ResponseWriter")
	RequestIdent         = HttpPackage.Ident("Request")
	RouterIdent          = HttpPackage.Ident("ServeMux")
)

var (
	FmtPackage   = protogen.GoImportPath("fmt")
	SprintfIdent = FmtPackage.Ident("Sprintf")
)

var (
	ProtoPackage     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoStringIdent = ProtoPackage.Ident("String")
)

var (
	GooseServerPackage           = protogen.GoImportPath("github.com/go-leo/goose/server")
	ErrorEncoderIdent            = GooseServerPackage.Ident("ErrorEncoder")
	ResponseTransformerIdent     = GooseServerPackage.Ident("ResponseTransformer")
	DefaultEncodeErrorIdent      = GooseServerPackage.Ident("DefaultEncodeError")
	EncodeResponseIdent          = GooseServerPackage.Ident("EncodeResponse")
	EncodeHttpBodyIdent          = GooseServerPackage.Ident("EncodeHttpBody")
	EncodeHttpResponseIdent      = GooseServerPackage.Ident("EncodeHttpResponse")
	DecodeRequestIdent           = GooseServerPackage.Ident("DecodeRequest")
	DecodeHttpBodyIdent          = GooseServerPackage.Ident("DecodeHttpBody")
	DecodeHttpRequestIdent       = GooseServerPackage.Ident("DecodeHttpRequest")
	DecodeFormIdent              = GooseServerPackage.Ident("DecodeForm")
	OptionIdent                  = GooseServerPackage.Ident("Option")
	NewOptionsIdent              = GooseServerPackage.Ident("NewOptions")
	ChainIdent                   = GooseServerPackage.Ident("Chain")
	CustomDecodeRequestIdent     = GooseServerPackage.Ident("CustomDecodeRequest")
	OnValidationErrCallbackIdent = GooseServerPackage.Ident("OnValidationErrCallback")
	ValidateRequestIdent         = GooseServerPackage.Ident("ValidateRequest")
)

var (
	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)
