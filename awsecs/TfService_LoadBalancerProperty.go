package awsecs


// Experimental.
type TfService_LoadBalancerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#container_name TfService#container_name}.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#container_port TfService#container_port}.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// advanced_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#advanced_configuration TfService#advanced_configuration}
	// Experimental.
	AdvancedConfiguration *TfService_AdvancedConfigurationProperty `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#elb_name TfService#elb_name}.
	// Experimental.
	ElbName *string `field:"optional" json:"elbName" yaml:"elbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#target_group_arn TfService#target_group_arn}.
	// Experimental.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

