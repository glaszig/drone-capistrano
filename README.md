# Drone Capistrano Plugin [![Build Status](https://travis-ci.org/glaszig/drone-capistrano.svg?branch=master)](https://travis-ci.org/glaszig/drone-capistrano) [![ImageLayers](https://badge.imagelayers.io/glaszig/drone-capistrano:latest.svg)](https://imagelayers.io/?images=glaszig/drone-capistrano:latest 'Get your own badge on imagelayers.io')

This is the Capistrano plugin for the Drone continuous integration platform.

## Usage

Configure your `drone.yml` like so.

```yaml
build:
  environment:
    BUNDLE_APP_CONFIG=.bundle
  commands:
    - bundle install --path vendor/bundle

deploy:
  capistrano:
    tasks: production deploy
    when:
      branch: master
```

Use this plugin for deployment via [Capistrano](http://capistranorb.com/).

The Docker image is based on the official `ruby:2.3-alpine` and should only
be used if your project is building on Ruby 2.3.

To have the Capistrano plugin properly pickup your gems make sure your Bundler
installs gems into the build path by setting a proper `BUNDLE_APP_PATH` env var
and running bundler with the `--path` option.
(see example above).

The following parameters are required:

- `tasks` - The Capistrano tasks to run, e.g. `production deploy`
