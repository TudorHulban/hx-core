package hxcore

import hxhelpers "github.com/TudorHulban/hx-core/helpers"

type HXInfo struct {
	Method                   string // if not provided will switch to new tab.
	Endpoint                 string
	CSSIDValidationDisable   string
	CSSIDValidationPasswords [3]string

	Swaps   []string // includes CSS ID sanitization
	Require []string // includes CSS ID sanitization
	Sends   []string // includes CSS ID sanitization

	Enable  []string // includes CSS ID sanitization
	Disable []string // includes CSS ID sanitization

	Show []string
	Hide []string

	OnChangeEnable []string // includes CSS ID sanitization

	LengthMax uint16
	LengthMin uint16
}

func (hx HXInfo) HavePasswordsValues() bool {
	return len(hx.CSSIDValidationPasswords[0]) > 0 && len(hx.CSSIDValidationPasswords[1]) > 0 && len(hx.CSSIDValidationPasswords[2]) > 0
}

func (hx HXInfo) PasswordsIDsSanitized() string {
	return hxhelpers.SanitizeCSSId(hx.CSSIDValidationPasswords[0]) +
		`,` +
		hxhelpers.SanitizeCSSId(hx.CSSIDValidationPasswords[1]) +
		`,` +
		hxhelpers.SanitizeCSSId(hx.CSSIDValidationPasswords[2])
}
