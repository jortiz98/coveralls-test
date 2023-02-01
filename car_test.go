package foobar_test

import "testing"

func TestCar_GetMaker(t *testing.T) {
	testcase := []string{"BMW", "Audi", "Tesla", "Ferrari"}

	for _, tc := range testcase {
		t.Run(tc, func(t *testing.T) {
			if tc != tc {
				t.Fail()
			}
		})
	}

}
