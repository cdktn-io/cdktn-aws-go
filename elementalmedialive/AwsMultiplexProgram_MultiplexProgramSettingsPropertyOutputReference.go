package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference interface {
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
	ServiceDescriptor() AwsMultiplexProgram_ServiceDescriptorPropertyList
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
	VideoSettings() AwsMultiplexProgram_VideoSettingsPropertyList
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

// The jsii proxy struct for AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference
type jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PreferredChannelPipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ProgramNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ProgramNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ServiceDescriptor() AwsMultiplexProgram_ServiceDescriptorPropertyList {
	var returns AwsMultiplexProgram_ServiceDescriptorPropertyList
	_jsii_.Get(
		j,
		"serviceDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ServiceDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) VideoSettings() AwsMultiplexProgram_VideoSettingsPropertyList {
	var returns AwsMultiplexProgram_VideoSettingsPropertyList
	_jsii_.Get(
		j,
		"videoSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) VideoSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMultiplexProgram.MultiplexProgramSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference_Override(a AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMultiplexProgram.MultiplexProgramSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetPreferredChannelPipeline(val *string) {
	if err := j.validateSetPreferredChannelPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetProgramNumber(val *float64) {
	if err := j.validateSetProgramNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNumber",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PutServiceDescriptor(value interface{}) {
	if err := a.validatePutServiceDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceDescriptor",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) PutVideoSettings(value interface{}) {
	if err := a.validatePutVideoSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVideoSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ResetServiceDescriptor() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceDescriptor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ResetVideoSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMultiplexProgram_MultiplexProgramSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

