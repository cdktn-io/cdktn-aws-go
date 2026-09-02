package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference interface {
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
	CustomPayload() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageCustomPayloadPropertyList
	// Experimental.
	CustomPayloadInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageResponseCard() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageImageResponseCardPropertyList
	// Experimental.
	ImageResponseCardInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PlainTextMessage() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePlainTextMessagePropertyList
	// Experimental.
	PlainTextMessageInput() interface{}
	// Experimental.
	SsmlMessage() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageSsmlMessagePropertyList
	// Experimental.
	SsmlMessageInput() interface{}
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
	PutCustomPayload(value interface{})
	// Experimental.
	PutImageResponseCard(value interface{})
	// Experimental.
	PutPlainTextMessage(value interface{})
	// Experimental.
	PutSsmlMessage(value interface{})
	// Experimental.
	ResetCustomPayload()
	// Experimental.
	ResetImageResponseCard()
	// Experimental.
	ResetPlainTextMessage()
	// Experimental.
	ResetSsmlMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference
type jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) CustomPayload() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageCustomPayloadPropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageCustomPayloadPropertyList
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) CustomPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ImageResponseCard() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageImageResponseCardPropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageImageResponseCardPropertyList
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ImageResponseCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageResponseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PlainTextMessage() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePlainTextMessagePropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePlainTextMessagePropertyList
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PlainTextMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) SsmlMessage() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageSsmlMessagePropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageSsmlMessagePropertyList
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) SsmlMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmlMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference_Override(t TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PutCustomPayload(value interface{}) {
	if err := t.validatePutCustomPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomPayload",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PutImageResponseCard(value interface{}) {
	if err := t.validatePutImageResponseCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageResponseCard",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PutPlainTextMessage(value interface{}) {
	if err := t.validatePutPlainTextMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlainTextMessage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) PutSsmlMessage(value interface{}) {
	if err := t.validatePutSsmlMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSsmlMessage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ResetCustomPayload() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPayload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ResetImageResponseCard() {
	_jsii_.InvokeVoid(
		t,
		"resetImageResponseCard",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ResetPlainTextMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetPlainTextMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ResetSsmlMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetSsmlMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

