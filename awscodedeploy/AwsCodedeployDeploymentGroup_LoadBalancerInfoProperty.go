package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty struct {
	// elb_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#elb_info AwsCodedeployDeploymentGroup#elb_info}
	// Experimental.
	ElbInfo interface{} `field:"optional" json:"elbInfo" yaml:"elbInfo"`
	// target_group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#target_group_info AwsCodedeployDeploymentGroup#target_group_info}
	// Experimental.
	TargetGroupInfo interface{} `field:"optional" json:"targetGroupInfo" yaml:"targetGroupInfo"`
	// target_group_pair_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#target_group_pair_info AwsCodedeployDeploymentGroup#target_group_pair_info}
	// Experimental.
	TargetGroupPairInfo *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty `field:"optional" json:"targetGroupPairInfo" yaml:"targetGroupPairInfo"`
}

