package dms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3EndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#bucket_name AwsS3Endpoint#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#endpoint_id AwsS3Endpoint#endpoint_id}.
	// Experimental.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#endpoint_type AwsS3Endpoint#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#service_access_role_arn AwsS3Endpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#add_column_name AwsS3Endpoint#add_column_name}.
	// Experimental.
	AddColumnName interface{} `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#add_trailing_padding_character AwsS3Endpoint#add_trailing_padding_character}.
	// Experimental.
	AddTrailingPaddingCharacter interface{} `field:"optional" json:"addTrailingPaddingCharacter" yaml:"addTrailingPaddingCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#bucket_folder AwsS3Endpoint#bucket_folder}.
	// Experimental.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#canned_acl_for_objects AwsS3Endpoint#canned_acl_for_objects}.
	// Experimental.
	CannedAclForObjects *string `field:"optional" json:"cannedAclForObjects" yaml:"cannedAclForObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#cdc_inserts_and_updates AwsS3Endpoint#cdc_inserts_and_updates}.
	// Experimental.
	CdcInsertsAndUpdates interface{} `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#cdc_inserts_only AwsS3Endpoint#cdc_inserts_only}.
	// Experimental.
	CdcInsertsOnly interface{} `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#cdc_max_batch_interval AwsS3Endpoint#cdc_max_batch_interval}.
	// Experimental.
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#cdc_min_file_size AwsS3Endpoint#cdc_min_file_size}.
	// Experimental.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#cdc_path AwsS3Endpoint#cdc_path}.
	// Experimental.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#certificate_arn AwsS3Endpoint#certificate_arn}.
	// Experimental.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#compression_type AwsS3Endpoint#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#csv_delimiter AwsS3Endpoint#csv_delimiter}.
	// Experimental.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#csv_no_sup_value AwsS3Endpoint#csv_no_sup_value}.
	// Experimental.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#csv_null_value AwsS3Endpoint#csv_null_value}.
	// Experimental.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#csv_row_delimiter AwsS3Endpoint#csv_row_delimiter}.
	// Experimental.
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#data_format AwsS3Endpoint#data_format}.
	// Experimental.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#data_page_size AwsS3Endpoint#data_page_size}.
	// Experimental.
	DataPageSize *float64 `field:"optional" json:"dataPageSize" yaml:"dataPageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#date_partition_delimiter AwsS3Endpoint#date_partition_delimiter}.
	// Experimental.
	DatePartitionDelimiter *string `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#date_partition_enabled AwsS3Endpoint#date_partition_enabled}.
	// Experimental.
	DatePartitionEnabled interface{} `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#date_partition_sequence AwsS3Endpoint#date_partition_sequence}.
	// Experimental.
	DatePartitionSequence *string `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#date_partition_timezone AwsS3Endpoint#date_partition_timezone}.
	// Experimental.
	DatePartitionTimezone *string `field:"optional" json:"datePartitionTimezone" yaml:"datePartitionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#detach_target_on_lob_lookup_failure_parquet AwsS3Endpoint#detach_target_on_lob_lookup_failure_parquet}.
	// Experimental.
	DetachTargetOnLobLookupFailureParquet interface{} `field:"optional" json:"detachTargetOnLobLookupFailureParquet" yaml:"detachTargetOnLobLookupFailureParquet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#dict_page_size_limit AwsS3Endpoint#dict_page_size_limit}.
	// Experimental.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#enable_statistics AwsS3Endpoint#enable_statistics}.
	// Experimental.
	EnableStatistics interface{} `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#encoding_type AwsS3Endpoint#encoding_type}.
	// Experimental.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#encryption_mode AwsS3Endpoint#encryption_mode}.
	// Experimental.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#expected_bucket_owner AwsS3Endpoint#expected_bucket_owner}.
	// Experimental.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#external_table_definition AwsS3Endpoint#external_table_definition}.
	// Experimental.
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#glue_catalog_generation AwsS3Endpoint#glue_catalog_generation}.
	// Experimental.
	GlueCatalogGeneration interface{} `field:"optional" json:"glueCatalogGeneration" yaml:"glueCatalogGeneration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#id AwsS3Endpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#ignore_header_rows AwsS3Endpoint#ignore_header_rows}.
	// Experimental.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#include_op_for_full_load AwsS3Endpoint#include_op_for_full_load}.
	// Experimental.
	IncludeOpForFullLoad interface{} `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#kms_key_arn AwsS3Endpoint#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#max_file_size AwsS3Endpoint#max_file_size}.
	// Experimental.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#parquet_timestamp_in_millisecond AwsS3Endpoint#parquet_timestamp_in_millisecond}.
	// Experimental.
	ParquetTimestampInMillisecond interface{} `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#parquet_version AwsS3Endpoint#parquet_version}.
	// Experimental.
	ParquetVersion *string `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#preserve_transactions AwsS3Endpoint#preserve_transactions}.
	// Experimental.
	PreserveTransactions interface{} `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#region AwsS3Endpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#rfc_4180 AwsS3Endpoint#rfc_4180}.
	// Experimental.
	Rfc4180 interface{} `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#row_group_length AwsS3Endpoint#row_group_length}.
	// Experimental.
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#server_side_encryption_kms_key_id AwsS3Endpoint#server_side_encryption_kms_key_id}.
	// Experimental.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#ssl_mode AwsS3Endpoint#ssl_mode}.
	// Experimental.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#tags AwsS3Endpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#tags_all AwsS3Endpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#timeouts AwsS3Endpoint#timeouts}
	// Experimental.
	Timeouts *AwsS3Endpoint_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#timestamp_column_name AwsS3Endpoint#timestamp_column_name}.
	// Experimental.
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#use_csv_no_sup_value AwsS3Endpoint#use_csv_no_sup_value}.
	// Experimental.
	UseCsvNoSupValue interface{} `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint#use_task_start_time_for_full_load_timestamp AwsS3Endpoint#use_task_start_time_for_full_load_timestamp}.
	// Experimental.
	UseTaskStartTimeForFullLoadTimestamp interface{} `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

