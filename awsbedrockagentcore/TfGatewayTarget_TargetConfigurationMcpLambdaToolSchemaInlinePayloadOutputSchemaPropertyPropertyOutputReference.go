package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference interface {
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
	Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyItemsPropertyList
	// Experimental.
	ItemsInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyPropertyList
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

// The jsii proxy struct for TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference
type jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Items() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyItemsPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyItemsPropertyList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Property() TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyPropertyList {
	var returns TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyPropertyList
	_jsii_.Get(
		j,
		"property",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) PropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference_Override(t TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) PutItems(value interface{}) {
	if err := t.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putItems",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) PutProperty(value interface{}) {
	if err := t.validatePutPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProperty",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		t,
		"resetItems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ResetProperty() {
	_jsii_.InvokeVoid(
		t,
		"resetProperty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		t,
		"resetRequired",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaPropertyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

