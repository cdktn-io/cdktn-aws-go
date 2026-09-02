package awsec2imagebuilder


// Experimental.
type TfDistributionConfiguration_S3ExportConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#disk_image_format TfDistributionConfiguration#disk_image_format}.
	// Experimental.
	DiskImageFormat *string `field:"required" json:"diskImageFormat" yaml:"diskImageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#role_name TfDistributionConfiguration#role_name}.
	// Experimental.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#s3_bucket TfDistributionConfiguration#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#s3_prefix TfDistributionConfiguration#s3_prefix}.
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

