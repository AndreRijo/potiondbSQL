package sql

//TODO: Uncomment this file and move it to some other project.
/*
import (
	"potionDB/potionDB/components"
	"potionDB/crdt/crdt"
	"potionDB/crdt/proto"
)

//This file is mostly for testing purposes.
//It provides items that can be used to test the SQL interface

type Lineitem struct {
	Orderkey      string
	Linenumber    string
	Extendedprice string
	Discount      string
	Orderdate     string
	Shippriority  string
	Shipdate      string
	Mktsegment    string
	N_Name        string
	C_Nationkey   string
	S_Nationkey   string
	R_Name        string
	P_Type        string
	Suppkey       string
	S_Name        string
	S_Address     string
	S_Phone       string
}

var L1 Lineitem = Lineitem{
	Orderkey: "1", Linenumber: "1", Suppkey: "1", C_Nationkey: "1", S_Nationkey: "1",
	Extendedprice: "10.5", Discount: "0.5", Mktsegment: "seg1",
	Orderdate: "1995-03-08", Shipdate: "1995-03-21", Shippriority: "LOW",
	N_Name: "INDONESIA", R_Name: "ASIA", P_Type: "toys",
	S_Name: "Lego", S_Address: "Somewhere", S_Phone: "123456789",
}
var L2 Lineitem = Lineitem{
	Orderkey: "2", Linenumber: "1", Suppkey: "2", C_Nationkey: "2", S_Nationkey: "2",
	Extendedprice: "43.5", Discount: "6.0", Mktsegment: "seg2",
	Orderdate: "1995-03-25", Shipdate: "1995-04-03", Shippriority: "LOW",
	N_Name: "CHINA", R_Name: "ASIA", P_Type: "toys",
	S_Name: "Supplier2", S_Address: "Somewhere2", S_Phone: "4328934278",
}
var L3 Lineitem = Lineitem{
	Orderkey: "3", Linenumber: "1", Suppkey: "1", C_Nationkey: "3", S_Nationkey: "1",
	Extendedprice: "10.5", Discount: "0.5", Mktsegment: "seg2",
	Orderdate: "1995-01-14", Shipdate: "1995-01-18", Shippriority: "HIGH",
	N_Name: "INDONESIA", R_Name: "ASIA", P_Type: "PROMOTION toys",
	S_Name: "Lego", S_Address: "Somewhere", S_Phone: "123456789",
}
var L4 Lineitem = Lineitem{
	Orderkey: "4", Linenumber: "1", Suppkey: "3", C_Nationkey: "3", S_Nationkey: "3",
	Extendedprice: "18", Discount: "3", Mktsegment: "seg1",
	Orderdate: "1995-02-26", Shipdate: "1995-03-03", Shippriority: "MEDIUM",
	N_Name: "GERMANY", R_Name: "EUROPE", P_Type: "tools",
	S_Name: "ToolsLDA", S_Address: "Somewhere over the rainbow", S_Phone: "3289123848",
}
var L5 Lineitem = Lineitem{
	Orderkey: "5", Linenumber: "1", Suppkey: "4", C_Nationkey: "4", S_Nationkey: "4",
	Extendedprice: "120", Discount: "12.0", Mktsegment: "seg3",
	Orderdate: "1995-03-16", Shipdate: "1995-03-16", Shippriority: "SAME_DAY_DELIVERY",
	N_Name: "FRANCE", R_Name: "EUROPE", P_Type: "PROMOTION perfume",
	S_Name: "Gucci", S_Address: "GucciLand", S_Phone: "4839279784",
}
var L6 Lineitem = Lineitem{
	Orderkey: "6", Linenumber: "1", Suppkey: "3", C_Nationkey: "3", S_Nationkey: "3",
	Extendedprice: "1345", Discount: "178", Mktsegment: "seg3",
	Orderdate: "1992-11-27", Shipdate: "1992-12-13", Shippriority: "LOW",
	N_Name: "INDONESIA", R_Name: "ASIA", P_Type: "PROMOTION jewellery",
	S_Name: "Pandora", S_Address: "Some factory", S_Phone: "3289147382",
}

type Orders struct {
	C_Name        string
	Custkey       string
	Orderkey      string
	Orderdate     string
	Totalprice    string
	Totalquantity string
}

var O1 Orders = Orders{
	Orderkey: "1", Custkey: "1", C_Name: "Edwin",
	Orderdate: "1995-03-08", Totalprice: "100", Totalquantity: "2",
}
var O2 Orders = Orders{
	Orderkey: "2", Custkey: "2", C_Name: "Caroline",
	Orderdate: "1995-03-25", Totalprice: "780", Totalquantity: "314",
}
var O3 Orders = Orders{
	Orderkey: "3", Custkey: "3", C_Name: "Kurt",
	Orderdate: "1995-01-14", Totalprice: "25", Totalquantity: "1",
}
var O4 Orders = Orders{
	Orderkey: "4", Custkey: "1", C_Name: "Edwin",
	Orderdate: "1995-02-26", Totalprice: "175", Totalquantity: "1",
}
var O5 Orders = Orders{
	Orderkey: "5", Custkey: "1", C_Name: "Edwin",
	Orderdate: "1995-03-16", Totalprice: "400", Totalquantity: "1",
}
var O6 Orders = Orders{
	Orderkey: "6", Custkey: "4", C_Name: "Elizabeth",
	Orderdate: "1992-11-27", Totalprice: "90", Totalquantity: "3",
}

type PartSupp struct {
	Partkey    string
	Suppkey    string
	Supplycost string
	Availqty   string
	N_Name     string
}

var P1 PartSupp = PartSupp{Partkey: "1", Suppkey: "1", Supplycost: "10", Availqty: "500", N_Name: "INDONESIA"}
var P2 PartSupp = PartSupp{Partkey: "2", Suppkey: "2", Supplycost: "2", Availqty: "10000", N_Name: "CHINA"}
var P3 PartSupp = PartSupp{Partkey: "3", Suppkey: "1", Supplycost: "7", Availqty: "2500", N_Name: "INDONESIA"}
var P4 PartSupp = PartSupp{Partkey: "4", Suppkey: "3", Supplycost: "40", Availqty: "400", N_Name: "GERMANY"}
var P5 PartSupp = PartSupp{Partkey: "5", Suppkey: "4", Supplycost: "15", Availqty: "150", N_Name: "FRANCE"}
var P6 PartSupp = PartSupp{Partkey: "6", Suppkey: "3", Supplycost: "5", Availqty: "1000", N_Name: "GERMANY"}

func (l Lineitem) ToMapUpdate() (params antidote.UpdateObjectParams) {
	upds := make(map[string]crdt.UpdateArguments)
	upds["Orderkey"] = crdt.SetValue{NewValue: l.Orderkey}
	upds["Linenumber"] = crdt.SetValue{NewValue: l.Linenumber}
	upds["Extendedprice"] = crdt.SetValue{NewValue: l.Extendedprice}
	upds["Discount"] = crdt.SetValue{NewValue: l.Discount}
	upds["Orderdate"] = crdt.SetValue{NewValue: l.Orderdate}
	upds["Shippriority"] = crdt.SetValue{NewValue: l.Shippriority}
	upds["Shipdate"] = crdt.SetValue{NewValue: l.Shipdate}
	upds["Mktsegment"] = crdt.SetValue{NewValue: l.Mktsegment}
	upds["N_Name"] = crdt.SetValue{NewValue: l.N_Name}
	upds["C_Nationkey"] = crdt.SetValue{NewValue: l.C_Nationkey}
	upds["S_Nationkey"] = crdt.SetValue{NewValue: l.S_Nationkey}
	upds["R_Name"] = crdt.SetValue{NewValue: l.R_Name}
	upds["P_Type"] = crdt.SetValue{NewValue: l.P_Type}
	upds["Suppkey"] = crdt.SetValue{NewValue: l.Suppkey}
	upds["S_Name"] = crdt.SetValue{NewValue: l.S_Name}
	upds["S_Address"] = crdt.SetValue{NewValue: l.S_Address}
	upds["S_Phone"] = crdt.SetValue{NewValue: l.S_Phone}
	upd := crdt.EmbMapUpdateAll{Upds: upds}
	var updArgs crdt.UpdateArguments = upd
	keyParams := antidote.KeyParams{Key: l.Orderkey + "_" + l.Linenumber, Bucket: "test", CrdtType: proto.CRDTType_RRMAP}
	return antidote.UpdateObjectParams{KeyParams: keyParams, UpdateArgs: &updArgs}
}

func (o Orders) ToMapUpdate() (params antidote.UpdateObjectParams) {
	upds := make(map[string]crdt.UpdateArguments)
	upds["C_Name"] = crdt.SetValue{NewValue: o.C_Name}
	upds["Custkey"] = crdt.SetValue{NewValue: o.Custkey}
	upds["Orderkey"] = crdt.SetValue{NewValue: o.Orderkey}
	upds["Orderdate"] = crdt.SetValue{NewValue: o.Orderdate}
	upds["Totalprice"] = crdt.SetValue{NewValue: o.Totalprice}
	upds["Totalquantity"] = crdt.SetValue{NewValue: o.Totalquantity}
	upd := crdt.EmbMapUpdateAll{Upds: upds}
	var updArgs crdt.UpdateArguments = upd
	keyParams := antidote.KeyParams{Key: o.Orderkey, Bucket: "test", CrdtType: proto.CRDTType_RRMAP}
	return antidote.UpdateObjectParams{KeyParams: keyParams, UpdateArgs: &updArgs}
}

func (p PartSupp) ToMapUpdate() (params antidote.UpdateObjectParams) {
	upds := make(map[string]crdt.UpdateArguments)
	upds["Partkey"] = crdt.SetValue{NewValue: p.Partkey}
	upds["Suppkey"] = crdt.SetValue{NewValue: p.Suppkey}
	upds["Supplycost"] = crdt.SetValue{NewValue: p.Supplycost}
	upds["Availqty"] = crdt.SetValue{NewValue: p.Availqty}
	upds["N_Name"] = crdt.SetValue{NewValue: p.N_Name}
	upd := crdt.EmbMapUpdateAll{Upds: upds}
	var updArgs crdt.UpdateArguments = upd
	keyParams := antidote.KeyParams{Key: p.Partkey, Bucket: "test", CrdtType: proto.CRDTType_RRMAP}
	return antidote.UpdateObjectParams{KeyParams: keyParams, UpdateArgs: &updArgs}
}

func (l Lineitem) GetKey() string {
	return l.Orderkey + "_" + l.Linenumber //The real key I think has more components
}

func (o Orders) GetKey() string {
	return o.Orderkey
}

func (p PartSupp) GetKey() string {
	return p.Partkey + "_" + p.Suppkey
}
*/
