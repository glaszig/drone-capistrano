# Docker image for the Drone Capistrano plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-capistrano
#     make deps build docker

FROM ruby:2.6-rc-alpine

ENV BUILD_PACKAGES="curl-dev build-base ca-certificates git openssh" \
    DEV_PACKAGES="zlib-dev libxml2-dev libxslt-dev tzdata yaml-dev sqlite-dev postgresql-dev mariadb-dev"

RUN \
  apk add --update --upgrade $BUILD_PACKAGES $DEV_PACKAGES && \
  gem install bundler --version "$BUNDLER_VERSION" && \
  find / -type f -iname \*.apk-new -delete && \
  rm -rf /var/cache/apk/* && \
  rm -rf /usr/lib/ruby/gems/*/cache/* && \
  rm -rf /usr/local/bundle/cache/* && \
  rm -rf ~/.gem

ADD bundle.sh /
ADD drone-capistrano /bin/
ENTRYPOINT ["/bin/drone-capistrano"]
