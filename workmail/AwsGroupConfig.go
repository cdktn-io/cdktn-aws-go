package workmail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGroupConfig struct {
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
	// Primary email address used to register the group with WorkMail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_group#email AwsGroup#email}
	// Experimental.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_group#name AwsGroup#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier of the WorkMail organization where the group is managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_group#organization_id AwsGroup#organization_id}
	// Experimental.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Whether to hide the group from the global address list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_group#hidden_from_global_address_list AwsGroup#hidden_from_global_address_list}
	// Experimental.
	HiddenFromGlobalAddressList interface{} `field:"optional" json:"hiddenFromGlobalAddressList" yaml:"hiddenFromGlobalAddressList"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_group#region AwsGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

