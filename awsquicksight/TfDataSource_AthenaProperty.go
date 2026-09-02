package awsquicksight


// Experimental.
type TfDataSource_AthenaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#role_arn TfDataSource#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#work_group TfDataSource#work_group}.
	// Experimental.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

