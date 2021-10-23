//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "unicode"
import "os"
import "reflect"
import "runtime"
import "sort"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// ref - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:28
//
// * NEATREFER - A REFER CLONE FOR NEATROFF
// *
// * Copyright (C) 2011-2017 Ali Gholami Rudi <ali at rudi dot ir>
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
type ref struct {
	keys  [128][]byte
	auth  [128][]byte
	id    int32
	nauth int32
}

// refs - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:35
// reference keys
// authors
// allocated reference id
// all references in refer database
var refs []ref = make([]ref, 16384)

// refs_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:36
var refs_n int32

// cites - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:37
// cited references
var cites [][]ref = make([][]ref, 16384)

// cites_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:38
var cites_n int32

// inserted - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:39
// number of inserted references
var inserted int32

// multiref - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:40
// allow specifying multiple references
var multiref int32

// accumulate - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:41
// accumulate all references
var accumulate int32

// initials - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:42
// initials for authors' first name
var initials int32

// refauth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:43
// use author-year citations
var refauth int32

// sortall - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:44
// sort references
var sortall int32

// refmac - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:45
// citation macro name
var refmac []byte

// refmac_auth - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:46
// author-year citation macro name
var refmac_auth []byte

// refdb - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:47
// the database file
var refdb *noarch.File

// lnget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:52
func lnget() []byte {
	// the next input line
	var buf []byte = make([]byte, 1024)
	return noarch.Fgets(buf, int32(1024), noarch.Stdin)
}

// lnput - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:59
func lnput(s []byte, n int32) {
	// write an output line
	noarch.Write(1, s, func() uint32 {
		if n >= 0 {
			return uint32(n)
		}
		return uint32(noarch.Strlen(s))
	}())
}

// dbget - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:65
func dbget() []byte {
	// the next refer database input line
	var buf []byte = make([]byte, 1024)
	if refdb != nil {
		return noarch.Fgets(buf, int32(1024), refdb)
	}
	return nil
}

