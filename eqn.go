//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

// #include </usr/include/string.h>
import "C"

import "unicode"
import "reflect"
import "runtime"
import "os"
import "fmt"
import "github.com/Konstantin8105/c4go/noarch"
import "unsafe"

// sbuf - transpiled function from  eqn.h:102
//
// * NEATEQN MAIN HEADER
// *
// * In Neateqn equations are recursively decomposed into boxes.  eqn.c
// * reads the input and makes eqn boxes by calling appropriate functions
// * from box.c.
//
// predefined array sizes
// registers used by neateqn
// helpers
// token and atom types
// spaces in hundredths of em
// small helper functions
// reading the source
// tokenizer
// default definitions and operators
// variable length string buffer
type sbuf struct {
	s  []byte
	sz int32
	n  int32
}

// box - transpiled function from  eqn.h:141
// allocated buffer
// buffer size
// length of the string stored in s
// tex styles
// equations
type box struct {
	raw    sbuf
	szreg  int32
	szown  int32
	reg    int32
	atoms  int32
	tbeg   int32
	tcur   int32
	style  int32
	tomark []byte
}

// box_alloc - transpiled function from  box.c:8
func box_alloc(szreg int32, pre int32, style int32) []box {
	// equation boxes
	var box_c4go_postfix []box = (*[10000]box)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	noarch.Memset((*[10000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&box_c4go_postfix[0]))) / int64(1))))[:], byte(0), 64)
	sbuf_init((*[10000]sbuf)(unsafe.Pointer(&box_c4go_postfix[0].raw))[:])
	box_c4go_postfix[0].szreg = szreg
	box_c4go_postfix[0].atoms = 0
	box_c4go_postfix[0].style = style
	if pre != 0 {
		box_c4go_postfix[0].tcur = pre
	}
	return box_c4go_postfix
}

// box_free - transpiled function from  box.c:21
func box_free(box_c4go_postfix []box) {
	if box_c4go_postfix[0].reg != 0 {
		sregrm(box_c4go_postfix[0].reg)
	}
	if box_c4go_postfix[0].szown != 0 {
		nregrm(box_c4go_postfix[0].szreg)
	}
	sbuf_done((*[10000]sbuf)(unsafe.Pointer(&box_c4go_postfix[0].raw))[:])
	_ = box_c4go_postfix
}

// box_put - transpiled function from  box.c:31
func box_put(box_c4go_postfix []box, s []byte) {
	sbuf_append((*[10000]sbuf)(unsafe.Pointer(&box_c4go_postfix[0].raw))[:], s)
	if box_c4go_postfix[0].reg != 0 {
		noarch.Printf([]byte(".as %s \"%s\n\x00"), sregname(box_c4go_postfix[0].reg), s)
	}
}

// box_putf - transpiled function from  box.c:38
func box_putf(box_c4go_postfix []box, s []byte, c4goArgs ...interface{}) {
	var buf []byte = make([]byte, 1000)
	var ap *va_list
	va_start(ap, s)
	noarch.Vsnprintf(buf, int32(1000), s, ap)
	va_end(ap)
	box_put(box_c4go_postfix, buf)
}

// box_buf - transpiled function from  box.c:48
func box_buf(box_c4go_postfix []box) []byte {
	return sbuf_buf((*[10000]sbuf)(unsafe.Pointer(&box_c4go_postfix[0].raw))[:])
}

// box_size - transpiled function from  box.c:54
func box_size(box_c4go_postfix []box, val []byte) int32 {
	// change box's point size; return the number register storing it
	var szreg int32 = box_c4go_postfix[0].szreg
	if val == nil || noarch.Not(val[0]) {
		return szreg
	}
	if noarch.Not(box_c4go_postfix[0].szown) {
		box_c4go_postfix[0].szown = 1
		box_c4go_postfix[0].szreg = nregmk()
	}
	if int32(val[0]) == int32('-') || int32(val[0]) == int32('+') {
		noarch.Printf([]byte(".nr %s %s%s\n\x00"), nregname(box_c4go_postfix[0].szreg), nreg(szreg), val)
	} else {
		noarch.Printf([]byte(".nr %s %s\n\x00"), nregname(box_c4go_postfix[0].szreg), val)
	}
	return box_c4go_postfix[0].szreg
}

// box_move - transpiled function from  box.c:70
func box_move(box_c4go_postfix []box, dy int32, dx int32) {
	if dy != 0 {
		box_putf(box_c4go_postfix, []byte("\\v'%du*%sp/100u'\x00"), dy, nreg(box_c4go_postfix[0].szreg))
	}
	if dx != 0 {
		box_putf(box_c4go_postfix, []byte("\\h'%du*%sp/100u'\x00"), dx, nreg(box_c4go_postfix[0].szreg))
	}
}

// spacing - transpiled function from  box.c:79
// T_ORD, T_BIGOP, T_BINOP, T_RELOP, T_LEFT, T_RIGHT, T_PUNC, T_INNER
var spacing [][]int32 = [][]int32{{0, 1, 2, 3, 0, 0, 0, 1}, {1, 1, 0, 3, 0, 0, 0, 1}, {2, 2, 0, 0, 2, 0, 0, 2}, {3, 3, 0, 0, 3, 0, 0, 3}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 2, 3, 0, 0, 0, 1}, {1, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 2, 3, 1, 0, 1, 1}}

// eqn_gaps - transpiled function from  box.c:91
func eqn_gaps(box_c4go_postfix []box, cur int32) int32 {
	// return the amount of automatic spacing before adding the given token
	var s int32
	// previous atom
	var a1 int32 = box_c4go_postfix[0].tcur & 240
	// current atom
	var a2 int32 = cur & 240
	if noarch.Not(box_c4go_postfix[0].style>>uint64(4)) && a1 != 0 && a2 != 0 {
		s = spacing[a1>>uint64(4)&15-1][a2>>uint64(4)&15-1]
	}
	if s == 3 {
		return e_thickspace
	}
	if s == 2 {
		return e_mediumspace
	}
	if s != 0 {
		return e_thinspace
	}
	return 0
}

// box_italiccorrection - transpiled function from  box.c:106
func box_italiccorrection(box_c4go_postfix []box) {
	if box_c4go_postfix[0].atoms != 0 && box_c4go_postfix[0].tcur&256 != 0 {
		// call just before inserting a non-italic character
		box_put(box_c4go_postfix, []byte("\\/\x00"))
	}
	box_c4go_postfix[0].tcur &= ^256
}

// box_fixatom - transpiled function from  box.c:113
func box_fixatom(cur int32, pre int32) int32 {
	if cur == 48 && (noarch.Not(pre) || pre == 64 || pre == 32 || pre == 80 || pre == 112) {
		return 16
	}
	if cur == 64 && (noarch.Not(pre) || pre == 80) {
		return 16
	}
	return cur
}

// box_beforeput - transpiled function from  box.c:124
func box_beforeput(box_c4go_postfix []box, type_ int32, breakable int32) {
	// call before inserting a token with box_put() and box_putf()
	// automatically inserted space before this token
	var autogaps int32
	if box_c4go_postfix[0].atoms != 0 {
		autogaps = eqn_gaps(box_c4go_postfix, type_&240)
		if noarch.Not(type_ & 256) {
			box_italiccorrection(box_c4go_postfix)
		}
		if autogaps != 0 && type_ != 20 && box_c4go_postfix[0].tcur != 20 {
			box_italiccorrection(box_c4go_postfix)
			if breakable != 0 {
				// enlarge a space to match autogaps
				box_putf(box_c4go_postfix, []byte("\\s[\\En(.s*%du*%sp/100u/\\w' 'u]\\j'%d' \\s0\x00"), autogaps, nreg(box_c4go_postfix[0].szreg), def_brcost(box_c4go_postfix[0].tcur&240))
			} else {
				box_putf(box_c4go_postfix, []byte("\\h'%du*%sp/100u'\x00"), autogaps, nreg(box_c4go_postfix[0].szreg))
			}
		}
	}
	if box_c4go_postfix[0].tomark != nil {
		noarch.Printf([]byte(".nr %s 0\\w'%s'\n\x00"), box_c4go_postfix[0].tomark, box_toreg(box_c4go_postfix))
		box_c4go_postfix[0].tomark = nil
	}
}

// box_afterput - transpiled function from  box.c:150
func box_afterput(box_c4go_postfix []box, type_ int32) {
	// call after inserting a token with box_put() and box_putf()
	box_c4go_postfix[0].atoms++
	box_c4go_postfix[0].tcur = type_&256 | box_fixatom(type_&255, box_c4go_postfix[0].tcur&255)
	if noarch.Not(box_c4go_postfix[0].tbeg) {
		box_c4go_postfix[0].tbeg = box_c4go_postfix[0].tcur
	}
}

// box_puttext - transpiled function from  box.c:159
func box_puttext(box_c4go_postfix []box, type_ int32, s []byte, c4goArgs ...interface{}) {
	// insert s with the given type
	var buf []byte = make([]byte, 1000)
	var ap *va_list
	va_start(ap, s)
	noarch.Vsnprintf(buf, int32(1000), s, ap)
	va_end(ap)
	box_beforeput(box_c4go_postfix, type_, 0)
	if noarch.Not(box_c4go_postfix[0].tcur&256) && type_&256 != 0 {
		box_put(box_c4go_postfix, []byte("\\,\x00"))
	}
	box_put(box_c4go_postfix, buf)
	box_afterput(box_c4go_postfix, type_)
}

// box_merge - transpiled function from  box.c:174
func box_merge(box_c4go_postfix []box, sub []box, breakable int32) {
	if box_empty(sub) != 0 {
		// append sub to box
		return
	}
	box_beforeput(box_c4go_postfix, sub[0].tbeg, breakable)
	box_toreg(box_c4go_postfix)
	box_put(box_c4go_postfix, box_toreg(sub))
	if noarch.Not(box_c4go_postfix[0].tbeg) {
		box_c4go_postfix[0].tbeg = sub[0].tbeg
	}
	if sub[0].atoms == 1 {
		// fix atom type only if merging a single atom
		box_afterput(box_c4go_postfix, sub[0].tcur)
	} else {
		box_c4go_postfix[0].tcur = sub[0].tcur
		box_c4go_postfix[0].atoms += sub[0].atoms
	}
}

// roff_max - transpiled function from  box.c:193
func roff_max(dst int32, a int32, b int32) {
	// put the maximum of number registers a and b into register dst
	noarch.Printf([]byte(".ie %s>=%s .nr %s 0+%s\n\x00"), nreg(a), nreg(b), nregname(dst), nreg(a))
	noarch.Printf([]byte(".el .nr %s 0+%s\n\x00"), nregname(dst), nreg(b))
}

// tok_dim - transpiled function from  box.c:201
func tok_dim(s []byte, wd int32, ht int32, dp int32) {
	// return the width, height and depth of a string
	noarch.Printf([]byte(".nr %s 0\\w'%s'\n\x00"), nregname(wd), s)
	if ht != 0 {
		noarch.Printf([]byte(".nr %s 0-\\n[bbury]\n\x00"), nregname(ht))
	}
	if dp != 0 {
		noarch.Printf([]byte(".nr %s 0\\n[bblly]\n\x00"), nregname(dp))
	}
}

// box_suprise - transpiled function from  box.c:210
func box_suprise(box_c4go_postfix []box) int32 {
	if box_c4go_postfix[0].style&1 != 0 {
		return e_sup3
	}
	if box_c4go_postfix[0].style == 0 {
		return e_sup1
	}
	return e_sup2
}

// box_sub - transpiled function from  box.c:217
func box_sub(box_c4go_postfix []box, sub []box, sup []box) {
	var box_wd int32 = nregmk()
	var box_wdnoic int32 = nregmk()
	var box_dp int32 = nregmk()
	var box_ht int32 = nregmk()
	var sub_wd int32 = nregmk()
	var sup_wd int32 = nregmk()
	var all_wd int32 = nregmk()
	var sup_dp int32 = nregmk()
	var sub_ht int32 = nregmk()
	var sup_rise int32 = nregmk()
	var sub_fall int32 = nregmk()
	var tmp_18e int32 = nregmk()
	var sub_cor int32 = nregmk()
	if sub != nil {
		box_italiccorrection(sub)
	}
	if sup != nil {
		box_italiccorrection(sup)
	}
	if sub != nil {
		tok_dim(box_toreg(box_c4go_postfix), box_wdnoic, 0, 0)
	}
	box_italiccorrection(box_c4go_postfix)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	tok_dim(box_toreg(box_c4go_postfix), box_wd, box_ht, box_dp)
	box_putf(box_c4go_postfix, []byte("\\h'5m/100u'\x00"))
	if sup != nil {
		tok_dim(box_toreg(sup), sup_wd, 0, sup_dp)
		// 18a
		noarch.Printf([]byte(".nr %s 0%su-(%dm/100u)\n\x00"), nregname(sup_rise), nreg(box_ht), e_supdrop)
		// 18c
		noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s (%dm/100u)\n\x00"), nreg(sup_rise), box_suprise(box_c4go_postfix), nregname(sup_rise), box_suprise(box_c4go_postfix))
		noarch.Printf([]byte(".if %s<(%s+(%dm/100u/4)) .nr %s 0%s+(%dm/100u/4)\n\x00"), nreg(sup_rise), nreg(sup_dp), e_xheight, nregname(sup_rise), nreg(sup_dp), e_xheight)
	}
	if sub != nil {
		tok_dim(box_toreg(sub), sub_wd, sub_ht, 0)
		// 18a
		noarch.Printf([]byte(".nr %s 0%su+(%dm/100u)\n\x00"), nregname(sub_fall), nreg(box_dp), e_subdrop)
	}
	if sub != nil && sup == nil {
		// 18b
		noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s (%dm/100u)\n\x00"), nreg(sub_fall), e_sub1, nregname(sub_fall), e_sub1)
		noarch.Printf([]byte(".if %s<(%s-(%dm/100u*4/5)) .nr %s 0%s-(%dm/100u*4/5)\n\x00"), nreg(sub_fall), nreg(sub_ht), e_xheight, nregname(sub_fall), nreg(sub_ht), e_xheight)
	}
	if len(sub) == 0 && len(sup) == 0 {
		// 18d
		noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s (%dm/100u)\n\x00"), nreg(sub_fall), e_sub2, nregname(sub_fall), e_sub2)
		// 18e
		noarch.Printf([]byte(".if (%s-%s)-(%s-%s)<(%dm/100u*4) \\{\\\n\x00"), nreg(sup_rise), nreg(sup_dp), nreg(sub_ht), nreg(sub_fall), e_rulethickness)
		noarch.Printf([]byte(".nr %s (%dm/100u*4)+%s-(%s-%s)\n\x00"), nregname(sub_fall), e_rulethickness, nreg(sub_ht), nreg(sup_rise), nreg(sup_dp))
		noarch.Printf([]byte(".nr %s (%dm/100u*4/5)-(%s-%s)\n\x00"), nregname(tmp_18e), e_xheight, nreg(sup_rise), nreg(sup_dp))
		noarch.Printf([]byte(".if %s>0 .nr %s +%s\n\x00"), nreg(tmp_18e), nregname(sup_rise), nreg(tmp_18e))
		noarch.Printf([]byte(".if %s>0 .nr %s -%s \\}\n\x00"), nreg(tmp_18e), nregname(sub_fall), nreg(tmp_18e))
	}
	if sup != nil {
		// writing the superscript
		box_putf(box_c4go_postfix, []byte("\\v'-%su'%s\\v'%su'\x00"), nreg(sup_rise), box_toreg(sup), nreg(sup_rise))
		if sub != nil {
			box_putf(box_c4go_postfix, []byte("\\h'-%su'\x00"), nreg(sup_wd))
		}
	}
	if sub != nil {
		// writing the subscript
		// subscript correction
		noarch.Printf([]byte(".nr %s (%s-%s)\n\x00"), nregname(sub_cor), nreg(box_wd), nreg(box_wdnoic))
		noarch.Printf([]byte(".if %s>0 .nr %s (%s+%s)*(%s-%s)/%s\n\x00"), nreg(box_ht), nregname(sub_cor), nreg(box_ht), nreg(sub_fall), nreg(box_wd), nreg(box_wdnoic), nreg(box_ht))
		noarch.Printf([]byte(".nr %s -%s\n\x00"), nregname(sub_wd), nreg(sub_cor))
		box_putf(box_c4go_postfix, []byte("\\h'-%su'\x00"), nreg(sub_cor))
		box_putf(box_c4go_postfix, []byte("\\v'%su'%s\\v'-%su'\x00"), nreg(sub_fall), box_toreg(sub), nreg(sub_fall))
		if sup != nil {
			box_putf(box_c4go_postfix, []byte("\\h'-%su'\x00"), nreg(sub_wd))
			roff_max(all_wd, sub_wd, sup_wd)
			box_putf(box_c4go_postfix, []byte("\\h'+%su'\x00"), nreg(all_wd))
		}
	}
	box_putf(box_c4go_postfix, []byte("\\h'%dm/100u'\x00"), e_scriptspace)
	nregrm(box_wd)
	nregrm(box_wdnoic)
	nregrm(box_dp)
	nregrm(box_ht)
	nregrm(sub_wd)
	nregrm(sup_wd)
	nregrm(all_wd)
	nregrm(sup_dp)
	nregrm(sub_ht)
	nregrm(sup_rise)
	nregrm(sub_fall)
	nregrm(tmp_18e)
	nregrm(sub_cor)
}

