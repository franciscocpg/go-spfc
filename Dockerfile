FROM milcom/centos7-systemd
# Install go
RUN /bin/bash -c 'bash <(curl -s https://gist.githubusercontent.com/franciscocpg/1ce8e61a9f915b95e12f/raw/)'

RUN /bin/bash -c 'yum install wget git -y'
# Install glide
RUN /bin/bash -c 'bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)'

RUN /bin/bash -c 'mkdir -p /root/go/src/github.com/franciscocpg/go-spfc'

ENV GOPATH /root/go

ADD . $GOPATH/src/github.com/franciscocpg/go-spfc