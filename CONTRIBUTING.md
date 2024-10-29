## How to develop with Go sources

```bash
go run ./src --lang python
```

## How to build

```bash
make all
```

## How to test NPM locally

```bash
npx ./npm
```

## How to test PyPi locally

```bash
pipx run --path restack_get_started/main.py
```

## How to release

1. Increment version in all occurences (makefile, pyproject, package.json)
2. `git tag v0.6.8`
3. `git push --tags`

At this moment, Github detects new version tag and builds and publishes
