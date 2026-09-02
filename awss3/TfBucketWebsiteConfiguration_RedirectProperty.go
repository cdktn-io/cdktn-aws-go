package awss3


// Experimental.
type TfBucketWebsiteConfiguration_RedirectProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#host_name TfBucketWebsiteConfiguration#host_name}.
	// Experimental.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#http_redirect_code TfBucketWebsiteConfiguration#http_redirect_code}.
	// Experimental.
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#protocol TfBucketWebsiteConfiguration#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#replace_key_prefix_with TfBucketWebsiteConfiguration#replace_key_prefix_with}.
	// Experimental.
	ReplaceKeyPrefixWith *string `field:"optional" json:"replaceKeyPrefixWith" yaml:"replaceKeyPrefixWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#replace_key_with TfBucketWebsiteConfiguration#replace_key_with}.
	// Experimental.
	ReplaceKeyWith *string `field:"optional" json:"replaceKeyWith" yaml:"replaceKeyWith"`
}

