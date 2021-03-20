package main

import (
	"testing"

	"github.com/pitakill/test/mapper"
)

func TestMapString(t *testing.T) {
	tests := []struct {
		ss   SkipString
		want string
	}{
		{
			NewSkipString(3, "Aspiration.com"),
			"asPirAtiOn.cOm",
		},
		{
			NewSkipString(3, "test"),
			"teSt",
		},
		{
			NewSkipString(3, "this is another test"),
			"thIs iS anOthEr tEst",
		},
		{
			NewSkipString(3, "4nd 4n0th3r 0n3"),
			"4nD 4n0th3r 0N3",
		},
		{
			NewSkipString(3, "._.._-"),
			"._.._-",
		},
		{
			NewSkipString(3, "This. Look, g00d: to; ME' "),
			"thIs. lOok, G00D: to; Me' ",
		},
	}

	for _, test := range tests {
		mapper.MapString(&test.ss)

		if string(test.ss.out) != test.want {
			t.Errorf("%q and %q are not equal\n", string(test.ss.out), test.want)
		}
	}
}
