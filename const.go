package bloom

import "crypto/subtle"

// ConstantTimeTest returns true if the data is in the BloomFilter,
// false otherwise. If true, the result might be a false positive.
// If false, the data is definitely not in the set.
//
// It does this without leaking any information about data through
// timing side-channels.
//
// Note: no other methods in this package should be presumed to be
// free of side-channels. In particular, Add leaks information about
// the data.
func (f *BloomFilter) ConstantTimeTest(data []byte) bool {
	h := baseHashes(data)
	var v uint64
	for i, word := range f.b.Bytes() {
		v |= ^word & f.mask(h, i)
	}
	return constantTimeEq64(v, 0) == 1
}

func (f *BloomFilter) mask(h [4]uint64, i int) uint64 {
	var mask uint64
	for j := uint(0); j < f.k; j++ {
		loc := f.location(h, j)
		match := constantTimeEq64(uint64(loc/64), uint64(i))
		mask |= constantTimeLeftShift64(uint64(match), loc%64)
	}
	return mask
}

func constantTimeLeftShift64(x uint64, y uint) uint64 {
	/* https://www.bearssl.org/constanttime.html
	 *
	 * "Shifts and rotations can have an execution time that
	 *  depends on the shift/rotation count. This is not the
	 *  case on CPU that feature a “barrel shifter”. Famously,
	 *  the Pentium IV (NetBurst architecture) does not have
	 *  barrel shifters, so shift and rotation counts may leak
	 *  through timing. The shifted or rotated data, though,
	 *  does not leak. This point impacts mostly algorithms
	 *  that use shift or rotations by amounts that depend on
	 *  potentially secret data (e.g. RC5 encryption)."
	 */
	for i, y := 32, int(y); i > 0; i >>= 1 {
		v := subtle.ConstantTimeLessOrEq(i, y)
		x = constantTimeSelect64(v, x<<uint(i), x)
		y -= i & (0 - v)
	}
	return x
}

func constantTimeSelect64(v int, x, y uint64) uint64 {
	return ^uint64(v-1)&x | uint64(v-1)&y
}

func constantTimeEq64(x, y uint64) int {
	eq := x ^ y
	eq = ^eq & (eq - 1)
	return int(eq >> 63)
}
