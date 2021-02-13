# qac-cli

![CI Linux Mac](https://github.com/enr/qac-cli/workflows/CI%20Linux%20Mac/badge.svg)
![CI Windows](https://github.com/enr/qac-cli/workflows/CI%20Windows/badge.svg)
[![Download](https://img.shields.io/badge/Download-Last%20release-brightgreen)](https://github.com/enr/qac-cli/releases/latest)

Command line tool to test _end to end_ other command line tools.

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

Copyright 2020-TODAY `qac` and `qac-cli` contributors.
