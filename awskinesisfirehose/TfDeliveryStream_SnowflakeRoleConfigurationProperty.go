package awskinesisfirehose


// Experimental.
type TfDeliveryStream_SnowflakeRoleConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#enabled TfDeliveryStream#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream#snowflake_role TfDeliveryStream#snowflake_role}.
	// Experimental.
	SnowflakeRole *string `field:"optional" json:"snowflakeRole" yaml:"snowflakeRole"`
}

