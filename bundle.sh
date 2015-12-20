#!/usr/bin/env sh

eval `ssh-agent -s`
ssh-add $GIT_SSH_KEY

cd "${BUILD_PATH}"
bundle exec $@
