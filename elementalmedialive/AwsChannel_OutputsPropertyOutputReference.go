package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_OutputsPropertyOutputReference interface {
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
	OutputSettings() AwsChannel_OutputSettingsPropertyOutputReference
	// Experimental.
	OutputSettingsInput() *AwsChannel_OutputSettingsProperty
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
	PutOutputSettings(value *AwsChannel_OutputSettingsProperty)
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

// The jsii proxy struct for AwsChannel_OutputsPropertyOutputReference
type jsiiProxy_AwsChannel_OutputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) AudioDescriptionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audioDescriptionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) AudioDescriptionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audioDescriptionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) CaptionDescriptionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"captionDescriptionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) CaptionDescriptionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"captionDescriptionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) OutputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) OutputSettings() AwsChannel_OutputSettingsPropertyOutputReference {
	var returns AwsChannel_OutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"outputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) OutputSettingsInput() *AwsChannel_OutputSettingsProperty {
	var returns *AwsChannel_OutputSettingsProperty
	_jsii_.Get(
		j,
		"outputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) VideoDescriptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoDescriptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) VideoDescriptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoDescriptionNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_OutputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsChannel_OutputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_OutputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_OutputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_OutputsPropertyOutputReference_Override(a AwsChannel_OutputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetAudioDescriptionNames(val *[]*string) {
	if err := j.validateSetAudioDescriptionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioDescriptionNames",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetCaptionDescriptionNames(val *[]*string) {
	if err := j.validateSetCaptionDescriptionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionDescriptionNames",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetOutputName(val *string) {
	if err := j.validateSetOutputNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputName",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputsPropertyOutputReference)SetVideoDescriptionName(val *string) {
	if err := j.validateSetVideoDescriptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoDescriptionName",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) PutOutputSettings(value *AwsChannel_OutputSettingsProperty) {
	if err := a.validatePutOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ResetAudioDescriptionNames() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioDescriptionNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ResetCaptionDescriptionNames() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionDescriptionNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ResetOutputName() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ResetVideoDescriptionName() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoDescriptionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

