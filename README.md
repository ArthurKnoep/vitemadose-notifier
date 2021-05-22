# Vitemadose Notifier
A small tool to watch vitemadose appointment slot availability. The tools
run as a daemon and will send you a notification as soon as a chronodose
slot is available.

### Usage
```
usage: notifier --department=DEPARTMENT [<flags>]

A simple tool to watch vitemadose appointment slot

Flags:
  -h, --help                   Show context-sensitive help (also try --help-long and --help-man).
  -d, --department=DEPARTMENT  The department number to watch
  -r, --reload=5m              The time to wait before reloading the data
```

##### Example
Watch slot availability for Paris
```
notifier --department 75
```

### Build instruction
```bash
cd cmd/notifier
go build .
```
