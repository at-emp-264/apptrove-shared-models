# Compiling Cap'n Proto Schemas

## Prerequisites

### 1. Install the `capnp` compiler

```sh
brew install capnp
```

### 2. Add the Go capnp dependency

```sh
go get capnproto.org/go/capnp/v3
```

### 3. Install the `capnpc-go` plugin

```sh
go install capnproto.org/go/capnp/v3/capnpc-go@latest
```

Make sure `$(go env GOPATH)/bin` is in your `PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Add the above line to your `~/.zshrc` (or `~/.bashrc`) to make it permanent.

---

## Compiling Schemas

Set the include path to the Cap'n Proto Go standard library once, then pass all `.capnp` files you want to compile:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"

CAPNP_GO_STD="$(go list -f '{{.Dir}}' -m capnproto.org/go/capnp/v3)/std"

capnp compile \
  -I"$CAPNP_GO_STD" \
  -ogo \
  dto/common/envelope.capnp \
  dto/clicks/click.capnp
```

> **Note:** Always compile a schema **and all schemas it imports** together. For example, `click.capnp` imports `envelope.capnp`, so both must be compiled.

Each `.capnp` file produces a corresponding `.go` file in the same directory:

| Schema | Generated file |
|---|---|
| `dto/common/envelope.capnp` | `dto/common/envelope.capnp.go` |
| `dto/clicks/click.capnp` | `dto/clicks/click.capnp.go` |

---

## Adding a New Schema

1. Create your `.capnp` file in the appropriate `dto/` subdirectory.
2. If it imports another schema, make sure that schema is also compiled.
3. Add the new file to the compile command above.

---

## Troubleshooting

| Error | Fix |
|---|---|
| `go: module capnproto.org/go/capnp/v3: not a known dependency` | Run `go get capnproto.org/go/capnp/v3` |
| `capnpc-go not found` | Run `go install capnproto.org/go/capnp/v3/capnpc-go@latest` and add `$(go env GOPATH)/bin` to `PATH` |
| `-I /std: no such directory` | The module is not in `go.mod` yet — run `go get capnproto.org/go/capnp/v3` first |
| `import not found: envelope.capnp` | Use the relative path from the importing file, e.g. `"../common/envelope.capnp"` |
