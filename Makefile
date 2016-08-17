CURDIR=$(shell pwd)
BINDIR=$(CURDIR)/bin
GOPATH=$(CURDIR):$(CURDIR)/vendor

all: daemon razerctl c-lib python-lib

daemon:
	go install daemon

razerctl:
	go install razerctl

c-lib:
	mkdir -p $(BINDIR)/lib/c
	go build -o $(BINDIR)/lib/c/openrazer.so -buildmode=c-shared lib/c

python-lib:
	mkdir -p $(BINDIR)/lib/python
	cp $(CURDIR)/src/lib/python/lib.py $(BINDIR)/lib/python/openrazer.py

samples: c-sample

c-sample: c-lib
	mkdir -p $(BINDIR)/samples
	gcc -I$(CURDIR)/bin/lib/c -Wall -o $(BINDIR)/samples/sample-c samples/c/main.c $(BINDIR)/lib/c/openrazer.so

clean:
	@rm $(BINDIR)/daemon
	@rm $(BINDIR)/razerctl
	@rm $(BINDIR)/openrazer.so
	@rm $(BINDIR)/openrazer.h
	@rm -r $(BINDIR)/lib
	@rm -r $(BINDIR)/samples
