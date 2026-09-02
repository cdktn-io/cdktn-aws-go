package awss3


// Experimental.
type TfBucketWebsiteConfiguration_RoutingRuleProperty struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#redirect TfBucketWebsiteConfiguration#redirect}
	// Experimental.
	Redirect *TfBucketWebsiteConfiguration_RedirectProperty `field:"required" json:"redirect" yaml:"redirect"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#condition TfBucketWebsiteConfiguration#condition}
	// Experimental.
	Condition *TfBucketWebsiteConfiguration_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
}

