# Clients

## Vault Client

Requests against a Vault server are handled by this client. Since most requests need authentication, a Vault token can be defined in the environment or fetched by the client for example with a Kubernetes service account.

**Config**
| Environment Variable | Type   | Default      | Description                                                                                                                    |
|----------------------|--------|--------------|--------------------------------------------------------------------------------------------------------------------------------|
| VAULT_ADDR           | string | --           | see [Vault Environment Variables](https://www.vaultproject.io/docs/commands#environment-variables)                             |
| VAULT_TOKEN          | string | --           | see [Vault Environment Variables](https://www.vaultproject.io/docs/commands#environment-variables)                             |
| VAULT_TOKEN_FILE     | string | --           | file containing the value for VAULT_TOKEN, will take precedence                                                                |
| VAULT_KUBE_AUTH_NAME | string | `kubernetes` | Name of the Kubernetes auth backend mount point, see [Vault Kubernetes Auth](https://www.vaultproject.io/docs/auth/kubernetes) |
| VAULT_KUBE_AUTH_ROLE | string | --           | Name of the Kubernetes auth backend role, see [Vault Kubernetes Auth](https://www.vaultproject.io/docs/auth/kubernetes)        |

## Redis Client

This client handles Redis requests.

**Config**
| Environment Variable | Type   | Default          | Description                                                        |
|----------------------|--------|------------------|--------------------------------------------------------------------|
| REDIS_ADDR           | string | `localhost:6379` | Host and port of the Redis instance                                |
| REDIS_PASSWORD       | string | --               | An optional password for authentication                            |
| REDIS_PASSWORD_FILE  | string | --               | file containing the value for REDIS_PASSWORD, will take precedence |

## Postgres Client

This client handles Postgres requests.

**Config**
| Environment Variable     | Type   | Default                                                                | Description                                                             |
|--------------------------|--------|------------------------------------------------------------------------|-------------------------------------------------------------------------|
| POSTGRES_CONNECTION      | string | `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable` | Host and port of the Postgres instance                                  |
| POSTGRES_CONNECTION_FILE | string | --                                                                     | File containing the value for POSTGRES_CONNECTION, will take precedence |
