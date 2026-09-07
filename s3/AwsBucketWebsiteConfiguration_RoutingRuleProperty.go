package s3


// Experimental.
type AwsBucketWebsiteConfiguration_RoutingRuleProperty struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#redirect AwsBucketWebsiteConfiguration#redirect}
	// Experimental.
	Redirect *AwsBucketWebsiteConfiguration_RedirectProperty `field:"required" json:"redirect" yaml:"redirect"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#condition AwsBucketWebsiteConfiguration#condition}
	// Experimental.
	Condition *AwsBucketWebsiteConfiguration_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
}

