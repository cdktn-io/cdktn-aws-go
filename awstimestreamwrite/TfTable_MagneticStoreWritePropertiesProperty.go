package awstimestreamwrite


// Experimental.
type TfTable_MagneticStoreWritePropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#enable_magnetic_store_writes TfTable#enable_magnetic_store_writes}.
	// Experimental.
	EnableMagneticStoreWrites interface{} `field:"optional" json:"enableMagneticStoreWrites" yaml:"enableMagneticStoreWrites"`
	// magnetic_store_rejected_data_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#magnetic_store_rejected_data_location TfTable#magnetic_store_rejected_data_location}
	// Experimental.
	MagneticStoreRejectedDataLocation *TfTable_MagneticStoreRejectedDataLocationProperty `field:"optional" json:"magneticStoreRejectedDataLocation" yaml:"magneticStoreRejectedDataLocation"`
}

