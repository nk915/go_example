package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	// todo test code
	expected := 1
	actual := 0
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")
}

// 참조 사이트: https://lejewk.github.io/go-test/
