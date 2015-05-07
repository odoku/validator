
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			shift(3),		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(4),		/* , */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Statement */
			reduce(1),		/* ,, reduce: Statement */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			shift(5),		/* () */
			shift(6),		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			shift(3),		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: Function */
			reduce(13),		/* ,, reduce: Function */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			shift(10),		/* identifire */
			shift(11),		/* nil */
			shift(12),		/* string */
			shift(13),		/* integer */
			shift(14),		/* float */
			shift(15),		/* boolean */
			shift(16),		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Statement */
			reduce(2),		/* ,, reduce: Statement */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(3),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(3),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(11),		/* ,, reduce: Arguments */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Arguments */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(4),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			shift(18),		/* () */
			shift(19),		/* ( */
			reduce(4),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(5),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(6),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(7),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(8),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(9),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(10),		/* ,, reduce: Expression */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Expression */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(20),		/* , */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			shift(21),		/* ) */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(13),		/* ,, reduce: Function */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(13),		/* ), reduce: Function */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			shift(10),		/* identifire */
			shift(11),		/* nil */
			shift(12),		/* string */
			shift(13),		/* integer */
			shift(14),		/* float */
			shift(15),		/* boolean */
			shift(16),		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* , */
			shift(10),		/* identifire */
			shift(11),		/* nil */
			shift(12),		/* string */
			shift(13),		/* integer */
			shift(14),		/* float */
			shift(15),		/* boolean */
			shift(16),		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: Function */
			reduce(14),		/* ,, reduce: Function */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			nil,		/* ) */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(20),		/* , */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			shift(24),		/* ) */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(12),		/* ,, reduce: Arguments */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(12),		/* ), reduce: Arguments */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(14),		/* ,, reduce: Function */
			nil,		/* identifire */
			nil,		/* nil */
			nil,		/* string */
			nil,		/* integer */
			nil,		/* float */
			nil,		/* boolean */
			nil,		/* regexp */
			nil,		/* () */
			nil,		/* ( */
			reduce(14),		/* ), reduce: Function */
			
		},

	},
	
}

