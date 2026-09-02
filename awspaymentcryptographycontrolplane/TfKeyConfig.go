package awspaymentcryptographycontrolplane

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKeyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#exportable TfKey#exportable}.
	// Experimental.
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#deletion_window_in_days TfKey#deletion_window_in_days}.
	// Experimental.
	DeletionWindowInDays *float64 `field:"optional" json:"deletionWindowInDays" yaml:"deletionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#enabled TfKey#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// key_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_attributes TfKey#key_attributes}
	// Experimental.
	KeyAttributes interface{} `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#key_check_value_algorithm TfKey#key_check_value_algorithm}.
	// Experimental.
	KeyCheckValueAlgorithm *string `field:"optional" json:"keyCheckValueAlgorithm" yaml:"keyCheckValueAlgorithm"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#region TfKey#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#tags TfKey#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/paymentcryptography_key#timeouts TfKey#timeouts}
	// Experimental.
	Timeouts *TfKey_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

