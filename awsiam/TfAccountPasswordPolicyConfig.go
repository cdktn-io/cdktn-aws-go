package awsiam

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAccountPasswordPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#allow_users_to_change_password TfAccountPasswordPolicy#allow_users_to_change_password}.
	// Experimental.
	AllowUsersToChangePassword interface{} `field:"optional" json:"allowUsersToChangePassword" yaml:"allowUsersToChangePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#hard_expiry TfAccountPasswordPolicy#hard_expiry}.
	// Experimental.
	HardExpiry interface{} `field:"optional" json:"hardExpiry" yaml:"hardExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#id TfAccountPasswordPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#max_password_age TfAccountPasswordPolicy#max_password_age}.
	// Experimental.
	MaxPasswordAge *float64 `field:"optional" json:"maxPasswordAge" yaml:"maxPasswordAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#minimum_password_length TfAccountPasswordPolicy#minimum_password_length}.
	// Experimental.
	MinimumPasswordLength *float64 `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#password_reuse_prevention TfAccountPasswordPolicy#password_reuse_prevention}.
	// Experimental.
	PasswordReusePrevention *float64 `field:"optional" json:"passwordReusePrevention" yaml:"passwordReusePrevention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#require_lowercase_characters TfAccountPasswordPolicy#require_lowercase_characters}.
	// Experimental.
	RequireLowercaseCharacters interface{} `field:"optional" json:"requireLowercaseCharacters" yaml:"requireLowercaseCharacters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#require_numbers TfAccountPasswordPolicy#require_numbers}.
	// Experimental.
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#require_symbols TfAccountPasswordPolicy#require_symbols}.
	// Experimental.
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy#require_uppercase_characters TfAccountPasswordPolicy#require_uppercase_characters}.
	// Experimental.
	RequireUppercaseCharacters interface{} `field:"optional" json:"requireUppercaseCharacters" yaml:"requireUppercaseCharacters"`
}

