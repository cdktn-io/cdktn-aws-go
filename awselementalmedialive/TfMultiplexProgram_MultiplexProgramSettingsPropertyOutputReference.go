package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PreferredChannelPipeline() *string
	// Experimental.
	SetPreferredChannelPipeline(val *string)
	// Experimental.
	PreferredChannelPipelineInput() *string
	// Experimental.
	ProgramNumber() *float64
	// Experimental.
	SetProgramNumber(val *float64)
	// Experimental.
	ProgramNumberInput() *float64
	// Experimental.
	ServiceDescriptor() TfMultiplexProgram_ServiceDescriptorPropertyList
	// Experimental.
	ServiceDescriptorInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VideoSettings() TfMultiplexProgram_VideoSettingsPropertyList
	// Experimental.
	VideoSettingsInput() interface{}
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
	PutServiceDescriptor(value interface{})
	// Experimental.
	PutVideoSettings(value interface{})
	// Experimental.
	ResetServiceDescriptor()
	// Experimental.
	ResetVideoSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference
type jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PreferredChannelPipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ProgramNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ProgramNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ServiceDescriptor() TfMultiplexProgram_ServiceDescriptorPropertyList {
	var returns TfMultiplexProgram_ServiceDescriptorPropertyList
	_jsii_.Get(
		j,
		"serviceDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ServiceDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) VideoSettings() TfMultiplexProgram_VideoSettingsPropertyList {
	var returns TfMultiplexProgram_VideoSettingsPropertyList
	_jsii_.Get(
		j,
		"videoSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) VideoSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfMultiplexProgram.MultiplexProgramSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference_Override(t TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfMultiplexProgram.MultiplexProgramSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetPreferredChannelPipeline(val *string) {
	if err := j.validateSetPreferredChannelPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetProgramNumber(val *float64) {
	if err := j.validateSetProgramNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNumber",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PutServiceDescriptor(value interface{}) {
	if err := t.validatePutServiceDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceDescriptor",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PutVideoSettings(value interface{}) {
	if err := t.validatePutVideoSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVideoSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ResetServiceDescriptor() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceDescriptor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ResetVideoSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

