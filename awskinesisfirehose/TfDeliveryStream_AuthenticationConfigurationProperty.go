package awskinesisfirehose


// Experimental.
type TfDeliveryStream_AuthenticationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#connectivity TfDeliveryStream#connectivity}.
	// Experimental.
	Connectivity *string `field:"required" json:"connectivity" yaml:"connectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#role_arn TfDeliveryStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

