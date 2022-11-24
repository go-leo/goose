package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"io"
	"log"
	"strconv"

	"github.com/go-leo/stringx"
)

const (
	contextPackage = goImportPath("context")
	gorsPackage    = goImportPath("github.com/go-leo/gors")
	ginPackage     = goImportPath("github.com/gin-gonic/gin")
	httpPackage    = goImportPath("net/http")
	ioPackage      = goImportPath("io")
	bindingPackage = goImportPath("github.com/gin-gonic/gin/binding")
	renderPackage  = goImportPath("github.com/gin-gonic/gin/render")
)

type generate struct {
	buf              *bytes.Buffer
	headerBuf        *bytes.Buffer
	importsBuf       *bytes.Buffer
	functionBuf      *bytes.Buffer
	header           string
	pkg              string
	imports          map[string]*goImport
	srvName          string
	usedPackageNames map[string]bool
	routerInfos      []*routerInfo
}

func (g *generate) checkResult2MustBeError(rpcType *ast.FuncType, methodName *ast.Ident) {
	result2 := rpcType.Results.List[1]
	result2Iden, ok := result2.Type.(*ast.Ident)
	if !ok {
		log.Fatalf("error: func %s 2th result is not error", methodName)
	}
	if result2Iden.Name != "error" {
		log.Fatalf("error: func %s 2th result is not error", methodName)
	}
}

