//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

// #include </usr/include/string.h>
// #include </usr/include/stdio.h>
import "C"

import (
	"os"
	"reflect"
	"runtime"
	"unicode"
	"unsafe"

	"github.com/Konstantin8105/c4go/noarch"
	"golang.org/x/sys/unix"
)

// glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.h:19
// predefined array limits
// device related variables
type glyph struct {
	id    [32]byte
	name  [32]byte
	font  []font
	wid   int32
	type_ int32
	pos   int32
}

// ps_title - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:24
//
// * NEATPOST: NEATROFF'S POSTSCRIPT/PDF POSTPROCESSOR
// *
// * Copyright (C) 2013-2020 Ali Gholami Rudi <ali at rudi dot ir>
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
// document title
var ps_title []byte

// ps_pagewidth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:25
// page width (tenths of a millimetre)
var ps_pagewidth int32 = 2159

// ps_pageheight - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:26
// page height (tenths of a millimetre)
var ps_pageheight int32 = 2794

// ps_linewidth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:27
// drawing line thickness in thousandths of an em
var ps_linewidth int32 = 40

// o_pages - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:28
// output pages
var o_pages int32

// mark_desc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:31
// bookmarks
// bookmark description
var mark_desc [][]byte = make([][]byte, 256)

// mark_page - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:32
// bookmark page
var mark_page []int32

// mark_offset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:33
// bookmark offset
var mark_offset []int32

// mark_level - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:34
// bookmark level
var mark_level []int32

// mark_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:35
// number of bookmarks
var mark_n int32

// mark_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:36
// allocated size of bookmark arrays
var mark_sz int32

// name_desc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:39
// named destinations
// reference name
var name_desc [][]byte = make([][]byte, 64)

// name_page - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:40
// reference page
var name_page []int32

// name_offset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:41
// reference offset
var name_offset []int32

// name_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:42
// number of references
var name_n int32

// name_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:43
// allocated size of name arrays
var name_sz int32

// next - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:45
func next() int32 {
	return noarch.Fgetc(noarch.Stdin)
}

// back - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:50
func back(c int32) {
	ungetc(c, noarch.Stdin)
}

// utf8len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:55
func utf8len(c int32) int32 {
	if ^c&192 != 0 {
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

// nextutf8 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:68
func nextutf8(s []byte) int32 {
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

// nextskip - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:83
func nextskip() {
	// skip blanks
	var c int32
	for {
		c = next()
		if noarch.Not(int32(((__ctype_b_loc())[0])[c]) & int32(uint16(noarch.ISspace))) {
			break
		}
	}
	back(c)
}

// nextnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:92
func nextnum() int32 {
	var c int32
	var n int32
	var neg int32
	nextskip()
	for 1 != 0 {
		c = next()
		if noarch.Not(n) && (c == int32('-') || c == int32('+')) {
			neg = noarch.BoolToInt(c == int32('-'))
			continue
		}
		if noarch.Not(int32(((__ctype_b_loc())[0])[c]) & int32(uint16(noarch.ISdigit))) {
			back(c)
		}
		if c < 0 || noarch.Not(int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit))) {
			break
		}
		n = n*10 + c - int32('0')
	}
	if neg != 0 {
		return -n
	}
	return n
}

// readnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:113
func readnum(n []int32) int32 {
	var c int32
	for {
		c = next()
		if !(c == int32(' ')) {
			break
		}
	}
	back(c)
	if c == int32('-') || c == int32('+') || c >= int32('0') && c <= int32('9') {
		n[0] = nextnum()
		return 0
	}
	return 1
}

// iseol - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:127
func iseol() int32 {
	var c int32
	for {
		c = next()
		if !(c == int32(' ')) {
			break
		}
	}
	back(c)
	return noarch.BoolToInt(c == int32('\n'))
}

// nexteol - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:138
func nexteol() {
	// skip until the end of line
	var c int32
	for {
		c = next()
		if !(c >= 0 && c != int32('\n')) {
			break
		}
	}
}

// nextword - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:146
func nextword(s []byte) {
	var c int32
	nextskip()
	c = next()
	for c >= 0 && noarch.Not(int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISspace))) {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(c)
		c = next()
	}
	if c >= 0 {
		back(c)
	}
	s[0] = '\x00'
}

// readln - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:161
func readln(s []byte) {
	// read until eol
	var c int32
	c = next()
	for c > 0 && c != int32('\n') {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(c)
		c = next()
	}
	if c == int32('\n') {
		back(c)
	}
	s[0] = '\x00'
}

// postline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:174
func postline() {
	var h int32
	var v int32
	for noarch.Not(readnum(c4goUnsafeConvert_int32(&h))) && noarch.Not(readnum(c4goUnsafeConvert_int32(&v))) {
		drawl(h, v)
	}
}

// postarc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:181
func postarc() {
	var h1 int32
	var v1 int32
	var h2 int32
	var v2 int32
	if noarch.Not(readnum(c4goUnsafeConvert_int32(&h1))) && noarch.Not(readnum(c4goUnsafeConvert_int32(&v1))) && noarch.Not(readnum(c4goUnsafeConvert_int32(&h2))) && noarch.Not(readnum(c4goUnsafeConvert_int32(&v2))) {
		drawa(h1, v1, h2, v2)
	}
}

// postspline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:188
func postspline() {
	var h2 int32
	var v2 int32
	var h1 int32 = nextnum()
	var v1 int32 = nextnum()
	if iseol() != 0 {
		drawl(h1, v1)
		return
	}
	for noarch.Not(readnum(c4goUnsafeConvert_int32(&h2))) && noarch.Not(readnum(c4goUnsafeConvert_int32(&v2))) {
		draws(h1, v1, h2, v2)
		h1 = h2
		v1 = v2
	}
	draws(h1, v1, 0, 0)
}

// postpoly - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:205
func postpoly() {
	var l int32 = int32('l')
	var c int32
	for noarch.Not(iseol()) && (l == int32('l') || l == int32('~') || l == int32('a')) {
		for {
			c = next()
			if !(c == int32(' ')) {
				break
			}
		}
		back(c)
		if c != int32('-') && c != int32('+') && (c < int32('0') || c > int32('9')) {
			l = c
			for c >= 0 && noarch.Not(int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISspace))) {
				c = next()
			}
			continue
		}
		if l == int32('l') {
			postline()
		}
		if l == int32('~') {
			postspline()
		}
		if l == int32('a') {
			postarc()
		}
	}
}

// postdraw - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:229
func postdraw() {
	var h1 int32
	var v1 int32
	var c int32 = next()
	drawbeg()
	switch tolower(c) {
	case 'l':
		h1 = nextnum()
		v1 = nextnum()
		drawl(h1, v1)
	case 'c':
		drawc(nextnum())
	case 'e':
		h1 = nextnum()
		v1 = nextnum()
		drawe(h1, v1)
	case 'a':
		postarc()
	case '~':
		postspline()
	case 'p':
		postpoly()
		break
	}
	drawend(noarch.BoolToInt(c == int32('p') || c == int32('P')), noarch.BoolToInt(c == int32('E') || c == int32('C') || c == int32('P')))
	nexteol()
}

// strcut - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:262
func strcut(dst []byte, src []byte) []byte {
	for int32(src[0]) == int32(' ') || int32(src[0]) == int32('\n') {
		src = src[0+1:]
	}
	if int32(src[0]) == int32('"') {
		src = src[0+1:]
		for int32(src[0]) != 0 && (int32(src[0]) != int32('"') || int32(src[1]) == int32('"')) {
			if int32(src[0]) == int32('"') {
				src = src[0+1:]
			}
			(func() []byte {
				defer func() {
					dst = dst[0+1:]
				}()
				return dst
			}())[0] = (func() []byte {
				defer func() {
					src = src[0+1:]
				}()
				return src
			}())[0]
		}
		if int32(src[0]) == int32('"') {
			src = src[0+1:]
		}
	} else {
		for int32(src[0]) != 0 && int32(src[0]) != int32(' ') && int32(src[0]) != int32('\n') {
			(func() []byte {
				defer func() {
					dst = dst[0+1:]
				}()
				return dst
			}())[0] = (func() []byte {
				defer func() {
					src = src[0+1:]
				}()
				return src
			}())[0]
		}
	}
	dst[0] = '\x00'
	return src
}

// postps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:283
func postps() {
	var cmd []byte = make([]byte, 1000)
	var arg []byte = make([]byte, 1000)
	nextword(cmd)
	readln(arg)
	if noarch.Not(noarch.Strcmp([]byte("PS\x00"), cmd)) || noarch.Not(noarch.Strcmp([]byte("ps\x00"), cmd)) {
		out([]byte("%s\n\x00"), arg)
	}
	if noarch.Not(noarch.Strcmp([]byte("rotate\x00"), cmd)) {
		outrotate(noarch.Atoi(arg))
	}
	if noarch.Not(noarch.Strcmp([]byte("eps\x00"), cmd)) || noarch.Not(noarch.Strcmp([]byte("pdf\x00"), cmd)) {
		var path []byte = make([]byte, 4096)
		var hwid int32
		var vwid int32
		var nspec int32
		var spec []byte = arg
		spec = strcut(path, spec)
		nspec = noarch.Sscanf(spec, []byte("%d %d\x00"), c4goUnsafeConvert_int32(&hwid), c4goUnsafeConvert_int32(&vwid))
		if nspec < 1 {
			hwid = 0
		}
		if nspec < 2 {
			vwid = 0
		}
		if int32(path[0]) != 0 && noarch.Not(noarch.Strcmp([]byte("eps\x00"), cmd)) {
			outeps(path, hwid, vwid)
		}
		if int32(path[0]) != 0 && noarch.Not(noarch.Strcmp([]byte("pdf\x00"), cmd)) {
			outpdf(path, hwid, vwid)
		}
	}
	if noarch.Not(noarch.Strcmp([]byte("name\x00"), cmd)) {
		var spec []byte = arg
		var nspec int32
		if name_n == name_sz {
			if name_sz == 0 {
				name_sz = 128
			} else {
				name_sz = name_sz * 2
			}
			name_desc = mextend(name_desc, name_n, name_sz, int32(64)).([][]byte)
			name_page = mextend(name_page, name_n, name_sz, int32(4)).([]int32)
			name_offset = mextend(name_offset, name_n, name_sz, int32(4)).([]int32)
		}
		spec = strcut(name_desc[name_n], spec)
		nspec = noarch.Sscanf(spec, []byte("%d %d\x00"), name_page[name_n:], name_offset[name_n:])
		if int32(name_desc[name_n][0]) != 0 && nspec > 0 {
			name_n++
		}
	}
	if noarch.Not(noarch.Strcmp([]byte("mark\x00"), cmd)) {
		var spec []byte = arg
		var nspec int32
		if mark_n == mark_sz {
			if mark_sz == 0 {
				mark_sz = 128
			} else {
				mark_sz = mark_sz * 2
			}
			mark_desc = mextend(mark_desc, mark_n, mark_sz, int32(256)).([][]byte)
			mark_page = mextend(mark_page, mark_n, mark_sz, int32(4)).([]int32)
			mark_offset = mextend(mark_offset, mark_n, mark_sz, int32(4)).([]int32)
			mark_level = mextend(mark_level, mark_n, mark_sz, int32(4)).([]int32)
		}
		spec = strcut(mark_desc[mark_n], spec)
		nspec = noarch.Sscanf(spec, []byte("%d %d %d\x00"), mark_page[mark_n:], mark_offset[mark_n:], mark_level[mark_n:])
		if int32(mark_desc[mark_n][0]) != 0 && nspec > 0 {
			mark_n++
		}
	}
	if noarch.Not(noarch.Strcmp([]byte("link\x00"), cmd)) {
		var link []byte = make([]byte, 4096)
		var hwid int32
		var vwid int32
		var nspec int32
		var spec []byte = arg
		spec = strcut(link, spec)
		nspec = noarch.Sscanf(spec, []byte("%d %d\x00"), c4goUnsafeConvert_int32(&hwid), c4goUnsafeConvert_int32(&vwid))
		if int32(link[0]) != 0 && nspec == 2 {
			outlink(link, hwid, vwid)
		}
	}
	if noarch.Not(noarch.Strcmp([]byte("info\x00"), cmd)) {
		var spec []byte = arg
		var kwd []byte = make([]byte, 128)
		var i int32
		for int32(spec[0]) == int32(' ') {
			spec = spec[0+1:]
		}
		for int32(spec[0]) != 0 && int32(spec[0]) != int32(' ') {
			if uint32(i) < 128-1 {
				kwd[func() int32 {
					defer func() {
						i++
					}()
					return i
				}()] = spec[0]
			}
			spec = spec[0+1:]
		}
		kwd[i] = '\x00'
		for int32(spec[0]) == int32(' ') {
			spec = spec[0+1:]
		}
		outinfo(kwd, spec)
	}
	if noarch.Not(noarch.Strcmp([]byte("set\x00"), cmd)) {
		var var_ []byte = make([]byte, 128)
		var val []byte = make([]byte, 128)
		if noarch.Sscanf(arg, []byte("%128s %128s\x00"), var_, val) == 2 {
			outset(var_, val)
		}
	}
	if noarch.Not(noarch.Strcmp([]byte("BeginObject\x00"), cmd)) {
		drawmbeg(arg)
	}
	if noarch.Not(noarch.Strcmp([]byte("EndObject\x00"), cmd)) {
		drawmend(arg)
	}
}

