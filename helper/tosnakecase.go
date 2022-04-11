package helper

import (
	stringy "github.com/gobeam/stringy"
)

func ToSnakeCase(kata string) string {
	str := stringy.New(kata)
	snakeStr := str.SnakeCase("?", "")
	return snakeStr.ToLower()
}
