NAME=pm-poppler
DEPEND=github.com/Masterminds/glide github.com/cheggaaa/go-poppler

.PHONY: depend clean build

clean-build: clean depend build

build: depend
	go build

run : depend
	go run *.go

depend:
	go get -v $(DEPEND)
	glide install

clean:
	rm -fr pm-poppler vendor/*
	touch vendor/.gitkeep