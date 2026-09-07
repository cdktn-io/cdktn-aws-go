package bedrockagentcore


// Experimental.
type AwsHarness_MemoryManagedMemoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#encryption_key_arn AwsHarness#encryption_key_arn}.
	// Experimental.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#event_expiry_duration AwsHarness#event_expiry_duration}.
	// Experimental.
	EventExpiryDuration *float64 `field:"optional" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#strategies AwsHarness#strategies}.
	// Experimental.
	Strategies *[]*string `field:"optional" json:"strategies" yaml:"strategies"`
}

