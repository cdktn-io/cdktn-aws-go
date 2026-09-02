package awsquicksight


// Experimental.
type TfDataSource_RdsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#database TfDataSource#database}.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#instance_id TfDataSource#instance_id}.
	// Experimental.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

