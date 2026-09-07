package rds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsOrderableDbInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#engine DataAwsOrderableDbInstance#engine}.
	// Experimental.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#availability_zone_group DataAwsOrderableDbInstance#availability_zone_group}.
	// Experimental.
	AvailabilityZoneGroup *string `field:"optional" json:"availabilityZoneGroup" yaml:"availabilityZoneGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#engine_latest_version DataAwsOrderableDbInstance#engine_latest_version}.
	// Experimental.
	EngineLatestVersion interface{} `field:"optional" json:"engineLatestVersion" yaml:"engineLatestVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#engine_version DataAwsOrderableDbInstance#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#id DataAwsOrderableDbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#instance_class DataAwsOrderableDbInstance#instance_class}.
	// Experimental.
	InstanceClass *string `field:"optional" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#license_model DataAwsOrderableDbInstance#license_model}.
	// Experimental.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#preferred_engine_versions DataAwsOrderableDbInstance#preferred_engine_versions}.
	// Experimental.
	PreferredEngineVersions *[]*string `field:"optional" json:"preferredEngineVersions" yaml:"preferredEngineVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#preferred_instance_classes DataAwsOrderableDbInstance#preferred_instance_classes}.
	// Experimental.
	PreferredInstanceClasses *[]*string `field:"optional" json:"preferredInstanceClasses" yaml:"preferredInstanceClasses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#read_replica_capable DataAwsOrderableDbInstance#read_replica_capable}.
	// Experimental.
	ReadReplicaCapable interface{} `field:"optional" json:"readReplicaCapable" yaml:"readReplicaCapable"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#region DataAwsOrderableDbInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#storage_type DataAwsOrderableDbInstance#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supported_engine_modes DataAwsOrderableDbInstance#supported_engine_modes}.
	// Experimental.
	SupportedEngineModes *[]*string `field:"optional" json:"supportedEngineModes" yaml:"supportedEngineModes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supported_network_types DataAwsOrderableDbInstance#supported_network_types}.
	// Experimental.
	SupportedNetworkTypes *[]*string `field:"optional" json:"supportedNetworkTypes" yaml:"supportedNetworkTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_clusters DataAwsOrderableDbInstance#supports_clusters}.
	// Experimental.
	SupportsClusters interface{} `field:"optional" json:"supportsClusters" yaml:"supportsClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_enhanced_monitoring DataAwsOrderableDbInstance#supports_enhanced_monitoring}.
	// Experimental.
	SupportsEnhancedMonitoring interface{} `field:"optional" json:"supportsEnhancedMonitoring" yaml:"supportsEnhancedMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_global_databases DataAwsOrderableDbInstance#supports_global_databases}.
	// Experimental.
	SupportsGlobalDatabases interface{} `field:"optional" json:"supportsGlobalDatabases" yaml:"supportsGlobalDatabases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_iam_database_authentication DataAwsOrderableDbInstance#supports_iam_database_authentication}.
	// Experimental.
	SupportsIamDatabaseAuthentication interface{} `field:"optional" json:"supportsIamDatabaseAuthentication" yaml:"supportsIamDatabaseAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_iops DataAwsOrderableDbInstance#supports_iops}.
	// Experimental.
	SupportsIops interface{} `field:"optional" json:"supportsIops" yaml:"supportsIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_kerberos_authentication DataAwsOrderableDbInstance#supports_kerberos_authentication}.
	// Experimental.
	SupportsKerberosAuthentication interface{} `field:"optional" json:"supportsKerberosAuthentication" yaml:"supportsKerberosAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_multi_az DataAwsOrderableDbInstance#supports_multi_az}.
	// Experimental.
	SupportsMultiAz interface{} `field:"optional" json:"supportsMultiAz" yaml:"supportsMultiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_performance_insights DataAwsOrderableDbInstance#supports_performance_insights}.
	// Experimental.
	SupportsPerformanceInsights interface{} `field:"optional" json:"supportsPerformanceInsights" yaml:"supportsPerformanceInsights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_storage_autoscaling DataAwsOrderableDbInstance#supports_storage_autoscaling}.
	// Experimental.
	SupportsStorageAutoscaling interface{} `field:"optional" json:"supportsStorageAutoscaling" yaml:"supportsStorageAutoscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#supports_storage_encryption DataAwsOrderableDbInstance#supports_storage_encryption}.
	// Experimental.
	SupportsStorageEncryption interface{} `field:"optional" json:"supportsStorageEncryption" yaml:"supportsStorageEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_orderable_db_instance#vpc DataAwsOrderableDbInstance#vpc}.
	// Experimental.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

