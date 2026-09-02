package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_InputAttachmentsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutomaticInputFailoverSettings() TfChannel_AutomaticInputFailoverSettingsPropertyOutputReference
	// Experimental.
	AutomaticInputFailoverSettingsInput() *TfChannel_AutomaticInputFailoverSettingsProperty
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
	InputAttachmentName() *string
	// Experimental.
	SetInputAttachmentName(val *string)
	// Experimental.
	InputAttachmentNameInput() *string
	// Experimental.
	InputId() *string
	// Experimental.
	SetInputId(val *string)
	// Experimental.
	InputIdInput() *string
	// Experimental.
	InputSettings() TfChannel_InputSettingsPropertyOutputReference
	// Experimental.
	InputSettingsInput() *TfChannel_InputSettingsProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutAutomaticInputFailoverSettings(value *TfChannel_AutomaticInputFailoverSettingsProperty)
	// Experimental.
	PutInputSettings(value *TfChannel_InputSettingsProperty)
	// Experimental.
	ResetAutomaticInputFailoverSettings()
	// Experimental.
	ResetInputSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_InputAttachmentsPropertyOutputReference
type jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) AutomaticInputFailoverSettings() TfChannel_AutomaticInputFailoverSettingsPropertyOutputReference {
	var returns TfChannel_AutomaticInputFailoverSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"automaticInputFailoverSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) AutomaticInputFailoverSettingsInput() *TfChannel_AutomaticInputFailoverSettingsProperty {
	var returns *TfChannel_AutomaticInputFailoverSettingsProperty
	_jsii_.Get(
		j,
		"automaticInputFailoverSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputAttachmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAttachmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputAttachmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAttachmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputSettings() TfChannel_InputSettingsPropertyOutputReference {
	var returns TfChannel_InputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"inputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InputSettingsInput() *TfChannel_InputSettingsProperty {
	var returns *TfChannel_InputSettingsProperty
	_jsii_.Get(
		j,
		"inputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_InputAttachmentsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_InputAttachmentsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_InputAttachmentsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_InputAttachmentsPropertyOutputReference_Override(t TfChannel_InputAttachmentsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetInputAttachmentName(val *string) {
	if err := j.validateSetInputAttachmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAttachmentName",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetInputId(val *string) {
	if err := j.validateSetInputIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputId",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) PutAutomaticInputFailoverSettings(value *TfChannel_AutomaticInputFailoverSettingsProperty) {
	if err := t.validatePutAutomaticInputFailoverSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutomaticInputFailoverSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) PutInputSettings(value *TfChannel_InputSettingsProperty) {
	if err := t.validatePutInputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ResetAutomaticInputFailoverSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAutomaticInputFailoverSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ResetInputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetInputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

