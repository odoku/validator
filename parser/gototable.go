
/*
*/
package parser

const numNTSymbols = 5
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Statement
		-1, // Expression
		-1, // Arguments
		2, // Function
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		7, // Function
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Statement
		9, // Expression
		17, // Arguments
		8, // Function
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Statement
		9, // Expression
		22, // Arguments
		8, // Function
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Statement
		23, // Expression
		-1, // Arguments
		8, // Function
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Statement
		-1, // Expression
		-1, // Arguments
		-1, // Function
		

	},
	
}
