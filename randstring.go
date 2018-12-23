package randstring

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// RandomStr is a structure providing the criteria when creating the random string
type RandomStr struct {

	// Length is the number of characters of the random string.
	// This excludes any prefix or posfix string length.
	Length int

	// CapitalsFlag determines if there should be at least one capital letter character in the
	// random string generated.
	CapitalsFlag bool

	// CapitalChars provides the list of capital characters to choose from when generating
	// the random string.
	CapitalChars []rune

	// LowercaseFlag determines if there should be at least one lowercase character in the
	// random string generated.
	LowercaseFlag bool

	// LowercaseChars provides the list of lowercase characters to choose from when generating
	// the random string.
	LowercaseChars []rune

	// DigitsFlag determines if there should be at least one numeral character in the
	// random string generated.
	DigitsFlag bool

	// DigitChars provides the list of numeral characters to choose from when generating
	// the random string.
	DigitChars string

	// SpecialsFlag determines if there should be at least one special character in the
	// random string generated.
	SpecialsFlag bool

	// SpecialChars provides the list of special characters to choose from when generating
	// the random string.
	SpecialChars string

	// Prefix provides a list of characters that should be added to the beginning of the random
	// string.
	Prefix []rune

	// Posfix provides a list of characters that should be added to the end of the random
	// string.
	Posfix []rune
}

// New creates a new RandomStr structure with default values.
func New(length int) RandomStr {

	return RandomStr{
		Length:         length,
		Prefix:         []rune(""),
		Posfix:         []rune(""),
		LowercaseFlag:  true,
		CapitalsFlag:   true,
		DigitsFlag:     true,
		SpecialsFlag:   true,
		CapitalChars:   []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		LowercaseChars: []rune("abcdefghijklmnopqrstuvwxyz"),
		DigitChars:     "0123456789",
		SpecialChars:   "^$*.[]{}()?-!@#%&/,><:;|_~=+",
	}
}

// SetPrefix sets a prefix value to be used as part of the string creation: [prefix][random string].
func (r *RandomStr) SetPrefix(prefix []rune) {
	r.Prefix = prefix
}

// SetPosfix sets the posfix value to be used as part of the string creation: [random string][posfix].
func (r *RandomStr) SetPosfix(posfix []rune) {
	r.Posfix = posfix
}

// Capitals determines whether there should be at least one capital letter character in the random string.
func (r *RandomStr) Capitals(flag bool) {
	r.CapitalsFlag = flag
}

// Lowercase determines whether there should be at least one lowercase character in the random string.
func (r *RandomStr) Lowercase(flag bool) {
	r.LowercaseFlag = flag
}

// Digits determines whether there should be at least one numeral character in the random string.
func (r *RandomStr) Digits(flag bool) {
	r.DigitsFlag = flag
}

// Specials determines whether there should be at least one special character in the random string.
func (r *RandomStr) Specials(flag bool) {
	r.SpecialsFlag = flag
}

// SetCapitalChars sets the list of capital characters to be used as part of the random string generation.
func (r *RandomStr) SetCapitalChars(chars []rune) {
	r.CapitalChars = chars
}

// SetLowercaseChars sets the list of lowercase characters to be used as part of the random string generation.
func (r *RandomStr) SetLowercaseChars(chars []rune) {
	r.LowercaseChars = chars
}

// SetDigitChars sets the list of number characters to be used as part of the random string generation
func (r *RandomStr) SetDigitChars(chars string) {
	r.DigitChars = chars
}

// SetSpecialChars sets the list of special characters to be used as part of the random string generation
func (r *RandomStr) SetSpecialChars(chars string) {
	r.SpecialChars = chars
}

func (r *RandomStr) requireCapitals() bool {
	return r.CapitalsFlag && len(r.CapitalChars) > 0
}

func (r *RandomStr) requireLowercase() bool {
	return r.LowercaseFlag && len(r.LowercaseChars) > 0
}

func (r *RandomStr) requireDigits() bool {
	return r.DigitsFlag && len(r.DigitChars) > 0
}

func (r *RandomStr) requireSpecialChars() bool {
	return r.SpecialsFlag && len(r.SpecialChars) > 0
}

func (r *RandomStr) buildAvailable() ([]rune, int) {
	var all []rune
	i := 0
	if r.requireCapitals() {
		all = append([]rune(r.CapitalChars))
		i++
	}
	if r.requireLowercase() {
		all = append([]rune(r.LowercaseChars))
		i++
	}
	if r.requireDigits() {
		all = append([]rune(r.DigitChars))
		i++
	}
	if r.requireSpecialChars() {
		all = append([]rune(r.SpecialChars))
		i++
	}
	return all, i
}

func (r *RandomStr) check() error {

	if r.CapitalsFlag && len(r.CapitalChars) == 0 {
		return errors.New("Restriction requires to have capitals, however no capital char list has been set")
	}

	if r.LowercaseFlag && len(r.LowercaseChars) == 0 {
		return errors.New("Restriction requires to have lowercase letters, however no lowercase char list has been set")
	}

	if r.DigitsFlag && len(r.DigitChars) == 0 {
		return errors.New("Restriction requires to have digits, however no digit char list has been set")
	}

	if r.SpecialsFlag && len(r.SpecialChars) == 0 {
		return errors.New("Restriction requires to have special characters, however no special char list has been set")
	}

	_, minLength := r.buildAvailable()
	if r.Length < minLength {
		return errors.New("The minimum length of the random string requires to be [" + string(minLength) + "] based on the restrictions selected")
	}
	return nil
}

// Build generates a random string based on the set of criteria.
func (r *RandomStr) Build() (string, error) {

	err := r.check()
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	all, _ := r.buildAvailable()
	index := 0
	var b strings.Builder

	if r.requireCapitals() {
		b.WriteRune(r.CapitalChars[rand.Intn(len(r.CapitalChars))])
		index++
	}
	if r.requireLowercase() {
		b.WriteRune(r.LowercaseChars[rand.Intn(len(r.LowercaseChars))])
		index++
	}
	if r.requireDigits() {
		b.WriteByte(r.DigitChars[rand.Intn(len(r.DigitChars))])
		index++
	}
	if r.requireSpecialChars() {
		b.WriteByte(r.SpecialChars[rand.Intn(len(r.SpecialChars))])
		index++
	}

	// fill the rest of the restricted length from any of the available charaters
	for i := index; i < r.Length; i++ {
		b.WriteRune(all[rand.Intn(len(all))])
	}

	buf := []rune(b.String())
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(r.Prefix) + string(buf) + string(r.Posfix), nil
}