// postdir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:375
// output device directory
var postdir []byte = []byte("./\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// postdev - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:376
// output device name
var postdev []byte = []byte("utf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// postx - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:378
func postx() {
	var cmd []byte = make([]byte, 128)
	var font []byte = make([]byte, 128)
	var pos int32
	nextword(cmd)
	switch int32(cmd[0]) {
	case 'f':
		pos = nextnum()
		nextword(font)
		dev_mnt(pos, font, font)
		outmnt(pos)
	case 'i':
		if dev_open(postdir, postdev) != 0 {
			noarch.Fprintf(noarch.Stderr, []byte("neatpost: cannot open device %s\n\x00"), postdev)
			unix.Exit(1)
		}
		docheader(ps_title, ps_pagewidth, ps_pageheight, ps_linewidth)
	case 'T':
		nextword(postdev)
	case 's':
	case 'X':
		postps()
		break
	}
	nexteol()
}

// postcmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:410
func postcmd(c int32) {
	var cs []byte = make([]byte, 32)
	if int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit)) != 0 {
		outrel((c-int32('0'))*10+next()-int32('0'), 0)
		nextutf8(cs)
		outc(cs)
		return
	}
	switch c {
	case 's':
		outsize(nextnum())
	case 'f':
		outfont(nextnum())
	case 'H':
		outh(nextnum())
	case 'V':
		outv(nextnum())
	case 'h':
		outrel(nextnum(), 0)
	case 'v':
		outrel(0, nextnum())
	case 'c':
		nextutf8(cs)
		outc(cs)
	case 'm':
		nextword(cs)
		outcolor(clr_get(cs))
	case 'N':
		nextnum()
	case 'C':
		nextword(cs)
		outc(cs)
	case 'p':
		if o_pages != 0 {
			docpageend(o_pages)
		}
		o_pages = nextnum()
		docpagebeg(o_pages)
		outpage()
	case 'w':
	case 'n':
		nextnum()
		nextnum()
	case 'D':
		postdraw()
	case 'x':
		postx()
	case '#':
		nexteol()
	default:
		noarch.Fprintf(noarch.Stderr, []byte("neatpost: unknown command %c\n\x00"), c)
		nexteol()
	}
}

// post - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:481
func post() {
	var c int32
	for (func() int32 {
		c = next()
		return c
	}()) >= 0 {
		if noarch.Not(int32(((__ctype_b_loc())[0])[c]) & int32(uint16(noarch.ISspace))) {
			postcmd(c)
		}
	}
	if o_pages != 0 {
		docpageend(o_pages)
	}
	if name_n != 0 {
		outname(name_n, name_desc, name_page, name_offset)
	}
	if mark_n != 0 {
		outmark(mark_n, mark_desc, mark_page, mark_offset, mark_level)
	}
}

// paper - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:495
type paper struct {
	name []byte
	w    int32
	h    int32
}

// papers - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:495
var papers []paper = []paper{{[]byte("letter\x00"), 2159, 2794}, {[]byte("legal\x00"), 2159, 3556}, {[]byte("ledger\x00"), 4318, 2794}, {[]byte("tabloid\x00"), 2794, 4318}}

// setpagesize - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:505
func setpagesize(s []byte) {
	var d1 int32
	var d2 int32
	var n int32
	var i int32
	{
		// predefined paper sizes
		for i = 0; uint32(i) < 96/24; i++ {
			if noarch.Not(noarch.Strcmp(papers[i].name, s)) {
				ps_pagewidth = papers[i].w
				ps_pageheight = papers[i].h
				return
			}
		}
	}
	if int32(((__ctype_b_loc())[0])[int32(s[0])])&int32(uint16(noarch.ISdigit)) != 0 && noarch.Strchr(s, int32('x')) != nil {
		// custom paper size in tenth of mm; example: 2100x2970 for a4
		ps_pagewidth = noarch.Atoi(s)
		ps_pageheight = noarch.Atoi((noarch.Strchr(s, int32('x')))[0+1:])
		return
	}
	if noarch.Strchr([]byte("abcABC\x00"), int32(s[0])) == nil || noarch.Not(int32(((__ctype_b_loc())[0])[int32(s[1])])&int32(uint16(noarch.ISdigit))) {
		// ISO paper sizes
		return
	}
	if tolower(int32(s[0])) == int32('a') {
		d1 = 8410
		d2 = 11890
	}
	if tolower(int32(s[0])) == int32('b') {
		d1 = 10000
		d2 = 14140
	}
	if tolower(int32(s[0])) == int32('c') {
		d1 = 9170
		d2 = 12970
	}
	n = int32(s[1]) - int32('0')
	ps_pagewidth = func() int32 {
		if n&1 != 0 {
			return d2
		}
		return d1
	}() >> uint64((n+1)>>uint64(1))
	ps_pageheight = func() int32 {
		if n&1 != 0 {
			return d1
		}
		return d2
	}() >> uint64(n>>uint64(1))
	ps_pagewidth -= ps_pagewidth % 10
	ps_pageheight -= ps_pageheight % 10
}

// mextend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:544
func mextend(old interface{}, oldsz int32, newsz int32, memsz int32) interface{} {
	var new_ interface{} = make([]byte, uint32(newsz*memsz))
	memcpy(new_, old, uint32(oldsz*memsz))
	noarch.Memset(new_[0+oldsz*memsz:].([]byte), byte(0), uint32((newsz-oldsz)*memsz))
	_ = old
	return new_
}

// utf8code - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:554
func utf8code(s []byte) int32 {
	// the unicode codepoint of the given utf-8 character
	var c int32 = int32(uint8(s[0]))
	if ^c&192 != 0 {
		// ASCII or invalid
		return c
	}
	if ^c&32 != 0 {
		return c&31<<uint64(6) | int32(s[1])&63
	}
	if ^c&16 != 0 {
		return c&15<<uint64(12) | int32(s[1])&63<<uint64(6) | int32(s[2])&63
	}
	if ^c&8 != 0 {
		return c&7<<uint64(18) | int32(s[1])&63<<uint64(12) | int32(s[2])&63<<uint64(6) | int32(s[3])&63
	}
	return c
}

// pdftext_ascii - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:568
func pdftext_ascii(s []byte) int32 {
	for ; s[0] != 0; func() []byte {
		tempVarUnary := s
		defer func() {
			s = s[0+1:]
		}()
		return tempVarUnary
	}() {
		if int32(uint8(s[0]))&128 != 0 || int32(s[0]) == int32('(') || int32(s[0]) == int32(')') {
			return 0
		}
	}
	return 1
}

// pdftext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:577
func pdftext(s []byte) []byte {
	// encode s as pdf text string
	var sb []sbuf = sbuf_make()
	if pdftext_ascii(s) != 0 {
		sbuf_chr(sb, int32('('))
		sbuf_str(sb, s)
		sbuf_chr(sb, int32(')'))
		return sbuf_done(sb)
	}
	// read utf-8 and write utf-16
	// unicode byte order marker
	sbuf_str(sb, []byte("<FEFF\x00"))
	for s[0] != 0 {
		var c int32 = utf8code(s)
		if c >= 0 && c <= 55295 || c >= 57344 && c <= 65535 {
			sbuf_printf(sb, []byte("%02X%02X\x00"), c>>uint64(8), c&255)
		}
		if c >= 65536 && c <= 1114111 {
			var c1 int32 = 55296 + (c-65536)>>uint64(10)
			var c2 int32 = 56320 + (c-65536)&1023
			sbuf_printf(sb, []byte("%02X%02X\x00"), c1>>uint64(8), c1&255)
			sbuf_printf(sb, []byte("%02X%02X\x00"), c2>>uint64(8), c2&255)
		}
		s = s[0+utf8len(int32(uint8(s[0]))):]
	}
	sbuf_chr(sb, int32('>'))
	return sbuf_done(sb)
}

// pdftext_static - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:606
func pdftext_static(s []byte) []byte {
	// encode s as pdf text string; returns a static buffer
	var buf []byte = make([]byte, 1024)
	var r []byte = pdftext(s)
	noarch.Snprintf(buf, int32(1024), []byte("%s\x00"), r)
	_ = r
	return buf
}

// usage - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:615
var usage []byte = []byte("Usage: neatpost [options] <input >output\nOptions:\n  -F dir  \tset font directory (./)\n  -p size \tset paper size (letter); e.g., a4, 2100x2970\n  -t title\tspecify document title\n  -w lwid \tdrawing line thickness in thousandths of an em (40)\n  -l      \tlandscape mode\n  -n      \talways draw glyphs by name (ps glyphshow)\n\x00")

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/post.c:625
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	var i int32
	var landscape int32
	for i = 1; i < argc; i++ {
		if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('F') {
			noarch.Strcpy(postdir, func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}())
		} else if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('p') {
			setpagesize(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}())
		} else if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('w') {
			ps_linewidth = noarch.Atoi(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}())
		} else if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('n') {
			outgname(1)
		} else if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('t') {
			if int32(argv[i][2]) != 0 {
				ps_title = (argv[i])[0+2:]
			} else {
				ps_title = argv[func() int32 {
					i++
					return i
				}()]
			}
		} else if int32(argv[i][0]) == int32('-') && int32(argv[i][1]) == int32('l') {
			landscape = 1
		} else {
			noarch.Printf([]byte("%s\x00"), usage)
			return
		}
	}
	if landscape != 0 {
		var t int32 = ps_pagewidth
		ps_pagewidth = ps_pageheight
		ps_pageheight = t
	}
	post()
	doctrailer(o_pages)
	dev_close()
	_ = mark_desc
	_ = mark_page
	_ = mark_offset
	_ = mark_level
	_ = name_desc
	_ = name_page
	_ = name_offset
	return
}

// pdf_title - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:11
// PDF post-processor functions
// document title
var pdf_title []byte = make([]byte, 256)

// pdf_author - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:12
// document author
var pdf_author []byte = make([]byte, 256)

// pdf_width - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:13
// page width
var pdf_width int32

// pdf_height - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:14
// page height
var pdf_height int32

// pdf_linewid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:15
// line width in thousands of ems
var pdf_linewid int32

// pdf_linecap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:16
// line cap style: 0 (butt), 1 (round), 2 (projecting square)
var pdf_linecap int32 = 1

// pdf_linejoin - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:17
// line join style: 0 (miter), 1 (round), 2 (bevel)
var pdf_linejoin int32 = 1

// pdf_pages - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:18
// pages object id
var pdf_pages int32

// pdf_root - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:19
// root object id
var pdf_root int32

// pdf_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:20
// current pdf file offset
var pdf_pos int32

// obj_off - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:21
// object offsets
var obj_off []int32

// obj_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:22
// number of pdf objects
var obj_sz int32

// obj_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:22
var obj_n int32

// page_id - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:23
// page object ids
var page_id []int32

// page_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:24
// number of pages
var page_sz int32

// page_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:24
var page_n int32

// pdf_outline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:25
// pdf outline hierarchiy
var pdf_outline int32

// pdf_dests - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:26
// named destinations
var pdf_dests int32

// pg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:28
// current page contents
var pg []sbuf

// o_f - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:29
// font and size
var o_f int32

// o_s - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:29
var o_s int32

// o_m - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:29
var o_m int32

// o_h - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:30
// current user position
var o_h int32

// o_v - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:30
var o_v int32

// p_h - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:31
// current output position
var p_h int32

// p_v - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:31
var p_v int32

// o_i - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:32
// output and pdf fonts (indices into pfont[])
var o_i int32

// p_i - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:32
var p_i int32

// p_f - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:33
// output font
var p_f int32

// p_s - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:33
var p_s int32

// p_m - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:33
var p_m int32

// o_queued - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:34
// queued character type
var o_queued int32

// o_iset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:35
// fonts accesssed in this page
var o_iset []byte = make([]byte, 1024)

// xobj - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:36
// page xobject object ids
var xobj []int32 = make([]int32, 128)

// xobj_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:37
// number of xobjects in this page
var xobj_n int32

// ann - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:38
// page annotations
var ann []int32 = make([]int32, 128)

// ann_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:39
// number of annotations in this page
var ann_n int32

// pfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:42
// loaded PDF fonts
type pfont struct {
	name [128]byte
	path [1024]byte
	desc [1024]byte
	gbeg int32
	gend int32
	sub  int32
	obj  int32
	des  int32
	cid  int32
}

// pfonts - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:54
// font PostScript name
// font path
// font descriptor path
// the first glyph
// the last glyph
// subfont number
// the font object
// font descriptor
// CID-indexed
var pfonts []pfont

// pfonts_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:55
var pfonts_n int32

// pfonts_sz - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:55
var pfonts_sz int32

