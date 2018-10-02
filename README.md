```bash
git subtree add --prefix deps/rocksdb https://github.com/facebook/rocksdb.git v5.15.10 --squash
git subtree add --prefix deps/lz4 git@github.com:lz4/lz4.git v1.8.3 --squash

pushd deps/rocksdb
./build_tools/build_detect_platform /tmp/rocksdb
```

```bash
CC=cc
CXX=g++
PLATFORM=OS_LINUX
PLATFORM_LDFLAGS= -lpthread -lrt -lsnappy -lgflags -lz -lbz2 -llz4 -lzstd
JAVA_LDFLAGS= -lpthread -lrt -lsnappy -lz -lbz2 -llz4 -lzstd
JAVA_STATIC_LDFLAGS= -lpthread -lrt
VALGRIND_VER=
PLATFORM_CCFLAGS= -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_FALLOCATE_PRESENT -DSNAPPY -DGFLAGS=1 -DZLIB -DBZIP2 -DLZ4 -DZSTD -DROCKSDB_MALLOC_USABLE_SIZE -DROCKSDB_PTHREAD_ADAPTIVE_MUTEX -DROCKSDB_BACKTRACE -DROCKSDB_RANGESYNC_PRESENT -DROCKSDB_SCHED_GETCPU_PRESENT -march=native  -DHAVE_SSE42 -DHAVE_PCLMUL -DROCKSDB_SUPPORT_THREAD_LOCAL
PLATFORM_CXXFLAGS=-std=c++11  -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_FALLOCATE_PRESENT -DSNAPPY -DGFLAGS=1 -DZLIB -DBZIP2 -DLZ4 -DZSTD -DROCKSDB_MALLOC_USABLE_SIZE -DROCKSDB_PTHREAD_ADAPTIVE_MUTEX -DROCKSDB_BACKTRACE -DROCKSDB_RANGESYNC_PRESENT -DROCKSDB_SCHED_GETCPU_PRESENT -march=native  -DHAVE_SSE42 -DHAVE_PCLMUL -DROCKSDB_SUPPORT_THREAD_LOCAL
PLATFORM_SHARED_CFLAGS=-fPIC
PLATFORM_SHARED_EXT=so
PLATFORM_SHARED_LDFLAGS=-Wl,--no-as-needed -shared -Wl,-soname -Wl,
PLATFORM_SHARED_VERSIONED=true
EXEC_LDFLAGS=
JEMALLOC_INCLUDE=
JEMALLOC_LIB=
ROCKSDB_MAJOR=5
ROCKSDB_MINOR=16
ROCKSDB_PATCH=0
CLANG_SCAN_BUILD=scan-build
CLANG_ANALYZER=
PROFILING_FLAGS=-pg
FIND=find
WATCH=watch
LUA_PATH=
```

```
GODEBUG=cgocheck=2 go test -a -v .
```

# gorocksdb, a Go wrapper for RocksDB

[![Build Status](https://travis-ci.org/tecbot/gorocksdb.svg)](https://travis-ci.org/tecbot/gorocksdb) [![GoDoc](https://godoc.org/github.com/tecbot/gorocksdb?status.svg)](http://godoc.org/github.com/tecbot/gorocksdb)

## Install

You'll need to build [RocksDB](https://github.com/facebook/rocksdb) v5.5+ on your machine.

After that, you can install gorocksdb using the following command:

    CGO_CFLAGS="-I/path/to/rocksdb/include" \
    CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
      go get github.com/tecbot/gorocksdb

Please note that this package might upgrade the required RocksDB version at any moment.
Vendoring is thus highly recommended if you require high stability.

*The [embedded CockroachDB RocksDB](https://github.com/cockroachdb/c-rocksdb) is no longer supported in gorocksdb.*
