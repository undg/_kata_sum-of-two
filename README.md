You have array of integers.

Find 2 numbers in array that will sum to target and return their indices.

e.g.

Numbers: [2,7,11,15]
Target: 9
Out: [0, 1]

## How to run tests

### Go

```bash
go test ./go/...
```

Or run tests on file changes with the included watch script:

```bash
./go/test-watch.sh
```

### Python

```bash
python3 -m pytest py/test_main.py
```

### Lua

```bash
busted init_test.lua
```

### JavaScript (Bun)

```bash
bun test
```
