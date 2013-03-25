package mt19937

const (
	_N          = 624
	_M          = 397
	_UPPER_MASK = uint32(1 << 31)
	_LOWER_MASK = _UPPER_MASK - 1
	_a          = uint32(0x9908b0df)
	_b          = uint32(0x9d2c5680)
	_c          = uint32(0xefc60000)
	_u          = 11
	_s          = 7
	_t          = 15
	_l          = 18
)

type MT19937 struct {
	state [_N]uint32
	index uint32
}

func (mt *MT19937) Seed(seed uint32) {
	mt.state[0] = seed
	for mt.index = 1; mt.index < _N; mt.index++ {
		mt.state[mt.index] =
			uint32(1812433253)*(mt.state[mt.index-1]^(mt.state[mt.index-1]>>30)) + mt.index
	}
}

func (mt *MT19937) Uint32() uint32 {

	if mt.index >= _N {
		var y uint32
		var k int
		var mag = [2]uint32{0, _a}

		if mt.index == _N+1 {
			mt.Seed(0x1571)
		}

		for ; k < _N-_M; k++ {
			y = (mt.state[k] & _UPPER_MASK) | (mt.state[k+1] & _LOWER_MASK)
			mt.state[k] = mt.state[k+_M] ^ (y >> 1) ^ mag[int(y&1)]
		}

		for ; k < _N-1; k++ {
			y = (mt.state[k] & _UPPER_MASK) | (mt.state[k+1] & _LOWER_MASK)
			mt.state[k] = mt.state[k+(_M-_N)] ^ (y >> 1) ^ mag[int(y&1)]
		}

		y = (mt.state[_N-1] & _UPPER_MASK) | (mt.state[0] & _LOWER_MASK)
		mt.state[_N-1] = mt.state[_M-1] ^ (y >> 1) ^ mag[int(y&1)]
		mt.index = 0
	}

	y := mt.state[mt.index]
	mt.index++

	y ^= (y >> _u)
	y ^= (y << _s) & _b
	y ^= (y << _t) & _c
	y ^= (y >> _l)

	return y
}

func New() *MT19937 {
	return &MT19937{index: _N + 1}
}
