package awsgamelift

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGameServerGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#game_server_group_name TfGameServerGroup#game_server_group_name}.
	// Experimental.
	GameServerGroupName *string `field:"required" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// instance_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#instance_definition TfGameServerGroup#instance_definition}
	// Experimental.
	InstanceDefinition interface{} `field:"required" json:"instanceDefinition" yaml:"instanceDefinition"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#launch_template TfGameServerGroup#launch_template}
	// Experimental.
	LaunchTemplate *TfGameServerGroup_LaunchTemplateProperty `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#max_size TfGameServerGroup#max_size}.
	// Experimental.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#min_size TfGameServerGroup#min_size}.
	// Experimental.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#role_arn TfGameServerGroup#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// auto_scaling_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#auto_scaling_policy TfGameServerGroup#auto_scaling_policy}
	// Experimental.
	AutoScalingPolicy *TfGameServerGroup_AutoScalingPolicyProperty `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#balancing_strategy TfGameServerGroup#balancing_strategy}.
	// Experimental.
	BalancingStrategy *string `field:"optional" json:"balancingStrategy" yaml:"balancingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#game_server_protection_policy TfGameServerGroup#game_server_protection_policy}.
	// Experimental.
	GameServerProtectionPolicy *string `field:"optional" json:"gameServerProtectionPolicy" yaml:"gameServerProtectionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#id TfGameServerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#region TfGameServerGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#tags TfGameServerGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#tags_all TfGameServerGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#timeouts TfGameServerGroup#timeouts}
	// Experimental.
	Timeouts *TfGameServerGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#vpc_subnets TfGameServerGroup#vpc_subnets}.
	// Experimental.
	VpcSubnets *[]*string `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

