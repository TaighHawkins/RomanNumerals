package roman

import (
	"testing"
)

func TestIToInteger(t *testing.T) {
	CheckToInt("I", 1, t)
}

func TestVToInteger(t *testing.T) {
	CheckToInt("V", 5, t)
}

func TestXToInteger(t *testing.T) {
	CheckToInt("X", 10, t)
}

func TestLToInteger(t *testing.T) {
	CheckToInt("L", 50, t)
}

func TestCToInteger(t *testing.T) {
	CheckToInt("C", 100, t)
}

func TestDToInteger(t *testing.T) {
	CheckToInt("D", 500, t)
}

func TestMToInteger(t *testing.T) {
	CheckToInt("M", 1000, t)
}
func TestToIntegerMultipleCharacters1(t *testing.T) {
	CheckToInt("LXXXVII", 87, t)
}

func TestToIntegerMultipleCharacters2(t *testing.T) {
	CheckToInt("XCIV", 94, t)
}

func TestToIntegerMultipleCharacters3(t *testing.T) {
	CheckToInt("LXIX", 69, t)
}

func TestToIntegerMultipleCharacters4(t *testing.T) {
	CheckToInt("XLIX", 49, t)
}

func TestToIntegerMultipleCharacters5(t *testing.T) {
	CheckToInt("MLXVI", 1066, t)
}

func TestToIntegerMultipleCharacters6(t *testing.T) {
	CheckToInt("MCMLXXXVIII", 1988, t)
}

func TestToNumerals1(t *testing.T) {
	CheckToNumeral(87, "LXXXVII", t)
}

func TestToNumerals2(t *testing.T) {
	CheckToNumeral(94, "XCIV", t)
}

func TestToNumerals3(t *testing.T) {
	CheckToNumeral(69, "LXIX", t)
}

func TestToNumerals4(t *testing.T) {
	CheckToNumeral(49, "XLIX", t)
}

func TestToNumerals5(t *testing.T) {
	CheckToNumeral(1066, "MLXVI", t)
}

func TestToNumerals6(t *testing.T) {
	CheckToNumeral(2421, "MMCDXXI", t)
}

func TestToNumerals7(t *testing.T) {
	CheckToNumeral(1988, "MCMLXXXVIII", t)
}

func CheckToInt(input string, want int, t *testing.T) {
	out, err := ToInt(input)
	if out != want || err != nil {
		t.Fatalf(`DecodeToNumber("%q") = %q, %v, want match for %#q, nil`, input, out, err, want)
	}
}

func CheckToNumeral(input int, want string, t *testing.T) {
	out, err := ToNumerals(input)
	if out != want || err != nil {
		t.Fatalf(`EncodeToNumerals("%d") = %q, %v, want match for %#q, nil`, input, out, err, want)
	}
}
