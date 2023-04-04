FROM alpine:latest

WORKDIR /

COPY output/tnote.linux-amd64 /app/tnote

CMD [ "/app/tnote" ]
