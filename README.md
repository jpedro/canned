# Canned

A go cli tool to store encrypted goods.


## Usage

```bash
### Install canned cli
$ go install github.com/jpedro/canned

### Check the version
$ canned version
v1.0.0

### Set the passsword and file
$ export CANNED_PASSWORD="test"

$ export CANNED_FILE="test.can"

### Start a whole new can
$ canned init
File test.can initialized.

### Add one item
$ canned set hello world
Item hello saved.

### List your items
$ canned ls
NAME    STRENTH   CREATED      UPDATED   TAGS
hello         *   2021-01-01

### Copy the hello item content to the clipboard
$ canned get hello
Item hello copied to the clipboard.
```
