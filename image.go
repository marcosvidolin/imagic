package main

import (
	"bytes"
	"errors"
)

// ImageFormat represents the type of an image
type ImageFormat int

const (
	UnknownFormat ImageFormat = iota
	JPEG
	PNG
	GIF
	BMP
	WEBP
	TIFF
	TIFF_BE // Big-endian TIFF
)

var imageFormatNames = map[ImageFormat]string{
	UnknownFormat: "Unknown",
	JPEG:          "JPEG",
	PNG:           "PNG",
	GIF:           "GIF",
	BMP:           "BMP",
	WEBP:          "WEBP",
	TIFF:          "TIFF",
	TIFF_BE:       "TIFF_BE",
}

// MagicNumbers is a map of image formats and their corresponding magic numbers
var MagicNumbers = map[ImageFormat][]byte{
	JPEG:    {0xFF, 0xD8, 0xFF},
	PNG:     {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
	GIF:     {0x47, 0x49, 0x46, 0x38},
	BMP:     {0x42, 0x4D},
	WEBP:    {0x52, 0x49, 0x46, 0x46}, // RIFF, requires additional check for "WEBP"
	TIFF:    {0x49, 0x49, 0x2A, 0x00}, // Little-endian TIFF
	TIFF_BE: {0x4D, 0x4D, 0x00, 0x2A}, // Big-endian TIFF
}

// String converts the ImageFormat to its string representation
func (f ImageFormat) String() string {
	if name, ok := imageFormatNames[f]; ok {
		return name
	}
	return "Unknown"
}

type Image struct {
	Format  ImageFormat
	Content []byte
}

func New(data []byte) (*Image, error) {
	return &Image{
		Format:  0,
		Content: []byte{},
	}, nil
}

// DetectImageFormat reads the bytes of an image and detects its format
func DetectImageFormat(data []byte) (*Image, error) {
	if len(data) < 16 {
		return nil, errors.New("insufficient data to determine format")
	}

	// Match magic numbers
	for format, magic := range MagicNumbers {
		if bytes.HasPrefix(data, magic) {
			// Special case for WEBP to check "WEBP" after "RIFF"
			if format == WEBP && !bytes.Contains(data, []byte("WEBP")) {
				continue
			}

			return &Image{Format: format, Content: data}, nil
		}
	}

	return &Image{Format: UnknownFormat, Content: data}, nil
}
