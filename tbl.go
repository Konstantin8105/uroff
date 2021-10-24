//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "unicode"
import "reflect"
import "runtime"
import "os"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// uintptr_ - transpiled function from  t.h:8
// t..c : external declarations
type uintptr_ = uint32

// vlong - transpiled function from  t.h:9
type vlong = int64

// colstr - transpiled function from  t.h:54
// Do NOT make MAXCOL bigger with adjusting nregs[] in tr.c
type colstr struct {
	col  []byte
	rcol []byte
}

// expflg - transpiled function from  t0.c:4
// t0.c: storage allocation
var expflg int32

// ctrflg - transpiled function from  t0.c:5
var ctrflg int32

// boxflg - transpiled function from  t0.c:6
var boxflg int32

// dboxflg - transpiled function from  t0.c:7
var dboxflg int32

// tab - transpiled function from  t0.c:8
var tab int32 = int32('\t')

// linsize - transpiled function from  t0.c:9
var linsize int32

// pr1403 - transpiled function from  t0.c:10
var pr1403 int32

// delim1 - transpiled function from  t0.c:11
var delim1 int32

// delim2 - transpiled function from  t0.c:11
var delim2 int32

// evenflg - transpiled function from  t0.c:12
var evenflg int32

// evenup - transpiled function from  t0.c:13
var evenup []int32

// F1 - transpiled function from  t0.c:14
var F1 int32

// F2 - transpiled function from  t0.c:15
var F2 int32

// allflg - transpiled function from  t0.c:16
var allflg int32

// leftover - transpiled function from  t0.c:17
var leftover []byte

// textflg - transpiled function from  t0.c:18
var textflg int32

// left1flg - transpiled function from  t0.c:19
var left1flg int32

// rightl - transpiled function from  t0.c:20
var rightl int32

// cstore - transpiled function from  t0.c:21
var cstore []byte

// cspace - transpiled function from  t0.c:21
var cspace []byte

// last - transpiled function from  t0.c:22
var last []byte

// table - transpiled function from  t0.c:23
var table [][]colstr = make([][]colstr, 250)

// stynum - transpiled function from  t0.c:24
var stynum []int32 = make([]int32, 251)

// fullbot - transpiled function from  t0.c:25
var fullbot []int32 = make([]int32, 250)

// instead - transpiled function from  t0.c:26
var instead [][]byte = make([][]byte, 250)

// linestop - transpiled function from  t0.c:27
var linestop []int32 = make([]int32, 250)

// style - transpiled function from  t0.c:28
var style [][]int32 = make([][]int32, 44)

// font - transpiled function from  t0.c:29
var font [][][]byte = make([][][]byte, 44)

// csize - transpiled function from  t0.c:30
var csize [][][]byte = make([][][]byte, 44)

// vsize - transpiled function from  t0.c:31
var vsize [][][]byte = make([][][]byte, 44)

// lefline - transpiled function from  t0.c:32
var lefline [][]int32 = make([][]int32, 44)

// cll - transpiled function from  t0.c:33
var cll [][]byte = make([][]byte, 10)

// flags - transpiled function from  t0.c:34
var flags [][]int32 = make([][]int32, 44)

// qcol - transpiled function from  t0.c:35
var qcol int32

// doubled - transpiled function from  t0.c:36
var doubled []int32

// acase - transpiled function from  t0.c:36
var acase []int32

// topat - transpiled function from  t0.c:36
var topat []int32

// nslin - transpiled function from  t0.c:37
var nslin int32

// nclin - transpiled function from  t0.c:37
var nclin int32

// sep - transpiled function from  t0.c:38
var sep []int32

// used - transpiled function from  t0.c:39
var used []int32

// lused - transpiled function from  t0.c:39
var lused []int32

// rused - transpiled function from  t0.c:39
var rused []int32

// nlin - transpiled function from  t0.c:40
var nlin int32

// ncol - transpiled function from  t0.c:40
var ncol int32

// iline - transpiled function from  t0.c:41
var iline int32 = 1

// ifile - transpiled function from  t0.c:42
var ifile []byte = []byte("Input\x00")

// texname - transpiled function from  t0.c:43
var texname int32 = int32('a')

// texct - transpiled function from  t0.c:44
var texct int32

// texstr - transpiled function from  t0.c:45
var texstr []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYXZ0123456789\x00")

// linstart - transpiled function from  t0.c:46
var linstart int32

// exstore - transpiled function from  t0.c:47
var exstore []byte

// exlim - transpiled function from  t0.c:47
var exlim []byte

// exspace - transpiled function from  t0.c:47
var exspace []byte

// tabin - transpiled function from  t0.c:48
//= stdin
var tabin *noarch.File

// tabout - transpiled function from  t0.c:49
// = stdout
var tabout *noarch.File

// main - transpiled function from  t1.c:7
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	// t1.c: main control and input switching
	var line []byte = make([]byte, 5120)
	tabout = noarch.Stdout
	setinp(argc, argv)
	for gets1(line, int32(5120)) != nil {
		noarch.Fprintf(tabout, []byte("%s\n\x00"), line)
		if prefix([]byte(".TS\x00"), line) != 0 {
			tableput()
		}
	}
	if tabin != nil && (int64(uintptr(unsafe.Pointer(tabin)))/int64(8)-int64(uintptr(unsafe.Pointer(noarch.Stdin)))/int64(8)) != 0 {
		noarch.Fclose(tabin)
	}
	return
}

// sargc - transpiled function from  t1.c:23
var sargc int32

// sargv - transpiled function from  t1.c:24
var sargv [][]byte

// setinp - transpiled function from  t1.c:26
func setinp(argc int32, argv [][]byte) {
	sargc = argc
	sargv = argv
	sargc--
	sargv = sargv[0+1:]
	if sargc == 0 || swapin() == 0 {
		tabin = noarch.Stdin
	}
}

// swapin - transpiled function from  t1.c:38
func swapin() int32 {
	var name []byte
	for sargc > 0 && int32((sargv[0])[0]) == int32('-') {
		if match([]byte("-ms\x00"), sargv[0]) != 0 {
			sargv[0] = []byte("/sys/lib/tmac/tmac.s\x00")
			break
		}
		if match([]byte("-mm\x00"), sargv[0]) != 0 {
			sargv[0] = []byte("/sys/lib/tmac/tmac.m\x00")
			break
		}
		if match([]byte("-TX\x00"), sargv[0]) != 0 {
			pr1403 = 1
		}
		if match([]byte("-\x00"), sargv[0]) != 0 {
			break
		}
		sargc--
		sargv = sargv[0+1:]
	}
	if sargc <= 0 {
		return 0
	}
	if tabin != nil && (int64(uintptr(unsafe.Pointer(tabin)))/int64(8)-int64(uintptr(unsafe.Pointer(noarch.Stdin)))/int64(8)) != 0 {
		// file closing is done by GCOS troff preprocessor
		noarch.Fclose(tabin)
	}
	ifile = sargv[0]
	name = ifile
	if match(ifile, []byte("-\x00")) != 0 {
		tabin = noarch.Stdin
	} else {
		tabin = noarch.Fopen(ifile, []byte("r\x00"))
	}
	iline = 1
	noarch.Fprintf(tabout, []byte(".ds f. %s\n\x00"), ifile)
	noarch.Fprintf(tabout, []byte(".lf %d %s\n\x00"), iline, name)
	if tabin == nil {
		error_([]byte("Can't open file\x00"))
	}
	sargc--
	sargv = sargv[0+1:]
	return 1
}

// tableput - transpiled function from  t2.c:3
func tableput() {
	// t2.c:  subroutine sequencing for one table
	saveline()
	savefill()
	ifdivert()
	cleanfc()
	getcomm()
	getspec()
	gettbl()
	getstop()
	checkuse()
	choochar()
	maktab()
	runout()
	release()
	rstofill()
	endoff()
	freearr()
	restline()
}

// optstr - transpiled function from  t3.c:5
// t3.c: interpret commands affecting whole table
type optstr struct {
	optnam []byte
	optadd []int32
}

// options - transpiled function from  t3.c:5
var options []optstr = []optstr{{[]byte("expand\x00"), c4goUnsafeConvert_int32(&expflg)}, {[]byte("EXPAND\x00"), c4goUnsafeConvert_int32(&expflg)}, {[]byte("center\x00"), c4goUnsafeConvert_int32(&ctrflg)}, {[]byte("CENTER\x00"), c4goUnsafeConvert_int32(&ctrflg)}, {[]byte("box\x00"), c4goUnsafeConvert_int32(&boxflg)}, {[]byte("BOX\x00"), c4goUnsafeConvert_int32(&boxflg)}, {[]byte("allbox\x00"), c4goUnsafeConvert_int32(&allflg)}, {[]byte("ALLBOX\x00"), c4goUnsafeConvert_int32(&allflg)}, {[]byte("doublebox\x00"), c4goUnsafeConvert_int32(&dboxflg)}, {[]byte("DOUBLEBOX\x00"), c4goUnsafeConvert_int32(&dboxflg)}, {[]byte("frame\x00"), c4goUnsafeConvert_int32(&boxflg)}, {[]byte("FRAME\x00"), c4goUnsafeConvert_int32(&boxflg)}, {[]byte("doubleframe\x00"), c4goUnsafeConvert_int32(&dboxflg)}, {[]byte("DOUBLEFRAME\x00"), c4goUnsafeConvert_int32(&dboxflg)}, {[]byte("tab\x00"), c4goUnsafeConvert_int32(&tab)}, {[]byte("TAB\x00"), c4goUnsafeConvert_int32(&tab)}, {[]byte("linesize\x00"), c4goUnsafeConvert_int32(&linsize)}, {[]byte("LINESIZE\x00"), c4goUnsafeConvert_int32(&linsize)}, {[]byte("delim\x00"), c4goUnsafeConvert_int32(&delim1)}, {[]byte("DELIM\x00"), c4goUnsafeConvert_int32(&delim1)}, {nil, nil}}

// getcomm - transpiled function from  t3.c:32
func getcomm() {
	var line []byte = make([]byte, 200)
	var cp []byte
	var nb []byte = make([]byte, 25)
	var t []byte
	var lp []optstr
	var c int32
	var ci int32
	var found int32
	for lp = options; lp[0].optnam != nil; func() []optstr {
		tempVarUnary := lp
		defer func() {
			lp = lp[0+1:]
		}()
		return tempVarUnary
	}() {
		lp[0].optadd[0] = 0
	}
	texname = int32(texstr[(func() int32 {
		texct = 0
		return texct
	}())])
	tab = int32('\t')
	noarch.Fprintf(tabout, []byte(".nr %d \\n(.s\n\x00"), 33)
	gets1(line, int32(200))
	if len(noarch.Strchr(line, int32(';'))) == 0 {
		// see if this is a command line
		backrest(line)
		return
	}
	for cp = line; (func() int32 {
		c = int32(cp[0])
		return c
	}()) != int32(';'); func() []byte {
		tempVarUnary := cp
		defer func() {
			cp = cp[0+1:]
		}()
		return tempVarUnary
	}() {
		if noarch.Not(letter(c)) {
			continue
		}
		found = 0
		for lp = options; lp[0].optadd != nil; func() []optstr {
			tempVarUnary := lp
			defer func() {
				lp = lp[0+1:]
			}()
			return tempVarUnary
		}() {
			if prefix(lp[0].optnam, cp) != 0 {
				lp[0].optadd[0] = 1
				cp = cp[0+noarch.Strlen(lp[0].optnam):]
				if letter(int32(cp[0])) != 0 {
					error_([]byte("Misspelled global option\x00"))
				}
				for int32(cp[0]) == int32(' ') {
					cp = cp[0+1:]
				}
				t = nb
				if int32(cp[0]) == int32('(') {
					for (func() int32 {
						ci = int32((func() []byte {
							cp = cp[0+1:]
							return cp
						}())[0])
						return ci
					}()) != int32(')') {
						(func() []byte {
							defer func() {
								t = t[0+1:]
							}()
							return t
						}())[0] = byte(ci)
					}
				} else {
					cp = c4goPointerArithByteSlice(cp, int(-1))
				}
				(func() []byte {
					defer func() {
						t = t[0+1:]
					}()
					return t
				}())[0] = byte(0)
				t[0] = byte(0)
				if (int64(uintptr(unsafe.Pointer(&lp[0].optadd[0])))/int64(4) - func() int64 {
					c4go_temp_name := c4goUnsafeConvert_int32(&tab)
					return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
				}()) == 0 {
					if nb[0] != 0 {
						lp[0].optadd[0] = int32(nb[0])
					}
				}
				if (int64(uintptr(unsafe.Pointer(&lp[0].optadd[0])))/int64(4) - func() int64 {
					c4go_temp_name := c4goUnsafeConvert_int32(&linsize)
					return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
				}()) == 0 {
					noarch.Fprintf(tabout, []byte(".nr %d %s\n\x00"), 33, nb)
				}
				if (int64(uintptr(unsafe.Pointer(&lp[0].optadd[0])))/int64(4) - func() int64 {
					c4go_temp_name := c4goUnsafeConvert_int32(&delim1)
					return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
				}()) == 0 {
					delim1 = int32(nb[0])
					delim2 = int32(nb[1])
				}
				found = 1
				break
			}
		}
		if noarch.Not(found) {
			error_([]byte("Illegal option\x00"))
		}
	}
	cp = cp[0+1:]
	backrest(cp)
}

// backrest - transpiled function from  t3.c:91
func backrest(cp []byte) {
	var s []byte
	for s = cp; s[0] != 0; func() []byte {
		tempVarUnary := s
		defer func() {
			s = s[0+1:]
		}()
		return tempVarUnary
	}() {
	}
	un1getc(int32('\n'))
	for (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&cp[0])))/int64(1)) > 0 {
		un1getc(int32((func() []byte {
			s = c4goPointerArithByteSlice(s, int(-1))
			return s
		}())[0]))
	}
}

// oncol - transpiled function from  t4.c:4
// t4.c: read table specification
var oncol int32

// getspec - transpiled function from  t4.c:6
func getspec() {
	var icol int32
	var i int32
	// must allow one extra for line at right
	qcol = findcol() + 1
	garray(qcol)
	c4goPointerArithInt32Slice(sep, int(-1))[0] = -1
	for icol = 0; icol < qcol; icol++ {
		sep[icol] = -1
		evenup[icol] = 0
		cll[icol][0] = byte(0)
		for i = 0; i < 44; i++ {
			csize[icol][i][0] = byte(0)
			vsize[icol][i][0] = byte(0)
			lefline[icol][i] = 0
			font[icol][i][0] = byte(lefline[icol][i])
			flags[icol][i] = 0
			style[icol][i] = int32('l')
		}
	}
	for i = 0; i < 44; i++ {
		// fixes sample55 looping
		lefline[qcol][i] = 0
	}
	ncol = 0
	nclin = ncol
	oncol = 0
	rightl = 0
	left1flg = rightl
	readspec()
	noarch.Fprintf(tabout, []byte(".rm\x00"))
	for i = 0; i < ncol; i++ {
		noarch.Fprintf(tabout, []byte(" %2s\x00"), reg(i, 2))
	}
	noarch.Fprintf(tabout, []byte("\n\x00"))
}

