package cloudfront


// Experimental.
type AwsMultitenantDistribution_GeoRestrictionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#restriction_type AwsMultitenantDistribution#restriction_type}.
	// Experimental.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#items AwsMultitenantDistribution#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

