package chatbot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTeamsChannelConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#channel_id AwsTeamsChannelConfiguration#channel_id}.
	// Experimental.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#configuration_name AwsTeamsChannelConfiguration#configuration_name}.
	// Experimental.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#iam_role_arn AwsTeamsChannelConfiguration#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#team_id AwsTeamsChannelConfiguration#team_id}.
	// Experimental.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#tenant_id AwsTeamsChannelConfiguration#tenant_id}.
	// Experimental.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#channel_name AwsTeamsChannelConfiguration#channel_name}.
	// Experimental.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#guardrail_policy_arns AwsTeamsChannelConfiguration#guardrail_policy_arns}.
	// Experimental.
	GuardrailPolicyArns *[]*string `field:"optional" json:"guardrailPolicyArns" yaml:"guardrailPolicyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#logging_level AwsTeamsChannelConfiguration#logging_level}.
	// Experimental.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#region AwsTeamsChannelConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#sns_topic_arns AwsTeamsChannelConfiguration#sns_topic_arns}.
	// Experimental.
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#tags AwsTeamsChannelConfiguration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#team_name AwsTeamsChannelConfiguration#team_name}.
	// Experimental.
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#timeouts AwsTeamsChannelConfiguration#timeouts}
	// Experimental.
	Timeouts *AwsTeamsChannelConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_teams_channel_configuration#user_authorization_required AwsTeamsChannelConfiguration#user_authorization_required}.
	// Experimental.
	UserAuthorizationRequired interface{} `field:"optional" json:"userAuthorizationRequired" yaml:"userAuthorizationRequired"`
}

