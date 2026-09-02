package awscloudfront


// Experimental.
type TfDistribution_RestrictionsProperty struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#geo_restriction TfDistribution#geo_restriction}
	// Experimental.
	GeoRestriction *TfDistribution_GeoRestrictionProperty `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

