package awsgamelift


// Experimental.
type TfGameServerGroup_InstanceDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#instance_type TfGameServerGroup#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#weighted_capacity TfGameServerGroup#weighted_capacity}.
	// Experimental.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

