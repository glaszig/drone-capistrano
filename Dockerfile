# Docker image for the Drone Capistrano plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-capistrano
#     make deps build
#     docker build --rm=true -t plugins/drone-capistrano .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    ruby && \
  gem install --no-ri --no-rdoc capistrano && \
  gem install --no-ri --no-rdoc io-console && \
  rm -rf /var/cache/apk/*

ADD drone-capistrano /bin/
ENTRYPOINT ["/bin/drone-capistrano"]
