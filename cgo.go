package gorocksdb

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I${SRCDIR}/deps/rocksdb/include -I${SRCDIR}/deps/rocksdb
// #cgo LDFLAGS: -lstdc++
// #cgo darwin LDFLAGS: -lrocksdb
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all -lrt
import "C"
