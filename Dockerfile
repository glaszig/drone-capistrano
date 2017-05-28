# Docker image for the Drone Capistrano plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-capistrano
#     make deps build docker

FROM ruby:2.3-alpine

RUN apk add --update build-base ca-certificates git openssh \
  && gem install bundler \
  && rm -rf /var/cache/apk/*

ENV BUNDLE_APP_CONFIG .bundle

# ADD bundle.sh /
ADD drone-capistrano /bin/
ENTRYPOINT ["/bin/drone-capistrano"]
