package awscodecatalyst

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfDevEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#env_id DataTfDevEnvironment#env_id}.
	// Experimental.
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#project_name DataTfDevEnvironment#project_name}.
	// Experimental.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#space_name DataTfDevEnvironment#space_name}.
	// Experimental.
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#alias DataTfDevEnvironment#alias}.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#creator_id DataTfDevEnvironment#creator_id}.
	// Experimental.
	CreatorId *string `field:"optional" json:"creatorId" yaml:"creatorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#id DataTfDevEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#region DataTfDevEnvironment#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// repositories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#repositories DataTfDevEnvironment#repositories}
	// Experimental.
	Repositories interface{} `field:"optional" json:"repositories" yaml:"repositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/codecatalyst_dev_environment#tags DataTfDevEnvironment#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

