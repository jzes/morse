package alphabet

import "github.com/jzes/morse/decoder"

var RootNode = decoder.CodeNode{
	Key:  "",
	Dot:  &ENode,
	Dash: &TNode,
}

// Nível 1
var ENode = decoder.CodeNode{
	Key:  "e",
	Dot:  &INode,
	Dash: &ANode,
}

var TNode = decoder.CodeNode{
	Key:  "t",
	Dot:  &NNode,
	Dash: &MNode,
}

// Nível 2
var INode = decoder.CodeNode{
	Key:  "i",
	Dot:  &SNode,
	Dash: &UNode,
}

var ANode = decoder.CodeNode{
	Key:  "a",
	Dot:  &RNode,
	Dash: &WNode,
}

var NNode = decoder.CodeNode{
	Key:  "n",
	Dot:  &DNode,
	Dash: &KNode,
}

var MNode = decoder.CodeNode{
	Key:  "m",
	Dot:  &GNode,
	Dash: &ONode,
}

// Nível 3
var SNode = decoder.CodeNode{
	Key:  "s",
	Dot:  &HNode,
	Dash: &VNode,
}

var UNode = decoder.CodeNode{
	Key: "u",
	Dot: &FNode,
}

var RNode = decoder.CodeNode{
	Key: "r",
	Dot: &LNode,
}

var WNode = decoder.CodeNode{
	Key:  "w",
	Dot:  &PNode,
	Dash: &JNode,
}

var DNode = decoder.CodeNode{
	Key:  "d",
	Dot:  &BNode,
	Dash: &XNode,
}

var KNode = decoder.CodeNode{
	Key:  "k",
	Dot:  &CNode,
	Dash: &YNode,
}

var GNode = decoder.CodeNode{
	Key:  "g",
	Dot:  &ZNode,
	Dash: &QNode,
}

var ONode = decoder.CodeNode{
	Key: "o",
}

// Nível 4 (folhas)
var HNode = decoder.CodeNode{Key: "h"}
var VNode = decoder.CodeNode{Key: "v"}
var FNode = decoder.CodeNode{Key: "f"}
var LNode = decoder.CodeNode{Key: "l"}
var PNode = decoder.CodeNode{Key: "p"}
var JNode = decoder.CodeNode{Key: "j"}
var BNode = decoder.CodeNode{Key: "b"}
var XNode = decoder.CodeNode{Key: "x"}
var CNode = decoder.CodeNode{Key: "c"}
var YNode = decoder.CodeNode{Key: "y"}
var ZNode = decoder.CodeNode{Key: "z"}
var QNode = decoder.CodeNode{Key: "q"}
