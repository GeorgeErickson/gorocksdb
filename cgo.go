package gorocksdb

// #cgo CPPFLAGS: -I${SRCDIR}/deps/rocksdb/include -I${SRCDIR}/deps/rocksdb
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -lrocksdb
// #cgo LDFLAGS: -lsnappy -llz4 -lzstd
// #cgo LDFLAGS: -lstdc++ -lm
import "C"
