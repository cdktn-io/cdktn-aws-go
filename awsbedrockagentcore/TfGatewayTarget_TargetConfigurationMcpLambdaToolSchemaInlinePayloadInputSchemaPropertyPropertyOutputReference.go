package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference interface {
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
	Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyList
	// Experimental.
	ItemsInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyPropertyList
	// Experimental.
	PropertyInput() interface{}
	// Experimental.
	Required() interface{}
	// Experimental.
	SetRequired(val interface{})
	// Experimental.
	RequiredInput() interface{}
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
	// Experimental.
	ResetRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference
type jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyItemsPropertyList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyPropertyList
	_jsii_.Get(
		j,
		"property",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) PropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference_Override(t TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) PutItems(value interface{}) {
	if err := t.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putItems",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) PutProperty(value interface{}) {
	if err := t.validatePutPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProperty",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		t,
		"resetItems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ResetProperty() {
	_jsii_.InvokeVoid(
		t,
		"resetProperty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		t,
		"resetRequired",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaPropertyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

