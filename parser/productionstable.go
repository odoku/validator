
package parser

import(
  "github.com/odoku/validator/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Statement	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Function	<< ast.NewFunctionList(X[0]) >>`,
		Id: "Statement",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionList(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : Statement "," Function	<< ast.AppendFunction(X[0], X[2]) >>`,
		Id: "Statement",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendFunction(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expression : Function	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : identifire	<< ast.NewString(X[0], "") >>`,
		Id: "Expression",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0], "")
		},
	},
	ProdTabEntry{
		String: `Expression : nil	<< nil, nil >>`,
		Id: "Expression",
		NTType: 2,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Expression : string	<< ast.NewString(X[0], "'") >>`,
		Id: "Expression",
		NTType: 2,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0], "'")
		},
	},
	ProdTabEntry{
		String: `Expression : integer	<< ast.NewInteger(X[0]) >>`,
		Id: "Expression",
		NTType: 2,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInteger(X[0])
		},
	},
	ProdTabEntry{
		String: `Expression : float	<< ast.NewFloat(X[0]) >>`,
		Id: "Expression",
		NTType: 2,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFloat(X[0])
		},
	},
	ProdTabEntry{
		String: `Expression : boolean	<< ast.NewBoolean(X[0]) >>`,
		Id: "Expression",
		NTType: 2,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolean(X[0])
		},
	},
	ProdTabEntry{
		String: `Expression : regexp	<< ast.NewRegexp(X[0]) >>`,
		Id: "Expression",
		NTType: 2,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRegexp(X[0])
		},
	},
	ProdTabEntry{
		String: `Arguments : Expression	<< ast.NewArgumentList(X[0]) >>`,
		Id: "Arguments",
		NTType: 3,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArgumentList(X[0])
		},
	},
	ProdTabEntry{
		String: `Arguments : Arguments "," Expression	<< ast.AppendArgument(X[0], X[2]) >>`,
		Id: "Arguments",
		NTType: 3,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendArgument(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Function : identifire "()"	<< ast.NewFunction(X[0], nil) >>`,
		Id: "Function",
		NTType: 4,
		Index: 13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Function : identifire "(" Arguments ")"	<< ast.NewFunction(X[0], X[2]) >>`,
		Id: "Function",
		NTType: 4,
		Index: 14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[0], X[2])
		},
	},
	
}
