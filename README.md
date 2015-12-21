# drone-capistrano

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-capistrano/status.svg)](http://beta.drone.io/drone-plugins/drone-capistrano)
[![](https://badge.imagelayers.io/plugins/drone-capistrano:latest.svg)](https://imagelayers.io/?images=plugins/drone-capistrano:latest 'Get your own badge on imagelayers.io')

Drone plugin for deployment via Capistrano.

## Usage

```sh
./drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone-plugins/drone-capistrano",
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
        "path": "/home/user/golang/src/github.com/drone-plugins/drone-capistrano"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone-plugins/drone-capistrano",
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
        "path": "/home/user/golang/src/github.com/drone-plugins/drone-capistrano"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```
