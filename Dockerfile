FROM golang:1.7.1

RUN apt update \
  && apt install -y libcairo2-dev libpoppler-glib-dev \
  && go get -u github.com/cheggaaa/go-poppler

RUN mkdir -p $GOPATH/src/github.com/kaakaa/pm-poppler

ENTRYPOINT ["tail", "-f", "/dev/null"]
