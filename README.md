# nutshttp

A http server for nutsdb

## Example

Run http server and listen ON ":8080".

```bash
go run examples/hello.go
```

**Check example data**

```bash
# Get all members in set
curl http://localhost:8080/set/bucket001/foo


# List all list
curl http://localhost:8080/list/bucket001/key1?start=0&end=10
```

**modify example data**

You can modify the tests in the sample program by modifying the file "example/init.yaml".

The file format is as follows:

```yaml
# kv data
kv:
  bucket-a:  # bucket name
    key-1:  # key
      str: value-1  # value
    key-2:  # key
      base64: dW50c2Ri  # data in base64 format
# list data
list:
  bucket-b:  # bucket name
    key-1:
      - str: value-1
      - str: value-2
# set data
set:
  bucket-c:  # bucket name
    key-1:
      - str: value-1
      - str: value-2
# sorted-set data
zset:
  bucket-d:  # bucket name
    key-1:
      - score: 1660575966082
        str: value

```
