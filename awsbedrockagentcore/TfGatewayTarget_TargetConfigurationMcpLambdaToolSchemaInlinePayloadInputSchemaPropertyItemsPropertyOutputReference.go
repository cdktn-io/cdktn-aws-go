package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsItemsPropertyList
	// Experimental.
	ItemsInput() interface{}
	// Experimental.
	Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyPropertyList
	// Experimental.
	PropertyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutItems(value interface{})
	// Experimental.
	PutProperty(value interface{})
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetItems()
	// Experimental.
	ResetProperty()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference
type jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsItemsPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsItemsPropertyList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyPropertyList
	_jsii_.Get(
		j,
		"property",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) PropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference_Override(t TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) PutItems(value interface{}) {
	if err := t.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putItems",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) PutProperty(value interface{}) {
	if err := t.validatePutPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProperty",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		t,
		"resetItems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ResetProperty() {
	_jsii_.InvokeVoid(
		t,
		"resetProperty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

