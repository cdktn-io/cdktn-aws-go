package lambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFunctionScalingConfigConfig struct {
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
	// Name or ARN of the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#function_name AwsFunctionScalingConfig#function_name}
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Qualifier for the scaling configuration. Valid values: $LATEST.PUBLISHED or a numeric version number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#qualifier AwsFunctionScalingConfig#qualifier}
	// Experimental.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// function_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#function_scaling_config AwsFunctionScalingConfig#function_scaling_config}
	// Experimental.
	FunctionScalingConfig interface{} `field:"optional" json:"functionScalingConfig" yaml:"functionScalingConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#region AwsFunctionScalingConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#timeouts AwsFunctionScalingConfig#timeouts}
	// Experimental.
	Timeouts *AwsFunctionScalingConfig_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

