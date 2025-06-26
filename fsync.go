//go:build !linux && !darwin

package voidDB

func (void *Void) fsync() error {
	return void.file.Sync()
}
