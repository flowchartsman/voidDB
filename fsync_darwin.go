package voidDB

import (
	"golang.org/x/sys/unix"
)

// Apple's implermentation of fsync() flushes to the OS buffer cache, but not
// necessarily to disk. To get as close as possible to Fdatasync, it's necessary
// to use fnctl with F_FULLFSYNC.
//
// From `man fsync` on OSX:
//
//	For applications that require tighter guarantees about the integrity of
//	their data, Mac OS X provides the F_FULLFSYNC fcntl.
//	The F_FULLFSYNC fcntl asks the drive to flush all buffered data to
//	permanent storage. Applications, such as databases, that require a strict
//	ordering of writes should use F_FULLFSYNC to ensure that their data is
//	written in the order they expect.
//
// According to `man 2 fcntl` this:
//
//	drains the entire queue of the device and acts as a barrier, [and] data
//	that had been fsync'd on the same device before is guaranteed to be
//	persisted when this call returns.
//
// Currently implemented on HFS, FAT, UDF and APFS.
func (void *Void) fsync() error {
	_, err := unix.FcntlInt(void.file.Fd(), unix.F_FULLFSYNC, 0)
	return err
}
