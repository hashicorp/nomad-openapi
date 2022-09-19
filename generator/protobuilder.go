package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"
)

type protoBuilder struct {
	handledModels map[string]interface{}
}

func (b *protoBuilder) buildProtosFromAPI() error {
	api := v1api{}

	b.handledModels = make(map[string]interface{})

	for _, path := range api.GetPaths() {
		for _, op := range path.Operations {
			err := b.buildOperation(op)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func (b *protoBuilder) buildOperation(op *operation) error {
	tmpl, err := template.ParseFiles("message.proto.tmpl")
	if err != nil {
		log.Print("parse files: ", err)
		return err
	}

	if op.RequestBody != nil {
		if _, ok := b.handledModels[op.RequestBody.Model.Name()]; !ok {
			err = tmpl.Execute(os.Stdout, newMessage(op.RequestBody.Model))
			if err != nil {
				log.Print("execute: ", err)
				return err
			}
			b.handledModels[op.RequestBody.Model.Name()] = op.RequestBody.Model
		}
	}

	if len(op.Responses) > 0 {
		for _, resp := range op.Responses {
			if resp.Content == nil || resp.Content.Model == nil {
				continue
			}

			if _, ok := b.handledModels[resp.Content.Model.Name()]; !ok {
				err = tmpl.Execute(os.Stdout, newMessage(resp.Content.Model))
				if err != nil {
					log.Print("execute: ", err)
					return err
				}
				b.handledModels[resp.Content.Model.Name()] = resp.Content.Model
			}
		}
	}

	return nil
}

func newMessage(typ reflect.Type) (message *Message) {
	message = &Message{}
	message.Name = typ.Name()

	if typ.Kind() != reflect.Struct {
		return message
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		message.fields = append(message.fields, newField(field, i))
	}
	return
}

type Message struct {
	Name   string
	fields []*Field
}

// TODO: Make field names more intelligent with tags
func (m *Message) Fields() []*Field {
	return m.fields
}

func newField(field reflect.StructField, index int) *Field {
	return &Field{
		Name:   field.Name,
		Type:   structTypeToProtobufType(field),
		number: index + 1,
	}
}

type Field struct {
	Name   string
	Type   string
	number int
}

// TODO: Make field numbers more intelligent with tags
func (f *Field) Number() int {
	return f.number
}

func structTypeToProtobufType(field reflect.StructField) string {
	name := field.Type.Name()

	if name == "" {
		name = "repeated " + field.Type.String()
	}

	switch name {
	case "float64":
		return "double"
	case "float32":
		return "float"
	case "int32":
		return "sint32"
	case "int64":
		return "sint64"
	case "uint32":
		return "fixed32"
	case "uint64":
		return "fixed64"
	case "[]byte":
		return "bytes"
	case "ptr":
		return "handlePtr"
	case "struct":
		return "handleStruct"
	default:
		return name
	}
}
