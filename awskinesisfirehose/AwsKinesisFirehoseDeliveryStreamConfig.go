package awskinesisfirehose

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisFirehoseDeliveryStreamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#destination AwsKinesisFirehoseDeliveryStream#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#name AwsKinesisFirehoseDeliveryStream#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#arn AwsKinesisFirehoseDeliveryStream#arn}.
	// Experimental.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#destination_id AwsKinesisFirehoseDeliveryStream#destination_id}.
	// Experimental.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// elasticsearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#elasticsearch_configuration AwsKinesisFirehoseDeliveryStream#elasticsearch_configuration}
	// Experimental.
	ElasticsearchConfiguration *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty `field:"optional" json:"elasticsearchConfiguration" yaml:"elasticsearchConfiguration"`
	// extended_s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#extended_s3_configuration AwsKinesisFirehoseDeliveryStream#extended_s3_configuration}
	// Experimental.
	ExtendedS3Configuration *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty `field:"optional" json:"extendedS3Configuration" yaml:"extendedS3Configuration"`
	// http_endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#http_endpoint_configuration AwsKinesisFirehoseDeliveryStream#http_endpoint_configuration}
	// Experimental.
	HttpEndpointConfiguration *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty `field:"optional" json:"httpEndpointConfiguration" yaml:"httpEndpointConfiguration"`
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#iceberg_configuration AwsKinesisFirehoseDeliveryStream#iceberg_configuration}
	// Experimental.
	IcebergConfiguration *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#id AwsKinesisFirehoseDeliveryStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kinesis_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#kinesis_source_configuration AwsKinesisFirehoseDeliveryStream#kinesis_source_configuration}
	// Experimental.
	KinesisSourceConfiguration *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty `field:"optional" json:"kinesisSourceConfiguration" yaml:"kinesisSourceConfiguration"`
	// msk_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#msk_source_configuration AwsKinesisFirehoseDeliveryStream#msk_source_configuration}
	// Experimental.
	MskSourceConfiguration *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty `field:"optional" json:"mskSourceConfiguration" yaml:"mskSourceConfiguration"`
	// opensearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#opensearch_configuration AwsKinesisFirehoseDeliveryStream#opensearch_configuration}
	// Experimental.
	OpensearchConfiguration *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
	// opensearchserverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#opensearchserverless_configuration AwsKinesisFirehoseDeliveryStream#opensearchserverless_configuration}
	// Experimental.
	OpensearchserverlessConfiguration *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty `field:"optional" json:"opensearchserverlessConfiguration" yaml:"opensearchserverlessConfiguration"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#redshift_configuration AwsKinesisFirehoseDeliveryStream#redshift_configuration}
	// Experimental.
	RedshiftConfiguration *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#region AwsKinesisFirehoseDeliveryStream#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#server_side_encryption AwsKinesisFirehoseDeliveryStream#server_side_encryption}
	// Experimental.
	ServerSideEncryption *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// snowflake_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#snowflake_configuration AwsKinesisFirehoseDeliveryStream#snowflake_configuration}
	// Experimental.
	SnowflakeConfiguration *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty `field:"optional" json:"snowflakeConfiguration" yaml:"snowflakeConfiguration"`
	// splunk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#splunk_configuration AwsKinesisFirehoseDeliveryStream#splunk_configuration}
	// Experimental.
	SplunkConfiguration *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty `field:"optional" json:"splunkConfiguration" yaml:"splunkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#tags AwsKinesisFirehoseDeliveryStream#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#tags_all AwsKinesisFirehoseDeliveryStream#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#timeouts AwsKinesisFirehoseDeliveryStream#timeouts}
	// Experimental.
	Timeouts *AwsKinesisFirehoseDeliveryStream_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#version_id AwsKinesisFirehoseDeliveryStream#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

