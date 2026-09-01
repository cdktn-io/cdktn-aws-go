package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference interface {
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
	InitialAudioGain() *float64
	// Experimental.
	SetInitialAudioGain(val *float64)
	// Experimental.
	InitialAudioGainInput() *float64
	// Experimental.
	InputEndAction() *string
	// Experimental.
	SetInputEndAction(val *string)
	// Experimental.
	InputEndActionInput() *string
	// Experimental.
	InputLossBehavior() AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference
	// Experimental.
	InputLossBehaviorInput() *AwsMedialiveChannel_InputLossBehaviorProperty
	// Experimental.
	InternalValue() *AwsMedialiveChannel_GlobalConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_GlobalConfigurationProperty)
	// Experimental.
	OutputLockingMode() *string
	// Experimental.
	SetOutputLockingMode(val *string)
	// Experimental.
	OutputLockingModeInput() *string
	// Experimental.
	OutputTimingSource() *string
	// Experimental.
	SetOutputTimingSource(val *string)
	// Experimental.
	OutputTimingSourceInput() *string
	// Experimental.
	SupportLowFramerateInputs() *string
	// Experimental.
	SetSupportLowFramerateInputs(val *string)
	// Experimental.
	SupportLowFramerateInputsInput() *string
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
	PutInputLossBehavior(value *AwsMedialiveChannel_InputLossBehaviorProperty)
	// Experimental.
	ResetInitialAudioGain()
	// Experimental.
	ResetInputEndAction()
	// Experimental.
	ResetInputLossBehavior()
	// Experimental.
	ResetOutputLockingMode()
	// Experimental.
	ResetOutputTimingSource()
	// Experimental.
	ResetSupportLowFramerateInputs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InitialAudioGain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InitialAudioGainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InputEndAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InputEndActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InputLossBehavior() AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference {
	var returns AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference
	_jsii_.Get(
		j,
		"inputLossBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InputLossBehaviorInput() *AwsMedialiveChannel_InputLossBehaviorProperty {
	var returns *AwsMedialiveChannel_InputLossBehaviorProperty
	_jsii_.Get(
		j,
		"inputLossBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InternalValue() *AwsMedialiveChannel_GlobalConfigurationProperty {
	var returns *AwsMedialiveChannel_GlobalConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) OutputLockingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) OutputLockingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) OutputTimingSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) OutputTimingSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) SupportLowFramerateInputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) SupportLowFramerateInputsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_GlobalConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_GlobalConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.GlobalConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_GlobalConfigurationPropertyOutputReference_Override(a AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.GlobalConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetInitialAudioGain(val *float64) {
	if err := j.validateSetInitialAudioGainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialAudioGain",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetInputEndAction(val *string) {
	if err := j.validateSetInputEndActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEndAction",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_GlobalConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetOutputLockingMode(val *string) {
	if err := j.validateSetOutputLockingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLockingMode",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetOutputTimingSource(val *string) {
	if err := j.validateSetOutputTimingSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputTimingSource",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetSupportLowFramerateInputs(val *string) {
	if err := j.validateSetSupportLowFramerateInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportLowFramerateInputs",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) PutInputLossBehavior(value *AwsMedialiveChannel_InputLossBehaviorProperty) {
	if err := a.validatePutInputLossBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputLossBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetInitialAudioGain() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialAudioGain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetInputEndAction() {
	_jsii_.InvokeVoid(
		a,
		"resetInputEndAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetInputLossBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetOutputLockingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputLockingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetOutputTimingSource() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputTimingSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ResetSupportLowFramerateInputs() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportLowFramerateInputs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

