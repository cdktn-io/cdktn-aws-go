package awsverifiedaccess


// Experimental.
type TfGroup_SseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_group#customer_managed_key_enabled TfGroup#customer_managed_key_enabled}.
	// Experimental.
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_group#kms_key_arn TfGroup#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

