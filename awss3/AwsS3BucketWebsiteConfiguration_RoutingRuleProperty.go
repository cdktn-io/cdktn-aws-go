package awss3


// Experimental.
type AwsS3BucketWebsiteConfiguration_RoutingRuleProperty struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#redirect AwsS3BucketWebsiteConfiguration#redirect}
	// Experimental.
	Redirect *AwsS3BucketWebsiteConfiguration_RedirectProperty `field:"required" json:"redirect" yaml:"redirect"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#condition AwsS3BucketWebsiteConfiguration#condition}
	// Experimental.
	Condition *AwsS3BucketWebsiteConfiguration_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
}

