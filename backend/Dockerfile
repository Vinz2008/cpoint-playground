FROM golang:1.21

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /prod
COPY ./* /prod/


RUN go build . 

EXPOSE $PORT

ENTRYPOINT ["./cpoint-playground-backend"]