// readspec - transpiled function from  t4.c:39
func readspec() {
	var icol int32
	var c int32
	var sawchar int32
	var stopc int32
	var i int32
	var sn []byte = make([]byte, 10)
	var snp []byte
	var temp []byte
	icol = 0
	sawchar = icol
	for (func() int32 {
		c = get1char()
		return c
	}()) != 0 {
		switch c {
		default:
			if c != tab {
				var buf []byte = make([]byte, 64)
				noarch.Sprintf(buf, []byte("bad table specification character %c\x00"), c)
				error_(buf)
			}
			fallthrough
		case ' ':
			// note this is also case tab
			continue
			fallthrough
		case '\n':
			if sawchar == 0 {
				continue
			}
			fallthrough
		case ',':
			fallthrough
		case '.':
			// end of table specification
			ncol = max(ncol, icol)
			if lefline[ncol][nclin] > 0 {
				ncol++
				rightl++
			}
			if sawchar != 0 {
				nclin++
			}
			if nclin >= 44 {
				error_([]byte("too many lines in specification\x00"))
			}
			icol = 0
			if ncol == 0 || nclin == 0 {
				error_([]byte("no specification\x00"))
			}
			if c == int32('.') {
				for (func() int32 {
					c = get1char()
					return c
				}()) != 0 && c != int32('\n') {
					if c != int32(' ') && c != int32('\t') {
						error_([]byte("dot not last character on format line\x00"))
					}
				}
				{
					// fix up sep - default is 3 except at edge
					for icol = 0; icol < ncol; icol++ {
						if sep[icol] < 0 {
							if icol+1 < ncol {
								sep[icol] = 3
							} else {
								sep[icol] = 2
							}
						}
					}
				}
				if oncol == 0 {
					oncol = ncol
				} else if oncol+2 < ncol {
					error_([]byte("tried to widen table in T&, not allowed\x00"))
				}
				return
			}
			sawchar = 0
			continue
			fallthrough
		case 'C':
			fallthrough
		case 'S':
			fallthrough
		case 'R':
			fallthrough
		case 'N':
			fallthrough
		case 'L':
			fallthrough
		case 'A':
			c += int32('a' - 'A')
			fallthrough
		case '_':
			if c == int32('_') {
				c = int32('-')
			}
			fallthrough
		case '=':
			fallthrough
		case '-':
			fallthrough
		case '^':
			fallthrough
		case 'c':
			fallthrough
		case 's':
			fallthrough
		case 'n':
			fallthrough
		case 'r':
			fallthrough
		case 'l':
			fallthrough
		case 'a':
			style[icol][nclin] = c
			if c == int32('s') && icol <= 0 {
				error_([]byte("first column can not be S-type\x00"))
			}
			if c == int32('s') && style[icol-1][nclin] == int32('a') {
				noarch.Fprintf(tabout, []byte(".tm warning: can't span a-type cols, changed to l\n\x00"))
				style[icol-1][nclin] = int32('l')
			}
			if c == int32('s') && style[icol-1][nclin] == int32('n') {
				noarch.Fprintf(tabout, []byte(".tm warning: can't span n-type cols, changed to c\n\x00"))
				style[icol-1][nclin] = int32('c')
			}
			icol++
			if c == int32('^') && nclin <= 0 {
				error_([]byte("first row can not contain vertical span\x00"))
			}
			if icol > qcol {
				error_([]byte("too many columns in table\x00"))
			}
			sawchar = 1
			continue
			fallthrough
		case 'b':
			fallthrough
		case 'i':
			c += int32('A' - 'a')
			fallthrough
		case 'B':
			fallthrough
		case 'I':
			if icol == 0 {
				continue
			}
			snp = font[icol-1][nclin]
			snp[0] = byte(func() int32 {
				if c == int32('I') {
					return int32('2')
				}
				return int32('3')
			}())
			snp[1] = byte(0)
			continue
			fallthrough
		case 't':
			fallthrough
		case 'T':
			if icol > 0 {
				flags[icol-1][nclin] |= 4
			}
			continue
			fallthrough
		case 'd':
			fallthrough
		case 'D':
			if icol > 0 {
				flags[icol-1][nclin] |= 8
			}
			continue
			fallthrough
		case 'f':
			fallthrough
		case 'F':
			if icol == 0 {
				continue
			}
			snp = font[icol-1][nclin]
			stopc = 0
			snp[1] = byte(stopc)
			snp[0] = snp[1]
			for i = 0; i < 2; i++ {
				c = get1char()
				if i == 0 && c == int32('(') {
					stopc = int32(')')
					c = get1char()
				}
				if c == 0 {
					break
				}
				if c == stopc {
					stopc = 0
					break
				}
				if stopc == 0 {
					if c == int32(' ') || c == tab {
						break
					}
				}
				if c == int32('\n') || c == int32('|') {
					un1getc(c)
					break
				}
				snp[i] = byte(c)
				if c >= int32('0') && c <= int32('9') {
					break
				}
			}
			if stopc != 0 {
				if get1char() != stopc {
					error_([]byte("Nonterminated font name\x00"))
				}
			}
			continue
			fallthrough
		case 'P':
			fallthrough
		case 'p':
			if icol <= 0 {
				continue
			}
			snp = csize[icol-1][nclin]
			temp = snp
			for (func() int32 {
				c = get1char()
				return c
			}()) != 0 {
				if c == int32(' ') || c == tab || c == int32('\n') {
					break
				}
				if c == int32('-') || c == int32('+') {
					if (int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&temp[0])))/int64(1)) > 0 {
						break
					} else {
						(func() []byte {
							defer func() {
								snp = snp[0+1:]
							}()
							return snp
						}())[0] = byte(c)
					}
				} else if digit(c) != 0 {
					(func() []byte {
						defer func() {
							snp = snp[0+1:]
						}()
						return snp
					}())[0] = byte(c)
				} else {
					break
				}
				if int32((int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&temp[0])))/int64(1))) > 4 {
					error_([]byte("point size too large\x00"))
				}
			}
			snp[0] = byte(0)
			if noarch.Atoi(temp) > 36 {
				error_([]byte("point size unreasonable\x00"))
			}
			un1getc(c)
			continue
			fallthrough
		case 'V':
			fallthrough
		case 'v':
			if icol <= 0 {
				continue
			}
			snp = vsize[icol-1][nclin]
			temp = snp
			for (func() int32 {
				c = get1char()
				return c
			}()) != 0 {
				if c == int32(' ') || c == tab || c == int32('\n') {
					break
				}
				if c == int32('-') || c == int32('+') {
					if (int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&temp[0])))/int64(1)) > 0 {
						break
					} else {
						(func() []byte {
							defer func() {
								snp = snp[0+1:]
							}()
							return snp
						}())[0] = byte(c)
					}
				} else if digit(c) != 0 {
					(func() []byte {
						defer func() {
							snp = snp[0+1:]
						}()
						return snp
					}())[0] = byte(c)
				} else {
					break
				}
				if int32((int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&temp[0])))/int64(1))) > 4 {
					error_([]byte("vertical spacing value too large\x00"))
				}
			}
			snp[0] = byte(0)
			un1getc(c)
			continue
			fallthrough
		case 'w':
			fallthrough
		case 'W':
			snp = cll[icol-1]
			stopc = 0
			for (func() int32 {
				c = get1char()
				return c
			}()) != 0 {
				if (int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&cll[icol-1])))/int64(1)) == 0 && c == int32('(') {
					stopc = int32(')')
					continue
				}
				if noarch.Not(stopc) && (c > int32('9') || c < int32('0')) {
					break
				}
				if stopc != 0 && c == stopc {
					break
				}
				(func() []byte {
					defer func() {
						snp = snp[0+1:]
					}()
					return snp
				}())[0] = byte(c)
			}
			snp[0] = byte(0)
			if int32((int64(uintptr(unsafe.Pointer(&snp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&cll[icol-1])))/int64(1))) > 10 {
				error_([]byte("column width too long\x00"))
			}
			if noarch.Not(stopc) {
				un1getc(c)
			}
			continue
			fallthrough
		case 'e':
			fallthrough
		case 'E':
			if icol < 1 {
				continue
			}
			evenup[icol-1] = 1
			evenflg = 1
			continue
			fallthrough
		case 'z':
			fallthrough
		case 'Z':
			if icol < 1 {
				// zero width-ignre width this item
				continue
			}
			flags[icol-1][nclin] |= 1
			continue
			fallthrough
		case 'u':
			fallthrough
		case 'U':
			if icol < 1 {
				// half line up
				continue
			}
			flags[icol-1][nclin] |= 2
			continue
			fallthrough
		case '0':
			fallthrough
		case '1':
			fallthrough
		case '2':
			fallthrough
		case '3':
			fallthrough
		case '4':
			fallthrough
		case '5':
			fallthrough
		case '6':
			fallthrough
		case '7':
			fallthrough
		case '8':
			fallthrough
		case '9':
			sn[0] = byte(c)
			snp = sn[0+1:]
			for digit(int32((func() byte {
				c = get1char()
				(func() []byte {
					defer func() {
						snp = snp[0+1:]
					}()
					return snp
				}())[0] = byte(c)
				return (func() []byte {
					defer func() {
						snp = snp[0+1:]
					}()
					return snp
				}())[0]
			}()))) != 0 {
			}
			un1getc(c)
			sep[icol-1] = max(sep[icol-1], numb(sn))
			continue
			fallthrough
		case '|':
			lefline[icol][nclin]++
			if icol == 0 {
				left1flg = 1
			}
			continue
		}
	}
	error_([]byte("EOF reading table specification\x00"))
}

