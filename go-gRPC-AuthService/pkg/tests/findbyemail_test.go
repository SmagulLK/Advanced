package tests

import "testing"

func Findbyemail(t *testing.T) {
	tests := []string{"Alkey", "Tima", "beka"}
	for _, test := range tests {
		if mockfindemail(test) == nil {
			t.Errorf("Findbyemail(%s) = %s; want %s", test, mockfindemail(test), test)
		}
	}
}
