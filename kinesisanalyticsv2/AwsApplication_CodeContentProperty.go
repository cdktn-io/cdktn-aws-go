package kinesisanalyticsv2


// Experimental.
type AwsApplication_CodeContentProperty struct {
	// s3_content_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#s3_content_location AwsApplication#s3_content_location}
	// Experimental.
	S3ContentLocation *AwsApplication_S3ContentLocationProperty `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#text_content AwsApplication#text_content}.
	// Experimental.
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
}

