package awscloudfront


// Experimental.
type AwsCloudfrontDistributionTenant_GeoRestrictionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#locations AwsCloudfrontDistributionTenant#locations}.
	// Experimental.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#restriction_type AwsCloudfrontDistributionTenant#restriction_type}.
	// Experimental.
	RestrictionType *string `field:"optional" json:"restrictionType" yaml:"restrictionType"`
}

