package awsquicksight


// Experimental.
type AwsQuicksightDataSet_RefreshConfigurationProperty struct {
	// incremental_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#incremental_refresh AwsQuicksightDataSet#incremental_refresh}
	// Experimental.
	IncrementalRefresh *AwsQuicksightDataSet_IncrementalRefreshProperty `field:"required" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

