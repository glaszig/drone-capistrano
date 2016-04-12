# Docker image for the Drone Capistrano plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-capistrano
#     make deps build docker

FROM ruby:2.3-alpine

RUN gem install capistrano

ADD bundle.sh /
ADD drone-capistrano /bin/
ENTRYPOINT ["/bin/drone-capistrano"]
