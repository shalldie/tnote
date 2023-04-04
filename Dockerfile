FROM alpine:latest
ARG TARGETPLATFORM
WORKDIR /

RUN echo "Building for $TARGETPLATFORM"

COPY output/tnote.linux-amd64 /app/tnote

CMD [ "/app/tnote" ]
