#/bin/bash


# NEATPOST
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neatpost/* tmp/
sed -i.bak '19,25s/^.*//g' ./tmp/dev.c
c4go transpile -o post.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/post.c\
	./tmp/pdf.c\
	./tmp/pdfext.c\
	./tmp/font.c\
	./tmp/dev.c\
	./tmp/clr.c\
	./tmp/dict.c\
	./tmp/iset.c\
	./tmp/sbuf.c\
	./tmp/post.h
rm -rf tmp

# NEATROFF
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neatroff/* tmp/
sed -i.bak '574,580s/^.*//g' ./tmp/font.c
sed -i.bak '66,81s/^.*//g' ./tmp/out.c
c4go transpile -o troff.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# NEATREFER
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neatrefer/* tmp/
c4go transpile -o refer.go\
	./tmp/*.c
rm -rf tmp


# NEATEQN
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neateqn/* tmp/
sed -i.bak '216,233s/^.*//g' ./tmp/eqn.c
c4go transpile -o eqn.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# NEATMKFN
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neatmkfn/* tmp/
sed -i.bak '54,58s/^.*//g' ./tmp/otf.c
c4go transpile -o mkfn.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# SHAPE
rm -rf tmp
mkdir tmp
cp ./neatroff_make/shape/* tmp/
c4go transpile -o shape.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# SOIL
rm -rf tmp
mkdir tmp
cp ./neatroff_make/soin/* tmp/
c4go transpile -o soil.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c
rm -rf tmp

# TBL
rm -rf tmp
mkdir tmp
cp ./neatroff_make/troff/tbl/* tmp/
mv ./tmp/t.h ./tmp/t.h.bak
echo "#ifndef TH"  >> ./tmp/t.h
echo "#define TH"  >> ./tmp/t.h
cat  ./tmp/t.h.bak >> ./tmp/t.h
echo "#endif"      >> ./tmp/t.h
c4go transpile -o tbl.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	-clang-flag="-Wall"\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# PIC
rm -rf tmp
mkdir tmp
cp ./neatroff_make/troff/pic/* tmp/
sed -i.bak '485s/void /int /g' ./tmp/input.c
c4go transpile -o pic.go\
 	-clang-flag="-DTROFFFDIR=\"./\neatroff_make\/fonts\""\
 	-clang-flag="-DTROFFMDIR=\".\/neatroff_make\/tmac\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp

# CLEAN folder names
sed -i.bak 's/GOPATH\/src\/github.com\/Konstantin8105\/uroff\/tmp\///g' *.go
sed -i.bak 's/1000000/10000/g' *.go
rm *.bak

# SHOW AMOUNT LINES
for entry in `ls *.go`
do
  wc -l "$entry"
done

