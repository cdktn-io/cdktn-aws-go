package awsxray

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResourcePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_resource_policy#policy_document TfResourcePolicy#policy_document}.
	// Experimental.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_resource_policy#policy_name TfResourcePolicy#policy_name}.
	// Experimental.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_resource_policy#bypass_policy_lockout_check TfResourcePolicy#bypass_policy_lockout_check}.
	// Experimental.
	BypassPolicyLockoutCheck interface{} `field:"optional" json:"bypassPolicyLockoutCheck" yaml:"bypassPolicyLockoutCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_resource_policy#policy_revision_id TfResourcePolicy#policy_revision_id}.
	// Experimental.
	PolicyRevisionId *string `field:"optional" json:"policyRevisionId" yaml:"policyRevisionId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/xray_resource_policy#region TfResourcePolicy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

