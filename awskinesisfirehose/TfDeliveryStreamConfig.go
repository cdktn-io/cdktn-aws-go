package awskinesisfirehose

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStreamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#destination TfDeliveryStream#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#name TfDeliveryStream#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#arn TfDeliveryStream#arn}.
	// Experimental.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#destination_id TfDeliveryStream#destination_id}.
	// Experimental.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// elasticsearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#elasticsearch_configuration TfDeliveryStream#elasticsearch_configuration}
	// Experimental.
	ElasticsearchConfiguration *TfDeliveryStream_ElasticsearchConfigurationProperty `field:"optional" json:"elasticsearchConfiguration" yaml:"elasticsearchConfiguration"`
	// extended_s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#extended_s3_configuration TfDeliveryStream#extended_s3_configuration}
	// Experimental.
	ExtendedS3Configuration *TfDeliveryStream_ExtendedS3ConfigurationProperty `field:"optional" json:"extendedS3Configuration" yaml:"extendedS3Configuration"`
	// http_endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#http_endpoint_configuration TfDeliveryStream#http_endpoint_configuration}
	// Experimental.
	HttpEndpointConfiguration *TfDeliveryStream_HttpEndpointConfigurationProperty `field:"optional" json:"httpEndpointConfiguration" yaml:"httpEndpointConfiguration"`
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#iceberg_configuration TfDeliveryStream#iceberg_configuration}
	// Experimental.
	IcebergConfiguration *TfDeliveryStream_IcebergConfigurationProperty `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#id TfDeliveryStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kinesis_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#kinesis_source_configuration TfDeliveryStream#kinesis_source_configuration}
	// Experimental.
	KinesisSourceConfiguration *TfDeliveryStream_KinesisSourceConfigurationProperty `field:"optional" json:"kinesisSourceConfiguration" yaml:"kinesisSourceConfiguration"`
	// msk_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#msk_source_configuration TfDeliveryStream#msk_source_configuration}
	// Experimental.
	MskSourceConfiguration *TfDeliveryStream_MskSourceConfigurationProperty `field:"optional" json:"mskSourceConfiguration" yaml:"mskSourceConfiguration"`
	// opensearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#opensearch_configuration TfDeliveryStream#opensearch_configuration}
	// Experimental.
	OpensearchConfiguration *TfDeliveryStream_OpensearchConfigurationProperty `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
	// opensearchserverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#opensearchserverless_configuration TfDeliveryStream#opensearchserverless_configuration}
	// Experimental.
	OpensearchserverlessConfiguration *TfDeliveryStream_OpensearchserverlessConfigurationProperty `field:"optional" json:"opensearchserverlessConfiguration" yaml:"opensearchserverlessConfiguration"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#redshift_configuration TfDeliveryStream#redshift_configuration}
	// Experimental.
	RedshiftConfiguration *TfDeliveryStream_RedshiftConfigurationProperty `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#region TfDeliveryStream#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#server_side_encryption TfDeliveryStream#server_side_encryption}
	// Experimental.
	ServerSideEncryption *TfDeliveryStream_ServerSideEncryptionProperty `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// snowflake_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#snowflake_configuration TfDeliveryStream#snowflake_configuration}
	// Experimental.
	SnowflakeConfiguration *TfDeliveryStream_SnowflakeConfigurationProperty `field:"optional" json:"snowflakeConfiguration" yaml:"snowflakeConfiguration"`
	// splunk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#splunk_configuration TfDeliveryStream#splunk_configuration}
	// Experimental.
	SplunkConfiguration *TfDeliveryStream_SplunkConfigurationProperty `field:"optional" json:"splunkConfiguration" yaml:"splunkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#tags TfDeliveryStream#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#tags_all TfDeliveryStream#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#timeouts TfDeliveryStream#timeouts}
	// Experimental.
	Timeouts *TfDeliveryStream_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#version_id TfDeliveryStream#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