// pdfout - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:58
func pdfout(s []byte, c4goArgs ...interface{}) {
	// print formatted pdf output
	var ap *va_list
	va_start(ap, s)
	pdf_pos += noarch.Vprintf(s, ap)
	va_end(ap)
}

// pdfmem - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:67
func pdfmem(s []byte, len_ int32) {
	// print pdf output
	noarch.Fwrite(s, len_, 1, noarch.Stdout)
	pdf_pos += len_
}

// obj_map - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:74
func obj_map() int32 {
	if obj_n == obj_sz {
		// allocate an object number
		obj_sz += 1024
		obj_off = mextend(obj_off, obj_n, obj_sz, int32(4)).([]int32)
	}
	defer func() {
		obj_n++
	}()
	return obj_n
}

// obj_beg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:84
func obj_beg(id int32) int32 {
	if id <= 0 {
		// start the definition of an object
		id = obj_map()
	}
	obj_off[id] = pdf_pos
	pdfout([]byte("%d 0 obj\n\x00"), id)
	return id
}

// obj_end - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:94
func obj_end() {
	// end an object definition
	pdfout([]byte("endobj\n\n\x00"))
}

// out - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:99
func out(s []byte, c4goArgs ...interface{}) {
}

// type1lengths - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:104
func type1lengths(t1 []byte, l int32, l1 []int32, l2 []int32, l3 []int32) int32 {
	// the length of the clear-text, encrypted, and fixed-content portions
	var i int32
	var cleartext []byte = t1
	var encrypted []byte
	var fixedcont []byte
	for i = 0; i < l-5 && encrypted == nil; i++ {
		if int32(t1[i]) == int32('e') && noarch.Not(noarch.Memcmp([]byte("eexec\x00"), t1[0+i:], 5)) {
			encrypted = t1[0+i:]
		}
	}
	if encrypted == nil {
		return 1
	}
	for ; i < l-512 && fixedcont == nil; i++ {
		if int32(t1[i]) == int32('0') && noarch.Not(noarch.Memcmp([]byte("00000\x00"), t1[0+i:], 5)) {
			fixedcont = t1[0+i:]
		}
	}
	l1[0] = int32((int64(uintptr(unsafe.Pointer(&encrypted[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&cleartext[0])))/int64(1)))
	if fixedcont != nil {
		l2[0] = int32((int64(uintptr(unsafe.Pointer(&fixedcont[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&cleartext[0])))/int64(1)))
	} else {
		l2[0] = 0
	}
	if fixedcont != nil {
		l3[0] = int32((int64(uintptr(unsafe.Pointer(&t1[0+l])))/int64(1) - int64(uintptr(unsafe.Pointer(&fixedcont[0])))/int64(1)))
	} else {
		l3[0] = 0
	}
	return 0
}

// fonttype - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:125
func fonttype(path []byte) int32 {
	// return font type: 't': TrueType, '1': Type 1, 'o': OpenType
	var ext []byte = noarch.Strrchr(path, int32('.'))
	if ext != nil && noarch.Not(noarch.Strcmp([]byte(".ttf\x00"), ext)) {
		return int32('t')
	}
	if ext != nil && noarch.Not(noarch.Strcmp([]byte(".otf\x00"), ext)) {
		return int32('t')
	}
	if ext != nil && (noarch.Not(noarch.Strcmp([]byte(".ttc\x00"), ext)) || noarch.Not(noarch.Strcmp([]byte(".otc\x00"), ext))) {
		return int32('t')
	}
	return int32('1')
}

// pfont_write - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:138
func pfont_write(ps []pfont) {
	// write the object corresponding to the given font
	var i int32
	var enc_obj int32
	var fn []font = dev_fontopen(ps[0].desc[:])
	// the encoding object
	enc_obj = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Encoding\n\x00"))
	pdfout([]byte("  /Differences [ %d\x00"), ps[0].gbeg%256)
	for i = ps[0].gbeg; i <= ps[0].gend; i++ {
		pdfout([]byte(" /%s\x00"), font_glget(fn, i)[0].id[:])
	}
	pdfout([]byte(" ]\n\x00"))
	pdfout([]byte(">>\n\x00"))
	obj_end()
	// the font object
	obj_beg(ps[0].obj)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Font\n\x00"))
	if fonttype(ps[0].path[:]) == int32('t') {
		pdfout([]byte("  /Subtype /TrueType\n\x00"))
	} else {
		pdfout([]byte("  /Subtype /Type1\n\x00"))
	}
	pdfout([]byte("  /BaseFont /%s\n\x00"), ps[0].name[:])
	pdfout([]byte("  /FirstChar %d\n\x00"), ps[0].gbeg%256)
	pdfout([]byte("  /LastChar %d\n\x00"), ps[0].gend%256)
	pdfout([]byte("  /Widths [\x00"))
	for i = ps[0].gbeg; i <= ps[0].gend; i++ {
		pdfout([]byte(" %d\x00"), font_glget(fn, i)[0].wid*100*72/dev_res)
	}
	pdfout([]byte(" ]\n\x00"))
	pdfout([]byte("  /FontDescriptor %d 0 R\n\x00"), ps[0].des)
	pdfout([]byte("  /Encoding %d 0 R\n\x00"), enc_obj)
	pdfout([]byte(">>\n\x00"))
	obj_end()
	font_close(fn)
}

// encodehex - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:175
func encodehex(d []sbuf, s []byte, n int32) {
	var hex []byte = []byte("0123456789ABCDEF\x00")
	var i int32
	for i = 0; i < n; i++ {
		sbuf_chr(d, int32(hex[int32(uint8(s[i]))>>uint64(4)]))
		sbuf_chr(d, int32(hex[int32(uint8(s[i]))&15]))
		if i%40 == 39 && i+1 < n {
			sbuf_chr(d, int32('\n'))
		}
	}
	sbuf_str(d, []byte(">\n\x00"))
}

// pfont_writecid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:189
func pfont_writecid(ps []pfont) {
	// write the object corresponding to this CID font
	var cid_obj int32
	var fn []font = dev_fontopen(ps[0].desc[:])
	var gcnt int32
	var i int32
	// CIDFont
	cid_obj = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Font\n\x00"))
	pdfout([]byte("  /Subtype /CIDFontType2\n\x00"))
	pdfout([]byte("  /BaseFont /%s\n\x00"), ps[0].name[:])
	pdfout([]byte("  /CIDSystemInfo <</Ordering(Identity)/Registry(Adobe)/Supplement 0>>\n\x00"))
	pdfout([]byte("  /FontDescriptor %d 0 R\n\x00"), ps[0].des)
	pdfout([]byte("  /DW 1000\n\x00"))
	for font_glget(fn, gcnt) != nil {
		gcnt++
	}
	pdfout([]byte("  /W [ %d [\x00"), ps[0].gbeg)
	for i = ps[0].gbeg; i <= ps[0].gend; i++ {
		pdfout([]byte(" %d\x00"), font_glget(fn, i)[0].wid*100*72/dev_res)
	}
	pdfout([]byte(" ] ]\n\x00"))
	pdfout([]byte(">>\n\x00"))
	obj_end()
	// the font object
	obj_beg(ps[0].obj)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Font\n\x00"))
	pdfout([]byte("  /Subtype /Type0\n\x00"))
	pdfout([]byte("  /BaseFont /%s\n\x00"), ps[0].name[:])
	pdfout([]byte("  /Encoding /Identity-H\n\x00"))
	pdfout([]byte("  /DescendantFonts [%d 0 R]\n\x00"), cid_obj)
	pdfout([]byte(">>\n\x00"))
	obj_end()
	font_close(fn)
}

// writedesc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:226
func writedesc(fn []font) int32 {
	// write font descriptor; returns its object ID
	var str_obj int32 = -1
	var des_obj int32
	var buf []byte = make([]byte, 1024)
	var fntype int32 = fonttype(font_path(fn))
	if fntype == int32('1') || fntype == int32('t') {
		var fd int32 = noarch.Open(font_path(fn), 0)
		var ffsb []sbuf = sbuf_make()
		var sb []sbuf = sbuf_make()
		var l1 int32
		var l2 int32
		var l3 int32
		var nr int32
		for (func() int32 {
			nr = int32(noarch.Read(fd, buf, 1024))
			return nr
		}()) > 0 {
			// reading the font file
			sbuf_mem(ffsb, buf, nr)
		}
		noarch.CloseOnExec(fd)
		l1 = sbuf_len(ffsb)
		if fntype == int32('1') {
			if type1lengths(sbuf_buf(ffsb), sbuf_len(ffsb), c4goUnsafeConvert_int32(&l1), c4goUnsafeConvert_int32(&l2), c4goUnsafeConvert_int32(&l3)) != 0 {
				// initialize Type 1 lengths
				l1 = 0
			}
			if l3 != 0 {
				// remove the fixed-content portion of the font
				sbuf_cut(ffsb, l1+l2)
			}
			l1 -= l3
			l3 = 0
		}
		// encoding file contents
		encodehex(sb, sbuf_buf(ffsb), sbuf_len(ffsb))
		if l1 != 0 {
			// write font data if it has nonzero length
			str_obj = obj_beg(0)
			pdfout([]byte("<<\n\x00"))
			pdfout([]byte("  /Filter /ASCIIHexDecode\n\x00"))
			pdfout([]byte("  /Length %d\n\x00"), sbuf_len(sb))
			pdfout([]byte("  /Length1 %d\n\x00"), l1)
			if fntype == int32('1') {
				pdfout([]byte("  /Length2 %d\n\x00"), l2)
			}
			if fntype == int32('1') {
				pdfout([]byte("  /Length3 %d\n\x00"), l3)
			}
			pdfout([]byte(">>\n\x00"))
			pdfout([]byte("stream\n\x00"))
			pdfmem(sbuf_buf(sb), sbuf_len(sb))
			pdfout([]byte("endstream\n\x00"))
			obj_end()
		}
		sbuf_free(ffsb)
		sbuf_free(sb)
	}
	// the font descriptor
	des_obj = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /FontDescriptor\n\x00"))
	pdfout([]byte("  /FontName /%s\n\x00"), font_name(fn))
	pdfout([]byte("  /Flags 32\n\x00"))
	pdfout([]byte("  /FontBBox [-1000 -1000 1000 1000]\n\x00"))
	pdfout([]byte("  /MissingWidth 1000\n\x00"))
	pdfout([]byte("  /StemV 100\n\x00"))
	pdfout([]byte("  /ItalicAngle 0\n\x00"))
	pdfout([]byte("  /CapHeight 100\n\x00"))
	pdfout([]byte("  /Ascent 100\n\x00"))
	pdfout([]byte("  /Descent 100\n\x00"))
	if str_obj >= 0 {
		pdfout([]byte("  /FontFile%s %d 0 R\n\x00"), func() []byte {
			if fntype == int32('t') {
				return []byte("2\x00")
			}
			return []byte("\x00")
		}(), str_obj)
	}
	pdfout([]byte(">>\n\x00"))
	obj_end()
	return des_obj
}

// pfont_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:297
func pfont_find(g []glyph) int32 {
	var fn []font = g[0].font
	var name []byte = font_name(fn)
	var ps []pfont
	var fntype int32 = fonttype(font_path(fn))
	var sub int32 = func() int32 {
		if fntype == int32('1') {
			return font_glnum(fn, g) / 256
		}
		return 0
	}()
	var i int32
	for i = 0; i < pfonts_n; i++ {
		if noarch.Not(noarch.Strcmp(name, pfonts[i].name[:])) && pfonts[i].sub == sub {
			return i
		}
	}
	if pfonts_n == pfonts_sz {
		pfonts_sz += 16
		pfonts = mextend(pfonts, pfonts_n, pfonts_sz, int32(2224)).([]pfont)
	}
	ps = pfonts[pfonts_n:]
	noarch.Snprintf(ps[0].name[:], int32(128), []byte("%s\x00"), name)
	noarch.Snprintf(ps[0].path[:], int32(1024), []byte("%s\x00"), font_path(fn))
	noarch.Snprintf(ps[0].desc[:], int32(1024), []byte("%s\x00"), font_desc(fn))
	ps[0].cid = noarch.BoolToInt(fntype == int32('t'))
	ps[0].obj = obj_map()
	ps[0].sub = sub
	ps[0].gbeg = 1 << uint64(20)
	for i = 0; i < pfonts_n; i++ {
		if noarch.Not(noarch.Strcmp(pfonts[i].name[:], ps[0].name[:])) {
			break
		}
	}
	if i < pfonts_n {
		ps[0].des = pfonts[i].des
	} else {
		ps[0].des = writedesc(fn)
	}
	defer func() {
		pfonts_n++
	}()
	return pfonts_n
}

// pfont_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:331
func pfont_done() {
	var i int32
	for i = 0; i < pfonts_n; i++ {
		if pfonts[i].cid != 0 {
			pfont_writecid(pfonts[i:])
		} else {
			pfont_write(pfonts[i:])
		}
	}
	_ = pfonts
}

// o_flush - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:343
func o_flush() {
	if o_queued == 1 {
		sbuf_printf(pg, []byte(">] TJ\n\x00"))
	}
	o_queued = 0
}

// o_loadfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:350
func o_loadfont(g []glyph) int32 {
	var fn int32 = pfont_find(g)
	o_iset[fn] = byte(1)
	return fn
}

// pdfpos00 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:358
func pdfpos00(uh int32, uv int32) []byte {
	// like pdfpos() but assume that uh and uv are multiplied by 100
	var buf []byte = make([]byte, 64)
	var h int32 = uh * 72 / dev_res
	var v int32 = pdf_height*100 - uv*72/dev_res
	noarch.Sprintf(buf, []byte("%s%d.%02d %s%d.%02d\x00"), func() []byte {
		if h < 0 {
			return []byte("-\x00")
		}
		return []byte("\x00")
	}(), noarch.Abs(h)/100, noarch.Abs(h)%100, func() []byte {
		if v < 0 {
			return []byte("-\x00")
		}
		return []byte("\x00")
	}(), noarch.Abs(v)/100, noarch.Abs(v)%100)
	return buf
}

// pdfpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:370
func pdfpos(uh int32, uv int32) []byte {
	// convert troff position to pdf position; returns a static buffer
	return pdfpos00(uh*100, uv*100)
}

// pdfunit - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:376
func pdfunit(uh int32, sz int32) []byte {
	// troff length to thousands of a unit of text space; returns a static buffer
	var buf []byte = make([]byte, 64)
	var h int32 = uh * 1000 * 72 / sz / dev_res
	noarch.Sprintf(buf, []byte("%s%d\x00"), func() []byte {
		if h < 0 {
			return []byte("-\x00")
		}
		return []byte("\x00")
	}(), noarch.Abs(h))
	return buf
}

// pdfcolor - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:385
func pdfcolor(m int32) []byte {
	// convert troff color to pdf color; returns a static buffer
	var buf []byte = make([]byte, 64)
	var r int32 = m >> uint64(16) & 255 * 1000 / 255
	var g int32 = m >> uint64(8) & 255 * 1000 / 255
	var b int32 = m & 255 * 1000 / 255
	sbuf_printf(pg, []byte("%d.%03d %d.%03d %d.%03d\x00"), r/1000, r%1000, g/1000, g%1000, b/1000, b%1000)
	return buf
}

// o_queue - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:396
func o_queue(g []glyph) {
	var gid int32
	if o_v != p_v {
		o_flush()
		sbuf_printf(pg, []byte("1 0 0 1 %s Tm\n\x00"), pdfpos(o_h, o_v))
		p_h = o_h
		p_v = o_v
	}
	if noarch.Not(o_queued) {
		sbuf_printf(pg, []byte("[<\x00"))
	}
	o_queued = 1
	if o_h != p_h {
		sbuf_printf(pg, []byte("> %s <\x00"), pdfunit(p_h-o_h, o_s))
	}
	// printing glyph identifier
	gid = font_glnum(g[0].font, g)
	if pfonts[o_i].cid != 0 {
		sbuf_printf(pg, []byte("%04x\x00"), gid)
	} else {
		sbuf_printf(pg, []byte("%02x\x00"), gid%256)
	}
	if gid < pfonts[o_i].gbeg {
		// updating gbeg and gend
		pfonts[o_i].gbeg = gid
	}
	if gid > pfonts[o_i].gend {
		pfonts[o_i].gend = gid
	}
	// advancing
	p_h = o_h + font_wid(g[0].font, o_s, g[0].wid)
}

// out_fontup - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:425
func out_fontup() {
	if o_m != p_m {
		o_flush()
		sbuf_printf(pg, []byte("%s rg\n\x00"), pdfcolor(o_m))
		p_m = o_m
	}
	if o_i >= 0 && (o_i != p_i || o_s != p_s) {
		var ps []pfont = pfonts[o_i:]
		o_flush()
		if ps[0].cid != 0 {
			sbuf_printf(pg, []byte("/%s %d Tf\n\x00"), ps[0].name[:], o_s)
		} else {
			sbuf_printf(pg, []byte("/%s.%d %d Tf\n\x00"), ps[0].name[:], ps[0].sub, o_s)
		}
		p_i = o_i
		p_s = o_s
	}
}

// outc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:444
func outc(c []byte) {
	var g []glyph
	var fn []font
	g = dev_glyph(c, o_f)
	if g != nil {
		fn = g[0].font
	} else {
		fn = dev_font(o_f)
	}
	if g == nil {
		outrel(func() int32 {
			if int32(c[0]) == int32(' ') && fn != nil {
				return font_swid(fn, o_s)
			}
			return 1
		}(), 0)
		return
	}
	o_i = o_loadfont(g)
	out_fontup()
	o_queue(g)
}

// outh - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:459
func outh(h int32) {
	o_h = h
}

// outv - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:464
func outv(v int32) {
	o_v = v
}

// outrel - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:469
func outrel(h int32, v int32) {
	o_h += h
	o_v += v
}

// outfont - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:475
func outfont(f int32) {
	if dev_font(f) != nil {
		o_f = f
	}
}

// outsize - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:481
func outsize(s int32) {
	if s > 0 {
		o_s = s
	}
}

// outcolor - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:487
func outcolor(c int32) {
	o_m = c
}

// outrotate - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:492
func outrotate(deg int32) {
}

// outeps - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:496
func outeps(eps []byte, hwid int32, vwid int32) {
}

// pdf_copy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:501
func pdf_copy(pdf []byte, len_ int32, pos int32) []byte {
	// return a copy of a PDF object; returns a static buffer
	var buf []byte = make([]byte, 4096)
	var datlen int32
	pos += pdf_ws(pdf, len_, pos)
	datlen = pdf_len(pdf, len_, pos)
	if uint32(datlen) > 4096-1 {
		datlen = int32(4096 - 1)
	}
	memcpy(buf, pdf[0+pos:], uint32(datlen))
	buf[datlen] = '\x00'
	return buf
}

// pdf_strcopy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:517
func pdf_strcopy(pdf []byte, len_ int32, pos int32, sb []sbuf) int32 {
	// write stream to sb
	var slen int32
	var val int32
	var beg int32
	if (func() int32 {
		val = pdf_dval_val(pdf, len_, pos, []byte("/Length\x00"))
		return val
	}()) < 0 {
		return -1
	}
	slen = noarch.Atoi(pdf[0+val:])
	pos = pos + pdf_len(pdf, len_, pos)
	pos += pdf_ws(pdf, len_, pos)
	if pos+slen+15 > len_ {
		return -1
	}
	beg = pos
	pos += noarch.Strlen([]byte("stream\x00"))
	if int32(pdf[pos]) == int32('\r') {
		pos++
	}
	pos += 1 + slen
	if int32(pdf[pos]) == int32('\r') || int32(pdf[pos]) == int32(' ') {
		pos++
	}
	if int32(pdf[pos]) == int32('\n') {
		pos++
	}
	pos += noarch.Strlen([]byte("endstream\x00")) + int32(1)
	sbuf_mem(sb, pdf[0+beg:], pos-beg)
	return 0
}

// pdf_objcopy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:543
func pdf_objcopy(pdf []byte, len_ int32, pos int32) int32 {
	// copy a PDF object and return its new identifier
	var id int32
	if (func() int32 {
		pos = pdf_ref(pdf, len_, pos)
		return pos
	}()) < 0 {
		return -1
	}
	if pdf_type(pdf, len_, pos) == int32('d') {
		var sb []sbuf = sbuf_make()
		pdf_dictcopy(pdf, len_, pos, sb)
		sbuf_chr(sb, int32('\n'))
		if pdf_dval(pdf, len_, pos, []byte("/Length\x00")) >= 0 {
			pdf_strcopy(pdf, len_, pos, sb)
		}
		id = obj_beg(0)
		pdfmem(sbuf_buf(sb), sbuf_len(sb))
		obj_end()
		sbuf_free(sb)
	} else {
		id = obj_beg(0)
		pdfmem(pdf[0+pos:], pdf_len(pdf, len_, pos))
		pdfout([]byte("\n\x00"))
		obj_end()
	}
	return id
}

// pdf_dictcopy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:568
func pdf_dictcopy(pdf []byte, len_ int32, pos int32, sb []sbuf) {
	// copy a PDF dictionary recursively
	var i int32
	var key int32
	var val int32
	var id int32
	sbuf_printf(sb, []byte("<<\x00"))
	for i = 0; ; i++ {
		if (func() int32 {
			key = pdf_dkey(pdf, len_, pos, i)
			return key
		}()) < 0 {
			break
		}
		sbuf_printf(sb, []byte(" %s\x00"), pdf_copy(pdf, len_, key))
		val = pdf_dval(pdf, len_, pos, pdf_copy(pdf, len_, key))
		if pdf_type(pdf, len_, val) == int32('r') {
			if (func() int32 {
				id = pdf_objcopy(pdf, len_, val)
				return id
			}()) >= 0 {
				sbuf_printf(sb, []byte(" %d 0 R\x00"), id)
			}
		} else {
			sbuf_printf(sb, []byte(" %s\x00"), pdf_copy(pdf, len_, val))
		}
	}
	sbuf_printf(sb, []byte(" >>\x00"))
}

// pdf_rescopy - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:589
func pdf_rescopy(pdf []byte, len_ int32, pos int32, sb []sbuf) {
	// copy resources dictionary
	var res_fields [][]byte = [][]byte{[]byte("/ProcSet\x00"), []byte("/ExtGState\x00"), []byte("/ColorSpace\x00"), []byte("/Pattern\x00"), []byte("/Shading\x00"), []byte("/Properties\x00"), []byte("/Font\x00"), []byte("/XObject\x00")}
	var res int32
	var i int32
	sbuf_printf(sb, []byte("  /Resources <<\n\x00"))
	for i = 0; uint32(i) < 64/8; i++ {
		if (func() int32 {
			res = pdf_dval_val(pdf, len_, pos, res_fields[i])
			return res
		}()) >= 0 {
			if pdf_type(pdf, len_, res) == int32('d') {
				sbuf_printf(sb, []byte("    %s \x00"), res_fields[i])
				pdf_dictcopy(pdf, len_, res, sb)
				sbuf_printf(sb, []byte("\n\x00"))
			} else {
				sbuf_printf(sb, []byte("    %s %s\n\x00"), res_fields[i], pdf_copy(pdf, len_, res))
			}
		}
	}
	sbuf_printf(sb, []byte("  >>\n\x00"))
}

// pdfbbox100 - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:610
func pdfbbox100(pdf []byte, len_ int32, pos int32, dim []int32) int32 {
	var val int32
	var i int32
	for i = 0; i < 4; i++ {
		var n int32
		var f1 int32
		var f2 int32
		if (func() int32 {
			val = pdf_lval(pdf, len_, pos, i)
			return val
		}()) < 0 {
			return -1
		}
		for ; int32(((__ctype_b_loc())[0])[int32(uint8(pdf[val]))])&int32(uint16(noarch.ISdigit)) != 0; val++ {
			n = n*10 + int32(pdf[val]) - int32('0')
		}
		if int32(pdf[val]) == int32('.') {
			if int32(((__ctype_b_loc())[0])[int32(uint8(pdf[val+1]))])&int32(uint16(noarch.ISdigit)) != 0 {
				f1 = int32(pdf[val+1]) - int32('0')
				if int32(((__ctype_b_loc())[0])[int32(uint8(pdf[val+2]))])&int32(uint16(noarch.ISdigit)) != 0 {
					f2 = int32(pdf[val+2]) - int32('0')
				}
			}
		}
		dim[i] = n*100 + f1*10 + f2
	}
	return 0
}

// pdfext - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:632
func pdfext(pdf []byte, len_ int32, hwid int32, vwid int32) int32 {
	var cont_fields [][]byte = [][]byte{[]byte("/Filter\x00"), []byte("/DecodeParms\x00")}
	var trailer int32
	var root int32
	var cont int32
	var pages int32
	var page1 int32
	var res int32
	var kids_val int32
	var page1_val int32
	var val int32
	var bbox int32
	var xobj_id int32
	var length int32
	var dim []int32 = make([]int32, 4)
	var hzoom int32 = 100
	var vzoom int32 = 100
	var sb []sbuf
	var i int32
	if uint32(xobj_n) == 512/4 {
		return -1
	}
	if (func() int32 {
		trailer = pdf_trailer(pdf, len_)
		return trailer
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		root = pdf_dval_obj(pdf, len_, trailer, []byte("/Root\x00"))
		return root
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		pages = pdf_dval_obj(pdf, len_, root, []byte("/Pages\x00"))
		return pages
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		kids_val = pdf_dval_val(pdf, len_, pages, []byte("/Kids\x00"))
		return kids_val
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		page1_val = pdf_lval(pdf, len_, kids_val, 0)
		return page1_val
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		page1 = pdf_ref(pdf, len_, page1_val)
		return page1
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		cont = pdf_dval_obj(pdf, len_, page1, []byte("/Contents\x00"))
		return cont
	}()) < 0 {
		return -1
	}
	if (func() int32 {
		val = pdf_dval_val(pdf, len_, cont, []byte("/Length\x00"))
		return val
	}()) < 0 {
		return -1
	}
	res = pdf_dval_val(pdf, len_, page1, []byte("/Resources\x00"))
	length = noarch.Atoi(pdf[0+val:])
	bbox = pdf_dval_val(pdf, len_, page1, []byte("/MediaBox\x00"))
	if bbox < 0 {
		bbox = pdf_dval_val(pdf, len_, pages, []byte("/MediaBox\x00"))
	}
	if bbox >= 0 && noarch.Not(pdfbbox100(pdf, len_, bbox, dim)) {
		if hwid > 0 {
			hzoom = hwid * (100 * 7200 / dev_res) / (dim[2] - dim[0])
		}
		if vwid > 0 {
			vzoom = vwid * (100 * 7200 / dev_res) / (dim[3] - dim[1])
		}
		if vwid <= 0 {
			vzoom = hzoom
		}
		if hwid <= 0 {
			hzoom = vzoom
		}
	}
	sb = sbuf_make()
	sbuf_printf(sb, []byte("<<\n\x00"))
	sbuf_printf(sb, []byte("  /Type /XObject\n\x00"))
	sbuf_printf(sb, []byte("  /Subtype /Form\n\x00"))
	sbuf_printf(sb, []byte("  /FormType 1\n\x00"))
	if bbox >= 0 {
		sbuf_printf(sb, []byte("  /BBox %s\n\x00"), pdf_copy(pdf, len_, bbox))
	}
	sbuf_printf(sb, []byte("  /Matrix [%d.%02d 0 0 %d.%02d %s]\n\x00"), hzoom/100, hzoom%100, vzoom/100, vzoom%100, pdfpos(o_h, o_v))
	if res >= 0 {
		pdf_rescopy(pdf, len_, res, sb)
	}
	sbuf_printf(sb, []byte("  /Length %d\n\x00"), length)
	for i = 0; uint32(i) < 16/8; i++ {
		if (func() int32 {
			val = pdf_dval_val(pdf, len_, cont, cont_fields[i])
			return val
		}()) >= 0 {
			sbuf_printf(sb, []byte("  %s %s\n\x00"), cont_fields[i], pdf_copy(pdf, len_, val))
		}
	}
	sbuf_printf(sb, []byte(">>\n\x00"))
	pdf_strcopy(pdf, len_, cont, sb)
	xobj_id = obj_beg(0)
	pdfmem(sbuf_buf(sb), sbuf_len(sb))
	obj_end()
	sbuf_free(sb)
	xobj[func() int32 {
		defer func() {
			xobj_n++
		}()
		return xobj_n
	}()] = xobj_id
	return xobj_n - 1
}

// outpdf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:702
func outpdf(pdf []byte, hwid int32, vwid int32) {
	var buf []byte = make([]byte, 4096)
	var sb []sbuf
	var xobj_id int32
	var fd int32
	var nr int32
	// reading the pdf file
	sb = sbuf_make()
	fd = noarch.Open(pdf, 0)
	for (func() int32 {
		nr = int32(noarch.Read(fd, buf, 4096))
		return nr
	}()) > 0 {
		sbuf_mem(sb, buf, nr)
	}
	noarch.CloseOnExec(fd)
	// the XObject
	xobj_id = pdfext(sbuf_buf(sb), sbuf_len(sb), hwid, vwid)
	sbuf_free(sb)
	o_flush()
	out_fontup()
	if xobj_id >= 0 {
		sbuf_printf(pg, []byte("ET /FO%d Do BT\n\x00"), xobj_id)
	}
	p_h = -1
	p_v = -1
}

// outlink - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:725
func outlink(lnk []byte, hwid int32, vwid int32) {
	if uint32(ann_n) == 512/4 {
		return
	}
	o_flush()
	ann[func() int32 {
		defer func() {
			ann_n++
		}()
		return ann_n
	}()] = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Annot\n\x00"))
	pdfout([]byte("  /Subtype /Link\n\x00"))
	pdfout([]byte("  /Rect [%s\x00"), pdfpos(o_h, o_v))
	pdfout([]byte(" %s]\n\x00"), pdfpos(o_h+hwid, o_v+vwid))
	if int32(lnk[0]) == int32('#') {
		// internal links
		pdfout([]byte("  /A << /S /GoTo /D (%s) >>\n\x00"), lnk[0+1:])
	} else {
		// external links
		pdfout([]byte("  /A << /S /URI /URI %s >>\n\x00"), pdftext_static(lnk))
	}
	pdfout([]byte(">>\n\x00"))
	obj_end()
}

// outname - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:745
func outname(n int32, desc [][]byte, page []int32, off []int32) {
	var i int32
	o_flush()
	pdf_dests = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	for i = 0; i < n; i++ {
		if page[i] > 0 && page[i]-1 < page_n {
			pdfout([]byte("  /%s [ %d 0 R /XYZ 0 %d 0 ]\n\x00"), desc[i], page_id[page[i]-1], pdf_height-off[i]*72/dev_res)
		}
	}
	pdfout([]byte(">>\n\x00"))
	obj_end()
}

// outmark - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:761
func outmark(n int32, desc [][]byte, page []int32, off []int32, level []int32) {
	var objs []int32 = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, uint32(n)*uint32(1))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	var i int32
	var j int32
	var cnt int32
	// allocating objects
	pdf_outline = obj_map()
	for i = 0; i < n; i++ {
		objs[i] = obj_map()
	}
	o_flush()
	// root object
	obj_beg(pdf_outline)
	pdfout([]byte("<<\n\x00"))
	for i = 0; i < n; i++ {
		if level[i] == level[0] {
			cnt++
		}
	}
	pdfout([]byte("  /Count %d\n\x00"), cnt)
	pdfout([]byte("  /First %d 0 R\n\x00"), objs[0])
	for i = n - 1; i > 0 && level[i] > level[0]; i-- {
	}
	pdfout([]byte("  /Last %d 0 R\n\x00"), objs[i])
	pdfout([]byte(">>\n\x00"))
	obj_end()
	{
		// other objects
		for i = 0; i < n; i++ {
			var cnt int32
			for j = i + 1; j < n && level[j] > level[i]; j++ {
				if level[j] == level[i]+1 {
					cnt++
				}
			}
			obj_beg(objs[i])
			pdfout([]byte("<<\n\x00"))
			pdfout([]byte("  /Title %s\n\x00"), pdftext_static(desc[i]))
			{
				// the parent field
				for j = i - 1; j >= 0 && level[j] >= level[i]; j-- {
				}
			}
			pdfout([]byte("  /Parent %d 0 R\n\x00"), func() int32 {
				if j >= 0 {
					return objs[j]
				}
				return pdf_outline
			}())
			{
				// the next field
				for j = i + 1; j < n && level[j] > level[i]; j++ {
				}
			}
			if j < n && level[j] == level[i] {
				pdfout([]byte("  /Next %d 0 R\n\x00"), objs[j])
			}
			{
				// the prev field
				for j = i - 1; j >= 0 && level[j] > level[i]; j-- {
				}
			}
			if j >= 0 && level[j] == level[i] {
				pdfout([]byte("  /Prev %d 0 R\n\x00"), objs[j])
			}
			if cnt != 0 {
				// node children
				var last int32
				pdfout([]byte("  /Count %d\n\x00"), cnt)
				pdfout([]byte("  /First %d 0 R\n\x00"), objs[i+1])
				for j = i + 1; j < n && level[j] > level[i]; j++ {
					if level[j] == level[i]+1 {
						last = j
					}
				}
				pdfout([]byte("  /Last %d 0 R\n\x00"), objs[last])
			}
			if page[i] > 0 && page[i]-1 < page_n {
				pdfout([]byte("  /Dest [ %d 0 R /XYZ 0 %d 0 ]\n\x00"), page_id[page[i]-1], pdf_height-off[i]*72/dev_res)
			}
			pdfout([]byte(">>\n\x00"))
			obj_end()
		}
	}
	_ = objs
}

// outinfo - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:827
func outinfo(kwd []byte, val []byte) {
	if noarch.Not(noarch.Strcmp([]byte("Author\x00"), kwd)) {
		noarch.Snprintf(pdf_author, int32(256), []byte("%s\x00"), val)
	}
	if noarch.Not(noarch.Strcmp([]byte("Title\x00"), kwd)) {
		noarch.Snprintf(pdf_title, int32(256), []byte("%s\x00"), val)
	}
}

// outset - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:835
func outset(var_ []byte, val []byte) {
	if noarch.Not(noarch.Strcmp([]byte("linewidth\x00"), var_)) {
		pdf_linewid = noarch.Atoi(val)
	}
	if noarch.Not(noarch.Strcmp([]byte("linecap\x00"), var_)) {
		pdf_linecap = noarch.Atoi(val)
	}
	if noarch.Not(noarch.Strcmp([]byte("linejoin\x00"), var_)) {
		pdf_linejoin = noarch.Atoi(val)
	}
}

// outpage - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:845
func outpage() {
	o_v = 0
	o_h = 0
	p_i = 0
	p_v = 0
	p_h = 0
	p_s = 0
	p_f = 0
	p_m = 0
	o_i = -1
}

// outmnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:858
func outmnt(f int32) {
	if p_f == f {
		p_f = -1
	}
}

// outgname - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:864
func outgname(g int32) {
}

// drawbeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:868
func drawbeg() {
	o_flush()
	out_fontup()
	sbuf_printf(pg, []byte("%s m\n\x00"), pdfpos(o_h, o_v))
}

// l_page - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:875
// drawing line properties
var l_page int32

// l_size - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:875
var l_size int32

// l_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:875
var l_wid int32

// l_cap - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:875
var l_cap int32

// l_join - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:875
var l_join int32

// drawend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:877
func drawend(close int32, fill int32) {
	if noarch.Not(fill) {
		fill = 2
	} else {
		fill = fill
	}
	if l_page != page_n || l_size != o_s || l_wid != pdf_linewid || l_cap != pdf_linecap || l_join != pdf_linejoin {
		var lwid int32 = pdf_linewid * o_s
		sbuf_printf(pg, []byte("%d.%03d w\n\x00"), lwid/1000, lwid%1000)
		sbuf_printf(pg, []byte("%d J %d j\n\x00"), pdf_linecap, pdf_linejoin)
		l_page = page_n
		l_size = o_s
		l_wid = pdf_linewid
		l_cap = pdf_linecap
		l_join = pdf_linejoin
	}
	if fill&2 != 0 {
		// stroking color
		sbuf_printf(pg, []byte("%s RG\n\x00"), pdfcolor(o_m))
	}
	if fill&1 != 0 {
		sbuf_printf(pg, func() []byte {
			if fill&2 != 0 {
				return []byte("b\n\x00")
			}
			return []byte("f\n\x00")
		}())
	} else {
		sbuf_printf(pg, func() []byte {
			if close != 0 {
				return []byte("s\n\x00")
			}
			return []byte("S\n\x00")
		}())
	}
	p_v = 0
	p_h = 0
}

// drawmbeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:901
func drawmbeg(s []byte) {
}

// drawmend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:905
func drawmend(s []byte) {
}

// drawl - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:909
func drawl(h int32, v int32) {
	outrel(h, v)
	sbuf_printf(pg, []byte("%s l\n\x00"), pdfpos(o_h, o_v))
}

// drawquad - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:916
func drawquad(ch int32, cv int32) {
	// draw circle/ellipse quadrant
	var b int32 = 551915
	var x0 int32 = o_h * 1000
	var y0 int32 = o_v * 1000
	var x3 int32 = x0 + ch*1000/2
	var y3 int32 = y0 + cv*1000/2
	var x1 int32 = x0
	var y1 int32 = y0 + cv*b/1000/2
	var x2 int32 = x0 + ch*b/1000/2
	var y2 int32 = y3
	if ch*cv < 0 {
		x1 = x3 - ch*b/1000/2
		y1 = y0
		x2 = x3
		y2 = y3 - cv*b/1000/2
	}
	sbuf_printf(pg, []byte("%s \x00"), pdfpos00(x1/10, y1/10))
	sbuf_printf(pg, []byte("%s \x00"), pdfpos00(x2/10, y2/10))
	sbuf_printf(pg, []byte("%s c\n\x00"), pdfpos00(x3/10, y3/10))
	outrel(ch/2, cv/2)
}

// drawc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:940
func drawc(c int32) {
	// draw a circle
	drawquad(+c, +c)
	drawquad(+c, -c)
	drawquad(-c, -c)
	drawquad(-c, +c)
	outrel(c, 0)
}

// drawe - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:950
func drawe(h int32, v int32) {
	// draw an ellipse
	drawquad(+h, +v)
	drawquad(+h, -v)
	drawquad(-h, -v)
	drawquad(-h, +v)
	outrel(h, 0)
}

// drawa - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:960
func drawa(h1 int32, v1 int32, h2 int32, v2 int32) {
	// draw an arc
	drawl(h1+h2, v1+v2)
}

// draws - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:966
func draws(h1 int32, v1 int32, h2 int32, v2 int32) {
	// draw an spline
	outrel(h1, v1)
	sbuf_printf(pg, []byte("%s l\n\x00"), pdfpos(o_h, o_v))
}

// docheader - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:972
func docheader(title []byte, pagewidth int32, pageheight int32, linewidth int32) {
	if title != nil {
		outinfo([]byte("Title\x00"), title)
	}
	obj_map()
	pdf_root = obj_map()
	pdf_pages = obj_map()
	pdfout([]byte("%%PDF-1.6\n\n\x00"))
	pdf_width = (pagewidth*72 + 127) / 254
	pdf_height = (pageheight*72 + 127) / 254
	pdf_linewid = linewidth
}

// doctrailer - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:985
func doctrailer(pages int32) {
	var i int32
	var xref_off int32
	var info_id int32
	// pdf pages object
	obj_beg(pdf_pages)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Pages\n\x00"))
	pdfout([]byte("  /MediaBox [ 0 0 %d %d ]\n\x00"), pdf_width, pdf_height)
	pdfout([]byte("  /Count %d\n\x00"), page_n)
	pdfout([]byte("  /Kids [\x00"))
	for i = 0; i < page_n; i++ {
		pdfout([]byte(" %d 0 R\x00"), page_id[i])
	}
	pdfout([]byte(" ]\n\x00"))
	pdfout([]byte(">>\n\x00"))
	obj_end()
	// pdf root object
	obj_beg(pdf_root)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Catalog\n\x00"))
	pdfout([]byte("  /Pages %d 0 R\n\x00"), pdf_pages)
	if pdf_dests > 0 {
		pdfout([]byte("  /Dests %d 0 R\n\x00"), pdf_dests)
	}
	if pdf_outline > 0 {
		pdfout([]byte("  /Outlines %d 0 R\n\x00"), pdf_outline)
	}
	pdfout([]byte(">>\n\x00"))
	obj_end()
	// fonts
	pfont_done()
	// info object
	info_id = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	if pdf_title[0] != 0 {
		pdfout([]byte("  /Title %s\n\x00"), pdftext_static(pdf_title))
	}
	if pdf_author[0] != 0 {
		pdfout([]byte("  /Author %s\n\x00"), pdftext_static(pdf_author))
	}
	pdfout([]byte("  /Creator (Neatroff)\n\x00"))
	pdfout([]byte("  /Producer (Neatpost)\n\x00"))
	pdfout([]byte(">>\n\x00"))
	obj_end()
	// the xref
	xref_off = pdf_pos
	pdfout([]byte("xref\n\x00"))
	pdfout([]byte("0 %d\n\x00"), obj_n)
	pdfout([]byte("0000000000 65535 f \n\x00"))
	for i = 1; i < obj_n; i++ {
		pdfout([]byte("%010d 00000 n \n\x00"), obj_off[i])
	}
	// the trailer
	pdfout([]byte("trailer\n\x00"))
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Size %d\n\x00"), obj_n)
	pdfout([]byte("  /Root %d 0 R\n\x00"), pdf_root)
	pdfout([]byte("  /Info %d 0 R\n\x00"), info_id)
	pdfout([]byte(">>\n\x00"))
	pdfout([]byte("startxref\n\x00"))
	pdfout([]byte("%d\n\x00"), xref_off)
	pdfout([]byte("%%%%EOF\n\x00"))
	_ = page_id
	_ = obj_off
}

// docpagebeg - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:1047
func docpagebeg(n int32) {
	pg = sbuf_make()
	sbuf_printf(pg, []byte("BT\n\x00"))
}

// docpageend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdf.c:1053
func docpageend(n int32) {
	var cont_id int32
	var i int32
	o_flush()
	sbuf_printf(pg, []byte("ET\n\x00"))
	// page contents
	cont_id = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Length %d\n\x00"), sbuf_len(pg)-1)
	pdfout([]byte(">>\n\x00"))
	pdfout([]byte("stream\n\x00"))
	pdfmem(sbuf_buf(pg), sbuf_len(pg))
	pdfout([]byte("endstream\n\x00"))
	obj_end()
	if page_n == page_sz {
		// the page object
		page_sz += 1024
		page_id = mextend(page_id, page_n, page_sz, int32(4)).([]int32)
	}
	page_id[func() int32 {
		defer func() {
			page_n++
		}()
		return page_n
	}()] = obj_beg(0)
	pdfout([]byte("<<\n\x00"))
	pdfout([]byte("  /Type /Page\n\x00"))
	pdfout([]byte("  /Parent %d 0 R\n\x00"), pdf_pages)
	pdfout([]byte("  /Resources <<\n\x00"))
	pdfout([]byte("    /Font <<\x00"))
	for i = 0; i < pfonts_n; i++ {
		if o_iset[i] != 0 {
			var ps []pfont = pfonts[i:]
			if ps[0].cid != 0 {
				pdfout([]byte(" /%s %d 0 R\x00"), ps[0].name[:], ps[0].obj)
			} else {
				pdfout([]byte(" /%s.%d %d 0 R\x00"), ps[0].name[:], ps[0].sub, ps[0].obj)
			}
		}
	}
	pdfout([]byte(" >>\n\x00"))
	if xobj_n != 0 {
		// XObjects
		pdfout([]byte("    /XObject <<\x00"))
		for i = 0; i < xobj_n; i++ {
			pdfout([]byte(" /FO%d %d 0 R\x00"), i, xobj[i])
		}
		pdfout([]byte(" >>\n\x00"))
	}
	pdfout([]byte("  >>\n\x00"))
	pdfout([]byte("  /Contents %d 0 R\n\x00"), cont_id)
	if ann_n != 0 {
		pdfout([]byte("  /Annots [\x00"))
		for i = 0; i < ann_n; i++ {
			pdfout([]byte(" %d 0 R\x00"), ann[i])
		}
		pdfout([]byte(" ]\n\x00"))
	}
	pdfout([]byte(">>\n\x00"))
	obj_end()
	sbuf_free(pg)
	noarch.Memset(o_iset, byte(0), uint32(pfonts_n)*1)
	xobj_n = 0
	ann_n = 0
}

// pdf_ws - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:9
func pdf_ws(pdf []byte, len_ int32, pos int32) int32 {
	// Parse and extract PDF objects
	// the number white space characters
	var i int32 = pos
	for i < len_ && int32(((__ctype_b_loc())[0])[int32(uint8(pdf[i]))])&int32(uint16(noarch.ISspace)) != 0 {
		i++
	}
	return i - pos
}

// pdf_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:18
func pdf_type(pdf []byte, len_ int32, pos int32) int32 {
	// s: string, d: dictionary, l: list, n: number, /: name, r: reference
	pos += pdf_ws(pdf, len_, pos)
	if int32(pdf[pos]) == int32('/') {
		return int32('/')
	}
	if int32(pdf[pos]) == int32('(') {
		return int32('s')
	}
	if int32(pdf[pos]) == int32('<') && int32(pdf[pos+1]) != int32('<') {
		return int32('s')
	}
	if int32(pdf[pos]) == int32('<') && int32(pdf[pos+1]) == int32('<') {
		return int32('d')
	}
	if int32(pdf[pos]) == int32('[') {
		return int32('l')
	}
	if noarch.Strchr([]byte("0123456789+-.\x00"), int32(uint8(pdf[pos]))) != nil {
		if noarch.Not(int32(((__ctype_b_loc())[0])[int32(uint8(pdf[pos]))]) & int32(uint16(noarch.ISdigit))) {
			return int32('n')
		}
		for pos < len_ && int32(((__ctype_b_loc())[0])[int32(uint8(pdf[pos]))])&int32(uint16(noarch.ISdigit)) != 0 {
			pos++
		}
		pos += pdf_ws(pdf, len_, pos)
		if noarch.Not(int32(((__ctype_b_loc())[0])[int32(uint8(pdf[pos]))]) & int32(uint16(noarch.ISdigit))) {
			return int32('n')
		}
		for pos < len_ && int32(((__ctype_b_loc())[0])[int32(uint8(pdf[pos]))])&int32(uint16(noarch.ISdigit)) != 0 {
			pos++
		}
		pos += pdf_ws(pdf, len_, pos)
		if pos < len_ && int32(pdf[pos]) == int32('R') {
			return int32('r')
		}
		return int32('n')
	}
	return -1
}

// pdf_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:48
func pdf_len(pdf []byte, len_ int32, pos int32) int32 {
	// the length of a pdf object
	var c int32
	var old int32 = pos
	if pos >= len_ {
		return 0
	}
	pos += pdf_ws(pdf, len_, pos)
	c = int32(uint8(pdf[pos]))
	if noarch.Strchr([]byte("0123456789+-.\x00"), c) != nil {
		if pdf_type(pdf, len_, pos) == int32('r') {
			var r []byte = memchr(pdf[0+pos:], int32('R'), uint32(len_-pos)).([]byte)
			return int32((int64(uintptr(unsafe.Pointer(&r[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&pdf[0+old])))/int64(1))) + 1
		}
		pos++
		for pos < len_ && noarch.Strchr([]byte("0123456789.\x00"), int32(uint8(pdf[pos]))) != nil {
			pos++
		}
	}
	if c == int32('(') {
		var depth int32 = 1
		pos++
		for pos < len_ && depth > 0 {
			if int32(pdf[pos]) == int32('(') {
				depth++
			}
			if int32(pdf[pos]) == int32(')') {
				depth--
			}
			if int32(pdf[pos]) == int32('\\') {
				pos++
			}
			pos++
		}
	}
	if c == int32('<') && pos+1 < len_ && int32(pdf[pos+1]) == int32('<') {
		pos += 2
		for pos+2 < len_ && (int32(pdf[pos]) != int32('>') || int32(pdf[pos+1]) != int32('>')) {
			pos += pdf_len(pdf, len_, pos)
			pos += pdf_len(pdf, len_, pos)
			pos += pdf_ws(pdf, len_, pos)
		}
		if pos+2 < len_ {
			pos += 2
		}
	} else if c == int32('<') {
		for pos < len_ && int32(pdf[pos]) != int32('>') {
			pos++
		}
		if pos < len_ {
			pos++
		}
	}
	if c == int32('/') {
		pos++
		for pos < len_ && noarch.Strchr([]byte(" \t\r\n\f()<>[]{}/%\x00"), int32(uint8(pdf[pos]))) == nil {
			pos++
		}
	}
	if c == int32('[') {
		pos++
		for pos < len_ && int32(pdf[pos]) != int32(']') {
			pos += pdf_len(pdf, len_, pos)
			pos += pdf_ws(pdf, len_, pos)
		}
		pos++
	}
	return pos - old
}

// startswith - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:110
func startswith(s []byte, t []byte) int32 {
	for int32(s[0]) != 0 && int32(t[0]) != 0 {
		if int32((func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]) != int32((func() []byte {
			defer func() {
				t = t[0+1:]
			}()
			return t
		}())[0]) {
			return 0
		}
	}
	return 1
}

// pdf_obj - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:119
func pdf_obj(pdf []byte, len_ int32, pos int32, obj []int32, rev []int32) int32 {
	if pdf_type(pdf, len_, pos) != int32('r') {
		// read an indirect reference
		return -1
	}
	obj[0] = noarch.Atoi(pdf[0+pos:])
	pos += pdf_len(pdf, len_, pos)
	rev[0] = noarch.Atoi(pdf[0+pos:])
	return 0
}

// pdf_dval - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:130
func pdf_dval(pdf []byte, len_ int32, pos int32, key []byte) int32 {
	// the value of a pdf dictionary key
	pos += 2
	for pos+2 < len_ && (int32(pdf[pos]) != int32('>') || int32(pdf[pos+1]) != int32('>')) {
		pos += pdf_ws(pdf, len_, pos)
		if uint32(pdf_len(pdf, len_, pos)) == uint32(noarch.Strlen(key)) && startswith(key, pdf[0+pos:]) != 0 {
			pos += pdf_len(pdf, len_, pos)
			pos += pdf_ws(pdf, len_, pos)
			return pos
		}
		pos += pdf_len(pdf, len_, pos)
		pos += pdf_len(pdf, len_, pos)
		pos += pdf_ws(pdf, len_, pos)
	}
	return -1
}

// pdf_dkey - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:148
func pdf_dkey(pdf []byte, len_ int32, pos int32, key int32) int32 {
	// return a dictionary key
	var i int32
	pos += 2
	for pos+2 < len_ && (int32(pdf[pos]) != int32('>') || int32(pdf[pos+1]) != int32('>')) {
		pos += pdf_ws(pdf, len_, pos)
		if func() int32 {
			defer func() {
				i++
			}()
			return i
		}() == key {
			return pos
		}
		pos += pdf_len(pdf, len_, pos)
		pos += pdf_len(pdf, len_, pos)
		pos += pdf_ws(pdf, len_, pos)
	}
	return -1
}

// pdf_lval - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:164
func pdf_lval(pdf []byte, len_ int32, pos int32, idx int32) int32 {
	// return a list entry
	var i int32
	pos++
	for pos < len_ && int32(pdf[pos]) != int32(']') {
		if func() int32 {
			defer func() {
				i++
			}()
			return i
		}() == idx {
			return pos
		}
		pos += pdf_len(pdf, len_, pos)
		pos += pdf_ws(pdf, len_, pos)
	}
	return -1
}

// my_memrchr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:177
func my_memrchr(m interface{}, c int32, n int32) interface{} {
	var i int32
	for i = 0; i < n; i++ {
		if int32((func() []byte {
			var position int32 = int32(-i)
			slice := func() []byte {
				var position int32 = int32(-1)
				slice := m[0+n:]
				if position < 0 {
					position = -position
					var hdr reflect.SliceHeader
					sliceLen := len(slice)
					hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
					runtime.KeepAlive(&slice[0])
					hdr.Len = sliceLen + int(position)
					hdr.Cap = hdr.Len
					slice = *((*[]byte)(unsafe.Pointer(&hdr)))
					return slice
				}
				return slice[position:]
			}()
			if position < 0 {
				position = -position
				var hdr reflect.SliceHeader
				sliceLen := len(slice)
				hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
				runtime.KeepAlive(&slice[0])
				hdr.Len = sliceLen + int(position)
				hdr.Cap = hdr.Len
				slice = *((*[]byte)(unsafe.Pointer(&hdr)))
				return slice
			}
			return slice[position:]
		}().([]uint8))[0]) == c {
			var position int32 = int32(-i)
			slice := func() []byte {
				var position int32 = int32(-1)
				slice := m[0+n:]
				if position < 0 {
					position = -position
					var hdr reflect.SliceHeader
					sliceLen := len(slice)
					hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
					runtime.KeepAlive(&slice[0])
					hdr.Len = sliceLen + int(position)
					hdr.Cap = hdr.Len
					slice = *((*[]byte)(unsafe.Pointer(&hdr)))
					return slice
				}
				return slice[position:]
			}()
			if position < 0 {
				position = -position
				var hdr reflect.SliceHeader
				sliceLen := len(slice)
				hdr.Data = uintptr(unsafe.Pointer(&slice[0])) - (uintptr(position))*unsafe.Sizeof(slice[0])
				runtime.KeepAlive(&slice[0])
				hdr.Len = sliceLen + int(position)
				hdr.Cap = hdr.Len
				slice = *((*[]byte)(unsafe.Pointer(&hdr)))
				return slice
			}
			return slice[position:]
		}
	}
	return nil
}

// prevline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:186
func prevline(pdf []byte, len_ int32, off int32) int32 {
	var nl []byte = my_memrchr(pdf, int32('\n'), off).([]byte)
	if nl != nil && (int64(uintptr(unsafe.Pointer(&nl[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&pdf[0])))/int64(1)) > 0 {
		var nl2 []byte = my_memrchr(pdf, int32('\n'), int32((int64(uintptr(unsafe.Pointer(&nl[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&pdf[0])))/int64(1)))-1).([]byte)
		if nl2 != nil {
			return int32((int64(uintptr(unsafe.Pointer(&nl2[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&pdf[0])))/int64(1))) + 1
		}
	}
	return -1
}

// nextline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:197
func nextline(pdf []byte, len_ int32, off int32) int32 {
	var nl []byte = memchr(pdf[0+off:], int32('\n'), uint32(len_-off)).([]byte)
	if nl != nil {
		return int32((int64(uintptr(unsafe.Pointer(&nl[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&pdf[0])))/int64(1))) + 1
	}
	return -1
}

// pdf_trailer - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:206
func pdf_trailer(pdf []byte, len_ int32) int32 {
	// the position of the trailer
	// %%EOF
	var pos int32 = prevline(pdf, len_, len_)
	for noarch.Not(startswith(pdf[0+pos:], []byte("trailer\x00"))) {
		if (func() int32 {
			pos = prevline(pdf, len_, pos)
			return pos
		}()) < 0 {
			return -1
		}
	}
	// skip trailer\n
	return nextline(pdf, len_, pos)
}

// pdf_xref - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:216
func pdf_xref(pdf []byte, len_ int32) int32 {
	// the position of the last xref table
	// %%EOF
	var pos int32 = prevline(pdf, len_, len_)
	if (func() int32 {
		pos = prevline(pdf, len_, pos)
		return pos
	}()) < 0 {
		return -1
	}
	if noarch.Sscanf(pdf[0+pos:], []byte("%d\x00"), c4goUnsafeConvert_int32(&pos)) != 1 || pos >= len_ || pos < 0 {
		// read startxref offset
		return -1
	}
	// skip xref\n
	return nextline(pdf, len_, pos)
}

// pdf_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:228
func pdf_find(pdf []byte, len_ int32, obj int32, rev int32) int32 {
	// find a pdf object
	var obj_beg int32
	var obj_cnt int32
	var cur_rev int32
	var cur_pos int32
	var beg []byte
	var i int32
	var pos int32 = pdf_xref(pdf, len_)
	if pos < 0 {
		return -1
	}
	for pos < len_ && noarch.Sscanf(pdf[0+pos:], []byte("%d %d\x00"), c4goUnsafeConvert_int32(&obj_beg), c4goUnsafeConvert_int32(&obj_cnt)) == 2 {
		{
			// the numbers after xref
			for i = 0; i < obj_cnt; i++ {
				if (func() int32 {
					pos = nextline(pdf, len_, pos)
					return pos
				}()) < 0 {
					return -1
				}
				if noarch.Sscanf(pdf[0+pos:], []byte("%d %d\x00"), c4goUnsafeConvert_int32(&cur_pos), c4goUnsafeConvert_int32(&cur_rev)) != 2 {
					return -1
				}
				if obj_beg+i == obj && cur_rev == rev {
					if cur_pos < 0 || cur_pos >= len_ {
						return -1
					}
					if (func() []byte {
						beg = noarch.Strstr(pdf[0+cur_pos:], []byte("obj\x00"))
						return beg
					}()) == nil {
						return -1
					}
					pos = int32((int64(uintptr(unsafe.Pointer(&beg[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&pdf[0])))/int64(1))) + 3
					pos += pdf_ws(pdf, len_, pos)
					return pos
				}
			}
		}
	}
	return -1
}

// pdf_ref - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:259
func pdf_ref(pdf []byte, len_ int32, pos int32) int32 {
	// read and dereference an indirect reference
	var obj int32
	var rev int32
	if pdf_obj(pdf, len_, pos, c4goUnsafeConvert_int32(&obj), c4goUnsafeConvert_int32(&rev)) != 0 {
		return -1
	}
	return pdf_find(pdf, len_, obj, rev)
}

// pdf_dval_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:268
func pdf_dval_val(pdf []byte, len_ int32, pos int32, key []byte) int32 {
	// retrieve and dereference a dictionary entry
	var val int32 = pdf_dval(pdf, len_, pos, key)
	var val_obj int32
	var val_rev int32
	if val < 0 {
		return -1
	}
	if pdf_type(pdf, len_, val) == int32('r') {
		pdf_obj(pdf, len_, val, c4goUnsafeConvert_int32(&val_obj), c4goUnsafeConvert_int32(&val_rev))
		return pdf_find(pdf, len_, val_obj, val_rev)
	}
	return val
}

// pdf_dval_obj - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/pdfext.c:282
func pdf_dval_obj(pdf []byte, len_ int32, pos int32, key []byte) int32 {
	// retrieve a dictionary entry, which is an indirect reference
	var val int32 = pdf_dval(pdf, len_, pos, key)
	if val < 0 {
		return -1
	}
	return pdf_ref(pdf, len_, val)
}

// font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:7
// font handling
type font struct {
	name     [64]byte
	desc     [1024]byte
	fontname [64]byte
	fontpath [1024]byte
	spacewid int32
	gl       []glyph
	gl_n     int32
	gl_sz    int32
	gl_dict  []dict
	ch_dict  []dict
	ch_map   []dict
}

// font_find - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:21
func font_find(fn []font, name []byte) []glyph {
	// glyphs present in the font
	// number of glyphs in the font
	// mapping from gl[i].id to i
	// charset mapping
	// character aliases
	// find a glyph by its name
	var i int32 = dict_get(fn[0].ch_dict, name)
	if i < 0 {
		// maybe a character alias
		i = dict_get(fn[0].ch_map, name)
	}
	if i >= 0 {
		return fn[0].gl[0+i:]
	}
	return nil
}

// font_glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:30
func font_glyph(fn []font, id []byte) []glyph {
	// find a glyph by its device-dependent identifier
	var i int32 = dict_get(fn[0].gl_dict, id)
	if i >= 0 {
		return fn[0].gl[i:]
	}
	return nil
}

// font_glyphput - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:36
func font_glyphput(fn []font, id []byte, name []byte, type_ int32) int32 {
	var g []glyph
	if fn[0].gl_n == fn[0].gl_sz {
		fn[0].gl_sz = fn[0].gl_sz + 1024
		fn[0].gl = mextend(fn[0].gl, fn[0].gl_n, fn[0].gl_sz, int32(112)).([]glyph)
	}
	g = fn[0].gl[fn[0].gl_n:]
	noarch.Snprintf(g[0].id[:], int32(32), []byte("%s\x00"), id)
	noarch.Snprintf(g[0].name[:], int32(32), []byte("%s\x00"), name)
	g[0].type_ = type_
	g[0].font = fn
	dict_put(fn[0].gl_dict, g[0].id[:], fn[0].gl_n)
	tempVar1 := &fn[0].gl_n
	defer func() {
		*tempVar1++
	}()
	return *tempVar1
}

// tilleol - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:52
func tilleol(fin *noarch.File, s []byte) {
	var c int32 = noarch.Fgetc(fin)
	for c != -1 && c != int32('\n') {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(c)
		c = noarch.Fgetc(fin)
	}
	s[0] = '\x00'
	if c != -1 {
		ungetc(c, fin)
	}
}

// font_readchar - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:64
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
		noarch.Sscanf(tok, []byte("%d\x00"), (*[1000000]int32)(unsafe.Pointer(&g[0].wid))[:])
		tilleol(fin, tok)
		if noarch.Sscanf(tok, []byte("%d\x00"), (*[1000000]int32)(unsafe.Pointer(&g[0].pos))[:]) != 1 {
			g[0].pos = 0
		}
		dict_put(fn[0].ch_dict, name, gid[0])
		n[0]++
	} else {
		dict_put(fn[0].ch_map, name, gid[0])
	}
	return 0
}

// skipline - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:92
func skipline(filp *noarch.File) {
	var c int32
	for {
		c = noarch.Fgetc(filp)
		if !(c != int32('\n') && c != -1) {
			break
		}
	}
}

// font_open - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:100
func font_open(path []byte) []font {
	var fn []font
	// last glyph in the charset
	var ch_g int32 = -1
	// number of glyphs in the charset
	var ch_n int32
	var tok []byte = make([]byte, 128)
	var fin *noarch.File
	fin = noarch.Fopen(path, []byte("r\x00"))
	if fin == nil {
		return nil
	}
	fn = (*[1000000]font)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	if fn == nil {
		noarch.Fclose(fin)
		return nil
	}
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&fn[0]))) / int64(1))))[:], byte(0), 2280)
	noarch.Snprintf(fn[0].desc[:], int32(1024), []byte("%s\x00"), path)
	fn[0].gl_dict = dict_make(-1, 1, 0)
	fn[0].ch_dict = dict_make(-1, 1, 0)
	fn[0].ch_map = dict_make(-1, 1, 0)
	for noarch.Fscanf(fin, []byte("%128s\x00"), tok) == 1 {
		if noarch.Not(noarch.Strcmp([]byte("char\x00"), tok)) {
			font_readchar(fn, fin, c4goUnsafeConvert_int32(&ch_n), c4goUnsafeConvert_int32(&ch_g))
		} else if noarch.Not(noarch.Strcmp([]byte("spacewidth\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%d\x00"), (*[1000000]int32)(unsafe.Pointer(&fn[0].spacewid))[:])
		} else if noarch.Not(noarch.Strcmp([]byte("name\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%s\x00"), fn[0].name[:])
		} else if noarch.Not(noarch.Strcmp([]byte("fontname\x00"), tok)) {
			noarch.Fscanf(fin, []byte("%s\x00"), fn[0].fontname[:])
		} else if noarch.Not(noarch.Strcmp([]byte("fontpath\x00"), tok)) {
			var c int32 = noarch.Fgetc(fin)
			for c == int32(' ') {
				c = noarch.Fgetc(fin)
			}
			ungetc(c, fin)
			tilleol(fin, fn[0].fontpath[:])
		} else if noarch.Not(noarch.Strcmp([]byte("ligatures\x00"), tok)) {
			for noarch.Fscanf(fin, []byte("%s\x00"), tok) == 1 {
				if noarch.Not(noarch.Strcmp([]byte("0\x00"), tok)) {
					break
				}
			}
		} else if noarch.Not(noarch.Strcmp([]byte("charset\x00"), tok)) {
			for noarch.Not(font_readchar(fn, fin, c4goUnsafeConvert_int32(&ch_n), c4goUnsafeConvert_int32(&ch_g))) {
			}
			break
		}
		skipline(fin)
	}
	noarch.Fclose(fin)
	return fn
}

