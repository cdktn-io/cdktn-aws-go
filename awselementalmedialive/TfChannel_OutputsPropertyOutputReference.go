package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_OutputsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioDescriptionNames() *[]*string
	// Experimental.
	SetAudioDescriptionNames(val *[]*string)
	// Experimental.
	AudioDescriptionNamesInput() *[]*string
	// Experimental.
	CaptionDescriptionNames() *[]*string
	// Experimental.
	SetCaptionDescriptionNames(val *[]*string)
	// Experimental.
	CaptionDescriptionNamesInput() *[]*string
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
	OutputName() *string
	// Experimental.
	SetOutputName(val *string)
	// Experimental.
	OutputNameInput() *string
	// Experimental.
	OutputSettings() TfChannel_OutputSettingsPropertyOutputReference
	// Experimental.
	OutputSettingsInput() *TfChannel_OutputSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VideoDescriptionName() *string
	// Experimental.
	SetVideoDescriptionName(val *string)
	// Experimental.
	VideoDescriptionNameInput() *string
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
	PutOutputSettings(value *TfChannel_OutputSettingsProperty)
	// Experimental.
	ResetAudioDescriptionNames()
	// Experimental.
	ResetCaptionDescriptionNames()
	// Experimental.
	ResetOutputName()
	// Experimental.
	ResetVideoDescriptionName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_OutputsPropertyOutputReference
type jsiiProxy_TfChannel_OutputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) AudioDescriptionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audioDescriptionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) AudioDescriptionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audioDescriptionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) CaptionDescriptionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"captionDescriptionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) CaptionDescriptionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"captionDescriptionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) OutputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) OutputSettings() TfChannel_OutputSettingsPropertyOutputReference {
	var returns TfChannel_OutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"outputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) OutputSettingsInput() *TfChannel_OutputSettingsProperty {
	var returns *TfChannel_OutputSettingsProperty
	_jsii_.Get(
		j,
		"outputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) VideoDescriptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoDescriptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference) VideoDescriptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoDescriptionNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_OutputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_OutputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_OutputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_OutputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_OutputsPropertyOutputReference_Override(t TfChannel_OutputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetAudioDescriptionNames(val *[]*string) {
	if err := j.validateSetAudioDescriptionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioDescriptionNames",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetCaptionDescriptionNames(val *[]*string) {
	if err := j.validateSetCaptionDescriptionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionDescriptionNames",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetOutputName(val *string) {
	if err := j.validateSetOutputNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputName",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputsPropertyOutputReference)SetVideoDescriptionName(val *string) {
	if err := j.validateSetVideoDescriptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoDescriptionName",
		val,
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) PutOutputSettings(value *TfChannel_OutputSettingsProperty) {
	if err := t.validatePutOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ResetAudioDescriptionNames() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioDescriptionNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ResetCaptionDescriptionNames() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionDescriptionNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ResetOutputName() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ResetVideoDescriptionName() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoDescriptionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_OutputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

