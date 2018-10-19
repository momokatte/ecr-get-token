package main

import (
	"testing"
)

func TestFormatToken_Base64(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "base64")
	if err != nil {
		t.Fatal(err.Error())
	}
	if out != token {
		t.Fatalf("Expected '%s', got '%s'", token, out)
	}
}

func TestFormatToken_Porcelain(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "porcelain")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "USERNAME PASSWORDPASSWORDPASSWORD\n"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Username(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "username")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "USERNAME"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Password(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "password")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "PASSWORDPASSWORDPASSWORD"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Json(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "json")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "{\"password\":\"PASSWORDPASSWORDPASSWORD\",\"username\":\"USERNAME\"}"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Shell(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "shell")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "USERNAME=USERNAME\nPASSWORD=PASSWORDPASSWORDPASSWORD\n"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Yaml(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	out, err := formatToken(token, "yaml")
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected := "username: \"USERNAME\"\npassword: \"PASSWORDPASSWORDPASSWORD\"\n"; out != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, out)
	}
}

func TestFormatToken_Unsupported(t *testing.T) {

	token := "VVNFUk5BTUU6UEFTU1dPUkRQQVNTV09SRFBBU1NXT1JE"

	_, err := formatToken(token, "rot13")
	if err == nil {
		t.Fatal("Expected error")
	}
}
