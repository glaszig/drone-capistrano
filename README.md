# drone-capistrano [![Build Status](https://travis-ci.org/glaszig/drone-capistrano.svg?branch=master)](https://travis-ci.org/glaszig/drone-capistrano) [![ImageLayers Size](https://img.shields.io/imagelayers/image-size/glaszig/drone-capistrano/latest.svg)](https://hub.docker.com/r/glaszig/drone-capistrano/)

Drone plugin for deployment via Capistrano.

## Usage

```
./drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/glaszig/drone-capistrano",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/home/glaszig/golang/src/github.com/glaszig/drone-capistrano"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build
docker build --rm=true -t glaszig/drone-capistrano .
```

### Example

```sh
docker run -i glaszig/drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/glaszig/drone-capistrano",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/home/glaszig/golang/src/github.com/glaszig/drone-capistrano"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```
