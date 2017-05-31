Use this plugin to deploy applications with capistrano.

## Config

The following parameters are used to configure the plugin:

* **capistrano_private_key** - Private SSH deploy key
* **capistrano_public_key** - Public SSH deploy key
* **tasks** - The Capistrano tasks to run

The following secret values can be set to configure the plugin.

* **CAPISTRANO_PRIVATE_KEY** - corresponds to **capistrano_private_key**
* **CAPISTRANO_PUBLIC_KEY** - corresponds to **capistrano_public_key**

It is highly recommended to put **CAPISTRANO_PRIVATE_KEY** into a secret so
it is not exposed to users. This can be done using the drone-cli.

```bash
drone secret add \
  --name capistrano_private_key \
  --value @$HOME/.ssh/drone-deploy \
  --image glaszig/drone-capistrano
  --repository octocat/hello-world

drone secret add \
  --name capistrano_public_key \
  --value @$HOME/.ssh/drone-deploy.pub \
  --image glaszig/drone-capistrano
  --repository octocat/hello-world
```

See [secrets](http://docs.drone.io/manage-secrets/) for additional
information on secrets

## Examples

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  deploy:
    image: glaszig/drone-capistrano
    repo: octocat/hello-world
    tasks: production deploy
    secrets:
      - capistrano_private_key
      - capistrano_public_key
    when:
      event: push
```
