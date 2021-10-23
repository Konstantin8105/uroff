//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

// Warning (*ast.BinaryOperator):  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:210 :cannot transpile BinaryOperator with type 'int' : result type = {PointerOperation_unknown05}. Error: operator is `==`. {'struct font *' == 'struct font *'}. for base type: `struct font`. PntCmpPnt:SubTwoPnts:GetPointerAddress:sizeof:0. not valid sizeof `struct font *`: 0
// Warning (*ast.IfStmt):  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:210 :cannot transpileToStmt : cannot transpileIfStmt. cannot transpile for condition. cannot transpileToExpr. err = cannot transpile BinaryOperator with type 'int' : result type = {PointerOperation_unknown05}. Error: operator is `==`. {'struct font *' == 'struct font *'}. for base type: `struct font`. PntCmpPnt:SubTwoPnts:GetPointerAddress:sizeof:0. not valid sizeof `struct font *`: 0

package main

// #include </usr/include/unistd.h>
// #include </usr/include/string.h>
import "C"

import "unicode"
import "os"
import "golang.org/x/sys/unix"
import "reflect"
import "runtime"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.h:141
//
// * Most functions and variables in neatroff are prefixed with tokens
// * that indicate their purpose, such as:
// *
// * + tr_xyz: the implementation of troff request .xyz (mostly tr.c)
// * + in_xyz: input layer (in.c)
// * + cp_xyz: copy-mode interpretation layer (cp.c)
// * + ren_xyz: rendering characters into lines (ren.c)
// * + out_xyz: output layer for generating troff output (out.c)
// * + dev_xyz: output devices (dev.c)
// * + num_xyz: number registers (reg.c)
// * + str_xyz: string registers (reg.c)
// * + env_xyz: environments (reg.c)
// * + eval_xyz: integer expression evaluation (eval.c)
// * + font_xyz: fonts (font.c)
// * + sbuf_xyz: variable length string buffers (sbuf.c)
// * + dict_xyz: dictionaries (dict.c)
// * + wb_xyz: word buffers (wb.c)
// * + fmt_xyz: line formatting buffers (fmt.c)
// * + n_xyz: builtin number register xyz
// * + c_xyz: characters for requests like hc and mc
// *
//
// predefined array limits
// converting scales
// escape sequences
// special characters
// escape character (\)
// basic control character (.)
// no-break control character (')
// page number character (%)
// number registers
// string registers
// saving and restoring registers before and after printing diverted lines
// enviroments
// mapping integers to sets
// mapping strings to longs
// device related variables
type glyph struct {
	id    [32]byte
	name  [32]byte
	font  []font
	wid   int16
	llx   int16
	lly   int16
	urx   int16
	ury   int16
	type_ int16
}

// sbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.h:211
// device-dependent glyph identifier
// the first character mapped to this glyph
// glyph font
// character width
// character bounding box
// character type; ascender/descender
// output device functions
// font-related functions
// different layers of neatroff
// input layer
// copy-mode layer
// troff layer
// .so request
// .nx request
// .ex request
// .lf request
// queue the given input file
// look up argument
// number of arguments
// shift the arguments
// push back input character
// the first pushed-back character
// current filename
// current line number
// skip or read the next line or block
// beginning of a request line
// do not interpret \w and \E
// read the next troff request
// execute a built-in troff request
// variable length string buffer
type sbuf struct {
	s  []byte
	sz int32
	n  int32
}

// wb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.h:229
// allocated buffer
// buffer size
// length of the string stored in s
// word buffer
type wb struct {
	sbuf        sbuf
	f           int32
	s           int32
	m           int32
	cd          int32
	r_f         int32
	r_s         int32
	r_m         int32
	r_cd        int32
	part        int32
	cost        int32
	els_neg     int32
	els_pos     int32
	h           int32
	v           int32
	ct          int32
	sb          int32
	st          int32
	llx         int32
	lly         int32
	urx         int32
	ury         int32
	icleft      int32
	sub_c       [256][]byte
	sub_n       int32
	sub_collect int32
}

// utf8len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:9
func utf8len(c int32) int32 {
	if ^c&192 != 0 {
		// reading characters and escapes
		// return the length of a utf-8 character based on its first byte
		// ASCII or invalid
		return noarch.BoolToInt(c > 0)
	}
	if ^c&32 != 0 {
		return 2
	}
	if ^c&16 != 0 {
		return 3
	}
	if ^c&8 != 0 {
		return 4
	}
	return 1
}

// utf8one - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:23
func utf8one(s []byte) int32 {
	// return nonzero if s is a single utf-8 character
	return noarch.BoolToInt(noarch.Not(s[utf8len(int32(uint8(s[0])))]))
}

// utf8read - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:29
func utf8read(s [][]byte, d []byte) int32 {
	// read a utf-8 character from s and copy it to d
	var l int32 = utf8len(int32(uint8((s[0])[0])))
	var i int32
	for i = 0; i < l; i++ {
		d[i] = (s[0])[i]
	}
	d[l] = '\x00'
	s[0] = (s[0])[0+l:]
	return l
}

// utf8next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:41
func utf8next(s []byte, next func() int32) int32 {
	// read a utf-8 character with next() and copy it to s
	var c int32 = next()
	var l int32 = utf8len(c)
	var i int32
	if c < 0 {
		return 0
	}
	s[0] = byte(c)
	for i = 1; i < l; i++ {
		s[i] = byte(next())
	}
	s[l] = '\x00'
	return l
}

// quotednext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:56
func quotednext(next func() int32, back func(int32)) []byte {
	// read quoted arguments of escape sequences (ESC_Q)
	var delim []byte = make([]byte, 32)
	var cs []byte = make([]byte, 32)
	var sb sbuf
	var d []byte = make([]byte, 32)
	charnext(delim, next, back)
	sbuf_init(c4goUnsafeConvert_sbuf(&sb))
	for charnext_delim(cs, next, back, delim) >= 0 {
		charnext_str(d, cs)
		sbuf_append(c4goUnsafeConvert_sbuf(&sb), d)
	}
	return sbuf_out(c4goUnsafeConvert_sbuf(&sb))
}

// unquotednext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:71
func unquotednext(cmd int32, next func() int32, back func(int32)) []byte {
	// read unquoted arguments of escape sequences (ESC_P)
	var c int32 = next()
	var sb sbuf
	sbuf_init(c4goUnsafeConvert_sbuf(&sb))
	if cmd == int32('s') && (c == int32('-') || c == int32('+')) {
		cmd = c
		sbuf_add(c4goUnsafeConvert_sbuf(&sb), c)
		c = next()
	}
	if c == int32('(') {
		sbuf_add(c4goUnsafeConvert_sbuf(&sb), next())
		sbuf_add(c4goUnsafeConvert_sbuf(&sb), next())
	} else if noarch.Not((nreg(int32('C')))[0]) && c == int32('[') {
		c = next()
		for c > 0 && c != int32('\n') && c != int32(']') {
			sbuf_add(c4goUnsafeConvert_sbuf(&sb), c)
			c = next()
		}
	} else {
		sbuf_add(c4goUnsafeConvert_sbuf(&sb), c)
		if (nreg(int32('C')))[0] != 0 && cmd == int32('s') && c >= int32('1') && c <= int32('3') {
			c = next()
			if int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit)) != 0 {
				sbuf_add(c4goUnsafeConvert_sbuf(&sb), c)
			} else {
				back(c)
			}
		}
	}
	return sbuf_out(c4goUnsafeConvert_sbuf(&sb))
}

// charnext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:114
func charnext(c []byte, next func() int32, back func(int32)) int32 {
	//
	// * read the next character or escape sequence (x, \x, \(xy, \[xyz], \C'xyz')
	// *
	// * character returned contents of c
	// * x                  '\0'  x
	// * \4x                c_ni  \4x
	// * \\x                '\\'  \\x
	// * \\(xy     '('           xy
	// * \\[xyz]   '['           xyz
	// * \\C'xyz'  'C'           xyz
	//
	var l int32
	var n int32
	if noarch.Not(utf8next(c, next)) {
		return -1
	}
	if int32(c[0]) == 4 {
		utf8next(c[0+1:], next)
		return 4
	}
	if int32(c[0]) == c_ec {
		utf8next(c[0+1:], next)
		if int32(c[1]) == int32('(') {
			l = utf8next(c, next)
			l += utf8next(c[0+l:], next)
			return int32('(')
		} else if noarch.Not((nreg(int32('C')))[0]) && int32(c[1]) == int32('[') {
			l = 0
			n = next()
			for n >= 0 && n != int32('\n') && n != int32(']') && l < 32-1 {
				c[func() int32 {
					defer func() {
						l++
					}()
					return l
				}()] = byte(n)
				n = next()
			}
			c[l] = '\x00'
			return int32('[')
		} else if int32(c[1]) == int32('C') {
			var chr []byte = quotednext(next, back)
			noarch.Snprintf(c, 32, []byte("%s\x00"), chr)
			_ = chr
			return int32('C')
		}
		return int32('\\')
	}
	return int32('\x00')
}

// charnext_delim - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:150
func charnext_delim(c []byte, next func() int32, back func(int32), delim []byte) int32 {
	// like nextchar(), but return -1 if delim was read
	var t int32 = charnext(c, next, back)
	if noarch.Strcmp(c, delim) != 0 {
		return t
	}
	return -1
}

// charnext_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:157
func charnext_str(d []byte, c []byte) {
	// convert back the character read from nextchar() (e.g. xy -> \\(xy)
	var c0 int32 = int32(uint8(c[0]))
	if c0 == c_ec || c0 == 4 || noarch.Not(c[1]) || utf8one(c) != 0 {
		noarch.Strcpy(d, c)
		return
	}
	if noarch.Not(c[2]) && utf8len(c0) == 1 {
		noarch.Sprintf(d, []byte("%c(%s\x00"), c_ec, c)
	} else {
		noarch.Sprintf(d, []byte("%cC'%s'\x00"), c_ec, c)
	}
}

// charread - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:171
func charread(s [][]byte, c []byte) int32 {
	// like charnext() for string buffers
	var ret int32
	sstr_push(s[0])
	ret = charnext(c, sstr_next, sstr_back)
	s[0] = sstr_pop()
	return ret
}

// charread_delim - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:181
func charread_delim(s [][]byte, c []byte, delim []byte) int32 {
	// like charnext_delim() for string buffers
	var ret int32
	sstr_push(s[0])
	ret = charnext_delim(c, sstr_next, sstr_back, delim)
	s[0] = sstr_pop()
	return ret
}

// quotedread - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:191
func quotedread(sp [][]byte, d []byte) {
	// read quoted arguments; this is called only for internal neatroff strings
	var s []byte = sp[0]
	var q int32 = int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0])
	for int32(s[0]) != 0 && int32(s[0]) != q {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	}
	if int32(s[0]) == q {
		s = s[0+1:]
	}
	d[0] = '\x00'
	sp[0] = s
}

// unquotedread - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:204
func unquotedread(sp [][]byte, d []byte) {
	// read unquoted arguments; this is called only for internal neatroff strings
	var s []byte = sp[0]
	if int32(s[0]) == int32('(') {
		s = s[0+1:]
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	} else if noarch.Not((nreg(int32('C')))[0]) && int32(s[0]) == int32('[') {
		s = s[0+1:]
		for int32(s[0]) != 0 && int32(s[0]) != int32(']') {
			(func() []byte {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		if int32(s[0]) == int32(']') {
			s = s[0+1:]
		}
	} else {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	}
	d[0] = '\x00'
	sp[0] = s
}

// escread - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:234
func escread(s [][]byte, d [][]byte) int32 {
	//
	// * read a glyph or an escape sequence
	// *
	// * This function reads from s either an output troff request
	// * (only the ones emitted by wb.c) or a glyph name and updates
	// * s.  The return value is the name of the troff request (the
	// * argument is copied into d) or zero for glyph names (it is
	// * copied into d).  Returns -1 when the end of s is reached.
	// * Note that to d, a pointer to a static array is assigned.
	//
	var buf []byte = make([]byte, 4096)
	var r []byte
	if noarch.Not((s[0])[0]) {
		return -1
	}
	r = buf
	d[0] = buf
	utf8read(s, r)
	if int32(r[0]) == c_ec {
		utf8read(s, r[0+1:])
		if int32(r[1]) == int32('(') {
			utf8read(s, r)
			utf8read(s, r[0+noarch.Strlen(r):])
		} else if noarch.Not((nreg(int32('C')))[0]) && int32(r[1]) == int32('[') {
			for int32((s[0])[0]) != 0 && int32((s[0])[0]) != int32(']') {
				(func() []byte {
					defer func() {
						r = r[0+1:]
					}()
					return r
				}())[0] = (func() []byte {
					tempVar1 := s[0]
					defer func() {
						s[0] = s[0][1:]
					}()
					return tempVar1
				}())[0]
			}
			r[0] = '\x00'
			if int32((s[0])[0]) == int32(']') {
				s[0] = (s[0])[0+1:]
			}
		} else if noarch.Strchr([]byte("CDfhmsvXx<>\x00"), int32(r[1])) != nil {
			var c int32 = int32(r[1])
			r[0] = '\x00'
			if noarch.Strchr([]byte("*fgkmns\x00"), c) != nil {
				unquotedread(s, r)
			}
			if noarch.Strchr([]byte("bCDhHjlLNoRSvwxXZ?\x00"), c) != nil {
				quotedread(s, r)
			}
			if c == int32('C') {
				return 0
			}
			return c
		}
	} else if int32(r[0]) == 4 {
		utf8read(s, r[0+1:])
	}
	return 0
}

// sstr_bufs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:282
//
// * string streams: provide next()/back() interface for string buffers
// *
// * Functions like charnext() require a next()/back() interface
// * for reading input streams.  In order to provide this interface
// * for string buffers, the following functions can be used:
// *
// *   sstr_push(s);
// *   charnext(c, sstr_next, sstr_back);
// *   sstr_pop();
// *
// * The calls to sstr_push()/sstr_pop() may be nested.
//
// buffer stack
var sstr_bufs [][]byte = make([][]byte, 32)

// sstr_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:283
// numbers of items in sstr_bufs[]
var sstr_n int32

// sstr_s - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:284
// current buffer
var sstr_s []byte

// sstr_push - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:286
func sstr_push(s []byte) {
	sstr_bufs[func() int32 {
		defer func() {
			sstr_n++
		}()
		return sstr_n
	}()] = sstr_s
	sstr_s = s
}

// sstr_pop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:292
func sstr_pop() []byte {
	var ret []byte = sstr_s
	sstr_s = sstr_bufs[func() int32 {
		sstr_n--
		return sstr_n
	}()]
	return ret
}

// sstr_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:299
func sstr_next() int32 {
	if int32(sstr_s[0]) != 0 {
		return int32(uint8((func() []byte {
			defer func() {
				sstr_s = sstr_s[0+1:]
			}()
			return sstr_s
		}())[0]))
	}
	return -1
}

// sstr_back - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/char.c:304
func sstr_back(c int32) {
	sstr_s = c4goPointerArithByteSlice(sstr_s, int(-1))
}

// color - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:8
// color management
type color struct {
	name  []byte
	value int32
}

// colors - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:8
var colors []color = []color{{[]byte("black\x00"), (0<<uint64(16) | 0<<uint64(8) | 0)}, {[]byte("red\x00"), (255<<uint64(16) | 0<<uint64(8) | 0)}, {[]byte("green\x00"), (0<<uint64(16) | 255<<uint64(8) | 0)}, {[]byte("yellow\x00"), (255<<uint64(16) | 255<<uint64(8) | 0)}, {[]byte("blue\x00"), (0<<uint64(16) | 0<<uint64(8) | 255)}, {[]byte("magenta\x00"), (255<<uint64(16) | 0<<uint64(8) | 255)}, {[]byte("cyan\x00"), (0<<uint64(16) | 255<<uint64(8) | 255)}, {[]byte("white\x00"), (255<<uint64(16) | 255<<uint64(8) | 255)}}

// clr_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:23
func clr_str(c int32) []byte {
	// returns a static buffer
	var clr_buf []byte = make([]byte, 32)
	if noarch.Not(c) {
		return []byte("0\x00")
	}
	noarch.Sprintf(clr_buf, []byte("#%02x%02x%02x\x00"), c>>uint64(16)&255, c>>uint64(8)&255, c&255)
	return clr_buf
}

// ccom - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:33
func ccom(s []byte, len_ int32) int32 {
	// read color component
	var digs []byte = []byte("0123456789abcdef\x00")
	var n int32
	var i int32
	for i = 0; i < len_; i++ {
		if noarch.Strchr(digs, tolower(int32(s[i]))) != nil {
			n = n*16 + int32((func() int64 {
				c4go_temp_name := noarch.Strchr(digs, tolower(int32(s[i])))
				return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
			}() - int64(uintptr(unsafe.Pointer(&digs[0])))/int64(1)))
		}
	}
	if len_ == 1 {
		return n * 255 / 15
	}
	return n
}

// clr_get - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:44
func clr_get(s []byte) int32 {
	var c int32 = int32(uint8(s[0]))
	var i int32
	if c == int32('#') && noarch.Strlen(s) == int32(7) {
		return ccom(s[0+1:], 2)<<uint64(16) | ccom(s[0+3:], 2)<<uint64(8) | ccom(s[0+5:], 2)
	}
	if c == int32('#') && noarch.Strlen(s) == int32(4) {
		return ccom(s[0+1:], 1)<<uint64(16) | ccom(s[0+2:], 1)<<uint64(8) | ccom(s[0+3:], 1)
	}
	if c == int32('#') && noarch.Strlen(s) == int32(3) {
		return ccom(s[0+1:], 2)<<uint64(16) | ccom(s[0+1:], 2)<<uint64(8) | ccom(s[0+1:], 2)
	}
	if c == int32('#') && noarch.Strlen(s) == int32(2) {
		return ccom(s[0+1:], 1)<<uint64(16) | ccom(s[0+1:], 1)<<uint64(8) | ccom(s[0+1:], 1)
	}
	if int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit)) != 0 && noarch.Atoi(s) >= 0 && uint32(noarch.Atoi(s)) < 192/24 {
		return colors[noarch.Atoi(s)].value
	}
	for i = 0; uint32(i) < 192/24; i++ {
		if noarch.Not(noarch.Strcmp(colors[i].name, s)) {
			return colors[i].value
		}
	}
	return 0
}

// cp_blkdep - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:7
// copy-mode character interpretation
// input block depth (text in \{ and \})
var cp_blkdep int32

// cp_cpmode - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:8
// disable the interpretation of \w and \E
var cp_cpmode int32

// cp_reqdep - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:9
// the block depth of current request line
var cp_reqdep int32

// cp_noninext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:12
func cp_noninext() int32 {
	// just like cp_next(), but remove c_ni characters
	var c int32 = cp_next()
	for c == 4 {
		c = cp_next()
	}
	return c
}

// cparg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:21
func cparg(d []byte, len_ int32) int32 {
	// return 1 if \*[] includes a space
	var c int32 = cp_noninext()
	var i int32
	if c == int32('(') {
		i += utf8next(d[0+i:], cp_noninext)
		i += utf8next(d[0+i:], cp_noninext)
	} else if noarch.Not((nreg(int32('C')))[0]) && c == int32('[') {
		c = cp_noninext()
		for c >= 0 && c != int32(']') && c != int32(' ') {
			if i+1 < len_ {
				d[func() int32 {
					defer func() {
						i++
					}()
					return i
				}()] = byte(c)
			}
			c = cp_noninext()
		}
		d[i] = '\x00'
		return noarch.BoolToInt(c == int32(' '))
	} else {
		in_back(c)
		utf8next(d, cp_noninext)
	}
	return 0
}

// regid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:44
func regid() int32 {
	var regname []byte = make([]byte, 128)
	cparg(regname, int32(128))
	return map_(regname)
}

// cp_num - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:52
func cp_num() {
	// interpolate \n(xy
	var id int32
	var c int32 = cp_noninext()
	if c != int32('-') && c != int32('+') {
		in_back(c)
	}
	id = regid()
	if c == int32('-') || c == int32('+') {
		num_inc(id, noarch.BoolToInt(c == int32('+')))
	}
	if num_str(id) != nil {
		in_push(num_str(id), nil)
	}
}

// cp_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:66
func cp_str() {
	// interpolate \*(xy
	var reg []byte = make([]byte, 128)
	var args [][]byte = [][]byte{reg, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	var buf []byte
	if cparg(reg, int32(128)) != 0 {
		buf = tr_args(args[0+1:], int32(']'), cp_noninext, in_back)
		cp_noninext()
	}
	if str_get(map_(reg)) != nil {
		in_push(str_get(map_(reg)), func() [][]byte {
			if buf != nil {
				return args
			}
			return nil
		}())
	} else if noarch.Not((nreg(int32('C')))[0]) {
		tr_req(map_(reg), args)
	}
	_ = buf
}

// cp_numfmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:83
func cp_numfmt() {
	// interpolate \g(xy
	in_push(num_getfmt(regid()), nil)
}

// cp_args - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:89
func cp_args(quote int32, escape int32) {
	// interpolate \$*, \$@, and \$^
	var sb sbuf
	var s []byte
	var i int32
	sbuf_init(c4goUnsafeConvert_sbuf(&sb))
	for i = 1; i < in_nargs(); i++ {
		sbuf_append(c4goUnsafeConvert_sbuf(&sb), func() []byte {
			if i > 1 {
				return []byte(" \x00")
			}
			return []byte("\x00")
		}())
		sbuf_append(c4goUnsafeConvert_sbuf(&sb), func() []byte {
			if quote != 0 {
				return []byte("\"\x00")
			}
			return []byte("\x00")
		}())
		s = in_arg(i)
		for s[0] != 0 {
			sbuf_append(c4goUnsafeConvert_sbuf(&sb), func() []byte {
				if escape != 0 && int32(s[0]) == int32('"') {
					return []byte("\"\x00")
				}
				return []byte("\x00")
			}())
			sbuf_add(c4goUnsafeConvert_sbuf(&sb), int32(uint8((func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0])))
		}
		sbuf_append(c4goUnsafeConvert_sbuf(&sb), func() []byte {
			if quote != 0 {
				return []byte("\"\x00")
			}
			return []byte("\x00")
		}())
	}
	in_push(sbuf_buf(c4goUnsafeConvert_sbuf(&sb)), nil)
	sbuf_done(c4goUnsafeConvert_sbuf(&sb))
}

// cp_arg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:110
func cp_arg() {
	// interpolate \$1
	var argname []byte = make([]byte, 128)
	var arg []byte
	var argnum int32
	cparg(argname, int32(128))
	if noarch.Not(noarch.Strcmp([]byte("@\x00"), argname)) {
		cp_args(1, 0)
		return
	}
	if noarch.Not(noarch.Strcmp([]byte("*\x00"), argname)) {
		cp_args(0, 0)
		return
	}
	if noarch.Not(noarch.Strcmp([]byte("^\x00"), argname)) {
		cp_args(1, 1)
		return
	}
	argnum = noarch.Atoi(argname)
	if argnum >= 0 && argnum < 32 {
		arg = in_arg(argnum)
	}
	if arg != nil {
		in_push(arg, nil)
	}
}

// cp_width - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:136
func cp_width() {
	// interpolate \w'xyz'
	var wid []byte = make([]byte, 16)
	noarch.Sprintf(wid, []byte("%d\x00"), ren_wid(cp_next, in_back))
	in_push(wid, nil)
}

// cp_numdef - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:144
func cp_numdef() {
	// define a register as \R'xyz expr'
	var arg []byte = quotednext(cp_noninext, in_back)
	var s []byte = arg
	for int32(s[0]) != 0 && int32(s[0]) != int32(' ') {
		s = s[0+1:]
	}
	if noarch.Not(s[0]) {
		_ = arg
		return
	}
	(func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0] = '\x00'
	num_set(map_(arg), eval_re(s, (nreg(map_(arg)))[0], int32('u')))
	_ = arg
}

// cp_cond - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:160
func cp_cond() {
	// conditional interpolation as \?'cond@expr1@expr2@'
	var delim []byte = make([]byte, 32)
	var cs []byte = make([]byte, 32)
	var r []byte
	var s []byte
	var s1 []byte
	var s2 []byte
	var n int32
	var arg []byte = quotednext(cp_noninext, in_back)
	s = arg
	n = eval_up((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('\x00'))
	if charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], delim) < 0 {
		_ = arg
		return
	}
	if noarch.Not(noarch.Strcmp(delim, []byte("\\&\x00"))) && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], delim) < 0 {
		_ = arg
		return
	}
	s1 = s
	r = s
	for charread_delim((*[1000000][]byte)(unsafe.Pointer(&s))[:], cs, delim) >= 0 {
		r = s
	}
	r[0] = '\x00'
	s2 = s
	r = s
	for charread_delim((*[1000000][]byte)(unsafe.Pointer(&s))[:], cs, delim) >= 0 {
		r = s
	}
	r[0] = '\x00'
	in_push(func() []byte {
		if n > 0 {
			return s1
		}
		return s2
	}(), nil)
	_ = arg
}

// cp_raw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:191
func cp_raw() int32 {
	var c int32
	if in_top() >= 0 {
		return in_next()
	}
	for {
		c = in_next()
		if !(c == 4) {
			break
		}
	}
	if c == c_ec {
		for {
			c = in_next()
			if !(c == 4) {
				break
			}
		}
		if c == int32('\n') {
			return cp_raw()
		}
		if c == int32('.') {
			return int32('.')
		}
		if c == int32('\\') {
			in_back(int32('\\'))
			return 4
		}
		if c == int32('t') {
			in_back(int32('\t'))
			return 4
		}
		if c == int32('a') {
			in_back(int32('\x01'))
			return 4
		}
		if c == int32('}') && noarch.Not(cp_cpmode) {
			// replace \{ and \} with a space if not in copy mode
			cp_blkdep--
			return int32(' ')
		}
		if c == int32('{') && noarch.Not(cp_cpmode) {
			cp_blkdep++
			return int32(' ')
		}
		in_back(c)
		return c_ec
	}
	return c
}

// cp_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:234
func cp_next() int32 {
	var c int32
	if in_top() >= 0 {
		return in_next()
	}
	c = cp_raw()
	if c == c_ec {
		c = cp_raw()
		if c == int32('E') && noarch.Not(cp_cpmode) {
			c = cp_next()
		}
		if c == int32('"') {
			for c >= 0 && c != int32('\n') {
				c = cp_raw()
			}
		} else if c == int32('w') && noarch.Not(cp_cpmode) {
			cp_width()
			c = cp_next()
		} else if c == int32('n') {
			cp_num()
			c = cp_next()
		} else if c == int32('*') {
			cp_str()
			c = cp_next()
		} else if c == int32('g') {
			cp_numfmt()
			c = cp_next()
		} else if c == int32('$') {
			cp_arg()
			c = cp_next()
		} else if c == int32('?') {
			cp_cond()
			c = cp_next()
		} else if c == int32('R') && noarch.Not(cp_cpmode) {
			cp_numdef()
			c = cp_next()
		} else {
			in_back(c)
			c = c_ec
		}
	}
	return c
}

// cp_blk - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:276
func cp_blk(skip int32) {
	if skip != 0 {
		var c int32 = cp_raw()
		for c >= 0 && (c != int32('\n') || cp_blkdep > cp_reqdep) {
			c = cp_raw()
		}
	} else {
		var c int32 = cp_next()
		for c == int32(' ') {
			c = cp_next()
		}
		if c != int32(' ') {
			// push back if the space is not inserted due to \{ and \}
			in_back(c)
		}
	}
}

// cp_copymode - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:292
func cp_copymode(mode int32) {
	cp_cpmode = mode
}

// cp_reqbeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/cp.c:298
func cp_reqbeg() {
	// beginning of a request; save current cp_blkdep
	cp_reqdep = cp_blkdep
}

// dev_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:8
// output device
// device directory
var dev_dir []byte = make([]byte, 1024)

// dev_dev - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:9
// output device name
var dev_dev []byte = make([]byte, 1024)

// dev_res - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:10
// device resolution
var dev_res int32

// dev_uwid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:11
// device unitwidth
var dev_uwid int32

// dev_hor - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:12
// minimum horizontal movement
var dev_hor int32

// dev_ver - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:13
// minimum vertical movement
var dev_ver int32

// fn_name - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:16
// mounted fonts
// font names
var fn_name [][]byte = make([][]byte, 32)

// fn_font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:17
// font structs
var fn_font [][]font = make([][]font, 32)

// fn_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:18
// number of device fonts
var fn_n int32

// fspecial_fn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:21
// .fspecial request
// .fspecial first arguments
var fspecial_fn [][]byte = make([][]byte, 32)

// fspecial_sp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:22
// .fspecial special fonts
var fspecial_sp [][]byte = make([][]byte, 32)

// fspecial_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:23
// number of fonts in fspecial_sp[]
var fspecial_n int32

// skipline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:25
func skipline(filp *noarch.File) {
	var c int32
	for {
		c = noarch.Fgetc(filp)
		if !(c != int32('\n') && c != -1) {
			break
		}
	}
}

// dev_prologue - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:33
func dev_prologue() {
	out([]byte("x T %s\n\x00"), dev_dev)
	out([]byte("x res %d %d %d\n\x00"), dev_res, dev_hor, dev_ver)
	out([]byte("x init\n\x00"))
}

// dev_position - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:41
func dev_position(id []byte) int32 {
	// find a position for the given font
	var i int32
	{
		// already mounted
		for i = 1; i < 32; i++ {
			if noarch.Not(noarch.Strcmp(fn_name[i], id)) {
				return i
			}
		}
	}
	{
		// the first empty position
		for i = 1; i < 32; i++ {
			if fn_font[i] == nil {
				return i
			}
		}
	}
	// no room left
	return 0
}

// dev_mnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:53
func dev_mnt(pos int32, id []byte, name []byte) int32 {
	var path []byte = make([]byte, 1024)
	var fn []font
	if pos >= 32 {
		return -1
	}
	if noarch.Strchr(name, int32('/')) != nil {
		noarch.Snprintf(path, int32(1024), []byte("%s\x00"), name)
	} else {
		noarch.Snprintf(path, int32(1024), []byte("%s/dev%s/%s\x00"), dev_dir, dev_dev, name)
	}
	fn = font_open(path)
	if fn == nil {
		return -1
	}
	if pos < 0 {
		pos = dev_position(id)
	}
	if fn_font[pos] != nil {
		font_close(fn_font[pos])
	}
	if (int64(uintptr(unsafe.Pointer(&fn_name[pos])))/int64(1) - int64(uintptr(unsafe.Pointer(&name[0])))/int64(1)) != 0 {
		// ignore if fn_name[pos] is passed
		noarch.Snprintf(fn_name[pos], int32(32), []byte("%s\x00"), id)
	}
	fn_font[pos] = fn
	out([]byte("x font %d %s\n\x00"), pos, name)
	return pos
}

// dev_open - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:77
func dev_open(dir []byte, dev []byte) int32 {
	var path []byte = make([]byte, 1024)
	var tok []byte = make([]byte, 128)
	var i int32
	var desc *noarch.File
	noarch.Snprintf(dev_dir, int32(1024), []byte("%s\x00"), dir)
	noarch.Snprintf(dev_dev, int32(1024), []byte("%s\x00"), dev)
	noarch.Snprintf(path, int32(1024), []byte("%s/dev%s/DESC\x00"), dir, dev)
	desc = noarch.Fopen(path, []byte("r\x00"))
	if desc == nil {
		return 1
	}
	for noarch.Fscanf(desc, []byte("%128s\x00"), tok) == 1 {
		if int32(tok[0]) == int32('#') {
			skipline(desc)
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("fonts\x00"), tok)) {
			noarch.Fscanf(desc, []byte("%d\x00"), c4goUnsafeConvert_int32(&fn_n))
			for i = 0; i < fn_n; i++ {
				noarch.Fscanf(desc, []byte("%s\x00"), fn_name[i+1])
			}
			fn_n++
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("sizes\x00"), tok)) {
			for noarch.Fscanf(desc, []byte("%128s\x00"), tok) == 1 {
				if noarch.Not(noarch.Strcmp([]byte("0\x00"), tok)) {
					break
				}
			}
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("res\x00"), tok)) {
			noarch.Fscanf(desc, []byte("%d\x00"), c4goUnsafeConvert_int32(&dev_res))
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("unitwidth\x00"), tok)) {
			noarch.Fscanf(desc, []byte("%d\x00"), c4goUnsafeConvert_int32(&dev_uwid))
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("hor\x00"), tok)) {
			noarch.Fscanf(desc, []byte("%d\x00"), c4goUnsafeConvert_int32(&dev_hor))
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("ver\x00"), tok)) {
			noarch.Fscanf(desc, []byte("%d\x00"), c4goUnsafeConvert_int32(&dev_ver))
			continue
		}
		if noarch.Not(noarch.Strcmp([]byte("charset\x00"), tok)) {
			break
		}
		skipline(desc)
	}
	noarch.Fclose(desc)
	dev_prologue()
	for i = 0; i < fn_n; i++ {
		if (fn_name[i])[0] != 0 {
			dev_mnt(i, fn_name[i], fn_name[i])
		}
	}
	return 0
}

// dev_epilogue - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:135
func dev_epilogue() {
	out([]byte("x trailer\n\x00"))
	out([]byte("x stop\n\x00"))
}

// dev_close - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:141
func dev_close() {
	var i int32
	dev_epilogue()
	for i = 0; i < 32; i++ {
		if fn_font[i] != nil {
			font_close(fn_font[i])
		}
		fn_font[i] = nil
	}
}

// dev_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:154
func dev_find(c []byte, fn int32, byid int32) []glyph {
	// glyph handling functions
	var find func([]font, []byte) []glyph
	var g []glyph
	var i int32
	if byid != 0 {
		find = font_glyph
	} else {
		find = font_find
	}
	if (func() []glyph {
		g = find(fn_font[fn], c)
		return g
	}()) != nil {
		return g
	}
	for i = 0; i < fspecial_n; i++ {
		if dev_pos(fspecial_fn[i]) == fn && dev_pos(fspecial_sp[i]) >= 0 {
			if (func() []glyph {
				g = find(dev_font(dev_pos(fspecial_sp[i])), c)
				return g
			}()) != nil {
				return g
			}
		}
	}
	for i = 0; i < 32; i++ {
		if fn_font[i] != nil && font_special(fn_font[i]) != 0 {
			if (func() []glyph {
				g = find(fn_font[i], c)
				return g
			}()) != nil {
				return g
			}
		}
	}
	return nil
}

// dev_glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:173
func dev_glyph(c []byte, fn int32) []glyph {
	if (int32(c[0]) == c_ec || int32(c[0]) == 4) && int32(c[1]) == c_ec {
		c = c[0+1:]
	}
	if int32(c[0]) == c_ec && int32(c[1]) == int32('(') {
		c = c[0+2:]
	}
	c = cmap_map(c)
	if noarch.Not(strncmp([]byte("GID=\x00"), c, 4)) {
		return dev_find(c[0+4:], fn, 1)
	}
	return dev_find(c, fn, 0)
}

// dev_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:186
func dev_pos(id []byte) int32 {
	// return the mounted position of a font
	var i int32
	if int32(((__ctype_b_loc())[0])[int32(id[0])])&int32(uint16(noarch.ISdigit)) != 0 {
		var num int32 = noarch.Atoi(id)
		if num < 0 || num >= 32 || fn_font[num] == nil {
			errmsg([]byte("neatroff: bad font position %s\n\x00"), id)
			return -1
		}
		return num
	}
	for i = 1; i < 32; i++ {
		if noarch.Not(noarch.Strcmp(fn_name[i], id)) {
			return i
		}
	}
	if noarch.Not(noarch.Strcmp(fn_name[0], id)) {
		return 0
	}
	return dev_mnt(0, id, id)
}

// dev_fontpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:206
func dev_fontpos(fn []font) int32 {
	// return the mounted position of a font struct
	var i int32
	// Warning (*ast.IfStmt):  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:210 :cannot transpileToStmt : cannot transpileIfStmt. cannot transpile for condition. cannot transpileToExpr. err = cannot transpile BinaryOperator with type 'int' : result type = {PointerOperation_unknown05}. Error: operator is `==`. {'struct font *' == 'struct font *'}. for base type: `struct font`. PntCmpPnt:SubTwoPnts:GetPointerAddress:sizeof:0. not valid sizeof `struct font *`: 0
	for i = 0; i < 32; i++ {
		// Warning (*ast.BinaryOperator):  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:210 :cannot transpile BinaryOperator with type 'int' : result type = {PointerOperation_unknown05}. Error: operator is `==`. {'struct font *' == 'struct font *'}. for base type: `struct font`. PntCmpPnt:SubTwoPnts:GetPointerAddress:sizeof:0. not valid sizeof `struct font *`: 0
	}
	return 0
}

// dev_font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:216
func dev_font(pos int32) []font {
	// return the font struct at pos
	if pos >= 0 && pos < 32 {
		return fn_font[pos]
	}
	return nil
}

// tr_fspecial - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:221
func tr_fspecial(args [][]byte) {
	var fn []byte = args[1]
	var i int32
	if fn == nil {
		fspecial_n = 0
		return
	}
	for i = 2; i < 32; i++ {
		if args[i] != nil && uint32(fspecial_n) < 1024/32 {
			noarch.Snprintf(fspecial_fn[fspecial_n], int32(32), []byte("%s\x00"), fn)
			noarch.Snprintf(fspecial_sp[fspecial_n], int32(32), []byte("%s\x00"), args[i])
			fspecial_n++
		}
	}
}

// dict - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:8
type dict struct {
	map_     []iset
	key      [][]byte
	val      []int32
	size     int32
	n        int32
	notfound int32
	hashlen  int32
	dupkeys  int32
}

// dict_extend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:19
func dict_extend(d []dict, size int32) {
	// the value returned for missing keys
	// the number of characters used for hashing
	// duplicate keys if set
	d[0].key = mextend(d[0].key, d[0].size, size, int32(8)).([][]byte)
	d[0].val = mextend(d[0].val, d[0].size, size, int32(4)).([]int32)
	d[0].size = size
}

// dict_make - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:33
func dict_make(notfound int32, dupkeys int32, hashlen int32) []dict {
	//
	// * initialise a dictionary
	// *
	// * notfound: the value returned for missing keys.
	// * dupkeys: if nonzero, store a copy of keys inserted via dict_put().
	// * hashlen: the number of characters used for hashing
	//
	var d []dict = xmalloc(int32(72)).([]dict)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&d[0]))) / int64(1))))[:], byte(0), 72)
	d[0].n = 1
	if hashlen != 0 {
		d[0].hashlen = hashlen
	} else {
		d[0].hashlen = 32
	}
	d[0].dupkeys = dupkeys
	d[0].notfound = notfound
	d[0].map_ = iset_make()
	dict_extend(d, 1<<uint64(10))
	return d
}

// dict_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:46
func dict_free(d []dict) {
	var i int32
	if d[0].dupkeys != 0 {
		for i = 0; i < d[0].size; i++ {
			_ = d[0].key[i]
		}
	}
	_ = d[0].val
	_ = d[0].key
	iset_free(d[0].map_)
	_ = d
}

// dict_hash - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:58
func dict_hash(d []dict, key []byte) int32 {
	var hash uint32 = uint32(uint8((func() []byte {
		defer func() {
			key = key[0+1:]
		}()
		return key
	}())[0]))
	var i int32 = d[0].hashlen
	for func() int32 {
		i--
		return i
	}() > 0 && int32(key[0]) != 0 {
		hash = hash<<uint64(5) + hash + uint32(uint8((func() []byte {
			defer func() {
				key = key[0+1:]
			}()
			return key
		}())[0]))
	}
	return int32(hash & 1023)
}

// dict_put - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:67
func dict_put(d []dict, key []byte, val int32) {
	var idx int32
	if d[0].n >= d[0].size {
		dict_extend(d, d[0].n+1<<uint64(10))
	}
	if d[0].dupkeys != 0 {
		var len_ int32 = noarch.Strlen(key) + int32(1)
		var dup []byte = xmalloc(len_).([]byte)
		memcpy(dup, key, uint32(len_))
		key = dup
	}
	idx = func() int32 {
		tempVar1 := &d[0].n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}()
	d[0].key[idx] = key
	d[0].val[idx] = val
	iset_put(d[0].map_, dict_hash(d, key), idx)
}

// dict_idx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:85
func dict_idx(d []dict, key []byte) int32 {
	// return the index of key in d
	var h int32 = dict_hash(d, key)
	var b []int32 = iset_get(d[0].map_, h)
	var r []int32 = b[0+iset_len(d[0].map_, h):]
	for b != nil && (func() int64 {
		c4go_temp_name := func() []int32 {
			r = c4goPointerArithInt32Slice(r, int(-1))
			return r
		}()
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}()-int64(uintptr(unsafe.Pointer(&b[0])))/int64(4)) >= 0 {
		if noarch.Not(noarch.Strcmp(d[0].key[r[0]], key)) {
			return r[0]
		}
	}
	return -1
}

// dict_key - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:96
func dict_key(d []dict, idx int32) []byte {
	return d[0].key[idx]
}

// dict_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:101
func dict_val(d []dict, idx int32) int32 {
	return d[0].val[idx]
}

// dict_get - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:106
func dict_get(d []dict, key []byte) int32 {
	var idx int32 = dict_idx(d, key)
	if idx >= 0 {
		return d[0].val[idx]
	}
	return d[0].notfound
}

// dict_prefix - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dict.c:113
func dict_prefix(d []dict, key []byte, pos []int32) int32 {
	// match a prefix of key; in the first call, *pos should be -1
	var r []int32 = iset_get(d[0].map_, dict_hash(d, key))
	for r != nil && r[func() int32 {
		tempVar1 := &pos[0]
		*tempVar1++
		return *tempVar1
	}()] >= 0 {
		var idx int32 = r[pos[0]]
		var plen int32 = noarch.Strlen(d[0].key[idx])
		if noarch.Not(strncmp(d[0].key[idx], key, uint32(plen))) {
			return d[0].val[idx]
		}
	}
	return d[0].notfound
}

// dir_do - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:7
// output text direction
// enable text direction processing
var dir_do int32

// dbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:9
// text in n_td direction
var dbuf []byte

// dbuf_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:10
// dbuf[] size and length
var dbuf_sz int32

// dbuf_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:10
var dbuf_n int32

// rbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:11
// text in (1 - n_td) direction
var rbuf []byte

// rbuf_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:12
// rbuf[] size and length
var rbuf_sz int32

// rbuf_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:12
var rbuf_n int32

// dir_cd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:13
// current direction
var dir_cd int32

// dir_copy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:16
func dir_copy(d [][]byte, d_n []int32, d_sz []int32, s []byte, s_n int32, dir int32) {
	for d_n[0]+s_n+1 > d_sz[0] {
		// append s to the start (dir == 0) or end (dir == 1) of d
		var sz int32 = func() int32 {
			if d_sz[0] != 0 {
				return d_sz[0] * 2
			}
			return 512
		}()
		var n []byte = make([]byte, uint32(sz+1))
		if d_sz[0] != 0 {
			memcpy(func() []byte {
				if dir != 0 {
					return n[0+d_sz[0]:]
				}
				return n
			}(), d[0], uint32(d_sz[0]))
		}
		_ = d[0]
		d_sz[0] = sz
		d[0] = n
	}
	if dir > 0 {
		memcpy(c4goPointerArithByteSlice(c4goPointerArithByteSlice((d[0])[0+d_sz[0]:], int(-d_n[0])), int(-s_n)), s, uint32(s_n))
	} else {
		memcpy((d[0])[0+d_n[0]:], s, uint32(s_n))
	}
	d_n[0] += s_n
}

// dir_flush - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:35
func dir_flush() {
	// copy rbuf (the text in reverse direction) to dbuf
	var s []byte = c4goPointerArithByteSlice(rbuf, int(func() int32 {
		if (nreg(map_([]byte(".td\x00"))))[0] > 0 {
			return 0
		}
		return rbuf_sz - rbuf_n
	}()))
	dir_copy((*[1000000][]byte)(unsafe.Pointer(&dbuf))[:], c4goUnsafeConvert_int32(&dbuf_n), c4goUnsafeConvert_int32(&dbuf_sz), s, rbuf_n, (nreg(map_([]byte(".td\x00"))))[0])
	rbuf_n = 0
}

// dir_append - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:43
func dir_append(s []byte) {
	// append s to dbuf or rbuf based on the current text direction
	var dir int32 = noarch.BoolToInt(dir_cd > 0)
	if dir == (nreg(map_([]byte(".td\x00"))))[0] && rbuf_n != 0 {
		dir_flush()
	}
	if dir == (nreg(map_([]byte(".td\x00"))))[0] {
		dir_copy((*[1000000][]byte)(unsafe.Pointer(&dbuf))[:], c4goUnsafeConvert_int32(&dbuf_n), c4goUnsafeConvert_int32(&dbuf_sz), s, noarch.Strlen(s), dir)
	} else {
		dir_copy((*[1000000][]byte)(unsafe.Pointer(&rbuf))[:], c4goUnsafeConvert_int32(&rbuf_n), c4goUnsafeConvert_int32(&rbuf_sz), s, noarch.Strlen(s), dir)
	}
}

// setfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:54
func setfont(f int32) {
	var cmd []byte = make([]byte, 32)
	noarch.Sprintf(cmd, []byte("%cf(%02d\x00"), c_ec, f)
	if f >= 0 {
		dir_append(cmd)
	}
}

// setsize - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:62
func setsize(s int32) {
	var cmd []byte = make([]byte, 32)
	noarch.Sprintf(cmd, func() []byte {
		if s <= 99 {
			return []byte("%cs(%02d\x00")
		}
		return []byte("%cs[%d]\x00")
	}(), c_ec, s)
	if s >= 0 {
		dir_append(cmd)
	}
}

// setcolor - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:70
func setcolor(m int32) {
	var cmd []byte = make([]byte, 32)
	noarch.Sprintf(cmd, []byte("%cm[%s]\x00"), c_ec, clr_str(m))
	if m >= 0 {
		dir_append(cmd)
	}
}

// dir_fix - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:78
func dir_fix(sbuf_c4go_postfix []sbuf, src []byte) {
	var cmd []byte = make([]byte, 1024)
	var prev_s []byte = src
	var r []byte
	var c []byte
	var f int32 = -1
	var s int32 = -1
	var m int32 = -1
	var t int32
	var n int32
	dir_cd = (nreg(map_([]byte(".td\x00"))))[0]
	for (func() int32 {
		t = escread((*[1000000][]byte)(unsafe.Pointer(&src))[:], (*[1000000][]byte)(unsafe.Pointer(&c))[:])
		return t
	}()) >= 0 {
		cmd[0] = '\x00'
		switch t {
		case 0:
			fallthrough
		case 'D':
			fallthrough
		case 'h':
			fallthrough
		case 'v':
			fallthrough
		case 'x':
			memcpy(cmd, prev_s, uint32(int32((int64(uintptr(unsafe.Pointer(&src[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&prev_s[0])))/int64(1)))))
			cmd[int32((int64(uintptr(unsafe.Pointer(&src[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&prev_s[0])))/int64(1)))] = '\x00'
			dir_append(cmd)
		case 'f':
			n = noarch.Atoi(c)
			if f != n {
				setfont(f)
				f = n
				setfont(f)
			}
		case 'm':
			n = clr_get(c)
			if m != n {
				setcolor(m)
				m = n
				setcolor(m)
			}
		case 's':
			n = noarch.Atoi(c)
			if s != n {
				setsize(s)
				s = n
				setsize(s)
			}
		case 'X':
			noarch.Sprintf(cmd, []byte("%c%c\x03%s\x03\x00"), c_ec, t, c)
			dir_append(cmd)
		case '<':
			setcolor(m)
			setfont(f)
			setsize(s)
			dir_cd = 1
			setsize(s)
			setfont(f)
			setcolor(m)
		case '>':
			setcolor(m)
			setfont(f)
			setsize(s)
			dir_cd = 0
			setsize(s)
			setfont(f)
			setcolor(m)
			break
		}
		prev_s = src
	}
	setcolor(m)
	setfont(f)
	setsize(s)
	dir_flush()
	if (nreg(map_([]byte(".td\x00"))))[0] > 0 {
		r = c4goPointerArithByteSlice(dbuf[0+dbuf_sz:], int(-dbuf_n))
	} else {
		r = dbuf
	}
	r[dbuf_n] = '\x00'
	dbuf_n = 0
	sbuf_append(sbuf_c4go_postfix, r)
}

// dir_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dir.c:157
func dir_done() {
	_ = rbuf
	_ = dbuf
}

// cwid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:7
func cwid(c []byte) int32 {
	// helper for drawing commands in ren.c
	var wb_c4go_postfix wb
	var w int32
	wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	wb_putexpand(c4goUnsafeConvert_wb(&wb_c4go_postfix), c)
	w = wb_wid(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	return w
}

// hchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:18
func hchar(c []byte) int32 {
	if int32(c[0]) == c_ec {
		return noarch.BoolToInt(int32(c[1]) == int32('_') || int32(c[1]) == int32('-'))
	}
	if noarch.Not(c[1]) {
		return noarch.BoolToInt(int32(c[0]) == int32('_'))
	}
	return noarch.BoolToInt(int32(c[0]) == int32('r') && int32(c[1]) == int32('u') || int32(c[0]) == int32('u') && int32(c[1]) == int32('l') || int32(c[0]) == int32('r') && int32(c[1]) == int32('n'))
}

// vchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:28
func vchar(c []byte) int32 {
	if noarch.Not(c[1]) {
		return noarch.BoolToInt(int32(c[0]) == int32('_'))
	}
	return noarch.BoolToInt(int32(c[0]) == int32('b') && int32(c[1]) == int32('v') || int32(c[0]) == int32('b') && int32(c[1]) == int32('r'))
}

// ren_hline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:35
func ren_hline(wb_c4go_postfix []wb, l int32, c []byte) {
	var w int32
	var n int32
	var i int32
	var rem int32
	w = cwid(c)
	if l < 0 {
		// negative length; moving backwards
		wb_hmov(wb_c4go_postfix, l)
		l = -l
	}
	if w != 0 {
		n = l / w
	} else {
		n = 0
	}
	if w != 0 {
		rem = l % w
	} else {
		rem = l
	}
	if l < w {
		// length less than character width
		n = 1
		rem = 0
		wb_hmov(wb_c4go_postfix, -(w-l)/2)
	}
	if rem != 0 {
		if hchar(c) != 0 {
			// the initial gap
			wb_putexpand(wb_c4go_postfix, c)
			wb_hmov(wb_c4go_postfix, rem-w)
		} else {
			wb_hmov(wb_c4go_postfix, rem)
		}
	}
	for i = 0; i < n; i++ {
		wb_putexpand(wb_c4go_postfix, c)
	}
	if l < w {
		// moving back
		wb_hmov(wb_c4go_postfix, -(w-l+1)/2)
	}
}

// ren_vline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:68
func ren_vline(wb_c4go_postfix []wb, l int32, c []byte) {
	var w int32
	var n int32
	var i int32
	var rem int32
	var hw int32
	var neg int32
	neg = noarch.BoolToInt(l < 0)
	// character height
	w = (nreg(int32('s')))[0] * dev_res / 72
	// character width
	hw = cwid(c)
	if l < 0 {
		// negative length; moving backwards
		wb_vmov(wb_c4go_postfix, l)
		l = -l
	}
	if w != 0 {
		n = l / w
	} else {
		n = 0
	}
	if w != 0 {
		rem = l % w
	} else {
		rem = l
	}
	if l < w {
		// length less than character width
		n = 1
		rem = 0
		wb_vmov(wb_c4go_postfix, -w+l/2)
	}
	if rem != 0 {
		if vchar(c) != 0 {
			// the initial gap
			wb_vmov(wb_c4go_postfix, w)
			wb_putexpand(wb_c4go_postfix, c)
			wb_hmov(wb_c4go_postfix, -hw)
			wb_vmov(wb_c4go_postfix, rem-w)
		} else {
			wb_vmov(wb_c4go_postfix, rem)
		}
	}
	for i = 0; i < n; i++ {
		wb_vmov(wb_c4go_postfix, w)
		wb_putexpand(wb_c4go_postfix, c)
		wb_hmov(wb_c4go_postfix, -hw)
	}
	if l < w {
		// moving back
		wb_vmov(wb_c4go_postfix, l/2)
	}
	if neg != 0 {
		wb_vmov(wb_c4go_postfix, -l)
	}
	wb_hmov(wb_c4go_postfix, hw)
}

// ren_hlcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:111
func ren_hlcmd(wb_c4go_postfix []wb, arg []byte) {
	var lc []byte = make([]byte, 32)
	var l int32 = eval_up((*[1000000][]byte)(unsafe.Pointer(&arg))[:], int32('m'))
	if int32(arg[0]) == c_ec && int32(arg[1]) == int32('&') {
		// \& can be used as a separator
		arg = arg[0+2:]
	}
	if noarch.Not(arg[0]) || charread((*[1000000][]byte)(unsafe.Pointer(&arg))[:], lc) < 0 {
		noarch.Strcpy(lc, []byte("ru\x00"))
	}
	if l != 0 {
		ren_hline(wb_c4go_postfix, l, lc)
	}
}

// ren_vlcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:123
func ren_vlcmd(wb_c4go_postfix []wb, arg []byte) {
	var lc []byte = make([]byte, 32)
	var l int32 = eval_up((*[1000000][]byte)(unsafe.Pointer(&arg))[:], int32('v'))
	if int32(arg[0]) == c_ec && int32(arg[1]) == int32('&') {
		// \& can be used as a separator
		arg = arg[0+2:]
	}
	if noarch.Not(arg[0]) || charread((*[1000000][]byte)(unsafe.Pointer(&arg))[:], lc) < 0 {
		noarch.Strcpy(lc, []byte("br\x00"))
	}
	if l != 0 {
		ren_vline(wb_c4go_postfix, l, lc)
	}
}

// tok_num - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:135
func tok_num(s [][]byte, scale int32) int32 {
	for int32((s[0])[0]) == int32(' ') || int32((s[0])[0]) == int32('\t') {
		s[0] = (s[0])[0+1:]
	}
	return eval_up(s, scale)
}

// tok_numpt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:142
func tok_numpt(s [][]byte, scale int32, i []int32) int32 {
	var o []byte
	for int32((s[0])[0]) == int32(' ') || int32((s[0])[0]) == int32('\t') {
		s[0] = (s[0])[0+1:]
	}
	o = s[0]
	i[0] = eval_up(s, scale)
	if (int64(uintptr(unsafe.Pointer(&o[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)) == 0 {
		return 1
	}
	return 0
}

// ren_dcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:152
func ren_dcmd(wb_c4go_postfix []wb, s []byte) {
	var h1 int32
	var h2 int32
	var v1 int32
	var v2 int32
	var c int32 = int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0])
	switch tolower(c) {
	case 'l':
		h1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'))
		v1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'))
		wb_drawl(wb_c4go_postfix, c, h1, v1)
	case 'c':
		h1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'))
		wb_drawc(wb_c4go_postfix, c, h1)
	case 'e':
		h1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'))
		v1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'))
		wb_drawe(wb_c4go_postfix, c, h1, v1)
	case 'a':
		h1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'))
		v1 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'))
		h2 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'))
		v2 = tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'))
		wb_drawa(wb_c4go_postfix, c, h1, v1, h2, v2)
	case '~':
		fallthrough
	case 'p':
		wb_drawxbeg(wb_c4go_postfix, c)
		for s[0] != 0 {
			if tok_numpt((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'), c4goUnsafeConvert_int32(&h1)) != 0 || tok_numpt((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'), c4goUnsafeConvert_int32(&v1)) != 0 {
				var tok []byte = make([]byte, 64)
				var i int32
				for uint32(i) < 64-1 && int32(s[0]) != 0 && int32(s[0]) != int32(' ') {
					tok[func() int32 {
						defer func() {
							i++
						}()
						return i
					}()] = (func() []byte {
						defer func() {
							s = s[0+1:]
						}()
						return s
					}())[0]
				}
				tok[i] = '\x00'
				wb_drawxcmd(wb_c4go_postfix, tok)
			} else {
				wb_drawxdot(wb_c4go_postfix, h1, v1)
			}
		}
		wb_drawxend(wb_c4go_postfix)
		break
	}
}

// ren_bcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:198
func ren_bcmd(wb_c4go_postfix []wb, arg []byte) {
	var wb2 wb
	var n int32
	var w int32
	var c int32
	var center int32
	// using ren_char()'s interface
	sstr_push(arg)
	wb_init(c4goUnsafeConvert_wb(&wb2))
	c = sstr_next()
	for c >= 0 {
		sstr_back(c)
		ren_char(c4goUnsafeConvert_wb(&wb2), sstr_next, sstr_back)
		if wb_wid(c4goUnsafeConvert_wb(&wb2)) > w {
			w = wb_wid(c4goUnsafeConvert_wb(&wb2))
		}
		wb_hmov(c4goUnsafeConvert_wb(&wb2), -wb_wid(c4goUnsafeConvert_wb(&wb2)))
		wb_vmov(c4goUnsafeConvert_wb(&wb2), (nreg(int32('s')))[0]*dev_res/72)
		n++
		c = sstr_next()
	}
	sstr_pop()
	center = -(n*((nreg(int32('s')))[0]*dev_res/72) + (nreg(int32('s')))[0]*dev_res/72) / 2
	wb_vmov(wb_c4go_postfix, center+(nreg(int32('s')))[0]*dev_res/72)
	wb_cat(wb_c4go_postfix, c4goUnsafeConvert_wb(&wb2))
	wb_done(c4goUnsafeConvert_wb(&wb2))
	wb_vmov(wb_c4go_postfix, center)
	wb_hmov(wb_c4go_postfix, w)
}

// ren_ocmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:225
func ren_ocmd(wb_c4go_postfix []wb, arg []byte) {
	var wb2 wb
	var wb3 wb
	var w int32
	var wc int32
	var c int32
	// using ren_char()'s interface
	sstr_push(arg)
	wb_init(c4goUnsafeConvert_wb(&wb2))
	wb_init(c4goUnsafeConvert_wb(&wb3))
	c = sstr_next()
	for c >= 0 {
		sstr_back(c)
		ren_char(c4goUnsafeConvert_wb(&wb3), sstr_next, sstr_back)
		wc = wb_wid(c4goUnsafeConvert_wb(&wb3))
		if wc > w {
			w = wc
		}
		wb_hmov(c4goUnsafeConvert_wb(&wb2), -wc/2)
		wb_cat(c4goUnsafeConvert_wb(&wb2), c4goUnsafeConvert_wb(&wb3))
		wb_hmov(c4goUnsafeConvert_wb(&wb2), -wc/2)
		c = sstr_next()
	}
	sstr_pop()
	wb_hmov(wb_c4go_postfix, w/2)
	wb_cat(wb_c4go_postfix, c4goUnsafeConvert_wb(&wb2))
	wb_hmov(wb_c4go_postfix, w/2)
	wb_done(c4goUnsafeConvert_wb(&wb3))
	wb_done(c4goUnsafeConvert_wb(&wb2))
}

// ren_zcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/draw.c:253
func ren_zcmd(wb_c4go_postfix []wb, arg []byte) {
	var h int32
	var v int32
	var c int32
	h = wb_hpos(wb_c4go_postfix)
	v = wb_vpos(wb_c4go_postfix)
	sstr_push(arg)
	for (func() int32 {
		c = sstr_next()
		return c
	}()) >= 0 {
		sstr_back(c)
		ren_char(wb_c4go_postfix, sstr_next, sstr_back)
	}
	sstr_pop()
	wb_hmov(wb_c4go_postfix, h-wb_hpos(wb_c4go_postfix))
	wb_vmov(wb_c4go_postfix, v-wb_vpos(wb_c4go_postfix))
}

// defunit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:9
// evaluation of integer expressions
// default scale indicator
var defunit int32

// abspos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:10
// absolute position like |1i
var abspos int32

// readunit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:12
func readunit(c int32, n int32) int32 {
	switch c {
	case 'i':
		return n * dev_res
	case 'c':
		return n * dev_res * 50 / 127
	case 'p':
		return n * dev_res / 72
	case 'P':
		return n * dev_res / 6
	case 'v':
		return n * (nreg(int32('v')))[0]
	case 'm':
		return n * (nreg(int32('s')))[0] * dev_res / 72
	case 'n':
		return n * (nreg(int32('s')))[0] * dev_res / 144
	case 'u':
		return n
	}
	return n
}

// evalnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:35
func evalnum(_s [][]byte) int32 {
	var s []byte = _s[0]
	// the result
	var n int32
	// n should be divided by mag
	var mag int32
	for int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISdigit)) != 0 || int32(s[0]) == int32('.') {
		if int32(s[0]) == int32('.') {
			mag = 1
			s = s[0+1:]
			continue
		}
		mag *= 10
		n = n*10 + int32((func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]) - int32('0')
	}
	if mag > 100000 {
		n /= mag / 100000
		mag /= mag / 100000
	}
	n = readunit(func() int32 {
		if int32(s[0]) != 0 && noarch.Strchr([]byte("icpPvmnu\x00"), int32(s[0])) != nil {
			return int32((func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0])
		}
		return defunit
	}(), n)
	_s[0] = s
	return n / func() int32 {
		if mag > 0 {
			return mag
		}
		return 1
	}()
}

// evaljmp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:58
func evaljmp(s [][]byte, c int32) int32 {
	if int32((s[0])[0]) == c {
		s[0] = (s[0])[0+1:]
		return 0
	}
	return 1
}

// evalisnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:67
func evalisnum(s [][]byte) int32 {
	return noarch.BoolToInt(int32((s[0])[0]) == int32('.') || int32(((__ctype_b_loc())[0])[int32(uint8((s[0])[0]))])&int32(uint16(noarch.ISdigit)) != 0)
}

// evalatom - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:75
func evalatom(s [][]byte) int32 {
	var ret int32
	if evalisnum(s) != 0 {
		return evalnum(s)
	}
	if noarch.Not(evaljmp(s, int32('-'))) {
		return -evalatom(s)
	}
	if noarch.Not(evaljmp(s, int32('+'))) {
		return evalatom(s)
	}
	if noarch.Not(evaljmp(s, int32('|'))) {
		return abspos + evalatom(s)
	}
	if noarch.Not(evaljmp(s, int32('('))) {
		ret = evalexpr(s)
		evaljmp(s, int32(')'))
		return ret
	}
	return 0
}

// nonzero - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:94
func nonzero(n int32) int32 {
	if noarch.Not(n) {
		errdie([]byte("neatroff: divide by zero\n\x00"))
	}
	return n
}

// evalexpr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:101
func evalexpr(s [][]byte) int32 {
	var ret int32 = evalatom(s)
	for (s[0])[0] != 0 {
		if noarch.Not(evaljmp(s, int32('+'))) {
			ret += evalatom(s)
		} else if noarch.Not(evaljmp(s, int32('-'))) {
			ret -= evalatom(s)
		} else if noarch.Not(evaljmp(s, int32('/'))) {
			ret /= nonzero(evalatom(s))
		} else if noarch.Not(evaljmp(s, int32('*'))) {
			ret *= evalatom(s)
		} else if noarch.Not(evaljmp(s, int32('%'))) {
			ret %= nonzero(evalatom(s))
		} else if noarch.Not(evaljmp(s, int32('<'))) {
			if noarch.Not(evaljmp(s, int32('='))) {
				ret = noarch.BoolToInt(ret <= evalatom(s))
			} else {
				ret = noarch.BoolToInt(ret < evalatom(s))
			}
		} else if noarch.Not(evaljmp(s, int32('>'))) {
			if noarch.Not(evaljmp(s, int32('='))) {
				ret = noarch.BoolToInt(ret >= evalatom(s))
			} else {
				ret = noarch.BoolToInt(ret > evalatom(s))
			}
		} else if noarch.BoolToInt(noarch.Not(evaljmp(s, int32('='))))+noarch.BoolToInt(noarch.Not(evaljmp(s, int32('=')))) != 0 {
			ret = noarch.BoolToInt(ret == evalatom(s))
		} else if noarch.Not(evaljmp(s, int32('&'))) {
			ret = noarch.BoolToInt(ret > 0 && evalatom(s) > 0)
		} else if noarch.Not(evaljmp(s, int32(':'))) {
			ret = noarch.BoolToInt(ret > 0 || evalatom(s) > 0)
		} else {
			break
		}
	}
	return ret
}

// eval_up - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:132
func eval_up(s [][]byte, unit int32) int32 {
	// evaluate *s and update s to point to the last character read
	defunit = unit
	if unit == int32('v') {
		abspos = -((nreg(int32('d')))[0])
	}
	if unit == int32('m') {
		abspos = (nreg(map_([]byte(".b0\x00"))))[0] - f_hpos()
	}
	return evalexpr(s)
}

// eval_re - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:143
func eval_re(s []byte, orig int32, unit int32) int32 {
	// evaluate s relative to its previous value
	var n int32
	// n should be added to orig
	var rel int32
	if int32(s[0]) == int32('+') || int32(s[0]) == int32('-') {
		if int32(s[0]) == int32('+') {
			rel = 1
		} else {
			rel = -1
		}
		s = s[0+1:]
	}
	n = eval_up((*[1000000][]byte)(unsafe.Pointer(&s))[:], unit)
	if rel != 0 {
		if rel > 0 {
			return orig + n
		}
		return orig - n
	}
	return n
}

// eval - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/eval.c:158
func eval(s []byte, unit int32) int32 {
	// evaluate s
	return eval_up((*[1000000][]byte)(unsafe.Pointer(&s))[:], unit)
}

// word - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:25
//
// * line formatting buffer for line adjustment and hyphenation
// *
// * The line formatting buffer does two main functions: breaking
// * words into lines (possibly after breaking them at their
// * hyphenation points), and, if requested, adjusting the space
// * between words in a line.  In this file the first step is
// * referred to as filling.
// *
// * Functions like fmt_word() return nonzero on failure, which
// * means the call should be repeated after fetching previously
// * formatted lines via fmt_nextline().
//
type word struct {
	s    []byte
	wid  int32
	elsn int32
	elsp int32
	gap  int32
	hy   int32
	str  int32
	cost int32
	swid int32
}

// line - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:36
// word's width
// els_neg and els_pos
// the space before this word
// hyphen width if inserted after this word
// does the space before it stretch
// the extra cost of line break after this word
// space width after this word (i.e., \w' ')
type line struct {
	sbuf sbuf
	wid  int32
	li   int32
	ll   int32
	lI   int32
	elsn int32
	elsp int32
}

// fmt_ - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:42
type fmt_ struct {
	words      []word
	words_n    int32
	words_sz   int32
	lines      []line
	lines_head int32
	lines_tail int32
	lines_sz   int32
	best       []int32
	best_pos   []int32
	best_dep   []int32
	gap        int32
	nls        int32
	nls_sup    int32
	li         int32
	ll         int32
	lI         int32
	filled     int32
	eos        int32
	fillreq    int32
}

// fmt_confupdate - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:64
func fmt_confupdate(f []fmt_) {
	// queued words
	// queued lines
	// for paragraph adjustment
	// current line
	// space before the next word
	// newlines before the next word
	// suppressed newlines
	// current line indentation and length
	// filled all words in the last fmt_fill()
	// last word ends a sentence
	// fill after the last word (\p)
	// .ll, .in and .ti are delayed until the partial line is output
	f[0].ll = (nreg(int32('l')))[0]
	if (nreg(map_([]byte(".ti\x00"))))[0] >= 0 {
		f[0].li = (nreg(map_([]byte(".ti\x00"))))[0]
	} else {
		f[0].li = (nreg(int32('i')))[0]
	}
	if (nreg(map_([]byte(".tI\x00"))))[0] >= 0 {
		f[0].lI = (nreg(map_([]byte(".tI\x00"))))[0]
	} else {
		f[0].lI = (nreg(int32('I')))[0]
	}
	(nreg(map_([]byte(".ti\x00"))))[0] = -1
	(nreg(map_([]byte(".tI\x00"))))[0] = -1
}

// fmt_confchanged - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:73
func fmt_confchanged(f []fmt_) int32 {
	return noarch.BoolToInt(f[0].ll != (nreg(int32('l')))[0] || f[0].li != func() int32 {
		if (nreg(map_([]byte(".ti\x00"))))[0] >= 0 {
			return (nreg(map_([]byte(".ti\x00"))))[0]
		}
		return (nreg(int32('i')))[0]
	}() || f[0].lI != func() int32 {
		if (nreg(map_([]byte(".tI\x00"))))[0] >= 0 {
			return (nreg(map_([]byte(".tI\x00"))))[0]
		}
		return (nreg(int32('I')))[0]
	}())
}

// fmt_movewords - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:80
func fmt_movewords(a []fmt_, dst int32, src int32, len_ int32) {
	// move words inside an fmt struct
	noarch.Memmove((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&a[0].words[0+dst]))) / int64(1))))[:], (*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&a[0].words[0+src]))) / int64(1))))[:], uint32(len_)*48)
}

// fmt_wordscopy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:86
func fmt_wordscopy(f []fmt_, beg int32, end int32, s []sbuf, els_neg []int32, els_pos []int32) int32 {
	// move words from the buffer to s
	var wcur []word
	var w int32
	var i int32
	els_neg[0] = 0
	els_pos[0] = 0
	for i = beg; i < end; i++ {
		wcur = f[0].words[i:]
		sbuf_printf(s, []byte("%ch'%du'\x00"), c_ec, wcur[0].gap)
		sbuf_append(s, wcur[0].s)
		w += wcur[0].wid + wcur[0].gap
		if wcur[0].elsn < els_neg[0] {
			els_neg[0] = wcur[0].elsn
		}
		if wcur[0].elsp > els_pos[0] {
			els_pos[0] = wcur[0].elsp
		}
		_ = wcur[0].s
	}
	if beg < end {
		wcur = f[0].words[end-1:]
		if wcur[0].hy != 0 {
			sbuf_append(s, []byte("\\(hy\x00"))
		}
		w += wcur[0].hy
	}
	return w
}

// fmt_nlines - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:114
func fmt_nlines(f []fmt_) int32 {
	return f[0].lines_head - f[0].lines_tail
}

// fmt_wordslen - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:120
func fmt_wordslen(f []fmt_, beg int32, end int32) int32 {
	// the total width of the specified words in f->words[]
	var i int32
	var w int32
	for i = beg; i < end; i++ {
		w += f[0].words[i].wid + f[0].words[i].gap
	}
	if beg < end {
		return w + f[0].words[end-1].hy
	}
	return 0
}

// fmt_spaces - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:129
func fmt_spaces(f []fmt_, beg int32, end int32) int32 {
	// the number of stretchable spaces in f
	var i int32
	var n int32
	for i = beg + 1; i < end; i++ {
		if f[0].words[i].str != 0 {
			n++
		}
	}
	return n
}

// fmt_spacessum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:139
func fmt_spacessum(f []fmt_, beg int32, end int32) int32 {
	// the amount of stretchable spaces in f
	var i int32
	var n int32
	for i = beg + 1; i < end; i++ {
		if f[0].words[i].str != 0 {
			n += f[0].words[i].gap
		}
	}
	return n
}

// fmt_nextline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:149
func fmt_nextline(f []fmt_, w []int32, li []int32, lI []int32, ll []int32, els_neg []int32, els_pos []int32) []byte {
	// return the next line in the buffer
	var l []line
	if f[0].lines_head == f[0].lines_tail {
		return nil
	}
	l = f[0].lines[func() int32 {
		tempVar1 := &f[0].lines_tail
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}():]
	li[0] = l[0].li
	lI[0] = l[0].lI
	ll[0] = l[0].ll
	w[0] = l[0].wid
	els_neg[0] = l[0].elsn
	els_pos[0] = l[0].elsp
	return sbuf_out((*[1000000]sbuf)(unsafe.Pointer(&l[0].sbuf))[:])
}

// fmt_mkline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:165
func fmt_mkline(f []fmt_) []line {
	var l []line
	if f[0].lines_head == f[0].lines_tail {
		f[0].lines_head = 0
		f[0].lines_tail = 0
	}
	if f[0].lines_head == f[0].lines_sz {
		f[0].lines_sz += 256
		f[0].lines = mextend(f[0].lines, f[0].lines_head, f[0].lines_sz, int32(48)).([]line)
	}
	l = f[0].lines[func() int32 {
		tempVar1 := &f[0].lines_head
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}():]
	l[0].li = f[0].li
	l[0].lI = f[0].lI
	l[0].ll = f[0].ll
	sbuf_init((*[1000000]sbuf)(unsafe.Pointer(&l[0].sbuf))[:])
	return l
}

// fmt_extractline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:188
func fmt_extractline(f []fmt_, beg int32, end int32, str int32) int32 {
	// extract words from beg to end; shrink or stretch spaces if needed
	var fmt_div int32
	var fmt_rem int32
	var w int32
	var i int32
	var nspc int32
	var llen int32
	var l []line
	if (func() []line {
		l = fmt_mkline(f)
		return l
	}()) == nil {
		return 1
	}
	if 0 < (f)[0].ll-(f)[0].li-(f)[0].lI {
		llen = (f)[0].ll - (f)[0].li - (f)[0].lI
	} else {
		llen = 0
	}
	w = fmt_wordslen(f, beg, end)
	if str != 0 && ((nreg(int32('u')))[0] != 0 && noarch.Not((nreg(map_([]byte(".na\x00"))))[0]) && noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('j')))[0]&3 == 3) && (nreg(int32('j')))[0]&8 != 0 {
		fmt_keshideh(f, beg, end, llen-w)
		w = fmt_wordslen(f, beg, end)
	}
	nspc = fmt_spaces(f, beg, end)
	if nspc != 0 && ((nreg(int32('u')))[0] != 0 && noarch.Not((nreg(map_([]byte(".na\x00"))))[0]) && noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('j')))[0]&3 == 3) && (llen < w || str != 0) {
		fmt_div = (llen - w) / nspc
		fmt_rem = (llen - w) % nspc
		if fmt_rem < 0 {
			fmt_div--
			fmt_rem += nspc
		}
		for i = beg + 1; i < end; i++ {
			if f[0].words[i].str != 0 {
				f[0].words[i].gap += fmt_div + noarch.BoolToInt(func() int32 {
					defer func() {
						fmt_rem--
					}()
					return fmt_rem
				}() > 0)
			}
		}
	}
	l[0].wid = fmt_wordscopy(f, beg, end, (*[1000000]sbuf)(unsafe.Pointer(&l[0].sbuf))[:], (*[1000000]int32)(unsafe.Pointer(&l[0].elsn))[:], (*[1000000]int32)(unsafe.Pointer(&l[0].elsp))[:])
	return 0
}

// fmt_sp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:217
func fmt_sp(f []fmt_) int32 {
	if fmt_fillwords(f, 1) != 0 {
		return 1
	}
	if fmt_extractline(f, 0, f[0].words_n, 0) != 0 {
		return 1
	}
	f[0].filled = 0
	f[0].nls--
	f[0].nls_sup = 0
	f[0].words_n = 0
	f[0].fillreq = 0
	return 0
}

// fmt_fill - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:232
func fmt_fill(f []fmt_, br int32) int32 {
	if fmt_fillwords(f, br) != 0 {
		// fill as many lines as possible; if br, put the remaining words in a line
		return 1
	}
	if br != 0 {
		f[0].filled = 0
		if f[0].words_n != 0 {
			if fmt_sp(f) != 0 {
				return 1
			}
		}
	}
	return 0
}

// fmt_space - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:245
func fmt_space(fmt__c4go_postfix []fmt_) {
	fmt__c4go_postfix[0].gap += font_swid(dev_font((nreg(int32('f')))[0]), (nreg(int32('s')))[0], (nreg(map_([]byte(".ss\x00"))))[0])
}

// fmt_newline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:250
func fmt_newline(f []fmt_) int32 {
	f[0].gap = 0
	if !(noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('u')))[0] != 0) {
		f[0].nls++
		fmt_sp(f)
		return 0
	}
	if f[0].nls >= 1 {
		if fmt_sp(f) != 0 {
			return 1
		}
	}
	if f[0].nls == 0 && noarch.Not(f[0].filled) && noarch.Not(f[0].words_n) {
		fmt_sp(f)
	}
	f[0].nls++
	return 0
}

// fmt_fillreq - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:268
func fmt_fillreq(f []fmt_) int32 {
	if f[0].fillreq > 0 {
		if fmt_fillwords(f, 0) != 0 {
			// format the paragraph after the next word (\p)
			return 1
		}
	}
	f[0].fillreq = f[0].words_n + 1
	return 0
}

// fmt_wb2word - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:277
func fmt_wb2word(f []fmt_, word_c4go_postfix []word, wb_c4go_postfix []wb, hy int32, str int32, gap int32, cost int32) {
	var len_ int32 = noarch.Strlen(wb_buf(wb_c4go_postfix))
	word_c4go_postfix[0].s = xmalloc(len_ + 1).([]byte)
	memcpy(word_c4go_postfix[0].s, wb_buf(wb_c4go_postfix), uint32(len_+1))
	word_c4go_postfix[0].wid = wb_wid(wb_c4go_postfix)
	word_c4go_postfix[0].elsn = wb_c4go_postfix[0].els_neg
	word_c4go_postfix[0].elsp = wb_c4go_postfix[0].els_pos
	if hy != 0 {
		word_c4go_postfix[0].hy = wb_hywid(wb_c4go_postfix)
	} else {
		word_c4go_postfix[0].hy = 0
	}
	word_c4go_postfix[0].str = str
	word_c4go_postfix[0].gap = gap
	word_c4go_postfix[0].cost = cost
	word_c4go_postfix[0].swid = wb_swid(wb_c4go_postfix)
}

// fmt_hyphmarks - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:294
func fmt_hyphmarks(word []byte, hyidx []int32, hyins []int32, hygap []int32) int32 {
	// find explicit break positions: dashes, \:, \%, and \~
	var s []byte = word
	var d []byte
	var c int32
	var n int32
	var lastchar int32
	for (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) > 0 {
	}
	if c < 0 || noarch.Not(noarch.Strcmp(env_hc(), d)) {
		return -1
	}
	for (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) >= 0 && n < 32 {
		if noarch.Not(c) {
			if noarch.Not(noarch.Strcmp(env_hc(), d)) {
				hyins[n] = 1
				hyidx[func() int32 {
					defer func() {
						n++
					}()
					return n
				}()] = int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&word[0])))/int64(1)))
			}
			if c_hydash(d) != 0 {
				hyins[n] = 0
				hyidx[func() int32 {
					defer func() {
						n++
					}()
					return n
				}()] = int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&word[0])))/int64(1)))
			}
			if noarch.Not(noarch.Strcmp([]byte("\\~\x00"), d)) {
				hygap[n] = 1
				hyidx[func() int32 {
					defer func() {
						n++
					}()
					return n
				}()] = int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&word[0])))/int64(1)))
			}
			lastchar = int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&word[0])))/int64(1)))
		}
	}
	for n > 0 && hyidx[n-1] == lastchar {
		// cannot break the end of a word
		n--
	}
	return n
}

// fmt_mkword - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:327
func fmt_mkword(f []fmt_) []word {
	if f[0].words_n == f[0].words_sz {
		f[0].words_sz += 256
		f[0].words = mextend(f[0].words, f[0].words_n, f[0].words_sz, int32(48)).([]word)
	}
	return f[0].words[func() int32 {
		tempVar1 := &f[0].words_n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}():]
}

// fmt_insertword - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:337
func fmt_insertword(f []fmt_, wb_c4go_postfix []wb, gap int32) {
	// sub-word boundaries
	var hyidx []int32 = make([]int32, 32)
	// insert dash
	var hyins []int32 = []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// stretchable no-break space
	var hygap []int32 = []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var src []byte = wb_buf(wb_c4go_postfix)
	var wbc wb
	var beg []byte
	var end []byte
	var n int32
	var i int32
	var cf int32
	var cs int32
	var cm int32
	var ccd int32
	n = fmt_hyphmarks(src, hyidx, hyins, hygap)
	if n <= 0 {
		fmt_wb2word(f, fmt_mkword(f), wb_c4go_postfix, 0, 1, gap, wb_cost(wb_c4go_postfix))
		return
	}
	if f[0].fillreq == f[0].words_n+1 {
		// update f->fillreq considering the new sub-words
		f[0].fillreq += n
	}
	wb_init(c4goUnsafeConvert_wb(&wbc))
	{
		// add sub-words
		for i = 0; i <= n; i++ {
			// dash width
			var ihy int32 = noarch.BoolToInt(i < n && hyins[i] != 0)
			// stretchable
			var istr int32 = noarch.BoolToInt(i == 0 || hygap[i-1] != 0)
			// gap width
			var igap int32
			// hyphenation cost
			var icost int32
			beg = c4goPointerArithByteSlice(src, int(func() int32 {
				if i > 0 {
					return hyidx[i-1]
				}
				return 0
			}()))
			end = src[0+func() int32 {
				if i < n {
					return hyidx[i]
				}
				return noarch.Strlen(src)
			}():]
			if i < n && hygap[i] != 0 {
				// remove \~
				end = c4goPointerArithByteSlice(end, int(-noarch.Strlen([]byte("\\~\x00"))))
			}
			wb_catstr(c4goUnsafeConvert_wb(&wbc), beg, end)
			wb_fnszget(c4goUnsafeConvert_wb(&wbc), c4goUnsafeConvert_int32(&cf), c4goUnsafeConvert_int32(&cs), c4goUnsafeConvert_int32(&cm), c4goUnsafeConvert_int32(&ccd))
			if i == n {
				icost = wb_cost(c4goUnsafeConvert_wb(&wbc))
			} else {
				icost = hygap[i] * 10000000
			}
			if i == 0 {
				igap = gap
			} else {
				igap = hygap[i-1] * wb_swid(c4goUnsafeConvert_wb(&wbc))
			}
			fmt_wb2word(f, fmt_mkword(f), c4goUnsafeConvert_wb(&wbc), ihy, istr, igap, icost)
			wb_reset(c4goUnsafeConvert_wb(&wbc))
			// restoring wbc
			wb_fnszset(c4goUnsafeConvert_wb(&wbc), cf, cs, cm, ccd)
		}
	}
	wb_done(c4goUnsafeConvert_wb(&wbc))
}

// fmt_wordgap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:379
func fmt_wordgap(f []fmt_) int32 {
	// the amount of space necessary before the next word
	var nls int32 = noarch.BoolToInt(f[0].nls != 0 || f[0].nls_sup != 0)
	var swid int32 = font_swid(dev_font((nreg(int32('f')))[0]), (nreg(int32('s')))[0], (nreg(map_([]byte(".ss\x00"))))[0])
	if f[0].eos != 0 && f[0].words_n != 0 {
		if nls != 0 && noarch.Not(f[0].gap) || noarch.Not(nls) && f[0].gap == 2*swid {
			return swid + font_swid(dev_font((nreg(int32('f')))[0]), (nreg(int32('s')))[0], (nreg(map_([]byte(".sss\x00"))))[0])
		}
	}
	if nls != 0 && noarch.Not(f[0].gap) && f[0].words_n != 0 {
		return swid
	}
	return f[0].gap
}

// fmt_word - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:390
func fmt_word(f []fmt_, wb_c4go_postfix []wb) int32 {
	if wb_empty(wb_c4go_postfix) != 0 {
		// insert wb into fmt
		return 0
	}
	if fmt_confchanged(f) != 0 {
		if fmt_fillwords(f, 0) != 0 {
			return 1
		}
	}
	if noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('u')))[0] != 0 && f[0].nls != 0 && f[0].gap != 0 {
		if fmt_sp(f) != 0 {
			return 1
		}
	}
	if noarch.Not(f[0].words_n) {
		// apply the new .l and .i
		fmt_confupdate(f)
	}
	f[0].gap = fmt_wordgap(f)
	f[0].eos = wb_eos(wb_c4go_postfix)
	fmt_insertword(f, wb_c4go_postfix, func() int32 {
		if f[0].filled != 0 {
			return 0
		}
		return f[0].gap
	}())
	f[0].filled = 0
	f[0].nls = 0
	f[0].nls_sup = 0
	f[0].gap = 0
	return 0
}

// fmt_keshideh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:413
func fmt_keshideh(f []fmt_, beg int32, end int32, wid int32) {
	// insert keshideh characters
	var wb_c4go_postfix wb
	var kw int32
	var i int32
	var c int32
	var w []word
	var cnt int32
	for {
		cnt = 0
		for c = 0; c < 2; c++ {
			for i = end - 1 - c; i >= beg; i -= 2 {
				w = f[0].words[i:]
				wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
				kw = wb_keshideh(w[0].s, c4goUnsafeConvert_wb(&wb_c4go_postfix), wid)
				if kw > 0 {
					_ = w[0].s
					w[0].s = xmalloc(noarch.Strlen(wb_buf(c4goUnsafeConvert_wb(&wb_c4go_postfix))) + int32(1)).([]byte)
					noarch.Strcpy(w[0].s, wb_buf(c4goUnsafeConvert_wb(&wb_c4go_postfix)))
					w[0].wid = wb_wid(c4goUnsafeConvert_wb(&wb_c4go_postfix))
					wid -= kw
					cnt++
				}
				wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
			}
		}
		if noarch.Not(cnt) {
			break
		}
	}
}

// scaledown - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:441
func scaledown(cost int32) int32 {
	// approximate 8 * sqrt(cost)
	var ret int32
	var i int32
	for i = 0; i < 14; i++ {
		ret += cost >> uint64(i*2) & 3 << uint64(i+3)
	}
	if ret < 1<<uint64(13) {
		return ret
	}
	return 1 << uint64(13)
}

// FMT_COST - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:451
func FMT_COST(llen int32, lwid int32, swid int32, nspc int32) int32 {
	// the cost of putting lwid words in a line of length llen
	// the ratio that the stretchable spaces of the line should be spread
	var ratio int32 = noarch.Labs((llen - lwid) * 100 / func() int32 {
		if swid != 0 {
			return swid
		}
		return 1
	}())
	if ratio > 4000 {
		// ratio too large; scaling it down
		ratio = 4000 + scaledown(ratio-4000)
	}
	// assigning a cost of 100 to each space stretching 100 percent
	return ratio * ratio / 100 * func() int32 {
		if nspc != 0 {
			return nspc
		}
		return 1
	}()
}

// fmt_hydepth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:463
func fmt_hydepth(f []fmt_, pos int32) int32 {
	// the number of hyphenations in consecutive lines ending at pos
	var n int32
	for pos > 0 && f[0].words[pos-1].hy != 0 && func() int32 {
		n++
		return n
	}() < 5 {
		pos = f[0].best_pos[pos]
	}
	return n
}

// hycost - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:471
func hycost(depth int32) int32 {
	if (nreg(map_([]byte(".hlm\x00"))))[0] > 0 && depth > (nreg(map_([]byte(".hlm\x00"))))[0] {
		return 10000000
	}
	if depth >= 3 {
		return (nreg(map_([]byte(".hycost\x00"))))[0] + (nreg(map_([]byte(".hycost2\x00"))))[0] + (nreg(map_([]byte(".hycost3\x00"))))[0]
	}
	if depth == 2 {
		return (nreg(map_([]byte(".hycost\x00"))))[0] + (nreg(map_([]byte(".hycost2\x00"))))[0]
	}
	if depth != 0 {
		return (nreg(map_([]byte(".hycost\x00"))))[0]
	}
	return 0
}

// fmt_findcost - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:483
func fmt_findcost(f []fmt_, pos int32) int32 {
	// the cost of putting a line break before word pos
	var i int32
	var hyphenated int32
	var cur int32
	var llen int32 = func() int32 {
		if 1 < func() int32 {
			if 0 < (f)[0].ll-(f)[0].li-(f)[0].lI {
				return (f)[0].ll - (f)[0].li - (f)[0].lI
			}
			return 0
		}() {
			if 0 < (f)[0].ll-(f)[0].li-(f)[0].lI {
				return (f)[0].ll - (f)[0].li - (f)[0].lI
			}
			return 0
		}
		return 1
	}()
	// current line length
	var lwid int32
	// amount of stretchable spaces
	var swid int32
	// number of stretchable spaces
	var nspc int32
	// equal to swid, unless swid is zero
	var dwid int32
	if pos <= 0 {
		return 0
	}
	if f[0].best_pos[pos] >= 0 {
		return f[0].best[pos] + f[0].words[pos-1].cost
	}
	// non-zero if the last word is hyphenated
	lwid = f[0].words[pos-1].hy
	hyphenated = noarch.BoolToInt(f[0].words[pos-1].hy != 0)
	i = pos - 1
	for i >= 0 {
		lwid += f[0].words[i].wid
		if i+1 < pos {
			lwid += f[0].words[i+1].gap
		}
		if i+1 < pos && f[0].words[i+1].str != 0 {
			swid += f[0].words[i+1].gap
			nspc++
		}
		if lwid > llen+swid*(nreg(map_([]byte(".ssh\x00"))))[0]/100 && i+1 < pos {
			break
		}
		dwid = swid
		if noarch.Not(dwid) && i > 0 {
			// no stretchable spaces
			dwid = f[0].words[i-1].swid
		}
		cur = fmt_findcost(f, i) + FMT_COST(llen, lwid, dwid, nspc)
		if hyphenated != 0 {
			cur += hycost(1 + fmt_hydepth(f, i))
		}
		if f[0].best_pos[pos] < 0 || cur < f[0].best[pos] {
			f[0].best_pos[pos] = i
			f[0].best_dep[pos] = f[0].best_dep[i] + 1
			f[0].best[pos] = cur
		}
		i--
	}
	return f[0].best[pos] + f[0].words[pos-1].cost
}

// fmt_bestpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:525
func fmt_bestpos(f []fmt_, pos int32) int32 {
	fmt_findcost(f, pos)
	if 0 < f[0].best_pos[pos] {
		return f[0].best_pos[pos]
	}
	return 0
}

// fmt_bestdep - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:531
func fmt_bestdep(f []fmt_, pos int32) int32 {
	fmt_findcost(f, pos)
	if 0 < f[0].best_dep[pos] {
		return f[0].best_dep[pos]
	}
	return 0
}

// fmt_breakparagraph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:538
func fmt_breakparagraph(f []fmt_, pos int32, br int32) int32 {
	// return the last filled word
	var i int32
	var best int32 = -1
	var cost int32
	var best_cost int32
	var llen int32 = func() int32 {
		if 0 < (f)[0].ll-(f)[0].li-(f)[0].lI {
			return (f)[0].ll - (f)[0].li - (f)[0].lI
		}
		return 0
	}()
	// current line length
	var lwid int32
	// amount of stretchable spaces
	var swid int32
	// number of stretchable spaces
	var nspc int32
	if f[0].fillreq > 0 && f[0].fillreq <= f[0].words_n {
		fmt_findcost(f, f[0].fillreq)
		return f[0].fillreq
	}
	if pos > 0 && f[0].words[pos-1].wid >= llen {
		fmt_findcost(f, pos)
		return pos
	}
	i = pos - 1
	lwid = 0
	if f[0].words[i].hy != 0 {
		// the last word is hyphenated
		lwid += f[0].words[i].hy
	}
	for i >= 0 {
		lwid += f[0].words[i].wid
		if i+1 < pos {
			lwid += f[0].words[i+1].gap
		}
		if i+1 < pos && f[0].words[i+1].str != 0 {
			swid += f[0].words[i+1].gap
			nspc++
		}
		if lwid > llen && i+1 < pos {
			break
		}
		cost = fmt_findcost(f, i)
		if br != 0 && (nreg(map_([]byte(".pmll\x00"))))[0] != 0 && lwid < llen*(nreg(map_([]byte(".pmll\x00"))))[0]/100 {
			// the cost of formatting short lines; should prevent widows
			var pmll int32 = llen * (nreg(map_([]byte(".pmll\x00"))))[0] / 100
			cost += (nreg(map_([]byte(".pmllcost\x00"))))[0] * (pmll - lwid) / pmll
		}
		if best < 0 || cost < best_cost {
			best = i
			best_cost = cost
		}
		i--
	}
	return best
}

// fmt_head - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:585
func fmt_head(f []fmt_, nreq int32, pos int32, nohy int32) int32 {
	// extract the first nreq formatted lines before the word at pos
	// best line break for nreq-th line
	var best int32 = pos
	// best line breaks without hyphenation
	var prev int32
	var next int32
	if nreq <= 0 || fmt_bestdep(f, pos) < nreq {
		return pos
	}
	for best > 0 && fmt_bestdep(f, best) > nreq {
		// finding the optimal line break for nreq-th line
		best = fmt_bestpos(f, best)
	}
	prev = best
	next = best
	if noarch.Not(nohy) {
		return best
	}
	for prev > 1 && f[0].words[prev-1].hy != 0 && fmt_bestdep(f, prev-1) == nreq {
		// finding closest line breaks without hyphenation
		prev--
	}
	for next < pos && f[0].words[next-1].hy != 0 && fmt_bestdep(f, next) == nreq {
		next++
	}
	if noarch.Not(f[0].words[prev-1].hy) && noarch.Not(f[0].words[next-1].hy) {
		// choosing the best of them
		if fmt_findcost(f, prev) <= fmt_findcost(f, next) {
			return prev
		}
		return next
	}
	if noarch.Not(f[0].words[prev-1].hy) {
		return prev
	}
	if noarch.Not(f[0].words[next-1].hy) {
		return next
	}
	return best
}

// fmt_break - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:616
func fmt_break(f []fmt_, end int32) int32 {
	// break f->words[0..end] into lines according to fmt_bestpos()
	var beg int32
	var ret int32
	beg = fmt_bestpos(f, end)
	if beg > 0 {
		ret += fmt_break(f, beg)
	}
	f[0].words[beg].gap = 0
	if fmt_extractline(f, beg, end, 1) != 0 {
		return ret
	}
	if beg > 0 {
		fmt_confupdate(f)
	}
	return ret + (end - beg)
}

// fmt_safelines - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:631
func fmt_safelines() int32 {
	// estimated number of lines until traps or the end of a page
	var lnht int32 = func() int32 {
		if 1 < (nreg(int32('L')))[0] {
			return (nreg(int32('L')))[0]
		}
		return 1
	}() * (nreg(int32('v')))[0]
	if (nreg(int32('v')))[0] > 0 {
		return (f_nexttrap() + lnht - 1) / lnht
	}
	return 1000
}

// fmt_fillwords - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:638
func fmt_fillwords(f []fmt_, br int32) int32 {
	// fill the words collected in the buffer
	// the number of lines until a trap
	var nreq int32
	// the final line ends before this word
	var end int32
	// like end, but only the first nreq lines included
	var end_head int32
	// only nreq first lines have been formatted
	var head int32
	// line length, taking shrinkable spaces into account
	var llen int32
	var n int32
	var i int32
	if !(noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('u')))[0] != 0) {
		return 0
	}
	llen = fmt_wordslen(f, 0, f[0].words_n) - fmt_spacessum(f, 0, f[0].words_n)*(nreg(map_([]byte(".ssh\x00"))))[0]/100
	if (f[0].fillreq <= 0 || f[0].words_n < f[0].fillreq) && llen <= func() int32 {
		if 0 < (f)[0].ll-(f)[0].li-(f)[0].lI {
			return (f)[0].ll - (f)[0].li - (f)[0].lI
		}
		return 0
	}() {
		// not enough words to fill
		return 0
	}
	// lines until a trap or page end
	nreq = fmt_safelines()
	if fmt_confchanged(f) != 0 {
		// if line settings are changed, output a single line
		nreq = 1
	}
	if nreq > 0 && nreq <= fmt_nlines(f) {
		// enough lines are collected already
		return 1
	}
	// resetting positions
	f[0].best = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, uint32(f[0].words_n+1)*uint32(1))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	f[0].best_pos = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, uint32(f[0].words_n+1)*uint32(1))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	f[0].best_dep = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, uint32(f[0].words_n+1)*uint32(1))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&f[0].best[0]))) / int64(1))))[:], byte(0), uint32(f[0].words_n+1)*8)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&f[0].best_dep[0]))) / int64(1))))[:], byte(0), uint32(f[0].words_n+1)*4)
	for i = 0; i < f[0].words_n+1; i++ {
		f[0].best_pos[i] = -1
	}
	end = fmt_breakparagraph(f, f[0].words_n, br)
	if nreq > 0 {
		// do not hyphenate the last line
		var nohy int32
		if (nreg(map_([]byte(".hy\x00"))))[0]&2 != 0 && nreq == fmt_nlines(f) {
			nohy = 1
		}
		end_head = fmt_head(f, nreq-fmt_nlines(f), end, nohy)
		head = noarch.BoolToInt(end_head < end)
		end = end_head
	}
	// recursively add lines
	if end > 0 {
		n = fmt_break(f, end)
	} else {
		n = 0
	}
	f[0].words_n -= n
	f[0].fillreq -= n
	fmt_movewords(f, 0, n, f[0].words_n)
	f[0].filled = noarch.BoolToInt(n != 0 && noarch.Not(f[0].words_n))
	if f[0].words_n != 0 {
		f[0].words[0].gap = 0
	}
	if f[0].words_n != 0 {
		// apply the new .l and .i
		fmt_confupdate(f)
	}
	_ = f[0].best
	_ = f[0].best_pos
	_ = f[0].best_dep
	f[0].best = nil
	f[0].best_pos = nil
	f[0].best_dep = nil
	return noarch.BoolToInt(head != 0 || n != end)
}

// fmt_alloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:697
func fmt_alloc() []fmt_ {
	var fmt__c4go_postfix []fmt_ = xmalloc(int32(136)).([]fmt_)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&fmt__c4go_postfix[0]))) / int64(1))))[:], byte(0), 136)
	return fmt__c4go_postfix
}

// fmt_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:704
func fmt_free(fmt__c4go_postfix []fmt_) {
	_ = fmt__c4go_postfix[0].lines
	_ = fmt__c4go_postfix[0].words
	_ = fmt__c4go_postfix
}

// fmt_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:711
func fmt_wid(fmt__c4go_postfix []fmt_) int32 {
	return fmt_wordslen(fmt__c4go_postfix, 0, fmt__c4go_postfix[0].words_n) + fmt_wordgap(fmt__c4go_postfix)
}

// fmt_morewords - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:716
func fmt_morewords(fmt__c4go_postfix []fmt_) int32 {
	return noarch.BoolToInt(fmt_morelines(fmt__c4go_postfix) != 0 || fmt__c4go_postfix[0].words_n != 0)
}

// fmt_morelines - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:721
func fmt_morelines(fmt__c4go_postfix []fmt_) int32 {
	return noarch.BoolToInt(fmt__c4go_postfix[0].lines_head != fmt__c4go_postfix[0].lines_tail)
}

// fmt_suppressnl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/fmt.c:727
func fmt_suppressnl(fmt__c4go_postfix []fmt_) {
	if fmt__c4go_postfix[0].nls != 0 {
		// suppress the last newline
		fmt__c4go_postfix[0].nls--
		fmt__c4go_postfix[0].nls_sup = 1
	}
}

// gpat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:17
// gpat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:18
// font handling
// convert wid in device unitwidth size to size sz
// flags for gpat->flg
// glyph substitution and positioning rules
// rule description
type gpat struct {
	g    int32
	flg  int16
	x    int16
	y    int16
	xadv int16
	yadv int16
}
type grule struct {
	pats []gpat
	sec  int32
	len_ int16
	feat int16
	scrp int16
	lang int16
}

// font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:28
// glyph index
// pattern flags; GF_*
// gpos data
// rule pattern
// rule section (OFF lookup)
// pats[] length
// rule's feature and script
type font struct {
	name      [32]byte
	fontname  [32]byte
	spacewid  int32
	special   int32
	cs        int32
	cs_ps     int32
	bd        int32
	zoom      int32
	s1        int32
	n1        int32
	s2        int32
	n2        int32
	gl        []glyph
	gl_n      int32
	gl_sz     int32
	gl_dict   []dict
	ch_dict   []dict
	ch_map    []dict
	feat_name [128][]byte
	feat_set  [128]int32
	scrp_name [64][]byte
	scrp      int32
	lang_name [64][]byte
	lang      int32
	secs      int32
	gsub      []grule
	gsub_n    int32
	gsub_sz   int32
	gpos      []grule
	gpos_n    int32
	gpos_sz   int32
	gsub0     []iset
	gpos0     []iset
	ggrp      []iset
}

// font_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:59
func font_find(fn []font, name []byte) []glyph {
	// for .cs, .bd, .fzoom requests
	// for .tkf request
	// glyphs present in the font
	// number of glyphs in the font
	// mapping from gl[i].id to i
	// charset mapping
	// characters mapped via font_map()
	// font features and scripts
	// feature names
	// feature enabled
	// script names
	// current script
	// language names
	// current language
	// number of font sections (OFF lookups)
	// glyph substitution and positioning
	// glyph substitution rules
	// glyph positioning rules
	// rules matching a glyph at pos 0
	// rules matching a glyph at pos 0
	// glyph groups
	// find a glyph by its name
	var i int32 = dict_get(fn[0].ch_map, name)
	if i == -1 {
		// -2 means the glyph has been unmapped
		i = dict_get(fn[0].ch_dict, name)
	}
	if i >= 0 {
		return fn[0].gl[0+i:]
	}
	return nil
}

// font_glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:68
func font_glyph(fn []font, id []byte) []glyph {
	// find a glyph by its device-dependent identifier
	var i int32 = dict_get(fn[0].gl_dict, id)
	if i >= 0 {
		return fn[0].gl[i:]
	}
	return nil
}

// font_glyphput - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:74
func font_glyphput(fn []font, id []byte, name []byte, type_ int32) int32 {
	var g []glyph
	if fn[0].gl_n == fn[0].gl_sz {
		fn[0].gl_sz = fn[0].gl_sz + 1024
		fn[0].gl = mextend(fn[0].gl, fn[0].gl_n, fn[0].gl_sz, int32(120)).([]glyph)
	}
	g = fn[0].gl[fn[0].gl_n:]
	noarch.Snprintf(g[0].id[:], int32(32), []byte("%s\x00"), id)
	noarch.Snprintf(g[0].name[:], int32(32), []byte("%s\x00"), name)
	g[0].type_ = int16(type_)
	g[0].font = fn
	dict_put(fn[0].gl_dict, g[0].id[:], fn[0].gl_n)
	tempVar1 := &fn[0].gl_n
	defer func() {
		*tempVar1++
	}()
	return *tempVar1
}

// font_map - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:91
func font_map(fn []font, name []byte, id []byte) int32 {
	// map character name to the given glyph; remove the mapping if id is NULL
	var gidx int32 = -1
	if id != nil {
		if font_glyph(fn, id) != nil {
			gidx = int32((func() int64 {
				c4go_temp_name := font_glyph(fn, id)
				return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
			}() - int64(uintptr(unsafe.Pointer(&fn[0].gl[0])))/int64(120)))
		} else {
			gidx = -2
		}
	}
	dict_put(fn[0].ch_map, name, gidx)
	return 0
}

// font_mapped - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:101
func font_mapped(fn []font, name []byte) int32 {
	// return nonzero if character name has been mapped with font_map()
	return noarch.BoolToInt(dict_get(fn[0].ch_map, name) != -1)
}

// font_featlg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:109
func font_featlg(fn []font, val int32) int32 {
	// enable/disable ligatures; first bit for liga and the second bit for rlig
	var ret int32
	ret |= font_feat(fn, []byte("liga\x00"), val&1)
	ret |= font_feat(fn, []byte("rlig\x00"), val&2) << uint64(1)
	return ret
}

// font_featkn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:118
func font_featkn(fn []font, val int32) int32 {
	// enable/disable pairwise kerning
	return font_feat(fn, []byte("kern\x00"), val)
}

// font_idx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:124
func font_idx(fn []font, g []glyph) int32 {
	// glyph index in fn->glyphs[]
	if g != nil {
		return int32((int64(uintptr(unsafe.Pointer(&g[0])))/int64(120) - int64(uintptr(unsafe.Pointer(&fn[0].gl[0])))/int64(120)))
	}
	return -1
}

// font_gpatmatch - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:129
func font_gpatmatch(fn []font, p []gpat, g int32) int32 {
	var r []int32
	if noarch.Not(int32(p[0].flg) & 8) {
		return noarch.BoolToInt(p[0].g == g)
	}
	r = iset_get(fn[0].ggrp, p[0].g)
	for r != nil && r[0] >= 0 {
		if (func() []int32 {
			defer func() {
				r = r[0+1:]
			}()
			return r
		}())[0] == g {
			return 1
		}
	}
	return 0
}

// font_rulematch - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:141
func font_rulematch(fn []font, rule []grule, src []int32, slen int32, dst []int32, dlen int32) int32 {
	// the index of matched glyphs in src
	var sidx int32
	// number of initial context glyphs
	var ncon int32
	var pats []gpat = rule[0].pats
	var j int32
	if fn[0].scrp >= 0 && fn[0].scrp != int32(rule[0].scrp) {
		// enable only the active script if set
		return 0
	}
	if int32(rule[0].lang) >= 0 && fn[0].lang != int32(rule[0].lang) {
		// enable common script features and those in the active language
		return 0
	}
	if noarch.Not(fn[0].feat_set[:][rule[0].feat]) {
		return 0
	}
	{
		// the number of initial context glyphs
		for j = 0; j < int32(rule[0].len_) && int32(pats[j].flg)&4 != 0; j++ {
			ncon++
		}
	}
	if dlen < ncon {
		return 0
	}
	for ; j < int32(rule[0].len_); j++ {
		if int32(pats[j].flg)&2 != 0 {
			// matching the base pattern
			continue
		}
		if sidx >= slen || noarch.Not(font_gpatmatch(fn, pats[j:], src[sidx])) {
			return 0
		}
		sidx++
	}
	{
		// matching the initial context
		for j = 0; j < int32(rule[0].len_) && int32(pats[j].flg)&4 != 0; j++ {
			if noarch.Not(font_gpatmatch(fn, pats[j:], dst[j-ncon])) {
				return 0
			}
		}
	}
	return 1
}

// font_findrule - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:177
func font_findrule(fn []font, gsub int32, pos int32, fwd []int32, fwdlen int32, ctx []int32, ctxlen int32, idx []int32) int32 {
	// find a matching gsub/gpos rule; *idx should be -1 initially
	var rules []grule = func() []grule {
		if gsub != 0 {
			return fn[0].gsub
		}
		return fn[0].gpos
	}()
	var r1 []int32 = iset_get(func() []iset {
		if gsub != 0 {
			return fn[0].gsub0
		}
		return fn[0].gpos0
	}(), fwd[0])
	for r1 != nil && r1[func() int32 {
		tempVar1 := &idx[0]
		*tempVar1++
		return *tempVar1
	}()] >= 0 {
		if r1[idx[0]] >= pos && font_rulematch(fn, rules[r1[idx[0]]:], fwd, fwdlen, ctx, ctxlen) != 0 {
			return r1[idx[0]]
		}
	}
	return -1
}

// font_performgpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:191
func font_performgpos(fn []font, src []int32, slen int32, x []int32, y []int32, xadv []int32, yadv []int32) {
	// perform all possible gpos rules on src
	var gpos []grule = fn[0].gpos
	var pats []gpat
	var curs_feat int32 = func() int32 {
		if dir_do != 0 {
			return font_findfeat(fn, []byte("curs\x00"))
		}
		return -1
	}()
	var curs_beg int32 = -1
	var curs_dif int32
	var i int32
	var k int32
	for i = 0; i < slen; i++ {
		var idx int32 = -1
		var curs_cur int32
		var lastsec int32 = -1
		for 1 != 0 {
			var r int32 = font_findrule(fn, 0, 0, src[0+i:], slen-i, src[0+i:], i, c4goUnsafeConvert_int32(&idx))
			if r < 0 {
				// no rule found
				break
			}
			if gpos[r].sec > 0 && gpos[r].sec <= lastsec {
				// perform at most one rule from each lookup
				continue
			}
			lastsec = gpos[r].sec
			pats = gpos[r].pats
			for k = 0; k < int32(gpos[r].len_); k++ {
				x[i+k] += int32(pats[k].x)
				y[i+k] += int32(pats[k].y)
				xadv[i+k] += int32(pats[k].xadv)
				yadv[i+k] += int32(pats[k].yadv)
			}
			if int32(gpos[r].feat) == curs_feat {
				curs_cur = 1
				if curs_beg < 0 {
					curs_beg = i
				}
				for k = 0; k < int32(gpos[r].len_); k++ {
					curs_dif += int32(pats[k].yadv)
				}
			}
		}
		if curs_beg >= 0 && noarch.Not(curs_cur) {
			yadv[curs_beg] -= curs_dif
			curs_beg = -1
			curs_dif = 0
		}
	}
	if curs_beg >= 0 {
		yadv[curs_beg] -= curs_dif
	}
}

// font_firstgsub - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:238
func font_firstgsub(fn []font, pos int32, src []int32, slen int32) int32 {
	// find the first gsub rule after pos that matches any glyph in src
	var best int32 = -1
	var i int32
	for i = 0; i < slen; i++ {
		var idx int32 = -1
		var r int32 = font_findrule(fn, 1, pos, src[0+i:], slen-i, src[0+i:], i, c4goUnsafeConvert_int32(&idx))
		if r >= 0 && (best < 0 || r < best) {
			best = r
		}
	}
	return best
}

// font_gsubapply - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:253
func font_gsubapply(fn []font, rule []grule, src []int32, slen int32, smap []int32) int32 {
	// apply the given gsub rule to all matches in src
	var dst []int32 = make([]int32, 256)
	var dlen int32
	var dmap []int32 = make([]int32, 256)
	var i int32
	var j int32
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&dmap[0]))) / int64(1))))[:], byte(0), uint32(slen)*4)
	for i = 0; i < slen; i++ {
		dmap[dlen] = smap[i]
		if font_rulematch(fn, rule, src[0+i:], slen-i, dst[0+dlen:], dlen) != 0 {
			for j = 0; j < int32(rule[0].len_); j++ {
				if int32(rule[0].pats[j].flg)&2 != 0 {
					dst[func() int32 {
						defer func() {
							dlen++
						}()
						return dlen
					}()] = rule[0].pats[j].g
				}
				if int32(rule[0].pats[j].flg)&1 != 0 {
					i++
				}
			}
			i--
		} else {
			dst[func() int32 {
				defer func() {
					dlen++
				}()
				return dlen
			}()] = src[i]
		}
	}
	memcpy(src, dst, uint32(dlen)*4)
	memcpy(smap, dmap, uint32(dlen)*4)
	return dlen
}

// font_performgsub - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:282
func font_performgsub(fn []font, src []int32, slen int32, smap []int32) int32 {
	// perform all possible gsub rules on src
	var i int32 = -1
	for func() int32 {
		i++
		return i
	}() >= 0 {
		if (func() int32 {
			i = font_firstgsub(fn, i, src, slen)
			return i
		}()) < 0 {
			break
		}
		slen = font_gsubapply(fn, fn[0].gsub[i:], src, slen, smap)
	}
	return slen
}

// font_layout - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:293
func font_layout(fn []font, gsrc [][]glyph, nsrc int32, sz int32, gdst [][]glyph, dmap []int32, x []int32, y []int32, xadv []int32, yadv []int32, lg int32, kn int32) int32 {
	var dst []int32 = make([]int32, 256)
	var ndst int32 = nsrc
	var i int32
	var featlg int32
	var featkn int32
	{
		// initialising dst
		for i = 0; i < nsrc; i++ {
			dst[i] = font_idx(fn, gsrc[i])
		}
	}
	for i = 0; i < ndst; i++ {
		dmap[i] = i
	}
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&x[0]))) / int64(1))))[:], byte(0), uint32(ndst)*4)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&y[0]))) / int64(1))))[:], byte(0), uint32(ndst)*4)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&xadv[0]))) / int64(1))))[:], byte(0), uint32(ndst)*4)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&yadv[0]))) / int64(1))))[:], byte(0), uint32(ndst)*4)
	if lg != 0 {
		// substitution rules
		featlg = font_featlg(fn, 3)
	}
	ndst = font_performgsub(fn, dst, ndst, dmap)
	if lg != 0 {
		font_featlg(fn, featlg)
	}
	if kn != 0 {
		// positioning rules
		featkn = font_featkn(fn, 1)
	}
	font_performgpos(fn, dst, ndst, x, y, xadv, yadv)
	if kn != 0 {
		font_featkn(fn, featkn)
	}
	for i = 0; i < ndst; i++ {
		gdst[i] = fn[0].gl[0+dst[i]:]
	}
	return ndst
}

// font_readchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:327
func font_readchar(fn []font, fin *noarch.File, n []int32, gid []int32) int32 {
	var g []glyph
	var tok []byte = make([]byte, 128)
	var name []byte = make([]byte, 32)
	var id []byte = make([]byte, 32)
	var type_ int32
	if noarch.Fscanf(fin, []byte("%32s %128s\x00"), name, tok) != 2 {
		return 1
	}
	if noarch.Not(noarch.Strcmp([]byte("---\x00"), name)) {
		noarch.Sprintf(name, []byte("c%04d\x00"), n[0])
	}
	if noarch.Strcmp([]byte("\"\x00"), tok) != 0 {
		if noarch.Fscanf(fin, []byte("%d %32s\x00"), c4goUnsafeConvert_int32(&type_), id) != 2 {
			return 1
		}
		gid[0] = font_glyphput(fn, id, name, type_)
		g = fn[0].gl[gid[0]:]
		noarch.Sscanf(tok, []byte("%hd,%hd,%hd,%hd,%hd\x00"), (*[1000000]int16)(unsafe.Pointer(&g[0].wid))[:], (*[1000000]int16)(unsafe.Pointer(&g[0].llx))[:], (*[1000000]int16)(unsafe.Pointer(&g[0].lly))[:], (*[1000000]int16)(unsafe.Pointer(&g[0].urx))[:], (*[1000000]int16)(unsafe.Pointer(&g[0].ury))[:])
		dict_put(fn[0].ch_dict, name, gid[0])
		n[0]++
	} else {
		dict_put(fn[0].ch_map, name, gid[0])
	}
	return 0
}

// font_findfeat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:353
func font_findfeat(fn []font, feat []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 1024/8 && int32(fn[0].feat_name[:][i][0]) != 0; i++ {
		if noarch.Not(noarch.Strcmp(feat, fn[0].feat_name[:][i])) {
			return i
		}
	}
	if uint32(i) < 1024/8 {
		noarch.Snprintf(fn[0].feat_name[:][i], int32(8), []byte("%s\x00"), feat)
		return i
	}
	return -1
}

// font_findscrp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:366
func font_findscrp(fn []font, scrp []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 512/8 && int32(fn[0].scrp_name[:][i][0]) != 0; i++ {
		if noarch.Not(noarch.Strcmp(scrp, fn[0].scrp_name[:][i])) {
			return i
		}
	}
	if uint32(i) == 512/8 {
		return -1
	}
	noarch.Snprintf(fn[0].scrp_name[:][i], int32(8), []byte("%s\x00"), scrp)
	return i
}

// font_findlang - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:378
func font_findlang(fn []font, lang []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 512/8 && int32(fn[0].lang_name[:][i][0]) != 0; i++ {
		if noarch.Not(noarch.Strcmp(lang, fn[0].lang_name[:][i])) {
			return i
		}
	}
	if uint32(i) == 512/8 {
		return -1
	}
	noarch.Snprintf(fn[0].lang_name[:][i], int32(8), []byte("%s\x00"), lang)
	return i
}

// font_gpat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:390
func font_gpat(fn []font, len_ int32) []gpat {
	var pats []gpat = xmalloc(int32(uint32(len_) * 16)).([]gpat)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&pats[0]))) / int64(1))))[:], byte(0), uint32(len_)*16)
	return pats
}

// font_gsub - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:397
func font_gsub(fn []font, len_ int32, feat int32, scrp int32, lang int32) []grule {
	var rule []grule
	var pats []gpat = font_gpat(fn, len_)
	if fn[0].gsub_n == fn[0].gsub_sz {
		fn[0].gsub_sz = fn[0].gsub_sz + 1024
		fn[0].gsub = mextend(fn[0].gsub, fn[0].gsub_n, fn[0].gsub_sz, int32(40)).([]grule)
	}
	rule = fn[0].gsub[func() int32 {
		tempVar1 := &fn[0].gsub_n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}():]
	rule[0].pats = pats
	rule[0].len_ = int16(len_)
	rule[0].feat = int16(feat)
	rule[0].scrp = int16(scrp)
	rule[0].lang = int16(lang)
	return rule
}

// font_gpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:415
func font_gpos(fn []font, len_ int32, feat int32, scrp int32, lang int32) []grule {
	var rule []grule
	var pats []gpat = font_gpat(fn, len_)
	if fn[0].gpos_n == fn[0].gpos_sz {
		fn[0].gpos_sz = fn[0].gpos_sz + 1024
		fn[0].gpos = mextend(fn[0].gpos, fn[0].gpos_n, fn[0].gpos_sz, int32(40)).([]grule)
	}
	rule = fn[0].gpos[func() int32 {
		tempVar1 := &fn[0].gpos_n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}():]
	rule[0].pats = pats
	rule[0].len_ = int16(len_)
	rule[0].feat = int16(feat)
	rule[0].scrp = int16(scrp)
	rule[0].lang = int16(lang)
	return rule
}

// font_readgpat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:433
func font_readgpat(fn []font, p []gpat, s []byte) int32 {
	if int32(s[0]) == int32('@') {
		p[0].g = noarch.Atoi(s[0+1:])
		if iset_len(fn[0].ggrp, p[0].g) == 1 {
			p[0].g = iset_get(fn[0].ggrp, p[0].g)[0]
		} else {
			p[0].flg |= int16(8)
		}
	} else {
		p[0].g = font_idx(fn, font_glyph(fn, s))
	}
	return noarch.BoolToInt(p[0].g < 0)
}

// font_readfeat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:447
func font_readfeat(fn []font, tok []byte, feat []int32, scrp []int32, lang []int32) {
	var ftag []byte = tok
	var stag []byte
	var ltag []byte
	scrp[0] = -1
	lang[0] = -1
	if noarch.Strchr(ftag, int32(':')) != nil {
		stag = (noarch.Strchr(ftag, int32(':')))[0+1:]
		c4goPointerArithByteSlice(stag, int(-1))[0] = '\x00'
	}
	if noarch.Strchr(stag, int32(':')) != nil {
		ltag = (noarch.Strchr(stag, int32(':')))[0+1:]
		c4goPointerArithByteSlice(ltag, int(-1))[0] = '\x00'
	}
	if stag != nil {
		scrp[0] = font_findscrp(fn, stag)
	}
	if ltag != nil {
		lang[0] = font_findlang(fn, ltag)
	}
	feat[0] = font_findfeat(fn, tok)
}

// font_readgsub - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:469
func font_readgsub(fn []font, fin *noarch.File) int32 {
	var tok []byte = make([]byte, 128)
	var rule []grule
	var feat int32
	var scrp int32
	var lang int32
	var i int32
	var n int32
	if noarch.Fscanf(fin, []byte("%128s %d\x00"), tok, c4goUnsafeConvert_int32(&n)) != 2 {
		return 1
	}
	font_readfeat(fn, tok, c4goUnsafeConvert_int32(&feat), c4goUnsafeConvert_int32(&scrp), c4goUnsafeConvert_int32(&lang))
	rule = font_gsub(fn, n, feat, scrp, lang)
	rule[0].sec = fn[0].secs
	for i = 0; i < n; i++ {
		if noarch.Fscanf(fin, []byte("%128s\x00"), tok) != 1 {
			return 1
		}
		if int32(tok[0]) == int32('-') {
			rule[0].pats[i].flg = 1
		}
		if int32(tok[0]) == int32('=') {
			rule[0].pats[i].flg = 4
		}
		if int32(tok[0]) == int32('+') {
			rule[0].pats[i].flg = 2
		}
		if noarch.Not(tok[0]) || font_readgpat(fn, rule[0].pats[i:], tok[0+1:]) != 0 {
			return 0
		}
	}
	return 0
}

// font_readgpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:495
func font_readgpos(fn []font, fin *noarch.File) int32 {
	var tok []byte = make([]byte, 128)
	var col []byte
	var rule []grule
	var feat int32
	var scrp int32
	var lang int32
	var i int32
	var n int32
	if noarch.Fscanf(fin, []byte("%128s %d\x00"), tok, c4goUnsafeConvert_int32(&n)) != 2 {
		return 1
	}
	font_readfeat(fn, tok, c4goUnsafeConvert_int32(&feat), c4goUnsafeConvert_int32(&scrp), c4goUnsafeConvert_int32(&lang))
	rule = font_gpos(fn, n, feat, scrp, lang)
	rule[0].sec = fn[0].secs
	for i = 0; i < n; i++ {
		if noarch.Fscanf(fin, []byte("%128s\x00"), tok) != 1 {
			return 1
		}
		col = noarch.Strchr(tok, int32(':'))
		if col != nil {
			col[0] = '\x00'
		}
		rule[0].pats[i].flg = 1
		if noarch.Not(tok[0]) || font_readgpat(fn, rule[0].pats[i:], tok) != 0 {
			return 0
		}
		if col != nil {
			noarch.Sscanf(col[0+1:], []byte("%hd%hd%hd%hd\x00"), (*[1000000]int16)(unsafe.Pointer(&rule[0].pats[i].x))[:], (*[1000000]int16)(unsafe.Pointer(&rule[0].pats[i].y))[:], (*[1000000]int16)(unsafe.Pointer(&rule[0].pats[i].xadv))[:], (*[1000000]int16)(unsafe.Pointer(&rule[0].pats[i].yadv))[:])
		}
	}
	return 0
}

// font_readggrp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:524
func font_readggrp(fn []font, fin *noarch.File) int32 {
	var tok []byte = make([]byte, 32)
	var id int32
	var n int32
	var i int32
	var g int32
	if noarch.Fscanf(fin, []byte("%d %d\x00"), c4goUnsafeConvert_int32(&id), c4goUnsafeConvert_int32(&n)) != 2 {
		return 1
	}
	for i = 0; i < n; i++ {
		if noarch.Fscanf(fin, []byte("%32s\x00"), tok) != 1 {
			return 1
		}
		g = font_idx(fn, font_glyph(fn, tok))
		if g >= 0 {
			iset_put(fn[0].ggrp, id, g)
		}
	}
	return 0
}

// font_readkern - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:540
func font_readkern(fn []font, fin *noarch.File) int32 {
	var c1 []byte = make([]byte, 32)
	var c2 []byte = make([]byte, 32)
	var rule []grule
	var val int32
	if noarch.Fscanf(fin, []byte("%32s %32s %d\x00"), c1, c2, c4goUnsafeConvert_int32(&val)) != 3 {
		return 1
	}
	rule = font_gpos(fn, 2, font_findfeat(fn, []byte("kern\x00")), -1, -1)
	rule[0].pats[0].g = font_idx(fn, font_glyph(fn, c1))
	rule[0].pats[1].g = font_idx(fn, font_glyph(fn, c2))
	rule[0].pats[0].xadv = int16(val)
	rule[0].pats[0].flg = 1
	rule[0].pats[1].flg = 1
	return 0
}

// font_lig - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:556
func font_lig(fn []font, lig []byte) {
	var c []byte = make([]byte, 32)
	var g []int32 = make([]int32, 256)
	var rule []grule
	var s []byte = lig
	var j int32
	var n int32
	for utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], c) > 0 {
		g[func() int32 {
			defer func() {
				n++
			}()
			return n
		}()] = font_idx(fn, font_find(fn, c))
	}
	rule = font_gsub(fn, n+1, font_findfeat(fn, []byte("liga\x00")), -1, -1)
	for j = 0; j < n; j++ {
		rule[0].pats[j].g = g[j]
		rule[0].pats[j].flg = 1
	}
	rule[0].pats[n].g = font_idx(fn, font_find(fn, lig))
	rule[0].pats[n].flg = 2
}

// font_rulefirstpat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:582
func font_rulefirstpat(fn []font, rule []grule) []gpat {
	var i int32
	for i = 0; i < int32(rule[0].len_); i++ {
		if noarch.Not(int32(rule[0].pats[i].flg) & (2 | 4)) {
			return rule[0].pats[i:]
		}
	}
	return nil
}

// font_isetinsert - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:591
func font_isetinsert(fn []font, iset_c4go_postfix []iset, rule int32, p []gpat) {
	if int32(p[0].flg)&8 != 0 {
		var r []int32 = iset_get(fn[0].ggrp, p[0].g)
		for r != nil && r[0] >= 0 {
			iset_put(iset_c4go_postfix, (func() []int32 {
				defer func() {
					r = r[0+1:]
				}()
				return r
			}())[0], rule)
		}
	} else {
		if p[0].g >= 0 {
			iset_put(iset_c4go_postfix, p[0].g, rule)
		}
	}
}

// font_open - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:603
func font_open(path []byte) []font {
	var fn []font
	// last glyph in the charset
	var ch_g int32 = -1
	// number of glyphs in the charset
	var ch_n int32
	var tok []byte = make([]byte, 128)
	var fin *noarch.File
	var ligs [][]byte = make([][]byte, 512)
	var ligs_n int32
	var sec int32
	var i int32
	fin = noarch.Fopen(path, []byte("r\x00"))
	if fin == nil {
		return nil
	}
	fn = xmalloc(int32(2888)).([]font)
	if fn == nil {
		noarch.Fclose(fin)
		return nil
	}
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&fn[0]))) / int64(1))))[:], byte(0), 2888)
	fn[0].gl_dict = dict_make(-1, 1, 0)
	fn[0].ch_dict = dict_make(-1, 1, 0)
	fn[0].ch_map = dict_make(-1, 1, 0)
	fn[0].ggrp = iset_make()
	for noarch.Fscanf(fin, []byte("%128s\x00"), tok) == 1 {
		if noarch.Not(noarch.Strcmp([]byte("char\x00"), tok)) {
			font_readchar(fn, fin, c4goUnsafeConvert_int32(&ch_n), c4goUnsafeConvert_int32(&ch_g))
		} else if noarch.Not(noarch.Strcmp([]byte("kern\x00"), tok)) {
			font_readkern(fn, fin)
		} else if noarch.Not(noarch.Strcmp([]byte("ligatures\x00"), tok)) {
			for noarch.Fscanf(fin, []byte("%s\x00"), ligs[ligs_n]) == 1 {
				if noarch.Not(noarch.Strcmp([]byte("0\x00"), ligs[ligs_n])) {
					break
				}
				if uint32(ligs_n) < 16384/32 {
					ligs_n++
				}
			}
		} else if noarch.Not(noarch.Strcmp([]byte("gsec\x00"), tok)) {
			if noarch.Fscanf(fin, []byte("%d\x00"), c4goUnsafeConvert_int32(&sec)) != 1 {
				fn[0].secs++
			}
		} else if noarch.Not(noarch.Strcmp([]byte("gsub\x00"), tok)) {
			font_readgsub(fn, fin)
		} else if noarch.Not(noarch.Strcmp([]byte("gpos\x00"), tok)) {
			font_readgpos(fn, fin)
		} else if noarch.Not(noarch.Strcmp([]byte("ggrp\x00"), tok)) {
			font_readggrp(fn, fin)
		} else if noarch.Not(noarch.Strcmp([]byte("spacewidth\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%d\x00"), (*[1000000]int32)(unsafe.Pointer(&fn[0].spacewid))[:])
		} else if noarch.Not(noarch.Strcmp([]byte("special\x00"), tok)) {
			fn[0].special = 1
		} else if noarch.Not(noarch.Strcmp([]byte("name\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%s\x00"), fn[0].name[:])
		} else if noarch.Not(noarch.Strcmp([]byte("fontname\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%s\x00"), fn[0].fontname[:])
		} else if noarch.Not(noarch.Strcmp([]byte("charset\x00"), tok)) {
			for noarch.Not(font_readchar(fn, fin, c4goUnsafeConvert_int32(&ch_n), c4goUnsafeConvert_int32(&ch_g))) {
			}
			break
		}
		skipline(fin)
	}
	for i = 0; i < ligs_n; i++ {
		font_lig(fn, ligs[i])
	}
	noarch.Fclose(fin)
	fn[0].gsub0 = iset_make()
	fn[0].gpos0 = iset_make()
	for i = 0; i < fn[0].gsub_n; i++ {
		font_isetinsert(fn, fn[0].gsub0, i, font_rulefirstpat(fn, fn[0].gsub[i:]))
	}
	for i = 0; i < fn[0].gpos_n; i++ {
		font_isetinsert(fn, fn[0].gpos0, i, font_rulefirstpat(fn, fn[0].gpos[i:]))
	}
	fn[0].scrp = -1
	fn[0].lang = -1
	return fn
}

// font_close - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:679
func font_close(fn []font) {
	var i int32
	for i = 0; i < fn[0].gsub_n; i++ {
		_ = fn[0].gsub[i].pats
	}
	for i = 0; i < fn[0].gpos_n; i++ {
		_ = fn[0].gpos[i].pats
	}
	dict_free(fn[0].gl_dict)
	dict_free(fn[0].ch_dict)
	dict_free(fn[0].ch_map)
	iset_free(fn[0].gsub0)
	iset_free(fn[0].gpos0)
	iset_free(fn[0].ggrp)
	_ = fn[0].gsub
	_ = fn[0].gpos
	_ = fn[0].gl
	_ = fn
}

// font_special - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:698
func font_special(fn []font) int32 {
	return fn[0].special
}

// font_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:704
func font_wid(fn []font, sz int32, w int32) int32 {
	// return width w for the given font and size
	sz = font_zoom(fn, sz)
	if w >= 0 {
		return (w*sz + dev_uwid/2) / dev_uwid
	}
	return -((-w*sz + dev_uwid/2) / dev_uwid)
}

// font_twid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:711
func font_twid(fn []font, sz int32) int32 {
	if fn[0].s1 >= 0 && sz <= fn[0].s1 {
		// return track kerning width for the given size
		return fn[0].n1 * (dev_res / 72)
	}
	if fn[0].s2 >= 0 && sz >= fn[0].s2 {
		return fn[0].n2 * (dev_res / 72)
	}
	if sz > fn[0].s1 && sz < fn[0].s2 {
		return ((sz-fn[0].s1)*fn[0].n1 + (fn[0].s2-sz)*fn[0].n2) * (dev_res / 72) / (fn[0].s2 - fn[0].s1)
	}
	return 0
}

// font_gwid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:724
func font_gwid(fn []font, cfn []font, sz int32, w int32) int32 {
	// glyph width, where cfn is the current font and fn is glyph's font
	var xfn []font = func() []font {
		if cfn != nil {
			return cfn
		}
		return fn
	}()
	if xfn[0].cs != 0 {
		return xfn[0].cs * (font_zoom(fn, func() int32 {
			if xfn[0].cs_ps != 0 {
				return xfn[0].cs_ps
			}
			return sz
		}()) * dev_res / 72) / 36
	}
	return font_wid(fn, sz, w) + func() int32 {
		if cfn != nil {
			return font_twid(fn, sz)
		}
		return 0
	}() + func() int32 {
		if font_getbd(xfn) != 0 {
			return font_getbd(xfn) - 1
		}
		return 0
	}()
}

// font_swid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:735
func font_swid(fn []font, sz int32, ss int32) int32 {
	// space width for the give word space or sentence space
	return font_gwid(fn, nil, sz, (fn[0].spacewid*ss+6)/12)
}

// font_getcs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:740
func font_getcs(fn []font) int32 {
	return fn[0].cs
}

// font_setcs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:745
func font_setcs(fn []font, cs int32, ps int32) {
	fn[0].cs = cs
	fn[0].cs_ps = ps
}

// font_getbd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:751
func font_getbd(fn []font) int32 {
	return fn[0].bd
}

// font_setbd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:756
func font_setbd(fn []font, bd int32) {
	fn[0].bd = bd
}

// font_track - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:761
func font_track(fn []font, s1 int32, n1 int32, s2 int32, n2 int32) {
	fn[0].s1 = s1
	fn[0].n1 = n1
	fn[0].s2 = s2
	fn[0].n2 = n2
}

// font_zoom - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:769
func font_zoom(fn []font, sz int32) int32 {
	if fn[0].zoom != 0 {
		return (sz*fn[0].zoom + 500) / 1000
	}
	return sz
}

// font_setzoom - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:774
func font_setzoom(fn []font, zoom int32) {
	fn[0].zoom = zoom
}

// font_feat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:780
func font_feat(fn []font, name []byte, val int32) int32 {
	// enable/disable font features; returns the previous value
	var idx int32 = font_findfeat(fn, name)
	var old int32 = func() int32 {
		if idx >= 0 {
			return fn[0].feat_set[:][idx]
		}
		return 0
	}()
	if idx >= 0 {
		fn[0].feat_set[:][idx] = noarch.BoolToInt(val != 0)
	}
	return old
}

// font_scrp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:790
func font_scrp(fn []font, name []byte) {
	// set font script
	if name != nil {
		fn[0].scrp = font_findscrp(fn, name)
	} else {
		fn[0].scrp = -1
	}
}

// font_lang - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:796
func font_lang(fn []font, name []byte) {
	// set font language
	if name != nil {
		fn[0].lang = font_findlang(fn, name)
	} else {
		fn[0].lang = -1
	}
}

// en_patterns - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyen.h:3
// english hyphenation patterns and exceptions
var en_patterns []byte = []byte(".ach4 .ad4der .af1t .al3t .am5at .an5c .ang4 .ani5m .ant4 .an3te .anti5s .ar5s .ar4tie .ar4ty .as3c .as1p .as1s .aster5 .atom5 .au1d .av4i .awn4 .ba4g .ba5na .bas4e .ber4 .be5ra .be3sm .be5sto .bri2 .but4ti .cam4pe .can5c .capa5b .car5ol .ca4t .ce4la .ch4 .chill5i .ci2 .cit5r .co3e .co4r .cor5ner .de4moi .de3o .de3ra .de3ri .des4c .dictio5 .do4t .du4c .dumb5 .earth5 .eas3i .eb4 .eer4 .eg2 .el5d .el3em .enam3 .en3g .en3s .eq5ui5t .er4ri .es3 .eu3 .eye5 .fes3 .for5mer .ga2 .ge2 .gen3t4 .ge5og .gi5a .gi4b .go4r .hand5i .han5k .he2 .hero5i .hes3 .het3 .hi3b .hi3er .hon5ey .hon3o .hov5 .id4l .idol3 .im3m .im5pin .in1 .in3ci .ine2 .in2k .in3s .ir5r .is4i .ju3r .la4cy .la4m .lat5er .lath5 .le2 .leg5e .len4 .lep5 .lev1 .li4g .lig5a .li2n .li3o .li4t .mag5a5 .mal5o .man5a .mar5ti .me2 .mer3c .me5ter .mis1 .mist5i .mon3e .mo3ro .mu5ta .muta5b .ni4c .od2 .odd5 .of5te .or5ato .or3c .or1d .or3t .os3 .os4tl .oth3 .out3 .ped5al .pe5te .pe5tit .pi4e .pio5n .pi2t .pre3m .ra4c .ran4t .ratio5na .ree2 .re5mit .res2 .re5stat .ri4g .rit5u .ro4q .ros5t .row5d .ru4d .sci3e .self5 .sell5 .se2n .se5rie .sh2 .si2 .sing4 .st4 .sta5bl .sy2 .ta4 .te4 .ten5an .th2 .ti2 .til4 .tim5o5 .ting4 .tin5k .ton4a .to4p .top5i .tou5s .trib5ut .un1a .un3ce .under5 .un1e .un5k .un5o .un3u .up3 .ure3 .us5a .ven4de .ve5ra .wil5i .ye4 4ab. a5bal a5ban abe2 ab5erd abi5a ab5it5ab ab5lat ab5o5liz 4abr ab5rog ab3ul a4car ac5ard ac5aro a5ceou ac1er a5chet 4a2ci a3cie ac1in a3cio ac5rob act5if ac3ul ac4um a2d ad4din ad5er. 2adi a3dia ad3ica adi4er a3dio a3dit a5diu ad4le ad3ow ad5ran ad4su 4adu a3duc ad5um ae4r aeri4e a2f aff4 a4gab aga4n ag5ell age4o 4ageu ag1i 4ag4l ag1n a2go 3agog ag3oni a5guer ag5ul a4gy a3ha a3he ah4l a3ho ai2 a5ia a3ic. ai5ly a4i4n ain5in ain5o ait5en a1j ak1en al5ab al3ad a4lar 4aldi 2ale al3end a4lenti a5le5o al1i al4ia. ali4e al5lev 4allic 4alm a5log. a4ly. 4alys 5a5lyst 5alyt 3alyz 4ama am5ab am3ag ama5ra am5asc a4matis a4m5ato am5era am3ic am5if am5ily am1in ami4no a2mo a5mon amor5i amp5en a2n an3age 3analy a3nar an3arc anar4i a3nati 4and ande4s an3dis an1dl an4dow a5nee a3nen an5est. a3neu 2ang ang5ie an1gl a4n1ic a3nies an3i3f an4ime a5nimi a5nine an3io a3nip an3ish an3it a3niu an4kli 5anniz ano4 an5ot anoth5 an2sa an4sco an4sn an2sp ans3po an4st an4sur antal4 an4tie 4anto an2tr an4tw an3ua an3ul a5nur 4ao apar4 ap5at ap5ero a3pher 4aphi a4pilla ap5illar ap3in ap3ita a3pitu a2pl apoc5 ap5ola apor5i apos3t aps5es a3pu aque5 2a2r ar3act a5rade ar5adis ar3al a5ramete aran4g ara3p ar4at a5ratio ar5ativ a5rau ar5av4 araw4 arbal4 ar4chan ar5dine ar4dr ar5eas a3ree ar3ent a5ress ar4fi ar4fl ar1i ar5ial ar3ian a3riet ar4im ar5inat ar3io ar2iz ar2mi ar5o5d a5roni a3roo ar2p ar3q arre4 ar4sa ar2sh 4as. as4ab as3ant ashi4 a5sia. a3sib a3sic 5a5si4t ask3i as4l a4soc as5ph as4sh as3ten as1tr asur5a a2ta at3abl at5ac at3alo at5ap ate5c at5ech at3ego at3en. at3era ater5n a5terna at3est at5ev 4ath ath5em a5then at4ho ath5om 4ati. a5tia at5i5b at1ic at3if ation5ar at3itu a4tog a2tom at5omiz a4top a4tos a1tr at5rop at4sk at4tag at5te at4th a2tu at5ua at5ue at3ul at3ura a2ty au4b augh3 au3gu au4l2 aun5d au3r au5sib aut5en au1th a2va av3ag a5van ave4no av3era av5ern av5ery av1i avi4er av3ig av5oc a1vor 3away aw3i aw4ly aws4 ax4ic ax4id ay5al aye4 ays4 azi4er azz5i 5ba. bad5ger ba4ge bal1a ban5dag ban4e ban3i barbi5 bari4a bas4si 1bat ba4z 2b1b b2be b3ber bbi4na 4b1d 4be. beak4 beat3 4be2d be3da be3de be3di be3gi be5gu 1bel be1li be3lo 4be5m be5nig be5nu 4bes4 be3sp be5str 3bet bet5iz be5tr be3tw be3w be5yo 2bf 4b3h bi2b bi4d 3bie bi5en bi4er 2b3if 1bil bi3liz bina5r4 bin4d bi5net bi3ogr bi5ou bi2t 3bi3tio bi3tr 3bit5ua b5itz b1j bk4 b2l2 blath5 b4le. blen4 5blesp b3lis b4lo blun4t 4b1m 4b3n bne5g 3bod bod3i bo4e bol3ic bom4bi bon4a bon5at 3boo 5bor. 4b1ora bor5d 5bore 5bori 5bos4 b5ota both5 bo4to bound3 4bp 4brit broth3 2b5s2 bsor4 2bt bt4l b4to b3tr buf4fer bu4ga bu3li bumi4 bu4n bunt4i bu3re bus5ie buss4e 5bust 4buta 3butio b5uto b1v 4b5w 5by. bys4 1ca cab3in ca1bl cach4 ca5den 4cag4 2c5ah ca3lat cal4la call5in 4calo can5d can4e can4ic can5is can3iz can4ty cany4 ca5per car5om cast5er cas5tig 4casy ca4th 4cativ cav5al c3c ccha5 cci4a ccompa5 ccon4 ccou3t 2ce. 4ced. 4ceden 3cei 5cel. 3cell 1cen 3cenc 2cen4e 4ceni 3cent 3cep ce5ram 4cesa 3cessi ces5si5b ces5t cet4 c5e4ta cew4 2ch 4ch. 4ch3ab 5chanic ch5a5nis che2 cheap3 4ched che5lo 3chemi ch5ene ch3er. ch3ers 4ch1in 5chine. ch5iness 5chini 5chio 3chit chi2z 3cho2 ch4ti 1ci 3cia ci2a5b cia5r ci5c 4cier 5cific. 4cii ci4la 3cili 2cim 2cin c4ina 3cinat cin3em c1ing c5ing. 5cino cion4 4cipe ci3ph 4cipic 4cista 4cisti 2c1it cit3iz 5ciz ck1 ck3i 1c4l4 4clar c5laratio 5clare cle4m 4clic clim4 cly4 c5n 1co co5ag coe2 2cog co4gr coi4 co3inc col5i 5colo col3or com5er con4a c4one con3g con5t co3pa cop3ic co4pl 4corb coro3n cos4e cov1 cove4 cow5a coz5e co5zi c1q cras5t 5crat. 5cratic cre3at 5cred 4c3reta cre4v cri2 cri5f c4rin cris4 5criti cro4pl crop5o cros4e cru4d 4c3s2 2c1t cta4b ct5ang c5tant c2te c3ter c4ticu ctim3i ctu4r c4tw cud5 c4uf c4ui cu5ity 5culi cul4tis 3cultu cu2ma c3ume cu4mi 3cun cu3pi cu5py cur5a4b cu5ria 1cus cuss4i 3c4ut cu4tie 4c5utiv 4cutr 1cy cze4 1d2a 5da. 2d3a4b dach4 4daf 2dag da2m2 dan3g dard5 dark5 4dary 3dat 4dativ 4dato 5dav4 dav5e 5day d1b d5c d1d4 2de. deaf5 deb5it de4bon decan4 de4cil de5com 2d1ed 4dee. de5if deli4e del5i5q de5lo d4em 5dem. 3demic dem5ic. de5mil de4mons demor5 1den de4nar de3no denti5f de3nu de1p de3pa depi4 de2pu d3eq d4erh 5derm dern5iz der5s des2 d2es. de1sc de2s5o des3ti de3str de4su de1t de2to de1v dev3il 4dey 4d1f d4ga d3ge4t dg1i d2gy d1h2 5di. 1d4i3a dia5b di4cam d4ice 3dict 3did 5di3en d1if di3ge di4lato d1in 1dina 3dine. 5dini di5niz 1dio dio5g di4pl dir2 di1re dirt5i dis1 5disi d4is3t d2iti 1di1v d1j d5k2 4d5la 3dle. 3dled 3dles. 4dless 2d3lo 4d5lu 2dly d1m 4d1n4 1do 3do. do5de 5doe 2d5of d4og do4la doli4 do5lor dom5iz do3nat doni4 doo3d dop4p d4or 3dos 4d5out do4v 3dox d1p 1dr drag5on 4drai dre4 drea5r 5dren dri4b dril4 dro4p 4drow 5drupli 4dry 2d1s2 ds4p d4sw d4sy d2th 1du d1u1a du2c d1uca duc5er 4duct. 4ducts du5el du4g d3ule dum4be du4n 4dup du4pe d1v d1w d2y 5dyn dy4se dys5p e1a4b e3act ead1 ead5ie ea4ge ea5ger ea4l eal5er eal3ou eam3er e5and ear3a ear4c ear5es ear4ic ear4il ear5k ear2t eart3e ea5sp e3ass east3 ea2t eat5en eath3i e5atif e4a3tu ea2v eav3en eav5i eav5o 2e1b e4bel. e4bels e4ben e4bit e3br e4cad ecan5c ecca5 e1ce ec5essa ec2i e4cib ec5ificat ec5ifie ec5ify ec3im eci4t e5cite e4clam e4clus e2col e4comm e4compe e4conc e2cor ec3ora eco5ro e1cr e4crem ec4tan ec4te e1cu e4cul ec3ula 2e2da 4ed3d e4d1er ede4s 4edi e3dia ed3ib ed3ica ed3im ed1it edi5z 4edo e4dol edon2 e4dri e4dul ed5ulo ee2c eed3i ee2f eel3i ee4ly ee2m ee4na ee4p1 ee2s4 eest4 ee4ty e5ex e1f e4f3ere 1eff e4fic 5efici efil4 e3fine ef5i5nite 3efit efor5es e4fuse. 4egal eger4 eg5ib eg4ic eg5ing e5git5 eg5n e4go. e4gos eg1ul e5gur 5egy e1h4 eher4 ei2 e5ic ei5d eig2 ei5gl e3imb e3inf e1ing e5inst eir4d eit3e ei3th e5ity e1j e4jud ej5udi eki4n ek4la e1la e4la. e4lac elan4d el5ativ e4law elaxa4 e3lea el5ebra 5elec e4led el3ega e5len e4l1er e1les el2f el2i e3libe e4l5ic. el3ica e3lier el5igib e5lim e4l3ing e3lio e2lis el5ish e3liv3 4ella el4lab ello4 e5loc el5og el3op. el2sh el4ta e5lud el5ug e4mac e4mag e5man em5ana em5b e1me e2mel e4met em3ica emi4e em5igra em1in2 em5ine em3i3ni e4mis em5ish e5miss em3iz 5emniz emo4g emoni5o em3pi e4mul em5ula emu3n e3my en5amo e4nant ench4er en3dic e5nea e5nee en3em en5ero en5esi en5est en3etr e3new en5ics e5nie e5nil e3nio en3ish en3it e5niu 5eniz 4enn 4eno eno4g e4nos en3ov en4sw ent5age 4enthes en3ua en5uf e3ny. 4en3z e5of eo2g e4oi4 e3ol eop3ar e1or eo3re eo5rol eos4 e4ot eo4to e5out e5ow e2pa e3pai ep5anc e5pel e3pent ep5etitio ephe4 e4pli e1po e4prec ep5reca e4pred ep3reh e3pro e4prob ep4sh ep5ti5b e4put ep5uta e1q equi3l e4q3ui3s er1a era4b 4erand er3ar 4erati. 2erb er4bl er3ch er4che 2ere. e3real ere5co ere3in er5el. er3emo er5ena er5ence 4erene er3ent ere4q er5ess er3est eret4 er1h er1i e1ria4 5erick e3rien eri4er er3ine e1rio 4erit er4iu eri4v e4riva er3m4 er4nis 4ernit 5erniz er3no 2ero er5ob e5roc ero4r er1ou er1s er3set ert3er 4ertl er3tw 4eru eru4t 5erwau e1s4a e4sage. e4sages es2c e2sca es5can e3scr es5cu e1s2e e2sec es5ecr es5enc e4sert. e4serts e4serva 4esh e3sha esh5en e1si e2sic e2sid es5iden es5igna e2s5im es4i4n esis4te esi4u e5skin es4mi e2sol es3olu e2son es5ona e1sp es3per es5pira es4pre 2ess es4si4b estan4 es3tig es5tim 4es2to e3ston 2estr e5stro estruc5 e2sur es5urr es4w eta4b eten4d e3teo ethod3 et1ic e5tide etin4 eti4no e5tir e5titio et5itiv 4etn et5ona e3tra e3tre et3ric et5rif et3rog et5ros et3ua et5ym et5z 4eu e5un e3up eu3ro eus4 eute4 euti5l eu5tr eva2p5 e2vas ev5ast e5vea ev3ell evel3o e5veng even4i ev1er e5verb e1vi ev3id evi4l e4vin evi4v e5voc e5vu e1wa e4wag e5wee e3wh ewil5 ew3ing e3wit 1exp 5eyc 5eye. eys4 1fa fa3bl fab3r fa4ce 4fag fain4 fall5e 4fa4ma fam5is 5far far5th fa3ta fa3the 4fato fault5 4f5b 4fd 4fe. feas4 feath3 fe4b 4feca 5fect 2fed fe3li fe4mo fen2d fend5e fer1 5ferr fev4 4f1f f4fes f4fie f5fin. f2f5is f4fly f2fy 4fh 1fi fi3a 2f3ic. 4f3ical f3ican 4ficate f3icen fi3cer fic4i 5ficia 5ficie 4fics fi3cu fi5del fight5 fil5i fill5in 4fily 2fin 5fina fin2d5 fi2ne f1in3g fin4n fis4ti f4l2 f5less flin4 flo3re f2ly5 4fm 4fn 1fo 5fon fon4de fon4t fo2r fo5rat for5ay fore5t for4i fort5a fos5 4f5p fra4t f5rea fres5c fri2 fril4 frol5 2f3s 2ft f4to f2ty 3fu fu5el 4fug fu4min fu5ne fu3ri fusi4 fus4s 4futa 1fy 1ga gaf4 5gal. 3gali ga3lo 2gam ga5met g5amo gan5is ga3niz gani5za 4gano gar5n4 gass4 gath3 4gativ 4gaz g3b gd4 2ge. 2ged geez4 gel4in ge5lis ge5liz 4gely 1gen ge4nat ge5niz 4geno 4geny 1geo ge3om g4ery 5gesi geth5 4geto ge4ty ge4v 4g1g2 g2ge g3ger gglu5 ggo4 gh3in gh5out gh4to 5gi. 1gi4a gia5r g1ic 5gicia g4ico gien5 5gies. gil4 g3imen 3g4in. gin5ge 5g4ins 5gio 3gir gir4l g3isl gi4u 5giv 3giz gl2 gla4 glad5i 5glas 1gle gli4b g3lig 3glo glo3r g1m g4my gn4a g4na. gnet4t g1ni g2nin g4nio g1no g4non 1go 3go. gob5 5goe 3g4o4g go3is gon2 4g3o3na gondo5 go3ni 5goo go5riz gor5ou 5gos. gov1 g3p 1gr 4grada g4rai gran2 5graph. g5rapher 5graphic 4graphy 4gray gre4n 4gress. 4grit g4ro gruf4 gs2 g5ste gth3 gu4a 3guard 2gue 5gui5t 3gun 3gus 4gu4t g3w 1gy 2g5y3n gy5ra h3ab4l hach4 hae4m hae4t h5agu ha3la hala3m ha4m han4ci han4cy 5hand. han4g hang5er hang5o h5a5niz han4k han4te hap3l hap5t ha3ran ha5ras har2d hard3e har4le harp5en har5ter has5s haun4 5haz haz3a h1b 1head 3hear he4can h5ecat h4ed he5do5 he3l4i hel4lis hel4ly h5elo hem4p he2n hena4 hen5at heo5r hep5 h4era hera3p her4ba here5a h3ern h5erou h3ery h1es he2s5p he4t het4ed heu4 h1f h1h hi5an hi4co high5 h4il2 himer4 h4ina hion4e hi4p hir4l hi3ro hir4p hir4r his3el his4s hith5er hi2v 4hk 4h1l4 hlan4 h2lo hlo3ri 4h1m hmet4 2h1n h5odiz h5ods ho4g hoge4 hol5ar 3hol4e ho4ma home3 hon4a ho5ny 3hood hoon4 hor5at ho5ris hort3e ho5ru hos4e ho5sen hos1p 1hous house3 hov5el 4h5p 4hr4 hree5 hro5niz hro3po 4h1s2 h4sh h4tar ht1en ht5es h4ty hu4g hu4min hun5ke hun4t hus3t4 hu4t h1w h4wart hy3pe hy3ph hy2s 2i1a i2al iam4 iam5ete i2an 4ianc ian3i 4ian4t ia5pe iass4 i4ativ ia4tric i4atu ibe4 ib3era ib5ert ib5ia ib3in ib5it. ib5ite i1bl ib3li i5bo i1br i2b5ri i5bun 4icam 5icap 4icar i4car. i4cara icas5 i4cay iccu4 4iceo 4ich 2ici i5cid ic5ina i2cip ic3ipa i4cly i2c5oc 4i1cr 5icra i4cry ic4te ictu2 ic4t3ua ic3ula ic4um ic5uo i3cur 2id i4dai id5anc id5d ide3al ide4s i2di id5ian idi4ar i5die id3io idi5ou id1it id5iu i3dle i4dom id3ow i4dr i2du id5uo 2ie4 ied4e 5ie5ga ield3 ien5a4 ien4e i5enn i3enti i1er. i3esc i1est i3et 4if. if5ero iff5en if4fr 4ific. i3fie i3fl 4ift 2ig iga5b ig3era ight3i 4igi i3gib ig3il ig3in ig3it i4g4l i2go ig3or ig5ot i5gre igu5i ig1ur i3h 4i5i4 i3j 4ik i1la il3a4b i4lade i2l5am ila5ra i3leg il1er ilev4 il5f il1i il3ia il2ib il3io il4ist 2ilit il2iz ill5ab 4iln il3oq il4ty il5ur il3v i4mag im3age ima5ry imenta5r 4imet im1i im5ida imi5le i5mini 4imit im4ni i3mon i2mu im3ula 2in. i4n3au 4inav incel4 in3cer 4ind in5dling 2ine i3nee iner4ar i5ness 4inga 4inge in5gen 4ingi in5gling 4ingo 4ingu 2ini i5ni. i4nia in3io in1is i5nite. 5initio in3ity 4ink 4inl 2inn 2i1no i4no4c ino4s i4not 2ins in3se insur5a 2int. 2in4th in1u i5nus 4iny 2io 4io. ioge4 io2gr i1ol io4m ion3at ion4ery ion3i io5ph ior3i i4os io5th i5oti io4to i4our 2ip ipe4 iphras4 ip3i ip4ic ip4re4 ip3ul i3qua iq5uef iq3uid iq3ui3t 4ir i1ra ira4b i4rac ird5e ire4de i4ref i4rel4 i4res ir5gi ir1i iri5de ir4is iri3tu 5i5r2iz ir4min iro4g 5iron. ir5ul 2is. is5ag is3ar isas5 2is1c is3ch 4ise is3er 3isf is5han is3hon ish5op is3ib isi4d i5sis is5itiv 4is4k islan4 4isms i2so iso5mer is1p is2pi is4py 4is1s is4sal issen4 is4ses is4ta. is1te is1ti ist4ly 4istral i2su is5us 4ita. ita4bi i4tag 4ita5m i3tan i3tat 2ite it3era i5teri it4es 2ith i1ti 4itia 4i2tic it3ica 5i5tick it3ig it5ill i2tim 2itio 4itis i4tism i2t5o5m 4iton i4tram it5ry 4itt it3uat i5tud it3ul 4itz. i1u 2iv iv3ell iv3en. i4v3er. i4vers. iv5il. iv5io iv1it i5vore iv3o3ro i4v3ot 4i5w ix4o 4iy 4izar izi4 5izont 5ja jac4q ja4p 1je jer5s 4jestie 4jesty jew3 jo4p 5judg 3ka. k3ab k5ag kais4 kal4 k1b k2ed 1kee ke4g ke5li k3en4d k1er kes4 k3est. ke4ty k3f kh4 k1i 5ki. 5k2ic k4ill kilo5 k4im k4in. kin4de k5iness kin4g ki4p kis4 k5ish kk4 k1l 4kley 4kly k1m k5nes 1k2no ko5r kosh4 k3ou kro5n 4k1s2 k4sc ks4l k4sy k5t k1w lab3ic l4abo laci4 l4ade la3dy lag4n lam3o 3land lan4dl lan5et lan4te lar4g lar3i las4e la5tan 4lateli 4lativ 4lav la4v4a 2l1b lbin4 4l1c2 lce4 l3ci 2ld l2de ld4ere ld4eri ldi4 ld5is l3dr l4dri le2a le4bi left5 5leg. 5legg le4mat lem5atic 4len. 3lenc 5lene. 1lent le3ph le4pr lera5b ler4e 3lerg 3l4eri l4ero les2 le5sco 5lesq 3less 5less. l3eva lev4er. lev4era lev4ers 3ley 4leye 2lf l5fr 4l1g4 l5ga lgar3 l4ges lgo3 2l3h li4ag li2am liar5iz li4as li4ato li5bi 5licio li4cor 4lics 4lict. l4icu l3icy l3ida lid5er 3lidi lif3er l4iff li4fl 5ligate 3ligh li4gra 3lik 4l4i4l lim4bl lim3i li4mo l4im4p l4ina 1l4ine lin3ea lin3i link5er li5og 4l4iq lis4p l1it l2it. 5litica l5i5tics liv3er l1iz 4lj lka3 l3kal lka4t l1l l4law l2le l5lea l3lec l3leg l3lel l3le4n l3le4t ll2i l2lin4 l5lina ll4o lloqui5 ll5out l5low 2lm l5met lm3ing l4mod lmon4 2l1n2 3lo. lob5al lo4ci 4lof 3logic l5ogo 3logu lom3er 5long lon4i l3o3niz lood5 5lope. lop3i l3opm lora4 lo4rato lo5rie lor5ou 5los. los5et 5losophiz 5losophy los4t lo4ta loun5d 2lout 4lov 2lp lpa5b l3pha l5phi lp5ing l3pit l4pl l5pr 4l1r 2l1s2 l4sc l2se l4sie 4lt lt5ag ltane5 l1te lten4 ltera4 lth3i l5ties. ltis4 l1tr ltu2 ltur3a lu5a lu3br luch4 lu3ci lu3en luf4 lu5id lu4ma 5lumi l5umn. 5lumnia lu3o luo3r 4lup luss4 lus3te 1lut l5ven l5vet4 2l1w 1ly 4lya 4lyb ly5me ly3no 2lys4 l5yse 1ma 2mab ma2ca ma5chine ma4cl mag5in 5magn 2mah maid5 4mald ma3lig ma5lin mal4li mal4ty 5mania man5is man3iz 4map ma5rine. ma5riz mar4ly mar3v ma5sce mas4e mas1t 5mate math3 ma3tis 4matiza 4m1b mba4t5 m5bil m4b3ing mbi4v 4m5c 4me. 2med 4med. 5media me3die m5e5dy me2g mel5on mel4t me2m mem1o3 1men men4a men5ac men4de 4mene men4i mens4 mensu5 3ment men4te me5on m5ersa 2mes 3mesti me4ta met3al me1te me5thi m4etr 5metric me5trie me3try me4v 4m1f 2mh 5mi. mi3a mid4a mid4g mig4 3milia m5i5lie m4ill min4a 3mind m5inee m4ingl min5gli m5ingly min4t m4inu miot4 m2is mis4er. mis5l mis4ti m5istry 4mith m2iz 4mk 4m1l m1m mma5ry 4m1n mn4a m4nin mn4o 1mo 4mocr 5mocratiz mo2d1 mo4go mois2 moi5se 4mok mo5lest mo3me mon5et mon5ge moni3a mon4ism mon4ist mo3niz monol4 mo3ny. mo2r 4mora. mos2 mo5sey mo3sp moth3 m5ouf 3mous mo2v 4m1p mpara5 mpa5rab mpar5i m3pet mphas4 m2pi mpi4a mp5ies m4p1in m5pir mp5is mpo3ri mpos5ite m4pous mpov5 mp4tr m2py 4m3r 4m1s2 m4sh m5si 4mt 1mu mula5r4 5mult multi3 3mum mun2 4mup mu4u 4mw 1na 2n1a2b n4abu 4nac. na4ca n5act nag5er. nak4 na4li na5lia 4nalt na5mit n2an nanci4 nan4it nank4 nar3c 4nare nar3i nar4l n5arm n4as nas4c nas5ti n2at na3tal nato5miz n2au nau3se 3naut nav4e 4n1b4 ncar5 n4ces. n3cha n5cheo n5chil n3chis nc1in nc4it ncour5a n1cr n1cu n4dai n5dan n1de nd5est. ndi4b n5d2if n1dit n3diz n5duc ndu4r nd2we 2ne. n3ear ne2b neb3u ne2c 5neck 2ned ne4gat neg5ativ 5nege ne4la nel5iz ne5mi ne4mo 1nen 4nene 3neo ne4po ne2q n1er nera5b n4erar n2ere n4er5i ner4r 1nes 2nes. 4nesp 2nest 4nesw 3netic ne4v n5eve ne4w n3f n4gab n3gel nge4n4e n5gere n3geri ng5ha n3gib ng1in n5git n4gla ngov4 ng5sh n1gu n4gum n2gy 4n1h4 nha4 nhab3 nhe4 3n4ia ni3an ni4ap ni3ba ni4bl ni4d ni5di ni4er ni2fi ni5ficat n5igr nik4 n1im ni3miz n1in 5nine. nin4g ni4o 5nis. nis4ta n2it n4ith 3nitio n3itor ni3tr n1j 4nk2 n5kero n3ket nk3in n1kl 4n1l n5m nme4 nmet4 4n1n2 nne4 nni3al nni4v nob4l no3ble n5ocl 4n3o2d 3noe 4nog noge4 nois5i no5l4i 5nologis 3nomic n5o5miz no4mo no3my no4n non4ag non5i n5oniz 4nop 5nop5o5li nor5ab no4rary 4nosc nos4e nos5t no5ta 1nou 3noun nov3el3 nowl3 n1p4 npi4 npre4c n1q n1r nru4 2n1s2 ns5ab nsati4 ns4c n2se n4s3es nsid1 nsig4 n2sl ns3m n4soc ns4pe n5spi nsta5bl n1t nta4b nter3s nt2i n5tib nti4er nti2f n3tine n4t3ing nti4p ntrol5li nt4s ntu3me nu1a nu4d nu5en nuf4fe n3uin 3nu3it n4um nu1me n5umi 3nu4n n3uo nu3tr n1v2 n1w4 nym4 nyp4 4nz n3za 4oa oad3 o5a5les oard3 oas4e oast5e oat5i ob3a3b o5bar obe4l o1bi o2bin ob5ing o3br ob3ul o1ce och4 o3chet ocif3 o4cil o4clam o4cod oc3rac oc5ratiz ocre3 5ocrit octor5a oc3ula o5cure od5ded od3ic odi3o o2do4 odor3 od5uct. od5ucts o4el o5eng o3er oe4ta o3ev o2fi of5ite ofit4t o2g5a5r og5ativ o4gato o1ge o5gene o5geo o4ger o3gie 1o1gis og3it o4gl o5g2ly 3ogniz o4gro ogu5i 1ogy 2ogyn o1h2 ohab5 oi2 oic3es oi3der oiff4 oig4 oi5let o3ing oint5er o5ism oi5son oist5en oi3ter o5j 2ok o3ken ok5ie o1la o4lan olass4 ol2d old1e ol3er o3lesc o3let ol4fi ol2i o3lia o3lice ol5id. o3li4f o5lil ol3ing o5lio o5lis. ol3ish o5lite o5litio o5liv olli4e ol5ogiz olo4r ol5pl ol2t ol3ub ol3ume ol3un o5lus ol2v o2ly om5ah oma5l om5atiz om2be om4bl o2me om3ena om5erse o4met om5etry o3mia om3ic. om3ica o5mid om1in o5mini 5ommend omo4ge o4mon om3pi ompro5 o2n on1a on4ac o3nan on1c 3oncil 2ond on5do o3nen on5est on4gu on1ic o3nio on1is o5niu on3key on4odi on3omy on3s onspi4 onspir5a onsu4 onten4 on3t4i ontif5 on5um onva5 oo2 ood5e ood5i oo4k oop3i o3ord oost5 o2pa ope5d op1er 3opera 4operag 2oph o5phan o5pher op3ing o3pit o5pon o4posi o1pr op1u opy5 o1q o1ra o5ra. o4r3ag or5aliz or5ange ore5a o5real or3ei ore5sh or5est. orew4 or4gu 4o5ria or3ica o5ril or1in o1rio or3ity o3riu or2mi orn2e o5rof or3oug or5pe 3orrh or4se ors5en orst4 or3thi or3thy or4ty o5rum o1ry os3al os2c os4ce o3scop 4oscopi o5scr os4i4e os5itiv os3ito os3ity osi4u os4l o2so os4pa os4po os2ta o5stati os5til os5tit o4tan otele4g ot3er. ot5ers o4tes 4oth oth5esi oth3i4 ot3ic. ot5ica o3tice o3tif o3tis oto5s ou2 ou3bl ouch5i ou5et ou4l ounc5er oun2d ou5v ov4en over4ne over3s ov4ert o3vis oviti4 o5v4ol ow3der ow3el ow5est ow1i own5i o4wo oy1a 1pa pa4ca pa4ce pac4t p4ad 5pagan p3agat p4ai pain4 p4al pan4a pan3el pan4ty pa3ny pa1p pa4pu para5bl par5age par5di 3pare par5el p4a4ri par4is pa2te pa5ter 5pathic pa5thy pa4tric pav4 3pay 4p1b pd4 4pe. 3pe4a pear4l pe2c 2p2ed 3pede 3pedi pedia4 ped4ic p4ee pee4d pek4 pe4la peli4e pe4nan p4enc pen4th pe5on p4era. pera5bl p4erag p4eri peri5st per4mal perme5 p4ern per3o per3ti pe5ru per1v pe2t pe5ten pe5tiz 4pf 4pg 4ph. phar5i phe3no ph4er ph4es. ph1ic 5phie ph5ing 5phisti 3phiz ph2l 3phob 3phone 5phoni pho4r 4phs ph3t 5phu 1phy pi3a pian4 pi4cie pi4cy p4id p5ida pi3de 5pidi 3piec pi3en pi4grap pi3lo pi2n p4in. pind4 p4ino 3pi1o pion4 p3ith pi5tha pi2tu 2p3k2 1p2l2 3plan plas5t pli3a pli5er 4plig pli4n ploi4 plu4m plum4b 4p1m 2p3n po4c 5pod. po5em po3et5 5po4g poin2 5point poly5t po4ni po4p 1p4or po4ry 1pos pos1s p4ot po4ta 5poun 4p1p ppa5ra p2pe p4ped p5pel p3pen p3per p3pet ppo5site pr2 pray4e 5preci pre5co pre3em pref5ac pre4la pre3r p3rese 3press pre5ten pre3v 5pri4e prin4t3 pri4s pris3o p3roca prof5it pro3l pros3e pro1t 2p1s2 p2se ps4h p4sib 2p1t pt5a4b p2te p2th pti3m ptu4r p4tw pub3 pue4 puf4 pul3c pu4m pu2n pur4r 5pus pu2t 5pute put3er pu3tr put4ted put4tin p3w qu2 qua5v 2que. 3quer 3quet 2rab ra3bi rach4e r5acl raf5fi raf4t r2ai ra4lo ram3et r2ami rane5o ran4ge r4ani ra5no rap3er 3raphy rar5c rare4 rar5ef 4raril r2as ration4 rau4t ra5vai rav3el ra5zie r1b r4bab r4bag rbi2 rbi4f r2bin r5bine rb5ing. rb4o r1c r2ce rcen4 r3cha rch4er r4ci4b rc4it rcum3 r4dal rd2i rdi4a rdi4er rdin4 rd3ing 2re. re1al re3an re5arr 5reav re4aw r5ebrat rec5oll rec5ompe re4cre 2r2ed re1de re3dis red5it re4fac re2fe re5fer. re3fi re4fy reg3is re5it re1li re5lu r4en4ta ren4te re1o re5pin re4posi re1pu r1er4 r4eri rero4 re5ru r4es. re4spi ress5ib res2t re5stal re3str re4ter re4ti4z re3tri reu2 re5uti rev2 re4val rev3el r5ev5er. re5vers re5vert re5vil rev5olu re4wh r1f rfu4 r4fy rg2 rg3er r3get r3gic rgi4n rg3ing r5gis r5git r1gl rgo4n r3gu rh4 4rh. 4rhal ri3a ria4b ri4ag r4ib rib3a ric5as r4ice 4rici 5ricid ri4cie r4ico rid5er ri3enc ri3ent ri1er ri5et rig5an 5rigi ril3iz 5riman rim5i 3rimo rim4pe r2ina 5rina. rin4d rin4e rin4g ri1o 5riph riph5e ri2pl rip5lic r4iq r2is r4is. ris4c r3ish ris4p ri3ta3b r5ited. rit5er. rit5ers rit3ic ri2tu rit5ur riv5el riv3et riv3i r3j r3ket rk4le rk4lin r1l rle4 r2led r4lig r4lis rl5ish r3lo4 r1m rma5c r2me r3men rm5ers rm3ing r4ming. r4mio r3mit r4my r4nar r3nel r4ner r5net r3ney r5nic r1nis4 r3nit r3niv rno4 r4nou r3nu rob3l r2oc ro3cr ro4e ro1fe ro5fil rok2 ro5ker 5role. rom5ete rom4i rom4p ron4al ron4e ro5n4is ron4ta 1room 5root ro3pel rop3ic ror3i ro5ro ros5per ros4s ro4the ro4ty ro4va rov5el rox5 r1p r4pea r5pent rp5er. r3pet rp4h4 rp3ing r3po r1r4 rre4c rre4f r4reo rre4st rri4o rri4v rron4 rros4 rrys4 4rs2 r1sa rsa5ti rs4c r2se r3sec rse4cr rs5er. rs3es rse5v2 r1sh r5sha r1si r4si4b rson3 r1sp r5sw rtach4 r4tag r3teb rten4d rte5o r1ti rt5ib rti4d r4tier r3tig rtil3i rtil4l r4tily r4tist r4tiv r3tri rtroph4 rt4sh ru3a ru3e4l ru3en ru4gl ru3in rum3pl ru2n runk5 run4ty r5usc ruti5n rv4e rvel4i r3ven rv5er. r5vest r3vey r3vic rvi4v r3vo r1w ry4c 5rynge ry3t sa2 2s1ab 5sack sac3ri s3act 5sai salar4 sal4m sa5lo sal4t 3sanc san4de s1ap sa5ta 5sa3tio sat3u sau4 sa5vor 5saw 4s5b scan4t5 sca4p scav5 s4ced 4scei s4ces sch2 s4cho 3s4cie 5scin4d scle5 s4cli scof4 4scopy scour5a s1cu 4s5d 4se. se4a seas4 sea5w se2c3o 3sect 4s4ed se4d4e s5edl se2g seg3r 5sei se1le 5self 5selv 4seme se4mol sen5at 4senc sen4d s5ened sen5g s5enin 4sentd 4sentl sep3a3 4s1er. s4erl ser4o 4servo s1e4s se5sh ses5t 5se5um 5sev sev3en sew4i 5sex 4s3f 2s3g s2h 2sh. sh1er 5shev sh1in sh3io 3ship shiv5 sho4 sh5old shon3 shor4 short5 4shw si1b s5icc 3side. 5sides 5sidi si5diz 4signa sil4e 4sily 2s1in s2ina 5sine. s3ing 1sio 5sion sion5a si2r sir5a 1sis 3sitio 5siu 1siv 5siz sk2 4ske s3ket sk5ine sk5ing s1l2 s3lat s2le slith5 2s1m s3ma small3 sman3 smel4 s5men 5smith smol5d4 s1n4 1so so4ce soft3 so4lab sol3d2 so3lic 5solv 3som 3s4on. sona4 son4g s4op 5sophic s5ophiz s5ophy sor5c sor5d 4sov so5vi 2spa 5spai spa4n spen4d 2s5peo 2sper s2phe 3spher spho5 spil4 sp5ing 4spio s4ply s4pon spor4 4spot squal4l s1r 2ss s1sa ssas3 s2s5c s3sel s5seng s4ses. s5set s1si s4sie ssi4er ss5ily s4sl ss4li s4sn sspend4 ss2t ssur5a ss5w 2st. s2tag s2tal stam4i 5stand s4ta4p 5stat. s4ted stern5i s5tero ste2w stew5a s3the st2i s4ti. s5tia s1tic 5stick s4tie s3tif st3ing 5stir s1tle 5stock stom3a 5stone s4top 3store st4r s4trad 5stratu s4tray s4trid 4stry 4st3w s2ty 1su su1al su4b3 su2g3 su5is suit3 s4ul su2m sum3i su2n su2r 4sv sw2 4swo s4y 4syc 3syl syn5o sy5rin 1ta 3ta. 2tab ta5bles 5taboliz 4taci ta5do 4taf4 tai5lo ta2l ta5la tal5en tal3i 4talk tal4lis ta5log ta5mo tan4de tanta3 ta5per ta5pl tar4a 4tarc 4tare ta3riz tas4e ta5sy 4tatic ta4tur taun4 tav4 2taw tax4is 2t1b 4tc t4ch tch5et 4t1d 4te. tead4i 4teat tece4 5tect 2t1ed te5di 1tee teg4 te5ger te5gi 3tel. teli4 5tels te2ma2 tem3at 3tenan 3tenc 3tend 4tenes 1tent ten4tag 1teo te4p te5pe ter3c 5ter3d 1teri ter5ies ter3is teri5za 5ternit ter5v 4tes. 4tess t3ess. teth5e 3teu 3tex 4tey 2t1f 4t1g 2th. than4 th2e 4thea th3eas the5at the3is 3thet th5ic. th5ica 4thil 5think 4thl th5ode 5thodic 4thoo thor5it tho5riz 2ths 1tia ti4ab ti4ato 2ti2b 4tick t4ico t4ic1u 5tidi 3tien tif2 ti5fy 2tig 5tigu till5in 1tim 4timp tim5ul 2t1in t2ina 3tine. 3tini 1tio ti5oc tion5ee 5tiq ti3sa 3tise tis4m ti5so tis4p 5tistica ti3tl ti4u 1tiv tiv4a 1tiz ti3za ti3zen 2tl t5la tlan4 3tle. 3tled 3tles. t5let. t5lo 4t1m tme4 2t1n2 1to to3b to5crat 4todo 2tof to2gr to5ic to2ma tom4b to3my ton4ali to3nat 4tono 4tony to2ra to3rie tor5iz tos2 5tour 4tout to3war 4t1p 1tra tra3b tra5ch traci4 trac4it trac4te tras4 tra5ven trav5es5 tre5f tre4m trem5i 5tria tri5ces 5tricia 4trics 2trim tri4v tro5mi tron5i 4trony tro5phe tro3sp tro3v tru5i trus4 4t1s2 t4sc tsh4 t4sw 4t3t2 t4tes t5to ttu4 1tu tu1a tu3ar tu4bi tud2 4tue 4tuf4 5tu3i 3tum tu4nis 2t3up. 3ture 5turi tur3is tur5o tu5ry 3tus 4tv tw4 4t1wa twis4 4two 1ty 4tya 2tyl type3 ty5ph 4tz tz4e 4uab uac4 ua5na uan4i uar5ant uar2d uar3i uar3t u1at uav4 ub4e u4bel u3ber u4bero u1b4i u4b5ing u3ble. u3ca uci4b uc4it ucle3 u3cr u3cu u4cy ud5d ud3er ud5est udev4 u1dic ud3ied ud3ies ud5is u5dit u4don ud4si u4du u4ene uens4 uen4te uer4il 3ufa u3fl ugh3en ug5in 2ui2 uil5iz ui4n u1ing uir4m uita4 uiv3 uiv4er. u5j 4uk u1la ula5b u5lati ulch4 5ulche ul3der ul4e u1len ul4gi ul2i u5lia ul3ing ul5ish ul4lar ul4li4b ul4lis 4ul3m u1l4o 4uls uls5es ul1ti ultra3 4ultu u3lu ul5ul ul5v um5ab um4bi um4bly u1mi u4m3ing umor5o um2p unat4 u2ne un4er u1ni un4im u2nin un5ish uni3v un3s4 un4sw unt3ab un4ter. un4tes unu4 un5y un5z u4ors u5os u1ou u1pe uper5s u5pia up3ing u3pl up3p upport5 upt5ib uptu4 u1ra 4ura. u4rag u4ras ur4be urc4 ur1d ure5at ur4fer ur4fr u3rif uri4fic ur1in u3rio u1rit ur3iz ur2l url5ing. ur4no uros4 ur4pe ur4pi urs5er ur5tes ur3the urti4 ur4tie u3ru 2us u5sad u5san us4ap usc2 us3ci use5a u5sia u3sic us4lin us1p us5sl us5tere us1tr u2su usur4 uta4b u3tat 4ute. 4utel 4uten uten4i 4u1t2i uti5liz u3tine ut3ing ution5a u4tis 5u5tiz u4t1l ut5of uto5g uto5matic u5ton u4tou uts4 u3u uu4m u1v2 uxu3 uz4e 1va 5va. 2v1a4b vac5il vac3u vag4 va4ge va5lie val5o val1u va5mo va5niz va5pi var5ied 3vat 4ve. 4ved veg3 v3el. vel3li ve4lo v4ely ven3om v5enue v4erd 5vere. v4erel v3eren ver5enc v4eres ver3ie vermi4n 3verse ver3th v4e2s 4ves. ves4te ve4te vet3er ve4ty vi5ali 5vian 5vide. 5vided 4v3iden 5vides 5vidi v3if vi5gn vik4 2vil 5vilit v3i3liz v1in 4vi4na v2inc vin5d 4ving vio3l v3io4r vi1ou vi4p vi5ro vis3it vi3so vi3su 4viti vit3r 4vity 3viv 5vo. voi4 3vok vo4la v5ole 5volt 3volv vom5i vor5ab vori4 vo4ry vo4ta 4votee 4vv4 v4y w5abl 2wac wa5ger wag5o wait5 w5al. wam4 war4t was4t wa1te wa5ver w1b wea5rie weath3 wed4n weet3 wee5v wel4l w1er west3 w3ev whi4 wi2 wil2 will5in win4de win4g wir4 3wise with3 wiz5 w4k wl4es wl3in w4no 1wo2 wom1 wo5ven w5p wra4 wri4 writa4 w3sh ws4l ws4pe w5s4t 4wt wy4 x1a xac5e x4ago xam3 x4ap xas5 x3c2 x1e xe4cuto x2ed xer4i xe5ro x1h xhi2 xhil5 xhu4 x3i xi5a xi5c xi5di x4ime xi5miz x3o x4ob x3p xpan4d xpecto5 xpe3d x1t2 x3ti x1u xu3a xx4 y5ac 3yar4 y5at y1b y1c y2ce yc5er y3ch ych4e ycom4 ycot4 y1d y5ee y1er y4erf yes4 ye4t y5gi 4y3h y1i y3la ylla5bl y3lo y5lu ymbol5 yme4 ympa3 yn3chr yn5d yn5g yn5ic 5ynx y1o4 yo5d y4o5g yom4 yo5net y4ons y4os y4ped yper5 yp3i y3po y4poc yp2ta y5pu yra5m yr5ia y3ro yr4r ys4c y3s2e ys3ica ys3io 3ysis y4so yss4 ys1t ys3ta ysur4 y3thin yt3ic y1w za1 z5a2b zar2 4zb 2ze ze4n ze4p z1er ze3ro zet4 2z1i z4il z4is 5zl 4zm 1zo zo4m zo5ol zte4 4z1z2 z4zy \x00")

// en_exceptions - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyen.h:406
var en_exceptions []byte = []byte("as-so-ciate as-so-ciates dec-li-na-tion oblig-a-tory phil-an-thropic present presents project projects reci-procity re-cog-ni-zance ref-or-ma-tion ret-ri-bu-tion ta-ble \x00")

// hwword - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:15
// hyphenation
// the hyphenation dictionary (.hw)
// buffer for .hw words
var hwword []byte = make([]byte, 262144)

// hwhyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:16
// buffer for .hw hyphenations
var hwhyph []byte = make([]byte, 262144)

// hwword_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:17
// used hwword[] length
var hwword_len int32

// hwdict - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:18
// map words to their index in hwoff[]
var hwdict []dict

// hwoff - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:19
// the offset of words in hwword[]
var hwoff []int32 = make([]int32, 16384)

// hw_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:20
// the number of dictionary words
var hw_n int32

// hy_cget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:23
func hy_cget(d []byte, s []byte) int32 {
	// read a single character from s into d; return the number of characters read
	var i int32
	if int32(s[0]) != int32('\\') {
		return utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], d)
	}
	if int32(s[1]) == int32('[') {
		s = s[0+2:]
		for int32(s[0]) != 0 && int32(s[0]) != int32(']') && i < 32-1 {
			d[func() int32 {
				defer func() {
					i++
				}()
				return i
			}()] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		d[i] = '\x00'
		if int32(s[0]) != 0 {
			return i + 3
		}
		return i + 2
	}
	if int32(s[1]) == int32('(') {
		s = s[0+2:]
		i += utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], d[0+i:])
		i += utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], d[0+i:])
		return 2 + i
	}
	if int32(s[1]) == int32('C') {
		var q int32 = int32(s[2])
		s = s[0+3:]
		for int32(s[0]) != 0 && int32(s[0]) != q && i < 32-1 {
			d[func() int32 {
				defer func() {
					i++
				}()
				return i
			}()] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		d[i] = '\x00'
		if int32(s[0]) != 0 {
			return i + 4
		}
		return i + 3
	}
	(func() []byte {
		defer func() {
			d = d[0+1:]
		}()
		return d
	}())[0] = (func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]
	return 1 + utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], d)
}

// hy_cput - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:54
func hy_cput(d []byte, s []byte) int32 {
	if noarch.Not(s[0]) || noarch.Not(s[1]) || utf8one(s) != 0 {
		// append character s to d; return the number of characters written
		noarch.Strcpy(d, s)
	} else if int32(s[0]) == int32('\\') {
		noarch.Strcpy(d, s)
	} else if noarch.Not(s[2]) {
		noarch.Snprintf(d, 32, []byte("\\[%s]\x00"), s)
	}
	return noarch.Strlen(d)
}

// hw_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:66
func hw_add(s []byte) {
	// insert word s into hwword[] and hwhyph[]
	var p []byte = hwword[0+hwword_len:]
	var n []byte = hwhyph[0+hwword_len:]
	var len_ int32 = noarch.Strlen(s) + int32(1)
	var i int32
	var c int32
	if hw_n == 16384 || uint32(hwword_len+len_) > 262144 {
		return
	}
	noarch.Memset(n, byte(0), uint32(len_))
	for (func() int32 {
		c = int32(uint8((func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]))
		return c
	}()) != 0 {
		if c == int32('-') {
			n[i] = byte(1)
		} else {
			p[func() int32 {
				defer func() {
					i++
				}()
				return i
			}()] = byte(c)
		}
	}
	p[i] = '\x00'
	hwoff[hw_n] = hwword_len
	dict_put(hwdict, hwword[0+hwoff[hw_n]:], hw_n)
	hwword_len += i + 1
	hw_n++
}

// hw_lookup - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:88
func hw_lookup(word []byte, hyph []byte) int32 {
	var word2 []byte = []byte{byte(0), '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}
	var hyph2 []byte
	var map_ []int32 = []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var off int32
	var i int32
	var j int32
	var idx int32 = -1
	hcode_strcpy(word2, word, map_, 0)
	for int32(word2[off]) == int32('.') {
		// skip unknown characters at the front
		off++
	}
	i = dict_prefix(hwdict, word2[0+off:], c4goUnsafeConvert_int32(&idx))
	if i < 0 {
		return 1
	}
	hyph2 = hwhyph[0+hwoff[i]:]
	for j = 0; word2[j+off] != 0; j++ {
		if hyph2[j] != 0 {
			hyph[map_[j+off]] = hyph2[j]
		}
	}
	return 0
}

// tr_hw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:108
func tr_hw(args [][]byte) {
	var word []byte = make([]byte, 256)
	var c []byte
	var i int32
	for i = 1; i < 32 && args[i] != nil; i++ {
		var s []byte = args[i]
		var d []byte = word
		for int32((int64(uintptr(unsafe.Pointer(&d[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&word[0])))/int64(1))) < 256-32 && noarch.Not(escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&c))[:])) {
			if noarch.Strcmp([]byte("-\x00"), c) != 0 {
				hcode_mapchar(c)
			}
			d = d[0+hy_cput(d, c):]
		}
		hw_add(word)
	}
}

// hyinit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:127
// the tex hyphenation algorithm
// hyphenation data initialized
var hyinit int32

// hypats - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:128
// hyphenation patterns
var hypats []byte = make([]byte, 262144)

// hynums - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:129
// hyphenation pattern numbers
var hynums []byte = make([]byte, 262144)

// hypats_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:130
// used hypats[] and hynums[] length
var hypats_len int32

// hydict - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:131
// map patterns to their index in hyoff[]
var hydict []dict

// hyoff - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:132
// the offset of this pattern in hypats[]
var hyoff []int32 = make([]int32, 16384)

// hy_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:133
// the number of patterns
var hy_n int32

// hy_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:136
func hy_find(s []byte, n []byte) {
	// find the patterns matching s and update hyphenation values in n
	var plen int32
	var p []byte
	var np []byte
	var i int32
	var j int32
	var idx int32 = -1
	for (func() int32 {
		i = dict_prefix(hydict, s, c4goUnsafeConvert_int32(&idx))
		return i
	}()) >= 0 {
		p = hypats[0+hyoff[i]:]
		np = c4goPointerArithByteSlice(hynums, int(int32((int64(uintptr(unsafe.Pointer(&p[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&hypats[0])))/int64(1)))))
		plen = noarch.Strlen(p) + int32(1)
		for j = 0; j < plen; j++ {
			if int32(n[j]) < int32(np[j]) {
				n[j] = np[j]
			}
		}
	}
}

// hy_dohyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:153
func hy_dohyph(hyph []byte, word []byte, flg int32) {
	// mark the hyphenation points of word in hyph
	// cleaned-up word[]; "Abc" -> ".abc."
	var w []byte = []byte{byte(0), '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}
	// the hyphenation value for w[]
	var n []byte = []byte{byte(0), '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}
	// start of the i-th character in w
	var c []int32 = make([]int32, 256)
	// w[i] corresponds to word[wmap[i]]
	var wmap []int32 = []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var ch []byte = make([]byte, 32)
	var nc int32
	var i int32
	var wlen int32
	hcode_strcpy(w, word, wmap, 1)
	wlen = noarch.Strlen(w)
	for i = 0; i < wlen-1; i += hy_cget(ch, w[0+i:]) {
		c[func() int32 {
			defer func() {
				nc++
			}()
			return nc
		}()] = i
	}
	for i = 0; i < nc-1; i++ {
		hy_find(w[0+c[i]:], n[0+c[i]:])
	}
	noarch.Memset(hyph, byte(0), uint32(wlen)*1)
	for i = 3; i < nc-2; i++ {
		if int32(n[c[i]])%2 != 0 && int32(w[c[i-1]]) != int32('.') && int32(w[c[i]]) != int32('.') && int32(w[c[i-2]]) != int32('.') && int32(w[c[i+1]]) != int32('.') && (^flg&4 != 0 || int32(w[c[i+2]]) != int32('.')) && (^flg&8 != 0 || int32(w[c[i-3]]) != int32('.')) {
			hyph[wmap[c[i]]] = byte(1)
		}
	}
}

// hy_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:178
func hy_add(s []byte) {
	// insert pattern s into hypats[] and hynums[]
	var p []byte = hypats[0+hypats_len:]
	var n []byte = hynums[0+hypats_len:]
	var len_ int32 = noarch.Strlen(s) + int32(1)
	var i int32
	var c int32
	if hy_n >= 16384 || uint32(hypats_len+len_) >= 262144 {
		return
	}
	noarch.Memset(n, byte(0), uint32(len_))
	for (func() int32 {
		c = int32(uint8((func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]))
		return c
	}()) != 0 {
		if c >= int32('0') && c <= int32('9') {
			n[i] = byte(c - int32('0'))
		} else {
			p[func() int32 {
				defer func() {
					i++
				}()
				return i
			}()] = byte(c)
		}
	}
	p[i] = '\x00'
	hyoff[hy_n] = hypats_len
	dict_put(hydict, hypats[0+hyoff[hy_n]:], hy_n)
	hypats_len += i + 1
	hy_n++
}

// hcodedict - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:201
// .hcode request
var hcodedict []dict

// hcodesrc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:202
var hcodesrc [][]byte = make([][]byte, 512)

// hcodedst - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:203
var hcodedst [][]byte = make([][]byte, 512)

// hcode_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:204
var hcode_n int32

// hcode_mapchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:207
func hcode_mapchar(s []byte) int32 {
	// replace the character in s after .hcode mapping; returns s's new length
	var i int32 = dict_get(hcodedict, s)
	if i >= 0 {
		noarch.Strcpy(s, hcodedst[i])
	} else if noarch.Not(s[1]) {
		s[0] = byte(func() int32 {
			if int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISalpha)) != 0 {
				return tolower(int32(uint8(s[0])))
			}
			return int32('.')
		}())
	}
	return noarch.Strlen(s)
}

// hcode_strcpy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:218
func hcode_strcpy(d []byte, s []byte, map_ []int32, dots int32) {
	// copy s to d after .hcode mappings; s[map[j]] corresponds to d[j]
	var c []byte = make([]byte, 32)
	var di int32
	var si int32
	if dots != 0 {
		d[func() int32 {
			defer func() {
				di++
			}()
			return di
		}()] = '.'
	}
	for di < 256-32 && int32(s[si]) != 0 {
		map_[di] = si
		si += hy_cget(c, s[0+si:])
		hcode_mapchar(c)
		di += hy_cput(d[0+di:], c)
	}
	if dots != 0 {
		d[func() int32 {
			defer func() {
				di++
			}()
			return di
		}()] = '.'
	}
	d[di] = '\x00'
}

// hcode_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:235
func hcode_add(c1 []byte, c2 []byte) {
	var i int32 = dict_get(hcodedict, c1)
	if i >= 0 {
		noarch.Strcpy(hcodedst[i], c2)
	} else if hcode_n < 512 {
		noarch.Strcpy(hcodesrc[hcode_n], c1)
		noarch.Strcpy(hcodedst[hcode_n], c2)
		dict_put(hcodedict, hcodesrc[hcode_n], hcode_n)
		hcode_n++
	}
}

// tr_hcode - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:248
func tr_hcode(args [][]byte) {
	var c1 []byte = make([]byte, 32)
	var c2 []byte = make([]byte, 32)
	var s []byte = args[1]
	for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], c1) >= 0 && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], c2) >= 0 {
		hcode_add(c1, c2)
	}
}

// hyph_readpatterns - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:256
func hyph_readpatterns(s []byte) {
	var word []byte = make([]byte, 256)
	var d []byte
	for s[0] != 0 {
		d = word
		for int32(s[0]) != 0 && noarch.Not(int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace))) {
			(func() []byte {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		d[0] = '\x00'
		hy_add(word)
		for int32(s[0]) != 0 && int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
			s = s[0+1:]
		}
	}
}

// hyph_readexceptions - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:271
func hyph_readexceptions(s []byte) {
	var word []byte = make([]byte, 256)
	var d []byte
	for s[0] != 0 {
		d = word
		for int32(s[0]) != 0 && noarch.Not(int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace))) {
			(func() []byte {
				defer func() {
					d = d[0+1:]
				}()
				return d
			}())[0] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		d[0] = '\x00'
		hw_add(word)
		for int32(s[0]) != 0 && int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
			s = s[0+1:]
		}
	}
}

// hyphenate - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:286
func hyphenate(hyph []byte, word []byte, flg int32) {
	if noarch.Not(hyinit) {
		hyinit = 1
		hyph_readpatterns(en_patterns)
		hyph_readexceptions(en_exceptions)
	}
	if hw_lookup(word, hyph) != 0 {
		hy_dohyph(hyph, word, flg)
	}
}

// hycase - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:298
// lowercase-uppercase character mapping
var hycase [][][]byte = [][][]byte{{[]byte("a\x00"), []byte("A\x00")}, {[]byte("á\x00"), []byte("Á\x00")}, {[]byte("à\x00"), []byte("À\x00")}, {[]byte("ă\x00"), []byte("Ă\x00")}, {[]byte("â\x00"), []byte("Â\x00")}, {[]byte("ǎ\x00"), []byte("Ǎ\x00")}, {[]byte("å\x00"), []byte("Å\x00")}, {[]byte("ä\x00"), []byte("Ä\x00")}, {[]byte("ã\x00"), []byte("Ã\x00")}, {[]byte("ą\x00"), []byte("Ą\x00")}, {[]byte("ā\x00"), []byte("Ā\x00")}, {[]byte("æ\x00"), []byte("Æ\x00")}, {[]byte("ǽ\x00"), []byte("Ǽ\x00")}, {[]byte("b\x00"), []byte("B\x00")}, {[]byte("c\x00"), []byte("C\x00")}, {[]byte("ć\x00"), []byte("Ć\x00")}, {[]byte("ĉ\x00"), []byte("Ĉ\x00")}, {[]byte("č\x00"), []byte("Č\x00")}, {[]byte("ç\x00"), []byte("Ç\x00")}, {[]byte("d\x00"), []byte("D\x00")}, {[]byte("ď\x00"), []byte("Ď\x00")}, {[]byte("đ\x00"), []byte("Đ\x00")}, {[]byte("ḍ\x00"), []byte("Ḍ\x00")}, {[]byte("ð\x00"), []byte("Ð\x00")}, {[]byte("e\x00"), []byte("E\x00")}, {[]byte("é\x00"), []byte("É\x00")}, {[]byte("è\x00"), []byte("È\x00")}, {[]byte("ê\x00"), []byte("Ê\x00")}, {[]byte("ě\x00"), []byte("Ě\x00")}, {[]byte("ë\x00"), []byte("Ë\x00")}, {[]byte("ė\x00"), []byte("Ė\x00")}, {[]byte("ę\x00"), []byte("Ę\x00")}, {[]byte("ē\x00"), []byte("Ē\x00")}, {[]byte("f\x00"), []byte("F\x00")}, {[]byte("g\x00"), []byte("G\x00")}, {[]byte("ğ\x00"), []byte("Ğ\x00")}, {[]byte("ĝ\x00"), []byte("Ĝ\x00")}, {[]byte("ģ\x00"), []byte("Ģ\x00")}, {[]byte("h\x00"), []byte("H\x00")}, {[]byte("ĥ\x00"), []byte("Ĥ\x00")}, {[]byte("ḥ\x00"), []byte("Ḥ\x00")}, {[]byte("ḫ\x00"), []byte("Ḫ\x00")}, {[]byte("i\x00"), []byte("I\x00")}, {[]byte("ı\x00"), []byte("I\x00")}, {[]byte("í\x00"), []byte("Í\x00")}, {[]byte("ì\x00"), []byte("Ì\x00")}, {[]byte("î\x00"), []byte("Î\x00")}, {[]byte("ǐ\x00"), []byte("Ǐ\x00")}, {[]byte("ï\x00"), []byte("Ï\x00")}, {[]byte("į\x00"), []byte("Į\x00")}, {[]byte("ī\x00"), []byte("Ī\x00")}, {[]byte("j\x00"), []byte("J\x00")}, {[]byte("ĵ\x00"), []byte("Ĵ\x00")}, {[]byte("k\x00"), []byte("K\x00")}, {[]byte("ķ\x00"), []byte("Ķ\x00")}, {[]byte("l\x00"), []byte("L\x00")}, {[]byte("ľ\x00"), []byte("Ľ\x00")}, {[]byte("ł\x00"), []byte("Ł\x00")}, {[]byte("ļ\x00"), []byte("Ļ\x00")}, {[]byte("ḷ\x00"), []byte("Ḷ\x00")}, {[]byte("m\x00"), []byte("M\x00")}, {[]byte("ṁ\x00"), []byte("Ṁ\x00")}, {[]byte("ṃ\x00"), []byte("Ṃ\x00")}, {[]byte("n\x00"), []byte("N\x00")}, {[]byte("ń\x00"), []byte("Ń\x00")}, {[]byte("ň\x00"), []byte("Ň\x00")}, {[]byte("ñ\x00"), []byte("Ñ\x00")}, {[]byte("ṅ\x00"), []byte("Ṅ\x00")}, {[]byte("ņ\x00"), []byte("Ņ\x00")}, {[]byte("ṇ\x00"), []byte("Ṇ\x00")}, {[]byte("œ\x00"), []byte("Œ\x00")}, {[]byte("o\x00"), []byte("O\x00")}, {[]byte("ó\x00"), []byte("Ó\x00")}, {[]byte("ò\x00"), []byte("Ò\x00")}, {[]byte("ô\x00"), []byte("Ô\x00")}, {[]byte("ǒ\x00"), []byte("Ǒ\x00")}, {[]byte("ö\x00"), []byte("Ö\x00")}, {[]byte("ő\x00"), []byte("Ő\x00")}, {[]byte("õ\x00"), []byte("Õ\x00")}, {[]byte("ø\x00"), []byte("Ø\x00")}, {[]byte("ō\x00"), []byte("Ō\x00")}, {[]byte("p\x00"), []byte("P\x00")}, {[]byte("q\x00"), []byte("Q\x00")}, {[]byte("r\x00"), []byte("R\x00")}, {[]byte("ŕ\x00"), []byte("Ŕ\x00")}, {[]byte("ř\x00"), []byte("Ř\x00")}, {[]byte("s\x00"), []byte("S\x00")}, {[]byte("ś\x00"), []byte("Ś\x00")}, {[]byte("ŝ\x00"), []byte("Ŝ\x00")}, {[]byte("š\x00"), []byte("Š\x00")}, {[]byte("ş\x00"), []byte("Ş\x00")}, {[]byte("ṣ\x00"), []byte("Ṣ\x00")}, {[]byte("t\x00"), []byte("T\x00")}, {[]byte("ť\x00"), []byte("Ť\x00")}, {[]byte("ț\x00"), []byte("Ț\x00")}, {[]byte("ṭ\x00"), []byte("Ṭ\x00")}, {[]byte("u\x00"), []byte("U\x00")}, {[]byte("ú\x00"), []byte("Ú\x00")}, {[]byte("ù\x00"), []byte("Ù\x00")}, {[]byte("ŭ\x00"), []byte("Ŭ\x00")}, {[]byte("û\x00"), []byte("Û\x00")}, {[]byte("ǔ\x00"), []byte("Ǔ\x00")}, {[]byte("ů\x00"), []byte("Ů\x00")}, {[]byte("ü\x00"), []byte("Ü\x00")}, {[]byte("ǘ\x00"), []byte("Ǘ\x00")}, {[]byte("ǜ\x00"), []byte("Ǜ\x00")}, {[]byte("ǚ\x00"), []byte("Ǚ\x00")}, {[]byte("ǖ\x00"), []byte("Ǖ\x00")}, {[]byte("ű\x00"), []byte("Ű\x00")}, {[]byte("ų\x00"), []byte("Ų\x00")}, {[]byte("ū\x00"), []byte("Ū\x00")}, {[]byte("v\x00"), []byte("V\x00")}, {[]byte("w\x00"), []byte("W\x00")}, {[]byte("x\x00"), []byte("X\x00")}, {[]byte("y\x00"), []byte("Y\x00")}, {[]byte("ý\x00"), []byte("Ý\x00")}, {[]byte("z\x00"), []byte("Z\x00")}, {[]byte("ź\x00"), []byte("Ź\x00")}, {[]byte("ž\x00"), []byte("Ž\x00")}, {[]byte("ż\x00"), []byte("Ż\x00")}, {[]byte("þ\x00"), []byte("Þ\x00")}, {[]byte("α\x00"), []byte("Α\x00")}, {[]byte("ά\x00"), []byte("Ά\x00")}, {[]byte("β\x00"), []byte("Β\x00")}, {[]byte("ϐ\x00"), []byte("Β\x00")}, {[]byte("γ\x00"), []byte("Γ\x00")}, {[]byte("δ\x00"), []byte("Δ\x00")}, {[]byte("ϫ\x00"), []byte("Ϫ\x00")}, {[]byte("ε\x00"), []byte("Ε\x00")}, {[]byte("έ\x00"), []byte("Έ\x00")}, {[]byte("ζ\x00"), []byte("Ζ\x00")}, {[]byte("ϩ\x00"), []byte("Ϩ\x00")}, {[]byte("η\x00"), []byte("Η\x00")}, {[]byte("ή\x00"), []byte("Ή\x00")}, {[]byte("θ\x00"), []byte("Θ\x00")}, {[]byte("ι\x00"), []byte("Ι\x00")}, {[]byte("ί\x00"), []byte("Ί\x00")}, {[]byte("ϊ\x00"), []byte("Ϊ\x00")}, {[]byte("κ\x00"), []byte("Κ\x00")}, {[]byte("ϧ\x00"), []byte("Ϧ\x00")}, {[]byte("λ\x00"), []byte("Λ\x00")}, {[]byte("μ\x00"), []byte("Μ\x00")}, {[]byte("ν\x00"), []byte("Ν\x00")}, {[]byte("ξ\x00"), []byte("Ξ\x00")}, {[]byte("ο\x00"), []byte("Ο\x00")}, {[]byte("ό\x00"), []byte("Ό\x00")}, {[]byte("π\x00"), []byte("Π\x00")}, {[]byte("ρ\x00"), []byte("Ρ\x00")}, {[]byte("ϲ\x00"), []byte("Ϲ\x00")}, {[]byte("σ\x00"), []byte("Σ\x00")}, {[]byte("ς\x00"), []byte("Σ\x00")}, {[]byte("ϭ\x00"), []byte("Ϭ\x00")}, {[]byte("τ\x00"), []byte("Τ\x00")}, {[]byte("ϯ\x00"), []byte("Ϯ\x00")}, {[]byte("υ\x00"), []byte("Υ\x00")}, {[]byte("ύ\x00"), []byte("Ύ\x00")}, {[]byte("ϋ\x00"), []byte("Ϋ\x00")}, {[]byte("φ\x00"), []byte("Φ\x00")}, {[]byte("ϥ\x00"), []byte("Ϥ\x00")}, {[]byte("χ\x00"), []byte("Χ\x00")}, {[]byte("ψ\x00"), []byte("Ψ\x00")}, {[]byte("ϣ\x00"), []byte("Ϣ\x00")}, {[]byte("ω\x00"), []byte("Ω\x00")}, {[]byte("ώ\x00"), []byte("Ώ\x00")}, {[]byte("а\x00"), []byte("А\x00")}, {[]byte("ӓ\x00"), []byte("Ӓ\x00")}, {[]byte("б\x00"), []byte("Б\x00")}, {[]byte("в\x00"), []byte("В\x00")}, {[]byte("г\x00"), []byte("Г\x00")}, {[]byte("ґ\x00"), []byte("Ґ\x00")}, {[]byte("д\x00"), []byte("Д\x00")}, {[]byte("ђ\x00"), []byte("Ђ\x00")}, {[]byte("е\x00"), []byte("Е\x00")}, {[]byte("ѐ\x00"), []byte("Ѐ\x00")}, {[]byte("є\x00"), []byte("Є\x00")}, {[]byte("ё\x00"), []byte("Ё\x00")}, {[]byte("ж\x00"), []byte("Ж\x00")}, {[]byte("з\x00"), []byte("З\x00")}, {[]byte("ѕ\x00"), []byte("Ѕ\x00")}, {[]byte("и\x00"), []byte("И\x00")}, {[]byte("ѝ\x00"), []byte("Ѝ\x00")}, {[]byte("ӥ\x00"), []byte("Ӥ\x00")}, {[]byte("і\x00"), []byte("І\x00")}, {[]byte("ї\x00"), []byte("Ї\x00")}, {[]byte("й\x00"), []byte("Й\x00")}, {[]byte("ј\x00"), []byte("Ј\x00")}, {[]byte("к\x00"), []byte("К\x00")}, {[]byte("л\x00"), []byte("Л\x00")}, {[]byte("љ\x00"), []byte("Љ\x00")}, {[]byte("м\x00"), []byte("М\x00")}, {[]byte("н\x00"), []byte("Н\x00")}, {[]byte("њ\x00"), []byte("Њ\x00")}, {[]byte("ᲂ\x00"), []byte("О\x00")}, {[]byte("о\x00"), []byte("О\x00")}, {[]byte("ӧ\x00"), []byte("Ӧ\x00")}, {[]byte("ө\x00"), []byte("Ө\x00")}, {[]byte("п\x00"), []byte("П\x00")}, {[]byte("р\x00"), []byte("Р\x00")}, {[]byte("с\x00"), []byte("С\x00")}, {[]byte("т\x00"), []byte("Т\x00")}, {[]byte("ћ\x00"), []byte("Ћ\x00")}, {[]byte("у\x00"), []byte("У\x00")}, {[]byte("ӱ\x00"), []byte("Ӱ\x00")}, {[]byte("ү\x00"), []byte("Ү\x00")}, {[]byte("ў\x00"), []byte("Ў\x00")}, {[]byte("ф\x00"), []byte("Ф\x00")}, {[]byte("х\x00"), []byte("Х\x00")}, {[]byte("ц\x00"), []byte("Ц\x00")}, {[]byte("ч\x00"), []byte("Ч\x00")}, {[]byte("џ\x00"), []byte("Џ\x00")}, {[]byte("ш\x00"), []byte("Ш\x00")}, {[]byte("щ\x00"), []byte("Щ\x00")}, {[]byte("ᲆ\x00"), []byte("Ъ\x00")}, {[]byte("ъ\x00"), []byte("Ъ\x00")}, {[]byte("ы\x00"), []byte("Ы\x00")}, {[]byte("ӹ\x00"), []byte("Ӹ\x00")}, {[]byte("ь\x00"), []byte("Ь\x00")}, {[]byte("э\x00"), []byte("Э\x00")}, {[]byte("ӭ\x00"), []byte("Ӭ\x00")}, {[]byte("ю\x00"), []byte("Ю\x00")}, {[]byte("я\x00"), []byte("Я\x00")}, {[]byte("ա\x00"), []byte("Ա\x00")}, {[]byte("բ\x00"), []byte("Բ\x00")}, {[]byte("գ\x00"), []byte("Գ\x00")}, {[]byte("դ\x00"), []byte("Դ\x00")}, {[]byte("ե\x00"), []byte("Ե\x00")}, {[]byte("զ\x00"), []byte("Զ\x00")}, {[]byte("է\x00"), []byte("Է\x00")}, {[]byte("ը\x00"), []byte("Ը\x00")}, {[]byte("թ\x00"), []byte("Թ\x00")}, {[]byte("ժ\x00"), []byte("Ժ\x00")}, {[]byte("ի\x00"), []byte("Ի\x00")}, {[]byte("լ\x00"), []byte("Լ\x00")}, {[]byte("խ\x00"), []byte("Խ\x00")}, {[]byte("ծ\x00"), []byte("Ծ\x00")}, {[]byte("կ\x00"), []byte("Կ\x00")}, {[]byte("հ\x00"), []byte("Հ\x00")}, {[]byte("ձ\x00"), []byte("Ձ\x00")}, {[]byte("ղ\x00"), []byte("Ղ\x00")}, {[]byte("ճ\x00"), []byte("Ճ\x00")}, {[]byte("մ\x00"), []byte("Մ\x00")}, {[]byte("յ\x00"), []byte("Յ\x00")}, {[]byte("ն\x00"), []byte("Ն\x00")}, {[]byte("շ\x00"), []byte("Շ\x00")}, {[]byte("ո\x00"), []byte("Ո\x00")}, {[]byte("չ\x00"), []byte("Չ\x00")}, {[]byte("պ\x00"), []byte("Պ\x00")}, {[]byte("ջ\x00"), []byte("Ջ\x00")}, {[]byte("ռ\x00"), []byte("Ռ\x00")}, {[]byte("ս\x00"), []byte("Ս\x00")}, {[]byte("վ\x00"), []byte("Վ\x00")}, {[]byte("տ\x00"), []byte("Տ\x00")}, {[]byte("ր\x00"), []byte("Ր\x00")}, {[]byte("ց\x00"), []byte("Ց\x00")}, {[]byte("փ\x00"), []byte("Փ\x00")}, {[]byte("ք\x00"), []byte("Ք\x00")}, {[]byte("օ\x00"), []byte("Օ\x00")}}

// tr_hpfa - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:353
func tr_hpfa(args [][]byte) {
	var tok []byte = make([]byte, 128)
	var c1 []byte = make([]byte, 32)
	var c2 []byte = make([]byte, 32)
	var filp *noarch.File
	hyinit = 1
	if args[1] == nil {
		// load english hyphenation patterns with no arguments
		hyph_readpatterns(en_patterns)
		hyph_readexceptions(en_exceptions)
	}
	if len(args[1]) == 0 && (func() *noarch.File {
		filp = noarch.Fopen(args[1], []byte("r\x00"))
		return filp
	}()) == nil {
		for noarch.Fscanf(filp, []byte("%128s\x00"), tok) == 1 {
			if noarch.Strlen(tok) < int32(256) {
				// reading patterns
				hy_add(tok)
			}
		}
		noarch.Fclose(filp)
	}
	if len(args[2]) == 0 && (func() *noarch.File {
		filp = noarch.Fopen(args[2], []byte("r\x00"))
		return filp
	}()) == nil {
		for noarch.Fscanf(filp, []byte("%128s\x00"), tok) == 1 {
			if noarch.Strlen(tok) < int32(256) {
				// reading exceptions
				hw_add(tok)
			}
		}
		noarch.Fclose(filp)
	}
	if len(args[3]) == 0 && (func() *noarch.File {
		filp = noarch.Fopen(args[3], []byte("r\x00"))
		return filp
	}()) == nil {
		for noarch.Fscanf(filp, []byte("%128s\x00"), tok) == 1 {
			// reading hcode mappings
			var s []byte = tok
			if utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], c1) != 0 && utf8read((*[1000000][]byte)(unsafe.Pointer(&s))[:], c2) != 0 && noarch.Not(s[0]) {
				// inverting
				hcode_add(c2, c1)
			}
		}
		noarch.Fclose(filp)
	}
	if args[3] != nil && noarch.Not(noarch.Strcmp([]byte("-\x00"), args[3])) {
		// lowercase-uppercase character hcode mappings
		var i int32
		for i = 0; uint32(i) < 4112/16; i++ {
			hcode_add(hycase[i][1], hycase[i][0])
		}
	}
}

// hyph_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:394
func hyph_init() {
	hwdict = dict_make(-1, 0, 2)
	hydict = dict_make(-1, 0, 2)
	hcodedict = dict_make(-1, 0, 1)
}

// hyph_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:401
func hyph_done() {
	if hwdict != nil {
		dict_free(hwdict)
	}
	if hydict != nil {
		dict_free(hydict)
	}
	if hcodedict != nil {
		dict_free(hcodedict)
	}
}

// tr_hpf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/hyph.c:411
func tr_hpf(args [][]byte) {
	// reseting the patterns
	hypats_len = 0
	hy_n = 0
	dict_free(hydict)
	// reseting the dictionary
	hwword_len = 0
	hw_n = 0
	dict_free(hwdict)
	// reseting hcode mappings
	hcode_n = 0
	dict_free(hcodedict)
	// reading
	hyph_init()
	tr_hpfa(args)
}

// inbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:7
// input stream management
type inbuf struct {
	path  [1024]byte
	fin   *noarch.File
	buf   []byte
	args  [][]byte
	unbuf [32]int32
	un    int32
	pos   int32
	len_  int32
	lnum  int32
	prev  []inbuf
}

// buf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:20
// for file buffers
// for string buffers
// unread characters
// number of unread characters
// file line number
var buf []inbuf

// files - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:21
var files [][]byte = make([][]byte, 16)

// nfiles - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:22
var nfiles int32

// cfile - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:23
var cfile int32

// in_new - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:28
func in_new() {
	var next []inbuf = xmalloc(int32(1248)).([]inbuf)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&next[0]))) / int64(1))))[:], byte(0), 1248)
	next[0].prev = buf
	buf = next
}

// in_push - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:36
func in_push(s []byte, args [][]byte) {
	var len_ int32 = noarch.Strlen(s)
	in_new()
	buf[0].buf = xmalloc(len_ + 1).([]byte)
	buf[0].len_ = len_
	noarch.Strcpy(buf[0].buf, s)
	if args != nil {
		buf[0].args = args_init(args)
	} else {
		buf[0].args = nil
	}
}

// in_so - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:46
func in_so(path []byte) {
	var fin *noarch.File = func() *noarch.File {
		if path != nil && int32(path[0]) != 0 {
			return noarch.Fopen(path, []byte("r\x00"))
		}
		return noarch.Stdin
	}()
	if fin == nil {
		errmsg([]byte("neatroff: failed to open <%s>\n\x00"), path)
		return
	}
	in_new()
	buf[0].fin = fin
	buf[0].lnum = 1
	if path != nil {
		noarch.Snprintf(buf[0].path[:], int32(1024), []byte("%s\x00"), path)
	}
}

// in_lf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:60
func in_lf(path []byte, lnum int32) {
	var cur []inbuf = buf
	for cur != nil && cur[0].fin == nil {
		cur = cur[0].prev
	}
	if path != nil {
		noarch.Snprintf(cur[0].path[:], int32(1024), []byte("%s\x00"), path)
	}
	cur[0].lnum = lnum
}

// in_queue - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:70
func in_queue(path []byte) {
	if nfiles < 16 {
		noarch.Snprintf(files[func() int32 {
			defer func() {
				nfiles++
			}()
			return nfiles
		}()], 1024, []byte("%s\x00"), func() []byte {
			if path != nil {
				return path
			}
			return []byte("\x00")
		}())
	}
}

// in_pop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:76
func in_pop() {
	var old []inbuf = buf
	buf = buf[0].prev
	if old[0].args != nil {
		args_free(old[0].args)
	}
	if old[0].fin != nil && (int64(uintptr(unsafe.Pointer(old[0].fin[0])))/int64(8)-int64(uintptr(unsafe.Pointer(noarch.Stdin)))/int64(8)) != 0 {
		noarch.Fclose(old[0].fin)
	}
	_ = old[0].buf
	_ = old
}

// in_nx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:88
func in_nx(path []byte) {
	for buf != nil {
		in_pop()
	}
	if path != nil {
		in_so(path)
	}
}

// in_ex - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:96
func in_ex() {
	for buf != nil {
		in_pop()
	}
	cfile = nfiles
}

// in_nextfile - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:103
func in_nextfile() int32 {
	for buf == nil && cfile < nfiles {
		in_so(files[func() int32 {
			defer func() {
				cfile++
			}()
			return cfile
		}()])
	}
	return noarch.BoolToInt(buf == nil)
}

// in_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:110
func in_next() int32 {
	var c int32
	for buf != nil || noarch.Not(in_nextfile()) {
		if buf[0].un != 0 {
			return buf[0].unbuf[:][func() int32 {
				tempVar1 := &buf[0].un
				*tempVar1--
				return *tempVar1
			}()]
		}
		if buf[0].buf != nil && buf[0].pos < buf[0].len_ {
			break
		}
		if buf[0].buf == nil && (func() int32 {
			c = noarch.Fgetc(buf[0].fin)
			return c
		}()) >= 0 {
			if c == int32('\n') {
				buf[0].lnum++
			}
			return c
		}
		in_pop()
	}
	if buf != nil {
		return int32(uint8(buf[0].buf[func() int32 {
			tempVar1 := &buf[0].pos
			defer func() {
				*tempVar1++
			}()
			return *tempVar1
		}()]))
	}
	return -1
}

// in_back - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:128
func in_back(c int32) {
	if c < 0 {
		return
	}
	if buf != nil && uint32(buf[0].un) < 128 {
		buf[0].unbuf[:][func() int32 {
			tempVar1 := &buf[0].un
			defer func() {
				*tempVar1++
			}()
			return *tempVar1
		}()] = c
	}
}

// in_top - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:136
func in_top() int32 {
	if buf != nil && buf[0].un != 0 {
		return buf[0].unbuf[:][buf[0].un-1]
	}
	return -1
}

// in_arg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:141
func in_arg(i int32) []byte {
	var cur []inbuf = buf
	for cur != nil && cur[0].args == nil {
		cur = cur[0].prev
	}
	if len(cur) == 0 && len(cur[0].args) == 0 && i < 32 && cur[0].args[i] != nil {
		return cur[0].args[i]
	}
	return []byte("\x00")
}

// in_nargs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:150
func in_nargs() int32 {
	var cur []inbuf = buf
	var n int32
	for cur != nil && cur[0].args == nil {
		cur = cur[0].prev
	}
	for len(cur) == 0 && len(cur[0].args) == 0 && cur[0].args[n] != nil {
		n++
	}
	return n
}

// in_shift - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:161
func in_shift() {
	var cur []inbuf = buf
	for cur != nil && cur[0].args == nil {
		cur = cur[0].prev
	}
	if len(cur) == 0 && len(cur[0].args) == 0 {
		_ = cur[0].args[1]
		noarch.Memmove((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&cur[0].args[0+1]))) / int64(1))))[:], (*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&cur[0].args[0+2]))) / int64(1))))[:], uint32(32-2)*8)
		cur[0].args[32-1] = nil
	}
}

// in_filename - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:174
func in_filename() []byte {
	var cur []inbuf = buf
	for cur != nil && cur[0].fin == nil {
		cur = cur[0].prev
	}
	if cur != nil && int32(cur[0].path[:][0]) != 0 {
		return cur[0].path[:]
	}
	return []byte("-\x00")
}

// in_lnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:182
func in_lnum() int32 {
	var cur []inbuf = buf
	for cur != nil && cur[0].fin == nil {
		cur = cur[0].prev
	}
	if cur != nil {
		return cur[0].lnum
	}
	return 0
}

// args_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:190
func args_init(args [][]byte) [][]byte {
	var out [][]byte = xmalloc(int32(32 * 8)).([][]byte)
	var i int32
	for i = 0; i < 32; i++ {
		out[i] = nil
		if args[i] != nil {
			var len_ int32 = noarch.Strlen(args[i]) + int32(1)
			out[i] = xmalloc(len_).([]byte)
			memcpy(out[i], args[i], uint32(len_))
		}
	}
	return out
}

// args_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/in.c:205
func args_free(args [][]byte) {
	var i int32
	for i = 0; i < 32; i++ {
		if args[i] != nil {
			_ = args[i]
		}
	}
	_ = args
}

// iset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:10
// iset structure to map integers to sets
type iset struct {
	set  [][]int32
	sz   []int32
	len_ []int32
	cnt  int32
}

// iset_extend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:17
func iset_extend(iset_c4go_postfix []iset, cnt int32) {
	iset_c4go_postfix[0].set = mextend(iset_c4go_postfix[0].set, iset_c4go_postfix[0].cnt, cnt, int32(8)).([][]int32)
	iset_c4go_postfix[0].sz = mextend(iset_c4go_postfix[0].sz, iset_c4go_postfix[0].cnt, cnt, int32(4)).([]int32)
	iset_c4go_postfix[0].len_ = mextend(iset_c4go_postfix[0].len_, iset_c4go_postfix[0].cnt, cnt, int32(4)).([]int32)
	iset_c4go_postfix[0].cnt = cnt
}

// iset_make - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:25
func iset_make() []iset {
	var iset_c4go_postfix []iset = xmalloc(int32(56)).([]iset)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&iset_c4go_postfix[0]))) / int64(1))))[:], byte(0), 56)
	iset_extend(iset_c4go_postfix, 1<<uint64(10))
	return iset_c4go_postfix
}

// iset_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:33
func iset_free(iset_c4go_postfix []iset) {
	var i int32
	for i = 0; i < iset_c4go_postfix[0].cnt; i++ {
		_ = iset_c4go_postfix[0].set[i]
	}
	_ = iset_c4go_postfix[0].set
	_ = iset_c4go_postfix[0].len_
	_ = iset_c4go_postfix[0].sz
	_ = iset_c4go_postfix
}

// iset_get - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:44
func iset_get(iset_c4go_postfix []iset, key int32) []int32 {
	if key >= 0 && key < iset_c4go_postfix[0].cnt {
		return iset_c4go_postfix[0].set[key]
	}
	return nil
}

// iset_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:49
func iset_len(iset_c4go_postfix []iset, key int32) int32 {
	if key >= 0 && key < iset_c4go_postfix[0].cnt {
		return iset_c4go_postfix[0].len_[key]
	}
	return 0
}

// iset_put - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/iset.c:54
func iset_put(iset_c4go_postfix []iset, key int32, ent int32) {
	if key < 0 || key >= 1<<uint64(20) {
		return
	}
	if key >= iset_c4go_postfix[0].cnt {
		iset_extend(iset_c4go_postfix, (key+1+1<<uint64(10)-1) & ^(1<<uint64(10)-1))
	}
	if key >= 0 && key < iset_c4go_postfix[0].cnt && iset_c4go_postfix[0].len_[key]+1 >= iset_c4go_postfix[0].sz[key] {
		var olen int32 = iset_c4go_postfix[0].sz[key]
		var nlen int32 = iset_c4go_postfix[0].sz[key]*2 + 8
		var nset interface{} = xmalloc(int32(uint32(nlen) * 4))
		if iset_c4go_postfix[0].set[key] != nil {
			memcpy(nset, iset_c4go_postfix[0].set[key], uint32(olen)*4)
			_ = iset_c4go_postfix[0].set[key]
		}
		iset_c4go_postfix[0].sz[key] = nlen
		iset_c4go_postfix[0].set[key] = nset.([]int32)
	}
	iset_c4go_postfix[0].set[key][func() int32 {
		tempVar1 := &iset_c4go_postfix[0].len_[key]
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}()] = ent
	iset_c4go_postfix[0].set[key][iset_c4go_postfix[0].len_[key]] = -1
}

// mapdict - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/map.c:9
// mapping register/macro names to indices
// register, macro, or environments names
var mapdict []dict

// map_ - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/map.c:12
func map_(s []byte) int32 {
	// map register names to [0..NREGS]
	var i int32
	if int32(s[0]) == int32('.') && int32(s[1]) != 0 && noarch.Not(s[2]) {
		// ".x" is mapped to 'x'
		return int32(uint8(s[1]))
	}
	if mapdict == nil {
		mapdict = dict_make(-1, 1, 2)
	}
	i = dict_idx(mapdict, s)
	if i < 0 {
		dict_put(mapdict, s, 0)
		i = dict_idx(mapdict, s)
		if 256+i >= 8192 {
			errdie([]byte("neatroff: increase NREGS\n\x00"))
		}
	}
	return 256 + i
}

// map_name - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/map.c:30
func map_name(id int32) []byte {
	// return the name mapped to id; returns a static buffer
	var map_buf []byte = make([]byte, 128)
	if id >= 256 {
		return dict_key(mapdict, id-256)
	}
	map_buf[0] = '.'
	map_buf[1] = byte(id)
	map_buf[2] = '\x00'
	return map_buf
}

// map_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/map.c:41
func map_done() {
	if mapdict != nil {
		dict_free(mapdict)
	}
}

// out_nl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:9
// generating troff output
var out_nl int32 = 1

// out_out - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:12
func out_out(s []byte, ap *va_list) {
	// output troff code; newlines may appear only at the end of s
	out_nl = noarch.BoolToInt(len(noarch.Strchr(s, int32('\n'))) != 0)
	noarch.Vfprintf(noarch.Stdout, s, ap)
}

// outnn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:19
func outnn(s []byte, c4goArgs ...interface{}) {
	// output troff code; no preceding newline is necessary
	var ap *va_list
	va_start(ap, s)
	out_out(s, ap)
	va_end(ap)
}

// out - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:28
func out(s []byte, c4goArgs ...interface{}) {
	// output troff cmd; should appear after a newline
	var ap *va_list
	if noarch.Not(out_nl) {
		outnn([]byte("\n\x00"))
	}
	va_start(ap, s)
	out_out(s, ap)
	va_end(ap)
}

// o_s - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:38
var o_s int32 = 10

// o_f - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:39
var o_f int32 = 1

// o_m - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:40
var o_m int32

// out_ps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:42
func out_ps(n int32) {
	if o_s != n {
		o_s = n
		out([]byte("s%d\n\x00"), o_s)
	}
}

// out_ft - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:50
func out_ft(n int32) {
	if n >= 0 && o_f != n {
		o_f = n
		out([]byte("f%d\n\x00"), o_f)
	}
}

// out_clr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:58
func out_clr(n int32) {
	if n >= 0 && o_m != n {
		o_m = n
		out([]byte("m%s\n\x00"), clr_str(o_m))
	}
}

// out_draw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:83
func out_draw(s []byte) {
	var c int32 = int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0])
	out([]byte("D%c\x00"), c)
	switch tolower(c) {
	case 'l':
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m')))
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v')))
		// dpost requires this
		outnn([]byte(" .\x00"))
	case 'c':
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m')))
	case 'e':
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m')))
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v')))
	case 'a':
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m')))
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v')))
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m')))
		outnn([]byte(" %d\x00"), tok_num((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v')))
	case '~':
		fallthrough
	case 'p':
		for s[0] != 0 {
			var h int32
			var v int32
			if tok_numpt((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('m'), c4goUnsafeConvert_int32(&h)) != 0 || tok_numpt((*[1000000][]byte)(unsafe.Pointer(&s))[:], int32('v'), c4goUnsafeConvert_int32(&v)) != 0 {
				outnn([]byte(" \x00"))
				for int32(s[0]) != 0 && int32(s[0]) != int32(' ') {
					outnn([]byte("%c\x00"), int32((func() []byte {
						defer func() {
							s = s[0+1:]
						}()
						return s
					}())[0]))
				}
			} else {
				outnn([]byte(" %d\x00"), h)
				outnn([]byte(" %d\x00"), v)
			}
		}
		break
	}
	outnn([]byte("\n\x00"))
}

// outg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:124
func outg(c []byte, fn int32, sz int32) {
	var ofn int32 = o_f
	var osz int32 = o_s
	out_ft(fn)
	out_ps(sz)
	if utf8one(c) != 0 {
		outnn([]byte("c%s%s\x00"), c, func() []byte {
			if int32(c[1]) != 0 {
				return []byte("\n\x00")
			}
			return []byte("\x00")
		}())
	} else {
		out([]byte("C%s\n\x00"), func() []byte {
			if int32(c[0]) == c_ec && int32(c[1]) == int32('(') {
				return c[0+2:]
			}
			return c
		}())
	}
	out_ft(ofn)
	out_ps(osz)
}

// outc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:138
func outc(c []byte) {
	var g []glyph = dev_glyph(c, o_f)
	var fn []font = dev_font(o_f)
	var cwid int32
	var bwid int32
	if g == nil {
		return
	}
	cwid = font_gwid(g[0].font, dev_font(o_f), o_s, int32(g[0].wid))
	bwid = font_wid(g[0].font, o_s, int32(g[0].wid))
	if font_mapped(g[0].font, c) != 0 {
		c = g[0].name[:]
	}
	if font_getcs(fn) != 0 {
		outnn([]byte("h%d\x00"), (cwid-bwid)/2)
	}
	outg(c, dev_fontpos(g[0].font), font_zoom(g[0].font, o_s))
	if font_getbd(fn) != 0 {
		outnn([]byte("h%d\x00"), font_getbd(fn)-1)
		outg(c, dev_fontpos(g[0].font), font_zoom(g[0].font, o_s))
		outnn([]byte("h%d\x00"), -font_getbd(fn)+1)
	}
	if font_getcs(fn) != 0 {
		outnn([]byte("h%d\x00"), -(cwid-bwid)/2)
	}
	outnn([]byte("h%d\x00"), cwid)
}

// out_x - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:162
func out_x(s []byte) {
	out([]byte("x X %s\n\x00"), s)
}

// out_line - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/out.c:167
func out_line(s []byte) {
	var c []byte
	var t int32
	for (func() int32 {
		t = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&c))[:])
		return t
	}()) >= 0 {
		if noarch.Not(t) {
			if int32(c[0]) == 4 || int32(c[0]) == int32('\\') && int32(c[1]) == int32('\\') {
				c[0] = c[1]
				c[1] = '\x00'
			}
			if int32(c[0]) == int32('\t') || int32(c[0]) == int32('\x01') || c_hymark(c) != 0 {
				continue
			}
			outc(cmap_map(c))
			continue
		}
		switch t {
		case 'D':
			out_draw(c)
		case 'f':
			out_ft(dev_pos(c))
		case 'h':
			outnn([]byte("h%d\x00"), eval(c, int32('m')))
		case 'm':
			if noarch.Not((nreg(int32('C')))[0]) {
				out_clr(clr_get(c))
			}
		case 's':
			out_ps(eval(c, 0))
		case 'v':
			outnn([]byte("v%d\x00"), eval(c, int32('v')))
		case 'X':
			out_x(c)
			break
		}
	}
}

// env - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:13
// registers and environments
type env struct {
	eregs     [64]int32
	tabs      [32]int32
	tabs_type [32]byte
	fmt_      []fmt_
	wb        wb
	tc        [32]byte
	lc        [32]byte
	hc        [32]byte
	mc        [32]byte
}

// nregs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:25
// environment-specific number registers
// tab stops
// type of tabs: L, C, R
// per environment line formatting buffer
// per environment partial word
// tab character (.tc)
// leader character (.lc)
// hyphenation character (.hc)
// margin character (.mc)
// global number registers
var nregs []int32 = make([]int32, 8192)

// nregs_inc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:26
// number register auto-increment size
var nregs_inc []int32 = make([]int32, 8192)

// nregs_fmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:27
// number register format
var nregs_fmt []int32 = make([]int32, 8192)

// sregs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:28
// global string registers
var sregs [][]byte = make([][]byte, 8192)

// sregs_dat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:29
// builtin function data
var sregs_dat []interface{} = make([]interface{}, 8192)

// envs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:30
// environments
var envs [][]env = make([][]env, 8192)

// env_c4go_postfix - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:31
// current enviroment
var env_c4go_postfix []env

// env_id - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:32
// current environment id
var env_id int32

// eregs_idx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:33
// register environment index in eregs[]
var eregs_idx []int32 = make([]int32, 8192)

// eregs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:35
// environment-specific number registers
var eregs [][]byte = [][]byte{[]byte("ln\x00"), []byte(".f\x00"), []byte(".i\x00"), []byte(".j\x00"), []byte(".l\x00"), []byte(".L\x00"), []byte(".nI\x00"), []byte(".nm\x00"), []byte(".nM\x00"), []byte(".nn\x00"), []byte(".nS\x00"), []byte(".m\x00"), []byte(".s\x00"), []byte(".u\x00"), []byte(".v\x00"), []byte(".it\x00"), []byte(".itn\x00"), []byte(".mc\x00"), []byte(".mcn\x00"), []byte(".ce\x00"), []byte(".f0\x00"), []byte(".i0\x00"), []byte(".l0\x00"), []byte(".hy\x00"), []byte(".hycost\x00"), []byte(".hycost2\x00"), []byte(".hycost3\x00"), []byte(".hlm\x00"), []byte(".L0\x00"), []byte(".m0\x00"), []byte(".n0\x00"), []byte(".s0\x00"), []byte(".ss\x00"), []byte(".ssh\x00"), []byte(".sss\x00"), []byte(".pmll\x00"), []byte(".pmllcost\x00"), []byte(".ti\x00"), []byte(".lt\x00"), []byte(".lt0\x00"), []byte(".v0\x00"), []byte(".I\x00"), []byte(".I0\x00"), []byte(".tI\x00"), []byte(".td\x00"), []byte(".cd\x00")}

// nreg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:48
func nreg(id int32) []int32 {
	if eregs_idx[id] != 0 {
		// return the address of a number register
		return env_c4go_postfix[0].eregs[:][eregs_idx[id]:]
	}
	return nregs[id:]
}

// directory - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:55
func directory(path []byte) []byte {
	var dst []byte = make([]byte, 1024)
	var s []byte = noarch.Strrchr(path, int32('/'))
	if s == nil {
		return []byte(".\x00")
	}
	if (int64(uintptr(unsafe.Pointer(&path[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)) == 0 {
		return []byte("/\x00")
	}
	memcpy(dst, path, uint32(int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&path[0])))/int64(1)))))
	dst[int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&path[0])))/int64(1)))] = '\x00'
	return dst
}

// num_tabs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:68
func num_tabs() []byte {
	var tabs []byte = make([]byte, 512)
	var i int32
	var s []byte = tabs
	for i = 0; i < 32 && int32(env_c4go_postfix[0].tabs_type[:][i]) != 0; i++ {
		s = s[0+noarch.Sprintf(s, []byte("%du%c \x00"), env_c4go_postfix[0].tabs[:][i], int32(env_c4go_postfix[0].tabs_type[:][i])):]
	}
	return tabs
}

// num_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:81
func num_str(id int32) []byte {
	// the contents of a number register (returns a static buffer)
	var numbuf []byte = make([]byte, 128)
	var s []byte = map_name(id)
	if noarch.Not(nregs_fmt[id]) {
		nregs_fmt[id] = int32('0')
	}
	numbuf[0] = '\x00'
	if int32(s[0]) == int32('.') && noarch.Not(s[2]) {
		switch int32(s[1]) {
		case 'b':
			noarch.Sprintf(numbuf, []byte("%d\x00"), font_getbd(dev_font((nreg(int32('f')))[0])))
			return numbuf
		case 'c':
			noarch.Sprintf(numbuf, []byte("%d\x00"), in_lnum())
			return numbuf
		case 'k':
			noarch.Sprintf(numbuf, []byte("%d\x00"), f_hpos())
			return numbuf
		case 'm':
			noarch.Sprintf(numbuf, []byte("#%02x%02x%02x\x00"), (nreg(int32('m')))[0]>>uint64(16)&255, (nreg(int32('m')))[0]>>uint64(8)&255, (nreg(int32('m')))[0]&255)
			return numbuf
		case 't':
			noarch.Sprintf(numbuf, []byte("%d\x00"), f_nexttrap())
			return numbuf
		case 'z':
			if f_divreg() >= 0 {
				noarch.Sprintf(numbuf, []byte("%s\x00"), map_name(f_divreg()))
			}
			return numbuf
		case 'F':
			noarch.Sprintf(numbuf, []byte("%s\x00"), in_filename())
			return numbuf
		case 'D':
			noarch.Sprintf(numbuf, []byte("%s\x00"), directory(in_filename()))
			return numbuf
		case '$':
			noarch.Sprintf(numbuf, []byte("%d\x00"), in_nargs()-1)
			return numbuf
		}
	}
	if int32(s[0]) == int32('.') && noarch.Not(noarch.Strcmp([]byte(".neat\x00"), s)) {
		return []byte("1\x00")
	}
	if int32(s[0]) == int32('.') && int32(s[1]) == int32('e') && int32(s[2]) == int32('v') && noarch.Not(s[3]) {
		return map_name(env_id)
	}
	if int32(s[0]) == int32('$') && int32(s[1]) == int32('$') && noarch.Not(s[2]) {
		noarch.Sprintf(numbuf, []byte("%d\x00"), getpid())
		return numbuf
	}
	if int32(s[0]) == int32('y') && int32(s[1]) == int32('r') && noarch.Not(s[2]) {
		noarch.Sprintf(numbuf, []byte("%02d\x00"), (nreg(id))[0])
		return numbuf
	}
	if int32(s[0]) == int32('.') && noarch.Not(noarch.Strcmp([]byte(".tabs\x00"), s)) {
		return num_tabs()
	}
	if noarch.Not(nregs_fmt[id]) || num_fmt(numbuf, (nreg(id))[0], nregs_fmt[id]) != 0 {
		noarch.Sprintf(numbuf, []byte("%d\x00"), (nreg(id))[0])
	}
	return numbuf
}

// num_set - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:140
func num_set(id int32, val int32) {
	if noarch.Not(nregs_fmt[id]) {
		nregs_fmt[id] = int32('0')
	}
	(nreg(id))[0] = val
}

// num_setinc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:147
func num_setinc(id int32, val int32) {
	nregs_inc[id] = val
}

// num_inc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:152
func num_inc(id int32, pos int32) {
	(nreg(id))[0] += func() int32 {
		if pos > 0 {
			return nregs_inc[id]
		}
		return -nregs_inc[id]
	}()
}

// num_del - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:157
func num_del(id int32) {
	(nreg(id))[0] = 0
	nregs_inc[id] = 0
	nregs_fmt[id] = 0
}

// str_set - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:164
func str_set(id int32, s []byte) {
	var len_ int32 = noarch.Strlen(s) + int32(1)
	if sregs[id] != nil {
		_ = sregs[id]
	}
	sregs[id] = xmalloc(len_).([]byte)
	memcpy(sregs[id], s, uint32(len_))
	sregs_dat[id] = nil
}

// str_get - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:174
func str_get(id int32) []byte {
	return sregs[id]
}

// str_dget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:179
func str_dget(id int32) interface{} {
	return sregs_dat[id]
}

// str_dset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:184
func str_dset(id int32, d interface{}) {
	sregs_dat[id] = d
}

// str_rm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:189
func str_rm(id int32) {
	if sregs[id] != nil {
		_ = sregs[id]
	}
	sregs[id] = nil
	sregs_dat[id] = nil
}

// str_rn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:197
func str_rn(src int32, dst int32) {
	if sregs[src] == nil && sregs_dat[src] == nil {
		return
	}
	str_rm(dst)
	sregs[dst] = sregs[src]
	sregs_dat[dst] = sregs_dat[src]
	sregs[src] = nil
	sregs_dat[src] = nil
}

// env_alloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:208
func env_alloc() []env {
	var env_c4go_postfix []env = xmalloc(int32(8936)).([]env)
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&env_c4go_postfix[0]))) / int64(1))))[:], byte(0), 8936)
	wb_init((*[1000000]wb)(unsafe.Pointer(&env_c4go_postfix[0].wb))[:])
	env_c4go_postfix[0].fmt_ = fmt_alloc()
	return env_c4go_postfix
}

// env_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:217
func env_free(env_c4go_postfix []env) {
	fmt_free(env_c4go_postfix[0].fmt_)
	wb_done((*[1000000]wb)(unsafe.Pointer(&env_c4go_postfix[0].wb))[:])
	_ = env_c4go_postfix
}

// env_set - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:224
func env_set(id int32) {
	var i int32
	env_c4go_postfix = envs[id]
	env_id = id
	if env_c4go_postfix == nil {
		envs[id] = env_alloc()
		env_c4go_postfix = envs[id]
		(nreg(int32('f')))[0] = 1
		(nreg(int32('i')))[0] = 0
		(nreg(int32('I')))[0] = 0
		(nreg(int32('j')))[0] = 3
		(nreg(int32('l')))[0] = dev_res * 65 / 10
		(nreg(int32('L')))[0] = 1
		(nreg(int32('s')))[0] = 10
		(nreg(int32('u')))[0] = 1
		(nreg(int32('v')))[0] = 12 * (dev_res / 72)
		(nreg(map_([]byte(".s0\x00"))))[0] = (nreg(int32('s')))[0]
		(nreg(map_([]byte(".f0\x00"))))[0] = (nreg(int32('f')))[0]
		(nreg(map_([]byte(".na\x00"))))[0] = 0
		(nreg(map_([]byte(".lt\x00"))))[0] = dev_res * 65 / 10
		(nreg(map_([]byte(".hy\x00"))))[0] = 1
		(nreg(map_([]byte(".ss\x00"))))[0] = 12
		(nreg(map_([]byte(".sss\x00"))))[0] = 12
		(nreg(map_([]byte(".nM\x00"))))[0] = 1
		(nreg(map_([]byte(".nS\x00"))))[0] = 1
		noarch.Strcpy(env_c4go_postfix[0].hc[:], []byte("\\%\x00"))
		noarch.Strcpy(env_c4go_postfix[0].lc[:], []byte(".\x00"))
		for i = 0; i < 32; i++ {
			env_c4go_postfix[0].tabs[:][i] = i * dev_res / 2
			env_c4go_postfix[0].tabs_type[:][i] = 'L'
		}
	}
}

// init_time - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:259
func init_time() {
	var t noarch.TimeT = noarch.Time(nil)
	var tm_c4go_postfix []noarch.Tm = noarch.LocalTime((*[1000000]noarch.TimeT)(unsafe.Pointer(&t))[:])
	num_set(map_([]byte("dw\x00")), tm_c4go_postfix[0].TmWday+1)
	num_set(map_([]byte("dy\x00")), tm_c4go_postfix[0].TmMday)
	num_set(map_([]byte("mo\x00")), tm_c4go_postfix[0].TmMon+1)
	num_set(map_([]byte("yr\x00")), tm_c4go_postfix[0].TmYear%100)
	num_set(map_([]byte(".yr\x00")), 1900+tm_c4go_postfix[0].TmYear)
}

// init_globals - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:270
func init_globals() {
	(nreg(int32('o')))[0] = dev_res
	(nreg(int32('p')))[0] = dev_res * 11
	(nreg(map_([]byte(".lg\x00"))))[0] = 1
	(nreg(map_([]byte(".kn\x00"))))[0] = 1
	num_set(map_([]byte(".H\x00")), 1)
	num_set(map_([]byte(".V\x00")), 1)
}

// env_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:280
func env_init() {
	var i int32
	init_time()
	init_globals()
	for i = 0; uint32(i) < 368/8; i++ {
		eregs_idx[map_(eregs[i])] = i + 1
	}
	env_set(map_([]byte("0\x00")))
}

// env_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:290
func env_done() {
	var i int32
	for i = 0; uint32(i) < 65536/8; i++ {
		if envs[i] != nil {
			env_free(envs[i])
		}
	}
	for i = 0; uint32(i) < 65536/8; i++ {
		_ = sregs[i]
	}
}

// oenv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:300
// environment stack
var oenv []int32 = make([]int32, 16)

// nenv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:301
var nenv int32

// tr_ev - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:303
func tr_ev(args [][]byte) {
	var id int32 = -1
	if args[1] != nil {
		id = map_(args[1])
	} else {
		if nenv != 0 {
			id = oenv[func() int32 {
				nenv--
				return nenv
			}()]
		} else {
			id = -1
		}
	}
	if id < 0 {
		return
	}
	if len(args[1]) == 0 && len(env_c4go_postfix) == 0 && nenv < 16 {
		oenv[func() int32 {
			defer func() {
				nenv++
			}()
			return nenv
		}()] = env_id
	}
	env_set(id)
}

// env_fmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:317
func env_fmt() []fmt_ {
	return env_c4go_postfix[0].fmt_
}

// env_wb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:322
func env_wb() []wb {
	return (*[1000000]wb)(unsafe.Pointer(&env_c4go_postfix[0].wb))[:]
}

// env_hc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:327
func env_hc() []byte {
	return env_c4go_postfix[0].hc[:]
}

// env_mc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:332
func env_mc() []byte {
	return env_c4go_postfix[0].mc[:]
}

// env_tc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:337
func env_tc() []byte {
	return env_c4go_postfix[0].tc[:]
}

// env_lc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:342
func env_lc() []byte {
	return env_c4go_postfix[0].lc[:]
}

// odiv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:348
// saving and restoring registers around diverted lines
type odiv struct {
	f  int32
	s  int32
	m  int32
	f0 int32
	s0 int32
	m0 int32
	cd int32
}

// odivs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:352
// state before diverted text
var odivs []odiv = make([]odiv, 16)

// nodivs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:353
var nodivs int32

// odiv_beg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:356
func odiv_beg() {
	// begin outputting diverted line
	var o []odiv = odivs[func() int32 {
		defer func() {
			nodivs++
		}()
		return nodivs
	}():]
	o[0].f = (nreg(int32('f')))[0]
	o[0].s = (nreg(int32('s')))[0]
	o[0].m = (nreg(int32('m')))[0]
	o[0].f0 = (nreg(map_([]byte(".f0\x00"))))[0]
	o[0].s0 = (nreg(map_([]byte(".s0\x00"))))[0]
	o[0].m0 = (nreg(map_([]byte(".m0\x00"))))[0]
	o[0].cd = (nreg(map_([]byte(".cd\x00"))))[0]
}

// odiv_end - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:369
func odiv_end() {
	// end outputting diverted line
	var o []odiv = odivs[func() int32 {
		nodivs--
		return nodivs
	}():]
	(nreg(int32('f')))[0] = o[0].f
	(nreg(int32('s')))[0] = o[0].s
	(nreg(int32('m')))[0] = o[0].m
	(nreg(map_([]byte(".f0\x00"))))[0] = o[0].f0
	(nreg(map_([]byte(".s0\x00"))))[0] = o[0].s0
	(nreg(map_([]byte(".m0\x00"))))[0] = o[0].m0
	(nreg(map_([]byte(".cd\x00"))))[0] = o[0].cd
}

// tr_ta - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:381
func tr_ta(args [][]byte) {
	var i int32
	var c int32
	for i = 0; i < 32; i++ {
		if i+1 < 32 && args[i+1] != nil {
			var a []byte = args[i+1]
			env_c4go_postfix[0].tabs[:][i] = eval_re(a, func() int32 {
				if i > 0 {
					return env_c4go_postfix[0].tabs[:][i-1]
				}
				return 0
			}(), int32('m'))
			if int32(a[0]) != 0 {
				c = int32(uint8(c4goPointerArithByteSlice(noarch.Strchr(a, int32('\x00')), int(-1))[0]))
			} else {
				c = 0
			}
			env_c4go_postfix[0].tabs_type[:][i] = byte(func() int32 {
				if noarch.Strchr([]byte("LRC\x00"), c) != nil {
					return c
				}
				return int32('L')
			}())
		} else {
			env_c4go_postfix[0].tabs[:][i] = 0
			env_c4go_postfix[0].tabs_type[:][i] = byte(0)
		}
	}
}

// tab_idx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:398
func tab_idx(pos int32) int32 {
	var i int32
	for i = 0; uint32(i) < 128/4; i++ {
		if env_c4go_postfix[0].tabs[:][i] > pos {
			return i
		}
	}
	return -1
}

// tab_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:407
func tab_next(pos int32) int32 {
	var i int32 = tab_idx(pos)
	if i >= 0 {
		return env_c4go_postfix[0].tabs[:][i]
	}
	return pos
}

// tab_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:413
func tab_type(pos int32) int32 {
	var i int32 = tab_idx(pos)
	if i >= 0 && int32(env_c4go_postfix[0].tabs_type[:][i]) != 0 {
		return int32(env_c4go_postfix[0].tabs_type[:][i])
	}
	return int32('L')
}

// num_getfmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:424
func num_getfmt(id int32) []byte {
	// number register format (.af)
	// the format of a number register (returns a static buffer)
	var fmtbuf []byte = make([]byte, 128)
	var s []byte = fmtbuf
	var fmt_ int32 = nregs_fmt[id] & 255
	var i int32
	if fmt_ == int32('0') || fmt_ == int32('x') || fmt_ == int32('X') {
		i = nregs_fmt[id] >> uint64(8)
		for func() int32 {
			defer func() {
				i--
			}()
			return i
		}() > 1 {
			(func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0] = '0'
		}
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(fmt_)
	} else if nregs_fmt[id] != 0 {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(fmt_)
	}
	s[0] = '\x00'
	return fmtbuf
}

// num_setfmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:442
func num_setfmt(id int32, s []byte) {
	var i int32
	if noarch.Strchr([]byte("iIaA\x00"), int32(s[0])) != nil {
		nregs_fmt[id] = int32(s[0])
	} else {
		for int32(((__ctype_b_loc())[0])[int32(uint8(s[i]))])&int32(uint16(noarch.ISdigit)) != 0 {
			i++
		}
		if int32(s[i]) == int32('x') || int32(s[i]) == int32('X') {
			nregs_fmt[id] = int32(s[i]) | (i+1)<<uint64(8)
		} else {
			nregs_fmt[id] = int32('0' | i<<uint64(8))
		}
	}
}

// nf_reverse - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:457
func nf_reverse(s []byte) {
	var r []byte = make([]byte, 128)
	var i int32
	var l int32
	noarch.Strcpy(r, s)
	l = noarch.Strlen(r)
	for i = 0; i < l; i++ {
		s[i] = r[l-i-1]
	}
}

// nf_roman - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:467
func nf_roman(s []byte, n int32, I []byte, V []byte) {
	var i int32
	if noarch.Not(n) {
		return
	}
	if n%5 == 4 {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(func() int32 {
			if n%10 == 9 {
				return int32(I[1])
			}
			return int32(V[0])
		}())
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = I[0]
	} else {
		for i = 0; i < n%5; i++ {
			(func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0] = I[0]
		}
		if n%10 >= 5 {
			(func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0] = V[0]
		}
	}
	s[0] = '\x00'
	nf_roman(s, n/10, I[0+1:], V[0+1:])
}

// nf_alpha - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:485
func nf_alpha(s []byte, n int32, a int32) {
	for n != 0 {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(a + (n-1)%26)
		n /= 26
	}
	s[0] = '\x00'
}

// num_fmt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/reg.c:495
func num_fmt(s []byte, n int32, fmt_ int32) int32 {
	// returns nonzero on failure
	var type_ int32 = fmt_ & 255
	if n < 0 {
		n = -n
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = '-'
	}
	if (type_ == int32('i') || type_ == int32('I')) && n > 0 && n < 40000 {
		if type_ == int32('i') {
			nf_roman(s, n, []byte("ixcmz\x00"), []byte("vldw\x00"))
		} else {
			nf_roman(s, n, []byte("IXCMZ\x00"), []byte("VLDW\x00"))
		}
		nf_reverse(s)
		return 0
	}
	if (type_ == int32('a') || type_ == int32('A')) && n > 0 {
		nf_alpha(s, n, type_)
		nf_reverse(s)
		return 0
	}
	if type_ == int32('0') || type_ == int32('x') || type_ == int32('X') {
		var pat []byte = make([]byte, 16)
		noarch.Sprintf(pat, []byte("%%0%d%c\x00"), fmt_>>uint64(8), func() int32 {
			if type_ == int32('0') {
				return int32('d')
			}
			return type_
		}())
		noarch.Sprintf(s, pat, n)
		return 0
	}
	return 1
}

// div - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:13
// rendering lines and managing traps
// diversions
type div struct {
	sbuf    sbuf
	reg     int32
	tpos    int32
	treg    int32
	dl      int32
	prev_d  int32
	prev_h  int32
	prev_mk int32
	prev_ns int32
}

// divs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:24
// diversion output
// diversion register
// diversion trap position
// diversion trap register
// diversion width
// previous \n(.d value
// previous \n(.h value
// previous .mk internal register
// previous .ns value
// diversion stack
var divs []div = make([]div, 16)

// cdiv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:25
// current diversion
var cdiv []div

// ren_div - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:26
// rendering a diversion
var ren_div int32

// ren_divvs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:27
// the amount of .v in diversions
var ren_divvs int32

// trap_em - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:28
// end macro
var trap_em int32 = -1

// ren_nl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:30
// just after a newline
var ren_nl int32

// ren_partial - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:31
// reading an input line in render_rec()
var ren_partial int32

// ren_unbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:32
// ren_back() buffer
var ren_unbuf []int32 = make([]int32, 8)

// ren_un - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:33
var ren_un int32

// ren_aborted - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:34
// .ab executed
var ren_aborted int32

// bp_first - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:36
// prior to the first page
var bp_first int32 = 1

// bp_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:37
// next page number
var bp_next int32 = 1073741824

// bp_ejected - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:38
// current ejected page
var bp_ejected int32

// bp_final - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:39
// 1: executing em, 2: the final page, 3: the 2nd final page
var bp_final int32

// ren_level - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:40
// the depth of render_rec() calls
var ren_level int32

// c_fa - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:42
// field delimiter
var c_fa []byte = make([]byte, 32)

// c_fb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:43
// field padding
var c_fb []byte = make([]byte, 32)

// ren_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:45
func ren_next() int32 {
	if ren_un > 0 {
		return ren_unbuf[func() int32 {
			ren_un--
			return ren_un
		}()]
	}
	return tr_next()
}

// ren_back - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:50
func ren_back(c int32) {
	ren_unbuf[func() int32 {
		defer func() {
			ren_un++
		}()
		return ren_un
	}()] = c
}

// tr_di - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:55
func tr_di(args [][]byte) {
	if args[1] != nil {
		if cdiv != nil {
			cdiv = cdiv[0+1:]
		} else {
			cdiv = divs
		}
		noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&cdiv[0]))) / int64(1))))[:], byte(0), 56)
		sbuf_init((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:])
		cdiv[0].reg = map_(args[1])
		cdiv[0].treg = -1
		if int32(args[0][2]) == int32('a') && str_get(cdiv[0].reg) != nil {
			// .da
			sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], str_get(cdiv[0].reg))
		}
		sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("%c%s\n\x00"), c_cc, []byte("\a<\x00"))
		cdiv[0].prev_d = (nreg(int32('d')))[0]
		cdiv[0].prev_h = (nreg(int32('h')))[0]
		cdiv[0].prev_mk = (nreg(map_([]byte(".mk\x00"))))[0]
		cdiv[0].prev_ns = (nreg(map_([]byte(".ns\x00"))))[0]
		(nreg(int32('d')))[0] = 0
		(nreg(int32('h')))[0] = 0
		(nreg(map_([]byte(".mk\x00"))))[0] = 0
		(nreg(map_([]byte(".ns\x00"))))[0] = 0
	} else if cdiv != nil {
		sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("%c%s\n\x00"), c_cc, []byte("\a>\x00"))
		str_set(cdiv[0].reg, sbuf_buf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:]))
		sbuf_done((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:])
		(nreg(map_([]byte("dl\x00"))))[0] = cdiv[0].dl
		(nreg(map_([]byte("dn\x00"))))[0] = (nreg(int32('d')))[0]
		(nreg(int32('d')))[0] = cdiv[0].prev_d
		(nreg(int32('h')))[0] = cdiv[0].prev_h
		(nreg(map_([]byte(".mk\x00"))))[0] = cdiv[0].prev_mk
		(nreg(map_([]byte(".ns\x00"))))[0] = cdiv[0].prev_ns
		if (int64(uintptr(unsafe.Pointer(&cdiv[0])))/int64(56) - int64(uintptr(unsafe.Pointer(&divs[0])))/int64(56)) > 0 {
			cdiv = c4goPointerArithDivSlice(cdiv, int(-1))
		} else {
			cdiv = nil
		}
	}
}

// f_divreg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:88
func f_divreg() int32 {
	if cdiv != nil {
		return cdiv[0].reg
	}
	return -1
}

// f_hpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:93
func f_hpos() int32 {
	return fmt_wid(env_fmt()) + wb_wid(env_wb())
}

// tr_divbeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:98
func tr_divbeg(args [][]byte) {
	odiv_beg()
	ren_div++
}

// tr_divend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:104
func tr_divend(args [][]byte) {
	if ren_div <= 0 {
		errdie([]byte("neatroff: diversion stack empty\n\x00"))
	}
	odiv_end()
	ren_div--
}

// tr_divvs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:112
func tr_divvs(args [][]byte) {
	ren_divvs = eval(args[1], int32('u'))
}

// tr_transparent - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:117
func tr_transparent(args [][]byte) {
	if cdiv != nil {
		sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("%s\n\x00"), args[1])
	} else {
		out([]byte("%s\n\x00"), args[1])
	}
}

// ren_page - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:129
func ren_page(force int32) {
	if noarch.Not(force) && bp_final >= 2 {
		return
	}
	(nreg(map_([]byte("nl\x00"))))[0] = 0
	(nreg(int32('d')))[0] = 0
	(nreg(int32('h')))[0] = 0
	if bp_next != 1073741824 {
		(nreg(map_([]byte("%\x00"))))[0] = bp_next
	} else {
		(nreg(map_([]byte("%\x00"))))[0] = (nreg(map_([]byte("%\x00"))))[0] + 1
	}
	bp_next = 1073741824
	(nreg(int32('%')))[0]++
	out([]byte("p%d\n\x00"), (nreg(int32('%')))[0])
	out([]byte("V%d\n\x00"), 0)
	if trap_pos(-1) == 0 {
		trap_exec(trap_reg(-1))
	}
}

// ren_first - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:145
func ren_first() int32 {
	if bp_first != 0 && cdiv == nil {
		bp_first = 0
		ren_page(1)
		return 0
	}
	return 1
}

// ren_sp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:156
func ren_sp(n int32, nodiv int32) {
	// when nodiv, do not append .sp to diversions
	// the vertical spacing before a line
	var linevs int32 = noarch.BoolToInt(noarch.Not(n))
	ren_first()
	if noarch.Not(n) && ren_div != 0 && ren_divvs != 0 && noarch.Not((nreg(int32('u')))[0]) {
		// .v at the time of diversion
		n = ren_divvs
	}
	ren_divvs = 0
	(nreg(map_([]byte(".ns\x00"))))[0] = 0
	(nreg(int32('d')))[0] += func() int32 {
		if n != 0 {
			return n
		}
		return (nreg(int32('v')))[0]
	}()
	if (nreg(int32('d')))[0] > (nreg(int32('h')))[0] {
		(nreg(int32('h')))[0] = (nreg(int32('d')))[0]
	}
	if cdiv != nil && noarch.Not(nodiv) {
		if linevs != 0 {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("%c%s %du\n\x00"), c_cc, []byte("\aV\x00"), (nreg(int32('v')))[0])
		} else {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("%csp %du\n\x00"), c_cc, func() int32 {
				if n != 0 {
					return n
				}
				return (nreg(int32('v')))[0]
			}())
		}
	}
	if cdiv == nil {
		(nreg(map_([]byte("nl\x00"))))[0] = (nreg(int32('d')))[0]
	}
}

// trap_exec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:179
func trap_exec(reg int32) {
	var cmd []byte = make([]byte, 16)
	var partial int32 = noarch.BoolToInt(ren_partial != 0 && (noarch.Not(ren_un) || ren_unbuf[0] != int32('\n')))
	if str_get(reg) != nil {
		noarch.Sprintf(cmd, []byte("%c%s %d\n\x00"), c_cc, []byte("\aP\x00"), ren_level)
		in_push(cmd, nil)
		in_push(str_get(reg), nil)
		if partial != 0 {
			in_push([]byte("\n\x00"), nil)
		}
		render_rec(func() int32 {
			ren_level++
			return ren_level
		}())
		if partial != 0 {
			// executed the trap while in the middle of an input line
			fmt_suppressnl(env_fmt())
		}
	}
}

// detect_traps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:196
func detect_traps(beg int32, end int32) int32 {
	var pos int32 = trap_pos(beg)
	return noarch.BoolToInt(pos >= 0 && (cdiv != nil || pos < (nreg(int32('p')))[0]) && pos <= end)
}

// ren_traps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:203
func ren_traps(beg int32, end int32, dosp int32) int32 {
	// return 1 if executed a trap
	var pos int32 = trap_pos(beg)
	if detect_traps(beg, end) != 0 {
		if dosp != 0 && pos > beg {
			ren_sp(pos-beg, 0)
		}
		trap_exec(trap_reg(beg))
		return 1
	}
	return 0
}

// detect_pagelimit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:215
func detect_pagelimit(ne int32) int32 {
	return noarch.BoolToInt(cdiv == nil && (nreg(map_([]byte("nl\x00"))))[0]+ne >= (nreg(int32('p')))[0])
}

// ren_pagelimit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:221
func ren_pagelimit(ne int32) int32 {
	if detect_pagelimit(ne) != 0 {
		// start a new page if needed
		ren_page(0)
		return 1
	}
	return 0
}

// down - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:231
func down(n int32) int32 {
	if ren_traps((nreg(int32('d')))[0], (nreg(int32('d')))[0]+func() int32 {
		if n != 0 {
			return n
		}
		return (nreg(int32('v')))[0]
	}(), 1) != 0 {
		// return 1 if triggered a trap
		return 1
	}
	ren_sp(func() int32 {
		if n < -((nreg(int32('d')))[0]) {
			return -((nreg(int32('d')))[0])
		}
		return n
	}(), 0)
	return ren_pagelimit(0)
}

// ren_ljust - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:240
func ren_ljust(spre []sbuf, w int32, ad int32, li int32, lI int32, ll int32) int32 {
	// line adjustment
	var ljust int32 = li
	var llen int32 = ll - lI - li
	(nreg(int32('n')))[0] = w
	if ad&3 == 0 {
		ljust += func() int32 {
			if llen > w {
				return (llen - w) / 2
			}
			return 0
		}()
	}
	if ad&3 == 2 {
		ljust += llen - w
	}
	if ljust != 0 {
		sbuf_printf(spre, []byte("%ch'%du'\x00"), c_ec, ljust)
	}
	if cdiv != nil && cdiv[0].dl < w+ljust {
		cdiv[0].dl = w + ljust
	}
	return ljust
}

// ren_out - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:257
func ren_out(beg []byte, mid []byte, end []byte) {
	if cdiv != nil {
		// append the line to the current diversion or send it to out.c
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], beg)
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], mid)
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], end)
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&cdiv[0].sbuf))[:], []byte("\n\x00"))
	} else {
		out([]byte("H%d\n\x00"), (nreg(int32('o')))[0])
		out([]byte("V%d\n\x00"), (nreg(int32('d')))[0])
		out_line(beg)
		out_line(mid)
		out_line(end)
	}
}

// ren_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:273
func ren_dir(sbuf_c4go_postfix []sbuf) {
	var fixed sbuf
	sbuf_init(c4goUnsafeConvert_sbuf(&fixed))
	dir_fix(c4goUnsafeConvert_sbuf(&fixed), sbuf_buf(sbuf_c4go_postfix))
	sbuf_done(sbuf_c4go_postfix)
	sbuf_init(sbuf_c4go_postfix)
	sbuf_append(sbuf_c4go_postfix, sbuf_buf(c4goUnsafeConvert_sbuf(&fixed)))
	sbuf_done(c4goUnsafeConvert_sbuf(&fixed))
}

// zwid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:284
func zwid() int32 {
	var g []glyph = dev_glyph([]byte("0\x00"), (nreg(int32('f')))[0])
	if g != nil {
		return font_gwid(g[0].font, dev_font((nreg(int32('f')))[0]), (nreg(int32('s')))[0], int32(g[0].wid))
	}
	return 0
}

// ren_lnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:291
func ren_lnum(spre []sbuf) {
	// append the line number to the output line
	var num []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var dig []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var wb_c4go_postfix wb
	var i int32
	wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	if (nreg(map_([]byte(".nn\x00"))))[0] <= 0 && (nreg(map_([]byte("ln\x00"))))[0]%(nreg(map_([]byte(".nM\x00"))))[0] == 0 {
		noarch.Sprintf(num, []byte("%d\x00"), (nreg(map_([]byte("ln\x00"))))[0])
	}
	wb_hmov(c4goUnsafeConvert_wb(&wb_c4go_postfix), (nreg(map_([]byte(".nI\x00"))))[0]*zwid())
	if noarch.Strlen(num) < int32(3) {
		wb_hmov(c4goUnsafeConvert_wb(&wb_c4go_postfix), int32((3-uint32(noarch.Strlen(num)))*uint32(zwid())))
	}
	for num[i] != 0 {
		dig[0] = num[func() int32 {
			defer func() {
				i++
			}()
			return i
		}()]
		wb_put(c4goUnsafeConvert_wb(&wb_c4go_postfix), dig)
	}
	wb_hmov(c4goUnsafeConvert_wb(&wb_c4go_postfix), (nreg(map_([]byte(".nS\x00"))))[0]*zwid())
	sbuf_append(spre, wb_buf(c4goUnsafeConvert_wb(&wb_c4go_postfix)))
	wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	if (nreg(map_([]byte(".nn\x00"))))[0] > 0 {
		(nreg(map_([]byte(".nn\x00"))))[0]--
	} else {
		(nreg(map_([]byte("ln\x00"))))[0]++
	}
}

// ren_mc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:317
func ren_mc(sbuf_c4go_postfix []sbuf, w int32, ljust int32) {
	// append margin character
	var wb_c4go_postfix wb
	wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	if w+ljust < (nreg(int32('l')))[0]+(nreg(map_([]byte(".mcn\x00"))))[0] {
		wb_hmov(c4goUnsafeConvert_wb(&wb_c4go_postfix), (nreg(int32('l')))[0]+(nreg(map_([]byte(".mcn\x00"))))[0]-w-ljust)
	}
	wb_putexpand(c4goUnsafeConvert_wb(&wb_c4go_postfix), env_mc())
	sbuf_append(sbuf_c4go_postfix, wb_buf(c4goUnsafeConvert_wb(&wb_c4go_postfix)))
	wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
}

// ren_line - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:329
func ren_line(line []byte, w int32, ad int32, body int32, li int32, lI int32, ll int32, els_neg int32, els_pos int32) int32 {
	// process a line and print it with ren_out()
	var sbeg sbuf
	var send sbuf
	var sbuf_c4go_postfix sbuf
	var prev_d int32
	var lspc int32
	var ljust int32
	ren_first()
	sbuf_init(c4goUnsafeConvert_sbuf(&sbeg))
	sbuf_init(c4goUnsafeConvert_sbuf(&send))
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	sbuf_append(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), line)
	// line space, ignoreing \x
	lspc = func() int32 {
		if 1 < (nreg(int32('L')))[0] {
			return (nreg(int32('L')))[0]
		}
		return 1
	}() * (nreg(int32('v')))[0]
	prev_d = (nreg(int32('d')))[0]
	if noarch.Not((nreg(map_([]byte(".ns\x00"))))[0]) || int32(line[0]) != 0 || els_neg != 0 || els_pos != 0 {
		if els_neg != 0 {
			ren_sp(-els_neg, 1)
		}
		ren_sp(0, 0)
		if int32(line[0]) != 0 && (nreg(map_([]byte(".nm\x00"))))[0] != 0 && body != 0 {
			ren_lnum(c4goUnsafeConvert_sbuf(&sbeg))
		}
		if noarch.Not(ren_div) && dir_do != 0 {
			ren_dir(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		}
		ljust = ren_ljust(c4goUnsafeConvert_sbuf(&sbeg), w, ad, li, lI, ll)
		if int32(line[0]) != 0 && body != 0 && (nreg(map_([]byte(".mc\x00"))))[0] != 0 {
			ren_mc(c4goUnsafeConvert_sbuf(&send), w, ljust)
		}
		ren_out(sbuf_buf(c4goUnsafeConvert_sbuf(&sbeg)), sbuf_buf(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix)), sbuf_buf(c4goUnsafeConvert_sbuf(&send)))
		(nreg(map_([]byte(".ns\x00"))))[0] = 0
		if els_pos != 0 {
			ren_sp(els_pos, 1)
		}
	}
	sbuf_done(c4goUnsafeConvert_sbuf(&sbeg))
	sbuf_done(c4goUnsafeConvert_sbuf(&send))
	sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	(nreg(int32('a')))[0] = els_pos
	if detect_traps(prev_d, (nreg(int32('d')))[0]) != 0 || detect_pagelimit(lspc-(nreg(int32('v')))[0]) != 0 {
		if noarch.Not(ren_pagelimit(lspc - (nreg(int32('v')))[0])) {
			ren_traps(prev_d, (nreg(int32('d')))[0], 0)
		}
		return 1
	}
	if lspc-(nreg(int32('v')))[0] != 0 && down(lspc-(nreg(int32('v')))[0]) != 0 {
		return 1
	}
	return 0
}

// ren_passline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:372
func ren_passline(fmt__c4go_postfix []fmt_) int32 {
	// read a line from fmt and send it to ren_line()
	var buf []byte
	var ll int32
	var li int32
	var lI int32
	var els_neg int32
	var els_pos int32
	var w int32
	var ret int32
	var ad int32 = (nreg(int32('j')))[0]
	ren_first()
	if noarch.Not(fmt_morewords(fmt__c4go_postfix)) {
		return 0
	}
	buf = fmt_nextline(fmt__c4go_postfix, c4goUnsafeConvert_int32(&w), c4goUnsafeConvert_int32(&li), c4goUnsafeConvert_int32(&lI), c4goUnsafeConvert_int32(&ll), c4goUnsafeConvert_int32(&els_neg), c4goUnsafeConvert_int32(&els_pos))
	if (nreg(int32('C')))[0] != 0 && noarch.Not((nreg(int32('u')))[0]) || (nreg(map_([]byte(".na\x00"))))[0] != 0 {
		ad = 1
	} else if ad&3 == 3 {
		if (nreg(map_([]byte(".td\x00"))))[0] > 0 {
			ad = 2
		} else {
			ad = 1
		}
	}
	if (nreg(map_([]byte(".ce\x00"))))[0] != 0 {
		ad = 0
	}
	ret = ren_line(buf, w, ad, 1, li, lI, ll, els_neg, els_pos)
	_ = buf
	return ret
}

// ren_fmtpop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:393
func ren_fmtpop(fmt__c4go_postfix []fmt_) int32 {
	// output formatted lines in fmt
	var ret int32
	for fmt_morelines(fmt__c4go_postfix) != 0 {
		ret = ren_passline(fmt__c4go_postfix)
	}
	return ret
}

// ren_fmtpopall - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:402
func ren_fmtpopall(fmt__c4go_postfix []fmt_) {
	for fmt_fill(fmt__c4go_postfix, 0) != 0 {
		// format and output all lines in fmt
		ren_fmtpop(fmt__c4go_postfix)
	}
	ren_fmtpop(fmt__c4go_postfix)
}

// ren_fmtword - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:410
func ren_fmtword(wb_c4go_postfix []wb) {
	for fmt_word(env_fmt(), wb_c4go_postfix) != 0 {
		// pass the given word buffer to the current line buffer (cfmt)
		ren_fmtpop(env_fmt())
	}
	wb_reset(wb_c4go_postfix)
}

// ren_br - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:418
func ren_br() int32 {
	// output current line; returns 1 if triggered a trap
	ren_first()
	ren_fmtword(env_wb())
	for fmt_fill(env_fmt(), 1) != 0 {
		ren_fmtpop(env_fmt())
	}
	return ren_fmtpop(env_fmt())
}

// tr_br - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:427
func tr_br(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	} else {
		// output the completed lines
		ren_fmtpopall(env_fmt())
	}
}

// tr_sp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:435
func tr_sp(args [][]byte) {
	var traps int32
	var n int32
	if int32(args[0][0]) == c_cc {
		traps = ren_br()
	}
	if args[1] != nil {
		n = eval(args[1], int32('v'))
	} else {
		n = (nreg(int32('v')))[0]
	}
	if n != 0 && (noarch.Not((nreg(map_([]byte(".ns\x00"))))[0]) || ren_div != 0) && noarch.Not(traps) {
		down(n)
	}
}

// tr_sv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:446
func tr_sv(args [][]byte) {
	var n int32 = eval(args[1], int32('v'))
	(nreg(map_([]byte(".sv\x00"))))[0] = 0
	if (nreg(int32('d')))[0]+n < f_nexttrap() {
		down(n)
	} else {
		(nreg(map_([]byte(".sv\x00"))))[0] = n
	}
}

// tr_ns - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:456
func tr_ns(args [][]byte) {
	(nreg(map_([]byte(".ns\x00"))))[0] = 1
}

// tr_rs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:461
func tr_rs(args [][]byte) {
	(nreg(map_([]byte(".ns\x00"))))[0] = 0
}

// tr_os - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:466
func tr_os(args [][]byte) {
	if (nreg(map_([]byte(".sv\x00"))))[0] != 0 {
		down((nreg(map_([]byte(".sv\x00"))))[0])
	}
	(nreg(map_([]byte(".sv\x00"))))[0] = 0
}

// tr_mk - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:473
func tr_mk(args [][]byte) {
	if args[1] != nil {
		num_set(map_(args[1]), (nreg(int32('d')))[0])
	} else {
		(nreg(map_([]byte(".mk\x00"))))[0] = (nreg(int32('d')))[0]
	}
}

// tr_rt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:481
func tr_rt(args [][]byte) {
	var n int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('d')))[0], int32('v'))
		}
		return (nreg(map_([]byte(".mk\x00"))))[0]
	}()
	if n >= 0 && n < (nreg(int32('d')))[0] {
		ren_sp(n-(nreg(int32('d')))[0], 0)
	}
}

// tr_ne - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:488
func tr_ne(args [][]byte) {
	var n int32 = func() int32 {
		if args[1] != nil {
			return eval(args[1], int32('v'))
		}
		return (nreg(int32('v')))[0]
	}()
	if noarch.Not(ren_first()) {
		return
	}
	if noarch.Not(ren_traps((nreg(int32('d')))[0], (nreg(int32('d')))[0]+n-1, 1)) {
		ren_pagelimit(n)
	}
}

// ren_ejectpage - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:497
func ren_ejectpage(br int32) {
	ren_first()
	bp_ejected = (nreg(int32('%')))[0]
	if br != 0 {
		ren_br()
	}
	for (nreg(int32('%')))[0] == bp_ejected && cdiv == nil {
		if detect_traps((nreg(int32('d')))[0], (nreg(int32('p')))[0]) != 0 {
			ren_traps((nreg(int32('d')))[0], (nreg(int32('p')))[0], 1)
		} else {
			bp_ejected = 0
			ren_page(0)
		}
	}
}

// tr_bp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:513
func tr_bp(args [][]byte) {
	if cdiv == nil && (args[1] != nil || noarch.Not((nreg(map_([]byte(".ns\x00"))))[0])) {
		if args[1] != nil {
			bp_next = eval_re(args[1], (nreg(map_([]byte("%\x00"))))[0], 0)
		}
		ren_ejectpage(noarch.BoolToInt(int32(args[0][0]) == c_cc))
	}
}

// tr_pn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:522
func tr_pn(args [][]byte) {
	if args[1] != nil {
		bp_next = eval_re(args[1], (nreg(map_([]byte("%\x00"))))[0], 0)
	}
}

// ren_ps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:528
func ren_ps(s []byte) {
	var ps int32 = func() int32 {
		if s == nil || noarch.Not(s[0]) || noarch.Not(noarch.Strcmp([]byte("0\x00"), s)) {
			return (nreg(map_([]byte(".s0\x00"))))[0] * (dev_res / 72)
		}
		return eval_re(s, (nreg(int32('s')))[0]*(dev_res/72), int32('p'))
	}()
	(nreg(map_([]byte(".s0\x00"))))[0] = (nreg(int32('s')))[0]
	if 1 < (ps+dev_res/72/2)/(dev_res/72) {
		(nreg(int32('s')))[0] = (ps + dev_res/72/2) / (dev_res / 72)
	} else {
		(nreg(int32('s')))[0] = 1
	}
}

// tr_ps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:536
func tr_ps(args [][]byte) {
	ren_ps(args[1])
}

// tr_ll - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:541
func tr_ll(args [][]byte) {
	var ll int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('l')))[0], int32('m'))
		}
		return (nreg(map_([]byte(".l0\x00"))))[0]
	}()
	(nreg(map_([]byte(".l0\x00"))))[0] = (nreg(int32('l')))[0]
	if 0 < ll {
		(nreg(int32('l')))[0] = ll
	} else {
		(nreg(int32('l')))[0] = 0
	}
}

// tr_in - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:548
func tr_in(args [][]byte) {
	var in int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('i')))[0], int32('m'))
		}
		return (nreg(map_([]byte(".i0\x00"))))[0]
	}()
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(map_([]byte(".i0\x00"))))[0] = (nreg(int32('i')))[0]
	if 0 < in {
		(nreg(int32('i')))[0] = in
	} else {
		(nreg(int32('i')))[0] = 0
	}
	(nreg(map_([]byte(".ti\x00"))))[0] = -1
}

// tr_ti - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:558
func tr_ti(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	if args[1] != nil {
		(nreg(map_([]byte(".ti\x00"))))[0] = eval_re(args[1], (nreg(int32('i')))[0], int32('m'))
	}
}

// tr_l2r - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:566
func tr_l2r(args [][]byte) {
	dir_do = 1
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(map_([]byte(".td\x00"))))[0] = 0
	(nreg(map_([]byte(".cd\x00"))))[0] = 0
}

// tr_r2l - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:575
func tr_r2l(args [][]byte) {
	dir_do = 1
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(map_([]byte(".td\x00"))))[0] = 1
	(nreg(map_([]byte(".cd\x00"))))[0] = 1
}

// tr_in2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:584
func tr_in2(args [][]byte) {
	var I int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('I')))[0], int32('m'))
		}
		return (nreg(map_([]byte(".I0\x00"))))[0]
	}()
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(map_([]byte(".I0\x00"))))[0] = (nreg(int32('I')))[0]
	if 0 < I {
		(nreg(int32('I')))[0] = I
	} else {
		(nreg(int32('I')))[0] = 0
	}
	(nreg(map_([]byte(".tI\x00"))))[0] = -1
}

// tr_ti2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:594
func tr_ti2(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	if args[1] != nil {
		(nreg(map_([]byte(".tI\x00"))))[0] = eval_re(args[1], (nreg(int32('I')))[0], int32('m'))
	}
}

// ren_ft - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:602
func ren_ft(s []byte) {
	var fn int32 = func() int32 {
		if s == nil || noarch.Not(s[0]) || noarch.Not(noarch.Strcmp([]byte("P\x00"), s)) {
			return (nreg(map_([]byte(".f0\x00"))))[0]
		}
		return dev_pos(s)
	}()
	if fn < 0 {
		errmsg([]byte("neatroff: failed to mount <%s>\n\x00"), s)
	} else {
		(nreg(map_([]byte(".f0\x00"))))[0] = (nreg(int32('f')))[0]
		(nreg(int32('f')))[0] = fn
	}
}

// tr_ft - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:613
func tr_ft(args [][]byte) {
	ren_ft(args[1])
}

// tr_fp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:618
func tr_fp(args [][]byte) {
	var pos int32
	if args[2] == nil {
		return
	}
	if int32(((__ctype_b_loc())[0])[int32(uint8(args[1][0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		pos = noarch.Atoi(args[1])
	} else {
		pos = -1
	}
	if dev_mnt(pos, args[2], func() []byte {
		if args[3] != nil {
			return args[3]
		}
		return args[2]
	}()) < 0 {
		errmsg([]byte("neatroff: failed to mount <%s>\n\x00"), args[2])
	}
}

// tr_nf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:628
func tr_nf(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(int32('u')))[0] = 0
}

// tr_fi - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:635
func tr_fi(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	(nreg(int32('u')))[0] = 1
}

// tr_ce - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:642
func tr_ce(args [][]byte) {
	if int32(args[0][0]) == c_cc {
		ren_br()
	}
	if args[1] != nil {
		(nreg(map_([]byte(".ce\x00"))))[0] = noarch.Atoi(args[1])
	} else {
		(nreg(map_([]byte(".ce\x00"))))[0] = 1
	}
}

// tr_fc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:649
func tr_fc(args [][]byte) {
	var fa []byte = args[1]
	var fb []byte = args[2]
	if fa != nil && charread((*[1000000][]byte)(unsafe.Pointer(&fa))[:], c_fa) >= 0 {
		if fb == nil || charread((*[1000000][]byte)(unsafe.Pointer(&fb))[:], c_fb) < 0 {
			noarch.Strcpy(c_fb, []byte(" \x00"))
		}
	} else {
		c_fa[0] = '\x00'
		c_fb[0] = '\x00'
	}
}

// ren_cl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:662
func ren_cl(s []byte) {
	var m int32 = func() int32 {
		if s == nil || noarch.Not(s[0]) {
			return (nreg(map_([]byte(".m0\x00"))))[0]
		}
		return clr_get(s)
	}()
	(nreg(map_([]byte(".m0\x00"))))[0] = (nreg(int32('m')))[0]
	(nreg(int32('m')))[0] = m
}

// tr_cl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:669
func tr_cl(args [][]byte) {
	ren_cl(args[1])
}

// tr_ab - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:674
func tr_ab(args [][]byte) {
	noarch.Fprintf(noarch.Stderr, []byte("%s\n\x00"), args[1])
	ren_aborted = 1
}

// ren_cmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:680
func ren_cmd(wb_c4go_postfix []wb, c int32, arg []byte) {
	switch c {
	case ' ':
		wb_hmov(wb_c4go_postfix, font_swid(dev_font((nreg(int32('f')))[0]), (nreg(int32('s')))[0], (nreg(map_([]byte(".ss\x00"))))[0]))
	case 'b':
		ren_bcmd(wb_c4go_postfix, arg)
	case 'c':
		wb_setpart(wb_c4go_postfix)
	case 'D':
		ren_dcmd(wb_c4go_postfix, arg)
	case 'd':
		wb_vmov(wb_c4go_postfix, (nreg(int32('s')))[0]*dev_res/72/2)
	case 'f':
		ren_ft(arg)
	case 'h':
		wb_hmov(wb_c4go_postfix, eval(arg, int32('m')))
	case 'j':
		wb_setcost(wb_c4go_postfix, eval(arg, 0))
	case 'k':
		num_set(map_(arg), func() int32 {
			if (int64(uintptr(unsafe.Pointer(&wb_c4go_postfix[0])))/int64(8320) - func() int64 {
				c4go_temp_name := env_wb()
				return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
			}()) == 0 {
				return f_hpos() - (nreg(map_([]byte(".b0\x00"))))[0]
			}
			return wb_wid(wb_c4go_postfix)
		}())
	case 'L':
		ren_vlcmd(wb_c4go_postfix, arg)
	case 'l':
		ren_hlcmd(wb_c4go_postfix, arg)
	case 'm':
		ren_cl(arg)
	case 'o':
		ren_ocmd(wb_c4go_postfix, arg)
	case 'p':
		if (int64(uintptr(unsafe.Pointer(&wb_c4go_postfix[0])))/int64(8320) - func() int64 {
			c4go_temp_name := env_wb()
			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
		}()) == 0 {
			for fmt_fillreq(env_fmt()) != 0 {
				ren_fmtpop(env_fmt())
			}
		}
	case 'r':
		wb_vmov(wb_c4go_postfix, -((nreg(int32('s')))[0] * dev_res / 72))
	case 's':
		ren_ps(arg)
	case 'u':
		wb_vmov(wb_c4go_postfix, -((nreg(int32('s')))[0]*dev_res/72)/2)
	case 'v':
		wb_vmov(wb_c4go_postfix, eval(arg, int32('v')))
	case 'X':
		wb_etc(wb_c4go_postfix, arg)
	case 'x':
		wb_els(wb_c4go_postfix, eval(arg, int32('v')))
	case 'Z':
		ren_zcmd(wb_c4go_postfix, arg)
	case '0':
		wb_hmov(wb_c4go_postfix, zwid())
	case '|':
		wb_hmov(wb_c4go_postfix, (nreg(int32('s')))[0]*dev_res/72/6)
	case '&':
		wb_hmov(wb_c4go_postfix, 0)
	case '^':
		wb_hmov(wb_c4go_postfix, (nreg(int32('s')))[0]*dev_res/72/12)
	case '/':
		wb_italiccorrection(wb_c4go_postfix)
	case ',':
		wb_italiccorrectionleft(wb_c4go_postfix)
	case '<':
		fallthrough
	case '>':
		(nreg(map_([]byte(".cd\x00"))))[0] = noarch.BoolToInt(c == int32('<'))
		wb_flushdir(wb_c4go_postfix)
		break
	}
}

// ren_put - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:778
func ren_put(wb_c4go_postfix []wb, c []byte, next func() int32, back func(int32)) {
	// insert a character, escape sequence, field or etc into wb
	var w int32
	if int32(c[0]) == int32(' ') || int32(c[0]) == int32('\n') {
		wb_put(wb_c4go_postfix, c)
		return
	}
	if int32(c[0]) == int32('\t') || int32(c[0]) == int32('\x01') {
		ren_tab(wb_c4go_postfix, func() []byte {
			if int32(c[0]) == int32('\t') {
				return env_tc()
			}
			return env_lc()
		}(), next, back)
		return
	}
	if int32(c_fa[0]) != 0 && noarch.Not(noarch.Strcmp(c_fa, c)) {
		ren_field(wb_c4go_postfix, next, back)
		return
	}
	if int32(c[0]) == c_ec {
		if int32(c[1]) == int32('z') {
			w = wb_wid(wb_c4go_postfix)
			ren_char(wb_c4go_postfix, next, back)
			wb_hmov(wb_c4go_postfix, w-wb_wid(wb_c4go_postfix))
			return
		}
		if noarch.Strchr([]byte(" bCcDdefHhjkLlmNoprSsuvXxZz0^|!{}&/,<>\x00"), int32(c[1])) != nil {
			var arg []byte
			if noarch.Strchr([]byte("*fgkmns\x00"), int32(c[1])) != nil {
				arg = unquotednext(int32(c[1]), next, back)
			}
			if noarch.Strchr([]byte("bCDhHjlLNoRSvwxXZ?\x00"), int32(c[1])) != nil {
				arg = quotednext(next, back)
			}
			if int32(c[1]) == int32('e') {
				noarch.Snprintf(c, 32, []byte("%c%c\x00"), c_ec, c_ec)
			} else if int32(c[1]) == int32('N') {
				noarch.Snprintf(c, 32, []byte("GID=%s\x00"), arg)
			} else {
				ren_cmd(wb_c4go_postfix, int32(c[1]), arg)
				_ = arg
				return
			}
			_ = arg
		}
	}
	if ren_div != 0 {
		wb_putraw(wb_c4go_postfix, c)
		return
	}
	if cdef_map(c, (nreg(int32('f')))[0]) != nil {
		// .char characters
		wb_putexpand(wb_c4go_postfix, c)
	} else {
		wb_put(wb_c4go_postfix, c)
	}
}

// ren_char - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:829
func ren_char(wb_c4go_postfix []wb, next func() int32, back func(int32)) int32 {
	// read one character and place it inside wb buffer
	var c []byte = make([]byte, 128)
	if charnext(c, next, back) < 0 {
		return -1
	}
	ren_put(wb_c4go_postfix, c, next, back)
	return 0
}

// ren_chardel - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:839
func ren_chardel(wb_c4go_postfix []wb, next func() int32, back func(int32), d1 []byte, d2 []byte) int32 {
	// like ren_char(); return 1 if d1 was read and 2 if d2 was read
	var c []byte = make([]byte, 128)
	if charnext(c, next, back) < 0 {
		return -1
	}
	if d1 != nil && noarch.Not(noarch.Strcmp(d1, c)) {
		return 1
	}
	if d2 != nil && noarch.Not(noarch.Strcmp(d2, c)) {
		return 2
	}
	ren_put(wb_c4go_postfix, c, next, back)
	return 0
}

// ren_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:854
func ren_wid(next func() int32, back func(int32)) int32 {
	// read the argument of \w and push its width
	var delim []byte = make([]byte, 32)
	var c int32
	var n int32
	var wb_c4go_postfix wb
	wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	charnext(delim, next, back)
	odiv_beg()
	c = next()
	for c >= 0 && c != int32('\n') {
		back(c)
		if ren_chardel(c4goUnsafeConvert_wb(&wb_c4go_postfix), next, back, delim, nil) != 0 {
			break
		}
		c = next()
	}
	odiv_end()
	n = wb_wid(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	wb_wconf(c4goUnsafeConvert_wb(&wb_c4go_postfix), (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("ct\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("st\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("sb\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("bbllx\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("bblly\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("bburx\x00"))))[0])))[:], (*[1000000]int32)(unsafe.Pointer(&((nreg(map_([]byte("bbury\x00"))))[0])))[:])
	wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	return n
}

// ren_until - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:877
func ren_until(wb_c4go_postfix []wb, next func() int32, back func(int32), d1 []byte, d2 []byte) int32 {
	// return 1 if d1 was read and 2 if d2 was read
	var c int32
	var ret int32
	c = next()
	for c >= 0 && c != int32('\n') {
		back(c)
		ret = ren_chardel(wb_c4go_postfix, next, back, d1, d2)
		if ret != 0 {
			return ret
		}
		c = next()
	}
	if c == int32('\n') {
		back(c)
	}
	return 0
}

// ren_untilmap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:895
func ren_untilmap(wb_c4go_postfix []wb, next func() int32, back func(int32), end []byte, src []byte, dst []byte) int32 {
	// like ren_until(); map src to dst
	var ret int32
	for (func() int32 {
		ret = ren_until(wb_c4go_postfix, next, back, src, end)
		return ret
	}()) == 1 {
		sstr_push(dst)
		ren_until(wb_c4go_postfix, sstr_next, sstr_back, end, nil)
		sstr_pop()
	}
	return 0
}

// wb_cpy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:907
func wb_cpy(dst []wb, src []wb, left int32) {
	wb_hmov(dst, left-wb_wid(dst))
	wb_cat(dst, src)
}

// ren_tl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:913
func ren_tl(next func() int32, back func(int32)) {
	var wb_c4go_postfix wb
	var wb2 wb
	var pgnum []byte
	var delim []byte = make([]byte, 32)
	ren_first()
	pgnum = num_str(map_([]byte("%\x00")))
	wb_init(c4goUnsafeConvert_wb(&wb_c4go_postfix))
	wb_init(c4goUnsafeConvert_wb(&wb2))
	charnext(delim, next, back)
	if noarch.Not(noarch.Strcmp([]byte("\n\x00"), delim)) {
		back(int32('\n'))
	}
	// the left-adjusted string
	ren_untilmap(c4goUnsafeConvert_wb(&wb2), next, back, delim, c_pc, pgnum)
	wb_cpy(c4goUnsafeConvert_wb(&wb_c4go_postfix), c4goUnsafeConvert_wb(&wb2), 0)
	// the centered string
	ren_untilmap(c4goUnsafeConvert_wb(&wb2), next, back, delim, c_pc, pgnum)
	wb_cpy(c4goUnsafeConvert_wb(&wb_c4go_postfix), c4goUnsafeConvert_wb(&wb2), ((nreg(map_([]byte(".lt\x00"))))[0]-wb_wid(c4goUnsafeConvert_wb(&wb2)))/2)
	// the right-adjusted string
	ren_untilmap(c4goUnsafeConvert_wb(&wb2), next, back, delim, c_pc, pgnum)
	wb_cpy(c4goUnsafeConvert_wb(&wb_c4go_postfix), c4goUnsafeConvert_wb(&wb2), (nreg(map_([]byte(".lt\x00"))))[0]-wb_wid(c4goUnsafeConvert_wb(&wb2)))
	// flushing the line
	ren_line(wb_buf(c4goUnsafeConvert_wb(&wb_c4go_postfix)), wb_wid(c4goUnsafeConvert_wb(&wb_c4go_postfix)), 1, 0, 0, 0, (nreg(map_([]byte(".lt\x00"))))[0], wb_c4go_postfix.els_neg, wb_c4go_postfix.els_pos)
	wb_done(c4goUnsafeConvert_wb(&wb2))
	wb_done(c4goUnsafeConvert_wb(&wb_c4go_postfix))
}

// ren_field - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:941
func ren_field(wb_c4go_postfix []wb, next func() int32, back func(int32)) {
	var wbs []wb = make([]wb, 32)
	var i int32
	var n int32
	var wid int32
	var left int32
	var right int32
	var cur_left int32
	var pad int32
	var rem int32
	for uint32(n) < 266240/8320 {
		wb_init(wbs[n:])
		if ren_until(wbs[func() int32 {
			defer func() {
				n++
			}()
			return n
		}():], next, back, c_fb, c_fa) != 1 {
			break
		}
	}
	if (int64(uintptr(unsafe.Pointer(&wb_c4go_postfix[0])))/int64(8320) - func() int64 {
		c4go_temp_name := env_wb()
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}()) == 0 {
		left = f_hpos()
	} else {
		left = wb_wid(wb_c4go_postfix)
	}
	right = tab_next(left)
	for i = 0; i < n; i++ {
		wid += wb_wid(wbs[i:])
	}
	pad = (right - left - wid) / func() int32 {
		if n > 1 {
			return n - 1
		}
		return 1
	}()
	rem = (right - left - wid) % func() int32 {
		if n > 1 {
			return n - 1
		}
		return 1
	}()
	for i = 0; i < n; i++ {
		if i == 0 {
			cur_left = left
		} else if i == n-1 {
			cur_left = right - wb_wid(wbs[i:])
		} else {
			cur_left = wb_wid(wb_c4go_postfix) + pad + noarch.BoolToInt(i+rem >= n)
		}
		wb_cpy(wb_c4go_postfix, wbs[i:], cur_left)
		wb_done(wbs[i:])
	}
}

// ren_tab - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:971
func ren_tab(wb_c4go_postfix []wb, tc []byte, next func() int32, back func(int32)) {
	var t wb
	var pos int32 = func() int32 {
		if (int64(uintptr(unsafe.Pointer(&wb_c4go_postfix[0])))/int64(8320) - func() int64 {
			c4go_temp_name := env_wb()
			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
		}()) == 0 {
			return f_hpos()
		}
		return wb_wid(wb_c4go_postfix)
	}()
	// insertion position
	var ins int32 = tab_next(pos)
	// tab type
	var typ int32 = tab_type(pos)
	var c int32
	wb_init(c4goUnsafeConvert_wb(&t))
	if typ == int32('R') || typ == int32('C') {
		c = next()
		for c >= 0 && c != int32('\n') && c != int32('\t') && c != int32('\x01') {
			back(c)
			ren_char(c4goUnsafeConvert_wb(&t), next, back)
			c = next()
		}
		back(c)
	}
	if typ == int32('C') {
		ins -= wb_wid(c4goUnsafeConvert_wb(&t)) / 2
	}
	if typ == int32('R') {
		ins -= wb_wid(c4goUnsafeConvert_wb(&t))
	}
	if noarch.Not(tc[0]) || ins <= pos {
		wb_hmov(wb_c4go_postfix, ins-pos)
	} else {
		ren_hline(wb_c4go_postfix, ins-pos, tc)
	}
	wb_cat(wb_c4go_postfix, c4goUnsafeConvert_wb(&t))
	wb_done(c4goUnsafeConvert_wb(&t))
}

// ren_parse - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1001
func ren_parse(wb_c4go_postfix []wb, s []byte) int32 {
	// parse characters and troff requests of s and append them to wb
	var c int32
	odiv_beg()
	sstr_push(s)
	c = sstr_next()
	for c >= 0 {
		sstr_back(c)
		if ren_char(wb_c4go_postfix, sstr_next, sstr_back) != 0 {
			break
		}
		c = sstr_next()
	}
	sstr_pop()
	odiv_end()
	return 0
}

// tr_popren - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1019
func tr_popren(args [][]byte) {
	// cause nested render_rec() to exit
	if args[1] != nil {
		ren_level = noarch.Atoi(args[1])
	} else {
		ren_level = 0
	}
}

// render_rec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1027
func render_rec(level int32) int32 {
	// read characters from tr.c and pass the rendered lines to out.c
	var c int32
	for ren_level >= level {
		for noarch.Not(ren_un) && noarch.Not(tr_nextreq()) {
			if ren_level < level {
				break
			}
		}
		if ren_level < level {
			break
		}
		if ren_aborted != 0 {
			return 1
		}
		c = ren_next()
		if c < 0 {
			if bp_final >= 2 {
				break
			}
			if bp_final == 0 {
				bp_final = 1
				if trap_em >= 0 {
					trap_exec(trap_em)
				}
			} else {
				bp_final = 2
				ren_ejectpage(1)
			}
		}
		if c >= 0 {
			ren_partial = noarch.BoolToInt(c != int32('\n'))
		}
		if c == int32(' ') || c == int32('\n') {
			if noarch.Not(wb_part(env_wb())) {
				// add cwb (the current word) to cfmt
				// not after a \c
				ren_fmtword(env_wb())
				if c == int32('\n') {
					for fmt_newline(env_fmt()) != 0 {
						ren_fmtpop(env_fmt())
					}
				}
				if !((nreg(int32('u')))[0] != 0 && noarch.Not((nreg(map_([]byte(".na\x00"))))[0]) && noarch.Not((nreg(map_([]byte(".ce\x00"))))[0]) && (nreg(int32('j')))[0]&4 == 4) {
					ren_fmtpopall(env_fmt())
				}
				if c == int32(' ') {
					fmt_space(env_fmt())
				}
			}
		}
		if c == int32(' ') || c == int32('\n') || c < 0 {
			// flush the line if necessary
			ren_fmtpop(env_fmt())
		}
		if c == int32('\n') || ren_nl != 0 {
			// end or start of input line
			(nreg(map_([]byte(".b0\x00"))))[0] = f_hpos()
		}
		if c == int32('\n') && (nreg(map_([]byte(".it\x00"))))[0] != 0 && func() int32 {
			tempVar1 := &(nreg(map_([]byte(".itn\x00"))))[0]
			*tempVar1--
			return *tempVar1
		}() == 0 {
			trap_exec((nreg(map_([]byte(".it\x00"))))[0])
		}
		if c == int32('\n') && noarch.Not(wb_part(env_wb())) {
			if 0 < (nreg(map_([]byte(".ce\x00"))))[0]-1 {
				(nreg(map_([]byte(".ce\x00"))))[0] = (nreg(map_([]byte(".ce\x00"))))[0] - 1
			} else {
				(nreg(map_([]byte(".ce\x00"))))[0] = 0
			}
		}
		if c != int32(' ') && c >= 0 {
			ren_back(c)
			ren_char(env_wb(), ren_next, ren_back)
		}
		if c >= 0 {
			ren_nl = noarch.BoolToInt(c == int32('\n'))
		}
	}
	return 0
}

// render - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1086
func render() int32 {
	// render input words
	(nreg(map_([]byte("nl\x00"))))[0] = -1
	for noarch.Not(tr_nextreq()) {
	}
	// transition to the first page
	ren_first()
	render_rec(0)
	bp_final = 3
	if fmt_morewords(env_fmt()) != 0 {
		ren_page(1)
	}
	ren_br()
	return 0
}

// tpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1104
// trap handling
// trap positions
var tpos []int32 = make([]int32, 1024)

// treg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1105
// trap registers
var treg []int32 = make([]int32, 1024)

// ntraps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1106
var ntraps int32

// trap_first - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1108
func trap_first(pos int32) int32 {
	var best int32 = -1
	var i int32
	for i = 0; i < ntraps; i++ {
		if treg[i] >= 0 && func() int32 {
			if tpos[i] < 0 {
				return (nreg(int32('p')))[0] + tpos[i]
			}
			return tpos[i]
		}() > pos {
			if best < 0 || func() int32 {
				if tpos[i] < 0 {
					return (nreg(int32('p')))[0] + tpos[i]
				}
				return tpos[i]
			}() < func() int32 {
				if tpos[best] < 0 {
					return (nreg(int32('p')))[0] + tpos[best]
				}
				return tpos[best]
			}() {
				best = i
			}
		}
	}
	return best
}

// trap_byreg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1119
func trap_byreg(reg int32) int32 {
	var i int32
	for i = 0; i < ntraps; i++ {
		if treg[i] == reg {
			return i
		}
	}
	return -1
}

// trap_bypos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1128
func trap_bypos(reg int32, pos int32) int32 {
	var i int32
	for i = 0; i < ntraps; i++ {
		if treg[i] >= 0 && func() int32 {
			if tpos[i] < 0 {
				return (nreg(int32('p')))[0] + tpos[i]
			}
			return tpos[i]
		}() == pos {
			if reg == -1 || treg[i] == reg {
				return i
			}
		}
	}
	return -1
}

// tr_wh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1138
func tr_wh(args [][]byte) {
	var reg int32
	var pos int32
	var id int32
	if args[1] == nil {
		return
	}
	pos = eval(args[1], int32('v'))
	id = trap_bypos(-1, pos)
	if args[2] == nil {
		if id >= 0 {
			treg[id] = -1
		}
		return
	}
	reg = map_(args[2])
	if id < 0 {
		// find an unused position in treg[]
		id = trap_byreg(-1)
	}
	if id < 0 {
		id = func() int32 {
			defer func() {
				ntraps++
			}()
			return ntraps
		}()
	}
	tpos[id] = pos
	treg[id] = reg
}

// tr_ch - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1159
func tr_ch(args [][]byte) {
	var reg int32
	var id int32
	if args[1] == nil {
		return
	}
	reg = map_(args[1])
	id = trap_byreg(reg)
	if id >= 0 {
		if args[2] != nil {
			tpos[id] = eval(args[2], int32('v'))
		} else {
			treg[id] = -1
		}
	}
}

// tr_dt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1175
func tr_dt(args [][]byte) {
	if cdiv == nil {
		return
	}
	if args[2] != nil {
		cdiv[0].tpos = eval(args[1], int32('v'))
		cdiv[0].treg = map_(args[2])
	} else {
		cdiv[0].treg = -1
	}
}

// tr_em - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1187
func tr_em(args [][]byte) {
	if args[1] != nil {
		trap_em = map_(args[1])
	} else {
		trap_em = -1
	}
}

// trap_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1192
func trap_pos(pos int32) int32 {
	var ret int32 = trap_first(pos)
	if bp_final >= 3 {
		return -1
	}
	if cdiv != nil {
		if cdiv[0].treg != 0 && cdiv[0].tpos > pos {
			return cdiv[0].tpos
		}
		return -1
	}
	if ret >= 0 {
		if tpos[ret] < 0 {
			return (nreg(int32('p')))[0] + tpos[ret]
		}
		return tpos[ret]
	}
	return -1
}

// trap_reg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1202
func trap_reg(pos int32) int32 {
	var ret int32 = trap_first(pos)
	if cdiv != nil {
		if cdiv[0].treg != 0 && cdiv[0].tpos > pos {
			return cdiv[0].treg
		}
		return -1
	}
	if ret >= 0 {
		return treg[ret]
	}
	return -1
}

// f_nexttrap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/ren.c:1210
func f_nexttrap() int32 {
	var pos int32 = trap_pos((nreg(int32('d')))[0])
	if cdiv != nil {
		if pos >= 0 {
			return pos
		}
		return 2147483647
	}
	return func() int32 {
		if pos >= 0 && pos < (nreg(int32('p')))[0] {
			return pos
		}
		return (nreg(int32('p')))[0]
	}() - (nreg(int32('d')))[0]
}

// errmsg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:24
func errmsg(fmt_ []byte, c4goArgs ...interface{}) {
	//
	// * NEATROFF TYPESETTING SYSTEM
	// *
	// * Copyright (C) 2012-2016 Ali Gholami Rudi <ali at rudi dot ir>
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
	var ap *va_list
	va_start(ap, fmt_)
	noarch.Vfprintf(noarch.Stderr, fmt_, ap)
	va_end(ap)
}

// errdie - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:32
func errdie(msg []byte) {
	noarch.Fprintf(noarch.Stderr, []byte("%s\x00"), msg)
	unix.Exit(1)
}

// mextend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:38
func mextend(old interface{}, oldsz int32, newsz int32, memsz int32) interface{} {
	var new_ interface{} = xmalloc(newsz * memsz)
	memcpy(new_, old, uint32(oldsz*memsz))
	noarch.Memset(new_[0+oldsz*memsz:].([]byte), byte(0), uint32((newsz-oldsz)*memsz))
	_ = old
	return new_
}

// xmalloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:47
func xmalloc(len_ int32) interface{} {
	var m interface{} = make([]byte, uint32(len_))
	if m == nil {
		errdie([]byte("neatroff: malloc() failed\n\x00"))
	}
	return m
}

// xopens - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:55
func xopens(path []byte) int32 {
	var filp *noarch.File = noarch.Fopen(path, []byte("r\x00"))
	if filp != nil {
		noarch.Fclose(filp)
	}
	return noarch.BoolToInt(filp != nil)
}

// cmddef - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:64
func cmddef(arg []byte, reg []int32, def [][]byte) {
	// parse the argument of -r and -d options
	var regname []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var eq []byte = noarch.Strchr(arg, int32('='))
	memcpy(regname, arg, uint32(func() int32 {
		if eq != nil {
			if 128-1 < int32((int64(uintptr(unsafe.Pointer(&eq[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&arg[0])))/int64(1))) {
				return 128 - 1
			}
			return int32((int64(uintptr(unsafe.Pointer(&eq[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&arg[0])))/int64(1)))
		}
		return 1
	}()))
	reg[0] = map_(regname)
	if eq != nil {
		def[0] = eq[0+1:]
	} else {
		def[0] = arg[0+1:]
	}
}

// cmdmac - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:74
func cmdmac(dir []byte, arg []byte) int32 {
	// find the macro specified with -m option
	var path []byte = make([]byte, 1024)
	noarch.Snprintf(path, int32(1024), []byte("%s/%s.tmac\x00"), dir, arg)
	if noarch.Not(xopens(path)) {
		noarch.Snprintf(path, int32(1024), []byte("%s/tmac.%s\x00"), dir, arg)
	}
	if noarch.Not(xopens(path)) {
		noarch.Snprintf(path, int32(1024), []byte("%s/%s\x00"), dir, arg)
	}
	if noarch.Not(xopens(path)) {
		return 1
	}
	in_queue(path)
	return 0
}

// usage - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:88
var usage []byte = []byte("Usage: neatroff [options] input\n\nOptions:\n  -mx   \tinclude macro x\n  -rx=y \tset number register x to y\n  -dx=y \tdefine string register x as y\n  -C    \tenable compatibility mode\n  -Tdev \tset output device\n  -Fdir \tset font directory (./)\n  -Mdir \tset macro directory (./)\n\x00")

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/roff.c:99
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	var fontdir []byte = []byte("./\x00")
	var macrodir []byte = []byte("./\x00")
	var mac []byte
	var def []byte
	var dev []byte = []byte("utf\x00")
	var reg int32
	var ret int32
	var i int32
	for i = 1; i < argc; i++ {
		if int32(argv[i][0]) != int32('-') || noarch.Not(argv[i][1]) {
			break
		}
		switch int32(argv[i][1]) {
		case 'C':
			(nreg(int32('C')))[0] = 1
		case 'm':
			mac = (argv[i])[0+2:]
			if noarch.Strchr(mac, int32('/')) != nil || cmdmac(macrodir, mac) != 0 && cmdmac([]byte(".\x00"), mac) != 0 {
				in_queue(mac)
			}
		case 'r':
			cmddef(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}(), c4goUnsafeConvert_int32(&reg), (*[1000000][]byte)(unsafe.Pointer(&def))[:])
			num_set(reg, eval_re(def, (nreg(reg))[0], int32('u')))
		case 'd':
			cmddef(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}(), c4goUnsafeConvert_int32(&reg), (*[1000000][]byte)(unsafe.Pointer(&def))[:])
			str_set(reg, def)
		case 'F':
			fontdir = func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}()
		case 'M':
			macrodir = func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}()
		case 'T':
			dev = func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}()
		default:
			noarch.Printf([]byte("%s\x00"), usage)
			return
		}
	}
	if dev_open(fontdir, dev) != 0 {
		noarch.Fprintf(noarch.Stderr, []byte("neatroff: cannot open device %s\n\x00"), dev)
		noarch.Exit(int32(1))
	}
	hyph_init()
	env_init()
	tr_init()
	if i == argc {
		// reading from standard input
		in_queue(nil)
	}
	for ; i < argc; i++ {
		in_queue(func() []byte {
			if noarch.Not(noarch.Strcmp([]byte("-\x00"), argv[i])) {
				return nil
			}
			return argv[i]
		}())
	}
	out([]byte("s%d\n\x00"), (nreg(int32('s')))[0])
	out([]byte("f%d\n\x00"), (nreg(int32('f')))[0])
	ret = render()
	out([]byte("V%d\n\x00"), (nreg(int32('p')))[0])
	hyph_done()
	tr_done()
	env_done()
	dev_close()
	map_done()
	dir_done()
	noarch.Exit(int32(ret))
}

// sbuf_extend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:11
func sbuf_extend(sbuf_c4go_postfix []sbuf, amount int32) {
	// variable length string buffer
	sbuf_c4go_postfix[0].sz = (amount + 512 - 1) & ^(512 - 1)
	sbuf_c4go_postfix[0].s = mextend(sbuf_c4go_postfix[0].s, sbuf_c4go_postfix[0].n, sbuf_c4go_postfix[0].sz, int32(1)).([]byte)
}

// sbuf_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:17
func sbuf_init(sbuf_c4go_postfix []sbuf) {
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&sbuf_c4go_postfix[0]))) / int64(1))))[:], byte(0), 24)
	sbuf_extend(sbuf_c4go_postfix, 512)
}

// sbuf_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:23
func sbuf_add(sbuf_c4go_postfix []sbuf, c int32) {
	if sbuf_c4go_postfix[0].n+2 >= sbuf_c4go_postfix[0].sz {
		sbuf_extend(sbuf_c4go_postfix, sbuf_c4go_postfix[0].sz*2)
	}
	sbuf_c4go_postfix[0].s[func() int32 {
		tempVar1 := &sbuf_c4go_postfix[0].n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}()] = byte(c)
}

// sbuf_append - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:30
func sbuf_append(sbuf_c4go_postfix []sbuf, s []byte) {
	var len_ int32 = noarch.Strlen(s)
	if sbuf_c4go_postfix[0].n+len_+1 >= sbuf_c4go_postfix[0].sz {
		sbuf_extend(sbuf_c4go_postfix, sbuf_c4go_postfix[0].n+len_+1)
	}
	memcpy(sbuf_c4go_postfix[0].s[0+sbuf_c4go_postfix[0].n:], s, uint32(len_))
	sbuf_c4go_postfix[0].n += len_
}

// sbuf_printf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:39
func sbuf_printf(sbuf_c4go_postfix []sbuf, s []byte, c4goArgs ...interface{}) {
	var buf []byte = make([]byte, 1024)
	var ap *va_list
	va_start(ap, s)
	noarch.Vsnprintf(buf, int32(1024), s, ap)
	va_end(ap)
	sbuf_append(sbuf_c4go_postfix, buf)
}

// sbuf_empty - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:49
func sbuf_empty(sbuf_c4go_postfix []sbuf) int32 {
	return noarch.BoolToInt(noarch.Not(sbuf_c4go_postfix[0].n))
}

// sbuf_buf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:54
func sbuf_buf(sbuf_c4go_postfix []sbuf) []byte {
	sbuf_c4go_postfix[0].s[sbuf_c4go_postfix[0].n] = '\x00'
	return sbuf_c4go_postfix[0].s
}

// sbuf_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:60
func sbuf_len(sbuf_c4go_postfix []sbuf) int32 {
	return sbuf_c4go_postfix[0].n
}

// sbuf_cut - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:66
func sbuf_cut(sbuf_c4go_postfix []sbuf, n int32) {
	if sbuf_c4go_postfix[0].n > n {
		// shorten the sbuf
		sbuf_c4go_postfix[0].n = n
	}
}

// sbuf_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:72
func sbuf_done(sbuf_c4go_postfix []sbuf) {
	_ = sbuf_c4go_postfix[0].s
}

// sbuf_out - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:77
func sbuf_out(sbuf_c4go_postfix []sbuf) []byte {
	var s []byte = sbuf_c4go_postfix[0].s
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&sbuf_c4go_postfix[0]))) / int64(1))))[:], byte(0), 24)
	return s
}

// tr_nl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:8
// built-in troff requests
// just read a newline
var tr_nl int32 = 1

// tr_bm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:9
// blank line macro
var tr_bm int32 = -1

// tr_sm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:10
// leading space macro
var tr_sm int32 = -1

// c_pc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:11
// page number character
var c_pc []byte = []byte("%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// c_ec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:12
// escape character
var c_ec int32 = int32('\\')

// c_cc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:13
// control character
var c_cc int32 = int32('.')

// c_c2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:14
// no-break control character
var c_c2 int32 = int32('\'')

// jmp_eol - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:17
func jmp_eol() {
	// skip everything until the end of line
	var c int32
	for {
		c = cp_next()
		if !(c >= 0 && c != int32('\n')) {
			break
		}
	}
}

// tr_vs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:25
func tr_vs(args [][]byte) {
	var vs int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('v')))[0], int32('p'))
		}
		return (nreg(map_([]byte(".v0\x00"))))[0]
	}()
	(nreg(map_([]byte(".v0\x00"))))[0] = (nreg(int32('v')))[0]
	if 0 < vs {
		(nreg(int32('v')))[0] = vs
	} else {
		(nreg(int32('v')))[0] = 0
	}
}

// tr_ls - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:32
func tr_ls(args [][]byte) {
	var ls int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('L')))[0], 0)
		}
		return (nreg(map_([]byte(".L0\x00"))))[0]
	}()
	(nreg(map_([]byte(".L0\x00"))))[0] = (nreg(int32('L')))[0]
	if 1 < ls {
		(nreg(int32('L')))[0] = ls
	} else {
		(nreg(int32('L')))[0] = 1
	}
}

// tr_pl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:39
func tr_pl(args [][]byte) {
	var n int32 = eval_re(func() []byte {
		if args[1] != nil {
			return args[1]
		}
		return []byte("11i\x00")
	}(), (nreg(int32('p')))[0], int32('v'))
	if 0 < n {
		(nreg(int32('p')))[0] = n
	} else {
		(nreg(int32('p')))[0] = 0
	}
}

// tr_nr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:45
func tr_nr(args [][]byte) {
	var id int32
	if args[2] == nil {
		return
	}
	id = map_(args[1])
	num_set(id, eval_re(args[2], (nreg(id))[0], int32('u')))
	if args[3] != nil {
		num_setinc(id, eval(args[3], int32('u')))
	}
}

// tr_rr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:56
func tr_rr(args [][]byte) {
	var i int32
	for i = 1; i < 32; i++ {
		if args[i] != nil {
			num_del(map_(args[i]))
		}
	}
}

// tr_af - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:64
func tr_af(args [][]byte) {
	if args[2] != nil {
		num_setfmt(map_(args[1]), args[2])
	}
}

// tr_ds - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:70
func tr_ds(args [][]byte) {
	str_set(map_(args[1]), func() []byte {
		if args[2] != nil {
			return args[2]
		}
		return []byte("\x00")
	}())
}

// tr_as - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:75
func tr_as(args [][]byte) {
	var reg int32
	var s1 []byte
	var s2 []byte
	var s []byte
	reg = map_(args[1])
	if str_get(reg) != nil {
		s1 = str_get(reg)
	} else {
		s1 = []byte("\x00")
	}
	if args[2] != nil {
		s2 = args[2]
	} else {
		s2 = []byte("\x00")
	}
	s = xmalloc(noarch.Strlen(s1) + noarch.Strlen(s2) + int32(1)).([]byte)
	noarch.Strcpy(s, s1)
	noarch.Strcat(s, s2)
	str_set(reg, s)
	_ = s
}

// tr_rm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:89
func tr_rm(args [][]byte) {
	var i int32
	for i = 1; i < 32; i++ {
		if args[i] != nil {
			str_rm(map_(args[i]))
		}
	}
}

// tr_rn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:97
func tr_rn(args [][]byte) {
	if args[2] == nil {
		return
	}
	str_rn(map_(args[1]), map_(args[2]))
}

// tr_po - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:104
func tr_po(args [][]byte) {
	var po int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(int32('o')))[0], int32('m'))
		}
		return (nreg(map_([]byte(".o0\x00"))))[0]
	}()
	(nreg(map_([]byte(".o0\x00"))))[0] = (nreg(int32('o')))[0]
	if 0 < po {
		(nreg(int32('o')))[0] = po
	} else {
		(nreg(int32('o')))[0] = 0
	}
}

// read_string - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:112
func read_string() []byte {
	// read a string argument of a macro
	var sbuf_c4go_postfix sbuf
	var c int32
	var empty int32
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	cp_copymode(1)
	for (func() int32 {
		c = cp_next()
		return c
	}()) == int32(' ') {
	}
	empty = noarch.BoolToInt(c <= 0 || c == int32('\n'))
	if c == int32('"') {
		c = cp_next()
	}
	for c > 0 && c != int32('\n') {
		if c != 4 {
			sbuf_add(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), c)
		}
		c = cp_next()
	}
	if c >= 0 {
		in_back(c)
	}
	cp_copymode(0)
	if empty != 0 {
		sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		return nil
	}
	return sbuf_out(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
}

// read_name - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:140
func read_name(two int32) []byte {
	// read a space separated macro argument; if two, read at most two characters
	var sbuf_c4go_postfix sbuf
	var c int32 = cp_next()
	var i int32
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	for c == int32(' ') || c == int32('\t') || c == 4 {
		c = cp_next()
	}
	for c > 0 && c != int32(' ') && c != int32('\t') && c != int32('\n') && (noarch.Not(two) || i < 2) {
		if c != 4 {
			sbuf_add(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), c)
			i++
		}
		c = cp_next()
	}
	if c >= 0 {
		in_back(c)
	}
	return sbuf_out(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
}

// macrobody - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:160
func macrobody(sbuf_c4go_postfix []sbuf, end []byte) {
	var first int32 = 1
	var c int32
	var req []byte
	in_back(int32('\n'))
	cp_copymode(1)
	for (func() int32 {
		c = cp_next()
		return c
	}()) >= 0 {
		if sbuf_c4go_postfix != nil && noarch.Not(first) {
			sbuf_add(sbuf_c4go_postfix, c)
		}
		first = 0
		if c == int32('\n') {
			if (func() int32 {
				c = cp_next()
				return c
			}()) != c_cc {
				in_back(c)
				continue
			}
			req = read_name((nreg(int32('C')))[0])
			if noarch.Not(noarch.Strcmp(end, req)) {
				in_push(end, nil)
				in_back(c_cc)
				break
			}
			if sbuf_c4go_postfix != nil {
				sbuf_add(sbuf_c4go_postfix, c_cc)
				sbuf_append(sbuf_c4go_postfix, req)
			}
			_ = req
			req = nil
		}
	}
	_ = req
	cp_copymode(0)
}

// tr_de - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:194
func tr_de(args [][]byte) {
	var sbuf_c4go_postfix sbuf
	var id int32
	if args[1] == nil {
		return
	}
	id = map_(args[1])
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	if int32(args[0][1]) == int32('a') && int32(args[0][2]) == int32('m') && str_get(id) != nil {
		sbuf_append(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), str_get(id))
	}
	macrobody(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), func() []byte {
		if args[2] != nil {
			return args[2]
		}
		return []byte(".\x00")
	}())
	str_set(id, sbuf_buf(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix)))
	sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	if noarch.Not((nreg(int32('C')))[0]) && args[3] != nil {
		// parse the arguments as request argv[3]
		str_dset(id, str_dget(map_(args[3])))
	}
}

// tr_ig - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:211
func tr_ig(args [][]byte) {
	macrobody(nil, func() []byte {
		if args[1] != nil {
			return args[1]
		}
		return []byte(".\x00")
	}())
}

// read_until - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:217
func read_until(sbuf_c4go_postfix []sbuf, stop []byte, next func() int32, back func(int32)) int32 {
	// read into sbuf until stop; if stop is NULL, stop at whitespace
	var cs []byte = make([]byte, 32)
	var cs2 []byte = make([]byte, 32)
	var c int32
	for (func() int32 {
		c = next()
		return c
	}()) >= 0 {
		if c == 4 {
			continue
		}
		back(c)
		if c == int32('\n') {
			return 1
		}
		if stop == nil && (c == int32(' ') || c == int32('\t')) {
			return 0
		}
		charnext(cs, next, back)
		if stop != nil && noarch.Not(noarch.Strcmp(stop, cs)) {
			return 0
		}
		charnext_str(cs2, cs)
		sbuf_append(sbuf_c4go_postfix, cs2)
	}
	return 1
}

// if_strcmp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:240
func if_strcmp(next func() int32, back func(int32)) int32 {
	// evaluate .if strcmp (i.e. 'str'str')
	var delim []byte = make([]byte, 32)
	var s1 sbuf
	var s2 sbuf
	var ret int32
	charnext(delim, next, back)
	sbuf_init(c4goUnsafeConvert_sbuf(&s1))
	sbuf_init(c4goUnsafeConvert_sbuf(&s2))
	read_until(c4goUnsafeConvert_sbuf(&s1), delim, next, back)
	read_until(c4goUnsafeConvert_sbuf(&s2), delim, next, back)
	cp_reqbeg()
	ret = noarch.BoolToInt(noarch.Not(noarch.Strcmp(sbuf_buf(c4goUnsafeConvert_sbuf(&s1)), sbuf_buf(c4goUnsafeConvert_sbuf(&s2)))))
	sbuf_done(c4goUnsafeConvert_sbuf(&s1))
	sbuf_done(c4goUnsafeConvert_sbuf(&s2))
	return ret
}

// if_cond - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:258
func if_cond(next func() int32, back func(int32)) int32 {
	switch cp_next() {
	case 'o':
		// evaluate .if condition letters
		return (nreg(map_([]byte("%\x00"))))[0] % 2
	case 'e':
		return noarch.BoolToInt(noarch.Not((nreg(map_([]byte("%\x00"))))[0] % 2))
	case 't':
		return 1
	case 'n':
		return 0
	}
	return 0
}

// if_eval - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:274
func if_eval(next func() int32, back func(int32)) int32 {
	// evaluate .if condition
	var sbuf_c4go_postfix sbuf
	var ret int32
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	read_until(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), nil, next, back)
	ret = noarch.BoolToInt(eval(sbuf_buf(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix)), int32('\x00')) > 0)
	sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	return ret
}

// eval_if - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:285
func eval_if(next func() int32, back func(int32)) int32 {
	var neg int32
	var ret int32
	var c int32
	for {
		c = next()
		if !(c == int32(' ') || c == int32('\t')) {
			break
		}
	}
	if c == int32('!') {
		neg = 1
		c = next()
	}
	back(c)
	if noarch.Strchr([]byte("oetn\x00"), c) != nil {
		ret = if_cond(next, back)
	} else if c == int32(' ') {
		ret = 0
	} else if noarch.Not(int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit))) && noarch.Strchr([]byte("-+*/%<=>&:.|()\x00"), c) == nil {
		ret = if_strcmp(next, back)
	} else {
		ret = if_eval(next, back)
	}
	return noarch.BoolToInt(ret != neg)
}

// ie_cond - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:310
// .ie condition stack
var ie_cond []int32 = make([]int32, 128)

// ie_depth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:311
var ie_depth int32

// tr_if - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:313
func tr_if(args [][]byte) {
	var c int32 = eval_if(cp_next, in_back)
	if int32(args[0][1]) == int32('i') && int32(args[0][2]) == int32('e') {
		if ie_depth < 128 {
			// .ie command
			ie_cond[func() int32 {
				defer func() {
					ie_depth++
				}()
				return ie_depth
			}()] = c
		}
	}
	cp_blk(noarch.BoolToInt(noarch.Not(c)))
}

// tr_el - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:322
func tr_el(args [][]byte) {
	cp_blk(func() int32 {
		if ie_depth > 0 {
			return ie_cond[func() int32 {
				ie_depth--
				return ie_depth
			}()]
		}
		return 1
	}())
}

// tr_na - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:327
func tr_na(args [][]byte) {
	(nreg(map_([]byte(".na\x00"))))[0] = 1
}

// adjmode - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:332
func adjmode(c int32, def int32) int32 {
	switch c {
	case 'l':
		return 1
	case 'r':
		return 2
	case 'c':
		return 0
	case 'b':
		fallthrough
	case 'n':
		return 3
	case 'k':
		return 3 | 8
	}
	return def
}

// tr_ad - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:350
func tr_ad(args [][]byte) {
	var s []byte = args[1]
	(nreg(map_([]byte(".na\x00"))))[0] = 0
	if s == nil {
		return
	}
	if int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		(nreg(int32('j')))[0] = noarch.Atoi(s) & 15
	} else {
		if int32(s[0]) == int32('p') {
			(nreg(int32('j')))[0] = 4 | adjmode(int32(s[1]), 3)
		} else {
			(nreg(int32('j')))[0] = adjmode(int32(s[0]), (nreg(int32('j')))[0])
		}
	}
}

// tr_tm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:362
func tr_tm(args [][]byte) {
	noarch.Fprintf(noarch.Stderr, []byte("%s\n\x00"), func() []byte {
		if args[1] != nil {
			return args[1]
		}
		return []byte("\x00")
	}())
}

// tr_so - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:367
func tr_so(args [][]byte) {
	if args[1] != nil {
		in_so(args[1])
	}
}

// tr_nx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:373
func tr_nx(args [][]byte) {
	in_nx(args[1])
}

// tr_shift - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:378
func tr_shift(args [][]byte) {
	var n int32 = func() int32 {
		if args[1] != nil {
			return noarch.Atoi(args[1])
		}
		return 1
	}()
	for func() int32 {
		defer func() {
			n--
		}()
		return n
	}() >= 1 {
		in_shift()
	}
}

// tr_ex - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:385
func tr_ex(args [][]byte) {
	in_ex()
}

// tr_sy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:390
func tr_sy(args [][]byte) {
	noarch.System(args[1])
}

// tr_lt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:395
func tr_lt(args [][]byte) {
	var lt int32 = func() int32 {
		if args[1] != nil {
			return eval_re(args[1], (nreg(map_([]byte(".lt\x00"))))[0], int32('m'))
		}
		return (nreg(map_([]byte(".lt0\x00"))))[0]
	}()
	(nreg(map_([]byte(".lt0\x00"))))[0] = (nreg(map_([]byte(".lt0\x00"))))[0]
	if 0 < lt {
		(nreg(map_([]byte(".lt\x00"))))[0] = lt
	} else {
		(nreg(map_([]byte(".lt\x00"))))[0] = 0
	}
}

// tr_pc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:402
func tr_pc(args [][]byte) {
	var s []byte = args[1]
	if s == nil || charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], c_pc) < 0 {
		noarch.Strcpy(c_pc, []byte("\x00"))
	}
}

// tr_tl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:409
func tr_tl(args [][]byte) {
	var c int32
	for {
		c = cp_next()
		if !(c >= 0 && (c == int32(' ') || c == int32('\t'))) {
			break
		}
	}
	in_back(c)
	ren_tl(cp_next, in_back)
	for {
		c = cp_next()
		if !(c >= 0 && c != int32('\n')) {
			break
		}
	}
}

// tr_ec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:422
func tr_ec(args [][]byte) {
	if args[1] != nil {
		c_ec = int32(args[1][0])
	} else {
		c_ec = int32('\\')
	}
}

// tr_cc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:427
func tr_cc(args [][]byte) {
	if args[1] != nil {
		c_cc = int32(args[1][0])
	} else {
		c_cc = int32('.')
	}
}

// tr_c2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:432
func tr_c2(args [][]byte) {
	if args[1] != nil {
		c_c2 = int32(args[1][0])
	} else {
		c_c2 = int32('\'')
	}
}

// tr_eo - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:437
func tr_eo(args [][]byte) {
	c_ec = -1
}

// tr_hc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:442
func tr_hc(args [][]byte) {
	var s []byte = args[1]
	if s == nil || charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], env_hc()) < 0 {
		noarch.Strcpy(env_hc(), []byte("\\%\x00"))
	}
}

// eos_sent - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:450
// sentence ending and their transparent characters
var eos_sent [][]byte = [][]byte{[]byte(".\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("?\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("!\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// eos_sentcnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:451
var eos_sentcnt int32 = 3

// eos_tran - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:452
var eos_tran [][]byte = [][]byte{[]byte("'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte(")\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("]\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("*\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// eos_trancnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:453
var eos_trancnt int32 = 5

// tr_eos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:455
func tr_eos(args [][]byte) {
	eos_sentcnt = 0
	eos_trancnt = 0
	if args[1] != nil {
		var s []byte = args[1]
		for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], eos_sent[eos_sentcnt]) >= 0 {
			if eos_sentcnt < 32-1 {
				eos_sentcnt++
			}
		}
	}
	if args[2] != nil {
		var s []byte = args[2]
		for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], eos_tran[eos_trancnt]) >= 0 {
			if eos_trancnt < 32-1 {
				eos_trancnt++
			}
		}
	}
}

// c_eossent - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:473
func c_eossent(s []byte) int32 {
	var i int32
	for i = 0; i < eos_sentcnt; i++ {
		if noarch.Not(noarch.Strcmp(eos_sent[i], s)) {
			return 1
		}
	}
	return 0
}

// c_eostran - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:482
func c_eostran(s []byte) int32 {
	var i int32
	for i = 0; i < eos_trancnt; i++ {
		if noarch.Not(noarch.Strcmp(eos_tran[i], s)) {
			return 1
		}
	}
	return 0
}

// hy_dash - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:492
// hyphenation dashes and hyphenation inhibiting character
var hy_dash [][]byte = [][]byte{[]byte("\\:\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("-\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("em\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("en\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\-\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("--\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("hy\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// hy_dashcnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:493
var hy_dashcnt int32 = 7

// hy_stop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:494
var hy_stop [][]byte = [][]byte{[]byte("\\%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// hy_stopcnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:495
var hy_stopcnt int32 = 1

// tr_nh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:497
func tr_nh(args [][]byte) {
	(nreg(map_([]byte(".hy\x00"))))[0] = 0
}

// tr_hy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:502
func tr_hy(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".hy\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".hy\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".hy\x00"))))[0] = 1
	}
}

// tr_hlm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:507
func tr_hlm(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".hlm\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".hlm\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".hlm\x00"))))[0] = 0
	}
}

// tr_hycost - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:512
func tr_hycost(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".hycost\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".hycost\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".hycost\x00"))))[0] = 0
	}
	if args[2] != nil {
		(nreg(map_([]byte(".hycost2\x00"))))[0] = eval_re(args[2], (nreg(map_([]byte(".hycost2\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".hycost2\x00"))))[0] = 0
	}
	if args[3] != nil {
		(nreg(map_([]byte(".hycost3\x00"))))[0] = eval_re(args[3], (nreg(map_([]byte(".hycost3\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".hycost3\x00"))))[0] = 0
	}
}

// tr_hydash - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:519
func tr_hydash(args [][]byte) {
	hy_dashcnt = 0
	if args[1] != nil {
		var s []byte = args[1]
		for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], hy_dash[hy_dashcnt]) >= 0 {
			if hy_dashcnt < 32-1 {
				hy_dashcnt++
			}
		}
	}
}

// tr_hystop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:530
func tr_hystop(args [][]byte) {
	hy_stopcnt = 0
	if args[1] != nil {
		var s []byte = args[1]
		for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], hy_stop[hy_stopcnt]) >= 0 {
			if hy_stopcnt < 32-1 {
				hy_stopcnt++
			}
		}
	}
}

// c_hydash - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:541
func c_hydash(s []byte) int32 {
	var i int32
	for i = 0; i < hy_dashcnt; i++ {
		if noarch.Not(noarch.Strcmp(hy_dash[i], s)) {
			return 1
		}
	}
	return 0
}

// c_hystop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:550
func c_hystop(s []byte) int32 {
	var i int32
	for i = 0; i < hy_stopcnt; i++ {
		if noarch.Not(noarch.Strcmp(hy_stop[i], s)) {
			return 1
		}
	}
	return 0
}

// c_hymark - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:559
func c_hymark(s []byte) int32 {
	return noarch.BoolToInt(noarch.Not(noarch.Strcmp([]byte("\\:\x00"), s)) || noarch.Not(noarch.Strcmp(env_hc(), s)))
}

// tr_pmll - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:564
func tr_pmll(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".pmll\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".pmll\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".pmll\x00"))))[0] = 0
	}
	if args[2] != nil {
		(nreg(map_([]byte(".pmllcost\x00"))))[0] = eval_re(args[2], (nreg(map_([]byte(".pmllcost\x00"))))[0], int32('\x00'))
	} else {
		(nreg(map_([]byte(".pmllcost\x00"))))[0] = 100
	}
}

// tr_lg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:570
func tr_lg(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".lg\x00"))))[0] = eval(args[1], int32('\x00'))
	}
}

// tr_kn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:576
func tr_kn(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".kn\x00"))))[0] = eval(args[1], int32('\x00'))
	}
}

// tr_cp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:582
func tr_cp(args [][]byte) {
	if args[1] != nil {
		(nreg(int32('C')))[0] = noarch.Atoi(args[1])
	}
}

// tr_ss - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:588
func tr_ss(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".ss\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".ss\x00"))))[0], 0)
		if args[2] != nil {
			(nreg(map_([]byte(".sss\x00"))))[0] = eval_re(args[2], (nreg(map_([]byte(".sss\x00"))))[0], 0)
		} else {
			(nreg(map_([]byte(".sss\x00"))))[0] = (nreg(map_([]byte(".ss\x00"))))[0]
		}
	}
}

// tr_ssh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:596
func tr_ssh(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".ssh\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte(".ssh\x00"))))[0], 0)
	} else {
		(nreg(map_([]byte(".ssh\x00"))))[0] = 0
	}
}

// tr_cs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:601
func tr_cs(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if fn != nil {
		font_setcs(fn, func() int32 {
			if args[2] != nil {
				return eval(args[2], 0)
			}
			return 0
		}(), func() int32 {
			if args[3] != nil {
				return eval(args[3], 0)
			}
			return 0
		}())
	}
}

// tr_fzoom - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:609
func tr_fzoom(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if fn != nil {
		font_setzoom(fn, func() int32 {
			if args[2] != nil {
				return eval(args[2], 0)
			}
			return 0
		}())
	}
}

// tr_tkf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:616
func tr_tkf(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if len(fn) == 0 && len(args[5]) == 0 {
		font_track(fn, eval(args[2], 0), eval(args[3], 0), eval(args[4], 0), eval(args[5], 0))
	}
}

// tr_ff - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:624
func tr_ff(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	var i int32
	for i = 2; i < 32; i++ {
		if len(fn) == 0 && len(args[i]) == 0 && int32(args[i][0]) != 0 && int32(args[i][1]) != 0 {
			font_feat(fn, (args[i])[0+1:], noarch.BoolToInt(int32(args[i][0]) == int32('+')))
		}
	}
}

// tr_ffsc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:633
func tr_ffsc(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if fn != nil {
		font_scrp(fn, args[2])
	}
	if fn != nil {
		font_lang(fn, args[3])
	}
}

// tr_nm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:642
func tr_nm(args [][]byte) {
	if args[1] == nil {
		(nreg(map_([]byte(".nm\x00"))))[0] = 0
		return
	}
	(nreg(map_([]byte(".nm\x00"))))[0] = 1
	(nreg(map_([]byte("ln\x00"))))[0] = eval_re(args[1], (nreg(map_([]byte("ln\x00"))))[0], 0)
	if 0 < (nreg(map_([]byte("ln\x00"))))[0] {
		(nreg(map_([]byte("ln\x00"))))[0] = (nreg(map_([]byte("ln\x00"))))[0]
	} else {
		(nreg(map_([]byte("ln\x00"))))[0] = 0
	}
	if args[2] != nil && int32(((__ctype_b_loc())[0])[int32(uint8(args[2][0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		if 1 < eval(args[2], 0) {
			(nreg(map_([]byte(".nM\x00"))))[0] = eval(args[2], 0)
		} else {
			(nreg(map_([]byte(".nM\x00"))))[0] = 1
		}
	}
	if args[3] != nil && int32(((__ctype_b_loc())[0])[int32(uint8(args[3][0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		if 0 < eval(args[3], 0) {
			(nreg(map_([]byte(".nS\x00"))))[0] = eval(args[3], 0)
		} else {
			(nreg(map_([]byte(".nS\x00"))))[0] = 0
		}
	}
	if args[4] != nil && int32(((__ctype_b_loc())[0])[int32(uint8(args[4][0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		if 0 < eval(args[4], 0) {
			(nreg(map_([]byte(".nI\x00"))))[0] = eval(args[4], 0)
		} else {
			(nreg(map_([]byte(".nI\x00"))))[0] = 0
		}
	}
}

// tr_nn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:659
func tr_nn(args [][]byte) {
	if args[1] != nil {
		(nreg(map_([]byte(".nn\x00"))))[0] = eval(args[1], 0)
	} else {
		(nreg(map_([]byte(".nn\x00"))))[0] = 1
	}
}

// tr_bd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:664
func tr_bd(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if args[1] == nil || noarch.Not(noarch.Strcmp([]byte("S\x00"), args[1])) {
		return
	}
	if fn != nil {
		font_setbd(fn, func() int32 {
			if args[2] != nil {
				return eval(args[2], int32('u'))
			}
			return 0
		}())
	}
}

// tr_it - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:673
func tr_it(args [][]byte) {
	if args[2] != nil {
		(nreg(map_([]byte(".it\x00"))))[0] = map_(args[2])
		(nreg(map_([]byte(".itn\x00"))))[0] = eval(args[1], 0)
	} else {
		(nreg(map_([]byte(".it\x00"))))[0] = 0
	}
}

// tr_mc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:683
func tr_mc(args [][]byte) {
	var s []byte = args[1]
	if s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], env_mc()) >= 0 {
		(nreg(map_([]byte(".mc\x00"))))[0] = 1
		if args[2] != nil {
			(nreg(map_([]byte(".mcn\x00"))))[0] = eval(args[2], int32('m'))
		} else {
			(nreg(map_([]byte(".mcn\x00"))))[0] = (nreg(int32('s')))[0] * dev_res / 72
		}
	} else {
		(nreg(map_([]byte(".mc\x00"))))[0] = 0
	}
}

// tr_tc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:694
func tr_tc(args [][]byte) {
	var s []byte = args[1]
	if s == nil || charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], env_tc()) < 0 {
		noarch.Strcpy(env_tc(), []byte("\x00"))
	}
}

// tr_lc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:701
func tr_lc(args [][]byte) {
	var s []byte = args[1]
	if s == nil || charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], env_lc()) < 0 {
		noarch.Strcpy(env_lc(), []byte("\x00"))
	}
}

// tr_lf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:708
func tr_lf(args [][]byte) {
	if args[1] != nil {
		in_lf(args[2], eval(args[1], 0))
	}
}

// tr_chop - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:714
func tr_chop(args [][]byte) {
	var sbuf_c4go_postfix sbuf
	var id int32
	id = map_(args[1])
	if str_get(id) != nil {
		sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		sbuf_append(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), str_get(id))
		if noarch.Not(sbuf_empty(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))) {
			sbuf_cut(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), sbuf_len(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))-1)
			str_set(id, sbuf_buf(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix)))
		}
		sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	}
}

// cmap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:731
// character translation (.tr)
// character mapping
var cmap []dict

// cmap_src - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:732
// source character
var cmap_src [][]byte = make([][]byte, 512)

// cmap_dst - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:733
// character mapping
var cmap_dst [][]byte = make([][]byte, 512)

// cmap_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:734
// number of translated character
var cmap_n int32

// cmap_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:736
func cmap_add(c1 []byte, c2 []byte) {
	var i int32 = dict_get(cmap, c1)
	if i >= 0 {
		noarch.Strcpy(cmap_dst[i], c2)
	} else if cmap_n < 512 {
		noarch.Strcpy(cmap_src[cmap_n], c1)
		noarch.Strcpy(cmap_dst[cmap_n], c2)
		dict_put(cmap, cmap_src[cmap_n], cmap_n)
		cmap_n++
	}
}

// cmap_map - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:749
func cmap_map(c []byte) []byte {
	var i int32 = dict_get(cmap, c)
	if i >= 0 {
		return cmap_dst[i]
	}
	return c
}

// tr_tr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:755
func tr_tr(args [][]byte) {
	var s []byte = args[1]
	var c1 []byte = make([]byte, 32)
	var c2 []byte = make([]byte, 32)
	for s != nil && charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], c1) >= 0 {
		if charread((*[1000000][]byte)(unsafe.Pointer(&s))[:], c2) < 0 {
			noarch.Strcpy(c2, []byte(" \x00"))
		}
		cmap_add(c1, c2)
	}
}

// cdef_src - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:767
// character definition (.char)
// source character
var cdef_src [][]byte = make([][]byte, 128)

// cdef_dst - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:768
// character definition
var cdef_dst [][]byte = make([][]byte, 128)

// cdef_fn - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:769
// owning font
var cdef_fn []int32 = make([]int32, 128)

// cdef_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:770
// number of defined characters
var cdef_n int32

// cdef_expanding - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:771
// inside cdef_expand() call
var cdef_expanding int32

// cdef_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:773
func cdef_find(c []byte, fn int32) int32 {
	var i int32
	for i = 0; i < cdef_n; i++ {
		if (noarch.Not(cdef_fn[i]) || cdef_fn[i] == fn) && noarch.Not(noarch.Strcmp(cdef_src[i], c)) {
			return i
		}
	}
	return -1
}

// cdef_map - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:783
func cdef_map(c []byte, fn int32) []byte {
	// return the definition of the given character
	var i int32 = cdef_find(c, fn)
	if noarch.Not(cdef_expanding) && i >= 0 {
		return cdef_dst[i]
	}
	return nil
}

// cdef_expand - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:789
func cdef_expand(wb_c4go_postfix []wb, s []byte, fn int32) int32 {
	var d []byte = cdef_map(s, fn)
	if d == nil {
		return 1
	}
	cdef_expanding = 1
	ren_parse(wb_c4go_postfix, d)
	cdef_expanding = 0
	return 0
}

// cdef_remove - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:800
func cdef_remove(fn []byte, cs []byte) {
	var c []byte = make([]byte, 32)
	var i int32
	var fp int32 = func() int32 {
		if fn != nil {
			return dev_pos(fn)
		}
		return -1
	}()
	if cs == nil || charread((*[1000000][]byte)(unsafe.Pointer(&cs))[:], c) < 0 {
		return
	}
	for i = 0; i < cdef_n; i++ {
		if noarch.Not(noarch.Strcmp(cdef_src[i], c)) {
			if fn == nil || fp > 0 && cdef_fn[i] == fp {
				_ = cdef_dst[i]
				cdef_dst[i] = nil
				cdef_src[i][0] = '\x00'
			}
		}
	}
}

// cdef_add - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:818
func cdef_add(fn []byte, cs []byte, def []byte) {
	var c []byte = make([]byte, 32)
	var i int32
	if def == nil || charread((*[1000000][]byte)(unsafe.Pointer(&cs))[:], c) < 0 {
		return
	}
	i = cdef_find(c, func() int32 {
		if fn != nil {
			return dev_pos(fn)
		}
		return -1
	}())
	if i < 0 {
		for i = 0; i < cdef_n; i++ {
			if cdef_dst[i] == nil {
				break
			}
		}
		if i == cdef_n && cdef_n < 128 {
			cdef_n++
		}
	}
	if i >= 0 && i < cdef_n {
		noarch.Snprintf(cdef_src[i], int32(32), []byte("%s\x00"), c)
		cdef_dst[i] = xmalloc(noarch.Strlen(def) + int32(1)).([]byte)
		noarch.Strcpy(cdef_dst[i], def)
		if fn != nil {
			cdef_fn[i] = dev_pos(fn)
		} else {
			cdef_fn[i] = 0
		}
	}
}

// tr_rchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:840
func tr_rchar(args [][]byte) {
	var i int32
	for i = 1; i < 32; i++ {
		if args[i] != nil {
			cdef_remove(nil, args[i])
		}
	}
}

// tr_char - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:848
func tr_char(args [][]byte) {
	if args[2] != nil {
		cdef_add(nil, args[1], args[2])
	} else {
		cdef_remove(nil, args[1])
	}
}

// tr_ochar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:856
func tr_ochar(args [][]byte) {
	if args[3] != nil {
		cdef_add(args[1], args[2], args[3])
	} else {
		cdef_remove(args[1], args[2])
	}
}

// tr_fmap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:864
func tr_fmap(args [][]byte) {
	var fn []font = func() []font {
		if args[1] != nil {
			return dev_font(dev_pos(args[1]))
		}
		return nil
	}()
	if len(fn) == 0 && len(args[2]) == 0 {
		font_map(fn, args[2], args[3])
	}
}

// tr_blm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:871
func tr_blm(args [][]byte) {
	if args[1] != nil {
		tr_bm = map_(args[1])
	} else {
		tr_bm = -1
	}
}

// tr_lsm - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:876
func tr_lsm(args [][]byte) {
	if args[1] != nil {
		tr_sm = map_(args[1])
	} else {
		tr_sm = -1
	}
}

// tr_co - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:881
func tr_co(args [][]byte) {
	var src []byte = args[1]
	var dst []byte = args[2]
	if len(src) == 0 && len(dst) == 0 && str_get(map_(src)) != nil {
		str_set(map_(dst), str_get(map_(src)))
	}
}

// tr_coa - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:889
func tr_coa(args [][]byte) {
	var src []byte = args[1]
	var dst []byte = args[2]
	if len(src) == 0 && len(dst) == 0 && str_get(map_(src)) != nil {
		var sb sbuf
		sbuf_init(c4goUnsafeConvert_sbuf(&sb))
		if str_get(map_(dst)) != nil {
			sbuf_append(c4goUnsafeConvert_sbuf(&sb), str_get(map_(dst)))
		}
		sbuf_append(c4goUnsafeConvert_sbuf(&sb), str_get(map_(src)))
		str_set(map_(dst), sbuf_buf(c4goUnsafeConvert_sbuf(&sb)))
		sbuf_done(c4goUnsafeConvert_sbuf(&sb))
	}
}

// tr_coo - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:904
func tr_coo(args [][]byte) {
	var reg []byte = args[1]
	var path []byte = args[2]
	var fp *noarch.File
	if reg == nil || noarch.Not(reg[0]) || path == nil || noarch.Not(path[0]) {
		return
	}
	if (func() *noarch.File {
		fp = noarch.Fopen(path, []byte("w\x00"))
		return fp
	}()) != nil {
		if str_get(map_(reg)) != nil {
			noarch.Fputs(str_get(map_(reg)), fp)
		}
		noarch.Fclose(fp)
	}
}

// tr_coi - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:918
func tr_coi(args [][]byte) {
	var reg []byte = args[1]
	var path []byte = args[2]
	var buf []byte = make([]byte, 1024)
	var fp *noarch.File
	if reg == nil || noarch.Not(reg[0]) || path == nil || noarch.Not(path[0]) {
		return
	}
	if (func() *noarch.File {
		fp = noarch.Fopen(path[0+1:], []byte("r\x00"))
		return fp
	}()) != nil {
		var sb sbuf
		sbuf_init(c4goUnsafeConvert_sbuf(&sb))
		for noarch.Fgets(buf, int32(1024), fp) != nil {
			sbuf_append(c4goUnsafeConvert_sbuf(&sb), buf)
		}
		str_set(map_(reg), sbuf_buf(c4goUnsafeConvert_sbuf(&sb)))
		sbuf_done(c4goUnsafeConvert_sbuf(&sb))
		noarch.Fclose(fp)
	}
}

// tr_dv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:937
func tr_dv(args [][]byte) {
	if args[1] != nil {
		out_x(args[1])
	}
}

// macroarg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:944
func macroarg(sbuf_c4go_postfix []sbuf, brk int32, next func() int32, back func(int32)) int32 {
	// read a single macro argument
	var quoted int32
	var c int32
	c = next()
	for c == int32(' ') {
		c = next()
	}
	if c == int32('\n') || c == brk {
		back(c)
	}
	if c < 0 || c == int32('\n') || c == brk {
		return 1
	}
	if c == int32('"') {
		quoted = 1
		c = next()
	}
	for c >= 0 && c != int32('\n') && (quoted != 0 || c != brk) {
		if noarch.Not(quoted) && c == int32(' ') {
			break
		}
		if quoted != 0 && c == int32('"') {
			c = next()
			if c != int32('"') {
				break
			}
		}
		if c == c_ec {
			sbuf_add(sbuf_c4go_postfix, c)
			c = next()
		}
		sbuf_add(sbuf_c4go_postfix, c)
		c = next()
	}
	sbuf_add(sbuf_c4go_postfix, 0)
	if c >= 0 {
		back(c)
	}
	return 0
}

// chopargs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:981
func chopargs(sbuf_c4go_postfix []sbuf, args [][]byte) {
	// split the arguments in sbuf, after calling one of mkargs_*()
	var s []byte = sbuf_buf(sbuf_c4go_postfix)
	var e []byte = s[0+sbuf_len(sbuf_c4go_postfix):]
	var n int32
	for n < 32 && s != nil && (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&e[0])))/int64(1)) < 0 {
		args[func() int32 {
			defer func() {
				n++
			}()
			return n
		}()] = s
		if (func() []byte {
			s = memchr(s, int32('\x00'), uint32(int32((int64(uintptr(unsafe.Pointer(&e[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1))))).([]byte)
			return s
		}()) != nil {
			s = s[0+1:]
		}
	}
}

// tr_args - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:994
func tr_args(args [][]byte, brk int32, next func() int32, back func(int32)) []byte {
	// read macro arguments; free the returned pointer when done
	var sbuf_c4go_postfix sbuf
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	for noarch.Not(macroarg(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), brk, next, back)) {
	}
	chopargs(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), args)
	return sbuf_out(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
}

// mkargs_macro - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1005
func mkargs_macro(sbuf_c4go_postfix []sbuf) {
	// read regular macro arguments
	cp_copymode(1)
	for noarch.Not(macroarg(sbuf_c4go_postfix, -1, cp_next, in_back)) {
	}
	jmp_eol()
	cp_copymode(0)
}

// mkargs_req - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1015
func mkargs_req(sbuf_c4go_postfix []sbuf) {
	// read request arguments; trims tabs too
	var n int32
	var c int32
	c = cp_next()
	for n < 32 {
		var ok int32
		for c == int32(' ') || c == int32('\t') {
			c = cp_next()
		}
		for c >= 0 && c != int32('\n') && c != int32(' ') && c != int32('\t') {
			if c != 4 {
				sbuf_add(sbuf_c4go_postfix, c)
			}
			c = cp_next()
			ok = 1
		}
		if ok != 0 {
			n++
			sbuf_add(sbuf_c4go_postfix, 0)
		}
		if c == int32('\n') {
			in_back(c)
		}
		if c < 0 || c == int32('\n') {
			break
		}
	}
	jmp_eol()
}

// mkargs_ds - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1043
func mkargs_ds(sbuf_c4go_postfix []sbuf) {
	// read arguments for .ds and .char
	var s []byte = read_name((nreg(int32('C')))[0])
	sbuf_append(sbuf_c4go_postfix, s)
	sbuf_add(sbuf_c4go_postfix, 0)
	_ = s
	s = read_string()
	if s != nil {
		sbuf_append(sbuf_c4go_postfix, s)
		sbuf_add(sbuf_c4go_postfix, 0)
		_ = s
	}
	jmp_eol()
}

// mkargs_ochar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1059
func mkargs_ochar(sbuf_c4go_postfix []sbuf) {
	// read arguments for .ochar
	var s []byte
	sbuf_append(sbuf_c4go_postfix, s)
	sbuf_add(sbuf_c4go_postfix, 0)
	_ = s
	mkargs_ds(sbuf_c4go_postfix)
}

// mkargs_reg1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1069
func mkargs_reg1(sbuf_c4go_postfix []sbuf) {
	// read arguments for .nr
	var s []byte = read_name((nreg(int32('C')))[0])
	sbuf_append(sbuf_c4go_postfix, s)
	sbuf_add(sbuf_c4go_postfix, 0)
	_ = s
	mkargs_req(sbuf_c4go_postfix)
}

// mkargs_null - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1079
func mkargs_null(sbuf_c4go_postfix []sbuf) {
	// do not read any arguments; for .if, .ie and .el
	{
	}
}

// mkargs_eol - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1084
func mkargs_eol(sbuf_c4go_postfix []sbuf) {
	// read the whole line for .tm
	var c int32
	cp_copymode(1)
	c = cp_next()
	for c == int32(' ') {
		c = cp_next()
	}
	for c >= 0 && c != int32('\n') {
		if c != 4 {
			sbuf_add(sbuf_c4go_postfix, c)
		}
		c = cp_next()
	}
	cp_copymode(0)
}

// cmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1099
type cmd struct {
	id   []byte
	f    func([][]byte)
	args func([]sbuf)
}

// cmds - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1099
var cmds []cmd = []cmd{{[]byte("\a<\x00"), tr_divbeg, nil}, {[]byte("\a>\x00"), tr_divend, nil}, {[]byte("\aV\x00"), tr_divvs, nil}, {[]byte("\aP\x00"), tr_popren, nil}, {[]byte(">>\x00"), tr_l2r, nil}, {[]byte("<<\x00"), tr_r2l, nil}, {[]byte("ab\x00"), tr_ab, mkargs_eol}, {[]byte("ad\x00"), tr_ad, nil}, {[]byte("af\x00"), tr_af, nil}, {[]byte("am\x00"), tr_de, mkargs_reg1}, {[]byte("as\x00"), tr_as, mkargs_ds}, {[]byte("bd\x00"), tr_bd, nil}, {[]byte("blm\x00"), tr_blm, nil}, {[]byte("bp\x00"), tr_bp, nil}, {[]byte("br\x00"), tr_br, nil}, {[]byte("c2\x00"), tr_c2, nil}, {[]byte("cc\x00"), tr_cc, nil}, {[]byte("ce\x00"), tr_ce, nil}, {[]byte("ch\x00"), tr_ch, nil}, {[]byte("char\x00"), tr_char, mkargs_ds}, {[]byte("chop\x00"), tr_chop, mkargs_reg1}, {[]byte("cl\x00"), tr_cl, nil}, {[]byte("co\x00"), tr_co, nil}, {[]byte("co+\x00"), tr_coa, nil}, {[]byte("co<\x00"), tr_coi, mkargs_ds}, {[]byte("co>\x00"), tr_coo, mkargs_ds}, {[]byte("cp\x00"), tr_cp, nil}, {[]byte("cs\x00"), tr_cs, nil}, {[]byte("da\x00"), tr_di, nil}, {[]byte("de\x00"), tr_de, mkargs_reg1}, {[]byte("di\x00"), tr_di, nil}, {[]byte("ds\x00"), tr_ds, mkargs_ds}, {[]byte("dt\x00"), tr_dt, nil}, {[]byte("dv\x00"), tr_dv, mkargs_eol}, {[]byte("ec\x00"), tr_ec, nil}, {[]byte("el\x00"), tr_el, mkargs_null}, {[]byte("em\x00"), tr_em, nil}, {[]byte("eo\x00"), tr_eo, nil}, {[]byte("eos\x00"), tr_eos, nil}, {[]byte("ev\x00"), tr_ev, nil}, {[]byte("ex\x00"), tr_ex, nil}, {[]byte("fc\x00"), tr_fc, nil}, {[]byte("ff\x00"), tr_ff, nil}, {[]byte("fi\x00"), tr_fi, nil}, {[]byte("fl\x00"), tr_br, nil}, {[]byte("fmap\x00"), tr_fmap, nil}, {[]byte("fp\x00"), tr_fp, nil}, {[]byte("ffsc\x00"), tr_ffsc, nil}, {[]byte("fspecial\x00"), tr_fspecial, nil}, {[]byte("ft\x00"), tr_ft, nil}, {[]byte("fzoom\x00"), tr_fzoom, nil}, {[]byte("hc\x00"), tr_hc, nil}, {[]byte("hcode\x00"), tr_hcode, nil}, {[]byte("hlm\x00"), tr_hlm, nil}, {[]byte("hpf\x00"), tr_hpf, nil}, {[]byte("hpfa\x00"), tr_hpfa, nil}, {[]byte("hy\x00"), tr_hy, nil}, {[]byte("hycost\x00"), tr_hycost, nil}, {[]byte("hydash\x00"), tr_hydash, nil}, {[]byte("hystop\x00"), tr_hystop, nil}, {[]byte("hw\x00"), tr_hw, nil}, {[]byte("ie\x00"), tr_if, mkargs_null}, {[]byte("if\x00"), tr_if, mkargs_null}, {[]byte("ig\x00"), tr_ig, nil}, {[]byte("in\x00"), tr_in, nil}, {[]byte("in2\x00"), tr_in2, nil}, {[]byte("it\x00"), tr_it, nil}, {[]byte("kn\x00"), tr_kn, nil}, {[]byte("lc\x00"), tr_lc, nil}, {[]byte("lf\x00"), tr_lf, nil}, {[]byte("lg\x00"), tr_lg, nil}, {[]byte("ll\x00"), tr_ll, nil}, {[]byte("ls\x00"), tr_ls, nil}, {[]byte("lsm\x00"), tr_lsm, nil}, {[]byte("lt\x00"), tr_lt, nil}, {[]byte("mc\x00"), tr_mc, nil}, {[]byte("mk\x00"), tr_mk, nil}, {[]byte("na\x00"), tr_na, nil}, {[]byte("ne\x00"), tr_ne, nil}, {[]byte("nf\x00"), tr_nf, nil}, {[]byte("nh\x00"), tr_nh, nil}, {[]byte("nm\x00"), tr_nm, nil}, {[]byte("nn\x00"), tr_nn, nil}, {[]byte("nr\x00"), tr_nr, mkargs_reg1}, {[]byte("ns\x00"), tr_ns, nil}, {[]byte("nx\x00"), tr_nx, nil}, {[]byte("ochar\x00"), tr_ochar, mkargs_ochar}, {[]byte("os\x00"), tr_os, nil}, {[]byte("pc\x00"), tr_pc, nil}, {[]byte("pl\x00"), tr_pl, nil}, {[]byte("pmll\x00"), tr_pmll, nil}, {[]byte("pn\x00"), tr_pn, nil}, {[]byte("po\x00"), tr_po, nil}, {[]byte("ps\x00"), tr_ps, nil}, {[]byte("rchar\x00"), tr_rchar, nil}, {[]byte("rm\x00"), tr_rm, nil}, {[]byte("rn\x00"), tr_rn, nil}, {[]byte("rr\x00"), tr_rr, nil}, {[]byte("rs\x00"), tr_rs, nil}, {[]byte("rt\x00"), tr_rt, nil}, {[]byte("shift\x00"), tr_shift, nil}, {[]byte("so\x00"), tr_so, nil}, {[]byte("sp\x00"), tr_sp, nil}, {[]byte("ss\x00"), tr_ss, nil}, {[]byte("ssh\x00"), tr_ssh, nil}, {[]byte("sv\x00"), tr_sv, nil}, {[]byte("sy\x00"), tr_sy, mkargs_eol}, {[]byte("ta\x00"), tr_ta, nil}, {[]byte("tc\x00"), tr_tc, nil}, {[]byte("ti\x00"), tr_ti, nil}, {[]byte("ti2\x00"), tr_ti2, nil}, {[]byte("tkf\x00"), tr_tkf, nil}, {[]byte("tl\x00"), tr_tl, mkargs_null}, {[]byte("tm\x00"), tr_tm, mkargs_eol}, {[]byte("tr\x00"), tr_tr, mkargs_eol}, {[]byte("vs\x00"), tr_vs, nil}, {[]byte("wh\x00"), tr_wh, nil}}

// dotted - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1223
func dotted(name []byte, dot int32) []byte {
	var out []byte = xmalloc(noarch.Strlen(name) + int32(2)).([]byte)
	out[0] = byte(dot)
	noarch.Strcpy(out[0+1:], name)
	return out
}

// tr_req - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1232
func tr_req(reg int32, args [][]byte) {
	// execute a built-in request
	var req []cmd = str_dget(reg).([]cmd)
	if req != nil {
		req[0].f(args)
	}
}

// tr_nextreq_exec - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1240
func tr_nextreq_exec(mac []byte, arg0 []byte, readargs int32) {
	// interpolate a macro for tr_nextreq()
	var args [][]byte = [][]byte{arg0, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	var req []cmd = str_dget(map_(mac)).([]cmd)
	var str []byte = str_get(map_(mac))
	var sbuf_c4go_postfix sbuf
	sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
	if readargs != 0 {
		if len(req) == 0 && req[0].args == nil {
			req[0].args(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		}
		if req != nil && req[0].args == nil {
			mkargs_req(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		}
		if req == nil {
			mkargs_macro(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
		}
		chopargs(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), args[0+1:])
	}
	if str != nil {
		in_push(str, args)
	}
	if str == nil && req != nil {
		req[0].f(args)
	}
	sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
}

// tr_nextreq - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1264
func tr_nextreq() int32 {
	// read the next troff request; return zero if a request was executed.
	var mac []byte
	var arg0 []byte
	var c int32
	if noarch.Not(tr_nl) {
		return 1
	}
	c = cp_next()
	if c == c_ec {
		// transparent line indicator
		var c2 int32 = cp_next()
		if c2 == int32('!') {
			var args [][]byte = [][]byte{[]byte("\\!\x00"), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
			var sbuf_c4go_postfix sbuf
			sbuf_init(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
			cp_copymode(1)
			mkargs_eol(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
			cp_copymode(0)
			chopargs(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix), args[0+1:])
			tr_transparent(args)
			sbuf_done(c4goUnsafeConvert_sbuf(&sbuf_c4go_postfix))
			return 0
		}
		in_back(c2)
	}
	if c < 0 || c != c_cc && c != c_c2 && (c != int32('\n') || tr_bm < 0) && (c != int32(' ') || tr_sm < 0) {
		// not a request, a blank line, or a line with leading spaces
		in_back(c)
		return 1
	}
	cp_reqbeg()
	if c == int32('\n') {
		// blank line macro
		mac = make([]byte, noarch.Strlen(map_name(tr_bm))+int32(1))
		noarch.Strcpy(mac, map_name(tr_bm))
		arg0 = dotted(mac, int32('.'))
		tr_nextreq_exec(mac, arg0, 0)
	} else if c == int32(' ') {
		// leading space macro
		var i int32
		mac = make([]byte, noarch.Strlen(map_name(tr_sm))+int32(1))
		noarch.Strcpy(mac, map_name(tr_sm))
		for i = 0; c == int32(' '); i++ {
			c = cp_next()
		}
		in_back(c)
		(nreg(map_([]byte("lsn\x00"))))[0] = i
		arg0 = dotted(mac, int32('.'))
		tr_nextreq_exec(mac, arg0, 0)
	} else {
		mac = read_name((nreg(int32('C')))[0])
		arg0 = dotted(mac, c)
		tr_nextreq_exec(mac, arg0, 1)
	}
	_ = arg0
	_ = mac
	return 0
}

// tr_next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1322
func tr_next() int32 {
	var c int32
	for noarch.Not(tr_nextreq()) {
	}
	c = cp_next()
	tr_nl = noarch.BoolToInt(c == int32('\n') || c < 0)
	return c
}

// tr_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1332
func tr_init() {
	var i int32
	for i = 0; uint32(i) < 4680/40; i++ {
		str_dset(map_(cmds[i].id), cmds[i:])
	}
	cmap = dict_make(-1, 0, 2)
}

// tr_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/tr.c:1340
func tr_done() {
	var i int32
	for i = 0; i < cdef_n; i++ {
		_ = cdef_dst[i]
	}
	dict_free(cmap)
}

// wb_init - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:22
func wb_init(wb_c4go_postfix []wb) {
	// word buffer
	// the current font, size and color
	// italic correction
	// the maximum and minimum values of bounding box coordinates
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&wb_c4go_postfix[0]))) / int64(1))))[:], byte(0), 8320)
	sbuf_init((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:])
	wb_c4go_postfix[0].sub_collect = 1
	wb_c4go_postfix[0].f = -1
	wb_c4go_postfix[0].s = -1
	wb_c4go_postfix[0].m = -1
	wb_c4go_postfix[0].cd = -1
	wb_c4go_postfix[0].r_f = -1
	wb_c4go_postfix[0].r_s = -1
	wb_c4go_postfix[0].r_m = -1
	wb_c4go_postfix[0].r_cd = -1
	wb_c4go_postfix[0].llx = 1 << uint64(29)
	wb_c4go_postfix[0].lly = 1 << uint64(29)
	wb_c4go_postfix[0].urx = -(1 << uint64(29))
	wb_c4go_postfix[0].ury = -(1 << uint64(29))
}

// wb_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:41
func wb_done(wb_c4go_postfix []wb) {
	sbuf_done((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:])
}

// wb_stsb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:47
func wb_stsb(wb_c4go_postfix []wb) {
	// update wb->st and wb->sb
	if wb_c4go_postfix[0].st < wb_c4go_postfix[0].v-wb_c4go_postfix[0].s*dev_res/72 {
		wb_c4go_postfix[0].st = wb_c4go_postfix[0].st
	} else {
		wb_c4go_postfix[0].st = wb_c4go_postfix[0].v - wb_c4go_postfix[0].s*dev_res/72
	}
	if wb_c4go_postfix[0].sb < wb_c4go_postfix[0].v {
		wb_c4go_postfix[0].sb = wb_c4go_postfix[0].v
	} else {
		wb_c4go_postfix[0].sb = wb_c4go_postfix[0].sb
	}
}

// wb_bbox - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:54
func wb_bbox(wb_c4go_postfix []wb, llx int32, lly int32, urx int32, ury int32) {
	// update bounding box
	if wb_c4go_postfix[0].llx < wb_c4go_postfix[0].h+llx {
		wb_c4go_postfix[0].llx = wb_c4go_postfix[0].llx
	} else {
		wb_c4go_postfix[0].llx = wb_c4go_postfix[0].h + llx
	}
	if wb_c4go_postfix[0].lly < -wb_c4go_postfix[0].v+lly {
		wb_c4go_postfix[0].lly = wb_c4go_postfix[0].lly
	} else {
		wb_c4go_postfix[0].lly = -wb_c4go_postfix[0].v + lly
	}
	if wb_c4go_postfix[0].urx < wb_c4go_postfix[0].h+urx {
		wb_c4go_postfix[0].urx = wb_c4go_postfix[0].h + urx
	} else {
		wb_c4go_postfix[0].urx = wb_c4go_postfix[0].urx
	}
	if wb_c4go_postfix[0].ury < -wb_c4go_postfix[0].v+ury {
		wb_c4go_postfix[0].ury = -wb_c4go_postfix[0].v + ury
	} else {
		wb_c4go_postfix[0].ury = wb_c4go_postfix[0].ury
	}
}

// wb_pendingfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:63
func wb_pendingfont(wb_c4go_postfix []wb) int32 {
	// pending font, size or color changes
	return noarch.BoolToInt(wb_c4go_postfix[0].f != func() int32 {
		if (wb_c4go_postfix)[0].r_f >= 0 {
			return (wb_c4go_postfix)[0].r_f
		}
		return (nreg(int32('f')))[0]
	}() || wb_c4go_postfix[0].s != func() int32 {
		if (wb_c4go_postfix)[0].r_s >= 0 {
			return (wb_c4go_postfix)[0].r_s
		}
		return (nreg(int32('s')))[0]
	}() || noarch.Not((nreg(int32('C')))[0]) && wb_c4go_postfix[0].m != func() int32 {
		if (wb_c4go_postfix)[0].r_m >= 0 {
			return (wb_c4go_postfix)[0].r_m
		}
		return (nreg(int32('m')))[0]
	}())
}

// wb_pendingdir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:70
func wb_pendingdir(wb_c4go_postfix []wb) int32 {
	// pending direction change
	return noarch.BoolToInt(wb_c4go_postfix[0].cd != func() int32 {
		if (wb_c4go_postfix)[0].r_cd >= 0 {
			return (wb_c4go_postfix)[0].r_cd
		}
		return (nreg(map_([]byte(".cd\x00"))))[0]
	}())
}

// wb_flushfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:76
func wb_flushfont(wb_c4go_postfix []wb) {
	if wb_c4go_postfix[0].f != func() int32 {
		if (wb_c4go_postfix)[0].r_f >= 0 {
			return (wb_c4go_postfix)[0].r_f
		}
		return (nreg(int32('f')))[0]
	}() {
		// append font and size to the buffer if needed
		sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cf(%02d\x00"), c_ec, func() int32 {
			if (wb_c4go_postfix)[0].r_f >= 0 {
				return (wb_c4go_postfix)[0].r_f
			}
			return (nreg(int32('f')))[0]
		}())
		if (wb_c4go_postfix)[0].r_f >= 0 {
			wb_c4go_postfix[0].f = (wb_c4go_postfix)[0].r_f
		} else {
			wb_c4go_postfix[0].f = (nreg(int32('f')))[0]
		}
	}
	if wb_c4go_postfix[0].s != func() int32 {
		if (wb_c4go_postfix)[0].r_s >= 0 {
			return (wb_c4go_postfix)[0].r_s
		}
		return (nreg(int32('s')))[0]
	}() {
		if func() int32 {
			if (wb_c4go_postfix)[0].r_s >= 0 {
				return (wb_c4go_postfix)[0].r_s
			}
			return (nreg(int32('s')))[0]
		}() < 100 {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cs(%02d\x00"), c_ec, func() int32 {
				if (wb_c4go_postfix)[0].r_s >= 0 {
					return (wb_c4go_postfix)[0].r_s
				}
				return (nreg(int32('s')))[0]
			}())
		} else {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cs[%d]\x00"), c_ec, func() int32 {
				if (wb_c4go_postfix)[0].r_s >= 0 {
					return (wb_c4go_postfix)[0].r_s
				}
				return (nreg(int32('s')))[0]
			}())
		}
		if (wb_c4go_postfix)[0].r_s >= 0 {
			wb_c4go_postfix[0].s = (wb_c4go_postfix)[0].r_s
		} else {
			wb_c4go_postfix[0].s = (nreg(int32('s')))[0]
		}
	}
	if noarch.Not((nreg(int32('C')))[0]) && wb_c4go_postfix[0].m != func() int32 {
		if (wb_c4go_postfix)[0].r_m >= 0 {
			return (wb_c4go_postfix)[0].r_m
		}
		return (nreg(int32('m')))[0]
	}() {
		sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cm[%s]\x00"), c_ec, clr_str(func() int32 {
			if (wb_c4go_postfix)[0].r_m >= 0 {
				return (wb_c4go_postfix)[0].r_m
			}
			return (nreg(int32('m')))[0]
		}()))
		if (wb_c4go_postfix)[0].r_m >= 0 {
			wb_c4go_postfix[0].m = (wb_c4go_postfix)[0].r_m
		} else {
			wb_c4go_postfix[0].m = (nreg(int32('m')))[0]
		}
	}
	wb_stsb(wb_c4go_postfix)
}

// wb_flushdir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:97
func wb_flushdir(wb_c4go_postfix []wb) {
	if wb_c4go_postfix[0].cd != func() int32 {
		if (wb_c4go_postfix)[0].r_cd >= 0 {
			return (wb_c4go_postfix)[0].r_cd
		}
		return (nreg(map_([]byte(".cd\x00"))))[0]
	}() {
		// append current text direction to the buffer if needed
		wb_flushsub(wb_c4go_postfix)
		if dir_do != 0 {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%c%c\x00"), c_ec, func() int32 {
				if func() int32 {
					if (wb_c4go_postfix)[0].r_cd >= 0 {
						return (wb_c4go_postfix)[0].r_cd
					}
					return (nreg(map_([]byte(".cd\x00"))))[0]
				}() > 0 {
					return int32('<')
				}
				return int32('>')
			}())
		}
		if (wb_c4go_postfix)[0].r_cd >= 0 {
			wb_c4go_postfix[0].cd = (wb_c4go_postfix)[0].r_cd
		} else {
			wb_c4go_postfix[0].cd = (nreg(map_([]byte(".cd\x00"))))[0]
		}
	}
}

// wb_flush - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:108
func wb_flush(wb_c4go_postfix []wb) {
	// apply font and size changes and flush the collected subword
	wb_flushsub(wb_c4go_postfix)
	wb_flushdir(wb_c4go_postfix)
	wb_flushfont(wb_c4go_postfix)
}

// wb_hmov - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:115
func wb_hmov(wb_c4go_postfix []wb, n int32) {
	wb_flushsub(wb_c4go_postfix)
	wb_c4go_postfix[0].h += n
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%ch'%du'\x00"), c_ec, n)
}

// wb_vmov - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:122
func wb_vmov(wb_c4go_postfix []wb, n int32) {
	wb_flushsub(wb_c4go_postfix)
	wb_c4go_postfix[0].v += n
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cv'%du'\x00"), c_ec, n)
}

// wb_els - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:129
func wb_els(wb_c4go_postfix []wb, els int32) {
	wb_flushsub(wb_c4go_postfix)
	if els > wb_c4go_postfix[0].els_pos {
		wb_c4go_postfix[0].els_pos = els
	}
	if els < wb_c4go_postfix[0].els_neg {
		wb_c4go_postfix[0].els_neg = els
	}
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cx'%du'\x00"), c_ec, els)
}

// wb_etc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:139
func wb_etc(wb_c4go_postfix []wb, x []byte) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cX\x02%s\x02\x00"), c_ec, x)
}

// wb_putbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:145
func wb_putbuf(wb_c4go_postfix []wb, c []byte) {
	var g []glyph
	var zerowidth int32
	if int32(c[0]) == int32('\t') || int32(c[0]) == int32('\x01') || int32(c[0]) == 4 && (int32(c[1]) == int32('\t') || int32(c[1]) == int32('\x01')) {
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], c)
		return
	}
	g = dev_glyph(c, wb_c4go_postfix[0].f)
	zerowidth = c_hymark(c)
	if g == nil && int32(c[0]) == c_ec && noarch.Not(zerowidth) {
		// unknown escape
		noarch.Memmove(c, c[0+1:], uint32(noarch.Strlen(c)))
		g = dev_glyph(c, wb_c4go_postfix[0].f)
	}
	if g != nil && noarch.Not(zerowidth) && wb_c4go_postfix[0].icleft != 0 && func() int32 {
		if 0 < -int32((g)[0].llx) {
			return -int32((g)[0].llx)
		}
		return 0
	}() != 0 {
		wb_hmov(wb_c4go_postfix, font_wid(g[0].font, wb_c4go_postfix[0].s, func() int32 {
			if 0 < -int32((g)[0].llx) {
				return -int32((g)[0].llx)
			}
			return 0
		}()))
	}
	wb_c4go_postfix[0].icleft = 0
	if noarch.Not(c[1]) || int32(c[0]) == c_ec || int32(c[0]) == 4 || utf8one(c) != 0 {
		if int32(c[0]) == 4 && int32(c[1]) == c_ec {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%c%c\x00"), c_ec, c_ec)
		} else {
			sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], c)
		}
	} else {
		if int32(c[1]) != 0 && noarch.Not(c[2]) {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%c(%s\x00"), c_ec, c)
		} else {
			sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cC'%s'\x00"), c_ec, c)
		}
	}
	if noarch.Not(zerowidth) {
		if noarch.Not((nreg(int32('C')))[0]) && g != nil {
			if int32(g[0].llx) != 0 || int32(g[0].lly) != 0 || int32(g[0].urx) != 0 || int32(g[0].ury) != 0 {
				var llx int32 = font_wid(g[0].font, wb_c4go_postfix[0].s, int32(g[0].llx))
				var lly int32 = font_wid(g[0].font, wb_c4go_postfix[0].s, int32(g[0].lly))
				var urx int32 = font_wid(g[0].font, wb_c4go_postfix[0].s, int32(g[0].urx))
				var ury int32 = font_wid(g[0].font, wb_c4go_postfix[0].s, int32(g[0].ury))
				wb_bbox(wb_c4go_postfix, llx, lly, urx, ury)
			} else {
				// no bounding box information
				var ht int32 = wb_c4go_postfix[0].s * (dev_res / 72)
				var urx int32 = font_wid(g[0].font, wb_c4go_postfix[0].s, int32(g[0].wid))
				var lly int32 = func() int32 {
					if int32(g[0].type_)&1 != 0 {
						return -ht / 2
					}
					return 0
				}()
				var ury int32 = func() int32 {
					if int32(g[0].type_)&2 != 0 {
						return ht
					}
					return ht / 2
				}()
				wb_bbox(wb_c4go_postfix, 0, lly, urx, ury)
			}
		}
		wb_c4go_postfix[0].h += func() int32 {
			if g != nil {
				return font_gwid(g[0].font, dev_font(wb_c4go_postfix[0].f), wb_c4go_postfix[0].s, int32(g[0].wid))
			}
			return 0
		}()
		wb_c4go_postfix[0].ct |= func() int32 {
			if g != nil {
				return int32(g[0].type_)
			}
			return 0
		}()
		wb_stsb(wb_c4go_postfix)
	}
}

// wb_hyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:197
func wb_hyph(src [][]byte, src_n int32, src_hyph []byte, flg int32) int32 {
	// return nonzero if it cannot be hyphenated
	// word to pass to hyphenate()
	var word []byte = make([]byte, 8192)
	// hyphenation points of word
	var hyph []byte = make([]byte, 8192)
	// the mapping from src[] to word[]
	var smap []int32 = make([]int32, 256)
	var s []byte
	var d []byte
	var i int32
	d = word
	d[0] = '\x00'
	for i = 0; i < src_n; i++ {
		s = src[i]
		smap[i] = int32((int64(uintptr(unsafe.Pointer(&d[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&word[0])))/int64(1)))
		if c_hystop(s) != 0 {
			return 1
		}
		if c_hymark(s) != 0 {
			continue
		}
		d = d[0+hy_cput(d, s):]
	}
	noarch.Memset(hyph, byte(0), uint32(int32((int64(uintptr(unsafe.Pointer(&d[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&word[0])))/int64(1))))*1)
	hyphenate(hyph, word, flg)
	for i = 0; i < src_n; i++ {
		src_hyph[i] = hyph[smap[i]]
	}
	return 0
}

// wb_collect - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:222
func wb_collect(wb_c4go_postfix []wb, val int32) int32 {
	var old int32 = wb_c4go_postfix[0].sub_collect
	wb_c4go_postfix[0].sub_collect = val
	return old
}

// wb_flushsub - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:230
func wb_flushsub(wb_c4go_postfix []wb) {
	// output the collected characters; only for those present in wb->f font
	var fn []font
	var gsrc [][]glyph = make([][]glyph, 256)
	var gdst [][]glyph = make([][]glyph, 256)
	var x []int32 = make([]int32, 256)
	var y []int32 = make([]int32, 256)
	var xadv []int32 = make([]int32, 256)
	var yadv []int32 = make([]int32, 256)
	var dmap []int32 = make([]int32, 256)
	var src_hyph []byte = make([]byte, 256)
	var dst_n int32
	var i int32
	var sidx int32
	if noarch.Not(wb_c4go_postfix[0].sub_n) || noarch.Not(wb_c4go_postfix[0].sub_collect) {
		return
	}
	wb_c4go_postfix[0].sub_collect = 0
	fn = dev_font(wb_c4go_postfix[0].f)
	if noarch.Not((nreg(map_([]byte(".hy\x00"))))[0]) || wb_hyph(wb_c4go_postfix[0].sub_c[:], wb_c4go_postfix[0].sub_n, src_hyph, (nreg(map_([]byte(".hy\x00"))))[0]) != 0 {
		noarch.Memset(src_hyph, byte(0), 256)
	}
	for sidx < wb_c4go_postfix[0].sub_n {
		// call font_layout() for collected glyphs; skip hyphenation marks
		var beg int32 = sidx
		for ; sidx < wb_c4go_postfix[0].sub_n && noarch.Not(c_hymark(wb_c4go_postfix[0].sub_c[:][sidx])); sidx++ {
			gsrc[sidx-beg] = font_find(fn, wb_c4go_postfix[0].sub_c[:][sidx])
		}
		dst_n = font_layout(fn, gsrc, sidx-beg, wb_c4go_postfix[0].s, gdst, dmap, x, y, xadv, yadv, (nreg(map_([]byte(".lg\x00"))))[0], (nreg(map_([]byte(".kn\x00"))))[0])
		for i = 0; i < dst_n; i++ {
			var xd []int32 = []int32{x[i], xadv[i] - x[i]}
			var yd []int32 = []int32{y[i], yadv[i] - y[i]}
			if xd[wb_c4go_postfix[0].cd] != 0 {
				wb_hmov(wb_c4go_postfix, font_wid(fn, wb_c4go_postfix[0].s, xd[wb_c4go_postfix[0].cd]))
			}
			if yd[wb_c4go_postfix[0].cd] != 0 {
				wb_vmov(wb_c4go_postfix, font_wid(fn, wb_c4go_postfix[0].s, yd[wb_c4go_postfix[0].cd]))
			}
			if src_hyph[beg+dmap[i]] != 0 {
				wb_putbuf(wb_c4go_postfix, env_hc())
			}
			if (int64(uintptr(unsafe.Pointer(&gdst[i])))/int64(120) - int64(uintptr(unsafe.Pointer(&gsrc[dmap[i]])))/int64(120)) == 0 {
				wb_putbuf(wb_c4go_postfix, wb_c4go_postfix[0].sub_c[:][beg+dmap[i]])
			} else {
				wb_putbuf(wb_c4go_postfix, gdst[i][0].name[:])
			}
			if xd[1-wb_c4go_postfix[0].cd] != 0 {
				wb_hmov(wb_c4go_postfix, font_wid(fn, wb_c4go_postfix[0].s, xd[1-wb_c4go_postfix[0].cd]))
			}
			if yd[1-wb_c4go_postfix[0].cd] != 0 {
				wb_vmov(wb_c4go_postfix, font_wid(fn, wb_c4go_postfix[0].s, yd[1-wb_c4go_postfix[0].cd]))
			}
		}
		for ; sidx < wb_c4go_postfix[0].sub_n && c_hymark(wb_c4go_postfix[0].sub_c[:][sidx]) != 0; sidx++ {
			wb_putbuf(wb_c4go_postfix, wb_c4go_postfix[0].sub_c[:][sidx])
		}
	}
	wb_c4go_postfix[0].sub_n = 0
	wb_c4go_postfix[0].icleft = 0
	wb_c4go_postfix[0].sub_collect = 1
}

// wb_put - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:279
func wb_put(wb_c4go_postfix []wb, c []byte) {
	if int32(c[0]) == int32('\n') {
		wb_c4go_postfix[0].part = 0
		return
	}
	if wb_pendingdir(wb_c4go_postfix) != 0 {
		wb_flushdir(wb_c4go_postfix)
	}
	if int32(c[0]) == int32(' ') {
		wb_flushsub(wb_c4go_postfix)
		wb_hmov(wb_c4go_postfix, font_swid(dev_font(func() int32 {
			if (wb_c4go_postfix)[0].r_f >= 0 {
				return (wb_c4go_postfix)[0].r_f
			}
			return (nreg(int32('f')))[0]
		}()), func() int32 {
			if (wb_c4go_postfix)[0].r_s >= 0 {
				return (wb_c4go_postfix)[0].r_s
			}
			return (nreg(int32('s')))[0]
		}(), (nreg(map_([]byte(".ss\x00"))))[0]))
		return
	}
	if noarch.Not(noarch.Strcmp([]byte("\\~\x00"), c)) {
		wb_flushsub(wb_c4go_postfix)
		sbuf_append((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], c)
		wb_c4go_postfix[0].h += font_swid(dev_font(func() int32 {
			if (wb_c4go_postfix)[0].r_f >= 0 {
				return (wb_c4go_postfix)[0].r_f
			}
			return (nreg(int32('f')))[0]
		}()), func() int32 {
			if (wb_c4go_postfix)[0].r_s >= 0 {
				return (wb_c4go_postfix)[0].r_s
			}
			return (nreg(int32('s')))[0]
		}(), (nreg(map_([]byte(".ss\x00"))))[0])
		return
	}
	if wb_pendingfont(wb_c4go_postfix) != 0 || uint32(wb_c4go_postfix[0].sub_n) == 8192/32 {
		wb_flush(wb_c4go_postfix)
	}
	if wb_c4go_postfix[0].sub_collect != 0 {
		if font_find(dev_font(wb_c4go_postfix[0].f), c) != nil || c_hymark(c) != 0 {
			noarch.Strcpy(wb_c4go_postfix[0].sub_c[:][func() int32 {
				tempVar1 := &wb_c4go_postfix[0].sub_n
				defer func() {
					*tempVar1++
				}()
				return *tempVar1
			}()], c)
		} else {
			wb_putraw(wb_c4go_postfix, c)
		}
	} else {
		wb_putbuf(wb_c4go_postfix, c)
	}
}

// wb_putraw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:311
func wb_putraw(wb_c4go_postfix []wb, c []byte) {
	// just like wb_put() but disable subword collection
	var collect int32
	wb_flushsub(wb_c4go_postfix)
	collect = wb_collect(wb_c4go_postfix, 0)
	wb_put(wb_c4go_postfix, c)
	wb_collect(wb_c4go_postfix, collect)
}

// wb_putexpand - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:321
func wb_putexpand(wb_c4go_postfix []wb, c []byte) {
	if cdef_expand(wb_c4go_postfix, c, func() int32 {
		if (wb_c4go_postfix)[0].r_f >= 0 {
			return (wb_c4go_postfix)[0].r_f
		}
		return (nreg(int32('f')))[0]
	}()) != 0 {
		// just like wb_put(), but call cdef_expand() if c is defined
		wb_put(wb_c4go_postfix, c)
	}
}

// wb_part - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:327
func wb_part(wb_c4go_postfix []wb) int32 {
	return wb_c4go_postfix[0].part
}

// wb_setpart - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:332
func wb_setpart(wb_c4go_postfix []wb) {
	wb_c4go_postfix[0].part = 1
}

// wb_cost - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:337
func wb_cost(wb_c4go_postfix []wb) int32 {
	return wb_c4go_postfix[0].cost
}

// wb_setcost - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:342
func wb_setcost(wb_c4go_postfix []wb, cost int32) {
	wb_c4go_postfix[0].cost = cost
}

// wb_drawl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:347
func wb_drawl(wb_c4go_postfix []wb, c int32, h int32, v int32) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cD'%c %du %du'\x00"), c_ec, c, h, v)
	wb_c4go_postfix[0].h += h
	wb_c4go_postfix[0].v += v
	wb_stsb(wb_c4go_postfix)
}

// wb_drawc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:356
func wb_drawc(wb_c4go_postfix []wb, c int32, r int32) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cD'%c %du'\x00"), c_ec, c, r)
	wb_c4go_postfix[0].h += r
}

// wb_drawe - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:363
func wb_drawe(wb_c4go_postfix []wb, c int32, h int32, v int32) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cD'%c %du %du'\x00"), c_ec, c, h, v)
	wb_c4go_postfix[0].h += h
}

// wb_drawa - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:370
func wb_drawa(wb_c4go_postfix []wb, c int32, h1 int32, v1 int32, h2 int32, v2 int32) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cD'%c %du %du %du %du'\x00"), c_ec, c, h1, v1, h2, v2)
	wb_c4go_postfix[0].h += h1 + h2
	wb_c4go_postfix[0].v += v1 + v2
	wb_stsb(wb_c4go_postfix)
}

// wb_drawxbeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:380
func wb_drawxbeg(wb_c4go_postfix []wb, c int32) {
	wb_flush(wb_c4go_postfix)
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("%cD'%c\x00"), c_ec, c)
}

// wb_drawxdot - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:386
func wb_drawxdot(wb_c4go_postfix []wb, h int32, v int32) {
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte(" %du %du\x00"), h, v)
	wb_c4go_postfix[0].h += h
	wb_c4go_postfix[0].v += v
	wb_stsb(wb_c4go_postfix)
}

// wb_drawxcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:394
func wb_drawxcmd(wb_c4go_postfix []wb, cmd []byte) {
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte(" %s\x00"), cmd)
}

// wb_drawxend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:399
func wb_drawxend(wb_c4go_postfix []wb) {
	sbuf_printf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:], []byte("'\x00"))
}

// wb_reset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:404
func wb_reset(wb_c4go_postfix []wb) {
	wb_done(wb_c4go_postfix)
	wb_init(wb_c4go_postfix)
}

// wb_buf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:410
func wb_buf(wb_c4go_postfix []wb) []byte {
	wb_flushsub(wb_c4go_postfix)
	return sbuf_buf((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:])
}

// wb_putc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:416
func wb_putc(wb_c4go_postfix []wb, t int32, s []byte) {
	if t != 0 && t != int32('C') {
		wb_flushsub(wb_c4go_postfix)
	}
	switch t {
	case 0:
		fallthrough
	case 'C':
		wb_put(wb_c4go_postfix, s)
	case 'D':
		ren_dcmd(wb_c4go_postfix, s)
	case 'f':
		wb_c4go_postfix[0].r_f = noarch.Atoi(s)
	case 'h':
		wb_hmov(wb_c4go_postfix, noarch.Atoi(s))
	case 'm':
		wb_c4go_postfix[0].r_m = clr_get(s)
	case 's':
		wb_c4go_postfix[0].r_s = noarch.Atoi(s)
	case 'v':
		wb_vmov(wb_c4go_postfix, noarch.Atoi(s))
	case 'x':
		wb_els(wb_c4go_postfix, noarch.Atoi(s))
	case 'X':
		wb_etc(wb_c4go_postfix, s)
	case '<':
		fallthrough
	case '>':
		wb_c4go_postfix[0].r_cd = noarch.BoolToInt(t == int32('<'))
		wb_flushdir(wb_c4go_postfix)
		break
	}
}

// wb_cat - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:457
func wb_cat(wb_c4go_postfix []wb, src []wb) {
	var s []byte
	var d []byte
	var c int32
	var part int32
	var collect int32
	wb_flushsub(src)
	wb_flushsub(wb_c4go_postfix)
	collect = wb_collect(wb_c4go_postfix, 0)
	s = sbuf_buf((*[1000000]sbuf)(unsafe.Pointer(&src[0].sbuf))[:])
	for (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) >= 0 {
		wb_putc(wb_c4go_postfix, c, d)
	}
	part = src[0].part
	wb_c4go_postfix[0].r_s = -1
	wb_c4go_postfix[0].r_f = -1
	wb_c4go_postfix[0].r_m = -1
	wb_c4go_postfix[0].r_cd = -1
	wb_reset(src)
	src[0].part = part
	wb_collect(wb_c4go_postfix, collect)
}

// wb_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:478
func wb_wid(wb_c4go_postfix []wb) int32 {
	wb_flushsub(wb_c4go_postfix)
	return wb_c4go_postfix[0].h
}

// wb_hpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:484
func wb_hpos(wb_c4go_postfix []wb) int32 {
	wb_flushsub(wb_c4go_postfix)
	return wb_c4go_postfix[0].h
}

// wb_vpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:490
func wb_vpos(wb_c4go_postfix []wb) int32 {
	wb_flushsub(wb_c4go_postfix)
	return wb_c4go_postfix[0].v
}

// wb_empty - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:496
func wb_empty(wb_c4go_postfix []wb) int32 {
	return noarch.BoolToInt(noarch.Not(wb_c4go_postfix[0].sub_n) && sbuf_empty((*[1000000]sbuf)(unsafe.Pointer(&wb_c4go_postfix[0].sbuf))[:]) != 0)
}

// wb_eos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:502
func wb_eos(wb_c4go_postfix []wb) int32 {
	// return 1 if wb ends a sentence (.?!)
	var i int32 = wb_c4go_postfix[0].sub_n - 1
	for i > 0 && c_eostran(wb_c4go_postfix[0].sub_c[:][i]) != 0 {
		i--
	}
	return noarch.BoolToInt(i >= 0 && c_eossent(wb_c4go_postfix[0].sub_c[:][i]) != 0)
}

// wb_wconf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:510
func wb_wconf(wb_c4go_postfix []wb, ct []int32, st []int32, sb []int32, llx []int32, lly []int32, urx []int32, ury []int32) {
	wb_flushsub(wb_c4go_postfix)
	ct[0] = wb_c4go_postfix[0].ct
	st[0] = -wb_c4go_postfix[0].st
	sb[0] = -wb_c4go_postfix[0].sb
	if wb_c4go_postfix[0].llx < 1<<uint64(29) {
		llx[0] = wb_c4go_postfix[0].llx
	} else {
		llx[0] = 0
	}
	if wb_c4go_postfix[0].lly < 1<<uint64(29) {
		lly[0] = -wb_c4go_postfix[0].lly
	} else {
		lly[0] = 0
	}
	if wb_c4go_postfix[0].urx > -(1 << uint64(29)) {
		urx[0] = wb_c4go_postfix[0].urx
	} else {
		urx[0] = 0
	}
	if wb_c4go_postfix[0].ury > -(1 << uint64(29)) {
		ury[0] = -wb_c4go_postfix[0].ury
	} else {
		ury[0] = 0
	}
}

// wb_prevglyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:523
func wb_prevglyph(wb_c4go_postfix []wb) []glyph {
	if wb_c4go_postfix[0].sub_n != 0 {
		return dev_glyph(wb_c4go_postfix[0].sub_c[:][wb_c4go_postfix[0].sub_n-1], wb_c4go_postfix[0].f)
	}
	return nil
}

// wb_italiccorrection - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:528
func wb_italiccorrection(wb_c4go_postfix []wb) {
	var g []glyph = wb_prevglyph(wb_c4go_postfix)
	if g != nil && func() int32 {
		if 0 < int32((g)[0].urx)-int32((g)[0].wid) {
			return int32((g)[0].urx) - int32((g)[0].wid)
		}
		return 0
	}() != 0 {
		wb_hmov(wb_c4go_postfix, font_wid(g[0].font, wb_c4go_postfix[0].s, func() int32 {
			if 0 < int32((g)[0].urx)-int32((g)[0].wid) {
				return int32((g)[0].urx) - int32((g)[0].wid)
			}
			return 0
		}()))
	}
}

// wb_italiccorrectionleft - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:535
func wb_italiccorrectionleft(wb_c4go_postfix []wb) {
	wb_flushsub(wb_c4go_postfix)
	wb_c4go_postfix[0].icleft = 1
}

// wb_fnszget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:541
func wb_fnszget(wb_c4go_postfix []wb, fn []int32, sz []int32, m []int32, cd []int32) {
	wb_flushsub(wb_c4go_postfix)
	fn[0] = wb_c4go_postfix[0].r_f
	sz[0] = wb_c4go_postfix[0].r_s
	m[0] = wb_c4go_postfix[0].r_m
	cd[0] = wb_c4go_postfix[0].r_cd
}

// wb_fnszset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:550
func wb_fnszset(wb_c4go_postfix []wb, fn int32, sz int32, m int32, cd int32) {
	wb_c4go_postfix[0].r_f = fn
	wb_c4go_postfix[0].r_s = sz
	wb_c4go_postfix[0].r_m = m
	wb_c4go_postfix[0].r_cd = cd
}

// wb_catstr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:558
func wb_catstr(wb_c4go_postfix []wb, s []byte, end []byte) {
	var collect int32
	var c int32
	var d []byte
	wb_flushsub(wb_c4go_postfix)
	collect = wb_collect(wb_c4go_postfix, 0)
	for (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&end[0])))/int64(1)) < 0 && (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) >= 0 {
		wb_putc(wb_c4go_postfix, c, d)
	}
	wb_collect(wb_c4go_postfix, collect)
}

// wb_hywid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:570
func wb_hywid(wb_c4go_postfix []wb) int32 {
	// return the size of \(hy if appended to wb
	var g []glyph = dev_glyph([]byte("hy\x00"), wb_c4go_postfix[0].f)
	if g != nil {
		return font_gwid(g[0].font, dev_font(func() int32 {
			if (wb_c4go_postfix)[0].r_f >= 0 {
				return (wb_c4go_postfix)[0].r_f
			}
			return (nreg(int32('f')))[0]
		}()), func() int32 {
			if (wb_c4go_postfix)[0].r_s >= 0 {
				return (wb_c4go_postfix)[0].r_s
			}
			return (nreg(int32('s')))[0]
		}(), int32(g[0].wid))
	}
	return 0
}

// wb_swid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:577
func wb_swid(wb_c4go_postfix []wb) int32 {
	// return the size of space if appended to wb
	return font_swid(dev_font(func() int32 {
		if (wb_c4go_postfix)[0].r_f >= 0 {
			return (wb_c4go_postfix)[0].r_f
		}
		return (nreg(int32('f')))[0]
	}()), func() int32 {
		if (wb_c4go_postfix)[0].r_s >= 0 {
			return (wb_c4go_postfix)[0].r_s
		}
		return (nreg(int32('s')))[0]
	}(), (nreg(map_([]byte(".ss\x00"))))[0])
}

// keshideh_chars - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:582
var keshideh_chars [][]byte = [][]byte{[]byte("ﺒ\x00"), []byte("ﺑ\x00"), []byte("ﭙ\x00"), []byte("ﭘ\x00"), []byte("ﺘ\x00"), []byte("ﺗ\x00"), []byte("ﺜ\x00"), []byte("ﺛ\x00"), []byte("ﺴ\x00"), []byte("ﺳ\x00"), []byte("ﺸ\x00"), []byte("ﺷ\x00"), []byte("ﻔ\x00"), []byte("ﻓ\x00"), []byte("ﻘ\x00"), []byte("ﻗ\x00"), []byte("ﮑ\x00"), []byte("ﮐ\x00"), []byte("ﮕ\x00"), []byte("ﮔ\x00"), []byte("ﻤ\x00"), []byte("ﻣ\x00"), []byte("ﻨ\x00"), []byte("ﻧ\x00"), []byte("ﻬ\x00"), []byte("ﻫ\x00"), []byte("ﯿ\x00"), []byte("ﯾ\x00")}

// keshideh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:588
func keshideh(c []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 224/8; i++ {
		if noarch.Not(noarch.Strcmp(keshideh_chars[i], c)) {
			return 1
		}
	}
	return 0
}

// wb_keshideh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/wb.c:598
func wb_keshideh(word []byte, dst []wb, wid int32) int32 {
	// insert keshideh
	var p []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var s []byte
	var d []byte
	var s_prev []byte
	var s_kesh []byte
	var ins int32
	var c int32
	// find the last keshideh position
	s = word
	for (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) >= 0 {
		wb_putc(dst, c, d)
		if noarch.Not(c) && keshideh(p) != 0 {
			var g []glyph = dev_glyph([]byte("ـ\x00"), func() int32 {
				if (dst)[0].r_f >= 0 {
					return (dst)[0].r_f
				}
				return (nreg(int32('f')))[0]
			}())
			var kw int32 = func() int32 {
				if g != nil {
					return font_gwid(g[0].font, dev_font(func() int32 {
						if (dst)[0].r_f >= 0 {
							return (dst)[0].r_f
						}
						return (nreg(int32('f')))[0]
					}()), func() int32 {
						if (dst)[0].r_s >= 0 {
							return (dst)[0].r_s
						}
						return (nreg(int32('s')))[0]
					}(), int32(g[0].wid))
				}
				return 0
			}()
			if g != nil && kw < wid {
				s_kesh = s_prev
				ins = kw
			}
		}
		s_prev = s
		noarch.Strcpy(p, func() []byte {
			if c != 0 {
				return []byte("\x00")
			}
			return d
		}())
	}
	// insert the keshideh at s_kesh
	s = word
	wb_reset(dst)
	for (func() int32 {
		c = escread((*[1000000][]byte)(unsafe.Pointer(&s))[:], (*[1000000][]byte)(unsafe.Pointer(&d))[:])
		return c
	}()) >= 0 {
		wb_putc(dst, c, d)
		if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s_kesh[0])))/int64(1)) == 0 {
			wb_putc(dst, 0, []byte("ـ\x00"))
		}
	}
	return ins
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[1000000]int32)(unsafe.Pointer(c4go_name))[:]
}

// c4goUnsafeConvert_sbuf : created by c4go
func c4goUnsafeConvert_sbuf(c4go_name *sbuf) []sbuf {
	return (*[1000000]sbuf)(unsafe.Pointer(c4go_name))[:]
}

// c4goUnsafeConvert_wb : created by c4go
func c4goUnsafeConvert_wb(c4go_name *wb) []wb {
	return (*[1000000]wb)(unsafe.Pointer(c4go_name))[:]
}

// the last output font and size
// current font and size; use n_f and n_s if -1
// partial input (\c)
// the extra cost of line break after this word
// extra line spacing
// buffer vertical and horizontal positions
// \w registers
// bounding box
// pending left italic correction
// queued subword
// the collected subword
// collected subword length
// enable subword collection
// character translation (.tr)
// character definition (.char)
// hyphenation flags
// adjustment types
// line formatting
// rendering
// the main loop
// horizontal line
// \l
// \L
// \b
// \o
// \D
// \Z
// out.c
// output rendered line
// output \X requests
// output troff cmd
// troff commands
// helpers
// utf-8 parsing
// reading escapes and characters
// string streams; nested next()/back() interface for string buffers
// internal commands
// mapping register, macro and environment names to indices
// map name s to an index
// return the name mapped to id
// text direction
// colors
// builtin number registers; n_X for .X register
// functions for implementing read-only registers
// .t
// .z
// .k

// __ctype_b_loc from ctype.h
// c function : const unsigned short int** __ctype_b_loc()
// dep pkg    : unicode
// dep func   :
func __ctype_b_loc() [][]uint16 {
	var characterTable []uint16

	for i := 0; i < 255; i++ {
		var c uint16

		// Each of the bitwise expressions below were copied from the enum
		// values, like _ISupper, etc.

		if unicode.IsUpper(rune(i)) {
			c |= ((1 << (0)) << 8)
		}

		if unicode.IsLower(rune(i)) {
			c |= ((1 << (1)) << 8)
		}

		if unicode.IsLetter(rune(i)) {
			c |= ((1 << (2)) << 8)
		}

		if unicode.IsDigit(rune(i)) {
			c |= ((1 << (3)) << 8)
		}

		if unicode.IsDigit(rune(i)) ||
			(i >= 'a' && i <= 'f') ||
			(i >= 'A' && i <= 'F') {
			// IsXDigit. This is the same implementation as the Mac version.
			// There may be a better way to do this.
			c |= ((1 << (4)) << 8)
		}

		if unicode.IsSpace(rune(i)) {
			c |= ((1 << (5)) << 8)
		}

		if unicode.IsPrint(rune(i)) {
			c |= ((1 << (6)) << 8)
		}

		// The IsSpace check is required because Go treats spaces as graphic
		// characters, which C does not.
		if unicode.IsGraphic(rune(i)) && !unicode.IsSpace(rune(i)) {
			c |= ((1 << (7)) << 8)
		}

		// http://www.cplusplus.com/reference/cctype/isblank/
		// The standard "C" locale considers blank characters the tab
		// character ('\t') and the space character (' ').
		if i == int('\t') || i == int(' ') {
			c |= ((1 << (8)) >> 8)
		}

		if unicode.IsControl(rune(i)) {
			c |= ((1 << (9)) >> 8)
		}

		if unicode.IsPunct(rune(i)) {
			c |= ((1 << (10)) >> 8)
		}

		if unicode.IsLetter(rune(i)) || unicode.IsDigit(rune(i)) {
			c |= ((1 << (11)) >> 8)
		}

		// Yes, I know this is a hideously slow way to do it but I just want to
		// test if this works right now.
		characterTable = append(characterTable, c)
	}
	return [][]uint16{characterTable}
}

// tolower from ctype.h
// c function : int tolower(int)
// dep pkg    : unicode
// dep func   :
func tolower(_c int32) int32 {
	return int32(unicode.ToLower(rune(_c)))
}

// memcpy is function from string.h.
// c function : void * memcpy( void * , const void * , size_t )
// dep pkg    : reflect
// dep func   :
func memcpy(dst, src interface{}, size uint32) interface{} {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)
		d := reflect.ValueOf(dst)
		if s.Len() == 0 {
			return dst
		}
		if s.Len() > 0 {
			size /= uint32(int(s.Index(0).Type().Size()))
		}
		var val reflect.Value
		for i := 0; i < int(size); i++ {
			if i < s.Len() {
				val = s.Index(i)
			}
			d.Index(i).Set(val)
		}
	}
	return dst
}

// Warning cannot generate return type binding function `getpid`: cannot resolve type '__pid_t' : I couldn't find an appropriate Go type for the C type '__pid_t'.

// memchr - add c-binding for implemention function
func memchr(arg0 interface{}, arg1 int32, arg2 uint32) interface{} {
	return interface{}(C.memchr(unsafe.Pointer(&arg0), C.int(arg1), C.ulong(arg2)))
}

// strncmp - add c-binding for implemention function
func strncmp(arg0 []byte, arg1 []byte, arg2 uint32) int32 {
	return int32(C.strncmp((*C.char)(unsafe.Pointer(&arg0[0])), (*C.char)(unsafe.Pointer(&arg1[0])), C.ulong(arg2)))
}

// va_list is C4GO implementation of va_list from "stdarg.h"
type va_list struct {
	position int
	Slice    []interface{}
}

func create_va_list(list []interface{}) *va_list {
	return &va_list{
		position: 0,
		Slice:    list,
	}
}

func va_start(v *va_list, count interface{}) {
	v.position = 0
}

func va_end(v *va_list) {
	// do nothing
}

func va_arg(v *va_list) interface{} {
	defer func() {
		v.position++
	}()
	value := v.Slice[v.position]
	switch value.(type) {
	case int:
		return int32(value.(int))
	default:
		return value
	}
}

// c4goPointerArithDivSlice - function of pointer arithmetic. generated by c4go
func c4goPointerArithDivSlice(slice []div, position int) []div {
	if position < 0 {
		// invert sign
		position = -position

		// Example from: go101.org/article/unsafe.html
		// repair size of slice
		var hdr reflect.SliceHeader
		sliceLen := len(slice)
		hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
		runtime.KeepAlive(&slice[0]) // needed!
		hdr.Len = sliceLen + int(position)
		hdr.Cap = hdr.Len
		slice = *((*[]div)(unsafe.Pointer(&hdr)))
		return slice
	}
	// position >= 0:
	return slice[position:]
}

// c4goPointerArithByteSlice - function of pointer arithmetic. generated by c4go
func c4goPointerArithByteSlice(slice []byte, position int) []byte {
	if position < 0 {
		// invert sign
		position = -position

		// Example from: go101.org/article/unsafe.html
		// repair size of slice
		var hdr reflect.SliceHeader
		sliceLen := len(slice)
		hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
		runtime.KeepAlive(&slice[0]) // needed!
		hdr.Len = sliceLen + int(position)
		hdr.Cap = hdr.Len
		slice = *((*[]byte)(unsafe.Pointer(&hdr)))
		return slice
	}
	// position >= 0:
	return slice[position:]
}

// c4goPointerArithInt32Slice - function of pointer arithmetic. generated by c4go
func c4goPointerArithInt32Slice(slice []int32, position int) []int32 {
	if position < 0 {
		// invert sign
		position = -position

		// Example from: go101.org/article/unsafe.html
		// repair size of slice
		var hdr reflect.SliceHeader
		sliceLen := len(slice)
		hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
		runtime.KeepAlive(&slice[0]) // needed!
		hdr.Len = sliceLen + int(position)
		hdr.Cap = hdr.Len
		slice = *((*[]int32)(unsafe.Pointer(&hdr)))
		return slice
	}
	// position >= 0:
	return slice[position:]
}
