package awspaymentcryptographycontrolplane


// Experimental.
type TfKey_KeyAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_algorithm TfKey#key_algorithm}.
	// Experimental.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_class TfKey#key_class}.
	// Experimental.
	KeyClass *string `field:"required" json:"keyClass" yaml:"keyClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_usage TfKey#key_usage}.
	// Experimental.
	KeyUsage *string `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// key_modes_of_use block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_modes_of_use TfKey#key_modes_of_use}
	// Experimental.
	KeyModesOfUse interface{} `field:"optional" json:"keyModesOfUse" yaml:"keyModesOfUse"`
}

