package awsssmcontacts


// Experimental.
type TfRotation_ShiftCoveragesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#map_block_key TfRotation#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// coverage_times block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_rotation#coverage_times TfRotation#coverage_times}
	// Experimental.
	CoverageTimes interface{} `field:"optional" json:"coverageTimes" yaml:"coverageTimes"`
}

