FROM golang:alpine 
LABEL maintainer="<chenhe@zuoyejia.com>"


WORKDIR /opt/sse
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/sse
COPY main /opt/sse/bin/

EXPOSE 8080

ENTRYPOINT ["/opt/sse/bin/main"]