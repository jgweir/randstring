package randstring

import (
	"strings"
	"testing"
)

func TestRandStringValid(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}
}

func TestRandStringValidNonASCII(t *testing.T) {

	// given
	chars := []rune("ÅÄÖ")
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(true)
	r.SetCapitalChars(chars)
	r.Lowercase(false)
	r.Digits(false)
	r.Specials(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if (strings.Count(str, "") - 1) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d] string [%s]", expectedLength, (strings.Count(str, "") - 1), str)
	}
}

func TestRandStringValidCapitals(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if !strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain at least one capital, actual string [%s]", str)
	}
}

func TestRandStringValidDigits(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if !strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain at least one digit, actual string [%s]", str)
	}
}

func TestRandStringValidLowercase(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if !strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain at least one lowercase, actual string [%s]", str)
	}
}

func TestRandStringValidSpecialChars(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if !strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain at least one special character, actual string [%s]", str)
	}
}

func TestRandStringValidNoCapitals(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain no capitals, actual string [%s]", str)
	}
}

func TestRandStringValidNoLowercase(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Lowercase(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain no lowercase characters, actual string [%s]", str)
	}
}

func TestRandStringValidNoDigits(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Digits(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain no digit characters, actual string [%s]", str)
	}
}

func TestRandStringValidNoSpecialCharacters(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Specials(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain no special characters, actual string [%s]", str)
	}
}

func TestRandStringValidOnlyCapitals(t *testing.T) {

	// given
	chars := []rune("ABCD")
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(true)
	r.SetCapitalChars(chars)
	r.Lowercase(false)
	r.Digits(false)
	r.Specials(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if !strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain only capital letter characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain only capital letter characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain only capital letter characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain only capital letter characters, actual string [%s]", str)
	}
}

func TestRandStringValidOnlyLowercase(t *testing.T) {

	// given
	chars := []rune("abcd")
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(false)
	r.Lowercase(true)
	r.SetLowercaseChars(chars)
	r.Digits(false)
	r.Specials(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain only lowercase characters, actual string [%s]", str)
	}
	if !strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain only lowercase characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain only lowercase characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain only capital characters, actual string [%s]", str)
	}
}

func TestRandStringValidOnlyDigits(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(false)
	r.Lowercase(false)
	r.Digits(true)
	r.SetDigitChars("1234")
	r.Specials(false)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain only digit characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain only digit characters, actual string [%s]", str)
	}
	if !strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain only digit characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain only digit characters, actual string [%s]", str)
	}
}

func TestRandStringValidOnlySpecialCharacters(t *testing.T) {

	// given
	expectedLength := 8
	r := New(expectedLength)
	r.Capitals(false)
	r.Lowercase(false)
	r.Digits(false)
	r.Specials(true)
	r.SetSpecialChars("$%")

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != 8 {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}

	if strings.ContainsAny(str, string(r.CapitalChars)) {
		t.Errorf("Expecting the string to contain only special characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.LowercaseChars)) {
		t.Errorf("Expecting the string to contain only special characters, actual string [%s]", str)
	}
	if strings.ContainsAny(str, string(r.DigitChars)) {
		t.Errorf("Expecting the string to contain only special characters, actual string [%s]", str)
	}
	if !strings.ContainsAny(str, string(r.SpecialChars)) {
		t.Errorf("Expecting the string to contain only special characters, actual string [%s]", str)
	}
}

func TestRandStringValidPrefix(t *testing.T) {

	// given
	prefix := []rune("Prefix-")
	length := 8
	r := New(length)
	r.SetPrefix(prefix)
	expectedLength := length + len(r.Prefix)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != expectedLength {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}
	if !strings.Contains(str, string(r.Prefix)) {
		t.Errorf("Expecting the string to contain the prefix [%s] , actual string [%s]", string(r.Prefix), str)
	}
}

func TestRandStringValidPosfix(t *testing.T) {

	// given
	posfix := []rune("-Posfix")
	length := 8
	r := New(length)
	r.SetPosfix(posfix)
	expectedLength := length + len(r.Posfix)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != expectedLength {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}
	if !strings.Contains(str, string(r.Posfix)) {
		t.Errorf("Expecting the string to contain the posfix [%s] , actual string [%s]", string(r.Posfix), str)
	}
}

func TestRandStringValidPrefixPosfix(t *testing.T) {

	// given
	length := 8
	prefix := []rune("Prefix-")
	posfix := []rune("-Posfix")
	r := New(length)
	r.SetPrefix(prefix)
	r.SetPosfix(posfix)
	expectedLength := len(r.Prefix) + length + len(r.Posfix)

	// when
	str, err := r.Build()

	// then
	if err != nil {
		t.Errorf("An error has been returned [%s]", err)
	}
	if len(str) != expectedLength {
		t.Errorf("Expecting a string with length [%d], actual [%d]", expectedLength, len(str))
	}
	if !strings.Contains(str, string(r.Prefix)) {
		t.Errorf("Expecting the string to contain the prefix [%s] , actual string [%s]", string(r.Prefix), str)
	}
	if !strings.Contains(str, string(r.Posfix)) {
		t.Errorf("Expecting the string to contain the posfix [%s] , actual string [%s]", string(r.Posfix), str)
	}
}

func TestRandStringInvalidRestrictionLength(t *testing.T) {

	// given
	r := New(3)

	// when
	_, err := r.Build()

	// then
	if err == nil {
		t.Errorf("Expecting an error")
	}

}

func TestRandStringInvalidRestrictionNoCapitalChars(t *testing.T) {

	// given
	var emptyList []rune
	r := RandomStr{
		Length:       8,
		CapitalsFlag: true,
		CapitalChars: emptyList,
	}

	// when
	_, err := r.Build()

	// then
	if err == nil {
		t.Errorf("Expecting error")
	}
}

func TestRandStringInvalidRestrictionNoLowercaseChars(t *testing.T) {

	// given
	var emptyList []rune
	r := RandomStr{
		Length:         8,
		LowercaseFlag:  true,
		LowercaseChars: emptyList,
	}

	// when
	_, err := r.Build()

	// then
	if err == nil {
		t.Errorf("Expecting error")
	}
}

func TestRandStringInvalidRestrictionNoDigitChars(t *testing.T) {

	// given
	r := RandomStr{
		Length:     8,
		DigitsFlag: true,
		DigitChars: "",
	}

	// when
	_, err := r.Build()

	// then
	if err == nil {
		t.Errorf("Expecting error")
	}
}

func TestRandStringInvalidRestrictionNoSpecialChars(t *testing.T) {

	// given
	r := RandomStr{
		Length:       8,
		SpecialsFlag: true,
		SpecialChars: "",
	}

	// when
	_, err := r.Build()

	// then
	if err == nil {
		t.Errorf("Expecting error")
	}
}
