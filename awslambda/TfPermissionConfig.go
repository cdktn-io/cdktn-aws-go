package awslambda

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPermissionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#action TfPermission#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#function_name TfPermission#function_name}.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#principal TfPermission#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#event_source_token TfPermission#event_source_token}.
	// Experimental.
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#function_url_auth_type TfPermission#function_url_auth_type}.
	// Experimental.
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#id TfPermission#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#invoked_via_function_url TfPermission#invoked_via_function_url}.
	// Experimental.
	InvokedViaFunctionUrl interface{} `field:"optional" json:"invokedViaFunctionUrl" yaml:"invokedViaFunctionUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#principal_org_id TfPermission#principal_org_id}.
	// Experimental.
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#qualifier TfPermission#qualifier}.
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#region TfPermission#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#source_account TfPermission#source_account}.
	// Experimental.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#source_arn TfPermission#source_arn}.
	// Experimental.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#statement_id TfPermission#statement_id}.
	// Experimental.
	StatementId *string `field:"optional" json:"statementId" yaml:"statementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#statement_id_prefix TfPermission#statement_id_prefix}.
	// Experimental.
	StatementIdPrefix *string `field:"optional" json:"statementIdPrefix" yaml:"statementIdPrefix"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_permission#timeouts TfPermission#timeouts}
	// Experimental.
	Timeouts *TfPermission_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