// font_close - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:150
func font_close(fn []font) {
	dict_free(fn[0].gl_dict)
	dict_free(fn[0].ch_dict)
	dict_free(fn[0].ch_map)
	_ = fn[0].gl
	_ = fn
}

// font_wid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:160
func font_wid(fn []font, sz int32, w int32) int32 {
	// return width w for the given font and size
	return (w*sz + dev_uwid/2) / dev_uwid
}

// font_swid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:166
func font_swid(fn []font, sz int32) int32 {
	// space width for the give word space or sentence space
	return font_wid(fn, sz, fn[0].spacewid)
}

// font_name - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:171
func font_name(fn []font) []byte {
	if int32(fn[0].fontname[:][0]) != 0 {
		return fn[0].fontname[:]
	}
	return fn[0].name[:]
}

// font_path - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:176
func font_path(fn []font) []byte {
	return fn[0].fontpath[:]
}

// font_glnum - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:181
func font_glnum(fn []font, g []glyph) int32 {
	return int32((int64(uintptr(unsafe.Pointer(&g[0])))/int64(112) - int64(uintptr(unsafe.Pointer(&fn[0].gl[0])))/int64(112)))
}

// font_glget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:186
func font_glget(fn []font, id int32) []glyph {
	if id >= 0 && id < fn[0].gl_n {
		return fn[0].gl[id:]
	}
	return nil
}

