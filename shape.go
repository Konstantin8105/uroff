//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "os"
import "github.com/Konstantin8105/c4go/noarch"
import "unsafe"

// shape - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:24
func shape(s []int32) {
	//
	// * FARSI/ARABIC SHAPING PREPROCESSOR FOR NEATROFF
	// *
	// * Copyright (C) 2010-2014 Ali Gholami Rudi <ali at rudi dot ir>
	// *
	// * Permission to use, copy, modify, and/or distribute this software for any
	// * purpose with or without fee is hereby granted, provided that the above
	// * copyright notice and this permission notice appear in all copies.
	// *
	// * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
	// * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
	// * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
	// * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
	// * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
	// * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
	// * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
	//
	var c []int32
	var n []int32
	var cold int32
	var cnew int32
	for s[0] != 0 {
		if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(4) - int64(uintptr(unsafe.Pointer(&n[0])))/int64(4)) == 0 {
			s = s[0+1:]
		}
		for uc_comb(s[0]) != 0 {
			s = s[0+1:]
		}
		c = n
		n = s
		if c == nil || !(c[0]&65280 == 1536 || c[0]&65532 == 8204 || c[0]&65280 == 64256 || c[0]&65280 == 64512) {
			cold = 0
			continue
		}
		cnew = uc_shape(c[0], cold, func() int32 {
			if n != nil {
				return n[0]
			}
			return 0
		}())
		cold = c[0]
		c[0] = cnew
	}
}

