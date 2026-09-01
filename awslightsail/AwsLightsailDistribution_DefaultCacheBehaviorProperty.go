package awslightsail


// Experimental.
type AwsLightsailDistribution_DefaultCacheBehaviorProperty struct {
	// The cache behavior of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#behavior AwsLightsailDistribution#behavior}
	// Experimental.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
}