// font_desc - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/font.c:191
func font_desc(fn []font) []byte {
	return fn[0].desc[:]
}

// dev_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:7
// device directory
var dev_dir []byte = make([]byte, 1024)

// dev_dev - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:8
// output device name
var dev_dev []byte = make([]byte, 1024)

// dev_res - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:9
// device resolution
var dev_res int32

// dev_uwid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:10
// device unitwidth
var dev_uwid int32

// dev_hor - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:11
// minimum horizontal movement
var dev_hor int32

// dev_ver - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:12
// minimum vertical movement
var dev_ver int32

// fn_name - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:15
// mounted fonts
// font names
var fn_name [][]byte = make([][]byte, 1024)

// fn_font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:16
// font structs
var fn_font [][]font = make([][]font, 1024)

// fn_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:17
// number of device fonts
var fn_n int32

// dev_fontopen - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:27
func dev_fontopen(name []byte) []font {
	var path []byte = make([]byte, 1024)
	if noarch.Strchr(name, int32('/')) != nil {
		noarch.Strcpy(path, name)
	} else {
		noarch.Sprintf(path, []byte("%s/dev%s/%s\x00"), dev_dir, dev_dev, name)
	}
	return font_open(path)
}

// dev_mnt - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:37
func dev_mnt(pos int32, id []byte, name []byte) int32 {
	var fn []font
	if pos >= 1024 {
		return -1
	}
	fn = dev_fontopen(name)
	if fn == nil {
		return -1
	}
	if fn_font[pos] != nil {
		font_close(fn_font[pos])
	}
	if (int64(uintptr(unsafe.Pointer(&fn_name[pos])))/int64(1) - int64(uintptr(unsafe.Pointer(&name[0])))/int64(1)) != 0 {
		// ignore if fn_name[pos] is passed
		noarch.Snprintf(fn_name[pos], int32(64), []byte("%s\x00"), id)
	}
	fn_font[pos] = fn
	return pos
}

