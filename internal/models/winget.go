package models

import (
	"context"

	"github.com/open-uem/ent/softwarepackage"
)

// LibraryPackageSource is the source, in the software package catalog, of the packages
// uploaded to the console's private winget repository (Deploy > Software library)
const LibraryPackageSource = "openuem"

// IsLibraryPackage reports if a package id is one of the console's winget repository's
// packages. They're installed from that repository instead of winget's community one.
func (m *Model) IsLibraryPackage(packageID string) (bool, error) {
	return m.Client.SoftwarePackage.Query().
		Where(softwarepackage.PackageID(packageID), softwarepackage.Source(LibraryPackageSource)).
		Exist(context.Background())
}
