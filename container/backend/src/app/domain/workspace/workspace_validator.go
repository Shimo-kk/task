package workspace

import (
	"regexp"
	"task/app/core"
	"unicode/utf8"
)

// 名称のバリデーション
func validateName(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "ワークスペース名は空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "ワークスペース名は50文字より大きい値を入力できません。")
	}

	reg := regexp.MustCompile(`[^a-zA-Z0-9!@#$%^&*()_+{}\[\]:;<>,.?/~\- ]`)
	if reg.MatchString(value) {
		return core.NewError(core.ValidationError, "ワークスペース名は半角の英数字と記号以外を入力することはできません。")
	}

	return nil
}
