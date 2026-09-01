package awscodeartifact

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodeartifactRepositoryPermissionsPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#domain AwsCodeartifactRepositoryPermissionsPolicy#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#policy_document AwsCodeartifactRepositoryPermissionsPolicy#policy_document}.
	// Experimental.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#repository AwsCodeartifactRepositoryPermissionsPolicy#repository}.
	// Experimental.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#domain_owner AwsCodeartifactRepositoryPermissionsPolicy#domain_owner}.
	// Experimental.
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#id AwsCodeartifactRepositoryPermissionsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#policy_revision AwsCodeartifactRepositoryPermissionsPolicy#policy_revision}.
	// Experimental.
	PolicyRevision *string `field:"optional" json:"policyRevision" yaml:"policyRevision"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeartifact_repository_permissions_policy#region AwsCodeartifactRepositoryPermissionsPolicy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

