package awsbedrock

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockProvisionedModelThroughputConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#model_arn AwsBedrockProvisionedModelThroughput#model_arn}.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#model_units AwsBedrockProvisionedModelThroughput#model_units}.
	// Experimental.
	ModelUnits *float64 `field:"required" json:"modelUnits" yaml:"modelUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#provisioned_model_name AwsBedrockProvisionedModelThroughput#provisioned_model_name}.
	// Experimental.
	ProvisionedModelName *string `field:"required" json:"provisionedModelName" yaml:"provisionedModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#commitment_duration AwsBedrockProvisionedModelThroughput#commitment_duration}.
	// Experimental.
	CommitmentDuration *string `field:"optional" json:"commitmentDuration" yaml:"commitmentDuration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#region AwsBedrockProvisionedModelThroughput#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#tags AwsBedrockProvisionedModelThroughput#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_provisioned_model_throughput#timeouts AwsBedrockProvisionedModelThroughput#timeouts}
	// Experimental.
	Timeouts *AwsBedrockProvisionedModelThroughput_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

