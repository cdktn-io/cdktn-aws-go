package s3glacier


// Experimental.
type AwsVault_NotificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glacier_vault#events AwsVault#events}.
	// Experimental.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glacier_vault#sns_topic AwsVault#sns_topic}.
	// Experimental.
	SnsTopic *string `field:"required" json:"snsTopic" yaml:"snsTopic"`
}