// findcol - transpiled function from  t4.c:296
func findcol() int32 {
	// this counts the number of columns and then puts the line back
	var s []byte
	var line []byte = make([]byte, 202)
	var p []byte
	var c int32
	var n int32
	var inpar int32
	for (func() int32 {
		c = get1char()
		return c
	}()) != 0 && c == int32(' ') {
	}
	if c != int32('\n') {
		un1getc(c)
	}
	for s = line; (func() byte {
		c = get1char()
		s[0] = byte(c)
		return s[0]
	}()) != 0; func() []byte {
		tempVarUnary := s
		defer func() {
			s = s[0+1:]
		}()
		return tempVarUnary
	}() {
		if c == int32(')') {
			inpar = 0
		}
		if inpar != 0 {
			continue
		}
		if c == int32('\n') || c == 0 || c == int32('.') || c == int32(',') {
			break
		} else if c == int32('(') {
			inpar = 1
		} else if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&line[0+200])))/int64(1)) >= 0 {
			error_([]byte("too long spec line\x00"))
		}
	}
	for p = line; (int64(uintptr(unsafe.Pointer(&p[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)) < 0; func() []byte {
		tempVarUnary := p
		defer func() {
			p = p[0+1:]
		}()
		return tempVarUnary
	}() {
		switch int32(p[0]) {
		case 'l':
			fallthrough
		case 'r':
			fallthrough
		case 'c':
			fallthrough
		case 'n':
			fallthrough
		case 'a':
			fallthrough
		case 's':
			fallthrough
		case 'L':
			fallthrough
		case 'R':
			fallthrough
		case 'C':
			fallthrough
		case 'N':
			fallthrough
		case 'A':
			fallthrough
		case 'S':
			fallthrough
		case '-':
			fallthrough
		case '=':
			fallthrough
		case '_':
			n++
		}
	}
	for (int64(uintptr(unsafe.Pointer(&p[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&line[0])))/int64(1)) >= 0 {
		un1getc(int32((func() []byte {
			defer func() {
				p = c4goPointerArithByteSlice(p, int(-1))
			}()
			return p
		}())[0]))
	}
	return n
}

// garray - transpiled function from  t4.c:345
func garray(qcol int32) {
	style = (*[10000][]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*qcol, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	evenup = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	//+1 for sample55 loop - others may need it too
	lefline = (*[10000][]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*(qcol+1), int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	font = (*[10000][][]byte)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*qcol, 2)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	csize = (*[10000][][]byte)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*qcol, 4)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	vsize = (*[10000][][]byte)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*qcol, 4)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	flags = (*[10000][]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(44*qcol, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	cll = (*[10000][]byte)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol, 10)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	sep = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	// sep[-1] must be legal
	sep = sep[0+1:]
	used = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	lused = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	rused = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	doubled = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	acase = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	topat = (*[10000]int32)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := getcore(qcol+1, int32(4))
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
}

// getcore - transpiled function from  t4.c:367
func getcore(a int32, b int32) []byte {
	var x []byte
	x = make([]byte, uint32(a)*uint32(b))
	if len(x) == 0 {
		error_([]byte("Couldn't get memory\x00"))
	}
	return x
}

// freearr - transpiled function from  t4.c:378
func freearr() {
	_ = style
	_ = evenup
	_ = lefline
	_ = flags
	_ = font
	_ = csize
	_ = vsize
	_ = cll
	_ = func() []int32 {
		tempVarUnary := sep
		defer func() {
			sep = c4goPointerArithInt32Slice(sep, int(-1))
		}()
		return tempVarUnary
	}()
	// netnews says this should be --sep because incremented earlier!
	_ = used
	_ = lused
	_ = rused
	_ = doubled
	_ = acase
	_ = topat
}

// gettbl - transpiled function from  t5.c:4
func gettbl() {
	// t5.c: read data for table
	var icol int32
	var ch int32
	cspace = chspace()
	cstore = cspace
	textflg = 0
	{
		nslin = 0
		for nlin = nslin; gets1(cstore, 2000-int32((int64(uintptr(unsafe.Pointer(&cstore[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&cspace[0])))/int64(1)))) != nil; nlin++ {
			stynum[nlin] = nslin
			if prefix([]byte(".TE\x00"), cstore) != 0 {
				leftover = nil
				break
			}
			if prefix([]byte(".TC\x00"), cstore) != 0 || prefix([]byte(".T&\x00"), cstore) != 0 {
				readspec()
				nslin++
			}
			if nlin >= 250 {
				leftover = cstore
				break
			}
			fullbot[nlin] = 0
			if int32(cstore[0]) == int32('.') && noarch.Not(int32(((__ctype_b_loc())[0])[int32(cstore[1])])&int32(uint16(noarch.ISdigit))) {
				instead[nlin] = cstore
				for (func() []byte {
					defer func() {
						cstore = cstore[0+1:]
					}()
					return cstore
				}())[0] != 0 {
				}
				continue
			} else {
				instead[nlin] = nil
			}
			if nodata(nlin) != 0 {
				if (func() int32 {
					ch = oneh(nlin)
					return ch
				}()) != 0 {
					fullbot[nlin] = ch
				}
				table[nlin] = (*[10000]colstr)(unsafe.Pointer(uintptr(func() int64 {
					c4go_temp_name := alocv(int32(uint32(ncol+2) * 32))
					return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
				}())))[:]
				for icol = 0; icol < ncol; icol++ {
					table[nlin][icol].rcol = []byte("\x00")
					table[nlin][icol].col = []byte("\x00")
				}
				nlin++
				nslin++
				fullbot[nlin] = 0
				instead[nlin] = nil
			}
			table[nlin] = (*[10000]colstr)(unsafe.Pointer(uintptr(func() int64 {
				c4go_temp_name := alocv(int32(uint32(ncol+2) * 32))
				return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
			}())))[:]
			if int32(cstore[1]) == 0 {
				switch int32(cstore[0]) {
				case '_':
					fullbot[nlin] = int32('-')
					continue
					fallthrough
				case '=':
					fullbot[nlin] = int32('=')
					continue
				}
			}
			stynum[nlin] = nslin
			nslin = min(nslin+1, nclin-1)
			for icol = 0; icol < ncol; icol++ {
				table[nlin][icol].col = cstore
				table[nlin][icol].rcol = nil
				ch = 1
				if match(cstore, []byte("T{\x00")) != 0 {
					// text follows
					table[nlin][icol].col = []byte(gettext_tbl(cstore, nlin, icol, font[icol][stynum[nlin]], csize[icol][stynum[nlin]]))
				} else {
					for ; (func() int32 {
						ch = int32(cstore[0])
						return ch
					}()) != int32('\x00') && ch != tab; func() []byte {
						tempVarUnary := cstore
						defer func() {
							cstore = cstore[0+1:]
						}()
						return tempVarUnary
					}() {
					}
					(func() []byte {
						defer func() {
							cstore = cstore[0+1:]
						}()
						return cstore
					}())[0] = '\x00'
					switch ctype(nlin, icol) {
					case 'n':
						// numerical or alpha, subcol
						table[nlin][icol].rcol = maknew(table[nlin][icol].col)
					case 'a':
						table[nlin][icol].rcol = table[nlin][icol].col
						table[nlin][icol].col = []byte("\x00")
						break
					}
				}
				for ctype(nlin, icol+1) == int32('s') {
					// spanning
					table[nlin][func() int32 {
						icol++
						return icol
					}()].col = []byte("\x00")
				}
				if ch == int32('\x00') {
					break
				}
			}
			for func() int32 {
				icol++
				return icol
			}() < ncol+2 {
				table[nlin][icol].col = []byte("\x00")
				table[nlin][icol].rcol = nil
			}
			for int32(cstore[0]) != int32('\x00') {
				cstore = cstore[0+1:]
			}
			if int32((int64(uintptr(unsafe.Pointer(&cstore[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&cspace[0])))/int64(1)))+300 > 2000 {
				cspace = chspace()
				cstore = cspace
			}
		}
	}
	last = cstore
	permute()
	if textflg != 0 {
		untext()
	}
}

// nodata - transpiled function from  t5.c:103
func nodata(il int32) int32 {
	var c int32
	for c = 0; c < ncol; c++ {
		switch ctype(il, c) {
		case 'c':
			fallthrough
		case 'n':
			fallthrough
		case 'r':
			fallthrough
		case 'l':
			fallthrough
		case 's':
			fallthrough
		case 'a':
			return 0
		}
	}
	return 1
}

// oneh - transpiled function from  t5.c:123
func oneh(lin int32) int32 {
	var k int32
	var icol int32
	k = ctype(lin, 0)
	for icol = 1; icol < ncol; icol++ {
		if k != ctype(lin, icol) {
			return 0
		}
	}
	return k
}

// permute - transpiled function from  t5.c:139
func permute() {
	var irow int32
	var jcol int32
	var is int32
	var start []byte
	var strig []byte
	for jcol = 0; jcol < ncol; jcol++ {
		for irow = 1; irow < nlin; irow++ {
			if vspand(irow, jcol, 0) != 0 {
				is = prev(irow)
				if is < 0 {
					error_([]byte("Vertical spanning in first row not allowed\x00"))
				}
				start = table[is][jcol].col
				strig = table[is][jcol].rcol
				for irow < nlin && vspand(irow, jcol, 0) != 0 {
					irow++
				}
				table[func() int32 {
					irow--
					return irow
				}()][jcol].col = start
				table[irow][jcol].rcol = strig
				for is < irow {
					table[is][jcol].rcol = nil
					table[is][jcol].col = []byte("\\^\x00")
					is = next(is)
				}
			}
		}
	}
}

// vspand - transpiled function from  t5.c:168
func vspand(ir int32, ij int32, ifform int32) int32 {
	if ir < 0 {
		return 0
	}
	if ir >= nlin {
		return 0
	}
	if instead[ir] != nil {
		return 0
	}
	if ifform == 0 && ctype(ir, ij) == int32('^') {
		return 1
	}
	if len(table[ir][ij].rcol) != 0 {
		return 0
	}
	if fullbot[ir] != 0 {
		return 0
	}
	return vspen(table[ir][ij].col)
}

// vspen - transpiled function from  t5.c:187
func vspen(s []byte) int32 {
	if len(s) == 0 {
		return 0
	}
	if noarch.Not(point(s)) {
		return 0
	}
	return match(s, []byte("\\^\x00"))
}

// maktab - transpiled function from  t6.c:9
func maktab() {
	// t6.c: compute tab stops
	// define the tab stops of the table
	var icol int32
	var ilin int32
	var tsep int32
	var k int32
	var ik int32
	var vforml int32
	var il int32
	var s int32
	var text int32
	var ss []byte
	for icol = 0; icol < ncol; icol++ {
		acase[icol] = 0
		doubled[icol] = acase[icol]
		noarch.Fprintf(tabout, []byte(".nr %2s 0\n\x00"), reg(icol, 2))
		for text = 0; text < 2; text++ {
			if text != 0 {
				noarch.Fprintf(tabout, []byte(".%2s\n.rm %2s\n\x00"), reg(icol, 2), reg(icol, 2))
			}
			for ilin = 0; ilin < nlin; ilin++ {
				if instead[ilin] != nil || fullbot[ilin] != 0 {
					continue
				}
				vforml = ilin
				for il = prev(ilin); il >= 0 && vspen(table[il][icol].col) != 0; il = prev(il) {
					vforml = il
				}
				if fspan(vforml, icol) != 0 {
					continue
				}
				if filler(table[ilin][icol].col) != 0 {
					continue
				}
				if flags[icol][stynum[ilin]]&1 != 0 {
					continue
				}
				switch ctype(vforml, icol) {
				case 'a':
					acase[icol] = 1
					ss = table[ilin][icol].col
					s = int32(uint32((uintptr_(ss))))
					if s > 0 && s < 128 && text != 0 {
						if doubled[icol] == 0 {
							noarch.Fprintf(tabout, []byte(".nr %d 0\n.nr %d 0\n\x00"), 31, 32)
						}
						doubled[icol] = 1
						noarch.Fprintf(tabout, []byte(".if \\n(%c->\\n(%d .nr %d \\n(%c-\n\x00"), s, 32, 32, s)
					}
					fallthrough
				case 'n':
					if len(table[ilin][icol].rcol) != 0 {
						if doubled[icol] == 0 && text == 0 {
							noarch.Fprintf(tabout, []byte(".nr %d 0\n.nr %d 0\n\x00"), 31, 32)
						}
						doubled[icol] = 1
						if real_(func() []byte {
							ss = table[ilin][icol].col
							tempVar3 := &ss
							return *tempVar3
						}()) != 0 && noarch.Not(vspen(ss)) {
							s = int32(uint32((uintptr_(ss))))
							if (s > 0 && s < 128) != (text != 0) {
								continue
							}
							noarch.Fprintf(tabout, []byte(".nr %d \x00"), 38)
							wide(ss, font[icol][stynum[vforml]], csize[icol][stynum[vforml]])
							noarch.Fprintf(tabout, []byte("\n\x00"))
							noarch.Fprintf(tabout, []byte(".if \\n(%d<\\n(%d .nr %d \\n(%d\n\x00"), 31, 38, 31, 38)
						}
						if text == 0 && real_(func() []byte {
							ss = table[ilin][icol].rcol
							tempVar3 := &ss
							return *tempVar3
						}()) != 0 && noarch.Not(vspen(ss)) && noarch.Not(barent(ss)) {
							noarch.Fprintf(tabout, []byte(".nr %d \\w%c%s%c\n\x00"), 38, F1, ss, F1)
							noarch.Fprintf(tabout, []byte(".if \\n(%d<\\n(%d .nr %d \\n(%d\n\x00"), 32, 38, 32, 38)
						}
						continue
					}
					fallthrough
				case 'r':
					fallthrough
				case 'c':
					fallthrough
				case 'l':
					if real_(func() []byte {
						ss = table[ilin][icol].col
						tempVar3 := &ss
						return *tempVar3
					}()) != 0 && noarch.Not(vspen(ss)) {
						s = int32(uint32((uintptr_(ss))))
						if (s > 0 && s < 128) != (text != 0) {
							continue
						}
						noarch.Fprintf(tabout, []byte(".nr %d \x00"), 38)
						wide(ss, font[icol][stynum[vforml]], csize[icol][stynum[vforml]])
						noarch.Fprintf(tabout, []byte("\n\x00"))
						noarch.Fprintf(tabout, []byte(".if \\n(%2s<\\n(%d .nr %2s \\n(%d\n\x00"), reg(icol, 2), 38, reg(icol, 2), 38)
					}
				}
			}
		}
		if acase[icol] != 0 {
			noarch.Fprintf(tabout, []byte(".if \\n(%d>=\\n(%2s .nr %2s \\n(%du+2n\n\x00"), 32, reg(icol, 2), reg(icol, 2), 32)
		}
		if doubled[icol] != 0 {
			noarch.Fprintf(tabout, []byte(".nr %2s \\n(%d\n\x00"), reg(icol, 1), 31)
			noarch.Fprintf(tabout, []byte(".nr %d \\n(%2s+\\n(%d\n\x00"), 38, reg(icol, 1), 32)
			noarch.Fprintf(tabout, []byte(".if \\n(%d>\\n(%2s .nr %2s \\n(%d\n\x00"), 38, reg(icol, 2), reg(icol, 2), 38)
			noarch.Fprintf(tabout, []byte(".if \\n(%d<\\n(%2s .nr %2s +(\\n(%2s-\\n(%d)/2\n\x00"), 38, reg(icol, 2), reg(icol, 1), reg(icol, 2), 38)
		}
		if cll[icol][0] != 0 {
			noarch.Fprintf(tabout, []byte(".nr %d %sn\n\x00"), 38, cll[icol])
			noarch.Fprintf(tabout, []byte(".if \\n(%2s<\\n(%d .nr %2s \\n(%d\n\x00"), reg(icol, 2), 38, reg(icol, 2), 38)
		}
		for ilin = 0; ilin < nlin; ilin++ {
			if (func() int32 {
				k = lspan(ilin, icol)
				return k
			}()) != 0 {
				ss = table[ilin][icol-k].col
				if noarch.Not(real_(ss)) || barent(ss) != 0 || vspen(ss) != 0 {
					continue
				}
				noarch.Fprintf(tabout, []byte(".nr %d \x00"), 38)
				wide(table[ilin][icol-k].col, font[icol-k][stynum[ilin]], csize[icol-k][stynum[ilin]])
				for ik = k; ik >= 0; ik-- {
					noarch.Fprintf(tabout, []byte("-\\n(%2s\x00"), reg(icol-ik, 2))
					if noarch.Not(expflg) && ik > 0 {
						noarch.Fprintf(tabout, []byte("-%dn\x00"), sep[icol-ik])
					}
				}
				noarch.Fprintf(tabout, []byte("\n\x00"))
				noarch.Fprintf(tabout, []byte(".if \\n(%d>0 .nr %d \\n(%d/%d\n\x00"), 38, 38, 38, k)
				noarch.Fprintf(tabout, []byte(".if \\n(%d<0 .nr %d 0\n\x00"), 38, 38)
				for ik = 1; ik <= k; ik++ {
					if doubled[icol-k+ik] != 0 {
						noarch.Fprintf(tabout, []byte(".nr %2s +\\n(%d/2\n\x00"), reg(icol-k+ik, 1), 38)
					}
					noarch.Fprintf(tabout, []byte(".nr %2s +\\n(%d\n\x00"), reg(icol-k+ik, 2), 38)
				}
			}
		}
	}
	if textflg != 0 {
		untext()
	}
	if evenflg != 0 {
		// if even requested, make all columns widest width
		noarch.Fprintf(tabout, []byte(".nr %d 0\n\x00"), 38)
		for icol = 0; icol < ncol; icol++ {
			if evenup[icol] == 0 {
				continue
			}
			noarch.Fprintf(tabout, []byte(".if \\n(%2s>\\n(%d .nr %d \\n(%2s\n\x00"), reg(icol, 2), 38, 38, reg(icol, 2))
		}
		for icol = 0; icol < ncol; icol++ {
			if evenup[icol] == 0 {
				// if column not evened just retain old interval
				continue
			}
			if doubled[icol] != 0 {
				noarch.Fprintf(tabout, []byte(".nr %2s (100*\\n(%2s/\\n(%2s)*\\n(%d/100\n\x00"), reg(icol, 1), reg(icol, 1), reg(icol, 2), 38)
			}
			// that nonsense with the 100's and parens tries
			//       to avoid overflow while proportionally shifting
			//       the middle of the number
			noarch.Fprintf(tabout, []byte(".nr %2s \\n(%d\n\x00"), reg(icol, 2), 38)
		}
	}
	{
		icol = 0
		// now adjust for total table width
		for tsep = icol; icol < ncol; icol++ {
			tsep += sep[icol]
		}
	}
	if expflg != 0 {
		noarch.Fprintf(tabout, []byte(".nr %d 0\x00"), 38)
		for icol = 0; icol < ncol; icol++ {
			noarch.Fprintf(tabout, []byte("+\\n(%2s\x00"), reg(icol, 2))
		}
		noarch.Fprintf(tabout, []byte("\n\x00"))
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.l-\\n(%d\n\x00"), 38, 38)
		if boxflg != 0 || dboxflg != 0 || allflg != 0 {
			// tsep += 1;
			{
			}
		} else {
			tsep -= sep[ncol-1]
		}
		noarch.Fprintf(tabout, []byte(".nr %d \\n(%d/%d\n\x00"), 38, 38, tsep)
		noarch.Fprintf(tabout, []byte(".if \\n(%d<0 .nr %d 0\n\x00"), 38, 38)
	} else {
		noarch.Fprintf(tabout, []byte(".nr %d 1n\n\x00"), 38)
	}
	noarch.Fprintf(tabout, []byte(".nr %2s 0\n\x00"), reg(-1, 2))
	if boxflg != 0 || allflg != 0 || dboxflg != 0 || left1flg != 0 {
		tsep = 2
	} else {
		tsep = 0
	}
	if c4goPointerArithInt32Slice(sep, int(-1))[0] >= 0 {
		tsep = c4goPointerArithInt32Slice(sep, int(-1))[0]
	}
	for icol = 0; icol < ncol; icol++ {
		noarch.Fprintf(tabout, []byte(".nr %2s \\n(%2s+((%d*\\n(%d)/2)\n\x00"), reg(icol, 0), reg(icol-1, 2), tsep, 38)
		noarch.Fprintf(tabout, []byte(".nr %2s +\\n(%2s\n\x00"), reg(icol, 2), reg(icol, 0))
		if doubled[icol] != 0 {
			// the next line is last-ditch effort to avoid zero field width
			//fprintf(tabout, ".if \\n(%2s=0 .nr %2s 1\n",reg(icol,CMID), reg(icol,CMID));
			noarch.Fprintf(tabout, []byte(".nr %2s +\\n(%2s\n\x00"), reg(icol, 1), reg(icol, 0))
		}
		//  fprintf(tabout, ".if n .if \\n(%s%%24>0 .nr %s +12u\n",reg(icol,CMID), reg(icol,CMID));
		tsep = sep[icol] * 2
	}
	if rightl != 0 {
		noarch.Fprintf(tabout, []byte(".nr %s (\\n(%s+\\n(%s)/2\n\x00"), reg(ncol-1, 2), reg(ncol-1, 0), reg(ncol-2, 2))
	}
	noarch.Fprintf(tabout, []byte(".nr TW \\n(%2s\n\x00"), reg(ncol-1, 2))
	tsep = sep[ncol-1]
	if boxflg != 0 || allflg != 0 || dboxflg != 0 {
		noarch.Fprintf(tabout, []byte(".nr TW +((%d*\\n(%d)/2)\n\x00"), tsep, 38)
	}
	noarch.Fprintf(tabout, []byte(".if t .if (\\n(TW+\\n(.o)>7.65i .tm Table at line %d file %s is too wide - \\n(TW units\n\x00"), iline-1, ifile)
}

// wide - transpiled function from  t6.c:199
func wide(s []byte, fn []byte, size []byte) {
	if point(s) != 0 {
		noarch.Fprintf(tabout, []byte("\\w%c\x00"), F1)
		if int32(fn[0]) > 0 {
			putfont(fn)
		}
		if size[0] != 0 {
			putsize(size)
		}
		noarch.Fprintf(tabout, []byte("%s\x00"), s)
		if int32(fn[0]) > 0 {
			putfont([]byte("P\x00"))
		}
		if size[0] != 0 {
			putsize([]byte("0\x00"))
		}
		noarch.Fprintf(tabout, []byte("%c\x00"), F1)
	} else {
		noarch.Fprintf(tabout, []byte("\\n(%c-\x00"), int32(uint32((uintptr_(s)))))
	}
}

// filler - transpiled function from  t6.c:219
func filler(s []byte) int32 {
	return noarch.BoolToInt(point(s) != 0 && int32(s[0]) == int32('\\') && int32(s[1]) == int32('R'))
}

// runout - transpiled function from  t7.c:5
func runout() {
	// t7.c: control to write table entries
	var i int32
	if boxflg != 0 || allflg != 0 || dboxflg != 0 {
		need()
	}
	if ctrflg != 0 {
		noarch.Fprintf(tabout, []byte(".nr #I \\n(.i\n\x00"))
		noarch.Fprintf(tabout, []byte(".in +(\\n(.lu-\\n(TWu-\\n(.iu)/2u\n\x00"))
	}
	noarch.Fprintf(tabout, []byte(".fc %c %c\n\x00"), F1, F2)
	noarch.Fprintf(tabout, []byte(".nr #T 0-1\n\x00"))
	deftail()
	for i = 0; i < nlin; i++ {
		putline(i, i)
	}
	if leftover != nil {
		yetmore()
	}
	noarch.Fprintf(tabout, []byte(".fc\n\x00"))
	noarch.Fprintf(tabout, []byte(".nr T. 1\n\x00"))
	noarch.Fprintf(tabout, []byte(".T# 1\n\x00"))
	if ctrflg != 0 {
		noarch.Fprintf(tabout, []byte(".in \\n(#Iu\n\x00"))
	}
}

// runtabs - transpiled function from  t7.c:31
func runtabs(lform int32, ldata int32) {
	var c int32
	var ct int32
	var vforml int32
	var lf int32
	noarch.Fprintf(tabout, []byte(".ta \x00"))
	for c = 0; c < ncol; c++ {
		vforml = lform
		for lf = prev(lform); lf >= 0 && vspen(table[lf][c].col) != 0; lf = prev(lf) {
			vforml = lf
		}
		if fspan(vforml, c) != 0 {
			continue
		}
		switch func() int32 {
			ct = ctype(vforml, c)
			return ct
		}() {
		case 'n':
			fallthrough
		case 'a':
			if table[ldata][c].rcol != nil {
				if lused[c] != 0 {
					//Zero field width
					noarch.Fprintf(tabout, []byte("\\n(%2su \x00"), reg(c, 1))
				}
			}
			fallthrough
		case 'c':
			fallthrough
		case 'l':
			fallthrough
		case 'r':
			if func() int32 {
				if (ct == int32('a') || ct == int32('n')) && table[ldata][c].rcol != nil {
					return rused[c]
				}
				return used[c] + lused[c]
			}() != 0 {
				noarch.Fprintf(tabout, []byte("\\n(%2su \x00"), reg(c, 2))
			}
			continue
			fallthrough
		case 's':
			if lspan(lform, c) != 0 {
				noarch.Fprintf(tabout, []byte("\\n(%2su \x00"), reg(c, 2))
			}
			continue
		}
	}
	noarch.Fprintf(tabout, []byte("\n\x00"))
}

// ifline - transpiled function from  t7.c:65
func ifline(s []byte) int32 {
	if noarch.Not(point(s)) {
		return 0
	}
	if int32(s[0]) == int32('\\') {
		s = s[0+1:]
	}
	if s[1] != 0 {
		return 0
	}
	if int32(s[0]) == int32('_') {
		return int32('-')
	}
	if int32(s[0]) == int32('=') {
		return int32('=')
	}
	return 0
}

// need - transpiled function from  t7.c:82
func need() {
	var texlin int32
	var horlin int32
	var i int32
	{
		i = 0
		horlin = i
		for texlin = horlin; i < nlin; i++ {
			if fullbot[i] != 0 {
				horlin++
			} else if len(instead[i]) != 0 {
				continue
			} else {
				texlin++
			}
		}
	}
	noarch.Fprintf(tabout, []byte(".ne %dv+%dp\n\x00"), texlin, 2*horlin)
}

// deftail - transpiled function from  t7.c:99
func deftail() {
	var i int32
	var c int32
	var lf int32
	var lwid int32
	for i = 0; i < 44; i++ {
		if linestop[i] != 0 {
			noarch.Fprintf(tabout, []byte(".nr #%c 0-1\n\x00"), linestop[i]+int32('a')-1)
		}
	}
	noarch.Fprintf(tabout, []byte(".nr #a 0-1\n\x00"))
	noarch.Fprintf(tabout, []byte(".eo\n\x00"))
	noarch.Fprintf(tabout, []byte(".de T#\n\x00"))
	noarch.Fprintf(tabout, []byte(".nr 35 1m\n\x00"))
	noarch.Fprintf(tabout, []byte(".ds #d .d\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\(ts\\n(.z\\(ts\\(ts .ds #d nl\n\x00"))
	noarch.Fprintf(tabout, []byte(".mk ##\n\x00"))
	noarch.Fprintf(tabout, []byte(".nr ## -1v\n\x00"))
	noarch.Fprintf(tabout, []byte(".ls 1\n\x00"))
	for i = 0; i < 44; i++ {
		if linestop[i] != 0 {
			noarch.Fprintf(tabout, []byte(".if \\n(#T>=0 .nr #%c \\n(#T\n\x00"), linestop[i]+int32('a')-1)
		}
	}
	if boxflg != 0 || allflg != 0 || dboxflg != 0 {
		if fullbot[nlin-1] == 0 {
			if noarch.Not(pr1403) {
				// bottom of table line
				noarch.Fprintf(tabout, []byte(".if \\n(T. .vs \\n(.vu-\\n(.sp\n\x00"))
			}
			noarch.Fprintf(tabout, []byte(".if \\n(T. \x00"))
			drawline(nlin, 0, ncol, func() int32 {
				if dboxflg != 0 {
					return int32('=')
				}
				return int32('-')
			}(), 1, 0)
			noarch.Fprintf(tabout, []byte("\n.if \\n(T. .vs\n\x00"))
		}
	}
	{
		// T. is really an argument to a macro but because of
		//     eqn we don't dare pass it as an argument and reference by $1
		for c = 0; c < ncol; c++ {
			if (func() int32 {
				lf = left(nlin-1, c, c4goUnsafeConvert_int32(&lwid))
				return lf
			}()) >= 0 {
				noarch.Fprintf(tabout, []byte(".if \\n(#%c>=0 .sp -1\n\x00"), linestop[lf]+int32('a')-1)
				noarch.Fprintf(tabout, []byte(".if \\n(#%c>=0 \x00"), linestop[lf]+int32('a')-1)
				tohcol(c)
				drawvert(lf, nlin-1, c, lwid)
				noarch.Fprintf(tabout, []byte("\\h'|\\n(TWu'\n\x00"))
			}
		}
	}
	if boxflg != 0 || allflg != 0 || dboxflg != 0 {
		// right hand line
		noarch.Fprintf(tabout, []byte(".if \\n(#a>=0 .sp -1\n\x00"))
		noarch.Fprintf(tabout, []byte(".if \\n(#a>=0 \\h'|\\n(TWu'\x00"))
		drawvert(0, nlin-1, ncol, func() int32 {
			if dboxflg != 0 {
				return 2
			}
			return 1
		}())
		noarch.Fprintf(tabout, []byte("\n\x00"))
	}
	noarch.Fprintf(tabout, []byte(".ls\n\x00"))
	noarch.Fprintf(tabout, []byte("..\n\x00"))
	noarch.Fprintf(tabout, []byte(".ec\n\x00"))
}

// watchout - transpiled function from  t8.c:4
// t8.c: write out one line of output table
var watchout int32

// once - transpiled function from  t8.c:5
var once int32

// putline - transpiled function from  t8.c:7
func putline(i int32, nl int32) {
	// i is line number for deciding format
	// nl is line number for finding data   usually identical
	var c int32
	var s int32
	var lf int32
	var ct int32
	var form int32
	var lwid int32
	var vspf int32
	var ip int32
	var cmidx int32
	var exvspen int32
	var vforml int32
	var vct int32
	var chfont int32
	var uphalf int32
	var ss []byte
	var size []byte
	var fn []byte
	var rct []byte
	exvspen = 0
	vspf = exvspen
	watchout = vspf
	cmidx = watchout
	if i == 0 {
		once = 0
	}
	if i == 0 && (allflg != 0 || boxflg != 0 || dboxflg != 0) {
		fullwide(0, func() int32 {
			if dboxflg != 0 {
				return int32('=')
			}
			return int32('-')
		}())
	}
	if len(instead[nl]) == 0 && fullbot[nl] == 0 {
		for c = 0; c < ncol; c++ {
			ss = table[nl][c].col
			if len(ss) == 0 {
				continue
			}
			if vspen(ss) != 0 {
				for ip = nl; ip < nlin; ip = next(ip) {
					ss = table[ip][c].col
					if noarch.Not(vspen(ss)) {
						break
					}
				}
				s = int32(uint32((uintptr_(ss))))
				if s > 0 && s < 128 {
					noarch.Fprintf(tabout, []byte(".ne \\n(%c|u+\\n(.Vu\n\x00"), s)
				}
				continue
			}
			if point(ss) != 0 {
				continue
			}
			s = int32(uint32((uintptr_(ss))))
			noarch.Fprintf(tabout, []byte(".ne \\n(%c|u+\\n(.Vu\n\x00"), s)
			watchout = 1
		}
	}
	if linestop[nl] != 0 {
		noarch.Fprintf(tabout, []byte(".mk #%c\n\x00"), linestop[nl]+int32('a')-1)
	}
	lf = prev(nl)
	if instead[nl] != nil {
		noarch.Fprintf(tabout, []byte("%s\n\x00"), instead[nl])
		return
	}
	if fullbot[nl] != 0 {
		switch func() int32 {
			ct = fullbot[nl]
			return ct
		}() {
		case '=':
			fallthrough
		case '-':
			fullwide(nl, ct)
		}
		return
	}
	for c = 0; c < ncol; c++ {
		if len(instead[nl]) == 0 && fullbot[nl] == 0 {
			if vspen(table[nl][c].col) != 0 {
				vspf = 1
			}
		}
		if lf >= 0 {
			if vspen(table[lf][c].col) != 0 {
				vspf = 1
			}
		}
	}
	if vspf != 0 {
		noarch.Fprintf(tabout, []byte(".nr #^ \\n(\\*(#du\n\x00"))
		// current line position relative to bottom
		noarch.Fprintf(tabout, []byte(".nr #- \\n(#^\n\x00"))
	}
	vspf = 0
	chfont = 0
	for c = 0; c < ncol; c++ {
		ss = table[nl][c].col
		if len(ss) == 0 {
			continue
		}
		if font[c][stynum[nl]] != nil {
			chfont = 1
		}
		if point(ss) != 0 {
			continue
		}
		s = int32(uint32((uintptr_(ss))))
		lf = prev(nl)
		if lf >= 0 && vspen(table[lf][c].col) != 0 {
			noarch.Fprintf(tabout, []byte(".if (\\n(%c|+\\n(^%c-1v)>\\n(#- .nr #- +(\\n(%c|+\\n(^%c-\\n(#--1v)\n\x00"), s, 'a'+byte(c), s, 'a'+byte(c))
		} else {
			noarch.Fprintf(tabout, []byte(".if (\\n(%c|+\\n(#^-1v)>\\n(#- .nr #- +(\\n(%c|+\\n(#^-\\n(#--1v)\n\x00"), s, s)
		}
	}
	if allflg != 0 && once > 0 {
		fullwide(i, int32('-'))
	}
	once = 1
	runtabs(i, nl)
	if allh(i) != 0 && noarch.Not(pr1403) {
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.v\n\x00"), 36)
		noarch.Fprintf(tabout, []byte(".vs \\n(.vu-\\n(.sp\n\x00"))
		noarch.Fprintf(tabout, []byte(".nr 35 \\n(.vu\n\x00"))
	} else {
		noarch.Fprintf(tabout, []byte(".nr 35 1m\n\x00"))
	}
	if chfont != 0 {
		noarch.Fprintf(tabout, []byte(".nr %2d \\n(.f\n\x00"), 31)
		noarch.Fprintf(tabout, []byte(".af %2d 01\n\x00"), 31)
	}
	noarch.Fprintf(tabout, []byte("\\&\x00"))
	vct = 0
	for c = 0; c < ncol; c++ {
		uphalf = 0
		if watchout == 0 && i+1 < nlin && (func() int32 {
			lf = left(i, c, c4goUnsafeConvert_int32(&lwid))
			return lf
		}()) >= 0 {
			tohcol(c)
			drawvert(lf, i, c, lwid)
			vct += 2
		}
		if rightl != 0 && c+1 == ncol {
			continue
		}
		vforml = i
		for lf = prev(nl); lf >= 0 && vspen(table[lf][c].col) != 0; lf = prev(lf) {
			vforml = lf
		}
		form = ctype(vforml, c)
		if form != int32('s') {
			rct = reg(c, 0)
			if form == int32('a') {
				rct = reg(c, 1)
			}
			if form == int32('n') && table[nl][c].rcol != nil && lused[c] == 0 {
				rct = reg(c, 1)
			}
			noarch.Fprintf(tabout, []byte("\\h'|\\n(%2su'\x00"), rct)
		}
		ss = table[nl][c].col
		fn = font[c][stynum[vforml]]
		size = csize[c][stynum[vforml]]
		if int32(size[0]) == 0 {
			size = nil
		}
		if flags[c][stynum[nl]]&2 != 0 && pr1403 == 0 {
			uphalf = 1
		}
		switch func() int32 {
			ct = ctype(vforml, c)
			return ct
		}() {
		case 'n':
			fallthrough
		case 'a':
			if table[nl][c].rcol != nil {
				if lused[c] != 0 {
					//Zero field width
					ip = prev(nl)
					if ip >= 0 {
						if vspen(table[ip][c].col) != 0 {
							if exvspen == 0 {
								noarch.Fprintf(tabout, []byte("\\v'-(\\n(\\*(#du-\\n(^%cu\x00"), c+int32('a'))
								if cmidx != 0 {
									// code folded from here
									noarch.Fprintf(tabout, []byte("-((\\n(#-u-\\n(^%cu)/2u)\x00"), c+int32('a'))
								}
								// unfolding
								vct++
								if pr1403 != 0 {
									// must round to whole lines
									// code folded from here
									noarch.Fprintf(tabout, []byte("/1v*1v\x00"))
								}
								// unfolding
								noarch.Fprintf(tabout, []byte("'\x00"))
								exvspen = 1
							}
						}
					}
					noarch.Fprintf(tabout, []byte("%c%c\x00"), F1, F2)
					if uphalf != 0 {
						noarch.Fprintf(tabout, []byte("\\u\x00"))
					}
					puttext(ss, fn, size)
					if uphalf != 0 {
						noarch.Fprintf(tabout, []byte("\\d\x00"))
					}
					noarch.Fprintf(tabout, []byte("%c\x00"), F1)
				}
				ss = table[nl][c].rcol
				form = 1
				break
			}
			fallthrough
		case 'c':
			form = 3
		case 'r':
			form = 2
		case 'l':
			form = 1
		case '-':
			fallthrough
		case '=':
			if real_(table[nl][c].col) != 0 {
				noarch.Fprintf(noarch.Stderr, []byte("%s: line %d: Data ignored on table line %d\n\x00"), ifile, iline-1, i+1)
			}
			makeline(i, c, ct)
			continue
			fallthrough
		default:
			continue
		}
		if func() int32 {
			if (ct == int32('a') || ct == int32('n')) && table[nl][c].rcol != nil {
				return rused[c]
			}
			return used[c]
		}() != 0 {
			if ifline(ss) != 0 {
				//Zero field width
				// form: 1 left, 2 right, 3 center adjust
				makeline(i, c, ifline(ss))
				continue
			}
			if filler(ss) != 0 {
				noarch.Fprintf(tabout, []byte("\\l'|\\n(%2su\\&%s'\x00"), reg(c, 2), ss[0+2:])
				continue
			}
			ip = prev(nl)
			cmidx = noarch.BoolToInt(flags[c][stynum[nl]]&(4|8) == 0)
			if ip >= 0 {
				if vspen(table[ip][c].col) != 0 {
					if exvspen == 0 {
						noarch.Fprintf(tabout, []byte("\\v'-(\\n(\\*(#du-\\n(^%cu\x00"), c+int32('a'))
						if cmidx != 0 {
							noarch.Fprintf(tabout, []byte("-((\\n(#-u-\\n(^%cu)/2u)\x00"), c+int32('a'))
						}
						vct++
						if pr1403 != 0 {
							// round to whole lines
							noarch.Fprintf(tabout, []byte("/1v*1v\x00"))
						}
						noarch.Fprintf(tabout, []byte("'\x00"))
					}
				}
			}
			noarch.Fprintf(tabout, []byte("%c\x00"), F1)
			if form != 1 {
				noarch.Fprintf(tabout, []byte("%c\x00"), F2)
			}
			if vspen(ss) != 0 {
				vspf = 1
			} else {
				if uphalf != 0 {
					noarch.Fprintf(tabout, []byte("\\u\x00"))
				}
				puttext(ss, fn, size)
				if uphalf != 0 {
					noarch.Fprintf(tabout, []byte("\\d\x00"))
				}
			}
			if form != 2 {
				noarch.Fprintf(tabout, []byte("%c\x00"), F2)
			}
			noarch.Fprintf(tabout, []byte("%c\x00"), F1)
		}
		ip = prev(nl)
		if ip >= 0 {
			if vspen(table[ip][c].col) != 0 {
				exvspen = noarch.BoolToInt(c+1 < ncol && vspen(table[ip][c+1].col) != 0 && topat[c] == topat[c+1] && cmidx == noarch.BoolToInt(flags[c+1][stynum[nl]]&(4|8) == 0) && left(i, c+1, c4goUnsafeConvert_int32(&lwid)) < 0)
				if exvspen == 0 {
					noarch.Fprintf(tabout, []byte("\\v'(\\n(\\*(#du-\\n(^%cu\x00"), c+int32('a'))
					if cmidx != 0 {
						noarch.Fprintf(tabout, []byte("-((\\n(#-u-\\n(^%cu)/2u)\x00"), c+int32('a'))
					}
					vct++
					if pr1403 != 0 {
						// round to whole lines
						noarch.Fprintf(tabout, []byte("/1v*1v\x00"))
					}
					noarch.Fprintf(tabout, []byte("'\x00"))
				}
			} else {
				exvspen = 0
			}
		}
		if vct > 7 && c < ncol {
			// if lines need to be split for gcos here is the place for a backslash
			noarch.Fprintf(tabout, []byte("\n.sp-1\n\\&\x00"))
			vct = 0
		}
	}
	noarch.Fprintf(tabout, []byte("\n\x00"))
	if allh(i) != 0 && noarch.Not(pr1403) {
		noarch.Fprintf(tabout, []byte(".vs \\n(%du\n\x00"), 36)
	}
	if watchout != 0 {
		funnies(i, nl)
	}
	if vspf != 0 {
		for c = 0; c < ncol; c++ {
			if vspen(table[nl][c].col) != 0 && (nl == 0 || (func() int32 {
				lf = prev(nl)
				return lf
			}()) < 0 || noarch.Not(vspen(table[lf][c].col))) {
				noarch.Fprintf(tabout, []byte(".nr ^%c \\n(#^u\n\x00"), 'a'+byte(c))
				topat[c] = nl
			}
		}
	}
}

// puttext - transpiled function from  t8.c:271
func puttext(s []byte, fn []byte, size []byte) {
	if point(s) != 0 {
		putfont(fn)
		putsize(size)
		noarch.Fprintf(tabout, []byte("%s\x00"), s)
		if int32(fn[0]) > 0 {
			noarch.Fprintf(tabout, []byte("\\f(\\n(%2d\x00"), 31)
		}
		if len(size) != 0 {
			putsize([]byte("0\x00"))
		}
	}
}

// funnies - transpiled function from  t8.c:286
func funnies(stl int32, lin int32) {
	// write out funny diverted things
	var c int32
	var s int32
	var pl int32
	var lwid int32
	var dv int32
	var lf int32
	var ct int32
	var fn []byte
	var ss []byte
	// rmember current vertical position
	noarch.Fprintf(tabout, []byte(".mk ##\n\x00"))
	// bottom position
	noarch.Fprintf(tabout, []byte(".nr %d \\n(##\n\x00"), 31)
	for c = 0; c < ncol; c++ {
		ss = table[lin][c].col
		if point(ss) != 0 {
			continue
		}
		if len(ss) == 0 {
			continue
		}
		s = int32(uint32((uintptr_(ss))))
		noarch.Fprintf(tabout, []byte(".sp |\\n(##u-1v\n\x00"))
		noarch.Fprintf(tabout, []byte(".nr %d \x00"), 37)
		ct = 0
		for pl = stl; pl >= 0 && noarch.Not(int32(((__ctype_b_loc())[0])[(func() int32 {
			ct = ctype(pl, c)
			return ct
		}())])&int32(uint16(noarch.ISalpha))); pl = prev(pl) {
		}
		switch ct {
		case 'n':
			fallthrough
		case 'c':
			noarch.Fprintf(tabout, []byte("(\\n(%2su+\\n(%2su-\\n(%c-u)/2u\n\x00"), reg(c, 0), reg(c-1+ctspan(lin, c), 2), s)
		case 'l':
			noarch.Fprintf(tabout, []byte("\\n(%2su\n\x00"), reg(c, 0))
		case 'a':
			noarch.Fprintf(tabout, []byte("\\n(%2su\n\x00"), reg(c, 1))
		case 'r':
			noarch.Fprintf(tabout, []byte("\\n(%2su-\\n(%c-u\n\x00"), reg(c, 2), s)
			break
		}
		noarch.Fprintf(tabout, []byte(".in +\\n(%du\n\x00"), 37)
		fn = font[c][stynum[stl]]
		if fn[0] != 0 {
			noarch.Fprintf(tabout, []byte(".ft %s\n\x00"), fn)
		}
		pl = prev(stl)
		if stl > 0 && pl >= 0 && vspen(table[pl][c].col) != 0 {
			noarch.Fprintf(tabout, []byte(".sp |\\n(^%cu\n\x00"), 'a'+byte(c))
			if flags[c][stynum[stl]]&(4|8) == 0 {
				noarch.Fprintf(tabout, []byte(".nr %d \\n(#-u-\\n(^%c-\\n(%c|+1v\n\x00"), 38, 'a'+byte(c), s)
				noarch.Fprintf(tabout, []byte(".if \\n(%d>0 .sp \\n(%du/2u\x00"), 38, 38)
				if pr1403 != 0 {
					// round
					noarch.Fprintf(tabout, []byte("/1v*1v\x00"))
				}
				noarch.Fprintf(tabout, []byte("\n\x00"))
			}
		}
		noarch.Fprintf(tabout, []byte(".%c+\n\x00"), s)
		noarch.Fprintf(tabout, []byte(".in -\\n(%du\n\x00"), 37)
		if fn[0] != 0 {
			noarch.Fprintf(tabout, []byte(".ft P\n\x00"))
		}
		noarch.Fprintf(tabout, []byte(".mk %d\n\x00"), 32)
		noarch.Fprintf(tabout, []byte(".if \\n(%d>\\n(%d .nr %d \\n(%d\n\x00"), 32, 31, 31, 32)
	}
	noarch.Fprintf(tabout, []byte(".sp |\\n(%du\n\x00"), 31)
	{
		dv = 0
		for c = dv; c < ncol; c++ {
			if stl+1 < nlin && (func() int32 {
				lf = left(stl, c, c4goUnsafeConvert_int32(&lwid))
				return lf
			}()) >= 0 {
				if func() int32 {
					defer func() {
						dv++
					}()
					return dv
				}() == 0 {
					noarch.Fprintf(tabout, []byte(".sp -1\n\x00"))
				}
				tohcol(c)
				dv++
				drawvert(lf, stl, c, lwid)
			}
		}
	}
	if dv != 0 {
		noarch.Fprintf(tabout, []byte("\n\x00"))
	}
}

// putfont - transpiled function from  t8.c:362
func putfont(fn []byte) {
	if fn != nil && int32(fn[0]) != 0 {
		noarch.Fprintf(tabout, func() []byte {
			if int32(fn[1]) != 0 {
				return []byte("\\f(%.2s\x00")
			}
			return []byte("\\f%.2s\x00")
		}(), fn)
	}
}

// putsize - transpiled function from  t8.c:370
func putsize(s []byte) {
	if s != nil && int32(s[0]) != 0 {
		noarch.Fprintf(tabout, []byte("\\s%s\x00"), s)
	}
}

// useln - transpiled function from  t9.c:3
// t9.c: write lines for tables over 200 lines
var useln int32

// yetmore - transpiled function from  t9.c:5
func yetmore() {
	for useln = 0; useln < 250 && len(table[useln]) == 0; useln++ {
	}
	if useln >= 250 {
		error_([]byte("Wierd.  No data in table.\x00"))
	}
	table[0] = table[useln]
	for useln = nlin - 1; useln >= 0 && (fullbot[useln] != 0 || instead[useln] != nil); useln-- {
	}
	if useln < 0 {
		error_([]byte("Wierd.  No real lines in table.\x00"))
	}
	domore(leftover)
	for gets1(func() []byte {
		cstore = cspace
		tempVar3 := &cstore
		return *tempVar3
	}(), 2000) != nil && domore(cstore) != 0 {
	}
	last = cstore
}

// domore - transpiled function from  t9.c:24
func domore(dataln []byte) int32 {
	var icol int32
	var ch int32
	if prefix([]byte(".TE\x00"), dataln) != 0 {
		return 0
	}
	if int32(dataln[0]) == int32('.') && noarch.Not(int32(((__ctype_b_loc())[0])[int32(dataln[1])])&int32(uint16(noarch.ISdigit))) {
		noarch.Fprintf(tabout, []byte("%s\n\x00"), dataln)
		return 1
	}
	fullbot[0] = 0
	instead[0] = nil
	if int32(dataln[1]) == 0 {
		switch int32(dataln[0]) {
		case '_':
			fullbot[0] = int32('-')
			putline(useln, 0)
			return 1
		case '=':
			fullbot[0] = int32('=')
			putline(useln, 0)
			return 1
		}
	}
	for icol = 0; icol < ncol; icol++ {
		table[0][icol].col = dataln
		table[0][icol].rcol = nil
		for ; (func() int32 {
			ch = int32(dataln[0])
			return ch
		}()) != int32('\x00') && ch != tab; func() []byte {
			tempVarUnary := dataln
			defer func() {
				dataln = dataln[0+1:]
			}()
			return tempVarUnary
		}() {
		}
		(func() []byte {
			defer func() {
				dataln = dataln[0+1:]
			}()
			return dataln
		}())[0] = '\x00'
		switch ctype(useln, icol) {
		case 'n':
			table[0][icol].rcol = maknew(table[0][icol].col)
		case 'a':
			table[0][icol].rcol = table[0][icol].col
			table[0][icol].col = []byte("\x00")
			break
		}
		for ctype(useln, icol+1) == int32('s') {
			// spanning
			table[0][func() int32 {
				icol++
				return icol
			}()].col = []byte("\x00")
		}
		if ch == int32('\x00') {
			break
		}
	}
	for func() int32 {
		icol++
		return icol
	}() < ncol {
		table[0][icol].col = []byte("\x00")
	}
	putline(useln, 0)
	// reuse space for numerical items
	exstore = exspace
	return 1
}

// checkuse - transpiled function from  tb.c:5
func checkuse() {
	// tb.c: check which entries exist, also storage allocation
	var i int32
	var c int32
	var k int32
	for c = 0; c < ncol; c++ {
		rused[c] = 0
		lused[c] = rused[c]
		used[c] = lused[c]
		for i = 0; i < nlin; i++ {
			if instead[i] != nil || fullbot[i] != 0 {
				continue
			}
			k = ctype(i, c)
			if k == int32('-') || k == int32('=') {
				continue
			}
			if k == int32('n') || k == int32('a') {
				rused[c] |= real_(table[i][c].rcol)
				if noarch.Not(real_(table[i][c].rcol)) {
					used[c] |= real_(table[i][c].col)
				}
				if table[i][c].rcol != nil {
					lused[c] |= real_(table[i][c].col)
				}
			} else {
				used[c] |= real_(table[i][c].col)
			}
		}
	}
}

// real_ - transpiled function from  tb.c:31
func real_(s []byte) int32 {
	if len(s) == 0 {
		return 0
	}
	if noarch.Not(point(s)) {
		return 1
	}
	if int32(s[0]) == 0 {
		return 0
	}
	return 1
}

// spcount - transpiled function from  tb.c:44
var spcount int32

// spvecs - transpiled function from  tb.c:46
var spvecs [][]byte = make([][]byte, 20)

// chspace - transpiled function from  tb.c:48
func chspace() []byte {
	var pp []byte
	if spvecs[spcount] != nil {
		return spvecs[func() int32 {
			defer func() {
				spcount++
			}()
			return spcount
		}()]
	}
	if spcount >= 20 {
		error_([]byte("Too many characters in table\x00"))
	}
	pp = make([]byte, uint32(2000+300)*1)
	spvecs[func() int32 {
		defer func() {
			spcount++
		}()
		return spcount
	}()] = pp
	if (int64(uintptr(unsafe.Pointer(&pp[0])))/int64(1)-func() int64 {
		c4go_temp_name := []byte(-1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}()) == 0 || len(pp) == 0 {
		error_([]byte("no space for characters\x00"))
	}
	return pp
}

// thisvec - transpiled function from  tb.c:65
var thisvec []byte

// tpcount - transpiled function from  tb.c:66
var tpcount int32 = -1

// tpvecs - transpiled function from  tb.c:67
var tpvecs [][]byte = make([][]byte, 50)

// alocv - transpiled function from  tb.c:69
func alocv(n int32) []int32 {
	var tp []int32
	var q []int32
	if tpcount < 0 || (int64(uintptr(unsafe.Pointer(&thisvec[0+n])))/int64(1)-int64(uintptr(unsafe.Pointer(&(tpvecs[tpcount])[0+2000])))/int64(1)) > 0 {
		tpcount++
		if len(tpvecs[tpcount]) == 0 {
			tpvecs[tpcount] = make([]byte, 2000*1)
		}
		thisvec = tpvecs[tpcount]
		if len(thisvec) == 0 {
			error_([]byte("no space for vectors\x00"))
		}
	}
	tp = (*[10000]int32)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&thisvec[0]))) / int64(1))))[:]
	thisvec = thisvec[0+n:]
	for q = tp; (int64(uintptr(unsafe.Pointer(&q[0])))/int64(4) - int64(uintptr(unsafe.Pointer(&(*[10000]int32)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&thisvec[0]))) / int64(1))))[0])))/int64(4)) < 0; func() []int32 {
		tempVarUnary := q
		defer func() {
			q = q[0+1:]
		}()
		return tempVarUnary
	}() {
		q[0] = 0
	}
	return tp
}

