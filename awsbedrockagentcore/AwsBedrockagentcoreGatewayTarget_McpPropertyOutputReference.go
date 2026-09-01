package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiGateway() AwsBedrockagentcoreGatewayTarget_ApiGatewayPropertyList
	// Experimental.
	ApiGatewayInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Lambda() AwsBedrockagentcoreGatewayTarget_LambdaPropertyList
	// Experimental.
	LambdaInput() interface{}
	// Experimental.
	McpServer() AwsBedrockagentcoreGatewayTarget_McpServerPropertyList
	// Experimental.
	McpServerInput() interface{}
	// Experimental.
	OpenApiSchema() AwsBedrockagentcoreGatewayTarget_OpenApiSchemaPropertyList
	// Experimental.
	OpenApiSchemaInput() interface{}
	// Experimental.
	SmithyModel() AwsBedrockagentcoreGatewayTarget_SmithyModelPropertyList
	// Experimental.
	SmithyModelInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutApiGateway(value interface{})
	// Experimental.
	PutLambda(value interface{})
	// Experimental.
	PutMcpServer(value interface{})
	// Experimental.
	PutOpenApiSchema(value interface{})
	// Experimental.
	PutSmithyModel(value interface{})
	// Experimental.
	ResetApiGateway()
	// Experimental.
	ResetLambda()
	// Experimental.
	ResetMcpServer()
	// Experimental.
	ResetOpenApiSchema()
	// Experimental.
	ResetSmithyModel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ApiGateway() AwsBedrockagentcoreGatewayTarget_ApiGatewayPropertyList {
	var returns AwsBedrockagentcoreGatewayTarget_ApiGatewayPropertyList
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ApiGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) Lambda() AwsBedrockagentcoreGatewayTarget_LambdaPropertyList {
	var returns AwsBedrockagentcoreGatewayTarget_LambdaPropertyList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) McpServer() AwsBedrockagentcoreGatewayTarget_McpServerPropertyList {
	var returns AwsBedrockagentcoreGatewayTarget_McpServerPropertyList
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) McpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) OpenApiSchema() AwsBedrockagentcoreGatewayTarget_OpenApiSchemaPropertyList {
	var returns AwsBedrockagentcoreGatewayTarget_OpenApiSchemaPropertyList
	_jsii_.Get(
		j,
		"openApiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) OpenApiSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openApiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) SmithyModel() AwsBedrockagentcoreGatewayTarget_SmithyModelPropertyList {
	var returns AwsBedrockagentcoreGatewayTarget_SmithyModelPropertyList
	_jsii_.Get(
		j,
		"smithyModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) SmithyModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smithyModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreGatewayTarget_McpPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreGatewayTarget.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference_Override(a AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreGatewayTarget.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) PutApiGateway(value interface{}) {
	if err := a.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) PutLambda(value interface{}) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) PutMcpServer(value interface{}) {
	if err := a.validatePutMcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMcpServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) PutOpenApiSchema(value interface{}) {
	if err := a.validatePutOpenApiSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenApiSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) PutSmithyModel(value interface{}) {
	if err := a.validatePutSmithyModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmithyModel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ResetApiGateway() {
	_jsii_.InvokeVoid(
		a,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ResetMcpServer() {
	_jsii_.InvokeVoid(
		a,
		"resetMcpServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ResetOpenApiSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenApiSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ResetSmithyModel() {
	_jsii_.InvokeVoid(
		a,
		"resetSmithyModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreGatewayTarget_McpPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

