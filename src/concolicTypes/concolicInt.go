package concolicTypes

import "symTypes"
import "github.com/aclements/go-z3/z3"

// type ConcolicBool struct {
// 	Value bool
// 	Sym   sym.SymBool
// }

// func (concBool *ConcolicBool) equals(other ConcolicBool) ConcolicBool {
//   // strange stuff...
//   res =
//   return ConcolicBool{
//     Value : res,
//     Sym   :
//   }
// }

type ConcolicInt struct {
	Value     int
	Sym       symTypes.SymInt
  Constant  bool
}

func (self ConcolicInt) equals(o interface{}) ConcolicBool {
	// return concInt.Value == other.Value
	switch o.(type) {
	case int:
		res := self.Value == int(o)
	case ConcolicInt:
		res := self.Value == ConcolicInt(o).Value
	default:
    return ConcolicBool{Value: false, Sym: ...}
	}
  return ConcolicBool{Value:res, Sym: ... }
}

func (self ConcolicInt) notEquals(o interface{}) ConcolicBool {
  return !self.equals(o)
}

func (self ConcolicInt) lt(o interface{}) ConcolicBool {
  switch o.(type) {
  case int:
    res := self.Value < int(o)
    sym := self.Sym.SymIntZ3Expr().LT(z3.Int(int(o)))
  case ConcolicInt:
    res := self.Value < ConcolicInt(o).Value
    sym := self.Sym.SymIntZ3Expr().LT(ConcolicInt(o).Sym.SymIntZ3Expr())
  default:
    reportError("cannot compare with < : incompatible types", self, o)
    // do something?
  }
  return ConcolicBool{Value: res, Sym: sym}
}

func (self ConcolicInt) le(o interface{}) ConcolicBool {
  switch o.(type) {
  case int:
    res := self.Value <= int(o)
    sym := self.Sym.SymIntZ3Expr().LE(z3.Int(int(o)))
  case ConcolicInt:
    res := self.Value <= ConcolicInt(o).Value
    sym := self.Sym.SymIntZ3Expr().LE(ConcolicInt(o).Sym.SymIntZ3Expr())
  default:
    reportError("cannot compare with <= : incompatible types", self, o)
    // do something?
  }
  return ConcolicBool{Value: res, Sym: sym}
}

func (self ConcolicInt) gt(o interface{}) ConcolicBool {
  switch o.(type) {
  case int:
    res := self.Value > int(o)
    sym := self.Sym.SymIntZ3Expr().GT(z3.Int(int(o)))
  case ConcolicInt:
    res := self.Value > ConcolicInt(o).Value
    sym := self.Sym.SymIntZ3Expr().GT(ConcolicInt(o).Sym.SymIntZ3Expr())
  default:
    reportError("cannot compare with > : incompatible types", self, o)
    // do something?
  }
  return ConcolicBool{Value: res, Sym: sym}
}

func (self ConcolicInt) ge(o interface{}) ConcolicBool {
  switch o.(type) {
  case int:
    res := self.Value >= int(o)
    sym := self.Sym.SymIntZ3Expr().GE(z3.Int(int(o)))
  case ConcolicInt:
    res := self.Value >= ConcolicInt(o).Value
    sym := self.Sym.SymIntZ3Expr().GE(ConcolicInt(o).Sym.SymIntZ3Expr())
  default:
    reportError("cannot compare with >= : incompatible types", self, o)
    // do something?
  }
  return ConcolicBool{Value: res, Sym: sym}
}

func (self ConcolicInt) add(o interface{}) ConcolicInt {
  switch o.(type) {
  case int:
    res := self.Value + int(o)
    sym := self.Sym.SymIntz3Expr().Add(z3.Int(int(o)))
  case ConcolicInt:
    res := self.Value + ConcolicInt(o).Value
    sym := self.Sym.SymIntZ3Expr().Add(ConcolicInt(o).Sym.SymIntZ3Expr())
  default:
    // something went very wrong.
    return nil
  }
  return ConcolicInt{Value:res, Sym:sym}
}



