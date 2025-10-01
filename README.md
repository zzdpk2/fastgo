## fastgo

## Deploy

```bash
$ mkdir -p  $HOME/golang/src/github.com/onexstack/
$ cd $HOME/golang/src/github.com/onexstack/
$ git clone https://github.com/onexstack/fastgo
$ cd fastgo/
$ ./build.sh
$ _output/fg-apiserver -c configs/fg-apiserver.yaml
```

**Notice：** 

1. Please login Mysql and execute `source configs/fastgo.sql;` to create `fastgo` database and tables.
2. Update `configs/fg-apiserver.yaml`  `mysql` settings.

