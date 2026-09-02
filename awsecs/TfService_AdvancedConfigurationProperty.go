package awsecs


// Experimental.
type TfService_AdvancedConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#alternate_target_group_arn TfService#alternate_target_group_arn}.
	// Experimental.
	AlternateTargetGroupArn *string `field:"required" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#production_listener_rule TfService#production_listener_rule}.
	// Experimental.
	ProductionListenerRule *string `field:"required" json:"productionListenerRule" yaml:"productionListenerRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn TfService#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#test_listener_rule TfService#test_listener_rule}.
	// Experimental.
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

