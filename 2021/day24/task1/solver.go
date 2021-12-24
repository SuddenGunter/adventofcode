package task1

import (
	"errors"
	"fmt"
)

func Solve() (int64, error) {
	alu := NewALU()

	old := int64(99999999999999)
	for i := int64(99999999999999); i >= 11111111111111; i-- {
		if old-i > 10000000 {
			fmt.Println("calculating ", i)
			old = i
		}
		if num, ok := alu.runMonadChecker(i); ok && num == 0 {
			return num, nil
		}

		alu.setState(registers{
			x: 0,
			y: 0,
			w: 0,
			z: 0,
		})
	}

	return 0, errors.New("solution not found")
}

type ALU struct {
	x, y, w, z *int64
	cache      map[cacheKey]registers
}

func NewALU() *ALU {
	cache := make(map[cacheKey]registers)
	x, y, z, w := int64(0), int64(0), int64(0), int64(0)
	return &ALU{cache: cache, x: &x, y: &y, z: &z, w: &w}
}

type cacheKey struct {
	num       int64
	registers registers
	operands  operands
}

type registers struct {
	x, y, w, z int64
}

type operands struct {
	op1, op2, op3 int64
}

func (a *ALU) runMonadChecker(num int64) (int64, bool) {
	nums, ok := split(num)
	if !ok {
		return 0, false
	}
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 11,
		op3: 1,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 11,
		op3: 11,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 14,
		op3: 1,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 11,
		op3: 11,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -8,
		op3: 2,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -5,
		op3: 9,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 11,
		op3: 7,
	})

	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -13,
		op3: 11,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 12,
		op3: 6,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -1,
		op3: 15,
	})
	a.func1(nums.pop(), operands{
		op1: 1,
		op2: 14,
		op3: 7,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -5,
		op3: 1,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -4,
		op3: 8,
	})
	a.func1(nums.pop(), operands{
		op1: 26,
		op2: -8,
		op3: 6,
	})

	return *a.z, true
}

func (a *ALU) setState(r registers) {
	*a.x = r.x
	*a.y = r.y
	*a.z = r.z
	*a.w = r.w
}

func (a *ALU) func1(num int64, operands operands) {
	key := cacheKey{
		num: num,
		registers: registers{
			x: *a.x,
			y: *a.y,
			w: *a.w,
			z: *a.z,
		},
		operands: operands,
	}
	cached, found := a.cache[key]

	if found {
		a.setState(cached)
		return
	}

	a.inp(a.w, num)
	a.mull(a.x, 0)
	a.add(a.x, a.z)
	a.modl(a.x, 26)
	a.divl(a.z, operands.op1)
	a.addl(a.x, operands.op2)
	a.eql(a.x, a.w)
	a.eqll(a.x, 0)
	a.mull(a.y, 0)
	a.addl(a.y, 25)
	a.mul(a.y, a.x)
	a.addl(a.y, 1)
	a.mul(a.z, a.y)
	a.mull(a.y, 0)
	a.add(a.y, a.w)
	a.addl(a.y, operands.op3)
	a.mul(a.y, a.x)
	a.add(a.z, a.y)

	a.cache[key] = registers{
		x: *a.x,
		y: *a.y,
		w: *a.w,
		z: *a.z,
	}
}

func split(num int64) (*stack, bool) {
	s := &stack{data: make([]int64, 0)}
	for i := 13.; i >= 0; i-- {
		x := num % 10
		if x == 0 {
			return nil, false
		}
		num /= 10
		s.push(x)
	}

	return s, true
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
