package awsdms


// Experimental.
type AwsDmsEndpoint_OracleSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#access_alternate_directly AwsDmsEndpoint#access_alternate_directly}.
	// Experimental.
	AccessAlternateDirectly interface{} `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#additional_archived_log_dest_id AwsDmsEndpoint#additional_archived_log_dest_id}.
	// Experimental.
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#add_supplemental_logging AwsDmsEndpoint#add_supplemental_logging}.
	// Experimental.
	AddSupplementalLogging interface{} `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#allow_selected_nested_tables AwsDmsEndpoint#allow_selected_nested_tables}.
	// Experimental.
	AllowSelectedNestedTables interface{} `field:"optional" json:"allowSelectedNestedTables" yaml:"allowSelectedNestedTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_log_dest_id AwsDmsEndpoint#archived_log_dest_id}.
	// Experimental.
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_logs_only AwsDmsEndpoint#archived_logs_only}.
	// Experimental.
	ArchivedLogsOnly interface{} `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_password AwsDmsEndpoint#asm_password}.
	// Experimental.
	AsmPassword *string `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_server AwsDmsEndpoint#asm_server}.
	// Experimental.
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_user AwsDmsEndpoint#asm_user}.
	// Experimental.
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method AwsDmsEndpoint#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#char_length_semantics AwsDmsEndpoint#char_length_semantics}.
	// Experimental.
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#convert_timestamp_with_zone_to_utc AwsDmsEndpoint#convert_timestamp_with_zone_to_utc}.
	// Experimental.
	ConvertTimestampWithZoneToUtc interface{} `field:"optional" json:"convertTimestampWithZoneToUtc" yaml:"convertTimestampWithZoneToUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_no_log AwsDmsEndpoint#direct_path_no_log}.
	// Experimental.
	DirectPathNoLog interface{} `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_parallel_load AwsDmsEndpoint#direct_path_parallel_load}.
	// Experimental.
	DirectPathParallelLoad interface{} `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#enable_homogenous_tablespace AwsDmsEndpoint#enable_homogenous_tablespace}.
	// Experimental.
	EnableHomogenousTablespace interface{} `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extra_archived_log_dest_ids AwsDmsEndpoint#extra_archived_log_dest_ids}.
	// Experimental.
	ExtraArchivedLogDestIds *[]*float64 `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#fail_task_on_lob_truncation AwsDmsEndpoint#fail_task_on_lob_truncation}.
	// Experimental.
	FailTaskOnLobTruncation interface{} `field:"optional" json:"failTaskOnLobTruncation" yaml:"failTaskOnLobTruncation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#number_datatype_scale AwsDmsEndpoint#number_datatype_scale}.
	// Experimental.
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#open_transaction_window AwsDmsEndpoint#open_transaction_window}.
	// Experimental.
	OpenTransactionWindow *float64 `field:"optional" json:"openTransactionWindow" yaml:"openTransactionWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#oracle_path_prefix AwsDmsEndpoint#oracle_path_prefix}.
	// Experimental.
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#parallel_asm_read_threads AwsDmsEndpoint#parallel_asm_read_threads}.
	// Experimental.
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_ahead_blocks AwsDmsEndpoint#read_ahead_blocks}.
	// Experimental.
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_table_space_name AwsDmsEndpoint#read_table_space_name}.
	// Experimental.
	ReadTableSpaceName interface{} `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#replace_path_prefix AwsDmsEndpoint#replace_path_prefix}.
	// Experimental.
	ReplacePathPrefix interface{} `field:"optional" json:"replacePathPrefix" yaml:"replacePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#retry_interval AwsDmsEndpoint#retry_interval}.
	// Experimental.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_access_role_arn AwsDmsEndpoint#secrets_manager_oracle_asm_access_role_arn}.
	// Experimental.
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_secret_id AwsDmsEndpoint#secrets_manager_oracle_asm_secret_id}.
	// Experimental.
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption AwsDmsEndpoint#security_db_encryption}.
	// Experimental.
	SecurityDbEncryption *string `field:"optional" json:"securityDbEncryption" yaml:"securityDbEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption_name AwsDmsEndpoint#security_db_encryption_name}.
	// Experimental.
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#spatial_data_option_to_geo_json_function_name AwsDmsEndpoint#spatial_data_option_to_geo_json_function_name}.
	// Experimental.
	SpatialDataOptionToGeoJsonFunctionName *string `field:"optional" json:"spatialDataOptionToGeoJsonFunctionName" yaml:"spatialDataOptionToGeoJsonFunctionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#standby_delay_time AwsDmsEndpoint#standby_delay_time}.
	// Experimental.
	StandbyDelayTime *float64 `field:"optional" json:"standbyDelayTime" yaml:"standbyDelayTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#trim_space_in_char AwsDmsEndpoint#trim_space_in_char}.
	// Experimental.
	TrimSpaceInChar interface{} `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_alternate_folder_for_online AwsDmsEndpoint#use_alternate_folder_for_online}.
	// Experimental.
	UseAlternateFolderForOnline interface{} `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_bfile AwsDmsEndpoint#use_bfile}.
	// Experimental.
	UseBfile interface{} `field:"optional" json:"useBfile" yaml:"useBfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_direct_path_full_load AwsDmsEndpoint#use_direct_path_full_load}.
	// Experimental.
	UseDirectPathFullLoad interface{} `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_logminer_reader AwsDmsEndpoint#use_logminer_reader}.
	// Experimental.
	UseLogminerReader interface{} `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_path_prefix AwsDmsEndpoint#use_path_prefix}.
	// Experimental.
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

