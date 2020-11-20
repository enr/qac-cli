# qac-cli

Command line tool to test end to end other command line tools.

A test plan is written in YAML format.

```yaml
preconditions:
  fs:
    - file: ../go.mod
specs:
  cat:
    command:
      working_dir: ../
      cli: cat go.mod
    expectations:
      status:
        equals_to: 0
      output:
        stdout:
          equals_to_file: ../go.mod
```

Usage:

```
qac path/to/testplay.yaml
```

You can see some example plan in `examples/` directory, and how `qac-cli` tests itself in `e2e/`.

## License

Apache 2.0 - see LICENSE file.

Copyright 2020 `qac` and `qac-cli` contributors.
