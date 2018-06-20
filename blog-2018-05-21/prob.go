package prob

// Here’s a problem that was asked by Cisco.
//
// Given an unsigned 8-bit integer, swap its even and odd bits. The 1st and 2nd
// bit should be swapped, the 3rd and 4th bit should be swapped, and so on.
//
// For example, 10101010 should be 01010101.
//
// Bonus: Can you do this in one line?

func SwapEvenOddBits(v byte) byte {
	const (
		evens byte = 85  // 0101 0101
		odds  byte = 170 // 1010 1010
	)
	return ((v & evens) << 1) | ((v & odds) >> 1)
}
