package sql

import (
	"potionDB/crdt/proto"
	"sqlToKeyValue/src/parser"
)

//Defines extra interfaces to be implemented by Listeners, or some common methods

type ProtoListener interface {
	parser.ViewSQLListener
	ToProtobuf() *proto.ApbTypedSQL
	FromProtobuf(*proto.ApbTypedSQL) ProtoListener
}
