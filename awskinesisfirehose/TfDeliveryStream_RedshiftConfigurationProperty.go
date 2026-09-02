package awskinesisfirehose


// Experimental.
type TfDeliveryStream_RedshiftConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cluster_jdbcurl TfDeliveryStream#cluster_jdbcurl}.
	// Experimental.
	ClusterJdbcurl *string `field:"required" json:"clusterJdbcurl" yaml:"clusterJdbcurl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#data_table_name TfDeliveryStream#data_table_name}.
	// Experimental.
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#role_arn TfDeliveryStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration TfDeliveryStream#s3_configuration}
	// Experimental.
	S3Configuration *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options TfDeliveryStream#cloudwatch_logging_options}
	// Experimental.
	CloudwatchLoggingOptions *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#copy_options TfDeliveryStream#copy_options}.
	// Experimental.
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#data_table_columns TfDeliveryStream#data_table_columns}.
	// Experimental.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#password TfDeliveryStream#password}.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration TfDeliveryStream#processing_configuration}
	// Experimental.
	ProcessingConfiguration *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration TfDeliveryStream#retry_duration}.
	// Experimental.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// s3_backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_configuration TfDeliveryStream#s3_backup_configuration}
	// Experimental.
	S3BackupConfiguration *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode TfDeliveryStream#s3_backup_mode}.
	// Experimental.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration TfDeliveryStream#secrets_manager_configuration}
	// Experimental.
	SecretsManagerConfiguration *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#username TfDeliveryStream#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

