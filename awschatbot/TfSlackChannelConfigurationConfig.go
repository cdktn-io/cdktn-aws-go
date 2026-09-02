package awschatbot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlackChannelConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#configuration_name TfSlackChannelConfiguration#configuration_name}.
	// Experimental.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#iam_role_arn TfSlackChannelConfiguration#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#slack_channel_id TfSlackChannelConfiguration#slack_channel_id}.
	// Experimental.
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#slack_team_id TfSlackChannelConfiguration#slack_team_id}.
	// Experimental.
	SlackTeamId *string `field:"required" json:"slackTeamId" yaml:"slackTeamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#guardrail_policy_arns TfSlackChannelConfiguration#guardrail_policy_arns}.
	// Experimental.
	GuardrailPolicyArns *[]*string `field:"optional" json:"guardrailPolicyArns" yaml:"guardrailPolicyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#logging_level TfSlackChannelConfiguration#logging_level}.
	// Experimental.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#region TfSlackChannelConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#sns_topic_arns TfSlackChannelConfiguration#sns_topic_arns}.
	// Experimental.
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#tags TfSlackChannelConfiguration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#timeouts TfSlackChannelConfiguration#timeouts}
	// Experimental.
	Timeouts *TfSlackChannelConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chatbot_slack_channel_configuration#user_authorization_required TfSlackChannelConfiguration#user_authorization_required}.
	// Experimental.
	UserAuthorizationRequired interface{} `field:"optional" json:"userAuthorizationRequired" yaml:"userAuthorizationRequired"`
}

