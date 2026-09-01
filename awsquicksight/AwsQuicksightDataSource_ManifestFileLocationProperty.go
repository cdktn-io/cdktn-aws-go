package awsquicksight


// Experimental.
type AwsQuicksightDataSource_ManifestFileLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#bucket AwsQuicksightDataSource#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#key AwsQuicksightDataSource#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

