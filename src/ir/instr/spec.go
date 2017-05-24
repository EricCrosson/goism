package instr

// Jump instruction names.
var (
	Jmp       = []byte("goto")
	JmpNil    = []byte("goto-if-nil")
	JmpNotNil = []byte("goto-if-not-nil")
	Label     = []byte("label")
)

// Binary ops.
var (
	NumEq    = binOp("num=")
	NumGt    = binOp("num>")
	NumLt    = binOp("num<")
	NumSub   = binOp("sub")
	NumAdd   = binOp("add")
	NumMul   = binOp("mul")
	NumQuo   = binOp("quo")
	ArrayRef = binOp("array-ref")

	SetCar = binOp("setcar")
	SetCdr = binOp("setcdr")
)

// Unary ops.
var (
	Add1 = unaryOp("add1")
	Sub1 = unaryOp("sub1")

	Not = unaryOp("not")
	Neg = unaryOp("neg")

	Car = unaryOp("car")
	Cdr = unaryOp("cdr")
)

// Other ops without explicit parameter.
var (
	ArraySet = Instr{
		Name:     []byte("array-set"),
		Encoding: AttrEnc0,
		Input:    AttrTake3,
		Output:   AttrPushAndDiscard,
	}

	Return = Instr{
		Name:     []byte("return"),
		Encoding: AttrEnc0,
		Input:    AttrTake1,
	}
)

func Concat(argc int) Instr {
	return Instr{
		Name:     []byte("concat"),
		Encoding: AttrEnc1,
		Input:    AttrTakeN,
		Output:   AttrPushTmp,
		Data:     uint16(argc),
	}
}

func StackSet(stIndex int) Instr {
	return Instr{
		Name:     []byte("stack-set"),
		Encoding: AttrEnc1,
		Input:    AttrReplaceNth,
		Data:     uint16(stIndex),
	}
}

func VarSet(cvIndex int) Instr {
	return Instr{
		Name:     []byte("var-set"),
		Encoding: AttrEnc1,
		Input:    AttrTake1,
		Data:     uint16(cvIndex),
	}
}

func Discard(n int) Instr {
	return Instr{
		Name:     []byte("discard"),
		Encoding: AttrEnc1,
		Input:    AttrTakeN,
		Data:     uint16(n),
	}
}

func ConstRef(cvIndex int) Instr {
	return Instr{
		Name:     []byte("constant"),
		Encoding: AttrEnc1,
		Output:   AttrPushConst,
		Data:     uint16(cvIndex),
	}
}

func StackRef(stIndex int) Instr {
	return Instr{
		Name:     []byte("stack-ref"),
		Encoding: AttrEnc1,
		Output:   AttrDupNth,
		Data:     uint16(stIndex),
	}
}

func VarRef(cvIndex int) Instr {
	return Instr{
		Name:     []byte("var-ref"),
		Encoding: AttrEnc1,
		Output:   AttrPushTmp,
		Data:     uint16(cvIndex),
	}
}

func Call(argc int) Instr {
	return Instr{
		Name:     []byte("call"),
		Encoding: AttrEnc1,
		Input:    AttrTakeNplus1,
		Output:   AttrPushTmp,
		Data:     uint16(argc),
	}
}