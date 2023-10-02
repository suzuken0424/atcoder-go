FROM golang:1.20.0

RUN apt update \
  && apt install -y vim \
  python3 \
  python3-pip \
  nodejs \
  npm \
  && pip3 install -U pip \
  && pip3 install online-judge-tools \
  && npm install -g atcoder-cli \
  && acc config default-task-choice all \
  && echo 'alias oj-test="oj t -c "go run ./main.go" -d ./tests/"' >> ~/.bashrc \
  && echo 'alias add-temp="cp /go/src/template.go ./main.go"' >> ~/.bashrc


WORKDIR /go/src

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/nsf/gocode@latest 
RUN go install github.com/acroca/go-symbols@latest 
RUN go install  github.com/fatih/gomodifytags@latest 
RUN go install  github.com/josharian/impl@latest 
RUN go install  github.com/haya14busa/goplay/cmd/goplay@latest 
RUN go install  github.com/go-delve/delve/cmd/dlv@latest 
RUN go install  golang.org/x/lint/golint@latest 
RUN go install  golang.org/x/tools/gopls@latest