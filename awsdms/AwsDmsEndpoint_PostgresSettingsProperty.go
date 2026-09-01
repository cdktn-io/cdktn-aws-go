package awsdms


// Experimental.
type AwsDmsEndpoint_PostgresSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#after_connect_script AwsDmsEndpoint#after_connect_script}.
	// Experimental.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method AwsDmsEndpoint#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#babelfish_database_name AwsDmsEndpoint#babelfish_database_name}.
	// Experimental.
	BabelfishDatabaseName *string `field:"optional" json:"babelfishDatabaseName" yaml:"babelfishDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#capture_ddls AwsDmsEndpoint#capture_ddls}.
	// Experimental.
	CaptureDdls interface{} `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#database_mode AwsDmsEndpoint#database_mode}.
	// Experimental.
	DatabaseMode *string `field:"optional" json:"databaseMode" yaml:"databaseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ddl_artifacts_schema AwsDmsEndpoint#ddl_artifacts_schema}.
	// Experimental.
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#execute_timeout AwsDmsEndpoint#execute_timeout}.
	// Experimental.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#fail_tasks_on_lob_truncation AwsDmsEndpoint#fail_tasks_on_lob_truncation}.
	// Experimental.
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_enable AwsDmsEndpoint#heartbeat_enable}.
	// Experimental.
	HeartbeatEnable interface{} `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_frequency AwsDmsEndpoint#heartbeat_frequency}.
	// Experimental.
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#heartbeat_schema AwsDmsEndpoint#heartbeat_schema}.
	// Experimental.
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_boolean_as_boolean AwsDmsEndpoint#map_boolean_as_boolean}.
	// Experimental.
	MapBooleanAsBoolean interface{} `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_jsonb_as_clob AwsDmsEndpoint#map_jsonb_as_clob}.
	// Experimental.
	MapJsonbAsClob interface{} `field:"optional" json:"mapJsonbAsClob" yaml:"mapJsonbAsClob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#map_long_varchar_as AwsDmsEndpoint#map_long_varchar_as}.
	// Experimental.
	MapLongVarcharAs *string `field:"optional" json:"mapLongVarcharAs" yaml:"mapLongVarcharAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#max_file_size AwsDmsEndpoint#max_file_size}.
	// Experimental.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#plugin_name AwsDmsEndpoint#plugin_name}.
	// Experimental.
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role_arn AwsDmsEndpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#slot_name AwsDmsEndpoint#slot_name}.
	// Experimental.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

