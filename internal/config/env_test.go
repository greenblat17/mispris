package config

import (
	"os"
	"testing"
)

func TestGetter(t *testing.T) {
	firstTest, secondTest := "first test", "second test"
	envForTest := "ENV_FOR_TEST"

	env := Getter(envForTest, firstTest)
	if env != firstTest {
		t.Fatal()
	}

	if err := os.Setenv(envForTest, secondTest); err != nil {
		env = Getter(envForTest, firstTest)
		if env == firstTest {
			t.Fatal()
		}
	}
}
