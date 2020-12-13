package day13

import "math/big"

// CongruenceConstraint a constraint for the searched number: number has to be equal to <remainder> mod <modulus>
type CongruenceConstraint struct {
	remainder int64
	// must be pairwise coprime
	modulus int64
}

// SimultaneousCongruence the set of congruence constraints that all must be satisfied by the searched number
type SimultaneousCongruence []CongruenceConstraint

// SolveSimultaneousCongruence search a number for which all constraints are satisfied (all modulus must be pairwise coprime)
func SolveSimultaneousCongruence(input SimultaneousCongruence) int64 {

	solution := new(big.Int)

	M := int64(1)
	for _, cc := range input {
		M *= cc.modulus
	}

	for _, cc := range input {
		remainder := new(big.Int).SetInt64(cc.remainder)
		m := new(big.Int).SetInt64(cc.modulus)
		Mi := new(big.Int).SetInt64(M / cc.modulus)

		r := new(big.Int)
		s := new(big.Int)
		new(big.Int).GCD(r, s, m, Mi)
		e := new(big.Int).Mul(s, Mi)
		solution.Add(solution, new(big.Int).Mul(e, remainder))
	}

	// a := new(big.Int).SetInt64(1)
	// b := new(big.Int)
	// new(big.Int).GCD(nil, nil, a, b)
	return solution.Mod(solution, new(big.Int).SetInt64(M)).Int64()
}
