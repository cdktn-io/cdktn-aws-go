package awsdlm


// Experimental.
type TfLifecyclePolicy_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cmk_arn TfLifecyclePolicy#cmk_arn}.
	// Experimental.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#encrypted TfLifecyclePolicy#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
}

