package packagemanager

import (
	"mime/multipart"
)

// Manager managing acutal package file.
type PackageManagerInterface interface {
	// Create the package file
	CreatePackageFile(template []byte, fileHeader *multipart.FileHeader) error

	// Get the package file
	GetPackageFile(fileName string) ([]byte, error)
}
