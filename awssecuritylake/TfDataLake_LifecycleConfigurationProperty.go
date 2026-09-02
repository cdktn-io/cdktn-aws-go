package awssecuritylake


// Experimental.
type TfDataLake_LifecycleConfigurationProperty struct {
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#expiration TfDataLake#expiration}
	// Experimental.
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#transition TfDataLake#transition}
	// Experimental.
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

