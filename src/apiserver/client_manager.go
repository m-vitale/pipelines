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

package main

import (
	"database/sql"
	"fmt"
	"ml/src/client"
	"ml/src/model"
	"ml/src/storage"
	"ml/src/util"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	minio "github.com/minio/minio-go"
)

const (
	minioServiceHost = "MINIO_SERVICE_SERVICE_HOST"
	minioServicePort = "MINIO_SERVICE_SERVICE_PORT"
	mysqlServiceHost = "MYSQL_SERVICE_HOST"
	mysqlServicePort = "MYSQL_SERVICE_PORT"
	podNamespace     = "POD_NAMESPACE"
	dbName           = "mlpipeline"
)

// Container for all service clients
type ClientManager struct {
	db            *gorm.DB
	packageStore  storage.PackageStoreInterface
	pipelineStore storage.PipelineStoreInterface
	jobStore      storage.JobStoreInterface
	objectStore   storage.ObjectStoreInterface
	time          util.TimeInterface
	uuid          util.UUIDGeneratorInterface
}

func (c *ClientManager) PackageStore() storage.PackageStoreInterface {
	return c.packageStore
}

func (c *ClientManager) PipelineStore() storage.PipelineStoreInterface {
	return c.pipelineStore
}

func (c *ClientManager) JobStore() storage.JobStoreInterface {
	return c.jobStore
}

func (c *ClientManager) ObjectStore() storage.ObjectStoreInterface {
	return c.objectStore
}

func (c *ClientManager) Time() util.TimeInterface {
	return c.time
}

func (c *ClientManager) UUID() util.UUIDGeneratorInterface {
	return c.uuid
}

func (c *ClientManager) init() {
	glog.Infof("Initializing client manager")

	db := initDBClient()

	// time
	c.time = util.NewRealTime()

	// UUID generator
	c.uuid = util.NewUUIDGenerator()

	// Initialize package store
	c.db = db
	c.packageStore = storage.NewPackageStore(db, c.time)

	// Initialize pipeline store
	c.db = db
	c.pipelineStore = storage.NewPipelineStore(db, c.time)

	// Initialize job store
	wfClient := client.CreateWorkflowClientOrFatal(getConfig(podNamespace))
	c.jobStore = storage.NewJobStore(db, wfClient, c.time)

	// Initialize package manager.
	c.objectStore = initMinioClient()

	glog.Infof("Client manager initialized successfully")
}

func (c *ClientManager) Close() {
	c.db.Close()
}

func initDBClient() *gorm.DB {
	driverName := getConfig("DBConfig.DriverName")
	var arg string
	var db *gorm.DB

	switch driverName {
	case "mysql":
		arg = initMysql(driverName)
	default:
		glog.Fatalf("Driver %v is not supported", driverName)
	}

	// db is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections.
	db, err := gorm.Open(driverName, arg)
	util.TerminateIfError(err)

	// Create table
	response := db.AutoMigrate(&model.Package{}, &model.Pipeline{},
		&model.Parameter{}, &model.Job{})
	if response.Error != nil {
		glog.Fatalf("Failed to initialize the databases.")
	}

	// Create a foreign key so deleting a pipeline will delete all jobs associated with it.
	response = db.Model(&model.Job{}).
		AddForeignKey("pipeline_id", "pipelines(id)", "CASCADE" /* onDelete */, "CASCADE" /* update */)
	if response.Error != nil {
		glog.Fatalf("Failed to create a foreign key for pipelineId in job table.")
	}
	return db
}

// Initialize the connection string for connecting to Mysql database
// Format would be something like root@tcp(ip:port)/dbname?charset=utf8&loc=Local&parseTime=True
func initMysql(driverName string) string {
	mysqlConfig := client.CreateMySQLConfig(
		"root",
		getConfig(mysqlServiceHost),
		getConfig(mysqlServicePort),
		"")
	db, err := sql.Open(driverName, mysqlConfig.FormatDSN())
	defer db.Close()
	util.TerminateIfError(err)

	// Create database if not exist
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	util.TerminateIfError(err)
	mysqlConfig.DBName = dbName
	return mysqlConfig.FormatDSN()
}

func initMinioClient() storage.ObjectStoreInterface {
	// Create minio client.
	minioServiceHost := getConfig(minioServiceHost)
	minioServicePort := getConfig(minioServicePort)
	accessKey := getConfig("ObjectStoreConfig.AccessKey")
	secretKey := getConfig("ObjectStoreConfig.SecretAccessKey")
	bucketName := getConfig("ObjectStoreConfig.BucketName")

	minioClient := client.CreateMinioClientOrFatal(minioServiceHost, minioServicePort, accessKey,
		secretKey)
	createMinioBucket(minioClient, bucketName)

	return storage.NewMinioObjectStore(&storage.MinioClient{Client: minioClient}, bucketName)
}

func createMinioBucket(minioClient *minio.Client, bucketName string) {
	// Create bucket if it does not exist
	err := minioClient.MakeBucket(bucketName, "")
	if err != nil {
		// Check to see if we already own this bucket.
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			glog.Infof("We already own %s\n", bucketName)
		} else {
			glog.Fatalf("Failed to create Minio bucket. Error: %v", err)
		}
	}
	glog.Infof("Successfully created bucket %s\n", bucketName)
}

// newClientManager creates and Init a new instance of ClientManager
func newClientManager() ClientManager {
	clientManager := ClientManager{}
	clientManager.init()

	return clientManager
}
