package awsapigateway


// Experimental.
type TfUsagePlan_ApiStagesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#api_id TfUsagePlan#api_id}.
	// Experimental.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#stage TfUsagePlan#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// throttle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#throttle TfUsagePlan#throttle}
	// Experimental.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

