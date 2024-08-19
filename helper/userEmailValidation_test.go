package helper

import "testing"

func TestDidUserEmailPassValidation(t *testing.T) {
	var email string
	email = "test@gmail.com"
	didPass := DidUserEmailPassValidation(email)
	if didPass == false {
		t.Errorf("DidUserEmailPassValidation failed")
	}

	email = "testgmail.com"
	didPass = DidUserEmailPassValidation(email)
	if didPass == true {
		t.Errorf("DidUserEmailPassValidation failed")
	}

	email = "testgmail.com"
	didPass = DidUserEmailPassValidation(email)
	if didPass == true {
		t.Errorf("DidUserEmailPassValidation failed")
	}

	email = ""
	didPass = DidUserEmailPassValidation(email)
	if didPass == true {
		t.Errorf("DidUserEmailPassValidation failed")
	}
}

func BenchmarkDidUserEmailPassValidation_Valid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DidUserEmailPassValidation("test@gmail.com")
	}
}
