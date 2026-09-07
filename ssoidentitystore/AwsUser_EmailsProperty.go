package ssoidentitystore


// Experimental.
type AwsUser_EmailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#primary AwsUser#primary}.
	// Experimental.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#type AwsUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#value AwsUser#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

