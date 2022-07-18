package mobile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*func Test_ParsePhoneNumber(t *testing.T) {
	got := ParsePhoneNumber("+3801112233", UA)
	t.Log(got)
	exp := "(UA)01112233"
	if got != exp {
		t.FailNow()
	}

}*/

func Test_ParsePhoneNumber(t *testing.T) {
	tests := map[string]struct {
		number  string
		country string
		exp     string
		expErr  string
	}{
		"success":         {number: "+3801112233", country: "UA", exp: "(UA)01112233", expErr: ""},
		"invalid country": {number: "+3801112233", country: "UA", exp: "(UA)01112233", expErr: ""},
	}
	for name, tt := range tests {

	}

}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Value)

	}

}
