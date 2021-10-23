//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "unicode"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// soin_cmd - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/soin.c:10
func soin_cmd(s []byte) int32 {
	// soin: inline troff .so requests
	var path []byte = make([]byte, 2048)
	var d []byte = path
	if int32(s[0]) != int32('.') || int32(s[1]) != int32('s') || int32(s[2]) != int32('o') || int32(s[3]) != int32(' ') {
		return 1
	}
	s = s[0+3:]
	for int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
		s = s[0+1:]
	}
	if int32(s[0]) == int32('"') {
		s = s[0+1:]
		for int32(s[0]) != 0 && int32(s[0]) != int32('\n') && int32(s[0]) != int32('"') {
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
	} else {
		for int32(s[0]) != 0 && int32(s[0]) != int32(' ') && int32(s[0]) != int32('\n') {
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
	}
	d[0] = '\x00'
	return soin(path)
}

// soin - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/soin.c:31
func soin(path []byte) int32 {
	var fp *noarch.File = func() *noarch.File {
		if path != nil {
			return noarch.Fopen(path, []byte("r\x00"))
		}
		return noarch.Stdin
	}()
	var lineno int32 = 1
	var ln []byte = make([]byte, 2048)
	if fp == nil {
		noarch.Fprintf(noarch.Stderr, []byte("soin: cannot open <%s>\n\x00"), path)
		return 1
	}
	noarch.Printf([]byte(".lf %d %s\n\x00"), lineno, func() []byte {
		if path != nil {
			return path
		}
		return []byte("stdin\x00")
	}())
	for noarch.Fgets(ln, int32(2048), fp) != nil {
		lineno++
		if noarch.Not(soin_cmd(ln)) {
			noarch.Printf([]byte(".lf %d %s\n\x00"), lineno, func() []byte {
				if path != nil {
					return path
				}
				return []byte("stdin\x00")
			}())
		} else {
			noarch.Fputs(ln, noarch.Stdout)
		}
		if int32(ln[0]) == int32('.') && int32(ln[1]) == int32('l') && int32(ln[2]) == int32('f') {
			noarch.Sscanf(ln, []byte(".lf %d\x00"), c4goUnsafeConvert_int32(&lineno))
		}
	}
	if path != nil {
		noarch.Fclose(fp)
	}
	return 0
}

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/uroff/tmp/soin.c:55
func main() {
	defer noarch.AtexitRun()
	soin(nil)
	return
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[1000000]int32)(unsafe.Pointer(c4go_name))[:]
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
