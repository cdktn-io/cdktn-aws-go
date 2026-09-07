package resiliencehubv2


// Experimental.
type AwsService_AssociatedSystemProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#system_arn AwsService#system_arn}.
	// Experimental.
	SystemArn *string `field:"required" json:"systemArn" yaml:"systemArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#user_journey_ids AwsService#user_journey_ids}.
	// Experimental.
	UserJourneyIds *[]*string `field:"optional" json:"userJourneyIds" yaml:"userJourneyIds"`
}

