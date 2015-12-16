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
    openssh \
    ruby \
    ruby-dev && \
  gem install --no-ri --no-rdoc bundler && \
  gem install --no-ri --no-rdoc capistrano && \
  rm -rf /var/cache/apk/*

ADD git_ssh.sh /
ADD ssh_config /etc/ssh/
ADD drone-capistrano /bin/
ENTRYPOINT ["/bin/drone-capistrano"]
