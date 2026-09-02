package awsorganizations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#email TfAccount#email}.
	// Experimental.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#name TfAccount#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#close_on_deletion TfAccount#close_on_deletion}.
	// Experimental.
	CloseOnDeletion interface{} `field:"optional" json:"closeOnDeletion" yaml:"closeOnDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#create_govcloud TfAccount#create_govcloud}.
	// Experimental.
	CreateGovcloud interface{} `field:"optional" json:"createGovcloud" yaml:"createGovcloud"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#iam_user_access_to_billing TfAccount#iam_user_access_to_billing}.
	// Experimental.
	IamUserAccessToBilling *string `field:"optional" json:"iamUserAccessToBilling" yaml:"iamUserAccessToBilling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#id TfAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#parent_id TfAccount#parent_id}.
	// Experimental.
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#role_name TfAccount#role_name}.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#tags TfAccount#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#tags_all TfAccount#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/organizations_account#timeouts TfAccount#timeouts}
	// Experimental.
	Timeouts *TfAccount_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

