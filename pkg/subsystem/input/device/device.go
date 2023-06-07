package device

import (
	"io"
	"os"
)

type DeviceSubsystem struct {
	ROMName string
	Size    uint
	file    *os.File
}

func New(fname string) (*DeviceSubsystem, error) {
	size, file, err := getFile(fname)
	if err != nil {
		goto err
	}

	return &DeviceSubsystem{
		ROMName: fname,
		Size:    size,
		file:    file,
	}, nil

err:
	return nil, err
}

func Destroy(input *DeviceSubsystem) {
	input.file.Close()
	input = nil
}

func getFile(fname string) (uint, *os.File, error) {
	var (
		size int64
		file *os.File
		err  error
	)
	file, err = os.Open(fname)
	if err != nil {
		goto err
	}

	size, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		goto err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		goto err
	}

	return uint(size), file, nil
err:
	return 0, &os.File{}, err
}
