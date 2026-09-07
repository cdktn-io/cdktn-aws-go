package bcmdataexports


// Experimental.
type AwsExport_S3DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#s3_bucket AwsExport#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#s3_prefix AwsExport#s3_prefix}.
	// Experimental.
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#s3_region AwsExport#s3_region}.
	// Experimental.
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// s3_output_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#s3_output_configurations AwsExport#s3_output_configurations}
	// Experimental.
	S3OutputConfigurations interface{} `field:"optional" json:"s3OutputConfigurations" yaml:"s3OutputConfigurations"`
}

