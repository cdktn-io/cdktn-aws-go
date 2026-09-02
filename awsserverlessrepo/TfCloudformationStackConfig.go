package awsserverlessrepo

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudformationStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#application_id TfCloudformationStack#application_id}.
	// Experimental.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#name TfCloudformationStack#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#capabilities TfCloudformationStack#capabilities}.
	// Experimental.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#id TfCloudformationStack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#parameters TfCloudformationStack#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#region TfCloudformationStack#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#semantic_version TfCloudformationStack#semantic_version}.
	// Experimental.
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#tags TfCloudformationStack#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#tags_all TfCloudformationStack#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#timeouts TfCloudformationStack#timeouts}
	// Experimental.
	Timeouts *TfCloudformationStack_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

