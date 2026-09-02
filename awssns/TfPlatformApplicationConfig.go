package awssns

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlatformApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#name TfPlatformApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#platform TfPlatformApplication#platform}.
	// Experimental.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#platform_credential TfPlatformApplication#platform_credential}.
	// Experimental.
	PlatformCredential *string `field:"required" json:"platformCredential" yaml:"platformCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#apple_platform_bundle_id TfPlatformApplication#apple_platform_bundle_id}.
	// Experimental.
	ApplePlatformBundleId *string `field:"optional" json:"applePlatformBundleId" yaml:"applePlatformBundleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#apple_platform_team_id TfPlatformApplication#apple_platform_team_id}.
	// Experimental.
	ApplePlatformTeamId *string `field:"optional" json:"applePlatformTeamId" yaml:"applePlatformTeamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#event_delivery_failure_topic_arn TfPlatformApplication#event_delivery_failure_topic_arn}.
	// Experimental.
	EventDeliveryFailureTopicArn *string `field:"optional" json:"eventDeliveryFailureTopicArn" yaml:"eventDeliveryFailureTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#event_endpoint_created_topic_arn TfPlatformApplication#event_endpoint_created_topic_arn}.
	// Experimental.
	EventEndpointCreatedTopicArn *string `field:"optional" json:"eventEndpointCreatedTopicArn" yaml:"eventEndpointCreatedTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#event_endpoint_deleted_topic_arn TfPlatformApplication#event_endpoint_deleted_topic_arn}.
	// Experimental.
	EventEndpointDeletedTopicArn *string `field:"optional" json:"eventEndpointDeletedTopicArn" yaml:"eventEndpointDeletedTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#event_endpoint_updated_topic_arn TfPlatformApplication#event_endpoint_updated_topic_arn}.
	// Experimental.
	EventEndpointUpdatedTopicArn *string `field:"optional" json:"eventEndpointUpdatedTopicArn" yaml:"eventEndpointUpdatedTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#failure_feedback_role_arn TfPlatformApplication#failure_feedback_role_arn}.
	// Experimental.
	FailureFeedbackRoleArn *string `field:"optional" json:"failureFeedbackRoleArn" yaml:"failureFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#id TfPlatformApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#platform_principal TfPlatformApplication#platform_principal}.
	// Experimental.
	PlatformPrincipal *string `field:"optional" json:"platformPrincipal" yaml:"platformPrincipal"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#region TfPlatformApplication#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#success_feedback_role_arn TfPlatformApplication#success_feedback_role_arn}.
	// Experimental.
	SuccessFeedbackRoleArn *string `field:"optional" json:"successFeedbackRoleArn" yaml:"successFeedbackRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application#success_feedback_sample_rate TfPlatformApplication#success_feedback_sample_rate}.
	// Experimental.
	SuccessFeedbackSampleRate *string `field:"optional" json:"successFeedbackSampleRate" yaml:"successFeedbackSampleRate"`
}

