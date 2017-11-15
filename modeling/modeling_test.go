package modeling

import "testing"

func TestGen(t *testing.T) {
	values, err := GenParetoN(0.05, 3.14, 100)
	if err != nil {
		t.Error(err)
	}
	t.Log(values)
}
