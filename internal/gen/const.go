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
	GoosePackage                 = protogen.GoImportPath("github.com/go-leo/goose")
	ErrorEncoderIdent            = GoosePackage.Ident("ErrorEncoder")
	ResponseTransformerIdent     = GoosePackage.Ident("ResponseTransformer")
	DefaultEncodeErrorIdent      = GoosePackage.Ident("DefaultEncodeError")
	EncodeResponseIdent          = GoosePackage.Ident("EncodeResponse")
	EncodeHttpBodyIdent          = GoosePackage.Ident("EncodeHttpBody")
	EncodeHttpResponseIdent      = GoosePackage.Ident("EncodeHttpResponse")
	DecodeRequestIdent           = GoosePackage.Ident("DecodeRequest")
	DecodeHttpBodyIdent          = GoosePackage.Ident("DecodeHttpBody")
	DecodeHttpRequestIdent       = GoosePackage.Ident("DecodeHttpRequest")
	DecodeFormIdent              = GoosePackage.Ident("DecodeForm")
	OptionIdent                  = GoosePackage.Ident("Option")
	NewOptionsIdent              = GoosePackage.Ident("NewOptions")
	ChainIdent                   = GoosePackage.Ident("Chain")
	CustomDecodeRequestIdent     = GoosePackage.Ident("CustomDecodeRequest")
	OnValidationErrCallbackIdent = GoosePackage.Ident("OnValidationErrCallback")
	ValidateRequestIdent         = GoosePackage.Ident("ValidateRequest")
)

var (
	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)
