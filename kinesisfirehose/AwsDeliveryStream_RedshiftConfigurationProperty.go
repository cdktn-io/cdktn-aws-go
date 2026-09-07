package kinesisfirehose


// Experimental.
type AwsDeliveryStream_RedshiftConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cluster_jdbcurl AwsDeliveryStream#cluster_jdbcurl}.
	// Experimental.
	ClusterJdbcurl *string `field:"required" json:"clusterJdbcurl" yaml:"clusterJdbcurl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#data_table_name AwsDeliveryStream#data_table_name}.
	// Experimental.
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#role_arn AwsDeliveryStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration AwsDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *AwsDeliveryStream_RedshiftConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options AwsDeliveryStream#cloudwatch_logging_options}
	// Experimental.
	CloudwatchLoggingOptions *AwsDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#copy_options AwsDeliveryStream#copy_options}.
	// Experimental.
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#data_table_columns AwsDeliveryStream#data_table_columns}.
	// Experimental.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#password AwsDeliveryStream#password}.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration AwsDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *AwsDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration AwsDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// s3_backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_configuration AwsDeliveryStream#s3_backup_configuration}
	// Experimental.
	S3BackupConfiguration *AwsDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode AwsDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration AwsDeliveryStream#secrets_manager_configuration}
	// Experimental.
	SecretsManagerConfiguration *AwsDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#username AwsDeliveryStream#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

