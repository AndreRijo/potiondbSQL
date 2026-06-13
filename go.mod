module sqlToKeyValue

go 1.22

require github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230321174746-8dcc6526cfb1

require google.golang.org/protobuf v1.34.2

require potionDB/crdt v0.0.0

require (
	github.com/planetscale/vtprotobuf v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)

replace potionDB/shared => ../potionDB/shared

replace potionDB/crdt => ../potionDB/crdt

replace github.com/AndreRijo/go-tools => ../goTools
