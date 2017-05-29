#!/usr/bin/env sh

eval `ssh-agent -s`
ssh-add /root/.ssh/capistrano
/usr/local/bundle/bin/bundle $@