// box_from - transpiled function from  box.c:332
func box_from(box_c4go_postfix []box, lim []box, llim []box, ulim []box) {
	// box's width
	var lim_wd int32 = nregmk()
	// box's height
	var lim_ht int32 = nregmk()
	// box's depth
	var lim_dp int32 = nregmk()
	// llim's width
	var llim_wd int32 = nregmk()
	// ulim's width
	var ulim_wd int32 = nregmk()
	// ulim's depth
	var ulim_dp int32 = nregmk()
	// llim's height
	var llim_ht int32 = nregmk()
	// the position of ulim
	var ulim_rise int32 = nregmk()
	// the position of llim
	var llim_fall int32 = nregmk()
	// the width of all
	var all_wd int32 = nregmk()
	box_italiccorrection(lim)
	box_beforeput(box_c4go_postfix, 32, 0)
	tok_dim(box_toreg(lim), lim_wd, lim_ht, lim_dp)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	if ulim != nil {
		tok_dim(box_toreg(ulim), ulim_wd, 0, ulim_dp)
	}
	if llim != nil {
		tok_dim(box_toreg(llim), llim_wd, llim_ht, 0)
	}
	if len(ulim) == 0 && len(llim) == 0 {
		roff_max(all_wd, llim_wd, ulim_wd)
	} else {
		noarch.Printf([]byte(".nr %s %s\n\x00"), nregname(all_wd), func() []byte {
			if ulim != nil {
				return nreg(ulim_wd)
			}
			return nreg(llim_wd)
		}())
	}
	noarch.Printf([]byte(".if %s>%s .nr %s 0%s\n\x00"), nreg(lim_wd), nreg(all_wd), nregname(all_wd), nreg(lim_wd))
	box_putf(box_c4go_postfix, []byte("\\h'%su-%su/2u'\x00"), nreg(all_wd), nreg(lim_wd))
	box_merge(box_c4go_postfix, lim, 0)
	box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\x00"), nreg(lim_wd))
	if ulim != nil {
		// 13a
		noarch.Printf([]byte(".nr %s (%dm/100u)-%s\n\x00"), nregname(ulim_rise), e_bigopspacing3, nreg(ulim_dp))
		noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s (%dm/100u)\n\x00"), nreg(ulim_rise), e_bigopspacing1, nregname(ulim_rise), e_bigopspacing1)
		noarch.Printf([]byte(".nr %s +%s+%s\n\x00"), nregname(ulim_rise), nreg(lim_ht), nreg(ulim_dp))
		box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\\v'-%su'%s\\v'%su'\\h'-%su/2u'\x00"), nreg(ulim_wd), nreg(ulim_rise), box_toreg(ulim), nreg(ulim_rise), nreg(ulim_wd))
	}
	if llim != nil {
		// 13a
		noarch.Printf([]byte(".nr %s (%dm/100u)-%s\n\x00"), nregname(llim_fall), e_bigopspacing4, nreg(llim_ht))
		noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s (%dm/100u)\n\x00"), nreg(llim_fall), e_bigopspacing2, nregname(llim_fall), e_bigopspacing2)
		noarch.Printf([]byte(".nr %s +%s+%s\n\x00"), nregname(llim_fall), nreg(lim_dp), nreg(llim_ht))
		box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\\v'%su'%s\\v'-%su'\\h'-%su/2u'\x00"), nreg(llim_wd), nreg(llim_fall), box_toreg(llim), nreg(llim_fall), nreg(llim_wd))
	}
	box_putf(box_c4go_postfix, []byte("\\h'%su/2u'\x00"), nreg(all_wd))
	box_afterput(box_c4go_postfix, 32)
	nregrm(lim_wd)
	nregrm(lim_ht)
	nregrm(lim_dp)
	nregrm(llim_wd)
	nregrm(ulim_wd)
	nregrm(ulim_dp)
	nregrm(llim_ht)
	nregrm(ulim_rise)
	nregrm(llim_fall)
	nregrm(all_wd)
}

// tok_len - transpiled function from  box.c:404
func tok_len(s []byte, wd int32, len_ int32, ht int32, dp int32) {
	// return the width of s; len is the height plus depth
	noarch.Printf([]byte(".nr %s 0\\w'%s'\n\x00"), nregname(wd), s)
	if len_ != 0 {
		noarch.Printf([]byte(".nr %s 0\\n[bblly]-\\n[bbury]-2\n\x00"), nregname(len_))
	}
	if dp != 0 {
		noarch.Printf([]byte(".nr %s 0\\n[bblly]-1\n\x00"), nregname(dp))
	}
	if ht != 0 {
		noarch.Printf([]byte(".nr %s 0-\\n[bbury]-1\n\x00"), nregname(ht))
	}
}

// blen_mk - transpiled function from  box.c:416
func blen_mk(s []byte, len_ []int32) {
	// len[0]: width, len[1]: vertical length, len[2]: height, len[3]: depth
	var i int32
	for i = 0; i < 4; i++ {
		len_[i] = nregmk()
	}
	tok_len(s, len_[0], len_[1], len_[2], len_[3])
}

// blen_rm - transpiled function from  box.c:425
func blen_rm(len_ []int32) {
	// free the registers allocated with blen_mk()
	var i int32
	for i = 0; i < 4; i++ {
		nregrm(len_[i])
	}
}

// box_over - transpiled function from  box.c:433
func box_over(box_c4go_postfix []box, num []box, den []box) {
	// build a fraction; the correct font should be set up beforehand
	var num_wd int32 = nregmk()
	var num_dp int32 = nregmk()
	var den_wd int32 = nregmk()
	var den_ht int32 = nregmk()
	var all_wd int32 = nregmk()
	var num_rise int32 = nregmk()
	var den_fall int32 = nregmk()
	var bar_wd int32 = nregmk()
	var bar_dp int32 = nregmk()
	var bar_ht int32 = nregmk()
	var bar_fall int32 = nregmk()
	var tmp_15d int32 = nregmk()
	var bargap int32 = func() int32 {
		if box_c4go_postfix[0].style == 0 || box_c4go_postfix[0].style == 1 {
			return 7
		}
		return 3
	}() * e_rulethickness / 2
	box_beforeput(box_c4go_postfix, 128, 0)
	box_italiccorrection(num)
	box_italiccorrection(den)
	tok_dim(box_toreg(num), num_wd, 0, num_dp)
	tok_dim(box_toreg(den), den_wd, den_ht, 0)
	roff_max(all_wd, num_wd, den_wd)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	tok_len([]byte("\\(ru\x00"), bar_wd, 0, bar_ht, bar_dp)
	// 15b
	noarch.Printf([]byte(".nr %s 0%dm/100u\n\x00"), nregname(num_rise), func() int32 {
		if box_c4go_postfix[0].style == 0 || box_c4go_postfix[0].style == 1 {
			return e_num1
		}
		return e_num2
	}())
	noarch.Printf([]byte(".nr %s 0%dm/100u\n\x00"), nregname(den_fall), func() int32 {
		if box_c4go_postfix[0].style == 0 || box_c4go_postfix[0].style == 1 {
			return e_denom1
		}
		return e_denom2
	}())
	// 15d
	noarch.Printf([]byte(".nr %s (%s-%s)-((%dm/100u)+(%dm/100u/2))\n\x00"), nregname(tmp_15d), nreg(num_rise), nreg(num_dp), e_axisheight, e_rulethickness)
	noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s +(%dm/100u)-%s\n\x00"), nreg(tmp_15d), bargap, nregname(num_rise), bargap, nreg(tmp_15d))
	noarch.Printf([]byte(".nr %s ((%dm/100u)-(%dm/100u/2))-(%s-%s)\n\x00"), nregname(tmp_15d), e_axisheight, e_rulethickness, nreg(den_ht), nreg(den_fall))
	noarch.Printf([]byte(".if %s<(%dm/100u) .nr %s +(%dm/100u)-%s\n\x00"), nreg(tmp_15d), bargap, nregname(den_fall), bargap, nreg(tmp_15d))
	// calculating the vertical position of the bar
	noarch.Printf([]byte(".nr %s 0-%s+%s/2-(%dm/100u)\n\x00"), nregname(bar_fall), nreg(bar_dp), nreg(bar_ht), e_axisheight)
	// making the bar longer
	noarch.Printf([]byte(".nr %s +2*(%dm/100u)\n\x00"), nregname(all_wd), e_overhang)
	// null delimiter space
	box_putf(box_c4go_postfix, []byte("\\h'%sp*%du/100u'\x00"), nreg(box_c4go_postfix[0].szreg), e_nulldelim)
	// drawing the bar
	box_putf(box_c4go_postfix, []byte("\\v'%su'\\f[\\n(.f]\\s[%s]\\l'%su'\\v'-%su'\\h'-%su/2u'\x00"), nreg(bar_fall), nreg(box_c4go_postfix[0].szreg), nreg(all_wd), nreg(bar_fall), nreg(all_wd))
	// output the numerator
	box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\x00"), nreg(num_wd))
	box_putf(box_c4go_postfix, []byte("\\v'-%su'%s\\v'%su'\x00"), nreg(num_rise), box_toreg(num), nreg(num_rise))
	box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\x00"), nreg(num_wd))
	// output the denominator
	box_putf(box_c4go_postfix, []byte("\\h'-%su/2u'\x00"), nreg(den_wd))
	box_putf(box_c4go_postfix, []byte("\\v'%su'%s\\v'-%su'\x00"), nreg(den_fall), box_toreg(den), nreg(den_fall))
	box_putf(box_c4go_postfix, []byte("\\h'(-%su+%su)/2u'\x00"), nreg(den_wd), nreg(all_wd))
	box_putf(box_c4go_postfix, []byte("\\h'%sp*%du/100u'\x00"), nreg(box_c4go_postfix[0].szreg), e_nulldelim)
	box_afterput(box_c4go_postfix, 128)
	box_toreg(box_c4go_postfix)
	nregrm(num_wd)
	nregrm(num_dp)
	nregrm(den_wd)
	nregrm(den_ht)
	nregrm(all_wd)
	nregrm(num_rise)
	nregrm(den_fall)
	nregrm(bar_wd)
	nregrm(bar_dp)
	nregrm(bar_ht)
	nregrm(bar_fall)
	nregrm(tmp_15d)
}

// box_bracketsel - transpiled function from  box.c:515
func box_bracketsel(dst int32, ht int32, dp int32, br [][]byte, any int32, both int32) {
	// choose the smallest bracket among br[], large enough for \n(ht+\n(dp
	var i int32
	for i = 0; br[i] != nil; i++ {
		noarch.Printf([]byte(".if '%s'' \x00"), sreg(dst))
		// is this bracket available?
		noarch.Printf([]byte(".if \\w'%s' \x00"), br[i])
		if both != 0 {
			// check both the height and the depth
			noarch.Printf([]byte(".if (%s-(%dm/100)*2)<=(-\\n[bbury]+\\n[bblly]+(%dm/100*2)) \x00"), nreg(ht), e_rulethickness, e_axisheight)
			noarch.Printf([]byte(".if (%s*2)<=(-\\n[bbury]+\\n[bblly]-(%dm/100*2)) \x00"), nreg(dp), e_axisheight)
		} else {
			noarch.Printf([]byte(".if (%s+%s)<=(-\\n[bbury]+\\n[bblly]) \x00"), nreg(ht), nreg(dp))
		}
		noarch.Printf([]byte(".ds %s \"%s\n\x00"), sregname(dst), br[i])
	}
	if any != 0 {
		for func() int32 {
			i--
			return i
		}() >= 0 {
			// choose the largest bracket, if any is 1
			noarch.Printf([]byte(".if '%s'' .if \\w'%s' .ds %s \"%s\n\x00"), sreg(dst), br[i], sregname(dst), br[i])
		}
	}
}

// box_bracketmk - transpiled function from  box.c:538
func box_bracketmk(dst int32, len_ int32, top []byte, mid []byte, bot []byte, cen []byte) {
	// build a bracket using the provided pieces
	var toplen []int32 = make([]int32, 4)
	var midlen []int32 = make([]int32, 4)
	var botlen []int32 = make([]int32, 4)
	var cenlen []int32 = make([]int32, 4)
	// number of mid glyphs to insert
	var mid_cnt int32 = nregmk()
	// the number of mid glyphs inserted
	var mid_cur int32 = nregmk()
	var cen_pos int32 = nregmk()
	var buildmacro int32 = sregmk()
	blen_mk(top, toplen)
	blen_mk(mid, midlen)
	blen_mk(bot, botlen)
	if cen != nil {
		blen_mk(cen, cenlen)
	}
	if cen == nil {
		// the number of mid tokens necessary to cover sub
		noarch.Printf([]byte(".nr %s %s*2-%s-%s*11/10/%s\n\x00"), nregname(mid_cnt), nreg(len_), nreg(toplen[1]), nreg(botlen[1]), nreg(midlen[1]))
		noarch.Printf([]byte(".if %s<0 .nr %s 0\n\x00"), nreg(mid_cnt), nregname(mid_cnt))
	} else {
		// for brackets with a center like {
		noarch.Printf([]byte(".nr %s %s-(%s+%s+%s/2)*11/10/%s\n\x00"), nregname(cen_pos), nreg(len_), nreg(cenlen[1]), nreg(toplen[1]), nreg(botlen[1]), nreg(midlen[1]))
		noarch.Printf([]byte(".if %s<0 .nr %s 0\n\x00"), nreg(cen_pos), nregname(cen_pos))
		noarch.Printf([]byte(".nr %s 0%s*2\n\x00"), nregname(mid_cnt), nreg(cen_pos))
	}
	// the macro to create the bracket; escaping backslashes
	noarch.Printf([]byte(".de %s\n\x00"), sregname(buildmacro))
	if cen != nil {
		// inserting cen
		noarch.Printf([]byte(".if \\%s=\\%s .as %s \"\\v'-\\%su'%s\\h'-\\%su'\\v'-\\%su'\n\x00"), nreg(mid_cur), nreg(cen_pos), sregname(dst), nreg(cenlen[3]), cen, nreg(cenlen[0]), nreg(cenlen[2]))
	}
	noarch.Printf([]byte(".if \\%s<\\%s .as %s \"\\v'-\\%su'%s\\h'-\\%su'\\v'-\\%su'\n\x00"), nreg(mid_cur), nreg(mid_cnt), sregname(dst), nreg(midlen[3]), mid, nreg(midlen[0]), nreg(midlen[2]))
	noarch.Printf([]byte(".if \\\\n+%s<\\%s .%s\n\x00"), escarg(nregname(mid_cur)), nreg(mid_cnt), sregname(buildmacro))
	fmt.Printf("..\n")
	// constructing the bracket
	noarch.Printf([]byte(".ds %s \"\\v'-%su'%s\\h'-%su'\\v'-%su'\n\x00"), sregname(dst), nreg(botlen[3]), bot, nreg(botlen[0]), nreg(botlen[2]))
	noarch.Printf([]byte(".nr %s 0 1\n\x00"), nregname(mid_cur))
	noarch.Printf([]byte(".%s\n\x00"), sregname(buildmacro))
	noarch.Printf([]byte(".as %s \"\\v'-%su'%s\\h'-%su'\\v'-%su'\n\x00"), sregname(dst), nreg(toplen[3]), top, nreg(toplen[0]), nreg(toplen[2]))
	// moving back vertically
	noarch.Printf([]byte(".as %s \"\\v'%su*%su+%su+%su+%su'\n\x00"), sregname(dst), nreg(mid_cnt), nreg(midlen[1]), nreg(botlen[1]), nreg(toplen[1]), func() []byte {
		if cen != nil {
			return nreg(cenlen[1])
		}
		return []byte("0\x00")
	}())
	// moving right
	noarch.Printf([]byte(".as %s \"\\h'%su'\n\x00"), sregname(dst), func() []byte {
		if cen != nil {
			return nreg(cenlen[0])
		}
		return nreg(midlen[0])
	}())
	blen_rm(toplen)
	blen_rm(midlen)
	blen_rm(botlen)
	if cen != nil {
		blen_rm(cenlen)
	}
	nregrm(mid_cnt)
	nregrm(mid_cur)
	nregrm(cen_pos)
	sregrm(buildmacro)
}

// box_bracket - transpiled function from  box.c:607
func box_bracket(box_c4go_postfix []box, brac []byte, ht int32, dp int32) {
	var sizes [][]byte = [][]byte{nil, nil, nil, nil, nil, nil, nil, nil}
	var top []byte
	var mid []byte
	var bot []byte
	var cen []byte
	var dst int32 = sregmk()
	var len_ int32 = nregmk()
	var fall int32 = nregmk()
	var parlen []int32 = make([]int32, 4)
	roff_max(len_, ht, dp)
	def_sizes(brac, sizes)
	noarch.Printf([]byte(".ds %s \"\n\x00"), sregname(dst))
	def_pieces(brac, (*[10000][]byte)(unsafe.Pointer(&top))[:], (*[10000][]byte)(unsafe.Pointer(&mid))[:], (*[10000][]byte)(unsafe.Pointer(&bot))[:], (*[10000][]byte)(unsafe.Pointer(&cen))[:])
	box_bracketsel(dst, ht, dp, sizes, noarch.BoolToInt(mid == nil), 1)
	if mid != nil {
		noarch.Printf([]byte(".if '%s'' \\{\\\n\x00"), sreg(dst))
		box_bracketmk(dst, len_, top, mid, bot, cen)
		fmt.Printf(".  \\}\n")
	}
	// calculating the total vertical length of the bracket
	blen_mk(sreg(dst), parlen)
	// calculating the amount the bracket should be moved downwards
	noarch.Printf([]byte(".nr %s 0-%s+%s/2-(%sp*%du/100u)\n\x00"), nregname(fall), nreg(parlen[3]), nreg(parlen[2]), nreg(box_c4go_postfix[0].szreg), e_axisheight)
	// printing the output
	box_putf(box_c4go_postfix, []byte("\\f[\\n(.f]\\s[\\n(.s]\\v'%su'%s\\v'-%su'\x00"), nreg(fall), sreg(dst), nreg(fall))
	box_toreg(box_c4go_postfix)
	blen_rm(parlen)
	sregrm(dst)
	nregrm(len_)
	nregrm(fall)
}

// bracsign - transpiled function from  box.c:640
func bracsign(brac []byte, left int32) []byte {
	if int32(brac[0]) == int32('c') && noarch.Not(strncmp([]byte("ceiling\x00"), brac, uint32(noarch.Strlen(brac)))) {
		if left != 0 {
			return []byte("\\(lc\x00")
		}
		return []byte("\\(rc\x00")
	}
	if int32(brac[0]) == int32('f') && noarch.Not(strncmp([]byte("floor\x00"), brac, uint32(noarch.Strlen(brac)))) {
		if left != 0 {
			return []byte("\\(lf\x00")
		}
		return []byte("\\(rf\x00")
	}
	if int32(brac[0]) == int32('<') && int32(brac[1]) == int32('\x00') {
		return []byte("\\(la\x00")
	}
	if int32(brac[0]) == int32('>') && int32(brac[1]) == int32('\x00') {
		return []byte("\\(ra\x00")
	}
	return brac
}

