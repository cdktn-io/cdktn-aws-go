package awsbedrockagentcore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMemoryStrategyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#memory_id TfMemoryStrategy#memory_id}.
	// Experimental.
	MemoryId *string `field:"required" json:"memoryId" yaml:"memoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#name TfMemoryStrategy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#type TfMemoryStrategy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#configuration TfMemoryStrategy#configuration}
	// Experimental.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#description TfMemoryStrategy#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#memory_execution_role_arn TfMemoryStrategy#memory_execution_role_arn}.
	// Experimental.
	MemoryExecutionRoleArn *string `field:"optional" json:"memoryExecutionRoleArn" yaml:"memoryExecutionRoleArn"`
	// memory_record_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#memory_record_schema TfMemoryStrategy#memory_record_schema}
	// Experimental.
	MemoryRecordSchema interface{} `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#namespaces TfMemoryStrategy#namespaces}.
	// Experimental.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#namespace_templates TfMemoryStrategy#namespace_templates}.
	// Experimental.
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
	// reflection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#reflection_configuration TfMemoryStrategy#reflection_configuration}
	// Experimental.
	ReflectionConfiguration interface{} `field:"optional" json:"reflectionConfiguration" yaml:"reflectionConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#region TfMemoryStrategy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#timeouts TfMemoryStrategy#timeouts}
	// Experimental.
	Timeouts *TfMemoryStrategy_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

