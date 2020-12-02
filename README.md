# Advent Of Code 2020

Solutions for [Advent Of Code 2020](https://adventofcode.com/2020).

## Create new Go app

In terminal:

```bash
nx g @nx-go/nx-go:app <app-name>
```

In `workspace.json`:

set `projects.<app-name>.architect.serve.options.cmd` to `"gow"`.

### Build

```bash
nx build <app-name>
```

### Test

Don't use `nx test` because it currently has a bug that causes it to only run main_test.go.

```bash
# Run all tests
go test ./...

# with watch
gow test ./...
```

```bash
# Run tests for single package
go test <relative-path-to-package-folder>

# with watch
gow test <relative-path-to-package-folder>
```

### Serve (with watch)

```bash
nx serve <app-name>
```