package awsquicksight


// Experimental.
type TfDataSet_RefreshConfigurationProperty struct {
	// incremental_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#incremental_refresh TfDataSet#incremental_refresh}
	// Experimental.
	IncrementalRefresh *TfDataSet_IncrementalRefreshProperty `field:"required" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

