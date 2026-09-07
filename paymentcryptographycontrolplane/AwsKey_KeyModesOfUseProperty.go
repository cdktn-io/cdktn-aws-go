package paymentcryptographycontrolplane


// Experimental.
type AwsKey_KeyModesOfUseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#decrypt AwsKey#decrypt}.
	// Experimental.
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#derive_key AwsKey#derive_key}.
	// Experimental.
	DeriveKey interface{} `field:"optional" json:"deriveKey" yaml:"deriveKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#encrypt AwsKey#encrypt}.
	// Experimental.
	Encrypt interface{} `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#generate AwsKey#generate}.
	// Experimental.
	Generate interface{} `field:"optional" json:"generate" yaml:"generate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#no_restrictions AwsKey#no_restrictions}.
	// Experimental.
	NoRestrictions interface{} `field:"optional" json:"noRestrictions" yaml:"noRestrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#sign AwsKey#sign}.
	// Experimental.
	Sign interface{} `field:"optional" json:"sign" yaml:"sign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#unwrap AwsKey#unwrap}.
	// Experimental.
	Unwrap interface{} `field:"optional" json:"unwrap" yaml:"unwrap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#verify AwsKey#verify}.
	// Experimental.
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#wrap AwsKey#wrap}.
	// Experimental.
	Wrap interface{} `field:"optional" json:"wrap" yaml:"wrap"`
}

