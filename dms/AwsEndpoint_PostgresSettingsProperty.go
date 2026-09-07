package dms


// Experimental.
type AwsEndpoint_PostgresSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#after_connect_script AwsEndpoint#after_connect_script}.
	// Experimental.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method AwsEndpoint#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#babelfish_database_name AwsEndpoint#babelfish_database_name}.
	// Experimental.
	BabelfishDatabaseName *string `field:"optional" json:"babelfishDatabaseName" yaml:"babelfishDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#capture_ddls AwsEndpoint#capture_ddls}.
	// Experimental.
	CaptureDdls interface{} `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#database_mode AwsEndpoint#database_mode}.
	// Experimental.
	DatabaseMode *string `field:"optional" json:"databaseMode" yaml:"databaseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ddl_artifacts_schema AwsEndpoint#ddl_artifacts_schema}.
	// Experimental.
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#execute_timeout AwsEndpoint#execute_timeout}.
	// Experimental.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#fail_tasks_on_lob_truncation AwsEndpoint#fail_tasks_on_lob_truncation}.
	// Experimental.
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_enable AwsEndpoint#heartbeat_enable}.
	// Experimental.
	HeartbeatEnable interface{} `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_frequency AwsEndpoint#heartbeat_frequency}.
	// Experimental.
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_schema AwsEndpoint#heartbeat_schema}.
	// Experimental.
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_boolean_as_boolean AwsEndpoint#map_boolean_as_boolean}.
	// Experimental.
	MapBooleanAsBoolean interface{} `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_jsonb_as_clob AwsEndpoint#map_jsonb_as_clob}.
	// Experimental.
	MapJsonbAsClob interface{} `field:"optional" json:"mapJsonbAsClob" yaml:"mapJsonbAsClob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_long_varchar_as AwsEndpoint#map_long_varchar_as}.
	// Experimental.
	MapLongVarcharAs *string `field:"optional" json:"mapLongVarcharAs" yaml:"mapLongVarcharAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#max_file_size AwsEndpoint#max_file_size}.
	// Experimental.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#plugin_name AwsEndpoint#plugin_name}.
	// Experimental.
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role_arn AwsEndpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#slot_name AwsEndpoint#slot_name}.
	// Experimental.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

