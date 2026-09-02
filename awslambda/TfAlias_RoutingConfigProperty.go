package awslambda


// Experimental.
type TfAlias_RoutingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_alias#additional_version_weights TfAlias#additional_version_weights}.
	// Experimental.
	AdditionalVersionWeights *map[string]*float64 `field:"optional" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

