package awss3glacier


// Experimental.
type AwsGlacierVault_NotificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glacier_vault#events AwsGlacierVault#events}.
	// Experimental.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glacier_vault#sns_topic AwsGlacierVault#sns_topic}.
	// Experimental.
	SnsTopic *string `field:"required" json:"snsTopic" yaml:"snsTopic"`
}

