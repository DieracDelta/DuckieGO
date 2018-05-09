package concolicTypes

import "fmt"
import "reflect"
import "github.com/aclements/go-z3/z3"

func concolicExecInput(testfunc reflect.Method, branchCtx *z3.Context, concreteValues *ConcreteValues) ([]reflect.Value, []z3.Bool) {
	var currPathConstrs []z3.Bool
	f := reflect.ValueOf(testfunc)
	args := []reflect.Value{ reflect.ValueOf(concreteValues), reflect.ValueOf(branchCtx), reflect.ValueOf(currPathConstrs) }
	res := f.Call(args)
	return res, currPathConstrs
}

func concolicForceBranch(branchNum int, branchCtx *z3.Context, branchConds ...z3.Bool) z3.Bool {
	if (branchNum < len(branchConds)) {
		cond := branchCtx.FromBool(true).And(branchConds[0:branchNum]...).And(branchConds[branchNum].Not())
		return cond
	} else {
		return branchCtx.FromBool(true)
	}
}

func concolicFindInput(branchCtx *z3.Context, constraint z3.Bool, names *ConcreteValues) (bool, *ConcreteValues) {
	solver := z3.NewSolver(branchCtx)
	solver.Assert(constraint)
	sat, err := solver.Check()
	newInput := newConcreteValues()
	if sat {
		model := solver.Model()
		for key, _ := range (names.getIntMappings()) {
			modelValue := model.Eval(branchCtx.IntConst(key), true)
			if modelValue != nil {
				value, isLiteral, ok := modelValue.(z3.Int).AsInt64()
				if isLiteral && ok {
					newInput.addIntValue(key, int(value))
				}
			}
		}
		return true, newInput
	} else if err != nil {
		panic(err)
	}
	return false, newInput
}

func concolicExec(testfunc reflect.Method, maxiter int) {
	// seenAlready := make(map[*z3.Bool]bool)
	inputs := initialConcreteValueQueue()
	iter := 0
	ctxConfig := z3.NewContextConfig()
	ctxConfig.SetUint("timeout", 5000)
	ctx := z3.NewContext(ctxConfig)
	for (iter < maxiter) && !(inputs.isEmpty()) {
		iter += 1
		inputThisTime := inputs.dequeue()
		_, branchConstrs := concolicExecInput(testfunc, ctx, inputThisTime)

		fmt.Printf(branchConstrs.AsAST().String())

		for b := 0; b < len(branchConstrs); b++ {
			oppConstr := concolicForceBranch(b, ctx, branchConstrs...)
			// if _, seen := seenAlready[oppConstr]; !seen {
				// seenAlready[oppConstr] = true
				newInputFound, newInput := concolicFindInput(ctx, oppConstr, inputThisTime)
				if newInputFound {
					newInput.inherit(inputThisTime)
					inputs.enqueue(newInput)
				}
			// }
		}
	}
}

func rubberducky(cv *ConcreteValues) {
	var i concolicTypes.ConcolicInt
	var j concolicTypes.ConcolicInt
	i = concolicTypes.ConcolicInt{cv.getIntValue("i"), symTypes.SymInt{"i", false}}
}

func rubberducky() {
	var i concolicTypes.ConcolicInt

	i = concolicTypes.ConcolicInt{5, symTypes.SymInt{"i", false}}

	i = i.Add(concolicTypes.ConcolicInt{1, symTypes.SymInt{true}})

	j := concolicTypes.ConcolicInt{69, symTypes.SymInt{"j", false}}

	i = i.Sub(concolicTypes.ConcolicInt{420, symTypes.SymInt{"", true}}.Add(j))

}

func main() {
	
}