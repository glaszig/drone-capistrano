# drone-capistrano

Drone plugin for deployment via Capistrano.

## Usage

```
./drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
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
        "path": "/home/tboerger/Projects/golang/src/github.com/drone-plugins/drone-capistrano"
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
docker build --rm=true -t plugins/drone-capistrano .
```

### Example

```sh
docker run -i plugins/drone-capistrano <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
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
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```
