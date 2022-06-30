package util

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "123abcABC!"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	if !(len(hashed) == 60) {
		t.Fatalf("Invalid password hash: %s", hashed)
	}
}

func TestCompareHashPasswordWithSamePassword(t *testing.T) {
	password := "123abcABC!"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	compareResult := CompareHashPassword(password, hashed)
	if !compareResult {
		t.Fatalf("Compare result is: %v, Expected true.\n", compareResult)
	}
}

func TestCompareHashPasswordWithDifferentPassword(t *testing.T) {
	var (
		password1 = "123abcABC!"
		password2 = "123ABCabc!"
	)
	hashed, err := HashPassword(password1)
	if err != nil {
		t.Fatal(err)
	}

	compareResult := CompareHashPassword(password2, hashed)
	if compareResult {
		t.Fatalf("Compare result is: %v, Expected false.\n", compareResult)
	}
}
