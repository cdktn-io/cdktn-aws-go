package awslambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type EphemeralTfInvocationConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#function_name EphemeralTfInvocation#function_name}.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#payload EphemeralTfInvocation#payload}.
	// Experimental.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#client_context EphemeralTfInvocation#client_context}.
	// Experimental.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#log_type EphemeralTfInvocation#log_type}.
	// Experimental.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#qualifier EphemeralTfInvocation#qualifier}.
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/lambda_invocation#region EphemeralTfInvocation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

