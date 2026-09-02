package awsec2


// Experimental.
type TfAllowedImagesSettings_CreationDateConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_allowed_images_settings#maximum_days_since_created TfAllowedImagesSettings#maximum_days_since_created}.
	// Experimental.
	MaximumDaysSinceCreated *float64 `field:"optional" json:"maximumDaysSinceCreated" yaml:"maximumDaysSinceCreated"`
}

