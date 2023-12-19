# OP-Secrets

1Password secrets management CLI to manage service secrets for containers and local development.

- We typically have one `secrets.yaml` secret config file per service. E.g. [secrets.yaml](./example/secrets.yaml)
- Provider precedence `order` applies from lower to upper
- Change directory `--chdir` option is useful when we have relative paths to `dotenv` files

## Development

Either run from the root of the repo or from [this](./) directory

```shell
# Run the golang code (aliases: run, dev)
$ task start -- --chdir=example -c=secrets.yaml list

# Reformat the golang code
$ task fmt

# Update golang module dependencies
$ task update-deps

# Compile the op-secrets binary (for all OS and architectures)
$ task build-all
```

## Usage

> If you have multiple 1password accounts you may have to use `$ OP_ACCOUNT=knock op-secrets ...`

```shell
$ op-secrets --help

Usage: op-secrets [options] command

Commands:
  list    Lists secrets in a human-friendly format (redacted)
  env     Get secrets in .env format
  inline  Get secrets in an inline format
  sh      Get secrets in shell export format

Options:
  --config file, -c file  Path to YAML secrets config file (required)
  --chdir directory       Switch to a different directory before executing the command
  --help, -h              show help
  --version, -v           print the version
```

```shell
# Lists secrets in a human-friendly format (redacted)
$ op-secrets --config=secrets.yaml list

# Get secrets in .env format
$ op-secrets --config=secrets.yaml env

# Populate secrets in local shell
$ eval "$(op-secrets --config=secrets.yaml sh)"

# Use it with Docker container or compose
$ docker container run --rm -it \
  --env-file=<(op-secrets --config=secrets.yaml env) alpine sh

# Test provided example/secrets.yaml file
$ op-secrets --chdir=example --config=secrets.yaml list
```
