package awsquicksight


// Experimental.
type TfDataSet_RefreshPropertiesProperty struct {
	// refresh_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#refresh_configuration TfDataSet#refresh_configuration}
	// Experimental.
	RefreshConfiguration *TfDataSet_RefreshConfigurationProperty `field:"required" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

