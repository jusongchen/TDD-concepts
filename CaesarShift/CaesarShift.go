package CaesarShift

import (
	"unicode/utf8"
)

const charsetSize = 26

//Encode encodes an input by replacing letters 'a'-'z' and 'A'-'Z'
// with a different one a fixed number of places down the alphabet
// https://en.wikipedia.org/wiki/Caesar_cipher
func Encode(in string, offset int) string {
	if offset == 0 {
		return in
	}

	//limit offset to [0,charsetSize)
	offset = normalizeOffset(offset)
	out := ""
	for len(in) > 0 {
		from, size := utf8.DecodeRuneInString(in)
		in = in[size:]
		to := from
		switch {
		case (from >= rune('A') && from <= rune('Z')):
			to = from + rune(offset)
			if to > rune('Z') {
				to -= charsetSize
			}
		case from >= rune('a') && from <= rune('z'):
			to = from + rune(offset)
			if to > rune('z') {
				to -= charsetSize
			}
		default:
			//do nothing
		}
		out += string(to)
	}

	return out
}

//Decode is the reverse of the Encode function
func Decode(in string, offset int) string {
	if offset == 0 {
		return in
	}
	offset = normalizeOffset(offset)
	out := ""
	for len(in) > 0 {
		from, size := utf8.DecodeRuneInString(in)
		in = in[size:]
		to := from
		switch {
		case (from >= rune('A') && from <= rune('Z')):
			to = from - rune(offset)
			if to < rune('A') {
				to += charsetSize
			}
		case from >= rune('a') && from <= rune('z'):
			to = from - rune(offset)
			if to < rune('a') {
				to += charsetSize
			}
		default:
			//do nothing
		}
		out += string(to)
	}

	return out
}

//limit offset to [0,charsetSize)
func normalizeOffset(offset int) int {

	if offset < 0 {
		//find the lowest positive
		offset = offset + (1-offset/charsetSize)*charsetSize
	}
	if offset > 0 {
		//find the lowest positive
		offset = offset % charsetSize
	}
	return offset
}

// func main() {
// 	CaesarShift(os.Args[1])
// }
