package kms


// Experimental.
type AwsCustomKeyStore_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#create AwsCustomKeyStore#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#delete AwsCustomKeyStore#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#update AwsCustomKeyStore#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