// box_wrap - transpiled function from  box.c:654
func box_wrap(box_c4go_postfix []box, sub []box, left []byte, right []byte) {
	// build large brackets; the correct font should be set up beforehand
	var sublen []int32 = make([]int32, 4)
	blen_mk(box_toreg(sub), sublen)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	if left != nil {
		box_beforeput(box_c4go_postfix, 80, 0)
		box_bracket(box_c4go_postfix, bracsign(left, 1), sublen[2], sublen[3])
		box_afterput(box_c4go_postfix, 80)
	}
	box_merge(box_c4go_postfix, sub, 0)
	if right != nil {
		box_beforeput(box_c4go_postfix, 96, 0)
		box_bracket(box_c4go_postfix, bracsign(right, 0), sublen[2], sublen[3])
		box_afterput(box_c4go_postfix, 96)
	}
	blen_rm(sublen)
}

// sqrt_rad - transpiled function from  box.c:674
func sqrt_rad(dst int32, len_ int32, wd int32) {
	// construct a radical with height at least len and width wd in dst register
	var sizes [][]byte = [][]byte{nil, nil, nil, nil, nil, nil, nil, nil}
	var srlen []int32 = make([]int32, 4)
	var rnlen []int32 = make([]int32, 4)
	var sr_sz int32 = nregmk()
	// if wd is shorter than \(rn
	var wd_diff int32 = nregmk()
	// the right-most horizontal position of \(sr
	var sr_rx int32 = nregmk()
	// horizontal displacement necessary for \(rn
	var rn_dx int32 = nregmk()
	var len2 int32 = nregmk()
	var rad int32 = sregmk()
	var top []byte
	var mid []byte
	var bot []byte
	var cen []byte
	noarch.Printf([]byte(".nr %s 0%s/2*11/10\n\x00"), nregname(len2), nreg(len_))
	noarch.Printf([]byte(".ds %s \"\n\x00"), sregname(rad))
	// selecting a radical of the appropriate size
	def_pieces([]byte("\\(sr\x00"), (*[10000][]byte)(unsafe.Pointer(&top))[:], (*[10000][]byte)(unsafe.Pointer(&mid))[:], (*[10000][]byte)(unsafe.Pointer(&bot))[:], (*[10000][]byte)(unsafe.Pointer(&cen))[:])
	def_sizes([]byte("\\(sr\x00"), sizes)
	box_bracketsel(rad, len2, len2, sizes, 0, 0)
	if mid != nil {
		// constructing the bracket if needed
		noarch.Printf([]byte(".if \\w'%s' \x00"), mid)
		noarch.Printf([]byte(".if '%s'' \\{\\\n\x00"), sreg(rad))
		box_bracketmk(rad, len2, top, mid, bot, nil)
		fmt.Printf(".  \\}\n")
	}
	// enlarging \(sr if no suitable glyph was found
	noarch.Printf([]byte(".if '%s'' \\{\\\n\x00"), sreg(rad))
	blen_mk([]byte("\\(sr\x00"), srlen)
	noarch.Printf([]byte(".ie %s<(%s+%s) .nr %s 0\\n(.s\n\x00"), nreg(len_), nreg(srlen[2]), nreg(srlen[3]), nregname(sr_sz))
	noarch.Printf([]byte(".el .nr %s 0%s*\\n(.s/(%s+%s-(%dm/100u))+1\n\x00"), nregname(sr_sz), nreg(len_), nreg(srlen[2]), nreg(srlen[3]), e_rulethickness)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(sr_sz))
	noarch.Printf([]byte(".ds %s \"\\(sr\n\x00"), sregname(rad))
	blen_rm(srlen)
	fmt.Printf(".  \\}\n")
	// adding the handle
	blen_mk(sreg(rad), srlen)
	noarch.Printf([]byte(".nr %s \\n[bburx]\n\x00"), nregname(sr_rx))
	blen_mk([]byte("\\(rn\x00"), rnlen)
	noarch.Printf([]byte(".nr %s 0%s-\\n[bbllx]-(%dm/100u)\n\x00"), nregname(rn_dx), nreg(sr_rx), e_rulethickness)
	noarch.Printf([]byte(".nr %s 0\n\x00"), nregname(wd_diff))
	noarch.Printf([]byte(".if %s<%s .nr %s 0%s-%s\n\x00"), nreg(wd), nreg(rnlen[0]), nregname(wd_diff), nreg(rnlen[0]), nreg(wd))
	// output the radical; align the top of the radical to the baseline
	noarch.Printf([]byte(".ds %s \"\\s[\\n(.s]\\f[\\n(.f]\\v'%su'\\h'%su'\\l'%su+%su\\(rn'\\h'-%su'\\v'-%su'\\h'-%su-%su'\\v'%su'%s\\v'-%su'\\h'%su+%su'\n\x00"), nregname(dst), nreg(rnlen[2]), nreg(rn_dx), nreg(wd), nreg(wd_diff), nreg(rn_dx), nreg(rnlen[2]), nreg(wd), nreg(wd_diff), nreg(srlen[2]), sreg(rad), nreg(srlen[2]), nreg(wd), nreg(wd_diff))
	blen_rm(srlen)
	blen_rm(rnlen)
	nregrm(sr_sz)
	nregrm(wd_diff)
	nregrm(sr_rx)
	nregrm(rn_dx)
	sregrm(rad)
}

// box_sqrt - transpiled function from  box.c:738
func box_sqrt(box_c4go_postfix []box, sub []box) {
	var sublen []int32 = make([]int32, 4)
	var radlen []int32 = make([]int32, 4)
	var rad int32 = sregmk()
	var rad_rise int32 = nregmk()
	var min_ht int32 = nregmk()
	box_italiccorrection(sub)
	box_beforeput(box_c4go_postfix, 16, 0)
	blen_mk(box_toreg(sub), sublen)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	// 11
	noarch.Printf([]byte(".nr %s 0%s+%s+(2*%dm/100u)+(%dm/100u/4)\n\x00"), nregname(min_ht), nreg(sublen[2]), nreg(sublen[3]), e_rulethickness, func() int32 {
		if box_c4go_postfix[0].style == 0 || box_c4go_postfix[0].style == 1 {
			return e_xheight
		}
		return e_rulethickness
	}())
	sqrt_rad(rad, min_ht, sublen[0])
	blen_mk(sreg(rad), radlen)
	noarch.Printf([]byte(".nr %s 0(%dm/100u)+(%dm/100u/4)\n\x00"), nregname(rad_rise), e_rulethickness, func() int32 {
		if box_c4go_postfix[0].style == 0 || box_c4go_postfix[0].style == 1 {
			return e_xheight
		}
		return e_rulethickness
	}())
	noarch.Printf([]byte(".if %s>(%s+%s+%s) .nr %s (%s+%s-%s-%s)/2\n\x00"), nreg(radlen[3]), nreg(sublen[2]), nreg(sublen[3]), nreg(rad_rise), nregname(rad_rise), nreg(rad_rise), nreg(radlen[3]), nreg(sublen[2]), nreg(sublen[3]))
	noarch.Printf([]byte(".nr %s +%s\n\x00"), nregname(rad_rise), nreg(sublen[2]))
	// output the radical
	box_putf(box_c4go_postfix, []byte("\\v'-%su'%s\\v'%su'\\h'-%su'%s\x00"), nreg(rad_rise), sreg(rad), nreg(rad_rise), nreg(sublen[0]), box_toreg(sub))
	box_afterput(box_c4go_postfix, 16)
	box_toreg(box_c4go_postfix)
	blen_rm(sublen)
	blen_rm(radlen)
	sregrm(rad)
	nregrm(rad_rise)
	nregrm(min_ht)
}

// box_bar - transpiled function from  box.c:778
func box_bar(box_c4go_postfix []box) {
	var box_wd int32 = nregmk()
	var box_ht int32 = nregmk()
	var bar_wd int32 = nregmk()
	var bar_dp int32 = nregmk()
	var bar_rise int32 = nregmk()
	box_italiccorrection(box_c4go_postfix)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	tok_len([]byte("\\(ru\x00"), bar_wd, 0, 0, bar_dp)
	tok_dim(box_toreg(box_c4go_postfix), box_wd, box_ht, 0)
	noarch.Printf([]byte(".if %su<(%dm/100u) .nr %s 0%dm/100u\n\x00"), nreg(box_ht), e_xheight, nregname(box_ht), e_xheight)
	noarch.Printf([]byte(".nr %s 0%su+%su+(3*%dm/100u)\n\x00"), nregname(bar_rise), nreg(box_ht), nreg(bar_dp), e_rulethickness)
	box_putf(box_c4go_postfix, []byte("\\v'-%su'\\s%s\\f[\\n(.f]\\l'-%su\\(ru'\\v'%su'\x00"), nreg(bar_rise), escarg(nreg(box_c4go_postfix[0].szreg)), nreg(box_wd), nreg(bar_rise))
	nregrm(box_wd)
	nregrm(box_ht)
	nregrm(bar_wd)
	nregrm(bar_dp)
	nregrm(bar_rise)
}

// box_accent - transpiled function from  box.c:804
func box_accent(box_c4go_postfix []box, c []byte) {
	var box_wd int32 = nregmk()
	var box_ht int32 = nregmk()
	var ac_rise int32 = nregmk()
	var ac_wd int32 = nregmk()
	var ac_dp int32 = nregmk()
	box_italiccorrection(box_c4go_postfix)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	tok_len(c, ac_wd, 0, 0, ac_dp)
	tok_dim(box_toreg(box_c4go_postfix), box_wd, box_ht, 0)
	noarch.Printf([]byte(".if %su<(%dm/100u) .nr %s 0%dm/100u\n\x00"), nreg(box_ht), e_xheight, nregname(box_ht), e_xheight)
	noarch.Printf([]byte(".nr %s 0%su+%su+(%sp*10u/100u)\n\x00"), nregname(ac_rise), nreg(box_ht), nreg(ac_dp), nreg(box_c4go_postfix[0].szreg))
	box_putf(box_c4go_postfix, []byte("\\v'-%su'\\h'-%su-%su/2u'\\s%s\\f[\\n(.f]%s\\h'%su-%su/2u'\\v'%su'\x00"), nreg(ac_rise), nreg(box_wd), nreg(ac_wd), escarg(nreg(box_c4go_postfix[0].szreg)), c, nreg(box_wd), nreg(ac_wd), nreg(ac_rise))
	nregrm(box_wd)
	nregrm(box_ht)
	nregrm(ac_rise)
	nregrm(ac_wd)
	nregrm(ac_dp)
}

// box_under - transpiled function from  box.c:831
func box_under(box_c4go_postfix []box) {
	var box_wd int32 = nregmk()
	var box_dp int32 = nregmk()
	var bar_wd int32 = nregmk()
	var bar_ht int32 = nregmk()
	var bar_fall int32 = nregmk()
	box_italiccorrection(box_c4go_postfix)
	noarch.Printf([]byte(".ps %s\n\x00"), nreg(box_c4go_postfix[0].szreg))
	tok_len([]byte("\\(ul\x00"), bar_wd, 0, bar_ht, 0)
	tok_dim(box_toreg(box_c4go_postfix), box_wd, 0, box_dp)
	noarch.Printf([]byte(".if %s<0 .nr %s 0\n\x00"), nreg(box_dp), nregname(box_dp))
	noarch.Printf([]byte(".nr %s 0%su+%su+(3*%dm/100u)\n\x00"), nregname(bar_fall), nreg(box_dp), nreg(bar_ht), e_rulethickness)
	box_putf(box_c4go_postfix, []byte("\\v'%su'\\s%s\\f[\\n(.f]\\l'-%su\\(ul'\\v'-%su'\x00"), nreg(bar_fall), escarg(nreg(box_c4go_postfix[0].szreg)), nreg(box_wd), nreg(bar_fall))
	nregrm(box_wd)
	nregrm(box_dp)
	nregrm(bar_wd)
	nregrm(bar_ht)
	nregrm(bar_fall)
}

// box_toreg - transpiled function from  box.c:856
func box_toreg(box_c4go_postfix []box) []byte {
	if noarch.Not(box_c4go_postfix[0].reg) {
		box_c4go_postfix[0].reg = sregmk()
		noarch.Printf([]byte(".ds %s \"%s\n\x00"), sregname(box_c4go_postfix[0].reg), box_buf(box_c4go_postfix))
	}
	return sreg(box_c4go_postfix[0].reg)
}

// box_empty - transpiled function from  box.c:865
func box_empty(box_c4go_postfix []box) int32 {
	return noarch.BoolToInt(noarch.Not(noarch.Strlen(box_buf(box_c4go_postfix))))
}

// box_vcenter - transpiled function from  box.c:870
func box_vcenter(box_c4go_postfix []box, sub []box) {
	var wd int32 = nregmk()
	var ht int32 = nregmk()
	var dp int32 = nregmk()
	var fall int32 = nregmk()
	box_beforeput(box_c4go_postfix, sub[0].tbeg, 0)
	tok_dim(box_toreg(sub), wd, ht, dp)
	noarch.Printf([]byte(".nr %s 0-%s+%s/2-(%sp*%du/100u)\n\x00"), nregname(fall), nreg(dp), nreg(ht), nreg(box_c4go_postfix[0].szreg), e_axisheight)
	box_putf(box_c4go_postfix, []byte("\\v'%su'%s\\v'-%su'\x00"), nreg(fall), box_toreg(sub), nreg(fall))
	box_toreg(box_c4go_postfix)
	box_afterput(box_c4go_postfix, sub[0].tcur)
	nregrm(wd)
	nregrm(ht)
	nregrm(dp)
	nregrm(fall)
}

// box_vertspace - transpiled function from  box.c:891
func box_vertspace(box_c4go_postfix []box) {
	// include line-space requests
	var box_wd int32 = nregmk()
	var htroom int32 = nregmk()
	var dproom int32 = nregmk()
	box_italiccorrection(box_c4go_postfix)
	// amount of room available before and after this line
	noarch.Printf([]byte(".nr %s 0+\\n(.vu-%sp+(%sp*%du/100u)\n\x00"), nregname(htroom), nreg(box_c4go_postfix[0].szreg), nreg(box_c4go_postfix[0].szreg), e_bodyheight)
	noarch.Printf([]byte(".nr %s 0+\\n(.vu-%sp+(%sp*%du/100u)\n\x00"), nregname(dproom), nreg(box_c4go_postfix[0].szreg), nreg(box_c4go_postfix[0].szreg), e_bodydepth)
	// appending \x requests
	tok_dim(box_toreg(box_c4go_postfix), box_wd, 0, 0)
	noarch.Printf([]byte(".if -\\n[bbury]>%s .as %s \"\\x'\\n[bbury]u+%su'\n\x00"), nreg(htroom), sregname(box_c4go_postfix[0].reg), nreg(htroom))
	noarch.Printf([]byte(".if \\n[bblly]>%s .as %s \"\\x'\\n[bblly]u-%su'\n\x00"), nreg(dproom), sregname(box_c4go_postfix[0].reg), nreg(dproom))
	nregrm(box_wd)
	nregrm(htroom)
	nregrm(dproom)
}

// box_markpos - transpiled function from  box.c:916
func box_markpos(box_c4go_postfix []box, reg []byte) {
	// put the current width to the given number register
	box_c4go_postfix[0].tomark = reg
}

// box_colinit - transpiled function from  box.c:922
func box_colinit(pile [][]box, n int32, plen [][]int32, wd int32, ht int32) {
	// initialize the length of a pile or column of a matrix
	var i int32
	for i = 0; i < n; i++ {
		if pile[i] != nil {
			box_italiccorrection(pile[i])
		}
	}
	for i = 0; i < n; i++ {
		blen_mk(func() []byte {
			if pile[i] != nil {
				return box_toreg(pile[i])
			}
			return []byte("\x00")
		}(), plen[i])
	}
	noarch.Printf([]byte(".nr %s 0%s\n\x00"), nregname(wd), nreg(plen[0][0]))
	noarch.Printf([]byte(".nr %s 0%s\n\x00"), nregname(ht), nreg(plen[0][2]))
	{
		// finding the maximum width
		for i = 1; i < n; i++ {
			noarch.Printf([]byte(".if %s>%s .nr %s 0+%s\n\x00"), nreg(plen[i][0]), nreg(wd), nregname(wd), nreg(plen[i][0]))
		}
	}
	{
		// finding the maximum height (vertical length)
		for i = 1; i < n; i++ {
			noarch.Printf([]byte(".if %s+%s>%s .nr %s 0+%s+%s\n\x00"), nreg(plen[i-1][3]), nreg(plen[i][2]), nreg(ht), nregname(ht), nreg(plen[i-1][3]), nreg(plen[i][2]))
		}
	}
	// maximum height and the depth of the last row
	noarch.Printf([]byte(".if %s>%s .nr %s 0+%s\n\x00"), nreg(plen[n-1][3]), nreg(ht), nregname(ht), nreg(plen[n-1][3]))
}

// box_colput - transpiled function from  box.c:952
func box_colput(pile [][]box, n int32, box_c4go_postfix []box, adj int32, plen [][]int32, wd int32, ht int32) {
	// append the give pile to box
	var i int32
	box_putf(box_c4go_postfix, []byte("\\v'-%du*%su/2u'\x00"), n-1, nreg(ht))
	{
		// adding the entries
		for i = 0; i < n; i++ {
			if adj == int32('c') {
				box_putf(box_c4go_postfix, []byte("\\h'%su-%su/2u'\x00"), nreg(wd), nreg(plen[i][0]))
			}
			if adj == int32('r') {
				box_putf(box_c4go_postfix, []byte("\\h'%su-%su'\x00"), nreg(wd), nreg(plen[i][0]))
			}
			box_putf(box_c4go_postfix, []byte("\\v'%su'%s\x00"), func() []byte {
				if i != 0 {
					return nreg(ht)
				}
				return []byte("0\x00")
			}(), func() []byte {
				if pile[i] != nil {
					return box_toreg(pile[i])
				}
				return []byte("\x00")
			}())
			if adj == int32('l') {
				box_putf(box_c4go_postfix, []byte("\\h'-%su'\x00"), nreg(plen[i][0]))
			}
			if adj == int32('c') {
				box_putf(box_c4go_postfix, []byte("\\h'-%su+(%su-%su/2u)'\x00"), nreg(wd), nreg(wd), nreg(plen[i][0]))
			}
			if adj == int32('r') {
				box_putf(box_c4go_postfix, []byte("\\h'-%su'\x00"), nreg(wd))
			}
		}
	}
	box_putf(box_c4go_postfix, []byte("\\v'-%du*%su/2u'\\h'%su'\x00"), n-1, nreg(ht), nreg(wd))
}

// box_coldone - transpiled function from  box.c:979
func box_coldone(pile [][]box, n int32, plen [][]int32) {
	// free the registers allocated for this pile
	var i int32
	for i = 0; i < n; i++ {
		blen_rm(plen[i])
	}
}

