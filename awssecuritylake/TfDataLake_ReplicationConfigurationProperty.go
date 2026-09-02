package awssecuritylake


// Experimental.
type TfDataLake_ReplicationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#regions TfDataLake#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#role_arn TfDataLake#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

