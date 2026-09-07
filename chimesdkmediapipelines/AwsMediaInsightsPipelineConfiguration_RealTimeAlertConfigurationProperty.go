package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_RealTimeAlertConfigurationProperty struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rules AwsMediaInsightsPipelineConfiguration#rules}
	// Experimental.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#disabled AwsMediaInsightsPipelineConfiguration#disabled}.
	// Experimental.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

