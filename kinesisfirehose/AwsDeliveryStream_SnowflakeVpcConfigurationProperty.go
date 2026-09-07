package kinesisfirehose


// Experimental.
type AwsDeliveryStream_SnowflakeVpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#private_link_vpce_id AwsDeliveryStream#private_link_vpce_id}.
	// Experimental.
	PrivateLinkVpceId *string `field:"required" json:"privateLinkVpceId" yaml:"privateLinkVpceId"`
}

