// Code generated by vfsgen; DO NOT EDIT.

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/source": &vfsgen۰DirInfo{
			name:    "source",
			modTime: time.Time{},
		},
		"/source/flag-types.json": &vfsgen۰CompressedFileInfo{
			name:             "flag-types.json",
			modTime:          time.Time{},
			uncompressedSize: 3364,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\x41\x6f\x9b\x40\x10\x85\xef\xfc\x8a\xe9\x56\x95\xa0\xb2\x20\x87\x28\x07\x4b\xbd\x54\x55\x2b\xdf\x2a\xb9\xe9\x25\x44\xd1\x06\x16\xb2\xea\x7a\x16\x2d\x43\xd4\xc8\xf2\x7f\xaf\x16\x9b\x04\x16\x70\x48\x5b\x2c\x5f\x7c\x78\xe3\xe5\xed\xfb\x66\x47\x73\xe3\x01\x6c\x3d\x00\x00\x86\x7c\x23\xd8\x12\xd8\x67\xad\x15\x5b\xec\x35\x7a\x2a\x6a\xed\xbe\xa5\x3d\x72\x55\x59\x31\xe3\xaa\x14\x07\x2d\x15\x25\xb1\x25\x90\xa9\x1a\x25\xd1\x48\xe2\x37\xdd\xa5\x22\xe3\x95\xb2\x45\x56\x1f\x68\xbe\x52\x70\x53\x0a\x63\xe5\x92\x4c\xa2\xf1\x31\xfc\x6e\x15\x6b\xee\x67\xe1\x4f\xeb\x11\xae\xc9\x48\xcc\xfd\x20\xe8\x58\xef\x55\x7b\xd2\x08\xaa\x0c\x42\xcc\x62\xc6\x3c\x80\xdd\x62\x38\xcc\x8f\x7f\x4f\x93\xea\x84\xb8\x54\xf6\x30\xd0\x03\x27\x90\x65\x5d\x86\xfb\x27\x68\x12\x9e\x57\xf0\x2f\x95\xe1\x24\x35\xba\xd9\x49\x6e\x44\xe8\x16\x1b\x08\xed\xc4\x47\x19\xf8\xa5\x10\xf0\x40\x54\x94\xcb\x28\xca\xb5\xe2\x98\x87\xda\xe4\x51\xf1\x2b\x8f\xac\x43\xf4\xbe\x0e\xd5\xf8\x04\x47\xe0\x5c\xf4\xc1\xd4\x77\xec\x7c\xe0\xad\x64\xdc\xbf\x8f\x51\xfa\xaa\x34\xa7\xab\x4b\x17\x52\xd6\x95\x27\xe1\x99\x16\xad\xd3\xf3\xda\xbd\x17\x6d\x01\x57\x97\xaf\xc5\xdb\x50\xb8\x2e\x8c\x44\xca\xfc\x98\x7d\xc8\x62\xb6\x68\x22\x8f\x46\xfd\x26\x50\x18\x99\xb8\x51\x1d\x79\x3c\x6a\x7b\x40\x06\xb2\xa2\x54\x6e\x93\x1b\x0b\x89\x24\x4c\xc6\x13\xb1\xdd\x8d\xc4\x92\x59\x73\x7d\x78\xf7\x09\x50\x2a\xd8\xc6\x18\x53\x4c\x23\xed\xb4\xc5\x9d\xfd\x99\x30\x08\x2b\x1c\x68\xb0\xc4\x13\xb4\x77\x85\x43\xcd\xbd\x78\x7b\x7f\xd3\x29\xfd\x5d\x21\x0d\xa4\x3c\x87\x8c\xfb\xa3\x77\x09\xaf\xfd\xec\xad\xfc\x5a\x4a\xe7\xa1\xb0\x56\x32\x11\x2e\x8a\x8f\x6e\x61\x8e\x87\x7e\x73\xdb\x22\xfe\x02\xec\x19\x90\xff\x7c\x8b\x20\xd8\x4b\x16\x96\x30\x46\x1b\x1f\xa5\x1a\xc3\x31\xfb\x74\x8c\x12\x73\x4a\xb3\x31\x7b\x99\xc5\x31\x6a\x87\x9b\x9c\x0d\xb7\x83\x8b\xc3\xac\xec\xa8\x7f\x3b\x73\xb5\x6f\x8f\x47\x7f\xca\x5e\x05\xd0\xcd\x79\x3c\xca\xf0\x1b\x18\xa8\xcd\xf3\x08\xba\xe4\x06\x5f\x41\xeb\x2e\x67\xf3\x0c\xae\xe5\xd0\x76\xa9\x4e\xb2\x5e\xac\xf7\x09\xf7\xcb\xb5\xec\x2f\x98\xea\x04\x1b\x66\x52\x4c\x67\xc5\x54\xff\x6f\xc7\x78\xb7\xde\x9f\x00\x00\x00\xff\xff\x61\x79\x63\x47\x24\x0d\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/altsrc_flags_generated.gotpl": &vfsgen۰CompressedFileInfo{
			name:             "altsrc_flags_generated.gotpl",
			modTime:          time.Time{},
			uncompressedSize: 1044,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x4f\x8b\xdb\x30\x10\xc5\xcf\xd6\xa7\x78\x84\xa5\xc4\x4b\xb0\xef\x29\x7b\x28\xdd\x2d\xf4\x92\x2d\xec\x42\xcf\x8a\x3c\xb2\x45\x15\xdb\x48\xe3\x84\x20\xfc\xdd\xcb\x38\x61\xfb\x4f\x29\xbd\xf4\x66\x69\xf4\xde\x9b\xdf\x8c\xeb\x1a\x1f\x87\x86\xd0\x52\x4f\x41\x33\x35\xd8\x9f\x61\xdb\xf7\x78\x7c\xc6\xee\xf9\x15\x4f\x8f\x9f\x5f\x2b\xa5\x46\x6d\xbe\xe9\x96\x90\x12\xaa\x2f\x97\xef\x9d\x3e\x10\xe6\x59\x29\x77\x18\x87\xc0\x58\xab\x62\x65\xbd\x6e\x57\xaa\x58\xb5\x8e\xbb\x69\x5f\x99\xe1\x50\x4f\xc1\xea\x23\xd5\xc6\xbb\x95\x2a\x55\x4a\x08\xba\x6f\x09\x77\x6e\x83\x3b\x79\x8e\xed\x03\xaa\x4f\x5e\xb7\x51\xcc\xea\x5a\x22\x96\x42\x75\x0d\x90\x1a\x5c\x04\x77\x84\x45\xc0\xe7\x91\xc0\x9d\x66\x9c\x82\x1e\x23\x8c\x77\x55\x56\xc4\x03\xb4\xf7\xc3\x49\x5c\xed\x10\x30\x70\x47\x01\x47\xed\x27\x8a\x52\xdc\x13\xe2\x48\xc6\x59\x47\x8d\x5a\x5c\xb3\x36\x91\xc3\x64\x18\x49\x15\xb7\x92\x54\x11\x89\x71\xbf\x5c\xcb\xf9\x85\x58\xcd\x4a\x62\x77\x74\xca\x7a\x9a\x40\x9a\x29\x42\xa3\xa7\x53\x36\x56\xd9\xa9\x37\xb7\xf4\x6b\xeb\x6f\x62\x97\xb8\xcf\x46\x26\x55\x04\xe2\x29\xf4\x78\x97\xab\xa7\x6c\x17\x5b\x58\xbf\x41\x24\xde\xa2\x77\x1e\xf3\x15\xeb\xc3\x38\xfa\x33\xa2\x3e\xd2\x8f\xbd\xbc\x10\x2f\x53\xf6\x9a\x29\x60\x8a\xf2\xbb\x18\xed\x7d\xdc\xc8\x93\xfe\xf2\x2d\x62\x11\xc8\xe6\x46\x6a\xb2\x99\xd5\xe2\x7e\xe1\x5f\xdb\x3c\x4d\x79\x69\x61\xfd\xc7\xdc\x4b\xe1\xb4\x95\xdc\x3f\x48\xdf\x72\xb8\x1d\x22\xfa\xf2\x67\xa6\xaf\x8e\xbb\xa7\x10\x86\xf0\x7f\xe1\xde\x62\xfe\x89\xf2\xed\x75\x0e\x97\x96\x6e\x7f\x87\xbe\x6e\xfa\x6f\xec\xbf\xb8\xca\x10\x52\x02\xf5\x0d\xe6\xf9\x7b\x00\x00\x00\xff\xff\x6d\x3d\x9e\xe0\x14\x04\x00\x00"),
		},
		"/templates/cli_flags_generated.gotpl": &vfsgen۰CompressedFileInfo{
			name:             "cli_flags_generated.gotpl",
			modTime:          time.Time{},
			uncompressedSize: 2591,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4d\x8f\xdb\x36\x10\x3d\x4b\xbf\x62\x6a\xe4\x20\x2d\x5c\xfb\xde\x20\xa7\x75\x36\x0d\x50\x24\x8b\xc6\xcd\x9d\x96\x86\x32\xb1\x34\xa9\x90\xa3\xdd\x2e\x0c\xff\xf7\x62\x48\xea\xc3\x1f\x58\xfb\xb2\x3d\x59\xc3\x19\xbe\xf7\xf8\x66\x48\x2f\x97\x70\x6f\x6b\x84\x06\x0d\x3a\x41\x58\xc3\xe6\x15\x64\xf3\x11\x56\xdf\xe1\xdb\xf7\x35\x7c\x5e\x7d\x5d\x2f\xf2\xbc\x15\xd5\x93\x68\x10\xf6\x7b\x58\x3c\xc6\xef\x6f\x62\x87\x70\x38\xe4\xb9\xda\xb5\xd6\x11\x14\x79\x36\x93\x5a\x34\x33\xfe\xdd\x11\xff\x78\x72\x95\x35\xcf\xfc\x49\x6a\x87\xb3\xbc\xcc\xf7\x7b\x70\xc2\x34\x08\x1f\xd4\x1c\x3e\x70\x3d\xfc\xf1\x09\x16\x0f\x5a\x34\x9e\xd1\x96\x4b\xe6\x08\x89\x45\x62\xe0\x1c\x28\x0f\x02\x42\xf9\x8b\xa2\x2d\xd0\x6b\x8b\x63\xe1\x9a\xa3\xc3\x61\x88\x57\xb6\x22\xa1\x34\xe3\x1d\x17\x4e\x11\x3d\xb9\xae\x22\xd8\xe7\x19\xaf\x66\x99\x27\xa7\x4c\x93\x67\xff\x78\xd1\x4c\xc2\xcf\xe6\xf9\xa7\x70\x63\xfc\xa0\x34\x3e\x0a\xda\x0e\x0b\x7f\xe3\xaf\x4e\x39\xac\xb3\x8d\xb5\x3a\xcf\xfe\x54\x75\x8d\x26\x4b\xd1\x5a\x3c\xa1\xe7\x2d\x29\xde\xef\x7f\x07\x25\x01\x7f\x25\x41\x3f\x85\xee\x10\xc8\x75\xc1\xca\x2c\x84\xd9\xe9\xb9\xe2\x36\x34\xf5\xf0\x3d\x85\x58\xa1\x27\x65\x04\x29\x6b\x46\xa0\xc9\x62\x76\xf7\x26\xde\x21\x67\xcf\x7f\x84\xc3\x80\x43\xea\x9c\x61\xab\x1d\x8a\x5a\x6c\x34\x82\xc3\xd6\xa1\x47\x43\x91\xc1\x4a\xa0\xad\xf2\xf0\xcc\x4a\x79\x67\x21\xad\x83\x8e\x3d\x83\x1a\xa5\xe8\x34\xf9\x32\x97\x9d\xa9\xa0\x90\x17\x8d\x2f\x13\x59\x51\x42\xb4\x90\x5b\x10\x89\x81\xf3\x31\x8b\xae\x90\x65\x12\xf7\x05\x29\xec\xef\xd5\xd1\x16\xc1\xf0\x42\x10\x83\x61\x2c\xae\x50\x26\x88\x4b\x9c\x32\x54\x26\xaa\xaf\xbe\x6f\xe7\xc0\xf6\xb2\x45\xda\xa2\x03\xeb\xc0\x58\x1a\x08\x79\x22\x5d\xaa\xbd\x42\x3e\x82\x16\x25\xf0\x1c\x1c\xb1\xf7\xb9\xa4\x20\x4c\x4c\x1c\x8b\xe1\xbc\xdc\xd5\xc9\x59\x81\xb8\x06\x44\x6c\xc2\x1c\x2c\xeb\x7b\x51\x1e\x41\x0a\xed\xf1\x8a\x9a\x91\xe0\x5c\xcd\xb0\x23\x0a\x18\xc6\xe3\x0b\x52\xb8\x16\x47\x2d\x88\x4d\x4f\x7e\xf2\x14\xdc\xde\x8b\x00\x76\xb9\x19\x21\x35\xd2\x9e\x38\x91\x18\xd2\xfc\x81\xf0\x3d\xc4\xc9\x98\x0a\x53\x83\x30\x80\xbb\x96\x5e\x19\x29\x55\xa9\x33\x13\x8d\xed\xa1\x08\x84\xd6\x8b\xeb\xda\x7b\xeb\x46\xed\xc7\xae\xa5\xab\x34\x78\x77\x0a\x04\xda\xda\x27\x0f\x5d\x1b\xa4\x44\x72\x2b\x41\x80\xb6\x95\xd0\x17\x79\xe7\xbd\x01\x47\x78\xf7\xd6\x10\xfe\x4b\xab\x78\xeb\x18\x59\xc9\x30\xa2\xd2\x76\xa6\x9f\xc9\x0a\xee\x52\x5d\x79\x06\x5d\x84\x4b\x14\xcf\x11\xb2\xbc\x1f\xa1\x07\x0e\xaf\xc5\x6c\x36\x7d\x58\xa7\x99\xb0\x8c\xda\x1f\xbd\xbc\x43\xa6\x7f\x5f\x26\xbd\xe5\x73\x77\xed\x45\x11\x73\xa8\x16\xbc\xf8\x03\x69\xb8\xf4\xda\x6e\x84\xbe\xd9\xbc\x26\x94\xbf\xa7\x7b\x97\x05\xfd\x5f\x1e\x2a\x09\xd2\xf3\x5f\x65\x74\x31\x8a\x79\x88\x96\xf5\x16\x96\x1f\xb9\xe6\xb7\x4f\x60\x54\xb8\xd4\x37\x19\x2f\x7d\x99\x67\x87\xf3\x17\xe0\xcc\x1f\xee\x4b\x70\xe5\x0d\xb8\x64\xc4\x1c\x3c\x12\xdc\x85\x74\xd2\xf8\x9e\xde\x48\xf6\xc5\x23\x2d\xfe\x0a\xca\x82\x92\x32\x5a\x36\x75\x63\x14\xf0\x28\x9c\x47\x17\xb9\x5b\xfe\xae\xe7\x80\xce\x31\xcc\xc0\x94\x6a\xa6\x2a\x4e\x4a\x65\xbc\xef\x21\xb6\xae\x30\x4a\x97\x5c\x9b\xfe\xa6\x99\x9e\x0b\x27\x02\x6e\xf0\x38\xe3\x56\x9c\x29\xbd\x17\x9e\xa2\xda\x53\x84\x49\x7e\xaa\x34\x95\x45\xc1\x13\x51\x37\xf6\x79\xdc\xf1\x5f\x00\x00\x00\xff\xff\xdb\x2f\x5d\x87\x1f\x0a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/source"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/source"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/source/flag-types.json"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/altsrc_flags_generated.gotpl"].(os.FileInfo),
		fs["/templates/cli_flags_generated.gotpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}