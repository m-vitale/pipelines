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
	"ml/apiserver/src/message/pipelinemanager"
	"ml/apiserver/src/util"

	"github.com/jinzhu/gorm"
)

type PackageStoreInterface interface {
	ListPackages() ([]pipelinemanager.Package, error)
	GetPackage(packageId uint) (pipelinemanager.Package, error)
	CreatePackage(pipelinemanager.Package) (pipelinemanager.Package, error)
}

type PackageStore struct {
	db *gorm.DB
}

func (s *PackageStore) ListPackages() ([]pipelinemanager.Package, error) {
	var packages []pipelinemanager.Package
	// List all packages.
	if r := s.db.Preload("Parameters").Find(&packages); r.Error != nil {
		return nil, util.NewInternalError("Failed to list packages", r.Error.Error())
	}
	return packages, nil
}

func (s *PackageStore) GetPackage(id uint) (pipelinemanager.Package, error) {
	var pkg pipelinemanager.Package
	if r := s.db.Preload("Parameters").First(&pkg, id); r.Error != nil {
		// Error returns when no package found.
		return pkg, util.NewResourceNotFoundError("Package", fmt.Sprint(id))
	}
	return pkg, nil
}

func (s *PackageStore) CreatePackage(p pipelinemanager.Package) (pipelinemanager.Package, error) {
	if r := s.db.Create(&p); r.Error != nil {
		return p, util.NewInternalError("Failed to add package to package table", r.Error.Error())
	}
	return p, nil
}

// factory function for package store
func NewPackageStore(db *gorm.DB) *PackageStore {
	return &PackageStore{db: db}
}
