FROM alpine:3.18

WORKDIR /game-store/gokiburi

RUN apk add wget
RUN apk add tar
RUN apk add go
RUN wget -q https://github.com/michenriksen/gokiburi/releases/download/v0.1.0/gokiburi_Linux_x86_64.tar.gz -O /tmp/gokiburi.tar.gz
RUN tar -xf /tmp/gokiburi.tar.gz -C /bin/

CMD [ "/bin/gokiburi", "--listen-address", "::", "--skip-paths", "web" ]
