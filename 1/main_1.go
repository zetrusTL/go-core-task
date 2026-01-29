package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func GetTypes(vars ...any) []string {
	res := make([]string, 0, len(vars))
	for _, v := range vars {
		res = append(res, reflect.TypeOf(v).String())
	}
	return res
}

func ToString(v any) string {
	return fmt.Sprint(v)
}

func ConcatToString(vars ...any) string {
	result := ""
	for _, v := range vars {
		result += ToString(v)
	}
	return result
}

func StringToRunes(s string) []rune {
	return []rune(s)
}

func HashRunesWithSalt(r []rune, salt string) string {
	middle := len(r) / 2

	withSalt := string(r[:middle]) + salt + string(r[middle:])

	hash := sha256.Sum256([]byte(withSalt))
	return hex.EncodeToString(hash[:])
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	values := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	types := GetTypes(values...)
	fmt.Println("Types:", types)

	joined := ConcatToString(values...)
	fmt.Println("Joined string:", joined)

	runes := StringToRunes(joined)
	fmt.Println("Runes:", runes)

	hash := HashRunesWithSalt(runes, "go-2024")
	fmt.Println("SHA256:", hash)
}
