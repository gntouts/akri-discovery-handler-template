# akri-discovery-handler-template

Quickly create the boilerplate for a simple Go CLI.

## Usage

This requires `gonew`:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Create your CLI:

```
gonew github.com/trstringer/akri-discovery-handler-template your-domain.com/your-project
```

Or to use a specific version/ref of the upstream template:

```
gonew github.com/trstringer/akri-discovery-handler-template@<git_ref> your-domain.com/your-project
```

The new project/module is created in the project directory:

```
cd ./your-project
```

Run setup:

```
make setup
```

## Tests

Run unit tests:

```
make test
```

Run e2e tests:

```
make e2e
```
