package awsquicksight


// Experimental.
type AwsQuicksightDataSet_RefreshPropertiesProperty struct {
	// refresh_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#refresh_configuration AwsQuicksightDataSet#refresh_configuration}
	// Experimental.
	RefreshConfiguration *AwsQuicksightDataSet_RefreshConfigurationProperty `field:"required" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

