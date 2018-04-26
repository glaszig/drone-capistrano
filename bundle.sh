#!/usr/bin/env sh

eval `ssh-agent -s`
ssh-add /root/.ssh/capistrano
bundle $@
