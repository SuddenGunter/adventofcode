package task1

import (
	"errors"
)

func Solve() (int64, error) {
	alu := &ALU{}
	for i := int64(99999999999999); i >= 11111111111111; i-- {
		if num := alu.runMonadChecker(i); num != 0 {
			return num, nil
		}
	}

	return 0, errors.New("solution not found")
}

type ALU struct {
	x, y, w, z *int64
}

func (a *ALU) runMonadChecker(num int64) int64 {
	nums := split(num)
	a.func1(nums.pop(), 1, 11, 1)
	a.func1(nums.pop(), 11, 11, 1)
	a.func1(nums.pop(), 1, 14, 1)
	a.func1(nums.pop(), 11, 11, 1)
	a.func1(nums.pop(), 2, -8, 26)
	a.func1(nums.pop(), 9, -5, 26)
	a.func1(nums.pop(), 7, 11, 1)
	a.func1(nums.pop(), 11, -13, 26)
	a.func1(nums.pop(), 6, 12, 1)
	a.func1(nums.pop(), 15, -1, 26)
	a.func1(nums.pop(), 7, 14, 1)
	a.func1(nums.pop(), 1, -5, 26)
	a.func1(nums.pop(), 8, -4, 26)
	a.func1(nums.pop(), 6, -8, 26)

	return *a.z
}

func (a *ALU) func1(num int64, v1, v2, v3 int64) {
	a.inp(a.w, num)
	a.mull(a.x, 0)
	a.add(a.x, a.z)
	a.modl(a.x, 26)
	a.divl(a.z, v3)
	a.addl(a.x, v2)
	a.eql(a.x, a.w)
	a.eqll(a.x, 0)
	a.mull(a.y, 0)
	a.addl(a.y, 25)
	a.mul(a.y, a.x)
	a.addl(a.y, 1)
	a.mul(a.z, a.y)
	a.mull(a.y, 0)
	a.add(a.y, a.w)
	a.addl(a.y, v1)
	a.mul(a.y, a.x)
	a.add(a.z, a.y)
}

func split(num int64) *stack {
	s := &stack{data: make([]int64, 0)}
	for i := 13.; i == 0; i-- {
		x := num % 10
		num /= 10
		s.push(x)
	}

	return s
}

func (a *ALU) add(l, r *int64) {
	*l += *r
}

func (a *ALU) addl(l *int64, r int64) {
	*l += r
}

func (a *ALU) mul(l, r *int64) {
	*l *= *r
}

func (a *ALU) mull(l *int64, r int64) {
	*l *= r
}

func (a *ALU) mod(l, r *int64) {
	*l %= *r
}

func (a *ALU) modl(l *int64, r int64) {
	*l %= r
}

func (a *ALU) div(l, r *int64) {
	*l %= *r
}

func (a *ALU) divl(l *int64, r int64) {
	*l %= r
}

func (a *ALU) inp(l *int64, r int64) {
	*l = r
}

func (a *ALU) eql(l, r *int64) {
	if *l == *r {
		*l = 1
	} else {
		*l = 0
	}
}

func (a *ALU) eqll(l *int64, r int64) {
	if *l == r {
		*l = 1
	} else {
		*l = 0
	}
}

func (a *ALU) clear() {
	*a = ALU{}
}