// dev_open - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:53
func dev_open(dir []byte, dev []byte) int32 {
	var path []byte = make([]byte, 1024)
	var tok []byte = make([]byte, 128)
	var i int32
	var desc *noarch.File
	noarch.Strcpy(dev_dir, dir)
	noarch.Strcpy(dev_dev, dev)
	noarch.Sprintf(path, []byte("%s/dev%s/DESC\x00"), dir, dev)
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
	return 0
}

// dev_close - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:107
func dev_close() {
	var i int32
	for i = 0; i < 1024; i++ {
		if fn_font[i] != nil {
			font_close(fn_font[i])
		}
		fn_font[i] = nil
	}
}

// dev_glyph - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:117
func dev_glyph(c []byte, fn int32) []glyph {
	if noarch.Not(strncmp([]byte("GID=\x00"), c, 4)) {
		return font_glyph(fn_font[fn], c[0+4:])
	}
	return font_find(fn_font[fn], c)
}

// dev_font - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:125
func dev_font(pos int32) []font {
	// return the font struct at pos
	if pos >= 0 && pos < 1024 {
		return fn_font[pos]
	}
	return nil
}

// dev_fontid - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/dev.c:130
func dev_fontid(fn []font) int32 {
	var i int32
	for i = 0; i < 1024; i++ {
		if (int64(uintptr(unsafe.Pointer(&fn_font[i])))/int64(2280) - int64(uintptr(unsafe.Pointer(&fn[0])))/int64(2280)) == 0 {
			return i
		}
	}
	return 0
}

