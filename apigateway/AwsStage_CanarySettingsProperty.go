package apigateway


// Experimental.
type AwsStage_CanarySettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#deployment_id AwsStage#deployment_id}.
	// Experimental.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#percent_traffic AwsStage#percent_traffic}.
	// Experimental.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#stage_variable_overrides AwsStage#stage_variable_overrides}.
	// Experimental.
	StageVariableOverrides *map[string]*string `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#use_stage_cache AwsStage#use_stage_cache}.
	// Experimental.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

