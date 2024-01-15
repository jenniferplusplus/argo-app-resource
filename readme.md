# Argo App Resource

A concourse resource to track and sync the health and state of applications managed by ArgoCD.

Supports checking, and fetching the current health and sync state of an application and all of its resources. Also supports syncing to latest, or to a specified revision.

## Resource Type Configuration

```yaml
resource_types:
- name: argo-app
  type: registry-image
  source:
    repository: ghcr.io/jenniferplusplus/argo-app-resource
    tag: '0.1.0'
```

## Source Configuration

| Field          | Type              | Description                                                                                                                               |
|----------------|-------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| `app`          | `string` required | The ArgoCD Application to track                                                                                                           |
| `token`        | `string` required | The access token to use for authentication                                                                                                |
| `host`         | `string` required | The hostname of the ArgoCD server                                                                                                         |
| `project`      | `string` optional | The ArgoCD Project that contains `app`. Default value is `default`                                                                        |
| `insecure`     | `bool` optional   | Whether to allow insecure connections to the host. Default value is `false`                                                               |
| `use_grpc_web` | `bool` optional   | Whether to use the fallback REST endpoints for the ArgoCD API. Equivalent to `--use-grpc-web` in the Argo CLI. Default value is `false`   |

Example configuration:

```yaml
resources:
- name: my-app
  type: argo-app
  source:
    token: ((vault:argocd.my-app-token))
    app: my-app
    project: my-app-project
    host: my.argocd.example
```

## `Check` Step

The version data recorded by `check` includes:
- The current target revision (git sha)
- The datetime when the last deployment completed
- The health of the application
- The sync state of the application

Check also records metadata about the status of the resources managed by the application. This metadata indicates the sync state and health of every individual resource reported by ArgoCD.

## `Get` Step Configuration

| Field     | Type              | Description                                     |
|-----------|-------------------|-------------------------------------------------|
| `debug`   | `bool` optional   | Enable debug logging. Default value is `false`  |

### Outputs created by `Get`

- `version.json` - Equivalent to the version information reported by `check`
- `resources.json` - Equivalent to the metadata included with the version information
- `application.json` - The full data reported by the ArgoCD API for the application.

Example `version.json`:

```json
{
  "revision": "edb47dc4fe8586c14d79437377a48d8746facb3e",
  "deployed_at": "2024-01-15T18:32:06Z",
  "health": "Degraded",
  "sync_status": "Synced"
}
```

`resources.json` identifies the resources using the format `kind`/`version`.`namespace`.`name`. Non-namespaced resources will use an underscore (`_`) instead. For example:
```json
[
  {
    "name": "Namespace/v1._.my-app",
    "value": "Synced"
  },
  {
    "name": "Service/v1.my-app.my-app-controller-service",
    "value": "Synced/Healthy"
  },
  {
    "name": "Deployment/v1.my-app.my-app-controller",
    "value": "Synced/Degraded"
  }
]
```

`application.json` is a large, complex object. You should consult the [ArgoCD API docs](https://argo-cd.readthedocs.io/en/stable/developer-guide/api-docs/) for a description of the structure and contents.

## `Put` Step Configuration

> [!WARNING]
> This feature is not thoroughly tested  
> Use with caution

I don't actually need or use the `put` feature myself. It's included for completeness' sake.


| Field                 | Type              | Description                                                                 |
|-----------------------|-------------------|-----------------------------------------------------------------------------|
| `debug`               | `bool` optional   | Enable debug logging. Default value is `false`                              |
| `rollback_revision`   | `string` optional | The desired revision to sync to. Default is to sync to the latest revision. |

## Contributing

Contributions are welcome. I built this quickly, because I needed it. It works for me, but it's not the best software engineering work I've ever done. Future work should include unit tests.
