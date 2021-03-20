package main

import "testing"

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"Aspiration.com",
			"asPirAtiOn.cOm",
		},
		{
			"test",
			"teSt",
		},
		{
			"this is another test",
			"thIs iS anOthEr tEst",
		},
		{
			"4nd 4n0th3r 0n3",
			"4nD 4n0th3r 0N3",
		},
		{
			"._.._-",
			"._.._-",
		},
		{
			"This. Look, g00d: to; ME' ",
			"thIs. lOok, G00D: to; Me' ",
		},
	}

	for _, test := range tests {
		output := CapitalizeEveryThirdAlphanumericChar(test.input)

		if output != test.expected {
			t.Errorf("%q and %q are not equal\n", output, test.expected)
		}
	}
}
