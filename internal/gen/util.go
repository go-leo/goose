package gen

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func FindField(name string, inMessage *protogen.Message) *protogen.Field {
	for _, field := range inMessage.Fields {
		if string(field.Desc.Name()) == name {
			return field
		}
		if field.Desc.JSONName() == name {
			return field
		}
	}
	return nil
}

func FullFieldName(fields []*protogen.Field) string {
	var fieldNames []string
	for _, p := range fields {
		fieldNames = append(fieldNames, p.GoName)
	}
	fullFieldName := strings.Join(fieldNames, ".")
	return fullFieldName
}

// FieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func FieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType []any, pointer bool) {
	if field.Desc.IsWeak() {
		return []any{"struct{}"}, false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = []any{"bool"}
	case protoreflect.EnumKind:
		goType = []any{field.Enum.GoIdent}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = []any{"int32"}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = []any{"uint32"}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = []any{"int64"}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = []any{"uint64"}
	case protoreflect.FloatKind:
		goType = []any{"float32"}
	case protoreflect.DoubleKind:
		goType = []any{"float64"}
	case protoreflect.StringKind:
		goType = []any{"string"}
	case protoreflect.BytesKind:
		goType = []any{"[]byte"}
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = []any{"*", field.Message.GoIdent}
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return append([]any{"[]"}, goType...), false
	case field.Desc.IsMap():
		keyType, _ := FieldGoType(g, field.Message.Fields[0])
		valType, _ := FieldGoType(g, field.Message.Fields[1])
		return []any{"map[", keyType, "]", valType}, false
	}
	return goType, pointer
}
