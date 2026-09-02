package awsbedrockagentcore


// Experimental.
type TfGateway_AuthorizerConfigurationProperty struct {
	// custom_jwt_authorizer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#custom_jwt_authorizer TfGateway#custom_jwt_authorizer}
	// Experimental.
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

