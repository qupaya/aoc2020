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

```bash
nx test <app-name>
```

### Serve (with watch)

```bash
nx serve <app-name>
```