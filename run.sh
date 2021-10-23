#/bin/bash


# NEATPOST
rm -rf tmp
mkdir tmp
cp ./neatroff_make/neatpost/* tmp/
sed -i.bak '19,25s/^.*//g' ./tmp/dev.c
c4go transpile -o post.go\
	-clang-flag="-DTROFFFDIR=\"./\""\
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
	-clang-flag="-DTROFFFDIR=\"./\""\
	-clang-flag="-DTROFFMDIR=\"./\""\
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
 	-clang-flag="-DTROFFFDIR=\"./\""\
 	-clang-flag="-DTROFFMDIR=\"./\""\
	./tmp/*.c\
	./tmp/*.h
rm -rf tmp
