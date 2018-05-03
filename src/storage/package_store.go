// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"
	"ml/src/model"
	"ml/src/util"

	"github.com/jinzhu/gorm"
)

type PackageStoreInterface interface {
	ListPackages() ([]model.Package, error)
	GetPackage(packageId uint) (*model.Package, error)
	DeletePackage(packageId uint) error
	CreatePackage(*model.Package) (*model.Package, error)
	UpdatePackageStatus(uint, model.PackageStatus) error
}

type PackageStore struct {
	db   *gorm.DB
	time util.TimeInterface
}

func (s *PackageStore) ListPackages() ([]model.Package, error) {
	var packages []model.Package
	// List all packages.
	if r := s.db.Preload("Parameters").Where("status = ?", model.PackageReady).Find(&packages); r.Error != nil {
		return nil, util.NewInternalServerError(r.Error, "Failed to list packages: %v", r.Error.Error())
	}
	return packages, nil
}

func (s *PackageStore) GetPackage(id uint) (*model.Package, error) {
	var pkg model.Package
	r := s.db.Preload("Parameters").Where("status = ?", model.PackageReady).First(&pkg, id)
	if r.RecordNotFound() {
		return nil, util.NewResourceNotFoundError("Package", fmt.Sprint(id))
	}
	if r.Error != nil {
		// TODO query can return multiple errors. log all of the errors when error handling v2 in place.
		return nil, util.NewInternalServerError(r.Error, "Failed to get package: %v", r.Error.Error())
	}
	return &pkg, nil
}

func (s *PackageStore) DeletePackage(id uint) error {
	r := s.db.Exec(`DELETE FROM packages WHERE id=?`, id)
	if r.Error != nil {
		return util.NewInternalServerError(r.Error, "Failed to delete package: %v", r.Error.Error())
	}
	return nil
}

func (s *PackageStore) CreatePackage(p *model.Package) (*model.Package, error) {
	newPackage := *p
	now := s.time.Now().Unix()
	newPackage.CreatedAtInSec = now
	if r := s.db.Create(&newPackage); r.Error != nil {
		return nil, util.NewInternalServerError(r.Error, "Failed to add package to package table: %v",
			r.Error.Error())
	}
	return &newPackage, nil
}

func (s *PackageStore) UpdatePackageStatus(id uint, status model.PackageStatus) error {
	r := s.db.Exec(`UPDATE packages SET status=? WHERE id=?`, status, id)
	if r.Error != nil {
		return util.NewInternalServerError(r.Error, "Failed to update the package metadata: %s", r.Error.Error())
	}
	return nil
}

// factory function for package store
func NewPackageStore(db *gorm.DB, time util.TimeInterface) *PackageStore {
	return &PackageStore{db: db, time: time}
}
