package awsapigateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMethodConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#authorization TfMethod#authorization}.
	// Experimental.
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#http_method TfMethod#http_method}.
	// Experimental.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#resource_id TfMethod#resource_id}.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#rest_api_id TfMethod#rest_api_id}.
	// Experimental.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#api_key_required TfMethod#api_key_required}.
	// Experimental.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#authorization_scopes TfMethod#authorization_scopes}.
	// Experimental.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#authorizer_id TfMethod#authorizer_id}.
	// Experimental.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#id TfMethod#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#operation_name TfMethod#operation_name}.
	// Experimental.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#region TfMethod#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#request_models TfMethod#request_models}.
	// Experimental.
	RequestModels *map[string]*string `field:"optional" json:"requestModels" yaml:"requestModels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#request_parameters TfMethod#request_parameters}.
	// Experimental.
	RequestParameters *map[string]interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method#request_validator_id TfMethod#request_validator_id}.
	// Experimental.
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

