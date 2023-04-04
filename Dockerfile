FROM alpine:latest
ARG TARGETPLATFORM
WORKDIR /

RUN echo "Building for $TARGETPLATFORM"

COPY output/ /app/output/

RUN cp /app/output/tnote.${TARGETPLATFORM/\//-} /app/tnote && rm -rf /app/output

CMD [ "/app/tnote" ]
