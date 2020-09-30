package types_test

import (
	"testing"

	"github.com/eudaimathia/src/token"
	"github.com/eudaimathia/src/types"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Create System with ur-types Nat, Bool
	nat := types.NewUr("Nat")
	boolT := types.NewUr("Bool")
	sys := types.NewSystem()
	sys.AddUr(nat)
	sys.AddUr(boolT)
	parse := func(s string) types.T {
		return sys.Parse(token.NewStream(s))
	}
	parseAndAssert := func(s string, from, to types.T) types.T {
		t.Helper()
		typ := parse(s)
		assert.True(t, typ.From() == from)
		assert.True(t, typ.To() == to)
		return typ
	}

	assert.True(t, parse("Nat") == nat)
	assert.True(t, parse("Bool") == boolT)

	natPred := parseAndAssert("Nat->Bool", nat, boolT)

	parseAndAssert("Nat->Nat->Bool", nat, natPred)

	transPred := parse("(Nat->Nat)->Bool")
	transPred2 := parseAndAssert(" ((Nat -> Nat) -> Bool)", transPred.From(), boolT)
	assert.True(t, transPred == transPred2)
}