// clr_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:8
func clr_str(c int32) []byte {
	// returns a static buffer
	var clr_buf []byte = make([]byte, 32)
	if noarch.Not(c) {
		return []byte("0\x00")
	}
	noarch.Sprintf(clr_buf, []byte("#%02x%02x%02x\x00"), c>>uint64(16)&255, c>>uint64(8)&255, c&255)
	return clr_buf
}

// color - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:17
type color struct {
	name  []byte
	value int32
}

// colors - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:17
var colors []color = []color{{[]byte("black\x00"), (0<<uint64(16) | 0<<uint64(8) | 0)}, {[]byte("red\x00"), (255<<uint64(16) | 0<<uint64(8) | 0)}, {[]byte("green\x00"), (0<<uint64(16) | 255<<uint64(8) | 0)}, {[]byte("yellow\x00"), (255<<uint64(16) | 255<<uint64(8) | 0)}, {[]byte("blue\x00"), (0<<uint64(16) | 0<<uint64(8) | 255)}, {[]byte("magenta\x00"), (255<<uint64(16) | 0<<uint64(8) | 255)}, {[]byte("cyan\x00"), (0<<uint64(16) | 255<<uint64(8) | 255)}, {[]byte("white\x00"), (255<<uint64(16) | 255<<uint64(8) | 255)}}

