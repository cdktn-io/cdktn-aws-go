package awscloudfront


// Experimental.
type AwsCloudfrontDistributionTenant_CustomizationsProperty struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#certificate AwsCloudfrontDistributionTenant#certificate}
	// Experimental.
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#geo_restriction AwsCloudfrontDistributionTenant#geo_restriction}
	// Experimental.
	GeoRestriction interface{} `field:"optional" json:"geoRestriction" yaml:"geoRestriction"`
	// web_acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#web_acl AwsCloudfrontDistributionTenant#web_acl}
	// Experimental.
	WebAcl interface{} `field:"optional" json:"webAcl" yaml:"webAcl"`
}

