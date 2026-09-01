package awsquicksight


// Experimental.
type AwsQuicksightDataSet_IncrementalRefreshProperty struct {
	// lookback_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#lookback_window AwsQuicksightDataSet#lookback_window}
	// Experimental.
	LookbackWindow *AwsQuicksightDataSet_LookbackWindowProperty `field:"required" json:"lookbackWindow" yaml:"lookbackWindow"`
}

