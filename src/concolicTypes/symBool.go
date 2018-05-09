package concolicTypes

import "github.com/aclements/go-z3/z3"

type SymBool struct {
	//id string
	z3Expr z3.Bool
}

func makeSymBoolVar(name string) SymBool {
	return SymBool{ctx.BoolConst(name)}
}

func (self *SymBool) SymBoolZ3Expr(ctx *z3.Context) z3.Bool {
	return self.z3Expr //ctx.BoolConst(self.id)
}
