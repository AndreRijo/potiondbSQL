module sqlToKeyValue

go 1.16

require github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230321174746-8dcc6526cfb1

require google.golang.org/protobuf v1.34.2

require potionDB/crdt v0.0.0

replace potionDB/shared => ../potionDB/shared

replace potionDB/crdt => ../potionDB/crdt