// release - transpiled function from  tb.c:91
func release() {
	// give back unwanted space in some vectors
	// this should call free; it does not because
	//    alloc() is so buggy
	spcount = 0
	tpcount = -1
	exstore = nil
}

// choochar - transpiled function from  tc.c:7
func choochar() {
	// tc.c: find character not in table to delimit fields
	// choose funny characters to delimit fields
	var had []int32 = make([]int32, 128)
	var ilin int32
	var icol int32
	var k int32
	var s []byte
	for icol = 0; icol < 128; icol++ {
		had[icol] = 0
	}
	F2 = 0
	F1 = F2
	for ilin = 0; ilin < nlin; ilin++ {
		if instead[ilin] != nil || fullbot[ilin] != 0 {
			continue
		}
		for icol = 0; icol < ncol; icol++ {
			k = ctype(ilin, icol)
			if k == 0 || k == int32('-') || k == int32('=') {
				continue
			}
			s = table[ilin][icol].col
			if point(s) != 0 {
				for ; s[0] != 0; func() []byte {
					tempVarUnary := s
					defer func() {
						s = s[0+1:]
					}()
					return tempVarUnary
				}() {
					if int32(uint8(s[0])) < 128 {
						had[uint8(s[0])] = 1
					}
				}
			}
			s = table[ilin][icol].rcol
			if point(s) != 0 {
				for ; s[0] != 0; func() []byte {
					tempVarUnary := s
					defer func() {
						s = s[0+1:]
					}()
					return tempVarUnary
				}() {
					if int32(uint8(s[0])) < 128 {
						had[uint8(s[0])] = 1
					}
				}
			}
		}
	}
	{
		// choose first funny character
		for s = []byte("\x02\x03\x05\x06\a!%&#/?,:;<=>@`^~_{}+-*ABCDEFGHIJKMNOPQRSTUVWXZabcdefgjkoqrstwxyzY\x00"); s[0] != 0; func() []byte {
			tempVarUnary := s
			defer func() {
				s = s[0+1:]
			}()
			return tempVarUnary
		}() {
			if had[s[0]] == 0 {
				F1 = int32(s[0])
				had[F1] = 1
				break
			}
		}
	}
	{
		// choose second funny character
		for s = []byte("\x02\x03\x05\x06\a!%&#/?,:;<=>@`^~_{}+-*ABCDEFGHIJKMNOPQRSTUVWXZabcdefgjkoqrstwxyzu\x00"); s[0] != 0; func() []byte {
			tempVarUnary := s
			defer func() {
				s = s[0+1:]
			}()
			return tempVarUnary
		}() {
			if had[s[0]] == 0 {
				F2 = int32(s[0])
				break
			}
		}
	}
	if F1 == 0 || F2 == 0 {
		error_([]byte("couldn't find characters to use for delimiters\x00"))
	}
}