// box_colnrows - transpiled function from  box.c:987
func box_colnrows(cols [][]box) int32 {
	// calculate the number of entries in the given pile
	var n int32
	for n < 32 && cols[n] != nil {
		n++
	}
	return n
}

// box_pile - transpiled function from  box.c:995
func box_pile(box_c4go_postfix []box, pile [][]box, adj int32, rowspace int32) {
	var plen [][]int32 = make([][]int32, 32)
	var max_wd int32 = nregmk()
	var max_ht int32 = nregmk()
	var n int32 = box_colnrows(pile)
	box_beforeput(box_c4go_postfix, 128, 0)
	box_colinit(pile, n, plen, max_wd, max_ht)
	// inserting spaces between entries
	noarch.Printf([]byte(".if %s<(%sp*%du/100u) .nr %s (%sp*%du/100u)\n\x00"), nreg(max_ht), nreg(box_c4go_postfix[0].szreg), e_baselinesep, nregname(max_ht), nreg(box_c4go_postfix[0].szreg), e_baselinesep)
	if rowspace != 0 {
		noarch.Printf([]byte(".nr %s +(%sp*%du/100u)\n\x00"), nregname(max_ht), nreg(box_c4go_postfix[0].szreg), rowspace)
	}
	// adding the entries
	box_colput(pile, n, box_c4go_postfix, adj, plen, max_wd, max_ht)
	box_coldone(pile, n, plen)
	box_afterput(box_c4go_postfix, 128)
	box_toreg(box_c4go_postfix)
	nregrm(max_wd)
	nregrm(max_ht)
}

// box_matrix - transpiled function from  box.c:1019
func box_matrix(box_c4go_postfix []box, ncols int32, cols [][][]box, adj []int32, colspace int32, rowspace int32) {
	var plen [][][]int32 = make([][][]int32, 32)
	var wd []int32 = make([]int32, 32)
	var ht []int32 = make([]int32, 32)
	var max_ht int32 = nregmk()
	var max_wd int32 = nregmk()
	var nrows int32
	var i int32
	box_beforeput(box_c4go_postfix, 128, 0)
	for i = 0; i < ncols; i++ {
		if box_colnrows(cols[i]) > nrows {
			nrows = box_colnrows(cols[i])
		}
	}
	for i = 0; i < ncols; i++ {
		wd[i] = nregmk()
	}
	for i = 0; i < ncols; i++ {
		ht[i] = nregmk()
	}
	{
		// initializing the columns
		for i = 0; i < ncols; i++ {
			box_colinit(cols[i], nrows, plen[i], wd[i], ht[i])
		}
	}
	// finding the maximum width and height
	noarch.Printf([]byte(".nr %s 0%s\n\x00"), nregname(max_wd), nreg(wd[0]))
	noarch.Printf([]byte(".nr %s 0%s\n\x00"), nregname(max_ht), nreg(ht[0]))
	for i = 1; i < ncols; i++ {
		noarch.Printf([]byte(".if %s>%s .nr %s 0+%s\n\x00"), nreg(wd[i]), nreg(max_wd), nregname(max_wd), nreg(wd[i]))
	}
	for i = 1; i < ncols; i++ {
		noarch.Printf([]byte(".if %s>%s .nr %s 0+%s\n\x00"), nreg(ht[i]), nreg(max_ht), nregname(max_ht), nreg(ht[i]))
	}
	// inserting spaces between rows
	noarch.Printf([]byte(".if %s<(%sp*%du/100u) .nr %s (%sp*%du/100u)\n\x00"), nreg(max_ht), nreg(box_c4go_postfix[0].szreg), e_baselinesep, nregname(max_ht), nreg(box_c4go_postfix[0].szreg), e_baselinesep)
	if rowspace != 0 {
		noarch.Printf([]byte(".nr %s +(%sp*%du/100u)\n\x00"), nregname(max_ht), nreg(box_c4go_postfix[0].szreg), rowspace)
	}
	{
		// printing the columns
		for i = 0; i < ncols; i++ {
			if i != 0 {
				// space between columns
				box_putf(box_c4go_postfix, []byte("\\h'%sp*%du/100u'\x00"), nreg(box_c4go_postfix[0].szreg), e_columnsep+colspace)
			}
			box_colput(cols[i], nrows, box_c4go_postfix, adj[i], plen[i], max_wd, max_ht)
		}
	}
	box_afterput(box_c4go_postfix, 128)
	box_toreg(box_c4go_postfix)
	for i = 0; i < ncols; i++ {
		box_coldone(cols[i], nrows, plen[i])
	}
	for i = 0; i < ncols; i++ {
		nregrm(ht[i])
	}
	for i = 0; i < ncols; i++ {
		nregrm(wd[i])
	}
	nregrm(max_wd)
	nregrm(max_ht)
}

// def_macros - transpiled function from  def.c:6
// null-terminated list of default macros
var def_macros [][][]byte = [][][]byte{{[]byte("<-\x00"), []byte("\\(<-\x00")}, {[]byte("<=\x00"), []byte("\\(<=\x00")}, {[]byte(">=\x00"), []byte("\\(>=\x00")}, {[]byte("==\x00"), []byte("\\(==\x00")}, {[]byte("->\x00"), []byte("\\(->\x00")}, {[]byte("!=\x00"), []byte("\\(!=\x00")}, {[]byte("+-\x00"), []byte("\\(+-\x00")}, {[]byte("...\x00"), []byte("vcenter roman \"\\ .\\ .\\ .\\ \"\x00")}, {[]byte(",...,\x00"), []byte("roman \",\\ .\\ .\\ .\\ ,\\|\"\x00")}, {[]byte("ALPHA\x00"), []byte("\\(*A\x00")}, {[]byte("BETA\x00"), []byte("\\(*B\x00")}, {[]byte("CHI\x00"), []byte("\\(*X\x00")}, {[]byte("DELTA\x00"), []byte("\\(*D\x00")}, {[]byte("EPSILON\x00"), []byte("\\(*E\x00")}, {[]byte("ETA\x00"), []byte("\\(*Y\x00")}, {[]byte("GAMMA\x00"), []byte("\\(*G\x00")}, {[]byte("IOTA\x00"), []byte("\\(*I\x00")}, {[]byte("KAPPA\x00"), []byte("\\(*K\x00")}, {[]byte("LAMBDA\x00"), []byte("\\(*L\x00")}, {[]byte("MU\x00"), []byte("\\(*M\x00")}, {[]byte("NU\x00"), []byte("\\(*N\x00")}, {[]byte("OMEGA\x00"), []byte("\\(*W\x00")}, {[]byte("OMICRON\x00"), []byte("\\(*O\x00")}, {[]byte("PHI\x00"), []byte("\\(*F\x00")}, {[]byte("PI\x00"), []byte("\\(*P\x00")}, {[]byte("PSI\x00"), []byte("\\(*Q\x00")}, {[]byte("RHO\x00"), []byte("\\(*R\x00")}, {[]byte("SIGMA\x00"), []byte("\\(*S\x00")}, {[]byte("TAU\x00"), []byte("\\(*T\x00")}, {[]byte("THETA\x00"), []byte("\\(*H\x00")}, {[]byte("UPSILON\x00"), []byte("\\(*U\x00")}, {[]byte("XI\x00"), []byte("\\(*C\x00")}, {[]byte("ZETA\x00"), []byte("\\(*Z\x00")}, {[]byte("alpha\x00"), []byte("\\(*a\x00")}, {[]byte("beta\x00"), []byte("\\(*b\x00")}, {[]byte("chi\x00"), []byte("\\(*x\x00")}, {[]byte("delta\x00"), []byte("\\(*d\x00")}, {[]byte("epsilon\x00"), []byte("\\(*e\x00")}, {[]byte("eta\x00"), []byte("\\(*y\x00")}, {[]byte("gamma\x00"), []byte("\\(*g\x00")}, {[]byte("iota\x00"), []byte("\\(*i\x00")}, {[]byte("kappa\x00"), []byte("\\(*k\x00")}, {[]byte("lambda\x00"), []byte("\\(*l\x00")}, {[]byte("mu\x00"), []byte("\\(*m\x00")}, {[]byte("nu\x00"), []byte("\\(*n\x00")}, {[]byte("omega\x00"), []byte("\\(*w\x00")}, {[]byte("omicron\x00"), []byte("\\(*o\x00")}, {[]byte("phi\x00"), []byte("\\(*f\x00")}, {[]byte("pi\x00"), []byte("\\(*p\x00")}, {[]byte("psi\x00"), []byte("\\(*q\x00")}, {[]byte("rho\x00"), []byte("\\(*r\x00")}, {[]byte("sigma\x00"), []byte("\\(*s\x00")}, {[]byte("tau\x00"), []byte("\\(*t\x00")}, {[]byte("theta\x00"), []byte("\\(*h\x00")}, {[]byte("upsilon\x00"), []byte("\\(*u\x00")}, {[]byte("xi\x00"), []byte("\\(*c\x00")}, {[]byte("zeta\x00"), []byte("\\(*z\x00")}, {[]byte("Im\x00"), []byte("roman \"Im\"\x00")}, {[]byte("Re\x00"), []byte("roman \"Re\"\x00")}, {[]byte("and\x00"), []byte("roman \"and\"\x00")}, {[]byte("approx\x00"), []byte("\"\\v'-.2m'\\z\\(ap\\v'.25m'\\(ap\\v'-.05m'\"\x00")}, {[]byte("arc\x00"), []byte("roman \"arc\"\x00")}, {[]byte("cdot\x00"), []byte("\\(c.\x00")}, {[]byte("cos\x00"), []byte("roman \"cos\"\x00")}, {[]byte("cosh\x00"), []byte("roman \"cosh\"\x00")}, {[]byte("coth\x00"), []byte("roman \"coth\"\x00")}, {[]byte("del\x00"), []byte("\\(gr\x00")}, {[]byte("det\x00"), []byte("roman \"det\"\x00")}, {[]byte("dollar\x00"), []byte("roman $\x00")}, {[]byte("exp\x00"), []byte("roman \"exp\"\x00")}, {[]byte("for\x00"), []byte("roman \"for\"\x00")}, {[]byte("grad\x00"), []byte("\\(gr\x00")}, {[]byte("half\x00"), []byte("roman \\(12\x00")}, {[]byte("if\x00"), []byte("roman \"if\"\x00")}, {[]byte("inf\x00"), []byte("\\(if\x00")}, {[]byte("infinity\x00"), []byte("\\(if\x00")}, {[]byte("int\x00"), []byte("{vcenter roman size +2 \\(is}\x00")}, {[]byte("inter\x00"), []byte("roman size +2 \\(ca\x00")}, {[]byte("lim\x00"), []byte("roman \"lim\"\x00")}, {[]byte("ln\x00"), []byte("roman \"ln\"\x00")}, {[]byte("log\x00"), []byte("roman \"log\"\x00")}, {[]byte("max\x00"), []byte("roman \"max\"\x00")}, {[]byte("min\x00"), []byte("roman \"min\"\x00")}, {[]byte("nothing\x00"), []byte("\x00")}, {[]byte("partial\x00"), []byte("\\(pd\x00")}, {[]byte("prime\x00"), []byte("roman \\(fm\x00")}, {[]byte("prod\x00"), []byte("{vcenter roman size +2 \\(pr}\x00")}, {[]byte("sin\x00"), []byte("roman \"sin\"\x00")}, {[]byte("sinh\x00"), []byte("roman \"sinh\"\x00")}, {[]byte("sum\x00"), []byte("{vcenter roman size +2 \\(su}\x00")}, {[]byte("tan\x00"), []byte("roman \"tan\"\x00")}, {[]byte("tanh\x00"), []byte("roman \"tanh\"\x00")}, {[]byte("times\x00"), []byte("\\(mu\x00")}, {[]byte("union\x00"), []byte("roman size +2 \\(cu\x00")}, {nil, nil}}

// binops - transpiled function from  def.c:105
// list of binary operations
var binops [][]byte = [][]byte{[]byte("+\x00"), []byte("\\(pl\x00"), []byte("−\x00"), []byte("-\x00"), []byte("\\(mi\x00"), []byte("÷\x00"), []byte("\\(-:\x00"), []byte("\\(di\x00"), []byte("×\x00"), []byte("xx\x00"), []byte("\\(mu\x00"), []byte("±\x00"), []byte("\\(+-\x00"), []byte("⊗\x00"), []byte("\\(Ox\x00"), []byte("⊕\x00"), []byte("\\(O+\x00"), []byte("∧\x00"), []byte("\\(l&\x00"), []byte("∨\x00"), []byte("\\(l|\x00"), []byte("∩\x00"), []byte("\\(ca\x00"), []byte("∪\x00"), []byte("\\(cu\x00"), []byte("⋅\x00"), []byte("\\(c.\x00")}

// relops - transpiled function from  def.c:121
// list of relations
var relops [][]byte = [][]byte{[]byte("<\x00"), []byte(">\x00"), []byte(":=\x00"), []byte("=\x00"), []byte("\\(eq\x00"), []byte("≅\x00"), []byte("\\(cg\x00"), []byte("≤\x00"), []byte("\\(<=\x00"), []byte("≥\x00"), []byte("\\(>=\x00"), []byte("≠\x00"), []byte("\\(!=\x00"), []byte("≡\x00"), []byte("\\(==\x00"), []byte("≈\x00"), []byte("\\(~~\x00"), []byte("⊃\x00"), []byte("\\(sp\x00"), []byte("⊇\x00"), []byte("\\(ip\x00"), []byte("⊄\x00"), []byte("\\(!b\x00"), []byte("⊂\x00"), []byte("\\(sb\x00"), []byte("⊆\x00"), []byte("\\(ib\x00"), []byte("∈\x00"), []byte("\\(mo\x00"), []byte("∉\x00"), []byte("\\(!m\x00"), []byte("↔\x00"), []byte("\\(ab\x00"), []byte("←\x00"), []byte("\\(<-\x00"), []byte("↑\x00"), []byte("\\(ua\x00"), []byte("→\x00"), []byte("\\(->\x00"), []byte("↓\x00"), []byte("\\(da\x00")}

// puncs - transpiled function from  def.c:145
// list of punctuations
var puncs [][]byte = [][]byte{[]byte(".\x00"), []byte(",\x00"), []byte(";\x00"), []byte(":\x00"), []byte("!\x00")}

// bracketleft - transpiled function from  def.c:148
// left and right brackets
var bracketleft [][]byte = [][]byte{[]byte("(\x00"), []byte("[\x00"), []byte("{\x00"), []byte("\\(lc\x00"), []byte("\\(lf\x00"), []byte("\\(la\x00")}

// bracketright - transpiled function from  def.c:149
var bracketright [][]byte = [][]byte{[]byte(")\x00"), []byte("]\x00"), []byte("}\x00"), []byte("\\(rc\x00"), []byte("\\(rf\x00"), []byte("\\(ra\x00")}

