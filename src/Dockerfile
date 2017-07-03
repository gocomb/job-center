FROM 10.110.18.107:5000/nodecontrol-base
MAINTAINER liyaoming
ENV GOPATH /app/get-resource
ADD . /app/get-resource/src
WORKDIR /app
EXPOSE 8080
RUN go build -i -o "/app/run" ./get-resource/src/cmd/get_resource.go
RUN chmod 777 /app/run
RUN /app/run
