package awssecuritylake


// Experimental.
type TfDataLake_TransitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#days TfDataLake#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#storage_class TfDataLake#storage_class}.
	// Experimental.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

