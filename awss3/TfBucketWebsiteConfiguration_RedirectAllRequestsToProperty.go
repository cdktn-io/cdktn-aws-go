package awss3


// Experimental.
type TfBucketWebsiteConfiguration_RedirectAllRequestsToProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#host_name TfBucketWebsiteConfiguration#host_name}.
	// Experimental.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_website_configuration#protocol TfBucketWebsiteConfiguration#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

