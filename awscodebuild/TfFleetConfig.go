package awscodebuild

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#base_capacity TfFleet#base_capacity}.
	// Experimental.
	BaseCapacity *float64 `field:"required" json:"baseCapacity" yaml:"baseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#compute_type TfFleet#compute_type}.
	// Experimental.
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#environment_type TfFleet#environment_type}.
	// Experimental.
	EnvironmentType *string `field:"required" json:"environmentType" yaml:"environmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#name TfFleet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// compute_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#compute_configuration TfFleet#compute_configuration}
	// Experimental.
	ComputeConfiguration *TfFleet_ComputeConfigurationProperty `field:"optional" json:"computeConfiguration" yaml:"computeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#fleet_service_role TfFleet#fleet_service_role}.
	// Experimental.
	FleetServiceRole *string `field:"optional" json:"fleetServiceRole" yaml:"fleetServiceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#image_id TfFleet#image_id}.
	// Experimental.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#overflow_behavior TfFleet#overflow_behavior}.
	// Experimental.
	OverflowBehavior *string `field:"optional" json:"overflowBehavior" yaml:"overflowBehavior"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#region TfFleet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#scaling_configuration TfFleet#scaling_configuration}
	// Experimental.
	ScalingConfiguration *TfFleet_ScalingConfigurationProperty `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#tags TfFleet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#tags_all TfFleet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#vpc_config TfFleet#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

