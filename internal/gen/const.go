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
)

var (
	FmtPackage   = protogen.GoImportPath("fmt")
	SprintfIdent = FmtPackage.Ident("Sprintf")
)

var (
	SlicesPackage = protogen.GoImportPath("slices")
	CloneIdent    = SlicesPackage.Ident("Clone")
)

var (
	GinPackage  = protogen.GoImportPath("github.com/gin-gonic/gin")
	RouterIdent = GinPackage.Ident("IRoutes")
	WrapHIdent  = GinPackage.Ident("WrapH")
	HandlerFuncIdent  = GinPackage.Ident("HandlerFunc")
	GinContextIdent  = GinPackage.Ident("Context")
)

var (
	ProtoPackage     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoStringIdent = ProtoPackage.Ident("String")
)

var (
	GonicPackage                 = protogen.GoImportPath("github.com/go-leo/gonic")
	ErrorEncoderIdent            = GonicPackage.Ident("ErrorEncoder")
	ResponseTransformerIdent     = GonicPackage.Ident("ResponseTransformer")
	DefaultEncodeErrorIdent      = GonicPackage.Ident("DefaultEncodeError")
	EncodeResponseIdent          = GonicPackage.Ident("EncodeResponse")
	EncodeHttpBodyIdent          = GonicPackage.Ident("EncodeHttpBody")
	EncodeHttpResponseIdent      = GonicPackage.Ident("EncodeHttpResponse")
	DecodeRequestIdent           = GonicPackage.Ident("DecodeRequest")
	DecodeHttpBodyIdent          = GonicPackage.Ident("DecodeHttpBody")
	DecodeHttpRequestIdent       = GonicPackage.Ident("DecodeHttpRequest")
	DecodeFormIdent              = GonicPackage.Ident("DecodeForm")
	OptionIdent                  = GonicPackage.Ident("Option")
	NewOptionsIdent              = GonicPackage.Ident("NewOptions")
	ChainIdent                   = GonicPackage.Ident("Chain")
	CustomDecodeRequestIdent     = GonicPackage.Ident("CustomDecodeRequest")
	OnValidationErrCallbackIdent = GonicPackage.Ident("OnValidationErrCallback")
	ValidateRequestIdent         = GonicPackage.Ident("ValidateRequest")
	VarsIdent                    = GonicPackage.Ident("Vars")
)

var (
	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)
