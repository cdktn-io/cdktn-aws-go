package awsdms


// Experimental.
type TfEndpoint_MysqlSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#after_connect_script TfEndpoint#after_connect_script}.
	// Experimental.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method TfEndpoint#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#clean_source_metadata_on_mismatch TfEndpoint#clean_source_metadata_on_mismatch}.
	// Experimental.
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#events_poll_interval TfEndpoint#events_poll_interval}.
	// Experimental.
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#execute_timeout TfEndpoint#execute_timeout}.
	// Experimental.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#max_file_size TfEndpoint#max_file_size}.
	// Experimental.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#parallel_load_threads TfEndpoint#parallel_load_threads}.
	// Experimental.
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#server_timezone TfEndpoint#server_timezone}.
	// Experimental.
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role_arn TfEndpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#target_db_type TfEndpoint#target_db_type}.
	// Experimental.
	TargetDbType *string `field:"optional" json:"targetDbType" yaml:"targetDbType"`
}