func (g *generate) checkAndGetResult1(rpcType *ast.FuncType, methodName *ast.Ident) *result {
	result1 := rpcType.Results.List[0]
	switch r1 := result1.Type.(type) {
	case *ast.ArrayType:
		ident, ok := r1.Elt.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ident.Name != "byte" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &result{bytes: true}
	case *ast.Ident:
		if r1.Name != "string" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &result{string: true}
	case *ast.StarExpr:
		switch x := r1.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &result{objectArgs: &objectArgs{name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &result{objectArgs: &objectArgs{name: x.Sel.Name, goImportPath: goImportPath(importPath)}}
				}
			}
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		default:
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		}
	case *ast.SelectorExpr:
		if r1.Sel == nil {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if r1.Sel.Name != "Reader" {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ident, ok := r1.X.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ioImport, ok := g.imports["io"]
		if !ok {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ioImport.PackageName != ident.Name {
			log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &result{reader: true}
	default:

	}
	return nil
}

func (g *generate) checkResults(rpcType *ast.FuncType, methodName *ast.Ident) {
	if rpcType.Results == nil {
		log.Fatalf("error: func %s results is empty", methodName)
	}
	if len(rpcType.Results.List) != 2 {
		log.Fatalf("error: func %s results count is not equal 2", methodName)
	}
}

func (g *generate) checkAndGetParam2(rpcType *ast.FuncType, methodName *ast.Ident) *param {
	param2 := rpcType.Params.List[1]
	switch p2 := param2.Type.(type) {
	case *ast.ArrayType:
		ident, ok := p2.Elt.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ident.Name != "byte" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &param{bytes: true}
	case *ast.Ident:
		if p2.Name != "string" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &param{string: true}
	case *ast.StarExpr:
		switch x := p2.X.(type) {
		case *ast.Ident:
			name := x.Name
			return &param{objectArgs: &objectArgs{name: name}}
		case *ast.SelectorExpr:
			ident, ok := x.X.(*ast.Ident)
			if !ok {
				log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			}
			for importPath, goImport := range g.imports {
				if goImport.PackageName == ident.Name {
					return &param{objectArgs: &objectArgs{name: x.Sel.Name, goImportPath: goImportPath(importPath)}}
				}
			}
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		default:
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
			return nil
		}

	case *ast.SelectorExpr:
		if p2.Sel == nil {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if p2.Sel.Name != "Reader" {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ident, ok := p2.X.(*ast.Ident)
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		ioImport, ok := g.imports["io"]
		if !ok {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		if ioImport.PackageName != ident.Name {
			log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		}
		return &param{reader: true}
	default:
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or io.Reader or *struct{}", methodName)
		return nil
	}
}

func (g *generate) checkParam1MustBeContext(rpcType *ast.FuncType, methodName *ast.Ident) {
	param1 := rpcType.Params.List[0]
	param0SelectorExpr, ok := param1.Type.(*ast.SelectorExpr)
	if !ok {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExpr.Sel.Name != "Context" {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	param0SelectorExprX, ok := param0SelectorExpr.X.(*ast.Ident)
	if !ok {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
	if param0SelectorExprX.Name != "context" {
		log.Fatalf("error: func %s 1th param is not context.Context", methodName)
	}
}

func (g *generate) checkParams(rpcType *ast.FuncType, methodName *ast.Ident) {
	if rpcType.Params == nil {
		log.Fatalf("error: func %s params is empty", methodName)
	}
	if len(rpcType.Params.List) != 2 {
		log.Fatalf("error: func %s params count is not equal 2", methodName)
	}
}

func (g *generate) content() []byte {
	// header
	g.printHeader()

	// function
	g.printFunction()

	// imports
	g.printImports()

	// all
	g.combine()
	return g.buf.Bytes()
}

func (g *generate) printHeader() {
	g.P(g.headerBuf, g.header)
	g.P(g.headerBuf, g.pkg)
}

func (g *generate) printFunction() {
	serviceName := g.srvName
	functionName := serviceName + "Routers"
	g.P(g.functionBuf, "func ", functionName, "(srv ", serviceName, ") []", gorsPackage.Ident("Router"), " {")
	g.P(g.functionBuf, "return []", gorsPackage.Ident("Router"), "{")
	for _, routerInfo := range g.routerInfos {
		g.printRouterInfo(routerInfo)
	}
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "}")
}

func (g *generate) printRouterInfo(info *routerInfo) {
	g.P(g.functionBuf, "{")
	g.P(g.functionBuf, "HTTPMethod: ", httpPackage.Ident(info.method), ",")
	g.P(g.functionBuf, "Path: ", strconv.Quote(info.path), ",")
	g.P(g.functionBuf, "HandlerFunc: func(c *", ginPackage.Ident("Context"), ") {")
	g.printHandler(info)
	g.P(g.functionBuf, "},")
	g.P(g.functionBuf, "},")
}

func (g *generate) printHandler(info *routerInfo) {
	if info.param2.bytes {
		g.printBytesReq(info)
	} else if info.param2.string {
		g.printStringReq(info)
	} else if info.param2.reader {
		g.printReaderReq(info)
	} else if info.param2.objectArgs != nil {
		g.printObjectReq(info)
	} else {
		log.Fatalf("error: func %s 2th param is invalid, must be []byte or string or *struct{}", info.rpcMethodName)
	}

	g.printRPCHandler(info)

	if info.result1.bytes {
		g.printBytesRender(info)
	} else if info.result1.string {
		g.printStringRender(info)
	} else if info.result1.reader {
		g.printReaderRender(info)
	} else if info.result1.objectArgs != nil {
		g.printObjectRender(info)
	} else {
		log.Fatalf("error: func %s 1th result is invalid, must be []byte or string or *struct{}", info.rpcMethodName)
	}
}

func (g *generate) printBytesReq(info *routerInfo) {
	g.P(g.functionBuf, "body, err := ", ioPackage.Ident("ReadAll"), "(c.Request.Body)")
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusBadRequest"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(", ginPackage.Ident("ErrorTypeBind"), ")")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "req := body")
}

func (g *generate) printStringReq(info *routerInfo) {
	g.P(g.functionBuf, "body, err := ", ioPackage.Ident("ReadAll"), "(c.Request.Body)")
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusBadRequest"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(gin.ErrorTypeBind)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "req := string(body)")
}

func (g *generate) printReaderReq(info *routerInfo) {
	g.P(g.functionBuf, "body := c.Request.Body")
	g.P(g.functionBuf, "req := body")
}

func (g *generate) printObjectReqInit(info *routerInfo) {
	objArgs := info.param2.objectArgs
	if stringx.IsBlank(string(objArgs.goImportPath)) {
		g.P(g.functionBuf, "req := new(", objArgs.name, ")")
	} else {
		g.P(g.functionBuf, "req := new(", objArgs.goImportPath.Ident(objArgs.name), ")")
	}
}

func (g *generate) printBindUriRequest() {
	g.P(g.functionBuf, "if err := c.BindUri(req); err != nil {")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusBadRequest"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(gin.ErrorTypeBind)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printBindRequest(binding string) {
	g.P(g.functionBuf, "if err := c.ShouldBindWith(req, ", bindingPackage.Ident(binding), "); err != nil {")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusBadRequest"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(gin.ErrorTypeBind)")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
}

func (g *generate) printObjectReq(info *routerInfo) {
	g.printObjectReqInit(info)
	if info.uriBinding {
		g.printBindUriRequest()
	}
	if info.queryBinding {
		g.printBindRequest("Query")
	}
	if info.headerBinding {
		g.printBindRequest("Header")
	}
	if info.jsonBinding {
		g.printBindRequest("JSON")
	}
	if info.xmlBinding {
		g.printBindRequest("XML")
	}
	if info.formBinding {
		g.printBindRequest("Form")
	}
	if info.formPostBinding {
		g.printBindRequest("FormPost")
	}
	if info.formMultipartBinding {
		g.printBindRequest("FormMultipart")
	}
	if info.protobufBinding {
		g.printBindRequest("ProtoBuf")
	}
	if info.msgpackBinding {
		g.printBindRequest("MsgPack")
	}
	if info.yamlBinding {
		g.printBindRequest("YAML")
	}
	if info.tomlBinding {
		g.printBindRequest("TOML")
	}

}

func (g *generate) printRPCHandler(info *routerInfo) {
	g.P(g.functionBuf, "var ctx ", contextPackage.Ident("Context"), " = c")
	g.P(g.functionBuf, "ctx = ", gorsPackage.Ident("InjectStatusCode"), "(ctx, ", httpPackage.Ident("StatusOK"), ")")
	g.P(g.functionBuf, "ctx = ", gorsPackage.Ident("InjectHeader"), "(ctx, c.Writer.Header())")
	g.P(g.functionBuf, "resp, err := srv.", info.rpcMethodName, "(ctx, req)")
	g.P(g.functionBuf, "if err != nil {")
	g.P(g.functionBuf, "if httpErr, ok := err.(*", gorsPackage.Ident("HttpError"), "); ok {")
	g.P(g.functionBuf, "c.String(httpErr.StatusCode(), httpErr.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(", ginPackage.Ident("ErrorTypePublic"), ")")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "c.String(", httpPackage.Ident("StatusInternalServerError"), ", err.Error())")
	g.P(g.functionBuf, "_ = c.Error(err).SetType(", ginPackage.Ident("ErrorTypePrivate"), ")")
	g.P(g.functionBuf, "return")
	g.P(g.functionBuf, "}")
	g.P(g.functionBuf, "statusCode := ", gorsPackage.Ident("ExtractStatusCode"), "(ctx)")
}

func (g *generate) printBytesRender(info *routerInfo) {
	switch {
	case info.bytesRender:
		g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Data"), "{ContentType: ", strconv.Quote(info.renderContentType), ", Data: resp})")
	default:
		log.Fatalf("error: func %s []byte result must be set BytesRender", info.rpcMethodName)
	}
}

func (g *generate) printStringRender(info *routerInfo) {
	switch {
	case info.stringRender, info.textRender, info.htmlRender:
		g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Data"), "{ContentType: ", strconv.Quote(info.renderContentType), ", Data: []byte(resp)})")
	case info.redirectRender:
		g.P(g.functionBuf, "c.Redirect(statusCode, resp)")
	default:
		log.Fatalf("error: func %s string result must be set BytesRender or StringRender or TextRender or HTMLRender or RedirectRender", info.rpcMethodName)
	}
}

func (g *generate) printReaderRender(info *routerInfo) {
	switch {
	case info.readerRender:
		g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("Reader"), "{ContentType: ", strconv.Quote(info.renderContentType), ", ContentLength: -1, Reader: resp})")
	default:
		log.Fatalf("error: func %s io.Reader result must be set ReaderRender", info.rpcMethodName)
	}
}

func (g *generate) printObjectRender(info *routerInfo) {
	switch {
	case info.jsonRender:
		g.P(g.functionBuf, "c.JSON(statusCode, resp)")
	case info.indentedJSONRender:
		g.P(g.functionBuf, "c.IndentedJSON(statusCode, resp)")
	case info.secureJSONRender:
		g.P(g.functionBuf, "c.SecureJSON(statusCode, resp)")
	case info.jsonpJSONRender:
		g.P(g.functionBuf, "c.JSONP(statusCode, resp)")
	case info.pureJSONRender:
		g.P(g.functionBuf, "c.PureJSON(statusCode, resp)")
	case info.asciiJSONRender:
		g.P(g.functionBuf, "c.AsciiJSON(statusCode, resp)")
	case info.xmlRender:
		g.P(g.functionBuf, "c.XML(statusCode, resp)")
	case info.yamlRender:
		g.P(g.functionBuf, "c.YAML(statusCode, resp)")
	case info.protobufRender:
		g.P(g.functionBuf, "c.ProtoBuf(statusCode, resp)")
	case info.msgpackRender:
		g.P(g.functionBuf, "c.Render(statusCode, ", renderPackage.Ident("MsgPack"), "{Data: resp})")
	case info.tomlRender:
		g.P(g.functionBuf, "c.TOML(statusCode, resp)")
	default:
		log.Fatalf("error: func %s *struct{} result must be set JSONRender or IndentedJSONRender or SecureJSONRender or JsonpJSONRender or PureJSONRender or AsciiJSONRender or XMLRender or YAMLRender or ProtoBufRender or MsgPackRender or TOMLRender", info.rpcMethodName)
	}
}

func (g *generate) printImports() {
	g.P(g.importsBuf, "import (")
	for _, imp := range g.imports {
		if !imp.enable {
			continue
		}
		g.P(g.importsBuf, imp.PackageName, " ", strconv.Quote(imp.ImportPath))
	}
	g.P(g.importsBuf, ")")

}

func (g *generate) combine() {
	_, _ = io.Copy(g.buf, g.headerBuf)
	_, _ = io.Copy(g.buf, g.importsBuf)
	_, _ = io.Copy(g.buf, g.functionBuf)
}

func (g *generate) P(w io.Writer, v ...any) {
	for _, x := range v {
		switch x := x.(type) {
		case *goIdent:
			x.GoImport.enable = true
			g.imports[x.GoImport.ImportPath] = x.GoImport
			_, _ = fmt.Fprint(w, x.Qualify())
		default:
			_, _ = fmt.Fprint(w, x)
		}
	}
	_, _ = fmt.Fprintln(w)
}
