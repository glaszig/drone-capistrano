Use this plugin for deployment via [Capistrano](http://capistranorb.com/).
The only requiree configuration option defines the Capistrano tasks to run.
Option `bundle_path` allows you to specify the path where Bundler should
install Gems.

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  capistrano:
    tasks: production deploy
    bundle_path: vendor/bundle
    when:
      branch: master
```
