package ec2


// Experimental.
type AwsAllowedImagesSettings_CreationDateConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#maximum_days_since_created AwsAllowedImagesSettings#maximum_days_since_created}.
	// Experimental.
	MaximumDaysSinceCreated *float64 `field:"optional" json:"maximumDaysSinceCreated" yaml:"maximumDaysSinceCreated"`
}

