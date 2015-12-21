Use this plugin for deployment via [Capistrano](http://capistranorb.com/).
The following parameters are required:

- `tasks` - The Capistrano tasks to run, e.g. `production deploy`
- `bundle_path` - Path where Bundler should install gems. You probably want
this to point to your drone cache path.

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
