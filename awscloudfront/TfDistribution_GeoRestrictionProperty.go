package awscloudfront


// Experimental.
type TfDistribution_GeoRestrictionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#restriction_type TfDistribution#restriction_type}.
	// Experimental.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#locations TfDistribution#locations}.
	// Experimental.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
}

