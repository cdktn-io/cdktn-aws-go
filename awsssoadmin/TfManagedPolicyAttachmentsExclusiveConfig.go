package awsssoadmin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfManagedPolicyAttachmentsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_managed_policy_attachments_exclusive#instance_arn TfManagedPolicyAttachmentsExclusive#instance_arn}.
	// Experimental.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_managed_policy_attachments_exclusive#managed_policy_arns TfManagedPolicyAttachmentsExclusive#managed_policy_arns}.
	// Experimental.
	ManagedPolicyArns *[]*string `field:"required" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_managed_policy_attachments_exclusive#permission_set_arn TfManagedPolicyAttachmentsExclusive#permission_set_arn}.
	// Experimental.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_managed_policy_attachments_exclusive#region TfManagedPolicyAttachmentsExclusive#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_managed_policy_attachments_exclusive#timeouts TfManagedPolicyAttachmentsExclusive#timeouts}
	// Experimental.
	Timeouts *TfManagedPolicyAttachmentsExclusive_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

