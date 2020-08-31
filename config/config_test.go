package config

import (
	"fmt"
	"testing"
)

func TestGetPort(t *testing.T) {
	port := GetPort()
	fmt.Printf("%s\n", port)
}

func TestGetSecret(t *testing.T) {
	secret := GetSecret()
	fmt.Printf("%s\n", secret)
}

func TestGetMySQL(t *testing.T) {
	mysql := GetMySQL()
	fmt.Printf("%#v\n", mysql)
}
