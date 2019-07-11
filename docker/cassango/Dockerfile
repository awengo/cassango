FROM golang:1.11

RUN apt-get update && apt-get install -y zlib1g-dev

ENV GOPATH /opt/go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/Masterminds/glide;
#RUN go get github.com/codegangsta/gin;

#WORKDIR /opt/go/src/ezapi/basic
#COPY ./apps/basic/glide.yaml glide.yaml
#COPY ./apps/basic/glide.lock glide.lock
#RUN glide install


#WORKDIR /opt/go/src/ezapi/member
#COPY ./apps/member/glide.yaml glide.yaml
#COPY ./apps/member/glide.lock glide.lock
#RUN glide install
