package cloudfront


// Experimental.
type AwsDistribution_GeoRestrictionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#restriction_type AwsDistribution#restriction_type}.
	// Experimental.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#locations AwsDistribution#locations}.
	// Experimental.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
}

