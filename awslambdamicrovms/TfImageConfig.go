package awslambdamicrovms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#base_image_arn TfImage#base_image_arn}.
	// Experimental.
	BaseImageArn *string `field:"required" json:"baseImageArn" yaml:"baseImageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#build_role_arn TfImage#build_role_arn}.
	// Experimental.
	BuildRoleArn *string `field:"required" json:"buildRoleArn" yaml:"buildRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#name TfImage#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#additional_os_capabilities TfImage#additional_os_capabilities}.
	// Experimental.
	AdditionalOsCapabilities *[]*string `field:"optional" json:"additionalOsCapabilities" yaml:"additionalOsCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#base_image_version TfImage#base_image_version}.
	// Experimental.
	BaseImageVersion *string `field:"optional" json:"baseImageVersion" yaml:"baseImageVersion"`
	// code_artifact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#code_artifact TfImage#code_artifact}
	// Experimental.
	CodeArtifact interface{} `field:"optional" json:"codeArtifact" yaml:"codeArtifact"`
	// cpu_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#cpu_configuration TfImage#cpu_configuration}
	// Experimental.
	CpuConfiguration interface{} `field:"optional" json:"cpuConfiguration" yaml:"cpuConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#description TfImage#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#egress_network_connectors TfImage#egress_network_connectors}.
	// Experimental.
	EgressNetworkConnectors *[]*string `field:"optional" json:"egressNetworkConnectors" yaml:"egressNetworkConnectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#environment_variables TfImage#environment_variables}.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#region TfImage#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#tags TfImage#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambdamicrovms_image#timeouts TfImage#timeouts}
	// Experimental.
	Timeouts *TfImage_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