// point - transpiled function from  tc.c:55
func point(ss []byte) int32 {
	var s vlong = vlong((int64(uint32((uintptr_(ss))))))
	return noarch.BoolToInt(s >= vlong((128)) || s < vlong((0)))
}

// error_ - transpiled function from  te.c:6
func error_(s []byte) {
	// te.c: error message control, input line count
	noarch.Fprintf(noarch.Stderr, []byte("\n%s:%d: %s\n\x00"), ifile, iline, s)
	noarch.Fprintf(noarch.Stderr, []byte("tbl quits\n\x00"))
	noarch.Exit(1)
}

// gets1 - transpiled function from  te.c:13
func gets1(s []byte, size int32) []byte {
	var ns []byte
	var nbl int32
	iline++
	ns = s
	if noarch.Fgets(s, size, tabin) == nil {
		if swapin() == 0 {
			return nil
		}
	}
	nbl = noarch.Strlen(s)
	// remove the newline
	s[nbl-1] = '\x00'
	s = c4goPointerArithByteSlice(s, int(nbl-2))
	for nbl = 0; (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&ns[0])))/int64(1)) > 0 && int32(s[0]) == int32('\\'); func() []byte {
		tempVarUnary := s
		defer func() {
			s = c4goPointerArithByteSlice(s, int(-1))
		}()
		return tempVarUnary
	}() {
		nbl++
	}
	if linstart != 0 && nbl%2 != 0 {
		// fold escaped nl if in table
		gets1(s[0+1:], size-int32((int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&ns[0])))/int64(1))))
	}
	return s
}

