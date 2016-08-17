CURDIR=$(shell pwd)
BINDIR=$(CURDIR)/bin
GOPATH=$(CURDIR):$(CURDIR)/vendor

all: daemon razerctl c-lib python-lib

daemon:
	mkdir -p $(BINDIR)
	go install daemon

razerctl:
	mkdir -p $(BINDIR)
	go install razerctl

c-lib:
	mkdir -p $(BINDIR)/lib/c
	go build -o $(BINDIR)/lib/c/openrazer.so -buildmode=c-shared lib/c

python-lib:
	mkdir -p $(BINDIR)/lib/python
	cp $(CURDIR)/src/lib/python/lib.py $(BINDIR)/lib/python/openrazer.py

samples: c-sample python-sample

c-sample: c-lib
	mkdir -p $(BINDIR)/samples
	gcc -I$(CURDIR)/bin/lib/c -Wall -o $(BINDIR)/samples/sample-c samples/c/main.c $(BINDIR)/lib/c/openrazer.so

python-sample:  python-lib
	mkdir -p $(BINDIR)/samples
	cp $(CURDIR)/samples/python/main.py $(BINDIR)/samples/sample-python.py
	chmod +x $(BINDIR)/samples/sample-python.py

run-python-sample:
	 PYTHONPATH="$(BINDIR)/lib/python" LD_LIBRARY_PATH="$(BINDIR)/lib/c/" $(BINDIR)/samples/sample-python.py

clean:
	@rm -r $(BINDIR)
