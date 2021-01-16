# Canned

A go cli tool to store encrypted goods.


## Usage

```bash
# Install canned
$ go get github.com/jpedro/canned/cli/canned

# Check the version
$ canned version
v0.1.0

# Set the passsword
$ export canPassword="test"

# Set the can file
$ export canFile="test.can"

# Start a whole new can
$ canned init
File test.can initialized.

# Add some items
$ canned set hello world
Item hello saved.

# List your items
$ canned ls
NAME    LENGTH   CREATED      UPDATED   TAGS
hello        5   2021-01-01

# Copy the hello item content to the clipboard
$ canned get hello
Item hello copied to the clipboard.
```
