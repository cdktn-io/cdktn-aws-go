package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty)
	// Experimental.
	M2TsSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference
	// Experimental.
	M2TsSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty
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
	PutM2TsSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty)
	// Experimental.
	ResetM2TsSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) InternalValue() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) M2TsSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"m2TsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) M2TsSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty
	_jsii_.Get(
		j,
		"m2TsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference_Override(t TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) PutM2TsSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty) {
	if err := t.validatePutM2TsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putM2TsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) ResetM2TsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetM2TsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

