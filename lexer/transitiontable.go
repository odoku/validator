
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 39 : // [''',''']
				return 2
			case r == 40 : // ['(','(']
				return 3
			case r == 41 : // [')',')']
				return 4
			case r == 43 : // ['+','+']
				return 5
			case r == 44 : // [',',',']
				return 6
			case r == 45 : // ['-','-']
				return 5
			case r == 46 : // ['.','.']
				return 7
			case r == 47 : // ['/','/']
				return 8
			case 48 <= r && r <= 57 : // ['0','9']
				return 9
			case 65 <= r && r <= 90 : // ['A','Z']
				return 10
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 101 : // ['a','e']
				return 10
			case r == 102 : // ['f','f']
				return 12
			case 103 <= r && r <= 109 : // ['g','m']
				return 10
			case r == 110 : // ['n','n']
				return 13
			case 111 <= r && r <= 115 : // ['o','s']
				return 10
			case r == 116 : // ['t','t']
				return 14
			case 117 <= r && r <= 122 : // ['u','z']
				return 10
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 15
			
			
			default:
				return 2
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 41 : // [')',')']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 7
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 18
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 47 : // ['/','/']
				return 19
			
			
			default:
				return 8
			}
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 7
			case 48 <= r && r <= 57 : // ['0','9']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case r == 97 : // ['a','a']
				return 22
			case 98 <= r && r <= 122 : // ['b','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 104 : // ['a','h']
				return 21
			case r == 105 : // ['i','i']
				return 23
			case 106 <= r && r <= 122 : // ['j','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 113 : // ['a','q']
				return 21
			case r == 114 : // ['r','r']
				return 24
			case 115 <= r && r <= 122 : // ['s','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 7
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 18
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 107 : // ['a','k']
				return 21
			case r == 108 : // ['l','l']
				return 25
			case 109 <= r && r <= 122 : // ['m','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 107 : // ['a','k']
				return 21
			case r == 108 : // ['l','l']
				return 26
			case 109 <= r && r <= 122 : // ['m','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 116 : // ['a','t']
				return 21
			case r == 117 : // ['u','u']
				return 27
			case 118 <= r && r <= 122 : // ['v','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 114 : // ['a','r']
				return 21
			case r == 115 : // ['s','s']
				return 28
			case 116 <= r && r <= 122 : // ['t','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 29
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 100 : // ['a','d']
				return 21
			case r == 101 : // ['e','e']
				return 29
			case 102 <= r && r <= 122 : // ['f','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			case 65 <= r && r <= 90 : // ['A','Z']
				return 21
			case r == 95 : // ['_','_']
				return 11
			case 97 <= r && r <= 122 : // ['a','z']
				return 21
			
			
			
			}
			return NoState
			
		},
	
}
