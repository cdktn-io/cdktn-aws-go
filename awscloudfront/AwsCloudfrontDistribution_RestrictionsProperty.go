package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_RestrictionsProperty struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#geo_restriction AwsCloudfrontDistribution#geo_restriction}
	// Experimental.
	GeoRestriction *AwsCloudfrontDistribution_GeoRestrictionProperty `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