// sdup - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:71
func sdup(s []byte) []byte {
	var e []byte = func() []byte {
		if noarch.Strchr(s, int32('\n')) != nil {
			return noarch.Strchr(s, int32('\n'))
		}
		return noarch.Strchr(s, int32('\x00'))
	}()
	var r []byte
	var n int32 = int32((int64(uintptr(unsafe.Pointer(&e[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)))
	r = make([]byte, uint32(n+1))
	memcpy(r, s, uint32(n))
	r[n] = '\x00'
	return r
}

// ref_author - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:83
func ref_author(ref []byte) []byte {
	// format author names as J. Smith
	var res []byte
	var out []byte
	var beg []byte
	if noarch.Not(initials) {
		return sdup(ref)
	}
	res = make([]byte, noarch.Strlen(ref)+int32(32))
	out = res
	for 1 != 0 {
		for int32(ref[0]) == int32(' ') || int32(ref[0]) == int32('.') {
			ref = ref[0+1:]
		}
		if int32(ref[0]) == int32('\x00') {
			break
		}
		beg = ref
		for int32(ref[0]) != 0 && int32(ref[0]) != int32(' ') && int32(ref[0]) != int32('.') {
			ref = ref[0+1:]
		}
		if (int64(uintptr(unsafe.Pointer(&out[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&res[0])))/int64(1)) != 0 {
			(func() []byte {
				defer func() {
					out = out[0+1:]
				}()
				return out
			}())[0] = ' '
		}
		if int32(((__ctype_b_loc())[0])[int32(uint8(beg[0]))])&int32(uint16(noarch.ISlower)) != 0 || int32(ref[0]) == int32('\x00') {
			for (int64(uintptr(unsafe.Pointer(&beg[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&ref[0])))/int64(1)) < 0 {
				(func() []byte {
					defer func() {
						out = out[0+1:]
					}()
					return out
				}())[0] = (func() []byte {
					defer func() {
						beg = beg[0+1:]
					}()
					return beg
				}())[0]
			}
		} else {
			for {
				// initials
				(func() []byte {
					defer func() {
						out = out[0+1:]
					}()
					return out
				}())[0] = (func() []byte {
					defer func() {
						beg = beg[0+1:]
					}()
					return beg
				}())[0]
				(func() []byte {
					defer func() {
						out = out[0+1:]
					}()
					return out
				}())[0] = '.'
				for (int64(uintptr(unsafe.Pointer(&beg[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&ref[0])))/int64(1)) < 0 && int32(beg[0]) != int32('-') {
					beg = beg[0+1:]
				}
				if int32(beg[0]) == int32('-') {
					// handling J.-K. Smith
					(func() []byte {
						defer func() {
							out = out[0+1:]
						}()
						return out
					}())[0] = (func() []byte {
						defer func() {
							beg = beg[0+1:]
						}()
						return beg
					}())[0]
				}
				if !((int64(uintptr(unsafe.Pointer(&beg[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&ref[0])))/int64(1)) < 0) {
					break
				}
			}
		}
	}
	out[0] = '\x00'
	return res
}

// rstrip - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:121
func rstrip(s []byte) {
	// strip excess whitespace
	var i int32
	var last int32 = -1
	for i = 0; s[i] != 0; i++ {
		if int32(s[i]) != int32(' ') && int32(s[i]) != int32('\n') {
			last = i
		}
	}
	s[last+1] = '\x00'
}

// db_ref - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:132
func db_ref(ref_c4go_postfix []ref, ln []byte) {
	for {
		if int32(ln[0]) == int32('%') && int32(ln[1]) >= int32('A') && int32(ln[1]) <= int32('Z') {
			// read a single refer record
			var r []byte = ln[0+2:]
			for int32(((__ctype_b_loc())[0])[int32(uint8(r[0]))])&int32(uint16(noarch.ISspace)) != 0 {
				r = r[0+1:]
			}
			rstrip(r)
			if int32(ln[1]) == int32('A') {
				ref_c4go_postfix[0].auth[:][func() int32 {
					tempVar1 := &ref_c4go_postfix[0].nauth
					defer func() {
						*tempVar1++
					}()
					return *tempVar1
				}()] = ref_author(r)
			} else {
				ref_c4go_postfix[0].keys[:][uint8(ln[1])] = sdup(r)
			}
			ref_c4go_postfix[0].id = -1
		}
		if !((func() []byte {
			ln = dbget()
			return ln
		}()) != nil && int32(ln[0]) != int32('\n')) {
			break
		}
	}
}

// db_parse - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:150
func db_parse() int32 {
	// parse a refer-style bib file and fill refs[]
	var ln []byte
	for (func() []byte {
		ln = dbget()
		return ln
	}()) != nil {
		if int32(ln[0]) != int32('\n') {
			db_ref(refs[func() int32 {
				defer func() {
					refs_n++
				}()
				return refs_n
			}():], ln)
		}
	}
	return 0
}

// fields - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:159
var fields []byte = []byte("LTABERJDVNPITOH\x00")

// fields_flag - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:160
var fields_flag []byte = []byte("OP\x00")

// kinds - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:161
var kinds [][]byte = [][]byte{[]byte("Other\x00"), []byte("Article\x00"), []byte("Book\x00"), []byte("In book\x00"), []byte("Report\x00")}

// ref_kind - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:163
func ref_kind(r []ref) int32 {
	if r[0].keys[:]['J'] != nil {
		return 1
	}
	if r[0].keys[:]['B'] != nil {
		return 3
	}
	if r[0].keys[:]['R'] != nil {
		return 4
	}
	if r[0].keys[:]['I'] != nil {
		return 2
	}
	return 0
}

// ref_ins - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:177
func ref_ins(ref_c4go_postfix []ref, id int32) {
	// print the given reference
	var buf []byte = make([]byte, 4096)
	var s []byte = buf
	var kind int32 = ref_kind(ref_c4go_postfix)
	var j int32
	s = s[0+noarch.Sprintf(s, []byte(".ds [F %d\n\x00"), id):]
	s = s[0+noarch.Sprintf(s, []byte(".]-\n\x00")):]
	if ref_c4go_postfix[0].nauth != 0 {
		s = s[0+noarch.Sprintf(s, []byte(".ds [A \x00")):]
		for j = 0; j < ref_c4go_postfix[0].nauth; j++ {
			s = s[0+noarch.Sprintf(s, []byte("%s%s\x00"), func() []byte {
				if j != 0 {
					return []byte(", \x00")
				}
				return []byte("\x00")
			}(), ref_c4go_postfix[0].auth[:][j]):]
		}
		s = s[0+noarch.Sprintf(s, []byte("\n\x00")):]
	}
	for j = int32('B'); j <= int32('Z'); j++ {
		var val []byte = ref_c4go_postfix[0].keys[:][j]
		if val == nil || noarch.Strchr(fields, j) == nil {
			continue
		}
		s = s[0+noarch.Sprintf(s, []byte(".ds [%c %s\n\x00"), j, func() []byte {
			if val != nil {
				return val
			}
			return []byte("\x00")
		}()):]
		if noarch.Strchr(fields_flag, j) != nil {
			s = s[0+noarch.Sprintf(s, []byte(".nr [%c 1\n\x00"), j):]
		}
	}
	s = s[0+noarch.Sprintf(s, []byte(".][ %d %s\n\x00"), kind, kinds[kind]):]
	lnput(buf, int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&buf[0])))/int64(1))))
}

// lastname - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:203
func lastname(name []byte) []byte {
	var last []byte = name
	for name[0] != 0 {
		if noarch.Not(int32(((__ctype_b_loc())[0])[int32(uint8(last[0]))]) & int32(uint16(noarch.ISlower))) {
			last = name
		}
		for int32(name[0]) != 0 && int32(name[0]) != int32(' ') {
			if int32((func() []byte {
				defer func() {
					name = name[0+1:]
				}()
				return name
			}())[0]) == int32('\\') {
				name = name[0+1:]
			}
		}
		for int32(name[0]) == int32(' ') {
			name = name[0+1:]
		}
	}
	return last
}

// refcmp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:218
func refcmp(r1 []ref, r2 []ref) int32 {
	if noarch.Not(r2[0].nauth) || r1[0].keys[:]['H'] != nil && r2[0].keys[:]['H'] == nil {
		return -1
	}
	if noarch.Not(r1[0].nauth) || r1[0].keys[:]['H'] == nil && r2[0].keys[:]['H'] != nil {
		return 1
	}
	return noarch.Strcmp(lastname(r1[0].auth[:][0]), lastname(r2[0].auth[:][0]))
}

// ref_all - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:228
func ref_all() {
	// print all references
	var i int32
	var j int32
	var sorted [][]ref
	sorted = (*[1000000][]ref)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, uint32(cites_n)*uint32(1))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	memcpy(sorted, cites, uint32(cites_n)*8)
	if sortall == int32('a') {
		for i = 1; i < cites_n; i++ {
			for j = i - 1; j >= 0 && refcmp(cites[i], sorted[j]) < 0; j-- {
				sorted[j+1] = sorted[j]
			}
			sorted[j+1] = cites[i]
		}
	}
	lnput([]byte(".]<\n\x00"), -1)
	for i = 0; i < cites_n; i++ {
		ref_ins(sorted[i], sorted[i][0].id+1)
	}
	lnput([]byte(".]>\x00"), -1)
	_ = sorted
}

// intcmp - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:248
func intcmp(v1 interface{}, v2 interface{}) int32 {
	return (v1.([]int32))[0] - (v2.([]int32))[0]
}

// refer_seen - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:254
func refer_seen(label []byte) int32 {
	// the given label was referenced; add it to cites[]
	var i int32
	for i = 0; i < refs_n; i++ {
		if (refs[i:])[0].keys[:]['L'] != nil && noarch.Not(noarch.Strcmp(label, ((refs[i:])[0].keys[:]['L']))) {
			break
		}
	}
	if i == refs_n {
		return -1
	}
	if refs[i].id < 0 {
		refs[i].id = func() int32 {
			defer func() {
				cites_n++
			}()
			return cites_n
		}()
		cites[refs[i].id] = refs[i:]
	}
	return refs[i].id
}

// refer_quote - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:269
func refer_quote(d []byte, s []byte) {
	if noarch.Strchr(s, int32(' ')) == nil && int32(s[0]) != int32('"') {
		noarch.Strcpy(d, s)
	} else {
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = '"'
		for s[0] != 0 {
			if int32(s[0]) == int32('"') {
				(func() []byte {
					defer func() {
						d = d[0+1:]
					}()
					return d
				}())[0] = '"'
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
		}
		(func() []byte {
			defer func() {
				d = d[0+1:]
			}()
			return d
		}())[0] = '"'
		d[0] = '\x00'
	}
}

// refer_cite - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:286
func refer_cite(id []int32, s []byte, auth int32) int32 {
	// replace .[ .] macros with reference numbers or author-year
	var msg []byte = make([]byte, 256)
	var label []byte = make([]byte, 256)
	var nid int32
	var i int32
	msg[0] = '\x00'
	for noarch.Not(nid) || multiref != 0 {
		var r []byte = label
		for int32(s[0]) != 0 && noarch.Strchr([]byte(" \t\n,\x00"), int32(uint8(s[0]))) != nil {
			s = s[0+1:]
		}
		for int32(s[0]) != 0 && noarch.Strchr([]byte(" \t\n,]\x00"), int32(uint8(s[0]))) == nil {
			(func() []byte {
				defer func() {
					r = r[0+1:]
				}()
				return r
			}())[0] = (func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0]
		}
		r[0] = '\x00'
		if noarch.Not(noarch.Strcmp([]byte("$LIST$\x00"), label)) {
			ref_all()
			break
		}
		id[nid] = refer_seen(label)
		if id[nid] < 0 {
			noarch.Fprintf(noarch.Stderr, []byte("refer: <%s> not found\n\x00"), label)
		} else {
			nid++
		}
		if noarch.Not(s[0]) || int32(s[0]) == int32('\n') || int32(s[0]) == int32(']') {
			break
		}
	}
	if noarch.Not(auth) {
		// numbered citations
		// sort references for cleaner reference intervals
		sort.SliceStable(id[:int32(uint32(nid))], func(a, b int) bool {
			return intcmp((*[1000000]int32)(unsafe.Pointer(&id[a]))[:], (*[1000000]int32)(unsafe.Pointer(&id[b]))[:]) <= 0
		})
		for i < nid {
			var beg int32 = func() int32 {
				defer func() {
					i++
				}()
				return i
			}()
			for i < nid && id[i] == id[i-1]+1 {
				// reading reference intervals
				i++
			}
			if beg != 0 {
				noarch.Sprintf(msg[0+noarch.Strlen(msg):], []byte(",\x00"))
			}
			if beg == i-1 {
				noarch.Sprintf(msg[0+noarch.Strlen(msg):], []byte("%d\x00"), id[beg]+1)
			} else {
				noarch.Sprintf(msg[0+noarch.Strlen(msg):], []byte("%d%s%d\x00"), id[beg]+1, func() []byte {
					if beg < i-2 {
						return []byte("\\-\x00")
					}
					return []byte(",\x00")
				}(), id[i-1]+1)
			}
		}
	} else if nid != 0 {
		// year + authors citations
		var ref_c4go_postfix []ref = cites[id[0]]
		noarch.Sprintf(msg, []byte("%s %d\x00"), func() []byte {
			if ref_c4go_postfix[0].keys[:]['D'] != nil {
				return ref_c4go_postfix[0].keys[:]['D']
			}
			return []byte("-\x00")
		}(), ref_c4go_postfix[0].nauth)
		for i = 0; i < ref_c4go_postfix[0].nauth; i++ {
			noarch.Sprintf(msg[0+noarch.Strlen(msg):], []byte(" \x00"))
			refer_quote(msg[0+noarch.Strlen(msg):], lastname(ref_c4go_postfix[0].auth[:][i]))
		}
	}
	lnput(msg, -1)
	return nid
}

// slen - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:340
func slen(s []byte, delim int32) int32 {
	var r []byte = noarch.Strchr(s, delim)
	if r != nil {
		return int32((int64(uintptr(unsafe.Pointer(&r[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)))
	}
	return int32((func() int64 {
		c4go_temp_name := noarch.Strchr(s, int32('\x00'))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}() - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)))
}

// refer_reqname - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:346
func refer_reqname(mac []byte, maclen int32, s []byte) int32 {
	var i int32
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('.') {
		return 1
	}
	for i = 0; i < maclen && int32(s[0]) != 0 && int32(s[0]) != int32(' '); i++ {
		mac[i] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	}
	mac[i] = '\x00'
	return noarch.BoolToInt(int32(s[0]) != int32(' '))
}

// refer_macname - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:357
func refer_macname(mac []byte, maclen int32, s []byte) int32 {
	var i int32
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('\\') {
		return 1
	}
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('*') {
		return 1
	}
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('[') {
		return 1
	}
	for i = 0; i < maclen && int32(s[0]) != 0 && int32(s[0]) != int32(' '); i++ {
		mac[i] = (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	}
	mac[i] = '\x00'
	return noarch.BoolToInt(int32(s[0]) != int32(' '))
}

// refer_refmac - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:373
func refer_refmac(pat []byte, mac []byte) int32 {
	// return 1 if mac is a citation macro
	var s []byte = func() []byte {
		if pat != nil {
			return noarch.Strstr(pat, mac)
		}
		return nil
	}()
	if noarch.Not(mac[0]) || s == nil {
		return 0
	}
	return noarch.BoolToInt(((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&pat[0])))/int64(1)) == 0 || int32(c4goPointerArithByteSlice(s, int(-1))[0]) == int32(',')) && (noarch.Not(s[noarch.Strlen(mac)]) || int32(s[noarch.Strlen(mac)]) == int32(',')))
}

// refer_insert - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:382
func refer_insert(id []int32, id_n int32) {
	var i int32
	for i = 0; i < id_n; i++ {
		ref_ins(cites[id[i]], func() int32 {
			inserted++
			return inserted
		}())
	}
}

// refer - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:389
func refer() {
	var mac []byte = make([]byte, 256)
	var id []int32 = make([]int32, 256)
	var s []byte
	var r []byte
	var ln []byte
	for (func() []byte {
		ln = lnget()
		return ln
	}()) != nil {
		var id_n int32
		if int32(ln[0]) == int32('.') && int32(ln[1]) == int32('[') {
			// multi-line citations: .[ rudi17 .]
			lnput(ln[0+2:], slen(ln[0+2:], int32('\n')))
			if (func() []byte {
				ln = lnget()
				return ln
			}()) != nil {
				id_n = refer_cite(id, ln, 0)
				for ln != nil && (int32(ln[0]) != int32('.') || int32(ln[1]) != int32(']')) {
					ln = lnget()
				}
				if ln != nil {
					lnput(ln[0+2:], -1)
				}
			}
			if noarch.Not(accumulate) {
				refer_insert(id, id_n)
			}
			continue
		}
		if int32(ln[0]) == int32('.') && noarch.Not(refer_reqname(mac, int32(256), ln)) && (refer_refmac(refmac, mac) != 0 || refer_refmac(refmac_auth, mac) != 0) {
			// single line citation .cite rudi17
			var i int32 = 1
			for int32(ln[i]) != 0 && int32(ln[i]) != int32(' ') {
				i++
			}
			for int32(ln[i]) != 0 && int32(ln[i]) == int32(' ') {
				i++
			}
			lnput(ln, i)
			id_n = refer_cite(id, ln[0+i:], refer_refmac(refmac_auth, mac))
			for int32(ln[i]) != 0 && int32(ln[i]) != int32(' ') && int32(ln[i]) != int32('\n') {
				i++
			}
			lnput(ln[0+i:], -1)
			if noarch.Not(accumulate) {
				refer_insert(id, id_n)
			}
			continue
		}
		s = ln
		r = s
		for (func() []byte {
			r = noarch.Strchr(r, int32('\\'))
			return r
		}()) != nil {
			// inline citations \*[cite rudi17]
			r = r[0+1:]
			if refer_macname(mac, int32(256), c4goPointerArithByteSlice(r, int(-1))) != 0 {
				continue
			}
			if noarch.Not(refer_refmac(refmac, mac)) && noarch.Not(refer_refmac(refmac_auth, mac)) {
				continue
			}
			if noarch.Strchr(r, int32(']')) == nil {
				continue
			}
			r = (noarch.Strchr(r, int32(' ')))[0+1:]
			lnput(s, int32((int64(uintptr(unsafe.Pointer(&r[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1))))
			id_n = refer_cite(id, r, refer_refmac(refmac_auth, mac))
			for int32(r[0]) != 0 && int32(r[0]) != int32(' ') && int32(r[0]) != int32(']') {
				r = r[0+1:]
			}
			s = r
		}
		lnput(s, -1)
		if noarch.Not(accumulate) {
			refer_insert(id, id_n)
		}
	}
}

// usage - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:451
var usage []byte = []byte("Usage neatrefer [options] <input >output\nOptions:\n\t-p bib    \tspecify the database file\n\t-e        \taccumulate references\n\t-m        \tmerge multiple references in a single .[/.] block\n\t-i        \tinitials for authors' first and middle names\n\t-o xy     \tcitation macro (\\*[xy label])\n\t-a xy     \tauthor-year citation macro (\\*[xy label])\n\t-sa       \tsort by author last names\n\x00")

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/refer.c:462
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	var i int32
	var j int32
	for i = 1; i < argc; i++ {
		switch func() int32 {
			if int32(argv[i][0]) == int32('-') {
				return int32(argv[i][1])
			}
			return int32('h')
		}() {
		case 'm':
			multiref = 1
		case 'e':
			accumulate = 1
		case 'p':
			refdb = noarch.Fopen(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}(), []byte("r\x00"))
			if refdb != nil {
				db_parse()
				noarch.Fclose(refdb)
			}
			refdb = nil
		case 'o':
			refmac = func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}()
		case 'i':
			initials = 1
		case 'a':
			refmac_auth = func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}()
		case 's':
			sortall = int32(uint8(func() int32 {
				if int32(argv[i][2]) != 0 {
					return int32(argv[i][2])
				}
				return int32(argv[func() int32 {
					i++
					return i
				}()][0])
			}()))
		default:
			noarch.Printf([]byte("%s\x00"), usage)
			noarch.Exit(int32(1))
		}
	}
	if refauth != 0 && multiref != 0 {
		noarch.Fprintf(noarch.Stderr, []byte("refer: cannot use -m with -a\n\x00"))
		noarch.Exit(int32(1))
	}
	refer()
	for i = 0; i < refs_n; i++ {
		for j = 0; uint32(j) < 1024/8; j++ {
			if refs[i].keys[:][j] != nil {
				_ = refs[i].keys[:][j]
			}
		}
	}
	for i = 0; i < refs_n; i++ {
		for j = 0; uint32(j) < 1024/8; j++ {
			if refs[i].auth[:][j] != nil {
				_ = refs[i].auth[:][j]
			}
		}
	}
	return
}

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
