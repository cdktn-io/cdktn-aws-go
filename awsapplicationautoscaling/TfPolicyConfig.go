package awsapplicationautoscaling

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#name TfPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#resource_id TfPolicy#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#scalable_dimension TfPolicy#scalable_dimension}.
	// Experimental.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#service_namespace TfPolicy#service_namespace}.
	// Experimental.
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#id TfPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#policy_type TfPolicy#policy_type}.
	// Experimental.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// predictive_scaling_policy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#predictive_scaling_policy_configuration TfPolicy#predictive_scaling_policy_configuration}
	// Experimental.
	PredictiveScalingPolicyConfiguration *TfPolicy_PredictiveScalingPolicyConfigurationProperty `field:"optional" json:"predictiveScalingPolicyConfiguration" yaml:"predictiveScalingPolicyConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#region TfPolicy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// step_scaling_policy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#step_scaling_policy_configuration TfPolicy#step_scaling_policy_configuration}
	// Experimental.
	StepScalingPolicyConfiguration *TfPolicy_StepScalingPolicyConfigurationProperty `field:"optional" json:"stepScalingPolicyConfiguration" yaml:"stepScalingPolicyConfiguration"`
	// target_tracking_scaling_policy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#target_tracking_scaling_policy_configuration TfPolicy#target_tracking_scaling_policy_configuration}
	// Experimental.
	TargetTrackingScalingPolicyConfiguration *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty `field:"optional" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

