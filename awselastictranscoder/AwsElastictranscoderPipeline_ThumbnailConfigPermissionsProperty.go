package awselastictranscoder


// Experimental.
type AwsElastictranscoderPipeline_ThumbnailConfigPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#access AwsElastictranscoderPipeline#access}.
	// Experimental.
	Access *[]*string `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#grantee AwsElastictranscoderPipeline#grantee}.
	// Experimental.
	Grantee *string `field:"optional" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#grantee_type AwsElastictranscoderPipeline#grantee_type}.
	// Experimental.
	GranteeType *string `field:"optional" json:"granteeType" yaml:"granteeType"`
}

