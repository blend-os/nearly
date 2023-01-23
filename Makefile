all: nearly

PREFIX?=/usr/local
DESTDIR?=

nearly:
	go build -o nearly main.go

install:
	cp nearly ${DESTDIR}/${PREFIX}/bin