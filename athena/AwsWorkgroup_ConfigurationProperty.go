package athena


// Experimental.
type AwsWorkgroup_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#bytes_scanned_cutoff_per_query AwsWorkgroup#bytes_scanned_cutoff_per_query}.
	// Experimental.
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// customer_content_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#customer_content_encryption_configuration AwsWorkgroup#customer_content_encryption_configuration}
	// Experimental.
	CustomerContentEncryptionConfiguration *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enable_minimum_encryption_configuration AwsWorkgroup#enable_minimum_encryption_configuration}.
	// Experimental.
	EnableMinimumEncryptionConfiguration interface{} `field:"optional" json:"enableMinimumEncryptionConfiguration" yaml:"enableMinimumEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enforce_workgroup_configuration AwsWorkgroup#enforce_workgroup_configuration}.
	// Experimental.
	EnforceWorkgroupConfiguration interface{} `field:"optional" json:"enforceWorkgroupConfiguration" yaml:"enforceWorkgroupConfiguration"`
	// engine_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#engine_version AwsWorkgroup#engine_version}
	// Experimental.
	EngineVersion *AwsWorkgroup_EngineVersionProperty `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#execution_role AwsWorkgroup#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// identity_center_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#identity_center_configuration AwsWorkgroup#identity_center_configuration}
	// Experimental.
	IdentityCenterConfiguration *AwsWorkgroup_IdentityCenterConfigurationProperty `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
	// managed_query_results_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#managed_query_results_configuration AwsWorkgroup#managed_query_results_configuration}
	// Experimental.
	ManagedQueryResultsConfiguration *AwsWorkgroup_ManagedQueryResultsConfigurationProperty `field:"optional" json:"managedQueryResultsConfiguration" yaml:"managedQueryResultsConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#monitoring_configuration AwsWorkgroup#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *AwsWorkgroup_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#publish_cloudwatch_metrics_enabled AwsWorkgroup#publish_cloudwatch_metrics_enabled}.
	// Experimental.
	PublishCloudwatchMetricsEnabled interface{} `field:"optional" json:"publishCloudwatchMetricsEnabled" yaml:"publishCloudwatchMetricsEnabled"`
	// query_results_s3_access_grants_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#query_results_s3_access_grants_configuration AwsWorkgroup#query_results_s3_access_grants_configuration}
	// Experimental.
	QueryResultsS3AccessGrantsConfiguration *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty `field:"optional" json:"queryResultsS3AccessGrantsConfiguration" yaml:"queryResultsS3AccessGrantsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#requester_pays_enabled AwsWorkgroup#requester_pays_enabled}.
	// Experimental.
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// result_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#result_configuration AwsWorkgroup#result_configuration}
	// Experimental.
	ResultConfiguration *AwsWorkgroup_ResultConfigurationProperty `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
}

