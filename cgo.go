package gorocksdb

// #cgo CPPFLAGS: -I${SRCDIR}/deps/rocksdb/include -I${SRCDIR}/deps/rocksdb -DNDEBUG -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX -fPIC -O2
// #cgo darwin CPPFLAGS: -DOS_MACOSX -DROCKSDB_BACKTRACE
// #cgo linux CPPFLAGS: -DOS_LINUX -DROCKSDB_MALLOC_USABLE_SIZE
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -lrocksdb
// #cgo LDFLAGS: -lsnappy -llz4 -lzstd
// #cgo LDFLAGS: -lstdc++ -lm
// #cgo linux LDFLAGS: -lrt
import "C"
