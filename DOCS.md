Use this plugin for deployment via [Capistrano](http://capistranorb.com/).
The only require configuration option defines the Capistrano tasks to run.

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  capistrano:
    tasks: production deploy
    when:
      branch: master
```
