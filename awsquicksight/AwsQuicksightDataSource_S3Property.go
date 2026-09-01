package awsquicksight


// Experimental.
type AwsQuicksightDataSource_S3Property struct {
	// manifest_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#manifest_file_location AwsQuicksightDataSource#manifest_file_location}
	// Experimental.
	ManifestFileLocation *AwsQuicksightDataSource_ManifestFileLocationProperty `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#role_arn AwsQuicksightDataSource#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

