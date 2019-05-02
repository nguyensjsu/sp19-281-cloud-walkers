FROM golang:latest
EXPOSE 3001
RUN mkdir /Database
ADD . /Database/
WORKDIR /Database
ENV GOPATH /Database
RUN cd /Database ; go install Database
CMD ["/Database/bin/LoginSystem"]
