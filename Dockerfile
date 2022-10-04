FROM golang:1.16
MAINTAINER King
RUN mkdir -p /acm_recruit
WORKDIR /acm_recruit
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN git clone https://github.com/dongzhiwei-git/acm_recruit.git
WORKDIR /acm_recruit/acm_recruit
RUN git checkout v1.0.0
RUN go build -o acm_recruit
EXPOSE 8000
CMD ["./acm_recruit"]

