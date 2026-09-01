package awscodedeploy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodedeployDeploymentGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#app_name AwsCodedeployDeploymentGroup#app_name}.
	// Experimental.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_group_name AwsCodedeployDeploymentGroup#deployment_group_name}.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#service_role_arn AwsCodedeployDeploymentGroup#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// alarm_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#alarm_configuration AwsCodedeployDeploymentGroup#alarm_configuration}
	// Experimental.
	AlarmConfiguration *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty `field:"optional" json:"alarmConfiguration" yaml:"alarmConfiguration"`
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#auto_rollback_configuration AwsCodedeployDeploymentGroup#auto_rollback_configuration}
	// Experimental.
	AutoRollbackConfiguration *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#autoscaling_groups AwsCodedeployDeploymentGroup#autoscaling_groups}.
	// Experimental.
	AutoscalingGroups *[]*string `field:"optional" json:"autoscalingGroups" yaml:"autoscalingGroups"`
	// blue_green_deployment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#blue_green_deployment_config AwsCodedeployDeploymentGroup#blue_green_deployment_config}
	// Experimental.
	BlueGreenDeploymentConfig *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty `field:"optional" json:"blueGreenDeploymentConfig" yaml:"blueGreenDeploymentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_config_name AwsCodedeployDeploymentGroup#deployment_config_name}.
	// Experimental.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// deployment_style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_style AwsCodedeployDeploymentGroup#deployment_style}
	// Experimental.
	DeploymentStyle *AwsCodedeployDeploymentGroup_DeploymentStyleProperty `field:"optional" json:"deploymentStyle" yaml:"deploymentStyle"`
	// ec2_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#ec2_tag_filter AwsCodedeployDeploymentGroup#ec2_tag_filter}
	// Experimental.
	Ec2TagFilter interface{} `field:"optional" json:"ec2TagFilter" yaml:"ec2TagFilter"`
	// ec2_tag_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#ec2_tag_set AwsCodedeployDeploymentGroup#ec2_tag_set}
	// Experimental.
	Ec2TagSet interface{} `field:"optional" json:"ec2TagSet" yaml:"ec2TagSet"`
	// ecs_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#ecs_service AwsCodedeployDeploymentGroup#ecs_service}
	// Experimental.
	EcsService *AwsCodedeployDeploymentGroup_EcsServiceProperty `field:"optional" json:"ecsService" yaml:"ecsService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#id AwsCodedeployDeploymentGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#load_balancer_info AwsCodedeployDeploymentGroup#load_balancer_info}
	// Experimental.
	LoadBalancerInfo *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty `field:"optional" json:"loadBalancerInfo" yaml:"loadBalancerInfo"`
	// on_premises_instance_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#on_premises_instance_tag_filter AwsCodedeployDeploymentGroup#on_premises_instance_tag_filter}
	// Experimental.
	OnPremisesInstanceTagFilter interface{} `field:"optional" json:"onPremisesInstanceTagFilter" yaml:"onPremisesInstanceTagFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#outdated_instances_strategy AwsCodedeployDeploymentGroup#outdated_instances_strategy}.
	// Experimental.
	OutdatedInstancesStrategy *string `field:"optional" json:"outdatedInstancesStrategy" yaml:"outdatedInstancesStrategy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#region AwsCodedeployDeploymentGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#tags AwsCodedeployDeploymentGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#tags_all AwsCodedeployDeploymentGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#termination_hook_enabled AwsCodedeployDeploymentGroup#termination_hook_enabled}.
	// Experimental.
	TerminationHookEnabled interface{} `field:"optional" json:"terminationHookEnabled" yaml:"terminationHookEnabled"`
	// trigger_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_configuration AwsCodedeployDeploymentGroup#trigger_configuration}
	// Experimental.
	TriggerConfiguration interface{} `field:"optional" json:"triggerConfiguration" yaml:"triggerConfiguration"`
}

