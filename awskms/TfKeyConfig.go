package awskms

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#bypass_policy_lockout_safety_check TfKey#bypass_policy_lockout_safety_check}.
	// Experimental.
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#customer_master_key_spec TfKey#customer_master_key_spec}.
	// Experimental.
	CustomerMasterKeySpec *string `field:"optional" json:"customerMasterKeySpec" yaml:"customerMasterKeySpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#custom_key_store_id TfKey#custom_key_store_id}.
	// Experimental.
	CustomKeyStoreId *string `field:"optional" json:"customKeyStoreId" yaml:"customKeyStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#deletion_window_in_days TfKey#deletion_window_in_days}.
	// Experimental.
	DeletionWindowInDays *float64 `field:"optional" json:"deletionWindowInDays" yaml:"deletionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#description TfKey#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#enable_key_rotation TfKey#enable_key_rotation}.
	// Experimental.
	EnableKeyRotation interface{} `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#id TfKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#is_enabled TfKey#is_enabled}.
	// Experimental.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#key_usage TfKey#key_usage}.
	// Experimental.
	KeyUsage *string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#multi_region TfKey#multi_region}.
	// Experimental.
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#policy TfKey#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#region TfKey#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#rotation_period_in_days TfKey#rotation_period_in_days}.
	// Experimental.
	RotationPeriodInDays *float64 `field:"optional" json:"rotationPeriodInDays" yaml:"rotationPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#tags TfKey#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#tags_all TfKey#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#timeouts TfKey#timeouts}
	// Experimental.
	Timeouts *TfKey_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_key#xks_key_id TfKey#xks_key_id}.
	// Experimental.
	XksKeyId *string `field:"optional" json:"xksKeyId" yaml:"xksKeyId"`
}

