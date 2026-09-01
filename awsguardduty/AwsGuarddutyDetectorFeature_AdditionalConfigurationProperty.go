package awsguardduty


// Experimental.
type AwsGuarddutyDetectorFeature_AdditionalConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector_feature#name AwsGuarddutyDetectorFeature#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_detector_feature#status AwsGuarddutyDetectorFeature#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
}