// bracketsizes - transpiled function from  def.c:152
// glyphs for different bracket sizes
var bracketsizes [][][]byte = [][][]byte{{[]byte("(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte(")\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte(")\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenrightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenrightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenrightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'parenrightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte("[\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("[\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte("]\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("]\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketrightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketrightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketrightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracketrightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte("{\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("{\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'braceleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'braceleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'braceleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'braceleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte("}\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("}\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracerightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracerightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracerightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'bracerightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}}, {[]byte("\\(lc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingleft'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("\\(rc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingright'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingrightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingrightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingrightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'ceilingrightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("\\(lf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorleft'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("\\(rf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorright'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorrightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorrightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorrightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'floorrightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("\\(la\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(la\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketleft'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketleftbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketleftBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketleftbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketleftBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("\\(ra\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(ra\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketright'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketrightbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketrightBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketrightbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'angbracketrightBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {[]byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}}, {[]byte("\\(sr\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(sr\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radical'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalbig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalBig'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalbigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalBigg'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// bracketpieces - transpiled function from  def.c:183
// large glyph pieces: name, top, mid, bot, centre
var bracketpieces [][][]byte = [][][]byte{{[]byte("(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(LT\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(LX\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(LB\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte(")\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(RT\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(RX\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(RB\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("[\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("]\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("{\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lt\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(bv\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lb\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lk\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}}, {[]byte("}\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rt\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(bv\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rb\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rk\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}}, {[]byte("\\(lc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("\\(rc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rc\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("\\(lf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(lf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("\\(rf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rx\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\(rf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("|\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {[]byte("\\(sr\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicaltp'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalvertex'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), []byte("\\N'radicalbt'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), {}, {}, {}, {}}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

// gtype - transpiled function from  def.c:199
// custom glyph types
type gtype struct {
	g     [32]byte
	type_ int32
}

// gtypes - transpiled function from  def.c:199
var gtypes []gtype = make([]gtype, 128)

// def_typeput - transpiled function from  def.c:204
func def_typeput(s []byte, type_ int32) {
	var i int32
	for i = 0; uint32(i) < 6144/48 && int32(gtypes[i].g[:][0]) != 0; i++ {
		if noarch.Not(noarch.Strcmp(s, gtypes[i].g[:])) {
			break
		}
	}
	if uint32(i) < 6144/48 {
		noarch.Strcpy(gtypes[i].g[:], s)
		gtypes[i].type_ = type_
	}
}

// alookup - transpiled function from  def.c:217
func alookup(a [][]byte, len_ int32, s []byte) []byte {
	// find an entry in an array
	var i int32
	for i = 0; i < len_; i++ {
		if noarch.Not(noarch.Strcmp(s, a[i])) {
			return a[i]
		}
	}
	return nil
}

// def_type - transpiled function from  def.c:226
func def_type(s []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 6144/48 && int32(gtypes[i].g[:][0]) != 0; i++ {
		if noarch.Not(noarch.Strcmp(s, gtypes[i].g[:])) {
			return gtypes[i].type_
		}
	}
	if alookup(puncs, int32(40/8), s) != nil {
		return 112
	}
	if alookup(binops, int32(216/8), s) != nil {
		return 48
	}
	if alookup(relops, int32(328/8), s) != nil {
		return 64
	}
	if alookup(bracketleft, int32(48/8), s) != nil {
		return 80
	}
	if alookup(bracketright, int32(48/8), s) != nil {
		return 96
	}
	return -1
}

// pieces_find - transpiled function from  def.c:245
func pieces_find(sign []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 16384/512; i++ {
		if noarch.Not(noarch.Strcmp(bracketpieces[i][0], sign)) {
			return i
		}
	}
	return -1
}

// def_pieces - transpiled function from  def.c:255
func def_pieces(sign []byte, top [][]byte, mid [][]byte, bot [][]byte, cen [][]byte) {
	// find the pieces for creating the given bracket
	var i int32 = pieces_find(sign)
	if i >= 0 {
		if int32(bracketpieces[i][1][0]) != 0 {
			top[0] = bracketpieces[i][1]
		} else {
			top[0] = nil
		}
		if int32(bracketpieces[i][2][0]) != 0 {
			mid[0] = bracketpieces[i][2]
		} else {
			mid[0] = nil
		}
		if int32(bracketpieces[i][3][0]) != 0 {
			bot[0] = bracketpieces[i][3]
		} else {
			bot[0] = nil
		}
		if int32(bracketpieces[i][4][0]) != 0 {
			cen[0] = bracketpieces[i][4]
		} else {
			cen[0] = nil
		}
	}
}

// def_piecesput - transpiled function from  def.c:266
func def_piecesput(sign []byte, top []byte, mid []byte, bot []byte, cen []byte) {
	var i int32 = pieces_find(sign)
	if i < 0 && (func() int32 {
		i = pieces_find([]byte("\x00"))
		return i
	}()) < 0 {
		return
	}
	noarch.Snprintf(bracketpieces[i][0], int32(64), []byte("%s\x00"), sign)
	noarch.Snprintf(bracketpieces[i][1], int32(64), []byte("%s\x00"), top)
	noarch.Snprintf(bracketpieces[i][2], int32(64), []byte("%s\x00"), mid)
	noarch.Snprintf(bracketpieces[i][3], int32(64), []byte("%s\x00"), bot)
	noarch.Snprintf(bracketpieces[i][4], int32(64), []byte("%s\x00"), cen)
}

// sizes_find - transpiled function from  def.c:278
func sizes_find(sign []byte) int32 {
	var i int32
	for i = 0; uint32(i) < 16384/512; i++ {
		if noarch.Not(noarch.Strcmp(bracketsizes[i][0], sign)) {
			return i
		}
	}
	return -1
}

// def_sizes - transpiled function from  def.c:288
func def_sizes(sign []byte, sizes [][]byte) {
	// return different sizes of the given bracket
	var idx int32 = sizes_find(sign)
	var i int32
	sizes[0] = sign
	for i = 1; idx >= 0 && i < 8; i++ {
		if int32(bracketsizes[idx][i][0]) != 0 {
			sizes[i-1] = bracketsizes[idx][i]
		} else {
			sizes[i-1] = nil
		}
	}
}

// def_sizesput - transpiled function from  def.c:297
func def_sizesput(sign []byte, sizes [][]byte) {
	var idx int32 = sizes_find(sign)
	var i int32
	if idx < 0 && (func() int32 {
		idx = sizes_find([]byte("\x00"))
		return idx
	}()) < 0 {
		return
	}
	noarch.Snprintf(bracketsizes[idx][0], int32(64), []byte("%s\x00"), sign)
	for i = 1; i < 8; i++ {
		noarch.Snprintf(bracketsizes[idx][i], int32(64), []byte("%s\x00"), func() []byte {
			if sizes[i-1] != nil {
				return sizes[i-1]
			}
			return []byte("\x00")
		}())
	}
}

// e_axisheight - transpiled function from  def.c:310
// global variables
// axis height
var e_axisheight int32 = 23

// e_minimumsize - transpiled function from  def.c:311
// minimum size
var e_minimumsize int32 = 5

// e_overhang - transpiled function from  def.c:312
var e_overhang int32 = 7

// e_nulldelim - transpiled function from  def.c:313
var e_nulldelim int32 = 12

// e_scriptspace - transpiled function from  def.c:314
var e_scriptspace int32 = 12

// e_thinspace - transpiled function from  def.c:315
var e_thinspace int32 = 17

// e_mediumspace - transpiled function from  def.c:316
var e_mediumspace int32 = 22

// e_thickspace - transpiled function from  def.c:317
var e_thickspace int32 = 28

// e_num1 - transpiled function from  def.c:318
// minimum numerator rise
var e_num1 int32 = 70

// e_num2 - transpiled function from  def.c:319
var e_num2 int32 = 40

// e_denom1 - transpiled function from  def.c:320
// minimum denominator fall
var e_denom1 int32 = 70

// e_denom2 - transpiled function from  def.c:321
var e_denom2 int32 = 36

// e_sup1 - transpiled function from  def.c:322
var e_sup1 int32 = 42

// e_sup2 - transpiled function from  def.c:323
var e_sup2 int32 = 37

// e_sup3 - transpiled function from  def.c:324
var e_sup3 int32 = 28

// e_sub1 - transpiled function from  def.c:325
var e_sub1 int32 = 20

// e_sub2 - transpiled function from  def.c:326
var e_sub2 int32 = 23

// e_supdrop - transpiled function from  def.c:327
var e_supdrop int32 = 38

// e_subdrop - transpiled function from  def.c:328
var e_subdrop int32 = 5

// e_xheight - transpiled function from  def.c:329
var e_xheight int32 = 45

// e_rulethickness - transpiled function from  def.c:330
var e_rulethickness int32 = 4

// e_bigopspacing1 - transpiled function from  def.c:331
var e_bigopspacing1 int32 = 11

// e_bigopspacing2 - transpiled function from  def.c:332
var e_bigopspacing2 int32 = 17

// e_bigopspacing3 - transpiled function from  def.c:333
var e_bigopspacing3 int32 = 20

// e_bigopspacing4 - transpiled function from  def.c:334
var e_bigopspacing4 int32 = 60

// e_bigopspacing5 - transpiled function from  def.c:335
var e_bigopspacing5 int32 = 10

// e_columnsep - transpiled function from  def.c:336
var e_columnsep int32 = 100

// e_baselinesep - transpiled function from  def.c:337
var e_baselinesep int32 = 140

// e_bodyheight - transpiled function from  def.c:338
var e_bodyheight int32 = 70

// e_bodydepth - transpiled function from  def.c:339
var e_bodydepth int32 = 25

// gvar - transpiled function from  def.c:341
type gvar struct {
	name []byte
	ref  []int32
}

// gvars - transpiled function from  def.c:341
var gvars []gvar = []gvar{{[]byte("axis_height\x00"), c4goUnsafeConvert_int32(&e_axisheight)}, {[]byte("minimum_size\x00"), c4goUnsafeConvert_int32(&e_minimumsize)}, {[]byte("over_hang\x00"), c4goUnsafeConvert_int32(&e_overhang)}, {[]byte("null_delimiter_space\x00"), c4goUnsafeConvert_int32(&e_nulldelim)}, {[]byte("script_space\x00"), c4goUnsafeConvert_int32(&e_scriptspace)}, {[]byte("thin_space\x00"), c4goUnsafeConvert_int32(&e_thinspace)}, {[]byte("medium_space\x00"), c4goUnsafeConvert_int32(&e_mediumspace)}, {[]byte("thick_space\x00"), c4goUnsafeConvert_int32(&e_thickspace)}, {[]byte("num1\x00"), c4goUnsafeConvert_int32(&e_num1)}, {[]byte("num2\x00"), c4goUnsafeConvert_int32(&e_num2)}, {[]byte("denom1\x00"), c4goUnsafeConvert_int32(&e_denom1)}, {[]byte("denom2\x00"), c4goUnsafeConvert_int32(&e_denom2)}, {[]byte("sup1\x00"), c4goUnsafeConvert_int32(&e_sup1)}, {[]byte("sup2\x00"), c4goUnsafeConvert_int32(&e_sup2)}, {[]byte("sup3\x00"), c4goUnsafeConvert_int32(&e_sup3)}, {[]byte("sub1\x00"), c4goUnsafeConvert_int32(&e_sub1)}, {[]byte("sub2\x00"), c4goUnsafeConvert_int32(&e_sub2)}, {[]byte("sup_drop\x00"), c4goUnsafeConvert_int32(&e_supdrop)}, {[]byte("sub_drop\x00"), c4goUnsafeConvert_int32(&e_subdrop)}, {[]byte("x_height\x00"), c4goUnsafeConvert_int32(&e_xheight)}, {[]byte("default_rule_thickness\x00"), c4goUnsafeConvert_int32(&e_rulethickness)}, {[]byte("big_op_spacing1\x00"), c4goUnsafeConvert_int32(&e_bigopspacing1)}, {[]byte("big_op_spacing2\x00"), c4goUnsafeConvert_int32(&e_bigopspacing2)}, {[]byte("big_op_spacing3\x00"), c4goUnsafeConvert_int32(&e_bigopspacing3)}, {[]byte("big_op_spacing4\x00"), c4goUnsafeConvert_int32(&e_bigopspacing4)}, {[]byte("big_op_spacing5\x00"), c4goUnsafeConvert_int32(&e_bigopspacing5)}, {[]byte("column_sep\x00"), c4goUnsafeConvert_int32(&e_columnsep)}, {[]byte("baseline_sep\x00"), c4goUnsafeConvert_int32(&e_baselinesep)}, {[]byte("body_height\x00"), c4goUnsafeConvert_int32(&e_bodyheight)}, {[]byte("body_depth\x00"), c4goUnsafeConvert_int32(&e_bodydepth)}}

// def_set - transpiled function from  def.c:377
func def_set(name []byte, val int32) {
	var i int32
	for i = 0; uint32(i) < 960/32; i++ {
		if noarch.Not(noarch.Strcmp(gvars[i].name, name)) {
			gvars[i].ref[0] = val
		}
	}
}

// ts_sup - transpiled function from  def.c:386
func ts_sup(style int32) int32 {
	// superscript style
	var sz int32 = func() int32 {
		if 2 < style>>uint64(4)+1 {
			return 2
		}
		return style>>uint64(4) + 1
	}()
	return sz<<uint64(4) | style&1
}

// ts_sub - transpiled function from  def.c:393
func ts_sub(style int32) int32 {
	// subscript style
	var sz int32 = func() int32 {
		if 2 < style>>uint64(4)+1 {
			return 2
		}
		return style>>uint64(4) + 1
	}()
	return sz<<uint64(4) | 1
}

// ts_num - transpiled function from  def.c:400
func ts_num(style int32) int32 {
	// numerator style
	var sz int32
	if style == 0 || style == 1 {
		if style&1 != 0 {
			return 3
		}
		return 2
	}
	if 2 < style>>uint64(4)+1 {
		sz = 2
	} else {
		sz = style>>uint64(4) + 1
	}
	return sz<<uint64(4) | style&1
}

// ts_denom - transpiled function from  def.c:410
func ts_denom(style int32) int32 {
	// denominator style
	var sz int32
	if style == 0 || style == 1 {
		return 3
	}
	if 2 < style>>uint64(4)+1 {
		sz = 2
	} else {
		sz = style>>uint64(4) + 1
	}
	return sz<<uint64(4) | 1
}

// brcost_type - transpiled function from  def.c:420
// extra line-break cost
var brcost_type []int32 = make([]int32, 32)

// brcost_cost - transpiled function from  def.c:421
var brcost_cost []int32 = make([]int32, 32)

// brcost_n - transpiled function from  def.c:422
var brcost_n int32

// def_brcost - transpiled function from  def.c:424
func def_brcost(type_ int32) int32 {
	var i int32
	for i = 0; i < brcost_n; i++ {
		if brcost_type[i] == type_ && brcost_cost[i] > 0 {
			return brcost_cost[i]
		}
	}
	return 100000
}

// def_brcostput - transpiled function from  def.c:433
func def_brcostput(type_ int32, cost int32) {
	var i int32
	if type_ == 0 {
		brcost_n = 0
	}
	for i = 0; i < brcost_n; i++ {
		if brcost_type[i] == type_ {
			break
		}
	}
	if type_ <= 0 || uint32(i+noarch.BoolToInt(i >= brcost_n)) >= 128/4 {
		return
	}
	brcost_type[i] = type_
	brcost_cost[i] = cost
	if i >= brcost_n {
		brcost_n = i + 1
	}
}

// chopped - transpiled function from  def.c:450
// at which characters equations are chopped
var chopped []byte = []byte("^~\"\t\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// def_chopped - transpiled function from  def.c:452
func def_chopped(c int32) int32 {
	return noarch.BoolToInt(len(noarch.Strchr([]byte("\n {}\x00"), c)) != 0 || len(noarch.Strchr(chopped, c)) != 0)
}

// def_choppedset - transpiled function from  def.c:457
func def_choppedset(c []byte) {
	noarch.Strcpy(chopped, c)
}

// gfont - transpiled function from  eqn.c:29
//
// * NEATEQN NEATROFF PREPROCESSOR
// *
// * Copyright (C) 2014-2017 Ali Gholami Rudi <ali at rudi dot ir>
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
// flags passed to eqn_box()
var gfont []byte = []byte("2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// grfont - transpiled function from  eqn.c:30
var grfont []byte = []byte("1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// gbfont - transpiled function from  eqn.c:31
var gbfont []byte = []byte("3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// gsize - transpiled function from  eqn.c:32
var gsize []byte = []byte("\\n[.eqnsz]\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

// eqn_lineup - transpiled function from  eqn.c:33
// the lineup horizontal request
var eqn_lineup []byte = make([]byte, 128)

// eqn_lineupreg - transpiled function from  eqn.c:34
// the number register holding lineup width
var eqn_lineupreg int32

// eqn_mk - transpiled function from  eqn.c:35
// the value of MK
var eqn_mk int32

// eqn_boxuntil - transpiled function from  eqn.c:40
func eqn_boxuntil(box_c4go_postfix []box, sz0 int32, fn0 []byte, delim []byte) int32 {
	// read equations until delim is read
	var sub []box
	for tok_get() != nil && tok_jmp(delim) != 0 {
		if noarch.Not(noarch.Strcmp([]byte("}\x00"), tok_get())) {
			return 1
		}
		sub = eqn_box(box_c4go_postfix[0].style, func() []box {
			if sub != nil {
				return box_c4go_postfix
			}
			return nil
		}(), sz0, fn0)
		box_merge(box_c4go_postfix, sub, 0)
		box_free(sub)
	}
	return 0
}

// sizesub - transpiled function from  eqn.c:54
func sizesub(dst int32, src int32, style int32, src_style int32) {
	if style>>uint64(4) > src_style>>uint64(4) {
		// subscript size
		noarch.Printf([]byte(".nr %s %s*7/10\n\x00"), nregname(dst), nreg(src))
		noarch.Printf([]byte(".if %s<%d .nr %s %d\n\x00"), nreg(dst), e_minimumsize, nregname(dst), e_minimumsize)
	} else {
		noarch.Printf([]byte(".nr %s %s\n\x00"), nregname(dst), nreg(src))
	}
}

// tok_quotes - transpiled function from  eqn.c:66
func tok_quotes(s []byte) []byte {
	if s != nil && int32(s[0]) == int32('"') {
		s[noarch.Strlen(s)-int32(1)] = '\x00'
		return s[0+1:]
	}
	if s != nil {
		return s
	}
	return []byte("\x00")
}

// tok_improve - transpiled function from  eqn.c:75
func tok_improve(s []byte) []byte {
	if s != nil && int32(s[0]) == int32('-') && int32(s[1]) == int32('\x00') {
		return []byte("\\(mi\x00")
	}
	if s != nil && int32(s[0]) == int32('+') && int32(s[1]) == int32('\x00') {
		return []byte("\\(pl\x00")
	}
	if s != nil && int32(s[0]) == int32('\'') && int32(s[1]) == int32('\x00') {
		return []byte("\\(fm\x00")
	}
	return tok_quotes(s)
}

// eqn_bracketsizes - transpiled function from  eqn.c:86
func eqn_bracketsizes() {
	var sign []byte = make([]byte, 64)
	var bufs [][]byte = make([][]byte, 8)
	var sizes [][]byte = [][]byte{nil, nil, nil, nil, nil, nil, nil, nil}
	var n int32
	var i int32
	noarch.Snprintf(sign, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	n = noarch.Atoi(tok_poptext(1))
	for i = 0; i < n; i++ {
		var size []byte = tok_quotes(tok_poptext(1))
		if i < 8 {
			noarch.Snprintf(bufs[i], int32(64), []byte("%s\x00"), size)
			sizes[i] = bufs[i]
		}
	}
	def_sizesput(sign, sizes)
}

// eqn_bracketpieces - transpiled function from  eqn.c:104
func eqn_bracketpieces() {
	var sign []byte = make([]byte, 64)
	var top []byte = make([]byte, 64)
	var mid []byte = make([]byte, 64)
	var bot []byte = make([]byte, 64)
	var cen []byte = make([]byte, 64)
	noarch.Snprintf(sign, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	noarch.Snprintf(top, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	noarch.Snprintf(mid, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	noarch.Snprintf(bot, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	noarch.Snprintf(cen, int32(64), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	def_piecesput(sign, top, mid, bot, cen)
}

// typenum - transpiled function from  eqn.c:115
func typenum(s []byte) int32 {
	if noarch.Not(noarch.Strcmp([]byte("ord\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("ordinary\x00"), s)) {
		return 16
	}
	if noarch.Not(noarch.Strcmp([]byte("op\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("operator\x00"), s)) {
		return 32
	}
	if noarch.Not(noarch.Strcmp([]byte("bin\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("binary\x00"), s)) {
		return 48
	}
	if noarch.Not(noarch.Strcmp([]byte("rel\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("relation\x00"), s)) {
		return 64
	}
	if noarch.Not(noarch.Strcmp([]byte("open\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("opening\x00"), s)) {
		return 80
	}
	if noarch.Not(noarch.Strcmp([]byte("close\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("closing\x00"), s)) {
		return 96
	}
	if noarch.Not(noarch.Strcmp([]byte("punct\x00"), s)) || noarch.Not(noarch.Strcmp([]byte("punctuation\x00"), s)) {
		return 112
	}
	if noarch.Not(noarch.Strcmp([]byte("inner\x00"), s)) {
		return 128
	}
	return 16
}

// eqn_chartype - transpiled function from  eqn.c:137
func eqn_chartype() {
	// read chartype command arguments and perform it
	var gl []byte = make([]byte, 32)
	var type_ []byte = make([]byte, 32)
	noarch.Snprintf(type_, int32(32), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	noarch.Snprintf(gl, int32(32), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	if typenum(type_) >= 0 {
		def_typeput(gl, typenum(type_))
	}
}

// eqn_breakcost - transpiled function from  eqn.c:147
func eqn_breakcost() {
	// read breakcost command arguments and perform it
	var tok []byte = make([]byte, 32)
	var cost int32
	var type_ int32
	noarch.Snprintf(tok, int32(32), []byte("%s\x00"), tok_quotes(tok_poptext(1)))
	cost = noarch.Atoi(tok_poptext(1))
	if noarch.Not(noarch.Strcmp([]byte("any\x00"), tok)) {
		type_ = 0
	} else {
		type_ = typenum(tok)
	}
	if type_ >= 0 {
		def_brcostput(type_, cost)
	}
}

// eqn_commands - transpiled function from  eqn.c:159
func eqn_commands() int32 {
	// read general eqn commands
	var var_ []byte = make([]byte, 1000)
	var sz []byte
	if noarch.Not(tok_jmp([]byte("delim\x00"))) {
		tok_delim()
		return 0
	}
	if noarch.Not(tok_jmp([]byte("define\x00"))) {
		tok_macro()
		return 0
	}
	if noarch.Not(tok_jmp([]byte("gfont\x00"))) {
		noarch.Strcpy(gfont, tok_quotes(tok_poptext(1)))
		return 0
	}
	if noarch.Not(tok_jmp([]byte("grfont\x00"))) {
		noarch.Strcpy(grfont, tok_quotes(tok_poptext(1)))
		return 0
	}
	if noarch.Not(tok_jmp([]byte("gbfont\x00"))) {
		noarch.Strcpy(gbfont, tok_quotes(tok_poptext(1)))
		return 0
	}
	if noarch.Not(tok_jmp([]byte("gsize\x00"))) {
		sz = tok_quotes(tok_poptext(1))
		if int32(sz[0]) == int32('-') || int32(sz[0]) == int32('+') {
			noarch.Sprintf(gsize, []byte("\\n%s%s\x00"), escarg([]byte(".eqnsz\x00")), sz)
		} else {
			noarch.Strcpy(gsize, sz)
		}
		return 0
	}
	if noarch.Not(tok_jmp([]byte("set\x00"))) {
		noarch.Strcpy(var_, tok_poptext(1))
		def_set(var_, noarch.Atoi(tok_poptext(1)))
		return 0
	}
	if noarch.Not(tok_jmp([]byte("bracketsizes\x00"))) {
		eqn_bracketsizes()
		return 0
	}
	if noarch.Not(tok_jmp([]byte("bracketpieces\x00"))) {
		eqn_bracketpieces()
		return 0
	}
	if noarch.Not(tok_jmp([]byte("chartype\x00"))) {
		eqn_chartype()
		return 0
	}
	if noarch.Not(tok_jmp([]byte("breakcost\x00"))) {
		eqn_breakcost()
		return 0
	}
	return 1
}

// tok_font - transpiled function from  eqn.c:236
func tok_font(tok int32, fn []byte) []byte {
	if fn != nil && int32(fn[0]) != 0 {
		// read user-specified spaces
		// return the font of the given token type
		return fn
	}
	if tok == 17 || tok == 19 {
		return gfont
	}
	return grfont
}

// tok_expect - transpiled function from  eqn.c:246
func tok_expect(s []byte) {
	if tok_jmp(s) != 0 {
		// check the next token
		noarch.Fprintf(noarch.Stderr, []byte("neateqn: expected %s bot got %s\n\x00"), s, tok_get())
		noarch.Exit(1)
	}
}

// eqn_pile - transpiled function from  eqn.c:256
func eqn_pile(box_c4go_postfix []box, sz0 int32, fn0 []byte, adj int32) {
	// read pile command
	var pile [][]box = [][]box{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	var i int32
	var n int32
	var rowspace int32
	if tok_jmp([]byte("{\x00")) != 0 {
		rowspace = noarch.Atoi(tok_poptext(1))
		tok_expect([]byte("{\x00"))
	}
	for {
		pile[func() int32 {
			defer func() {
				n++
			}()
			return n
		}()] = box_alloc(sz0, 0, box_c4go_postfix[0].style)
		if !noarch.Not(eqn_boxuntil(pile[n-1], sz0, fn0, []byte("above\x00"))) {
			break
		}
	}
	tok_expect([]byte("}\x00"))
	box_pile(box_c4go_postfix, pile, adj, rowspace)
	for i = 0; i < n; i++ {
		box_free(pile[i])
	}
}

// eqn_matrix - transpiled function from  eqn.c:276
func eqn_matrix(box_c4go_postfix []box, sz0 int32, fn0 []byte) {
	// read matrix command
	var cols [][][]box = [][][]box{{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	var adj []int32 = make([]int32, 32)
	var nrows int32
	var ncols int32
	var colspace int32
	var rowspace int32
	var i int32
	var j int32
	if tok_jmp([]byte("{\x00")) != 0 {
		colspace = noarch.Atoi(tok_poptext(1))
		tok_expect([]byte("{\x00"))
	}
	for 1 != 0 {
		if noarch.Not(tok_jmp([]byte("col\x00"))) || noarch.Not(tok_jmp([]byte("ccol\x00"))) {
			adj[ncols] = int32('c')
		} else if noarch.Not(tok_jmp([]byte("lcol\x00"))) {
			adj[ncols] = int32('l')
		} else if noarch.Not(tok_jmp([]byte("rcol\x00"))) {
			adj[ncols] = int32('r')
		} else {
			break
		}
		nrows = 0
		if tok_jmp([]byte("{\x00")) != 0 {
			i = noarch.Atoi(tok_poptext(1))
			if i > rowspace {
				rowspace = i
			}
			tok_expect([]byte("{\x00"))
		}
		for {
			cols[ncols][func() int32 {
				defer func() {
					nrows++
				}()
				return nrows
			}()] = box_alloc(sz0, 0, box_c4go_postfix[0].style)
			if !noarch.Not(eqn_boxuntil(cols[ncols][nrows-1], sz0, fn0, []byte("above\x00"))) {
				break
			}
		}
		tok_expect([]byte("}\x00"))
		ncols++
	}
	tok_expect([]byte("}\x00"))
	box_matrix(box_c4go_postfix, ncols, cols, adj, colspace, rowspace)
	for i = 0; i < ncols; i++ {
		for j = 0; j < 32; j++ {
			if cols[i][j] != nil {
				box_free(cols[i][j])
			}
		}
	}
}

// italic - transpiled function from  eqn.c:321
func italic(fn []byte) int32 {
	// return nonzero if fn is italic
	if noarch.Not(noarch.Strcmp([]byte("I\x00"), fn)) || noarch.Not(noarch.Strcmp([]byte("2\x00"), fn)) || (int64(uintptr(unsafe.Pointer(&gfont[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)) == 0 || noarch.Not(noarch.Strcmp(gfont, fn)) {
		return 256
	}
	return 0
}

// eqn_left - transpiled function from  eqn.c:328
func eqn_left(flg int32, pre []box, sz0 int32, fn0 []byte) []box {
	// read a box without fractions
	var box_c4go_postfix []box
	var sub_sub []box
	var sub_sup []box
	var sub_from []box
	var sub_to []box
	var sqrt []box
	var inner []box
	var left []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var right []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var fn []byte = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var sz int32 = sz0
	var subsz int32
	var dx int32
	var dy int32
	var style int32 = 65535 & flg
	if fn0 != nil {
		noarch.Strcpy(fn, fn0)
	}
	for noarch.Not(eqn_commands()) {
	}
	box_c4go_postfix = box_alloc(sz, func() int32 {
		if pre != nil {
			return pre[0].tcur
		}
		return 0
	}(), style)
	if noarch.Not(eqn_gaps(box_c4go_postfix, sz)) {
		for noarch.Not(eqn_gaps(box_c4go_postfix, sz)) {
		}
		return box_c4go_postfix
	}
	for 1 != 0 {
		if noarch.Not(tok_jmp([]byte("fat\x00"))) {
		} else if noarch.Not(tok_jmp([]byte("roman\x00"))) {
			noarch.Strcpy(fn, grfont)
		} else if noarch.Not(tok_jmp([]byte("italic\x00"))) {
			noarch.Strcpy(fn, gfont)
		} else if noarch.Not(tok_jmp([]byte("bold\x00"))) {
			noarch.Strcpy(fn, gbfont)
		} else if noarch.Not(tok_jmp([]byte("font\x00"))) {
			noarch.Strcpy(fn, tok_poptext(1))
		} else if noarch.Not(tok_jmp([]byte("size\x00"))) {
			sz = box_size(box_c4go_postfix, tok_poptext(1))
		} else if noarch.Not(tok_jmp([]byte("fwd\x00"))) {
			dx += noarch.Atoi(tok_poptext(1))
		} else if noarch.Not(tok_jmp([]byte("back\x00"))) {
			dx -= noarch.Atoi(tok_poptext(1))
		} else if noarch.Not(tok_jmp([]byte("down\x00"))) {
			dy += noarch.Atoi(tok_poptext(1))
		} else if noarch.Not(tok_jmp([]byte("up\x00"))) {
			dy -= noarch.Atoi(tok_poptext(1))
		} else {
			break
		}
	}
	if noarch.Not(tok_jmp([]byte("sqrt\x00"))) {
		sqrt = eqn_left(style|1, nil, sz, fn)
		noarch.Printf([]byte(".ft %s\n\x00"), grfont)
		box_sqrt(box_c4go_postfix, sqrt)
		box_free(sqrt)
	} else if noarch.Not(tok_jmp([]byte("pile\x00"))) || noarch.Not(tok_jmp([]byte("cpile\x00"))) {
		eqn_pile(box_c4go_postfix, sz, fn, int32('c'))
	} else if noarch.Not(tok_jmp([]byte("lpile\x00"))) {
		eqn_pile(box_c4go_postfix, sz, fn, int32('l'))
	} else if noarch.Not(tok_jmp([]byte("rpile\x00"))) {
		eqn_pile(box_c4go_postfix, sz, fn, int32('r'))
	} else if noarch.Not(tok_jmp([]byte("matrix\x00"))) {
		eqn_matrix(box_c4go_postfix, sz, fn)
	} else if noarch.Not(tok_jmp([]byte("vcenter\x00"))) {
		inner = eqn_left(flg, pre, sz, fn)
		box_vcenter(box_c4go_postfix, inner)
		box_free(inner)
	} else if noarch.Not(tok_jmp([]byte("{\x00"))) {
		eqn_boxuntil(box_c4go_postfix, sz, fn, []byte("}\x00"))
	} else if noarch.Not(tok_jmp([]byte("left\x00"))) {
		inner = box_alloc(sz, 0, style)
		noarch.Snprintf(left, int32(32), []byte("%s\x00"), tok_quotes(tok_poptext(0)))
		eqn_boxuntil(inner, sz, fn, []byte("right\x00"))
		noarch.Snprintf(right, int32(32), []byte("%s\x00"), tok_quotes(tok_poptext(0)))
		noarch.Printf([]byte(".ft %s\n\x00"), grfont)
		box_wrap(box_c4go_postfix, inner, func() []byte {
			if int32(left[0]) != 0 {
				return left
			}
			return nil
		}(), func() []byte {
			if int32(right[0]) != 0 {
				return right
			}
			return nil
		}())
		box_free(inner)
	} else if tok_get() != nil && tok_type() != 3 {
		if dx != 0 || dy != 0 {
			box_move(box_c4go_postfix, dy, dx)
		}
		box_putf(box_c4go_postfix, []byte("\\s%s\x00"), escarg(nreg(sz)))
		for {
			var cfn []byte = tok_font(tok_type(), fn)
			var chops int32
			box_puttext(box_c4go_postfix, tok_type()|italic(cfn), []byte("\\f%s%s\x00"), escarg(cfn), tok_improve(tok_get()))
			chops = tok_chops(0)
			tok_pop()
			if chops != 0 {
				// what we read was a splitting
				break
			}
			if !noarch.Not(tok_chops(0)) {
				break
			}
		}
		if dx != 0 || dy != 0 {
			// the next token is splitting
			box_move(box_c4go_postfix, -dy, -dx)
		}
	}
	for tok_get() != nil {
		if noarch.Not(tok_jmp([]byte("dyad\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte("\\(ab\x00"))
		} else if noarch.Not(tok_jmp([]byte("bar\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_bar(box_c4go_postfix)
		} else if noarch.Not(tok_jmp([]byte("under\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_under(box_c4go_postfix)
		} else if noarch.Not(tok_jmp([]byte("vec\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte("\\s[\\n(.s/2u]\\(->\\s0\x00"))
		} else if noarch.Not(tok_jmp([]byte("tilde\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte("\\s[\\n(.s*3u/4u]\\(ap\\s0\x00"))
		} else if noarch.Not(tok_jmp([]byte("hat\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte("ˆ\x00"))
		} else if noarch.Not(tok_jmp([]byte("dot\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte(".\x00"))
		} else if noarch.Not(tok_jmp([]byte("dotdot\x00"))) {
			noarch.Printf([]byte(".ft %s\n\x00"), grfont)
			box_accent(box_c4go_postfix, []byte("..\x00"))
		} else {
			break
		}
	}
	subsz = nregmk()
	if noarch.Not(tok_jmp([]byte("sub\x00"))) {
		sizesub(subsz, sz0, ts_sup(style), style)
		sub_sub = eqn_left(ts_sup(style)|131072, nil, subsz, fn0)
	}
	if (sub_sub != nil || noarch.Not(flg&131072)) && noarch.Not(tok_jmp([]byte("sup\x00"))) {
		sizesub(subsz, sz0, ts_sub(style), style)
		sub_sup = eqn_left(ts_sub(style), nil, subsz, fn0)
	}
	if len(sub_sub) == 0 || len(sub_sup) == 0 {
		box_sub(box_c4go_postfix, sub_sub, sub_sup)
	}
	if noarch.Not(tok_jmp([]byte("from\x00"))) {
		sizesub(subsz, sz0, ts_sub(style), style)
		sub_from = eqn_left(ts_sub(style)|262144, nil, subsz, fn0)
	}
	if (sub_from != nil || noarch.Not(flg&262144)) && noarch.Not(tok_jmp([]byte("to\x00"))) {
		sizesub(subsz, sz0, ts_sup(style), style)
		sub_to = eqn_left(ts_sup(style), nil, subsz, fn0)
	}
	if len(sub_from) == 0 || len(sub_to) == 0 {
		inner = box_alloc(sz0, 0, style)
		box_from(inner, box_c4go_postfix, sub_from, sub_to)
		box_free(box_c4go_postfix)
		box_c4go_postfix = inner
	}
	nregrm(subsz)
	if sub_sub != nil {
		box_free(sub_sub)
	}
	if sub_sup != nil {
		box_free(sub_sup)
	}
	if sub_from != nil {
		box_free(sub_from)
	}
	if sub_to != nil {
		box_free(sub_to)
	}
	return box_c4go_postfix
}

// eqn_box - transpiled function from  eqn.c:486
func eqn_box(flg int32, pre []box, sz0 int32, fn0 []byte) []box {
	// read a box
	var box_c4go_postfix []box
	var sub_num []box
	var sub_den []box
	var style int32 = flg & 65535
	box_c4go_postfix = eqn_left(flg, pre, sz0, fn0)
	for noarch.Not(tok_jmp([]byte("over\x00"))) {
		sub_num = box_c4go_postfix
		sub_den = eqn_left(style|1, nil, sz0, fn0)
		box_c4go_postfix = box_alloc(sz0, func() int32 {
			if pre != nil {
				return pre[0].tcur
			}
			return 0
		}(), style)
		noarch.Printf([]byte(".ft %s\n\x00"), grfont)
		box_over(box_c4go_postfix, sub_num, sub_den)
		box_free(sub_num)
		box_free(sub_den)
	}
	return box_c4go_postfix
}

// eqn_read - transpiled function from  eqn.c:505
func eqn_read(style int32) []box {
	// read an equation, either inline or block
	var box_c4go_postfix []box
	var sub []box
	var szreg int32 = nregmk()
	noarch.Printf([]byte(".nr %s %s\n\x00"), nregname(szreg), gsize)
	box_c4go_postfix = box_alloc(szreg, 0, style)
	for tok_get() != nil {
		if noarch.Not(tok_jmp([]byte("mark\x00"))) {
			if noarch.Not(eqn_mk) {
				eqn_mk = 1
			} else {
				eqn_mk = eqn_mk
			}
			box_markpos(box_c4go_postfix, []byte(".eqnmk\x00"))
			continue
		}
		if noarch.Not(tok_jmp([]byte("lineup\x00"))) {
			eqn_mk = 2
			box_markpos(box_c4go_postfix, nregname(eqn_lineupreg))
			noarch.Sprintf(eqn_lineup, []byte("\\h'\\n%su-%su'\x00"), escarg([]byte(".eqnmk\x00")), nreg(eqn_lineupreg))
			continue
		}
		sub = eqn_box(style, box_c4go_postfix, szreg, nil)
		box_merge(box_c4go_postfix, sub, 1)
		box_free(sub)
	}
	box_vertspace(box_c4go_postfix)
	nregrm(szreg)
	return box_c4go_postfix
}

// errdie - transpiled function from  eqn.c:533
func errdie(msg []byte) {
	noarch.Fprintf(noarch.Stderr, msg)
	noarch.Exit(1)
}

// main - transpiled function from  eqn.c:539
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	var box_c4go_postfix []box
	var eqnblk []byte = make([]byte, 128)
	var i int32
	for i = 1; i < argc; i++ {
		if int32(argv[i][0]) != int32('-') || noarch.Not(argv[i][1]) {
			break
		}
		if int32(argv[i][1]) == int32('c') {
			def_choppedset(func() []byte {
				if int32(argv[i][2]) != 0 {
					return (argv[i])[0+2:]
				}
				return argv[func() int32 {
					i++
					return i
				}()]
			}())
		} else {
			fmt.Printf("Usage: neateqn [options] <input >output\n\n")
			fmt.Printf("Options:\n")
			fmt.Printf("  -c chars  \tcharacters that chop equations\n")
			noarch.Exit(int32(1))
		}
	}
	for i = 0; def_macros[i][0] != nil; i++ {
		src_define(def_macros[i][0], def_macros[i][1])
	}
	for noarch.Not(tok_eqn()) {
		reg_reset()
		eqn_mk = 0
		tok_pop()
		noarch.Printf([]byte(".nr %s \\n(.s\n\x00"), []byte(".eqnsz\x00"))
		noarch.Printf([]byte(".nr %s \\n(.f\n\x00"), []byte(".eqnfn\x00"))
		eqn_lineupreg = nregmk()
		box_c4go_postfix = eqn_read(func() int32 {
			if tok_inline() != 0 {
				return 2
			}
			return 0
		}())
		noarch.Printf([]byte(".nr MK %d\n\x00"), eqn_mk)
		if noarch.Not(box_empty(box_c4go_postfix)) {
			noarch.Sprintf(eqnblk, []byte("%s%s\x00"), eqn_lineup, box_toreg(box_c4go_postfix))
			tok_eqnout(eqnblk)
			noarch.Printf([]byte(".ps \\n%s\n\x00"), escarg([]byte(".eqnsz\x00")))
			noarch.Printf([]byte(".ft \\n%s\n\x00"), escarg([]byte(".eqnfn\x00")))
		}
		noarch.Printf([]byte(".lf %d\n\x00"), src_lineget())
		eqn_lineup[0] = '\x00'
		nregrm(eqn_lineupreg)
		box_free(box_c4go_postfix)
	}
	src_done()
	return
}

// sreg_max - transpiled function from  reg.c:9
// maximum allocated string register
var sreg_max int32

// sreg_free - transpiled function from  reg.c:10
// free string registers
var sreg_free []int32 = make([]int32, 2048)

// sreg_n - transpiled function from  reg.c:11
// number of items in sreg_free[]
var sreg_n int32

// sreg_name - transpiled function from  reg.c:12
var sreg_name [][]byte = make([][]byte, 2048)

// sreg_read - transpiled function from  reg.c:13
var sreg_read [][]byte = make([][]byte, 2048)

// nreg_max - transpiled function from  reg.c:15
var nreg_max int32

// nreg_free - transpiled function from  reg.c:16
var nreg_free []int32 = make([]int32, 2048)

// nreg_n - transpiled function from  reg.c:17
var nreg_n int32

// nreg_name - transpiled function from  reg.c:18
var nreg_name [][]byte = make([][]byte, 2048)

// nreg_read - transpiled function from  reg.c:19
var nreg_read [][]byte = make([][]byte, 2048)

// sregmk - transpiled function from  reg.c:22
func sregmk() int32 {
	// allocate a troff string register
	var id int32 = func() int32 {
		if sreg_n != 0 {
			return sreg_free[func() int32 {
				sreg_n--
				return sreg_n
			}()]
		}
		sreg_max++
		return sreg_max
	}()
	noarch.Sprintf(sreg_name[id], []byte("%s%02d\x00"), []byte("\x00"), id)
	noarch.Sprintf(sreg_read[id], []byte("\\*%s\x00"), escarg(sreg_name[id]))
	return id
}

// sregrm - transpiled function from  reg.c:31
func sregrm(id int32) {
	// free a troff string register
	sreg_free[func() int32 {
		defer func() {
			sreg_n++
		}()
		return sreg_n
	}()] = id
}

// sregname - transpiled function from  reg.c:36
func sregname(id int32) []byte {
	return sreg_name[id]
}

// sreg - transpiled function from  reg.c:41
func sreg(id int32) []byte {
	return sreg_read[id]
}

// nregmk - transpiled function from  reg.c:47
func nregmk() int32 {
	// allocate a troff number register
	var id int32 = func() int32 {
		if nreg_n != 0 {
			return nreg_free[func() int32 {
				nreg_n--
				return nreg_n
			}()]
		}
		nreg_max++
		return nreg_max
	}()
	noarch.Sprintf(nreg_name[id], []byte("%s%02d\x00"), []byte("\x00"), id)
	noarch.Sprintf(nreg_read[id], []byte("\\n%s\x00"), escarg(nreg_name[id]))
	return id
}

// nregrm - transpiled function from  reg.c:56
func nregrm(id int32) {
	// free a troff number register
	nreg_free[func() int32 {
		defer func() {
			nreg_n++
		}()
		return nreg_n
	}()] = id
}

// nregname - transpiled function from  reg.c:61
func nregname(id int32) []byte {
	return nreg_name[id]
}

// nreg - transpiled function from  reg.c:66
func nreg(id int32) []byte {
	return nreg_read[id]
}

// reg_reset - transpiled function from  reg.c:72
func reg_reset() {
	// free all allocated registers
	nreg_max = 0
	nreg_n = 0
	sreg_max = 11
	sreg_n = 0
}

// escarg - transpiled function from  reg.c:81
func escarg(arg []byte) []byte {
	// format the argument of a troff escape like \s or \f
	var buf []byte = make([]byte, 256)
	if noarch.Not(arg[1]) {
		noarch.Sprintf(buf, []byte("%c\x00"), int32(arg[0]))
	} else if noarch.Not(arg[2]) {
		noarch.Sprintf(buf, []byte("(%c%c\x00"), int32(arg[0]), int32(arg[1]))
	} else {
		noarch.Sprintf(buf, []byte("[%s]\x00"), arg)
	}
	return buf
}

// sbuf_extend - transpiled function from  sbuf.c:9
func sbuf_extend(sbuf_c4go_postfix []sbuf, amount int32) {
	var s []byte = sbuf_c4go_postfix[0].s
	sbuf_c4go_postfix[0].sz = (func() int32 {
		if 1 < amount {
			return amount
		}
		return 1
	}() + 512 - 1) & ^(512 - 1)
	sbuf_c4go_postfix[0].s = make([]byte, uint32(sbuf_c4go_postfix[0].sz))
	if sbuf_c4go_postfix[0].n != 0 {
		memcpy(sbuf_c4go_postfix[0].s, s, uint32(sbuf_c4go_postfix[0].n))
	}
	_ = s
}

// sbuf_init - transpiled function from  sbuf.c:19
func sbuf_init(sbuf_c4go_postfix []sbuf) {
	noarch.Memset((*[10000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&sbuf_c4go_postfix[0]))) / int64(1))))[:], byte(0), 24)
	sbuf_extend(sbuf_c4go_postfix, 512)
}

// sbuf_add - transpiled function from  sbuf.c:25
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

// sbuf_append - transpiled function from  sbuf.c:32
func sbuf_append(sbuf_c4go_postfix []sbuf, s []byte) {
	var len_ int32 = noarch.Strlen(s)
	if sbuf_c4go_postfix[0].n+len_+1 >= sbuf_c4go_postfix[0].sz {
		sbuf_extend(sbuf_c4go_postfix, sbuf_c4go_postfix[0].n+len_+1)
	}
	memcpy(sbuf_c4go_postfix[0].s[0+sbuf_c4go_postfix[0].n:], s, uint32(len_))
	sbuf_c4go_postfix[0].n += len_
}

// sbuf_printf - transpiled function from  sbuf.c:41
func sbuf_printf(sbuf_c4go_postfix []sbuf, s []byte, c4goArgs ...interface{}) {
	var buf []byte = make([]byte, 1000)
	var ap *va_list
	va_start(ap, s)
	noarch.Vsprintf(buf, s, ap)
	va_end(ap)
	sbuf_append(sbuf_c4go_postfix, buf)
}

// sbuf_empty - transpiled function from  sbuf.c:51
func sbuf_empty(sbuf_c4go_postfix []sbuf) int32 {
	return noarch.BoolToInt(noarch.Not(sbuf_c4go_postfix[0].n))
}

// sbuf_buf - transpiled function from  sbuf.c:56
func sbuf_buf(sbuf_c4go_postfix []sbuf) []byte {
	sbuf_c4go_postfix[0].s[sbuf_c4go_postfix[0].n] = '\x00'
	return sbuf_c4go_postfix[0].s
}

// sbuf_len - transpiled function from  sbuf.c:62
func sbuf_len(sbuf_c4go_postfix []sbuf) int32 {
	return sbuf_c4go_postfix[0].n
}

// sbuf_cut - transpiled function from  sbuf.c:68
func sbuf_cut(sbuf_c4go_postfix []sbuf, n int32) {
	if sbuf_c4go_postfix[0].n > n {
		// shorten the sbuf
		sbuf_c4go_postfix[0].n = n
	}
}

// sbuf_done - transpiled function from  sbuf.c:74
func sbuf_done(sbuf_c4go_postfix []sbuf) {
	_ = sbuf_c4go_postfix[0].s
}

// esrc - transpiled function from  src.c:12
// reading input
// eqn input stream
type esrc struct {
	prev  []esrc
	buf   []byte
	pos   int32
	unbuf [1000]int32
	uncnt int32
	args  [10][]byte
	call  int32
}

// esrc_stdin - transpiled function from  src.c:22
// previous buffer
// input buffer; NULL for stdin
// current position in buf
// push-back buffer
// macro arguments
// is a macro call
// the default input stream
var esrc_stdin esrc

// esrc_c4go_postfix - transpiled function from  src.c:23
var esrc_c4go_postfix []esrc = c4goUnsafeConvert_esrc(&esrc_stdin)

// lineno - transpiled function from  src.c:24
// current line number
var lineno int32 = 1

// esrc_depth - transpiled function from  src.c:25
// the length of esrc chain
var esrc_depth int32

// src_strdup - transpiled function from  src.c:27
func src_strdup(s []byte) []byte {
	var d []byte = make([]byte, noarch.Strlen(s)+int32(1))
	noarch.Strcpy(d, s)
	return d
}

// src_push - transpiled function from  src.c:35
func src_push(buf []byte, args [][]byte) {
	// push buf in the input stream; this is a macro call if args is not NULL
	var next []esrc
	var i int32
	if esrc_depth > 512 {
		errdie([]byte("neateqn: macro recursion limit reached\n\x00"))
	}
	next = (*[10000]esrc)(unsafe.Pointer(uintptr(func() int64 {
		c4go_temp_name := make([]uint32, 1)
		return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	}())))[:]
	noarch.Memset((*[10000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&next[0]))) / int64(1))))[:], byte(0), 4136)
	next[0].prev = esrc_c4go_postfix
	next[0].buf = src_strdup(buf)
	next[0].call = noarch.BoolToInt(len(args) != 0)
	if args != nil {
		for i = 0; i < 10; i++ {
			if args[i] != nil {
				next[0].args[:][i] = src_strdup(args[i])
			} else {
				next[0].args[:][i] = nil
			}
		}
	}
	esrc_c4go_postfix = next
	esrc_depth++
}

// src_pop - transpiled function from  src.c:54
func src_pop() {
	// back to the previous esrc buffer
	var prev []esrc = esrc_c4go_postfix[0].prev
	var i int32
	if prev != nil {
		for i = 0; i < 10; i++ {
			_ = esrc_c4go_postfix[0].args[:][i]
		}
		_ = esrc_c4go_postfix[0].buf
		_ = esrc_c4go_postfix
		esrc_c4go_postfix = prev
		esrc_depth--
	}
}

// src_stdin - transpiled function from  src.c:68
func src_stdin() int32 {
	var c int32 = noarch.Fgetc(noarch.Stdin)
	if c == int32('\n') {
		lineno++
	}
	return c
}

// src_next - transpiled function from  src.c:77
func src_next() int32 {
	for 1 != 0 {
		if esrc_c4go_postfix[0].uncnt != 0 {
			// read the next character
			return esrc_c4go_postfix[0].unbuf[:][func() int32 {
				tempVar1 := &esrc_c4go_postfix[0].uncnt
				*tempVar1--
				return *tempVar1
			}()]
		}
		if esrc_c4go_postfix[0].prev == nil {
			return src_stdin()
		}
		if esrc_c4go_postfix[0].buf[esrc_c4go_postfix[0].pos] != 0 {
			return int32(uint8(esrc_c4go_postfix[0].buf[func() int32 {
				tempVar1 := &esrc_c4go_postfix[0].pos
				defer func() {
					*tempVar1++
				}()
				return *tempVar1
			}()]))
		}
		src_pop()
	}
	return 0
}

// src_back - transpiled function from  src.c:92
func src_back(c int32) {
	if c > 0 {
		// push back c
		esrc_c4go_postfix[0].unbuf[:][func() int32 {
			tempVar1 := &esrc_c4go_postfix[0].uncnt
			defer func() {
				*tempVar1++
			}()
			return *tempVar1
		}()] = c
	}
}

// src_lineget - transpiled function from  src.c:98
func src_lineget() int32 {
	return lineno
}

// src_lineset - transpiled function from  src.c:103
func src_lineset(n int32) {
	lineno = n
}

// macro - transpiled function from  src.c:109
// eqn macros
type macro struct {
	name [32]byte
	def  []byte
}

// macros - transpiled function from  src.c:113
var macros []macro = make([]macro, 512)

// nmacros - transpiled function from  src.c:114
var nmacros int32

// src_findmacro - transpiled function from  src.c:116
func src_findmacro(name []byte) int32 {
	var i int32
	for i = 0; i < nmacros; i++ {
		if noarch.Not(noarch.Strcmp(macros[i].name[:], name)) {
			return i
		}
	}
	return -1
}

// src_macro - transpiled function from  src.c:126
func src_macro(name []byte) int32 {
	// return nonzero if name is a macro
	return noarch.BoolToInt(src_findmacro(name) >= 0)
}

// src_define - transpiled function from  src.c:132
func src_define(name []byte, def []byte) {
	// define a macro
	var idx int32 = src_findmacro(name)
	if idx < 0 && nmacros < 512 {
		idx = func() int32 {
			defer func() {
				nmacros++
			}()
			return nmacros
		}()
	}
	if idx >= 0 {
		noarch.Strcpy(macros[idx].name[:], name)
		_ = macros[idx].def
		macros[idx].def = src_strdup(def)
	}
}

// src_done - transpiled function from  src.c:144
func src_done() {
	var i int32
	for i = 0; i < nmacros; i++ {
		_ = macros[i].def
	}
}

// src_expand - transpiled function from  src.c:152
func src_expand(name []byte, args [][]byte) int32 {
	// expand macro
	var i int32 = src_findmacro(name)
	if i >= 0 {
		src_push(macros[i].def, args)
	}
	return noarch.BoolToInt(i < 0)
}

// src_arg - transpiled function from  src.c:161
func src_arg(i int32) int32 {
	// expand argument
	var call int32 = esrc_c4go_postfix[0].call
	if call != 0 && esrc_c4go_postfix[0].args[:][i-1] != nil {
		src_push(esrc_c4go_postfix[0].args[:][i-1], nil)
	}
	if call != 0 {
		return 0
	}
	return 1
}

// src_top - transpiled function from  src.c:170
func src_top() int32 {
	// return one if not reading macros and their arguments
	return noarch.BoolToInt(esrc_c4go_postfix[0].prev == nil)
}

// kwds - transpiled function from  tok.c:13
// the preprocessor and tokenizer
var kwds [][]byte = [][]byte{[]byte("fwd\x00"), []byte("down\x00"), []byte("back\x00"), []byte("up\x00"), []byte("bold\x00"), []byte("italic\x00"), []byte("roman\x00"), []byte("font\x00"), []byte("fat\x00"), []byte("size\x00"), []byte("bar\x00"), []byte("dot\x00"), []byte("dotdot\x00"), []byte("dyad\x00"), []byte("hat\x00"), []byte("under\x00"), []byte("vec\x00"), []byte("tilde\x00"), []byte("sub\x00"), []byte("sup\x00"), []byte("from\x00"), []byte("to\x00"), []byte("vcenter\x00"), []byte("left\x00"), []byte("right\x00"), []byte("over\x00"), []byte("sqrt\x00"), []byte("pile\x00"), []byte("lpile\x00"), []byte("cpile\x00"), []byte("rpile\x00"), []byte("above\x00"), []byte("matrix\x00"), []byte("col\x00"), []byte("ccol\x00"), []byte("lcol\x00"), []byte("rcol\x00"), []byte("delim\x00"), []byte("define\x00"), []byte("gfont\x00"), []byte("grfont\x00"), []byte("gbfont\x00"), []byte("gsize\x00"), []byte("set\x00"), []byte("chartype\x00"), []byte("mark\x00"), []byte("lineup\x00"), []byte("bracketsizes\x00"), []byte("bracketpieces\x00"), []byte("breakcost\x00")}

// tok_eqen - transpiled function from  tok.c:26
// non-zero if inside .EQ/.EN
var tok_eqen int32

// tok_line - transpiled function from  tok.c:27
// inside inline eqn block
var tok_line int32

// tok_part - transpiled function from  tok.c:28
// partial line with inline eqn blocks
var tok_part int32

// tok - transpiled function from  tok.c:29
// current token
var tok []byte = make([]byte, 1000)

// tok_prev - transpiled function from  tok.c:30
// previous token
var tok_prev []byte = make([]byte, 1000)

// tok_curtype - transpiled function from  tok.c:31
// type of current token
var tok_curtype int32

// tok_cursep - transpiled function from  tok.c:32
// current character is a separator
var tok_cursep int32

// tok_prevsep - transpiled function from  tok.c:33
// previous character was a separator
var tok_prevsep int32

// eqn_beg - transpiled function from  tok.c:34
// inline eqn delimiters
var eqn_beg int32

// eqn_end - transpiled function from  tok.c:34
var eqn_end int32

// tok_req - transpiled function from  tok.c:37
func tok_req(a int32, b int32) int32 {
	// return zero if troff request .ab is read
	var eqln []int32 = make([]int32, 1000)
	var i int32
	var ret int32
	eqln[func() int32 {
		defer func() {
			i++
		}()
		return i
	}()] = src_next()
	if eqln[i-1] != int32('.') {
		goto failed
	}
	eqln[func() int32 {
		defer func() {
			i++
		}()
		return i
	}()] = src_next()
	for eqln[i-1] == int32(' ') && uint32(i) < 4000-4 {
		eqln[func() int32 {
			defer func() {
				i++
			}()
			return i
		}()] = src_next()
	}
	if eqln[i-1] != a {
		goto failed
	}
	eqln[func() int32 {
		defer func() {
			i++
		}()
		return i
	}()] = src_next()
	if eqln[i-1] != b {
		goto failed
	}
	ret = 1
failed:
	;
	for i > 0 {
		src_back(eqln[func() int32 {
			i--
			return i
		}()])
	}
	return ret
}

// tok_en - transpiled function from  tok.c:61
func tok_en() int32 {
	// read .EN
	return tok_req(int32('E'), int32('N'))
}

// tok_eq - transpiled function from  tok.c:67
func tok_eq(s []byte) int32 {
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('.') {
		// does the line start with eq
		return 0
	}
	for int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
		s = s[0+1:]
	}
	return noarch.BoolToInt(int32(s[0]) == int32('E') && int32(s[1]) == int32('Q'))
}

// tok_lf - transpiled function from  tok.c:77
func tok_lf(s []byte) int32 {
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('.') {
		// read an lf request
		return 0
	}
	for int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
		s = s[0+1:]
	}
	if int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('l') || int32((func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0]) != int32('f') {
		return 0
	}
	for int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISspace)) != 0 {
		s = s[0+1:]
	}
	if int32(((__ctype_b_loc())[0])[int32(uint8(s[0]))])&int32(uint16(noarch.ISdigit)) != 0 {
		src_lineset(noarch.Atoi(s))
	}
	return 1
}

// tok_next - transpiled function from  tok.c:93
func tok_next() int32 {
	// read the next input character
	var c int32
	if noarch.Not(tok_eqen) && noarch.Not(tok_line) {
		return 0
	}
	c = src_next()
	if tok_eqen != 0 && c == int32('\n') && tok_en() != 0 {
		tok_eqen = 0
	}
	if tok_line != 0 && (src_top() != 0 && c == eqn_end) {
		tok_line = 0
		return 0
	}
	return c
}

// tok_back - transpiled function from  tok.c:109
func tok_back(c int32) {
	if tok_eqen != 0 || tok_line != 0 {
		// push back the last character read
		src_back(c)
	}
}

// tok_preview - transpiled function from  tok.c:116
func tok_preview(s []byte) {
	// read the next word
	var c int32 = src_next()
	var n int32
	if c > 0 && def_chopped(c) != 0 {
		s[func() int32 {
			defer func() {
				n++
			}()
			return n
		}()] = byte(c)
		s[n] = '\x00'
		return
	}
	for c > 0 && noarch.Not(def_chopped(c)) && (noarch.Not(tok_line) || (noarch.Not(src_top()) || c != eqn_end)) {
		s[func() int32 {
			defer func() {
				n++
			}()
			return n
		}()] = byte(c)
		c = src_next()
	}
	s[n] = '\x00'
	src_back(c)
}

// tok_unpreview - transpiled function from  tok.c:134
func tok_unpreview(s []byte) {
	// push back the given word
	var n int32 = noarch.Strlen(s)
	for n > 0 {
		src_back(int32(uint8(s[func() int32 {
			n--
			return n
		}()])))
	}
}

// tok_keyword - transpiled function from  tok.c:142
func tok_keyword() int32 {
	// read a keyword; return zero on success
	var i int32
	tok_preview(tok)
	for i = 0; uint32(i) < 400/8; i++ {
		if noarch.Not(noarch.Strcmp(kwds[i], tok)) {
			return 0
		}
	}
	tok_unpreview(tok)
	return 1
}

// tok_readarg - transpiled function from  tok.c:154
func tok_readarg(sbuf_c4go_postfix []sbuf) int32 {
	// read the next argument of a macro call; return zero if read a ','
	var c int32 = src_next()
	// number of nested parenthesis
	var pdepth int32
	// inside double quotes
	var quotes int32
	for c > 0 && (pdepth != 0 || quotes != 0 || c != int32(',') && c != int32(')')) {
		sbuf_add(sbuf_c4go_postfix, c)
		if noarch.Not(quotes) && c == int32(')') {
			pdepth++
		}
		if noarch.Not(quotes) && c == int32('(') {
			pdepth--
		}
		if c == int32('"') {
			quotes = 1 - quotes
		}
		if c == int32('\\') {
			sbuf_add(sbuf_c4go_postfix, func() int32 {
				c = src_next()
				tempVar3 := &c
				return *tempVar3
			}())
			if c == int32('*') || c == int32('n') {
				sbuf_add(sbuf_c4go_postfix, func() int32 {
					c = src_next()
					tempVar3 := &c
					return *tempVar3
				}())
			}
			if c == int32('(') {
				sbuf_add(sbuf_c4go_postfix, func() int32 {
					c = src_next()
					tempVar3 := &c
					return *tempVar3
				}())
				sbuf_add(sbuf_c4go_postfix, func() int32 {
					c = src_next()
					tempVar3 := &c
					return *tempVar3
				}())
			} else if c == int32('[') {
				for c > 0 && c != int32(']') {
					sbuf_add(sbuf_c4go_postfix, func() int32 {
						c = src_next()
						tempVar3 := &c
						return *tempVar3
					}())
				}
			}
		}
		c = src_next()
	}
	if c == int32(',') {
		return 0
	}
	return 1
}

// tok_expand - transpiled function from  tok.c:185
func tok_expand() int32 {
	// expand a macro; return zero on success
	var args [][]byte = [][]byte{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	var sbufs []sbuf = make([]sbuf, 10)
	var i int32
	var n int32
	tok_preview(tok)
	if src_macro(tok) != 0 {
		var c int32 = src_next()
		src_back(c)
		if c == int32('(') {
			// macro arguments follow
			src_next()
			for n <= 9 {
				sbuf_init(sbufs[n:])
				if tok_readarg(sbufs[func() int32 {
					defer func() {
						n++
					}()
					return n
				}():]) != 0 {
					break
				}
			}
		}
		for i = 0; i < n; i++ {
			args[i] = sbuf_buf(sbufs[i:])
		}
		src_expand(tok, args)
		for i = 0; i < n; i++ {
			sbuf_done(sbufs[i:])
		}
		return 0
	}
	tok_unpreview(tok)
	return 1
}

// tok_eqn - transpiled function from  tok.c:214
func tok_eqn() int32 {
	// read until .EQ or eqn_beg
	var ln sbuf
	var c int32
	tok_cursep = 1
	sbuf_init(c4goUnsafeConvert_sbuf(&ln))
	for (func() int32 {
		c = src_next()
		return c
	}()) > 0 {
		if c == eqn_beg {
			fmt.Printf(".eo\n")
			noarch.Printf([]byte(".%s %s \"%s\n\x00"), func() []byte {
				if tok_part != 0 {
					return []byte("as\x00")
				}
				return []byte("ds\x00")
			}(), []byte("10\x00"), sbuf_buf(c4goUnsafeConvert_sbuf(&ln)))
			sbuf_done(c4goUnsafeConvert_sbuf(&ln))
			fmt.Printf(".ec\n")
			tok_part = 1
			tok_line = 1
			return 0
		}
		sbuf_add(c4goUnsafeConvert_sbuf(&ln), c)
		if c == int32('\n') && noarch.Not(tok_part) {
			noarch.Printf([]byte("%s\x00"), sbuf_buf(c4goUnsafeConvert_sbuf(&ln)))
			tok_lf(sbuf_buf(c4goUnsafeConvert_sbuf(&ln)))
			if tok_eq(sbuf_buf(c4goUnsafeConvert_sbuf(&ln))) != 0 && noarch.Not(tok_en()) {
				tok_eqen = 1
				sbuf_done(c4goUnsafeConvert_sbuf(&ln))
				return 0
			}
		}
		if c == int32('\n') && tok_part != 0 {
			noarch.Printf([]byte(".lf %d\n\x00"), src_lineget())
			noarch.Printf([]byte("\\*%s%s\x00"), escarg([]byte("10\x00")), sbuf_buf(c4goUnsafeConvert_sbuf(&ln)))
			tok_part = 0
		}
		if c == int32('\n') {
			sbuf_cut(c4goUnsafeConvert_sbuf(&ln), 0)
		}
	}
	sbuf_done(c4goUnsafeConvert_sbuf(&ln))
	return 1
}

// tok_eqnout - transpiled function from  tok.c:254
func tok_eqnout(s []byte) {
	if noarch.Not(tok_part) {
		// collect the output of this eqn block
		noarch.Printf([]byte(".ds %s \"%s%s%s\n\x00"), []byte("10\x00"), []byte("\\E*[.eqnbeg]\\R'.eqnfn0 \\En(.f'\\R'.eqnsz0 \\En(.s'\x00"), s, []byte("\\f[\\En[.eqnfn0]]\\s[\\En[.eqnsz0]]\\E*[.eqnend]\x00"))
		noarch.Printf([]byte(".lf %d\n\x00"), src_lineget()-1)
		noarch.Printf([]byte("\\&\\*%s\n\x00"), escarg([]byte("10\x00")))
	} else {
		noarch.Printf([]byte(".as %s \"%s%s%s\n\x00"), []byte("10\x00"), []byte("\\E*[.eqnbeg]\\R'.eqnfn0 \\En(.f'\\R'.eqnsz0 \\En(.s'\x00"), s, []byte("\\f[\\En[.eqnfn0]]\\s[\\En[.eqnsz0]]\\E*[.eqnend]\x00"))
	}
}

// utf8len - transpiled function from  tok.c:266
func utf8len(c int32) int32 {
	if ^c&128 != 0 {
		// return the length of a utf-8 character based on its first byte
		return noarch.BoolToInt(c > 0)
	}
	if ^c&64 != 0 {
		return 1
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

// char_type - transpiled function from  tok.c:282
func char_type(s []byte) int32 {
	// return the type of a token
	var c int32 = int32(uint8(s[0]))
	var t int32
	if int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISdigit)) != 0 {
		return 18
	}
	if c == int32('"') {
		return 19
	}
	if (func() int32 {
		t = def_type(s)
		return t
	}()) >= 0 {
		return t
	}
	if c == int32('~') || c == int32('^') {
		return 20
	}
	if int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISpunct)) != 0 && (c != int32('\\') || noarch.Not(s[1])) {
		return 16
	}
	return 17
}

// tok_read - transpiled function from  tok.c:300
func tok_read() int32 {
	// read the next token
	var s []byte = tok
	var e []byte = c4goPointerArithByteSlice(tok[0+1000:], int(-2))
	var c int32
	var c2 int32
	var i int32
	s[0] = '\x00'
	c = tok_next()
	if c <= 0 {
		return 1
	}
	tok_prevsep = tok_cursep
	tok_cursep = def_chopped(c)
	if tok_cursep != 0 {
		tok_prevsep = 1
	}
	if c == int32(' ') || c == int32('\n') {
		for c > 0 && (c == int32(' ') || c == int32('\n')) {
			c = tok_next()
		}
		tok_back(c)
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = ' '
		s[0] = '\x00'
		tok_curtype = 1
		return 0
	}
	if c == int32('\t') {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = '\t'
		s[0] = '\x00'
		tok_curtype = 2
		return 0
	}
	if tok_prevsep != 0 {
		if c == int32('$') {
			c2 = tok_next()
			if c2 >= int32('1') && c2 <= int32('9') && noarch.Not(src_arg(c2-int32('0'))) {
				tok_cursep = 1
				return tok_read()
			}
			tok_back(c2)
		}
		tok_back(c)
		if noarch.Not(tok_keyword()) {
			tok_curtype = 3
			tok_cursep = 1
			return 0
		}
		if noarch.Not(tok_expand()) {
			tok_cursep = 1
			return tok_read()
		}
		c = tok_next()
	}
	if noarch.Strchr(([]byte("^~{}(),\"\n\t =:|.+-*/\\,()[]<>!\x00")), c) != nil {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(c)
		if c == int32('\\') {
			c = tok_next()
			if c == int32('(') {
				(func() []byte {
					defer func() {
						s = s[0+1:]
					}()
					return s
				}())[0] = byte(c)
				(func() []byte {
					defer func() {
						s = s[0+1:]
					}()
					return s
				}())[0] = byte(tok_next())
				(func() []byte {
					defer func() {
						s = s[0+1:]
					}()
					return s
				}())[0] = byte(tok_next())
			} else if c == int32('[') {
				for c != 0 && c != int32(']') {
					if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&e[0])))/int64(1)) < 0 {
						(func() []byte {
							defer func() {
								s = s[0+1:]
							}()
							return s
						}())[0] = byte(c)
					}
					c = tok_next()
				}
				(func() []byte {
					defer func() {
						s = s[0+1:]
					}()
					return s
				}())[0] = ']'
			}
		} else if c == int32('"') {
			c = tok_next()
			for c > 0 && c != int32('"') {
				if c == int32('\\') {
					c2 = tok_next()
					if c2 == int32('"') {
						c = int32('"')
					} else {
						tok_back(c2)
					}
				}
				if (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&e[0])))/int64(1)) < 0 {
					(func() []byte {
						defer func() {
							s = s[0+1:]
						}()
						return s
					}())[0] = byte(c)
				}
				c = tok_next()
			}
			(func() []byte {
				defer func() {
					s = s[0+1:]
				}()
				return s
			}())[0] = '"'
		} else {
			// two-character operators
			c2 = tok_next()
			switch c<<uint64(8) | c2 {
			case (int32('<'<<uint64(8) | '=')):
				fallthrough
			case (int32('>'<<uint64(8) | '=')):
				fallthrough
			case (int32('='<<uint64(8) | '=')):
				fallthrough
			case (int32('!'<<uint64(8) | '=')):
				fallthrough
			case (int32('>'<<uint64(8) | '>')):
				fallthrough
			case (int32('<'<<uint64(8) | '<')):
				fallthrough
			case (int32(':'<<uint64(8) | '=')):
				fallthrough
			case (int32('-'<<uint64(8) | '>')):
				fallthrough
			case (int32('<'<<uint64(8) | '-')):
				fallthrough
			case (int32('-'<<uint64(8) | '+')):
				(func() []byte {
					defer func() {
						s = s[0+1:]
					}()
					return s
				}())[0] = byte(c2)
			default:
				tok_back(c2)
			}
		}
		s[0] = '\x00'
		tok_curtype = char_type(tok)
		return 0
	}
	(func() []byte {
		defer func() {
			s = s[0+1:]
		}()
		return s
	}())[0] = byte(c)
	i = utf8len(c)
	for func() int32 {
		i--
		return i
	}() > 0 && (int64(uintptr(unsafe.Pointer(&s[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&e[0])))/int64(1)) < 0 {
		(func() []byte {
			defer func() {
				s = s[0+1:]
			}()
			return s
		}())[0] = byte(tok_next())
	}
	s[0] = '\x00'
	tok_curtype = char_type(tok)
	return 0
}

// tok_get - transpiled function from  tok.c:415
func tok_get() []byte {
	// current token
	if int32(tok[0]) != 0 {
		return tok
	}
	return nil
}

// tok_type - transpiled function from  tok.c:421
func tok_type() int32 {
	// current token type
	if int32(tok[0]) != 0 {
		return tok_curtype
	}
	return 0
}

// tok_chops - transpiled function from  tok.c:427
func tok_chops(soft int32) int32 {
	if tok_get() == nil || tok_curtype == 3 {
		// return nonzero if current token chops the equation
		return 1
	}
	if soft != 0 {
		return noarch.BoolToInt(len(noarch.Strchr(([]byte("^~{}(),\"\n\t =:|.+-*/\\,()[]<>!\x00")), int32(uint8(tok_get()[0])))) != 0)
	}
	return def_chopped(int32(uint8(tok_get()[0])))
}

// tok_pop - transpiled function from  tok.c:437
func tok_pop() []byte {
	// read the next token, return the previous
	noarch.Strcpy(tok_prev, tok)
	tok_read()
	if int32(tok_prev[0]) != 0 {
		return tok_prev
	}
	return nil
}

// tok_poptext - transpiled function from  tok.c:445
func tok_poptext(sep int32) []byte {
	for tok_type() == 1 {
		// like tok_pop() but ignore T_SPACE tokens; if sep, read until chopped
		tok_read()
	}
	tok_prev[0] = '\x00'
	for {
		noarch.Strcat(tok_prev, tok)
		tok_read()
		if !(int32(tok[0]) != 0 && noarch.Not(tok_chops(noarch.BoolToInt(noarch.Not(sep))))) {
			break
		}
	}
	if int32(tok_prev[0]) != 0 {
		return tok_prev
	}
	return nil
}

// tok_blanks - transpiled function from  tok.c:458
func tok_blanks() {
	for tok_type() == 1 {
		// skip spaces
		tok_pop()
	}
}

// tok_jmp - transpiled function from  tok.c:465
func tok_jmp(s []byte) int32 {
	// if the next token is s, return zero and skip it
	tok_blanks()
	if tok_get() != nil && noarch.Not(s[1]) && noarch.Strchr([]byte("{}~^\t\x00"), int32(s[0])) != nil && noarch.Not(noarch.Strcmp(s, tok_get())) {
		tok_pop()
		return 0
	}
	if tok_type() != 3 || tok_get() == nil || noarch.Strcmp(s, tok_get()) != 0 {
		return 1
	}
	tok_pop()
	return 0
}

// tok_delim - transpiled function from  tok.c:479
func tok_delim() {
	// read delim command
	var delim []byte = make([]byte, 32)
	tok_preview(delim)
	if noarch.Not(noarch.Strcmp([]byte("off\x00"), delim)) {
		eqn_beg = 0
		eqn_end = 0
	} else {
		eqn_beg = int32(delim[0])
		eqn_end = int32(delim[1])
	}
}

// tok_macrodef - transpiled function from  tok.c:493
func tok_macrodef(def []sbuf) {
	// read macro definition
	var c int32
	var delim int32
	c = src_next()
	for c > 0 && int32(((__ctype_b_loc())[0])[c])&int32(uint16(noarch.ISspace)) != 0 {
		c = src_next()
	}
	delim = c
	c = src_next()
	for c > 0 && c != delim {
		sbuf_add(def, c)
		c = src_next()
	}
}

// tok_macro - transpiled function from  tok.c:509
func tok_macro() {
	// read the next macro command
	var name []byte = make([]byte, 32)
	var def sbuf
	tok_preview(name)
	sbuf_init(c4goUnsafeConvert_sbuf(&def))
	tok_macrodef(c4goUnsafeConvert_sbuf(&def))
	src_define(name, sbuf_buf(c4goUnsafeConvert_sbuf(&def)))
	sbuf_done(c4goUnsafeConvert_sbuf(&def))
}

// tok_inline - transpiled function from  tok.c:521
func tok_inline() int32 {
	// return 1 if inside inline equations
	return tok_line
}

// c4goUnsafeConvert_esrc : created by c4go
func c4goUnsafeConvert_esrc(c4go_name *esrc) []esrc {
	return (*[10000]esrc)(unsafe.Pointer(c4go_name))[:]
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[10000]int32)(unsafe.Pointer(c4go_name))[:]
}

// c4goUnsafeConvert_sbuf : created by c4go
func c4goUnsafeConvert_sbuf(c4go_name *sbuf) []sbuf {
	return (*[10000]sbuf)(unsafe.Pointer(c4go_name))[:]
}

// the contents
// number register holding box size
// register holding the contents
// the number of atoms inserted
// type of the first and the last atoms
// tex style (TS_*)
// register for saving box width
// managing registers
// eqn global variables

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
