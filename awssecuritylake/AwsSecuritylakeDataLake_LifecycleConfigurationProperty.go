package awssecuritylake


// Experimental.
type AwsSecuritylakeDataLake_LifecycleConfigurationProperty struct {
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#expiration AwsSecuritylakeDataLake#expiration}
	// Experimental.
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#transition AwsSecuritylakeDataLake#transition}
	// Experimental.
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