// clrcomp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:32
func clrcomp(s []byte, len_ int32) int32 {
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

// clr_get - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/clr.c:43
func clr_get(s []byte) int32 {
	var i int32
	if int32(s[0]) == int32('#') && noarch.Strlen(s) == int32(7) {
		return clrcomp(s[0+1:], 2)<<uint64(16) | clrcomp(s[0+3:], 2)<<uint64(8) | clrcomp(s[0+5:], 2)
	}
	if int32(s[0]) == int32('#') && noarch.Strlen(s) == int32(4) {
		return clrcomp(s[0+1:], 1)<<uint64(16) | clrcomp(s[0+2:], 1)<<uint64(8) | clrcomp(s[0+3:], 1)
	}
	if int32(((__ctype_b_loc())[0])[int32(s[0])])&int32(uint16(noarch.ISdigit)) != 0 && noarch.Atoi(s) >= 0 && uint32(noarch.Atoi(s)) < 192/24 {
		return colors[noarch.Atoi(s)].value
	}
	for i = 0; uint32(i) < 192/24; i++ {
		if noarch.Not(noarch.Strcmp(colors[i].name, s)) {
			return colors[i].value
		}
	}
	return 0
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
	var d []dict = (*[1000000]dict)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
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
		var dup []byte = make([]byte, uint32(len_))
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
	// match a prefix of key; in the first call, *idx should be -1
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
	var iset_c4go_postfix []iset = (*[1000000]iset)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
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
		var nset interface{} = make([]uint32, uint32(nlen)*uint32(1))
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

// sbuf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:12
// variable length string buffer
type sbuf struct {
	s    []byte
	s_n  int32
	s_sz int32
}

// sbuf_extend - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:18
func sbuf_extend(sbuf_c4go_postfix []sbuf, newsz int32) {
	// allocated buffer
	// length of the string stored in s[]
	// size of memory allocated for s[]
	var s []byte = sbuf_c4go_postfix[0].s
	sbuf_c4go_postfix[0].s_sz = newsz
	sbuf_c4go_postfix[0].s = make([]byte, uint32(sbuf_c4go_postfix[0].s_sz))
	if sbuf_c4go_postfix[0].s_n != 0 {
		memcpy(sbuf_c4go_postfix[0].s, s, uint32(sbuf_c4go_postfix[0].s_n))
	}
	_ = s
}

// sbuf_make - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:28
func sbuf_make() []sbuf {
	var sb []sbuf = (*[1000000]sbuf)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&sb[0]))) / int64(1))))[:], byte(0), 24)
	return sb
}

// sbuf_buf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:35
func sbuf_buf(sb []sbuf) []byte {
	if sb[0].s == nil {
		sbuf_extend(sb, 1)
	}
	sb[0].s[sb[0].s_n] = '\x00'
	return sb[0].s
}

// sbuf_done - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:43
func sbuf_done(sb []sbuf) []byte {
	var s []byte = sbuf_buf(sb)
	_ = sb
	return s
}

// sbuf_free - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:50
func sbuf_free(sb []sbuf) {
	_ = sb[0].s
	_ = sb
}

// sbuf_chr - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:56
func sbuf_chr(sbuf_c4go_postfix []sbuf, c int32) {
	if sbuf_c4go_postfix[0].s_n+2 >= sbuf_c4go_postfix[0].s_sz {
		sbuf_extend(sbuf_c4go_postfix, (func() int32 {
			if sbuf_c4go_postfix[0].s_sz*2 < sbuf_c4go_postfix[0].s_sz+1 {
				return sbuf_c4go_postfix[0].s_sz + 1
			}
			return sbuf_c4go_postfix[0].s_sz * 2
		}()+128-1) & ^(128-1))
	}
	sbuf_c4go_postfix[0].s[func() int32 {
		tempVar1 := &sbuf_c4go_postfix[0].s_n
		defer func() {
			*tempVar1++
		}()
		return *tempVar1
	}()] = byte(c)
}

// sbuf_mem - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:63
func sbuf_mem(sbuf_c4go_postfix []sbuf, s []byte, len_ int32) {
	if sbuf_c4go_postfix[0].s_n+len_+1 >= sbuf_c4go_postfix[0].s_sz {
		sbuf_extend(sbuf_c4go_postfix, (func() int32 {
			if sbuf_c4go_postfix[0].s_sz*2 < sbuf_c4go_postfix[0].s_sz+(len_+1) {
				return sbuf_c4go_postfix[0].s_sz + (len_ + 1)
			}
			return sbuf_c4go_postfix[0].s_sz * 2
		}()+128-1) & ^(128-1))
	}
	memcpy(sbuf_c4go_postfix[0].s[0+sbuf_c4go_postfix[0].s_n:], s, uint32(len_))
	sbuf_c4go_postfix[0].s_n += len_
}

// sbuf_str - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:71
func sbuf_str(sbuf_c4go_postfix []sbuf, s []byte) {
	sbuf_mem(sbuf_c4go_postfix, s, noarch.Strlen(s))
}

// sbuf_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:76
func sbuf_len(sbuf_c4go_postfix []sbuf) int32 {
	return sbuf_c4go_postfix[0].s_n
}

// sbuf_cut - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:81
func sbuf_cut(sb []sbuf, len_ int32) {
	if sb[0].s_n > len_ {
		sb[0].s_n = len_
	}
}

// sbuf_printf - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/sbuf.c:87
func sbuf_printf(sbuf_c4go_postfix []sbuf, s []byte, c4goArgs ...interface{}) {
	var buf []byte = make([]byte, 256)
	var ap *va_list
	va_start(ap, s)
	noarch.Vsnprintf(buf, int32(256), s, ap)
	va_end(ap)
	sbuf_str(sbuf_c4go_postfix, buf)
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[1000000]int32)(unsafe.Pointer(c4go_name))[:]
}

// device-dependent glyph identifier
// the first character mapped to this glyph
// glyph font
// character width
// character type; ascender/descender
// glyph code
// output device functions
// font-related functions
// output functions
// colors
// mapping integers to sets
// mapping strings to longs
// memory allocation
// helper functions
// string buffers
// reading PDF files

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

// tolower from ctype.h
// c function : int tolower(int)
// dep pkg    : unicode
// dep func   :
func tolower(_c int32) int32 {
	return int32(unicode.ToLower(rune(_c)))
}

// memchr - add c-binding for implemention function
func memchr(arg0 interface{}, arg1 int32, arg2 uint32) interface{} {
	return interface{}(C.memchr(unsafe.Pointer(&arg0), C.int(arg1), C.ulong(arg2)))
}

// strncmp - add c-binding for implemention function
func strncmp(arg0 []byte, arg1 []byte, arg2 uint32) int32 {
	return int32(C.strncmp((*C.char)(unsafe.Pointer(&arg0[0])), (*C.char)(unsafe.Pointer(&arg1[0])), C.ulong(arg2)))
}

// ungetc - add c-binding for implemention function
func ungetc(arg0 int32, arg1 *noarch.File) int32 {
	return int32(C.ungetc(C.int(arg0), (*C.FILE)(unsafe.Pointer(&arg1))))
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
