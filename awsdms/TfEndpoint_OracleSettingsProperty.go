package awsdms


// Experimental.
type TfEndpoint_OracleSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#access_alternate_directly TfEndpoint#access_alternate_directly}.
	// Experimental.
	AccessAlternateDirectly interface{} `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#additional_archived_log_dest_id TfEndpoint#additional_archived_log_dest_id}.
	// Experimental.
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#add_supplemental_logging TfEndpoint#add_supplemental_logging}.
	// Experimental.
	AddSupplementalLogging interface{} `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#allow_selected_nested_tables TfEndpoint#allow_selected_nested_tables}.
	// Experimental.
	AllowSelectedNestedTables interface{} `field:"optional" json:"allowSelectedNestedTables" yaml:"allowSelectedNestedTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_log_dest_id TfEndpoint#archived_log_dest_id}.
	// Experimental.
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_logs_only TfEndpoint#archived_logs_only}.
	// Experimental.
	ArchivedLogsOnly interface{} `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_password TfEndpoint#asm_password}.
	// Experimental.
	AsmPassword *string `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_server TfEndpoint#asm_server}.
	// Experimental.
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_user TfEndpoint#asm_user}.
	// Experimental.
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method TfEndpoint#authentication_method}.
	// Experimental.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#char_length_semantics TfEndpoint#char_length_semantics}.
	// Experimental.
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#convert_timestamp_with_zone_to_utc TfEndpoint#convert_timestamp_with_zone_to_utc}.
	// Experimental.
	ConvertTimestampWithZoneToUtc interface{} `field:"optional" json:"convertTimestampWithZoneToUtc" yaml:"convertTimestampWithZoneToUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_no_log TfEndpoint#direct_path_no_log}.
	// Experimental.
	DirectPathNoLog interface{} `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_parallel_load TfEndpoint#direct_path_parallel_load}.
	// Experimental.
	DirectPathParallelLoad interface{} `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#enable_homogenous_tablespace TfEndpoint#enable_homogenous_tablespace}.
	// Experimental.
	EnableHomogenousTablespace interface{} `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extra_archived_log_dest_ids TfEndpoint#extra_archived_log_dest_ids}.
	// Experimental.
	ExtraArchivedLogDestIds *[]*float64 `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#fail_task_on_lob_truncation TfEndpoint#fail_task_on_lob_truncation}.
	// Experimental.
	FailTaskOnLobTruncation interface{} `field:"optional" json:"failTaskOnLobTruncation" yaml:"failTaskOnLobTruncation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#number_datatype_scale TfEndpoint#number_datatype_scale}.
	// Experimental.
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#open_transaction_window TfEndpoint#open_transaction_window}.
	// Experimental.
	OpenTransactionWindow *float64 `field:"optional" json:"openTransactionWindow" yaml:"openTransactionWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#oracle_path_prefix TfEndpoint#oracle_path_prefix}.
	// Experimental.
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#parallel_asm_read_threads TfEndpoint#parallel_asm_read_threads}.
	// Experimental.
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_ahead_blocks TfEndpoint#read_ahead_blocks}.
	// Experimental.
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_table_space_name TfEndpoint#read_table_space_name}.
	// Experimental.
	ReadTableSpaceName interface{} `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#replace_path_prefix TfEndpoint#replace_path_prefix}.
	// Experimental.
	ReplacePathPrefix interface{} `field:"optional" json:"replacePathPrefix" yaml:"replacePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#retry_interval TfEndpoint#retry_interval}.
	// Experimental.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_access_role_arn TfEndpoint#secrets_manager_oracle_asm_access_role_arn}.
	// Experimental.
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_secret_id TfEndpoint#secrets_manager_oracle_asm_secret_id}.
	// Experimental.
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption TfEndpoint#security_db_encryption}.
	// Experimental.
	SecurityDbEncryption *string `field:"optional" json:"securityDbEncryption" yaml:"securityDbEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption_name TfEndpoint#security_db_encryption_name}.
	// Experimental.
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#spatial_data_option_to_geo_json_function_name TfEndpoint#spatial_data_option_to_geo_json_function_name}.
	// Experimental.
	SpatialDataOptionToGeoJsonFunctionName *string `field:"optional" json:"spatialDataOptionToGeoJsonFunctionName" yaml:"spatialDataOptionToGeoJsonFunctionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#standby_delay_time TfEndpoint#standby_delay_time}.
	// Experimental.
	StandbyDelayTime *float64 `field:"optional" json:"standbyDelayTime" yaml:"standbyDelayTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#trim_space_in_char TfEndpoint#trim_space_in_char}.
	// Experimental.
	TrimSpaceInChar interface{} `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_alternate_folder_for_online TfEndpoint#use_alternate_folder_for_online}.
	// Experimental.
	UseAlternateFolderForOnline interface{} `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_bfile TfEndpoint#use_bfile}.
	// Experimental.
	UseBfile interface{} `field:"optional" json:"useBfile" yaml:"useBfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_direct_path_full_load TfEndpoint#use_direct_path_full_load}.
	// Experimental.
	UseDirectPathFullLoad interface{} `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_logminer_reader TfEndpoint#use_logminer_reader}.
	// Experimental.
	UseLogminerReader interface{} `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_path_prefix TfEndpoint#use_path_prefix}.
	// Experimental.
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

