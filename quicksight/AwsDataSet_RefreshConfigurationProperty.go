package quicksight


// Experimental.
type AwsDataSet_RefreshConfigurationProperty struct {
	// incremental_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#incremental_refresh AwsDataSet#incremental_refresh}
	// Experimental.
	IncrementalRefresh *AwsDataSet_IncrementalRefreshProperty `field:"required" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

