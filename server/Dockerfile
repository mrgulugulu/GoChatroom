FROM golang

RUN mkdir -p /app

WORKDIR /app
 
ADD main /app/main

EXPOSE 8080

RUN echo "10.101.110.38    kafka-0.kafka-headless.kafka.svc.cluster.local">>/etc/hosts
 
CMD ["./main"]