// shape_ligs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:45
func shape_ligs(d []int32, s []int32) {
	var l int32
	for s[0] != 0 {
		if (func() int32 {
			l = uc_lig(d, s)
			return l
		}()) != 0 {
			s = s[0+l:]
		} else {
			d[0] = (func() []int32 {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		d = d[0+1:]
	}
	d[0] = 0
}

// raw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:60
var raw []byte = make([]byte, 8388608)

// utf8 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:61
var utf8 []int32 = make([]int32, 8388608)

// ligs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:62
var ligs []int32 = make([]int32, 8388608)

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/shape.c:64
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	xread(0, raw, int32(8388608))
	utf8_dec(utf8, raw)
	shape(utf8)
	shape_ligs(ligs, utf8)
	utf8_enc(raw, ligs)
	xwrite(1, raw, noarch.Strlen(raw))
	return
}

// achar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:6
// sorted list of characters that can be shaped
type achar struct {
	c uint32
	s uint32
	i uint32
	m uint32
	f uint32
}

// achars - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:6
var achars []achar = []achar{{1569, 65152, 0, 0, 0}, {1570, 65153, 0, 0, 65154}, {1571, 65155, 0, 0, 65156}, {1572, 65157, 0, 0, 65158}, {1573, 65159, 0, 0, 65160}, {1574, 65161, 65163, 65164, 65162}, {1575, 65165, 0, 0, 65166}, {1576, 65167, 65169, 65170, 65168}, {1577, 65171, 0, 0, 65172}, {1578, 65173, 65175, 65176, 65174}, {1579, 65177, 65179, 65180, 65178}, {1580, 65181, 65183, 65184, 65182}, {1581, 65185, 65187, 65188, 65186}, {1582, 65189, 65191, 65192, 65190}, {1583, 65193, 0, 0, 65194}, {1584, 65195, 0, 0, 65196}, {1585, 65197, 0, 0, 65198}, {1586, 65199, 0, 0, 65200}, {1587, 65201, 65203, 65204, 65202}, {1588, 65205, 65207, 65208, 65206}, {1589, 65209, 65211, 65212, 65210}, {1590, 65213, 65215, 65216, 65214}, {1591, 65217, 65219, 65220, 65218}, {1592, 65221, 65223, 65224, 65222}, {1593, 65225, 65227, 65228, 65226}, {1594, 65229, 65231, 65232, 65230}, {1600, 1600, 1600, 1600, 0}, {1601, 65233, 65235, 65236, 65234}, {1602, 65237, 65239, 65240, 65238}, {1603, 65241, 65243, 65244, 65242}, {1604, 65245, 65247, 65248, 65246}, {1605, 65249, 65251, 65252, 65250}, {1606, 65253, 65255, 65256, 65254}, {1607, 65257, 65259, 65260, 65258}, {1608, 65261, 0, 0, 65262}, {1609, 65263, 0, 0, 65264}, {1610, 65265, 65267, 65268, 65266}, {1662, 64342, 64344, 64345, 64343}, {1670, 64378, 64380, 64381, 64379}, {1688, 64394, 0, 0, 64395}, {1705, 64398, 64400, 64401, 64399}, {1711, 64402, 64404, 64405, 64403}, {1740, 64508, 64510, 64511, 64509}, {8204, 0, 0, 0, 0}, {8205, 0, 8205, 8205, 0}}

// find_achar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:60
func find_achar(c int32) []achar {
	// utf-8 code
	// single form
	// initial form
	// medial form
	// final form
	// hamza
	// alef madda
	// alef hamza above
	// waw hamza
	// alef hamza below
	// yeh hamza
	// alef
	// beh
	// teh marbuta
	// teh
	// theh
	// jeem
	// hah
	// khah
	// dal
	// thal
	// reh
	// zain
	// seen
	// sheen
	// sad
	// dad
	// tah
	// zah
	// ain
	// ghain
	// tatweel
	// feh
	// qaf
	// kaf
	// lam
	// meem
	// noon
	// heh
	// waw
	// alef maksura
	// yeh
	// peh
	// tcheh
	// jeh
	// fkaf
	// gaf
	// fyeh
	// ZWNJ
	// ZWJ
	var h int32
	var m int32
	var l int32
	h = int32(1080 / 24)
	l = 0
	for l < h {
		// using binary search to find c
		m = (h + l) >> uint64(1)
		if achars[m].c == uint32(c) {
			return achars[m:]
		}
		if uint32(c) < achars[m].c {
			h = m
		} else {
			l = m + 1
		}
	}
	return nil
}

// can_join - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:78
func can_join(c1 int32, c2 int32) int32 {
	var a1 []achar = find_achar(c1)
	var a2 []achar = find_achar(c2)
	return noarch.BoolToInt(len(a1) == 0 && len(a2) == 0 && (a1[0].i != 0 || a1[0].m != 0) && (a2[0].f != 0 || a2[0].m != 0))
}

// uc_shape - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:85
func uc_shape(cur int32, prev int32, next int32) int32 {
	var c int32 = cur
	var join_prev int32
	var join_next int32
	var ac []achar = find_achar(c)
	if ac == nil {
		// ignore non-Arabic characters
		return c
	}
	join_prev = can_join(prev, c)
	join_next = can_join(c, next)
	if join_prev != 0 && join_next != 0 {
		c = int32(ac[0].m)
	}
	if join_prev != 0 && noarch.Not(join_next) {
		c = int32(ac[0].f)
	}
	if noarch.Not(join_prev) && join_next != 0 {
		c = int32(ac[0].i)
	}
	if noarch.Not(join_prev) && noarch.Not(join_next) {
		// some fonts do not have a glyph for ac->s
		c = int32(ac[0].c)
	}
	if c != 0 {
		return c
	}
	return cur
}

// uc_comb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:122
func uc_comb(c int32) int32 {
	//
	// * return nonzero for Arabic combining characters
	// *
	// * The standard Arabic diacritics:
	// * + 0x064b: fathatan
	// * + 0x064c: dammatan
	// * + 0x064d: kasratan
	// * + 0x064e: fatha
	// * + 0x064f: damma
	// * + 0x0650: kasra
	// * + 0x0651: shadda
	// * + 0x0652: sukun
	// * + 0x0653: madda above
	// * + 0x0654: hamza above
	// * + 0x0655: hamza below
	// * + 0x0670: superscript alef
	//
	// the standard diacritics
	return noarch.BoolToInt(c >= 1611 && c <= 1621 || c >= 64606 && c <= 64611 || c == 1648)
}

// uc_lig2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:129
func uc_lig2(a1 int32, a2 int32) int32 {
	if a1 == 65247 && a2 == 65166 {
		// shadda ligatures
		// superscript alef
		// lam alef isolated
		return 65275
	}
	if a1 == 65248 && a2 == 65166 {
		// lam alef final
		return 65276
	}
	if a1 == 1617 && a2 == 1612 {
		// shadda dammatan
		return 64606
	}
	if a1 == 1617 && a2 == 1613 {
		// shadda kasratan
		return 64607
	}
	if a1 == 1617 && a2 == 1614 {
		// shadda fatha
		return 64608
	}
	if a1 == 1617 && a2 == 1615 {
		// shadda damma
		return 64609
	}
	if a1 == 1617 && a2 == 1616 {
		// shadda kasra
		return 64610
	}
	return 0
}

// uc_lig3 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:148
func uc_lig3(a1 int32, a2 int32, a3 int32) int32 {
	if a1 == 65247 && a3 == 65166 {
		if a2 == 1619 {
			// lam alef isolated
			return 65269
		}
		if a2 == 1620 {
			return 65271
		}
		if a2 == 1621 {
			return 65273
		}
	}
	if a1 == 65248 && a3 == 65166 {
		if a2 == 1619 {
			// lam alef final
			return 65270
		}
		if a2 == 1620 {
			return 65272
		}
		if a2 == 1621 {
			return 65274
		}
	}
	return 0
}

// uc_lig - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/uc.c:170
func uc_lig(dst []int32, src []int32) int32 {
	if src[1] != 0 && uc_lig2(src[0], src[1]) != 0 {
		// return the length of the ligature in src; writes the ligature to dst
		dst[0] = uc_lig2(src[0], src[1])
		return 2
	}
	if src[1] != 0 && src[2] != 0 && uc_lig3(src[0], src[1], src[2]) != 0 {
		dst[0] = uc_lig3(src[0], src[1], src[2])
		return 3
	}
	return 0
}

// xread - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:4
func xread(fd int32, buf interface{}, len_ int32) int32 {
	var n int32
	var r int32
	for (func() int32 {
		r = int32(noarch.Read(fd, buf[0+n:], uint32(len_-n)))
		return r
	}()) > 0 {
		n += r
	}
	return n
}

// xwrite - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:13
func xwrite(fd int32, buf interface{}, len_ int32) int32 {
	var n int32
	var w int32
	for (func() int32 {
		w = int32(noarch.Write(fd, buf[0+n:], uint32(len_-n)))
		return w
	}()) > 0 {
		n += w
	}
	return n
}

// readutf8 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:22
func readutf8(src [][]byte) int32 {
	var result int32
	var l int32 = 1
	var s []byte = src[0]
	for l < 6 && int32(uint8(s[0]))&(64>>uint64(l)) != 0 {
		l++
	}
	result = 63 >> uint64(l) & int32(uint8((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]))
	for func() int32 {
		defer func() {
			l--
		}()
		return l
	}() != 0 {
		result = result<<uint64(6) | int32(uint8((func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]))&63
	}
	src[0] = s
	return result
}

// utf8_dec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:36
func utf8_dec(dst []int32, src []byte) {
	var s []byte = src
	var d []int32 = dst
	for s[0] != 0 {
		if noarch.Not(^int32(uint8(s[0])) & 192) {
			(func() []int32 {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = readutf8((*[1000000][]byte)(unsafe.Pointer(&s))[:])
		} else {
			(func() []int32 {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = int32((func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0])
		}
	}
	d[0] = int32('\x00')
}

// writeutf8 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:49
func writeutf8(dst [][]byte, c int32) {
	var d []byte = dst[0]
	var l int32
	if c > 65535 {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = byte(240 | c>>uint64(18))
		l = 3
	} else if c > 2047 {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = byte(224 | c>>uint64(12))
		l = 2
	} else if c > 127 {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = byte(192 | c>>uint64(6))
		l = 1
	}
	for func() int32 {
		defer func() {
			l--
		}()
		return l
	}() != 0 {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = byte(128 | c>>uint64(l*6)&63)
	}
	dst[0] = d
}

// utf8_enc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/util.c:68
func utf8_enc(dst []byte, src []int32) {
	var s []int32 = src
	var d []byte = dst
	for s[0] != 0 {
		if s[0] & ^127 != 0 {
			writeutf8((*[1000000][]byte)(unsafe.Pointer(&d))[:], (func() []int32 {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0])
		} else {
			(func() []byte {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = byte((func() []int32 {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0])
		}
	}
	d[0] = '\x00'
}
