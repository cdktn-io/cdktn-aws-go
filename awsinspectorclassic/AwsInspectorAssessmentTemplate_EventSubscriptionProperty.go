package awsinspectorclassic


// Experimental.
type AwsInspectorAssessmentTemplate_EventSubscriptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector_assessment_template#event AwsInspectorAssessmentTemplate#event}.
	// Experimental.
	Event *string `field:"required" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector_assessment_template#topic_arn AwsInspectorAssessmentTemplate#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

