package awscloudfront


// Experimental.
type TfMultitenantDistribution_RestrictionsProperty struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#geo_restriction TfMultitenantDistribution#geo_restriction}
	// Experimental.
	GeoRestriction interface{} `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
}

