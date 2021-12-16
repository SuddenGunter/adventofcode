package task2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	val, err := Solve("C200B40A82")

	assert.NoError(t, err)
	assert.Equal(t, 3, val)
}

func TestSolve_Example2(t *testing.T) {
	val, err := Solve("04005AC33890")

	assert.NoError(t, err)
	assert.Equal(t, 54, val)
}

func TestSolve_Example3(t *testing.T) {
	val, err := Solve("880086C3E88112")

	assert.NoError(t, err)
	assert.Equal(t, 7, val)
}

func TestSolve_Example4(t *testing.T) {
	val, err := Solve("CE00C43D881120")

	assert.NoError(t, err)
	assert.Equal(t, 9, val)
}

func TestSolve_Example5(t *testing.T) {
	val, err := Solve("D8005AC2A8F0")

	assert.NoError(t, err)
	assert.Equal(t, 1, val)
}

func TestSolve_Example6(t *testing.T) {
	val, err := Solve("D8005AC2A8F0")

	assert.NoError(t, err)
	assert.Equal(t, 0, val)
}

func TestSolve_Example7(t *testing.T) {
	val, err := Solve("9C005AC2F8F0")

	assert.NoError(t, err)
	assert.Equal(t, 0, val)
}

func TestSolve_Example8(t *testing.T) {
	val, err := Solve("9C0141080250320F1802104A08")

	assert.NoError(t, err)
	assert.Equal(t, 1, val)
}
