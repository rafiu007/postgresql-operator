package utils

import (
	"github.com/dev4devs-com/postgresql-operator/pkg/apis/postgresql-operator/v1alpha1"
	"github.com/dev4devs-com/postgresql-operator/pkg/config"
)

var defaultPostgreSQLConfig = config.NewPostgreSQLConfig()

// AddPostgresqlMandatorySpecs will add the specs which are mandatory for PostgreSQL CR in the case them
// not be applied
func AddPostgresqlMandatorySpecs(db *v1alpha1.Postgresql) {

	/*
	   CR DB Resource
	   ---------------------
	*/

	if db.Spec.Size == 0 {
		db.Spec.Size = defaultPostgreSQLConfig.Size
	}

	/*
		Environment Variables
		---------------------
		The following values are used to create the ConfigMap and the Environment Variables which will use these values
	*/

	if db.Spec.DatabaseName == "" {
		db.Spec.DatabaseName = defaultPostgreSQLConfig.DatabaseName
	}

	if db.Spec.DatabasePassword == "" {
		db.Spec.DatabasePassword = defaultPostgreSQLConfig.DatabasePassword
	}

	if db.Spec.DatabaseUser == "" {
		db.Spec.DatabaseUser = defaultPostgreSQLConfig.DatabaseUser
	}

	/*
	   Database Container
	   ---------------------------------
	*/

	//Following are the values which will be used as the key label for the environment variable of the database image.
	if db.Spec.DatabaseNameKeyEnvVar == "" {
		db.Spec.DatabaseNameKeyEnvVar = defaultPostgreSQLConfig.DatabaseNameKeyEnvVar
	}

	if db.Spec.DatabasePasswordKeyEnvVar == "" {
		db.Spec.DatabasePasswordKeyEnvVar = defaultPostgreSQLConfig.DatabasePasswordKeyEnvVar
	}

	if db.Spec.DatabaseUserKeyEnvVar == "" {
		db.Spec.DatabaseUserKeyEnvVar = defaultPostgreSQLConfig.DatabaseUserKeyEnvVar
	}

	if db.Spec.Image == "" {
		db.Spec.Image = defaultPostgreSQLConfig.Image
	}

	if db.Spec.ContainerName == "" {
		db.Spec.ContainerName = defaultPostgreSQLConfig.ContainerName
	}

	if db.Spec.DatabaseMemoryLimit == "" {
		db.Spec.DatabaseMemoryLimit = defaultPostgreSQLConfig.DatabaseMemoryLimit
	}

	if db.Spec.DatabaseMemoryRequest == "" {
		db.Spec.DatabaseMemoryRequest = defaultPostgreSQLConfig.DatabaseMemoryRequest
	}

	if db.Spec.DatabaseStorageRequest == "" {
		db.Spec.DatabaseStorageRequest = defaultPostgreSQLConfig.DatabaseStorageRequest
	}

	if db.Spec.DatabaseCpu == "" {
		db.Spec.DatabaseCpu = defaultPostgreSQLConfig.DatabaseCpu
	}

	if db.Spec.DatabaseCpuLimit == "" {
		db.Spec.DatabaseCpuLimit = defaultPostgreSQLConfig.DatabaseCpuLimit
	}

	if db.Spec.DatabasePort == 0 {
		db.Spec.DatabasePort = defaultPostgreSQLConfig.DatabasePort
	}
}
