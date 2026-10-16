package alphabet

type CodeNode struct {
	Key  string
	Dot  *CodeNode
	Dash *CodeNode
}

var RootNode = CodeNode{
	Key:  "",
	Dot:  &ENode,
	Dash: &TNode,
}

// Nível 1
var ENode = CodeNode{
	Key:  "e",
	Dot:  &INode,
	Dash: &ANode,
}

var TNode = CodeNode{
	Key:  "t",
	Dot:  &NNode,
	Dash: &MNode,
}

// Nível 2
var INode = CodeNode{
	Key:  "i",
	Dot:  &SNode,
	Dash: &UNode,
}

var ANode = CodeNode{
	Key:  "a",
	Dot:  &RNode,
	Dash: &WNode,
}

var NNode = CodeNode{
	Key:  "n",
	Dot:  &DNode,
	Dash: &KNode,
}

var MNode = CodeNode{
	Key:  "m",
	Dot:  &GNode,
	Dash: &ONode,
}

// Nível 3
var SNode = CodeNode{
	Key:  "s",
	Dot:  &HNode,
	Dash: &VNode,
}

var UNode = CodeNode{
	Key: "u",
	Dot: &FNode,
}

var RNode = CodeNode{
	Key: "r",
	Dot: &LNode,
}

var WNode = CodeNode{
	Key:  "w",
	Dot:  &PNode,
	Dash: &JNode,
}

var DNode = CodeNode{
	Key:  "d",
	Dot:  &BNode,
	Dash: &XNode,
}

var KNode = CodeNode{
	Key:  "k",
	Dot:  &CNode,
	Dash: &YNode,
}

var GNode = CodeNode{
	Key:  "g",
	Dot:  &ZNode,
	Dash: &QNode,
}

var ONode = CodeNode{
	Key: "o",
}

// Nível 4 (folhas)
var HNode = CodeNode{Key: "h"}
var VNode = CodeNode{Key: "v"}
var FNode = CodeNode{Key: "f"}
var LNode = CodeNode{Key: "l"}
var PNode = CodeNode{Key: "p"}
var JNode = CodeNode{Key: "j"}
var BNode = CodeNode{Key: "b"}
var XNode = CodeNode{Key: "x"}
var CNode = CodeNode{Key: "c"}
var YNode = CodeNode{Key: "y"}
var ZNode = CodeNode{Key: "z"}
var QNode = CodeNode{Key: "q"}
