package awsecr

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRepositoryCreationTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#applied_for TfRepositoryCreationTemplate#applied_for}.
	// Experimental.
	AppliedFor *[]*string `field:"required" json:"appliedFor" yaml:"appliedFor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#prefix TfRepositoryCreationTemplate#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#custom_role_arn TfRepositoryCreationTemplate#custom_role_arn}.
	// Experimental.
	CustomRoleArn *string `field:"optional" json:"customRoleArn" yaml:"customRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#description TfRepositoryCreationTemplate#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#encryption_configuration TfRepositoryCreationTemplate#encryption_configuration}
	// Experimental.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#id TfRepositoryCreationTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#image_tag_mutability TfRepositoryCreationTemplate#image_tag_mutability}.
	// Experimental.
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// image_tag_mutability_exclusion_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#image_tag_mutability_exclusion_filter TfRepositoryCreationTemplate#image_tag_mutability_exclusion_filter}
	// Experimental.
	ImageTagMutabilityExclusionFilter interface{} `field:"optional" json:"imageTagMutabilityExclusionFilter" yaml:"imageTagMutabilityExclusionFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#lifecycle_policy TfRepositoryCreationTemplate#lifecycle_policy}.
	// Experimental.
	LifecyclePolicy *string `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#region TfRepositoryCreationTemplate#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#repository_policy TfRepositoryCreationTemplate#repository_policy}.
	// Experimental.
	RepositoryPolicy *string `field:"optional" json:"repositoryPolicy" yaml:"repositoryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_repository_creation_template#resource_tags TfRepositoryCreationTemplate#resource_tags}.
	// Experimental.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

