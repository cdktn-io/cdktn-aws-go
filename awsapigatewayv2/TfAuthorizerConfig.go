package awsapigatewayv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAuthorizerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#api_id TfAuthorizer#api_id}.
	// Experimental.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#authorizer_type TfAuthorizer#authorizer_type}.
	// Experimental.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#name TfAuthorizer#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#authorizer_credentials_arn TfAuthorizer#authorizer_credentials_arn}.
	// Experimental.
	AuthorizerCredentialsArn *string `field:"optional" json:"authorizerCredentialsArn" yaml:"authorizerCredentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#authorizer_payload_format_version TfAuthorizer#authorizer_payload_format_version}.
	// Experimental.
	AuthorizerPayloadFormatVersion *string `field:"optional" json:"authorizerPayloadFormatVersion" yaml:"authorizerPayloadFormatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#authorizer_result_ttl_in_seconds TfAuthorizer#authorizer_result_ttl_in_seconds}.
	// Experimental.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#authorizer_uri TfAuthorizer#authorizer_uri}.
	// Experimental.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#enable_simple_responses TfAuthorizer#enable_simple_responses}.
	// Experimental.
	EnableSimpleResponses interface{} `field:"optional" json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#id TfAuthorizer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#identity_sources TfAuthorizer#identity_sources}.
	// Experimental.
	IdentitySources *[]*string `field:"optional" json:"identitySources" yaml:"identitySources"`
	// jwt_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#jwt_configuration TfAuthorizer#jwt_configuration}
	// Experimental.
	JwtConfiguration *TfAuthorizer_JwtConfigurationProperty `field:"optional" json:"jwtConfiguration" yaml:"jwtConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#region TfAuthorizer#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_authorizer#timeouts TfAuthorizer#timeouts}
	// Experimental.
	Timeouts *TfAuthorizer_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

