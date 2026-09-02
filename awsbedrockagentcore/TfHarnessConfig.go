package awsbedrockagentcore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHarnessConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#execution_role_arn TfHarness#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#harness_name TfHarness#harness_name}.
	// Experimental.
	HarnessName *string `field:"required" json:"harnessName" yaml:"harnessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#allowed_tools TfHarness#allowed_tools}.
	// Experimental.
	AllowedTools *[]*string `field:"optional" json:"allowedTools" yaml:"allowedTools"`
	// authorizer_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#authorizer_configuration TfHarness#authorizer_configuration}
	// Experimental.
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#environment TfHarness#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// environment_artifact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#environment_artifact TfHarness#environment_artifact}
	// Experimental.
	EnvironmentArtifact interface{} `field:"optional" json:"environmentArtifact" yaml:"environmentArtifact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#environment_variables TfHarness#environment_variables}.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#max_iterations TfHarness#max_iterations}.
	// Experimental.
	MaxIterations *float64 `field:"optional" json:"maxIterations" yaml:"maxIterations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#max_tokens TfHarness#max_tokens}.
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// memory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#memory TfHarness#memory}
	// Experimental.
	Memory interface{} `field:"optional" json:"memory" yaml:"memory"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#model TfHarness#model}
	// Experimental.
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#region TfHarness#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// skill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#skill TfHarness#skill}
	// Experimental.
	Skill interface{} `field:"optional" json:"skill" yaml:"skill"`
	// system_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#system_prompt TfHarness#system_prompt}
	// Experimental.
	SystemPrompt interface{} `field:"optional" json:"systemPrompt" yaml:"systemPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#tags TfHarness#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#timeouts TfHarness#timeouts}
	// Experimental.
	Timeouts *TfHarness_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#timeout_seconds TfHarness#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#tool TfHarness#tool}
	// Experimental.
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#truncation TfHarness#truncation}.
	// Experimental.
	Truncation interface{} `field:"optional" json:"truncation" yaml:"truncation"`
}

