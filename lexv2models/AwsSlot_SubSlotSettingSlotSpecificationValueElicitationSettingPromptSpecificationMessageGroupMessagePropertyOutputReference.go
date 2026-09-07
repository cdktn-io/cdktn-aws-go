package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference interface {
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
	CustomPayload() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageCustomPayloadPropertyList
	// Experimental.
	CustomPayloadInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageResponseCard() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageImageResponseCardPropertyList
	// Experimental.
	ImageResponseCardInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PlainTextMessage() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePlainTextMessagePropertyList
	// Experimental.
	PlainTextMessageInput() interface{}
	// Experimental.
	SsmlMessage() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageSsmlMessagePropertyList
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

// The jsii proxy struct for AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference
type jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) CustomPayload() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageCustomPayloadPropertyList {
	var returns AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageCustomPayloadPropertyList
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) CustomPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ImageResponseCard() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageImageResponseCardPropertyList {
	var returns AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageImageResponseCardPropertyList
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ImageResponseCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageResponseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PlainTextMessage() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePlainTextMessagePropertyList {
	var returns AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePlainTextMessagePropertyList
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PlainTextMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) SsmlMessage() AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageSsmlMessagePropertyList {
	var returns AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessageSsmlMessagePropertyList
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) SsmlMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmlMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference_Override(a AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PutCustomPayload(value interface{}) {
	if err := a.validatePutCustomPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPayload",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PutImageResponseCard(value interface{}) {
	if err := a.validatePutImageResponseCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageResponseCard",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PutPlainTextMessage(value interface{}) {
	if err := a.validatePutPlainTextMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlainTextMessage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) PutSsmlMessage(value interface{}) {
	if err := a.validatePutSsmlMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSsmlMessage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ResetCustomPayload() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPayload",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ResetImageResponseCard() {
	_jsii_.InvokeVoid(
		a,
		"resetImageResponseCard",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ResetPlainTextMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetPlainTextMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ResetSsmlMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSsmlMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroupMessagePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

