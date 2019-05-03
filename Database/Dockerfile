FROM golang:latest
EXPOSE 3001
RUN mkdir /userprofile
ADD . /userprofile/
WORKDIR /userprofile
ENV GOPATH /userprofile
RUN cd /userprofile ; go install userprofile
CMD ["/userprofile/bin/userprofile"]
