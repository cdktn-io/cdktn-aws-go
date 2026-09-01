package awsglue


// Experimental.
type AwsGlueCrawler_SchemaChangePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#delete_behavior AwsGlueCrawler#delete_behavior}.
	// Experimental.
	DeleteBehavior *string `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#update_behavior AwsGlueCrawler#update_behavior}.
	// Experimental.
	UpdateBehavior *string `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

