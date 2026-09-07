package costoptimizationhub

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPreferencesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/costoptimizationhub_preferences#member_account_discount_visibility AwsPreferences#member_account_discount_visibility}.
	// Experimental.
	MemberAccountDiscountVisibility *string `field:"optional" json:"memberAccountDiscountVisibility" yaml:"memberAccountDiscountVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/costoptimizationhub_preferences#savings_estimation_mode AwsPreferences#savings_estimation_mode}.
	// Experimental.
	SavingsEstimationMode *string `field:"optional" json:"savingsEstimationMode" yaml:"savingsEstimationMode"`
}

