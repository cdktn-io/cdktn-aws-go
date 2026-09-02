package awskinesisfirehose


// Experimental.
type TfDeliveryStream_MskSourceConfigurationProperty struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#authentication_configuration TfDeliveryStream#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *TfDeliveryStream_AuthenticationConfigurationProperty `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#msk_cluster_arn TfDeliveryStream#msk_cluster_arn}.
	// Experimental.
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#topic_name TfDeliveryStream#topic_name}.
	// Experimental.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#read_from_timestamp TfDeliveryStream#read_from_timestamp}.
	// Experimental.
	ReadFromTimestamp *string `field:"optional" json:"readFromTimestamp" yaml:"readFromTimestamp"`
}

