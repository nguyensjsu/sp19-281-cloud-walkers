FROM golang:1.12
RUN mkdir /cloudwalker
ADD . /cloudwalker/
WORKDIR /cloudwalker
ENV GOPATH /cloudwalker
RUN cd /cloudwalker; go install cloudwalker
CMD ["/cloudwalker/bin/cloudwalker"]                         
