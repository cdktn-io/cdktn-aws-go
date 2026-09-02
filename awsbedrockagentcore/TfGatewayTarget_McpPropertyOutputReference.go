package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayTarget_McpPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiGateway() TfGatewayTarget_ApiGatewayPropertyList
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
	Lambda() TfGatewayTarget_LambdaPropertyList
	// Experimental.
	LambdaInput() interface{}
	// Experimental.
	McpServer() TfGatewayTarget_McpServerPropertyList
	// Experimental.
	McpServerInput() interface{}
	// Experimental.
	OpenApiSchema() TfGatewayTarget_OpenApiSchemaPropertyList
	// Experimental.
	OpenApiSchemaInput() interface{}
	// Experimental.
	SmithyModel() TfGatewayTarget_SmithyModelPropertyList
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

// The jsii proxy struct for TfGatewayTarget_McpPropertyOutputReference
type jsiiProxy_TfGatewayTarget_McpPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ApiGateway() TfGatewayTarget_ApiGatewayPropertyList {
	var returns TfGatewayTarget_ApiGatewayPropertyList
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ApiGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) Lambda() TfGatewayTarget_LambdaPropertyList {
	var returns TfGatewayTarget_LambdaPropertyList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) McpServer() TfGatewayTarget_McpServerPropertyList {
	var returns TfGatewayTarget_McpServerPropertyList
	_jsii_.Get(
		j,
		"mcpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) McpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) OpenApiSchema() TfGatewayTarget_OpenApiSchemaPropertyList {
	var returns TfGatewayTarget_OpenApiSchemaPropertyList
	_jsii_.Get(
		j,
		"openApiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) OpenApiSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openApiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) SmithyModel() TfGatewayTarget_SmithyModelPropertyList {
	var returns TfGatewayTarget_SmithyModelPropertyList
	_jsii_.Get(
		j,
		"smithyModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) SmithyModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smithyModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayTarget_McpPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGatewayTarget_McpPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayTarget_McpPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayTarget_McpPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayTarget_McpPropertyOutputReference_Override(t TfGatewayTarget_McpPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) PutApiGateway(value interface{}) {
	if err := t.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) PutLambda(value interface{}) {
	if err := t.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambda",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) PutMcpServer(value interface{}) {
	if err := t.validatePutMcpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMcpServer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) PutOpenApiSchema(value interface{}) {
	if err := t.validatePutOpenApiSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpenApiSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) PutSmithyModel(value interface{}) {
	if err := t.validatePutSmithyModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSmithyModel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ResetApiGateway() {
	_jsii_.InvokeVoid(
		t,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ResetMcpServer() {
	_jsii_.InvokeVoid(
		t,
		"resetMcpServer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ResetOpenApiSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenApiSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ResetSmithyModel() {
	_jsii_.InvokeVoid(
		t,
		"resetSmithyModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_McpPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

