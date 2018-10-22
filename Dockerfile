FROM plugins/base:amd64

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Apex Up" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates curl && \
  curl -sf https://up.apex.sh/install | sh && \
  rm -rf /var/cache/apk/*

ADD release/linux/amd64/drone-apex-up /bin/
ENTRYPOINT ["/bin/drone-apex-up"]