// backup - transpiled function from  te.c:37
var backup []byte = make([]byte, 500)

// backp - transpiled function from  te.c:38
var backp []byte = backup

// un1getc - transpiled function from  te.c:40
func un1getc(c int32) {
	if c == int32('\n') {
		iline--
	}
	(func() []byte {
		defer func() {
			backp = backp[0+1:]
		}()
		return backp
	}())[0] = byte(c)
	if (int64(uintptr(unsafe.Pointer(&backp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&backup[0+500])))/int64(1)) >= 0 {
		error_([]byte("too much backup\x00"))
	}
}

// get1char - transpiled function from  te.c:50
func get1char() int32 {
	var c int32
	if (int64(uintptr(unsafe.Pointer(&backp[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&backup[0])))/int64(1)) > 0 {
		c = int32((func() []byte {
			backp = c4goPointerArithByteSlice(backp, int(-1))
			return backp
		}())[0])
	} else {
		c = noarch.Fgetc(tabin)
	}
	if c == -1 {
		if swapin() == 0 {
			error_([]byte("unexpected EOF\x00"))
		}
		c = noarch.Fgetc(tabin)
	}
	if c == int32('\n') {
		iline++
	}
	return c
}

// savefill - transpiled function from  tf.c:4
func savefill() {
	// tf.c: save and restore fill mode around table
	// remembers various things: fill mode, vs, ps in mac 35 (SF)
	noarch.Fprintf(tabout, []byte(".de %d\n\x00"), 35)
	noarch.Fprintf(tabout, []byte(".ps \\n(.s\n\x00"))
	noarch.Fprintf(tabout, []byte(".vs \\n(.vu\n\x00"))
	noarch.Fprintf(tabout, []byte(".in \\n(.iu\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\n(.u .fi\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\n(.j .ad\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\n(.j=0 .na\n\x00"))
	noarch.Fprintf(tabout, []byte("..\n\x00"))
	noarch.Fprintf(tabout, []byte(".nf\n\x00"))
	// set obx offset if useful
	noarch.Fprintf(tabout, []byte(".nr #~ 0\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\n(.T .if n .nr #~ 0.6n\n\x00"))
}

// rstofill - transpiled function from  tf.c:23
func rstofill() {
	noarch.Fprintf(tabout, []byte(".%d\n\x00"), 35)
}

// endoff - transpiled function from  tf.c:30
func endoff() {
	var i int32
	for i = 0; i < 44; i++ {
		if linestop[i] != 0 {
			noarch.Fprintf(tabout, []byte(".nr #%c 0\n\x00"), linestop[i]+int32('a')-1)
		}
	}
	for i = 0; i < texct; i++ {
		noarch.Fprintf(tabout, []byte(".rm %c+\n\x00"), int32(texstr[i]))
	}
	noarch.Fprintf(tabout, []byte("%s\n\x00"), last)
}

// ifdivert - transpiled function from  tf.c:44
func ifdivert() {
	noarch.Fprintf(tabout, []byte(".ds #d .d\n\x00"))
	noarch.Fprintf(tabout, []byte(".if \\(ts\\n(.z\\(ts\\(ts .ds #d nl\n\x00"))
}

// saveline - transpiled function from  tf.c:52
func saveline() {
	noarch.Fprintf(tabout, []byte(".if \\n+(b.=1 .nr d. \\n(.c-\\n(c.-1\n\x00"))
	linstart = iline
}

// restline - transpiled function from  tf.c:60
func restline() {
	noarch.Fprintf(tabout, []byte(".if \\n-(b.=0 .nr c. \\n(.c-\\n(d.-%d\n\x00"), iline-linstart)
	linstart = 0
}

// cleanfc - transpiled function from  tf.c:68
func cleanfc() {
	noarch.Fprintf(tabout, []byte(".fc\n\x00"))
}

// gettext_tbl - transpiled function from  tg.c:4
func gettext_tbl(sp []byte, ilin int32, icol int32, fn []byte, sz []byte) int32 {
	// tg.c: process included text blocks
	// get a section of text
	var line []byte = make([]byte, 4096)
	var oname int32
	var startline int32
	var vs []byte
	startline = iline
	if texname == 0 {
		error_([]byte("Too many text block diversions\x00"))
	}
	if textflg == 0 {
		// remember old line length
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.lu\n\x00"), 34)
		textflg = 1
	}
	noarch.Fprintf(tabout, []byte(".eo\n\x00"))
	noarch.Fprintf(tabout, []byte(".am %s\n\x00"), reg(icol, 2))
	noarch.Fprintf(tabout, []byte(".br\n\x00"))
	noarch.Fprintf(tabout, []byte(".di %c+\n\x00"), texname)
	rstofill()
	if fn != nil && int32(fn[0]) != 0 {
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.f\n.ft %s\n\x00"), 31, fn)
	}
	// protect font
	noarch.Fprintf(tabout, []byte(".ft \\n(.f\n\x00"))
	vs = vsize[icol][stynum[ilin]]
	if sz != nil && int32(sz[0]) != 0 || vs != nil && int32(vs[0]) != 0 {
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.v\n\x00"), 39)
		if len(vs) == 0 || int32(vs[0]) == 0 {
			vs = []byte("\\n(.s+2\x00")
		}
		if sz != nil && int32(sz[0]) != 0 {
			noarch.Fprintf(tabout, []byte(".ps %s\n\x00"), sz)
		}
		noarch.Fprintf(tabout, []byte(".vs %s\n\x00"), vs)
		noarch.Fprintf(tabout, []byte(".if \\n(%du>\\n(.vu .sp \\n(%du-\\n(.vu\n\x00"), 39, 39)
	}
	if cll[icol][0] != 0 {
		noarch.Fprintf(tabout, []byte(".ll %sn\n\x00"), cll[icol])
	} else {
		noarch.Fprintf(tabout, []byte(".ll \\n(%du*%du/%du\n\x00"), 34, ctspan(ilin, icol), ncol+1)
	}
	noarch.Fprintf(tabout, []byte(".if \\n(.l<\\n(%2s .ll \\n(%2su\n\x00"), reg(icol, 2), reg(icol, 2))
	if ctype(ilin, icol) == int32('a') {
		noarch.Fprintf(tabout, []byte(".ll -2n\n\x00"))
	}
	noarch.Fprintf(tabout, []byte(".in 0\n\x00"))
	for {
		if len(gets1(line, int32(4096))) == 0 {
			iline = startline
			error_([]byte("missing closing T}\x00"))
		}
		if int32(line[0]) == int32('T') && int32(line[1]) == int32('}') && int32(line[2]) == tab {
			break
		}
		if match([]byte("T}\x00"), line) != 0 {
			break
		}
		noarch.Fprintf(tabout, []byte("%s\n\x00"), line)
	}
	if fn != nil && int32(fn[0]) != 0 {
		noarch.Fprintf(tabout, []byte(".ft \\n(%d\n\x00"), 31)
	}
	if sz != nil && int32(sz[0]) != 0 {
		noarch.Fprintf(tabout, []byte(".br\n.ps\n.vs\n\x00"))
	}
	noarch.Fprintf(tabout, []byte(".br\n\x00"))
	noarch.Fprintf(tabout, []byte(".di\n\x00"))
	noarch.Fprintf(tabout, []byte(".nr %c| \\n(dn\n\x00"), texname)
	noarch.Fprintf(tabout, []byte(".nr %c- \\n(dl\n\x00"), texname)
	noarch.Fprintf(tabout, []byte("..\n\x00"))
	noarch.Fprintf(tabout, []byte(".ec \\\n\x00"))
	if line[2] != 0 {
		// copy remainder of line
		tcopy(sp, line[0+3:])
	} else {
		sp[0] = byte(0)
	}
	oname = texname
	texname = int32(texstr[func() int32 {
		texct++
		return texct
	}()])
	return oname
}

// untext - transpiled function from  tg.c:78
func untext() {
	rstofill()
	noarch.Fprintf(tabout, []byte(".nf\n\x00"))
	noarch.Fprintf(tabout, []byte(".ll \\n(%du\n\x00"), 34)
}

// interv - transpiled function from  ti.c:5
func interv(i int32, c int32) int32 {
	// ti.c: classify line intersections
	// determine local environment for intersections
	var ku int32
	var kl int32
	if c >= ncol || c == 0 {
		if dboxflg != 0 {
			if i == 0 {
				return 2
			}
			if i >= nlin {
				return 1
			}
			return 3
		}
		if c >= ncol {
			return 0
		}
	}
	if i > 0 {
		ku = lefdata(i-1, c)
	} else {
		ku = 0
	}
	if i+1 >= nlin && allh(i) != 0 {
		kl = 0
	} else {
		kl = lefdata(func() int32 {
			if allh(i) != 0 {
				return i + 1
			}
			return i
		}(), c)
	}
	if ku == 2 && kl == 2 {
		return 3
	}
	if ku == 2 {
		return 1
	}
	if kl == 2 {
		return 2
	}
	return 0
}

// interh - transpiled function from  ti.c:36
func interh(i int32, c int32) int32 {
	var kl int32
	var kr int32
	if fullbot[i] == int32('=') || dboxflg != 0 && (i == 0 || i >= nlin-1) {
		if c == ncol {
			return 1
		}
		if c == 0 {
			return 2
		}
		return 3
	}
	if i >= nlin {
		return 0
	}
	if c > 0 {
		kl = thish(i, c-1)
	} else {
		kl = 0
	}
	if kl <= 1 && i > 0 && allh(up1(i)) != 0 {
		if c > 0 {
			kl = thish(up1(i), c-1)
		} else {
			kl = 0
		}
	}
	kr = thish(i, c)
	if kr <= 1 && i > 0 && allh(up1(i)) != 0 {
		if c > 0 {
			kr = thish(up1(i), c)
		} else {
			kr = 0
		}
	}
	if kl == int32('=') && kr == int32('=') {
		return 3
	}
	if kl == int32('=') {
		return 1
	}
	if kr == int32('=') {
		return 2
	}
	return 0
}

// up1 - transpiled function from  ti.c:66
func up1(i int32) int32 {
	i--
	for instead[i] != nil && i > 0 {
		i--
	}
	return i
}

// maknew - transpiled function from  tm.c:4
func maknew(str []byte) []byte {
	// tm.c: split numerical fields
	// make two numerical fields
	var c int32
	var p []byte
	var q []byte
	var ba []byte
	var dpoint []byte
	p = str
	for ba = nil; (func() int32 {
		c = int32(str[0])
		return c
	}()) != 0; func() []byte {
		tempVarUnary := str
		defer func() {
			str = str[0+1:]
		}()
		return tempVarUnary
	}() {
		if c == int32('\\') && int32(str[0+1]) == int32('&') {
			ba = str
		}
	}
	str = p
	if len(ba) == 0 {
		for dpoint = nil; str[0] != 0; func() []byte {
			tempVarUnary := str
			defer func() {
				str = str[0+1:]
			}()
			return tempVarUnary
		}() {
			if int32(str[0]) == int32('.') && noarch.Not(ineqn(str, p)) && ((int64(uintptr(unsafe.Pointer(&str[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&p[0])))/int64(1)) > 0 && digit(int32(str[0-1])) != 0 || digit(int32(str[0+1])) != 0) {
				dpoint = str
			}
		}
		if len(dpoint) == 0 {
			for ; (int64(uintptr(unsafe.Pointer(&str[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&p[0])))/int64(1)) > 0; func() []byte {
				tempVarUnary := str
				defer func() {
					str = c4goPointerArithByteSlice(str, int(-1))
				}()
				return tempVarUnary
			}() {
				if digit(int32(str[0-1])) != 0 && noarch.Not(ineqn(str, p)) {
					break
				}
			}
		}
		if dpoint == nil && (int64(uintptr(unsafe.Pointer(&p[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&str[0])))/int64(1)) == 0 {
			// not numerical, don't split
			return []byte(0)
		}
		if dpoint != nil {
			str = dpoint
		}
	} else {
		str = ba
	}
	p = str
	if len(exstore) == 0 || (int64(uintptr(unsafe.Pointer(&exstore[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&exlim[0])))/int64(1)) > 0 {
		exspace = chspace()
		exstore = exspace
		exlim = exstore[0+2000:]
	}
	q = exstore
	for (func() byte {
		(func() []byte {
			defer func() {
				exstore = exstore[0+1:]
			}()
			return exstore
		}())[0] = (func() []byte {
			defer func() {
				str = str[0+1:]
			}()
			return str
		}())[0]
		return (func() []byte {
			defer func() {
				exstore = exstore[0+1:]
			}()
			return exstore
		}())[0]
	}()) != 0 {
	}
	p[0] = byte(0)
	return q
}

// ineqn - transpiled function from  tm.c:47
func ineqn(s []byte, p []byte) int32 {
	// true if s is in a eqn within p
	var ineq int32
	var c int32
	for (func() int32 {
		c = int32(p[0])
		return c
	}()) != 0 {
		if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&p[0])))/int64(1)) == 0 {
			return ineq
		}
		p = p[0+1:]
		if ineq == 0 && c == delim1 {
			ineq = 1
		} else if ineq == 1 && c == delim2 {
			ineq = 0
		}
	}
	return 0
}

// nregs - transpiled function from  tr.c:3
// tr.c: number register allocation
var nregs [][]byte = [][]byte{[]byte("40\x00"), []byte("41\x00"), []byte("42\x00"), []byte("43\x00"), []byte("44\x00"), []byte("45\x00"), []byte("46\x00"), []byte("47\x00"), []byte("48\x00"), []byte("49\x00"), []byte("50\x00"), []byte("51\x00"), []byte("52\x00"), []byte("53\x00"), []byte("54\x00"), []byte("55\x00"), []byte("56\x00"), []byte("57\x00"), []byte("58\x00"), []byte("59\x00"), []byte("60\x00"), []byte("61\x00"), []byte("62\x00"), []byte("63\x00"), []byte("64\x00"), []byte("65\x00"), []byte("66\x00"), []byte("67\x00"), []byte("68\x00"), []byte("69\x00"), []byte("70\x00"), []byte("71\x00"), []byte("72\x00"), []byte("73\x00"), []byte("74\x00"), []byte("75\x00"), []byte("76\x00"), []byte("77\x00"), []byte("78\x00"), []byte("79\x00"), []byte("80\x00"), []byte("81\x00"), []byte("82\x00"), []byte("83\x00"), []byte("84\x00"), []byte("85\x00"), []byte("86\x00"), []byte("87\x00"), []byte("88\x00"), []byte("89\x00"), []byte("90\x00"), []byte("91\x00"), []byte("92\x00"), []byte("93\x00"), []byte("94\x00"), []byte("95\x00"), []byte("96\x00"), []byte("97\x00"), []byte("4q\x00"), []byte("4r\x00"), []byte("4s\x00"), []byte("4t\x00"), []byte("4u\x00"), []byte("4v\x00"), []byte("4w\x00"), []byte("4x\x00"), []byte("4y\x00"), []byte("4z\x00"), []byte("4;\x00"), []byte("4.\x00"), []byte("4a\x00"), []byte("4b\x00"), []byte("4c\x00"), []byte("4d\x00"), []byte("4e\x00"), []byte("4f\x00"), []byte("4g\x00"), []byte("4h\x00"), []byte("4i\x00"), []byte("4j\x00"), []byte("4k\x00"), []byte("4l\x00"), []byte("4m\x00"), []byte("4n\x00"), []byte("4o\x00"), []byte("4p\x00"), []byte("5a\x00"), []byte("5b\x00"), []byte("5c\x00"), []byte("5d\x00"), []byte("5e\x00"), []byte("5f\x00"), []byte("5g\x00"), []byte("5h\x00"), []byte("5i\x00"), []byte("5j\x00"), []byte("5k\x00"), []byte("5l\x00"), []byte("5m\x00"), []byte("5n\x00"), []byte("5o\x00"), []byte("5p\x00"), []byte("5q\x00"), []byte("5r\x00"), []byte("5s\x00"), []byte("5t\x00"), []byte("5u\x00"), []byte("5v\x00"), []byte("5w\x00"), []byte("5x\x00"), nil}

// reg - transpiled function from  tr.c:20
func reg(col int32, place int32) []byte {
	if 888 < uint32(2*3*qcol) {
		// this array must have at least 3*qcol entries
		//    or illegal register names will result
		error_([]byte("Too many columns for registers\x00"))
	}
	return nregs[qcol*place+col]
}

// match - transpiled function from  ts.c:4
func match(s1 []byte, s2 []byte) int32 {
	for int32(s1[0]) == int32(s2[0]) {
		if int32((func() []byte {
			defer func() {
				s1 = s1[0+1:]
			}()
			return s1
		}())[0]) == int32('\x00') {
			// ts.c: minor string processing subroutines
			return 1
		} else {
			s2 = s2[0+1:]
		}
	}
	return 0
}

// prefix - transpiled function from  ts.c:16
func prefix(small []byte, big []byte) int32 {
	var c int32
	for (func() int32 {
		c = int32((func() []byte {
			defer func() {
				small = small[0+1:]
			}()
			return small
		}())[0])
		return c
	}()) == int32((func() []byte {
		defer func() {
			big = big[0+1:]
		}()
		return big
	}())[0]) {
		if c == 0 {
			return 1
		}
	}
	return noarch.BoolToInt(c == 0)
}

// letter - transpiled function from  ts.c:28
func letter(ch int32) int32 {
	if ch >= int32('a') && ch <= int32('z') {
		return 1
	}
	if ch >= int32('A') && ch <= int32('Z') {
		return 1
	}
	return 0
}

// numb - transpiled function from  ts.c:39
func numb(str []byte) int32 {
	// convert to integer
	var k int32
	for k = 0; int32(str[0]) >= int32('0') && int32(str[0]) <= int32('9'); func() []byte {
		tempVarUnary := str
		defer func() {
			str = str[0+1:]
		}()
		return tempVarUnary
	}() {
		k = k*10 + int32(str[0]) - int32('0')
	}
	return k
}

// digit - transpiled function from  ts.c:50
func digit(x int32) int32 {
	return noarch.BoolToInt(x >= int32('0') && x <= int32('9'))
}

// max - transpiled function from  ts.c:57
func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// tcopy - transpiled function from  ts.c:64
func tcopy(s []byte, t []byte) {
	for (func() byte {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = (func() []byte {
			defer func() {
				t = t[0+1:]
			}()
			return t
		}())[0]
		return (func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0]
	}()) != 0 {
	}
}

// ctype - transpiled function from  tt.c:4
func ctype(il int32, ic int32) int32 {
	if instead[il] != nil {
		// tt.c: subroutines for drawing horizontal lines
		return 0
	}
	if fullbot[il] != 0 {
		return 0
	}
	il = stynum[il]
	return style[ic][il]
}

// min - transpiled function from  tt.c:16
func min(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// fspan - transpiled function from  tt.c:23
func fspan(i int32, c int32) int32 {
	c++
	return noarch.BoolToInt(c < ncol && ctype(i, c) == int32('s'))
}

// lspan - transpiled function from  tt.c:31
func lspan(i int32, c int32) int32 {
	var k int32
	if ctype(i, c) != int32('s') {
		return 0
	}
	c++
	if c < ncol && ctype(i, c) == int32('s') {
		return 0
	}
	for k = 0; ctype(i, func() int32 {
		c--
		return c
	}()) == int32('s'); k++ {
	}
	return k
}

// ctspan - transpiled function from  tt.c:47
func ctspan(i int32, c int32) int32 {
	var k int32
	c++
	for k = 1; c < ncol && ctype(i, c) == int32('s'); k++ {
		c++
	}
	return k
}

// tohcol - transpiled function from  tt.c:58
func tohcol(ic int32) {
	if ic == 0 {
		noarch.Fprintf(tabout, []byte("\\h'|0'\x00"))
	} else {
		noarch.Fprintf(tabout, []byte("\\h'(|\\n(%2su+|\\n(%2su)/2u'\x00"), reg(ic, 0), reg(ic-1, 2))
	}
}

// allh - transpiled function from  tt.c:69
func allh(i int32) int32 {
	// return true if every element in line i is horizontal
	// also at least one must be horizontl
	var c int32
	var one int32
	var k int32
	if fullbot[i] != 0 {
		return 1
	}
	if i >= nlin {
		return noarch.BoolToInt(dboxflg != 0 || boxflg != 0)
	}
	{
		c = 0
		for one = c; c < ncol; c++ {
			k = thish(i, c)
			if k == 0 {
				return 0
			}
			if k == 1 {
				continue
			}
			one = 1
		}
	}
	return one
}

// thish - transpiled function from  tt.c:92
func thish(i int32, c int32) int32 {
	var t int32
	var s []byte
	var pc []colstr
	if c < 0 {
		return 0
	}
	if i < 0 {
		return 0
	}
	t = ctype(i, c)
	if t == int32('_') || t == int32('-') {
		return int32('-')
	}
	if t == int32('=') {
		return int32('=')
	}
	if t == int32('^') {
		return 1
	}
	if fullbot[i] != 0 {
		return fullbot[i]
	}
	if t == int32('s') {
		return thish(i, c-1)
	}
	if t == 0 {
		return 1
	}
	pc = table[i][c:]
	if t == int32('a') {
		s = pc[0].rcol
	} else {
		s = pc[0].col
	}
	if len(s) == 0 || point(s) != 0 && int32(s[0]) == 0 {
		return 1
	}
	if vspen(s) != 0 {
		return 1
	}
	if (func() int32 {
		t = barent(s)
		return t
	}()) != 0 {
		return t
	}
	return 0
}

// makeline - transpiled function from  tu.c:5
func makeline(i int32, c int32, lintype int32) {
	// tu.c: draws horizontal lines
	var cr int32
	var type_ int32
	var shortl int32
	type_ = thish(i, c)
	if type_ == 0 {
		return
	}
	shortl = noarch.BoolToInt(int32(table[i][c].col[0]) == int32('\\'))
	if c > 0 && noarch.Not(shortl) && thish(i, c-1) == type_ {
		return
	}
	if shortl == 0 {
		for cr = c; cr < ncol && (ctype(i, cr) == int32('s') || type_ == thish(i, cr)); cr++ {
		}
	} else {
		for cr = c + 1; cr < ncol && ctype(i, cr) == int32('s'); cr++ {
		}
	}
	drawline(i, c, cr-1, lintype, 0, shortl)
}

// fullwide - transpiled function from  tu.c:26
func fullwide(i int32, lintype int32) {
	var cr int32
	var cl int32
	if noarch.Not(pr1403) {
		noarch.Fprintf(tabout, []byte(".nr %d \\n(.v\n.vs \\n(.vu-\\n(.sp\n\x00"), 36)
	}
	cr = 0
	for cr < ncol {
		cl = cr
		for i > 0 && vspand(prev(i), cl, 1) != 0 {
			cl++
		}
		for cr = cl; cr < ncol; cr++ {
			if i > 0 && vspand(prev(i), cr, 1) != 0 {
				break
			}
		}
		if cl < ncol {
			drawline(i, cl, func() int32 {
				if cr < ncol {
					return cr - 1
				}
				return cr
			}(), lintype, 1, 0)
		}
	}
	noarch.Fprintf(tabout, []byte("\n\x00"))
	if noarch.Not(pr1403) {
		noarch.Fprintf(tabout, []byte(".vs \\n(%du\n\x00"), 36)
	}
}

// drawline - transpiled function from  tu.c:50
func drawline(i int32, cl int32, cr int32, lintype int32, noheight int32, shortl int32) {
	var exhr []byte
	var exhl []byte
	var lnch []byte
	var lcount int32
	var ln int32
	var linpos int32
	var oldpos int32
	var nodata int32
	lcount = 0
	exhl = []byte("\x00")
	exhr = exhl
	switch lintype {
	case '-':
		lcount = 1
	case '=':
		lcount = func() int32 {
			if pr1403 != 0 {
				return 1
			}
			return 2
		}()
	case 4:
		lcount = 1
		break
	}
	if lcount <= 0 {
		return
	}
	nodata = noarch.BoolToInt(cr-cl >= ncol || noheight != 0 || allh(i) != 0)
	if noarch.Not(nodata) {
		noarch.Fprintf(tabout, []byte("\\v'-.5m'\x00"))
	}
	{
		oldpos = 0
		for ln = oldpos; ln < lcount; ln++ {
			linpos = 2*ln - lcount + 1
			if linpos != oldpos {
				noarch.Fprintf(tabout, []byte("\\v'%dp'\x00"), linpos-oldpos)
			}
			oldpos = linpos
			if shortl == 0 {
				tohcol(cl)
				if lcount > 1 {
					switch interv(i, cl) {
					case 1:
						exhl = func() []byte {
							if ln == 0 {
								return []byte("1p\x00")
							}
							return []byte("-1p\x00")
						}()
					case 2:
						exhl = func() []byte {
							if ln == 1 {
								return []byte("1p\x00")
							}
							return []byte("-1p\x00")
						}()
					case 3:
						exhl = []byte("1p\x00")
						break
					}
					if exhl[0] != 0 {
						noarch.Fprintf(tabout, []byte("\\h'%s'\x00"), exhl)
					}
				} else if lcount == 1 {
					switch interv(i, cl) {
					case 1:
						fallthrough
					case 2:
						exhl = []byte("-1p\x00")
					case 3:
						exhl = []byte("1p\x00")
						break
					}
					if exhl[0] != 0 {
						noarch.Fprintf(tabout, []byte("\\h'%s'\x00"), exhl)
					}
				}
				if lcount > 1 {
					switch interv(i, cr+1) {
					case 1:
						exhr = func() []byte {
							if ln == 0 {
								return []byte("-1p\x00")
							}
							return []byte("+1p\x00")
						}()
					case 2:
						exhr = func() []byte {
							if ln == 1 {
								return []byte("-1p\x00")
							}
							return []byte("+1p\x00")
						}()
					case 3:
						exhr = []byte("-1p\x00")
						break
					}
				} else if lcount == 1 {
					switch interv(i, cr+1) {
					case 1:
						fallthrough
					case 2:
						exhr = []byte("+1p\x00")
					case 3:
						exhr = []byte("-1p\x00")
						break
					}
				}
			} else {
				noarch.Fprintf(tabout, []byte("\\h'|\\n(%2su'\x00"), reg(cl, 0))
			}
			noarch.Fprintf(tabout, []byte("\\s\\n(%d\x00"), 33)
			if linsize != 0 {
				noarch.Fprintf(tabout, []byte("\\v'-\\n(%dp/6u'\x00"), 33)
			}
			if shortl != 0 {
				noarch.Fprintf(tabout, []byte("\\l'|\\n(%2su'\x00"), reg(cr, 2))
			} else {
				lnch = []byte("\\(ul\x00")
				if pr1403 != 0 {
					if lintype == 2 {
						lnch = []byte("=\x00")
					} else {
						lnch = []byte("\\(ru\x00")
					}
				}
				if cr+1 >= ncol {
					noarch.Fprintf(tabout, []byte("\\l'|\\n(TWu%s%s'\x00"), exhr, lnch)
				} else {
					noarch.Fprintf(tabout, []byte("\\l'(|\\n(%2su+|\\n(%2su)/2u%s%s'\x00"), reg(cr, 2), reg(cr+1, 0), exhr, lnch)
				}
			}
			if linsize != 0 {
				noarch.Fprintf(tabout, []byte("\\v'\\n(%dp/6u'\x00"), 33)
			}
			noarch.Fprintf(tabout, []byte("\\s0\x00"))
		}
	}
	if oldpos != 0 {
		noarch.Fprintf(tabout, []byte("\\v'%dp'\x00"), -oldpos)
	}
	if noarch.Not(nodata) {
		noarch.Fprintf(tabout, []byte("\\v'+.5m'\x00"))
	}
}

// getstop - transpiled function from  tu.c:160
func getstop() {
	var i int32
	var c int32
	var k int32
	var junk int32
	var stopp int32
	stopp = 1
	for i = 0; i < 250; i++ {
		linestop[i] = 0
	}
	for i = 0; i < nlin; i++ {
		for c = 0; c < ncol; c++ {
			k = left(i, c, c4goUnsafeConvert_int32(&junk))
			if k >= 0 && linestop[k] == 0 {
				linestop[k] = func() int32 {
					stopp++
					return stopp
				}()
			}
		}
	}
	if boxflg != 0 || allflg != 0 || dboxflg != 0 {
		linestop[0] = 1
	}
}

// left - transpiled function from  tu.c:179
func left(i int32, c int32, lwidp []int32) int32 {
	var kind int32
	var li int32
	var lj int32
	// returns -1 if no line to left
	// returns number of line where it starts
	// stores into lwid the kind of line
	lwidp[0] = 0
	if i < 0 {
		return -1
	}
	kind = lefdata(i, c)
	if kind == 0 {
		return -1
	}
	if i+1 < nlin {
		if lefdata(next(i), c) == kind {
			return -1
		}
	}
	li = i
	for i >= 0 && lefdata(i, c) == kind {
		i = prev(func() int32 {
			li = i
			tempVar3 := &li
			return *tempVar3
		}())
	}
	if prev(li) == -1 {
		li = 0
	}
	lwidp[0] = kind
	for lj = i + 1; lj < li; lj++ {
		if instead[lj] != nil && noarch.Strcmp(instead[lj], []byte(".TH\x00")) == 0 {
			return li
		}
	}
	for i = i + 1; i < li; i++ {
		if fullbot[i] != 0 {
			li = i
		}
	}
	return li
}

// lefdata - transpiled function from  tu.c:211
func lefdata(i int32, c int32) int32 {
	var ck int32
	if i >= nlin {
		i = nlin - 1
	}
	if ctype(i, c) == int32('s') {
		for ck = c; ctype(i, ck) == int32('s'); ck-- {
		}
		if thish(i, ck) == 0 {
			return 0
		}
	}
	i = stynum[i]
	i = lefline[c][i]
	if i > 0 {
		return i
	}
	if dboxflg != 0 && c == 0 {
		return 2
	}
	if allflg != 0 {
		return 1
	}
	if boxflg != 0 && c == 0 {
		return 1
	}
	return 0
}

// next - transpiled function from  tu.c:238
func next(i int32) int32 {
	for i+1 < nlin {
		i++
		if noarch.Not(fullbot[i]) && instead[i] == nil {
			break
		}
	}
	return i
}

// prev - transpiled function from  tu.c:250
func prev(i int32) int32 {
	for func() int32 {
		i--
		return i
	}() >= 0 && (fullbot[i] != 0 || instead[i] != nil) {
	}
	return i
}

// drawvert - transpiled function from  tv.c:4
func drawvert(start int32, end int32, c int32, lwid int32) {
	// tv.c: draw vertical lines
	var exb []byte
	var ext []byte
	var tp int32
	var sl int32
	var ln int32
	var pos int32
	var epb int32
	var ept int32
	var vm int32
	end++
	vm = int32('v')
	for instead[end] != nil {
		// note: nr 35 has value of 1m outside of linesize
		end++
	}
	for ln = 0; ln < lwid; ln++ {
		ept = 0
		epb = ept
		pos = -ln - lwid + 1
		if pos != tp {
			noarch.Fprintf(tabout, []byte("\\h'%dp'\x00"), pos-tp)
		}
		tp = pos
		if end < nlin {
			if fullbot[end] != 0 || instead[end] == nil && allh(end) != 0 {
				epb = 2
			} else {
				switch midbar(end, c) {
				case '-':
					exb = []byte("1v-.5m\x00")
				case '=':
					exb = []byte("1v-.5m\x00")
					epb = 1
					break
				}
			}
		}
		if lwid > 1 {
			switch interh(end, c) {
			case 3:
				epb--
			case 2:
				epb += func() int32 {
					if ln == 0 {
						return 1
					}
					return -1
				}()
			case 1:
				epb += func() int32 {
					if ln == 1 {
						return 1
					}
					return -1
				}()
				break
			}
		}
		if lwid == 1 {
			switch interh(end, c) {
			case 3:
				epb--
			case 2:
				fallthrough
			case 1:
				epb++
				break
			}
		}
		if start > 0 {
			sl = start - 1
			for sl >= 0 && instead[sl] != nil {
				sl--
			}
			if sl >= 0 && (fullbot[sl] != 0 || allh(sl) != 0) {
				ept = 0
			} else if sl >= 0 {
				switch midbar(sl, c) {
				case '-':
					ext = []byte(".5m\x00")
				case '=':
					ext = []byte(".5m\x00")
					ept = -1
				default:
					vm = int32('m')
					break
				}
			} else {
				ept = -4
			}
		} else if start == 0 && allh(0) != 0 {
			ept = 0
			vm = int32('m')
		}
		if lwid > 1 {
			switch interh(start, c) {
			case 3:
				ept++
			case 1:
				ept += func() int32 {
					if ln == 0 {
						return 1
					}
					return -1
				}()
			case 2:
				ept += func() int32 {
					if ln == 1 {
						return 1
					}
					return -1
				}()
				break
			}
		} else if lwid == 1 {
			switch interh(start, c) {
			case 3:
				ept++
			case 1:
				fallthrough
			case 2:
				ept--
				break
			}
		}
		if exb != nil {
			noarch.Fprintf(tabout, []byte("\\v'%s'\x00"), exb)
		}
		if epb != 0 {
			noarch.Fprintf(tabout, []byte("\\v'%dp'\x00"), epb)
		}
		noarch.Fprintf(tabout, []byte("\\s\\n(%d\x00"), 33)
		if linsize != 0 {
			noarch.Fprintf(tabout, []byte("\\v'-\\n(%dp/6u'\x00"), 33)
		}
		// adjustment for T450 nroff boxes
		noarch.Fprintf(tabout, []byte("\\h'-\\n(#~u'\x00"))
		// adjustment to make vertical and horizontal lines meet properly
		noarch.Fprintf(tabout, []byte("\\h'-\\n(%dp*7u/100u'\x00"), 33)
		noarch.Fprintf(tabout, []byte("\\v'\\n(%dp*2u/100u'\x00"), 33)
		noarch.Fprintf(tabout, []byte("\\L'|\\n(#%cu-%s\x00"), linestop[start]+int32('a')-1, func() []byte {
			if vm == int32('v') {
				return []byte("1v\x00")
			}
			return []byte("\\n(35u\x00")
		}())
		if ext != nil {
			noarch.Fprintf(tabout, []byte("-(%s)\x00"), ext)
		}
		if exb != nil {
			noarch.Fprintf(tabout, []byte("-(%s)\x00"), exb)
		}
		pos = ept - epb
		if pos != 0 {
			noarch.Fprintf(tabout, []byte("%s%dp\x00"), func() []byte {
				if pos >= 0 {
					return []byte("+\x00")
				}
				return []byte("\x00")
			}(), pos)
		}
		// the string #d is either "nl" or ".d" depending
		//    on diversions; on GCOS not the same
		noarch.Fprintf(tabout, []byte("'\\s0\\v'\\n(\\*(#du-\\n(#%cu+%s\x00"), linestop[start]+int32('a')-1, func() []byte {
			if vm == int32('v') {
				return []byte("1v\x00")
			}
			return []byte("\\n(35u\x00")
		}())
		if ext != nil {
			noarch.Fprintf(tabout, []byte("+%s\x00"), ext)
		}
		if ept != 0 {
			noarch.Fprintf(tabout, []byte("%s%dp\x00"), func() []byte {
				if -ept > 0 {
					return []byte("+\x00")
				}
				return []byte("\x00")
			}(), -ept)
		}
		noarch.Fprintf(tabout, []byte("'\x00"))
		if linsize != 0 {
			noarch.Fprintf(tabout, []byte("\\v'\\n(%dp/6u'\x00"), 33)
		}
	}
}

// midbar - transpiled function from  tv.c:140
func midbar(i int32, c int32) int32 {
	var k int32
	k = midbcol(i, c)
	if k == 0 && c > 0 {
		k = midbcol(i, c-1)
	}
	return k
}

// midbcol - transpiled function from  tv.c:152
func midbcol(i int32, c int32) int32 {
	var ct int32
	for (func() int32 {
		ct = ctype(i, c)
		return ct
	}()) == int32('s') {
		c--
	}
	if ct == int32('-') || ct == int32('=') {
		return ct
	}
	if (func() int32 {
		ct = barent(table[i][c].col)
		return ct
	}()) != 0 {
		return ct
	}
	return 0
}

// barent - transpiled function from  tv.c:167
func barent(s []byte) int32 {
	if len(s) == 0 {
		return 1
	}
	if noarch.Not(point(s)) {
		return 0
	}
	if int32(s[0]) == int32('\\') {
		s = s[0+1:]
	}
	if int32(s[1]) != 0 {
		return 0
	}
	switch int32(s[0]) {
	case '_':
		return int32('-')
	case '=':
		return int32('=')
	}
	return 0
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[10000]int32)(unsafe.Pointer(c4go_name))[:]
}

// this refers to the relative position of lines
//t1.c
//t2.c
//t3.c
//t4.c
//t5.c
//t6.c
//t7.c
//t8.c
//t9.c
//tb.c
//tc.c
//te.c
//tf.c
//tg.c
//ti.c
//tm.c
//tr.c
//ts.c
//tt.c
//tu.c
//tv.c

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
