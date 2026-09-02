package awscloudfront


// Experimental.
type TfContinuousDeploymentPolicy_TrafficConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#type TfContinuousDeploymentPolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// single_header_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#single_header_config TfContinuousDeploymentPolicy#single_header_config}
	// Experimental.
	SingleHeaderConfig interface{} `field:"optional" json:"singleHeaderConfig" yaml:"singleHeaderConfig"`
	// single_weight_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#single_weight_config TfContinuousDeploymentPolicy#single_weight_config}
	// Experimental.
	SingleWeightConfig interface{} `field:"optional" json:"singleWeightConfig" yaml:"singleWeightConfig"`
}

