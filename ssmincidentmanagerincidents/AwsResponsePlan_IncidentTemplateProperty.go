package ssmincidentmanagerincidents


// Experimental.
type AwsResponsePlan_IncidentTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#impact AwsResponsePlan#impact}.
	// Experimental.
	Impact *float64 `field:"required" json:"impact" yaml:"impact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#title AwsResponsePlan#title}.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#dedupe_string AwsResponsePlan#dedupe_string}.
	// Experimental.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#incident_tags AwsResponsePlan#incident_tags}.
	// Experimental.
	IncidentTags *map[string]*string `field:"optional" json:"incidentTags" yaml:"incidentTags"`
	// notification_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#notification_target AwsResponsePlan#notification_target}
	// Experimental.
	NotificationTarget interface{} `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#summary AwsResponsePlan#summary}.
	// Experimental.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

