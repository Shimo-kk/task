package user

import (
	"regexp"
	"task/app/core"
	"unicode/utf8"

	emailverifier "github.com/AfterShip/email-verifier"
)

// ワークスペースIDのバリデーション
func validateWorkspaceId(value int) error {
	if value <= 0 {
		return core.NewError(core.ValidationError, "ワークスペースIDは空の値を入力することはできません。")
	}

	return nil
}

// ユーザー名のバリデーション
func validateName(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "ユーザー名は空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "ユーザー名は50文字より大きい値を入力できません。")
	}

	return nil
}

// Emailのバリデーション
func validateEmail(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "Emailは空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 256 {
		return core.NewError(core.ValidationError, "Emailは256文字より大きい値を入力できません。")
	}

	verifier := emailverifier.NewVerifier()
	ret, err := verifier.Verify(value)

	if err != nil {
		return core.NewError(core.ValidationError, "Emailの検証に失敗しました。: "+err.Error())
	}

	if !ret.Syntax.Valid {
		return core.NewError(core.ValidationError, "Emailが不正です。")
	}

	return nil
}

// パスワードのバリデーション
func validatePassword(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "パスワードは空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) < 6 {
		return core.NewError(core.ValidationError, "パスワードは6文字より小さい値を入力できません。")
	}

	if utf8.RuneCountInString(value) > 128 {
		return core.NewError(core.ValidationError, "パスワードは128文字より大きい値を入力できません。")
	}

	reg := regexp.MustCompile(`[^a-zA-Z0-9!@#$%^&*()_+{}\[\]:;<>,.?/~\- ]`)
	if reg.MatchString(value) {
		return core.NewError(core.ValidationError, "パスワードは半角の英数字と記号以外を入力することはできません。")
	}

	return nil
}
