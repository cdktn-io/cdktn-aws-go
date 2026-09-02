package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_CaptionDescriptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Accessibility() *string
	// Experimental.
	SetAccessibility(val *string)
	// Experimental.
	AccessibilityInput() *string
	// Experimental.
	CaptionSelectorName() *string
	// Experimental.
	SetCaptionSelectorName(val *string)
	// Experimental.
	CaptionSelectorNameInput() *string
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
	DestinationSettings() TfChannel_DestinationSettingsPropertyOutputReference
	// Experimental.
	DestinationSettingsInput() *TfChannel_DestinationSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LanguageCode() *string
	// Experimental.
	SetLanguageCode(val *string)
	// Experimental.
	LanguageCodeInput() *string
	// Experimental.
	LanguageDescription() *string
	// Experimental.
	SetLanguageDescription(val *string)
	// Experimental.
	LanguageDescriptionInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
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
	PutDestinationSettings(value *TfChannel_DestinationSettingsProperty)
	// Experimental.
	ResetAccessibility()
	// Experimental.
	ResetDestinationSettings()
	// Experimental.
	ResetLanguageCode()
	// Experimental.
	ResetLanguageDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_CaptionDescriptionsPropertyOutputReference
type jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) Accessibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) AccessibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) CaptionSelectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionSelectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) CaptionSelectorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionSelectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) DestinationSettings() TfChannel_DestinationSettingsPropertyOutputReference {
	var returns TfChannel_DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) DestinationSettingsInput() *TfChannel_DestinationSettingsProperty {
	var returns *TfChannel_DestinationSettingsProperty
	_jsii_.Get(
		j,
		"destinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) LanguageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) LanguageDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_CaptionDescriptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_CaptionDescriptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_CaptionDescriptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.CaptionDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_CaptionDescriptionsPropertyOutputReference_Override(t TfChannel_CaptionDescriptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.CaptionDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetAccessibility(val *string) {
	if err := j.validateSetAccessibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibility",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetCaptionSelectorName(val *string) {
	if err := j.validateSetCaptionSelectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionSelectorName",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetLanguageDescription(val *string) {
	if err := j.validateSetLanguageDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageDescription",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) PutDestinationSettings(value *TfChannel_DestinationSettingsProperty) {
	if err := t.validatePutDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ResetAccessibility() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessibility",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ResetDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		t,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ResetLanguageDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetLanguageDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_CaptionDescriptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

