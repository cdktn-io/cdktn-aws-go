package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_GlobalConfigurationPropertyOutputReference interface {
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
	InputLossBehavior() TfChannel_InputLossBehaviorPropertyOutputReference
	// Experimental.
	InputLossBehaviorInput() *TfChannel_InputLossBehaviorProperty
	// Experimental.
	InternalValue() *TfChannel_GlobalConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfChannel_GlobalConfigurationProperty)
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
	PutInputLossBehavior(value *TfChannel_InputLossBehaviorProperty)
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

// The jsii proxy struct for TfChannel_GlobalConfigurationPropertyOutputReference
type jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InitialAudioGain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InitialAudioGainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InputEndAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InputEndActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InputLossBehavior() TfChannel_InputLossBehaviorPropertyOutputReference {
	var returns TfChannel_InputLossBehaviorPropertyOutputReference
	_jsii_.Get(
		j,
		"inputLossBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InputLossBehaviorInput() *TfChannel_InputLossBehaviorProperty {
	var returns *TfChannel_InputLossBehaviorProperty
	_jsii_.Get(
		j,
		"inputLossBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InternalValue() *TfChannel_GlobalConfigurationProperty {
	var returns *TfChannel_GlobalConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) OutputLockingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) OutputLockingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) OutputTimingSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) OutputTimingSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) SupportLowFramerateInputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) SupportLowFramerateInputsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_GlobalConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_GlobalConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_GlobalConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.GlobalConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_GlobalConfigurationPropertyOutputReference_Override(t TfChannel_GlobalConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.GlobalConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetInitialAudioGain(val *float64) {
	if err := j.validateSetInitialAudioGainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialAudioGain",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetInputEndAction(val *string) {
	if err := j.validateSetInputEndActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEndAction",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetInternalValue(val *TfChannel_GlobalConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetOutputLockingMode(val *string) {
	if err := j.validateSetOutputLockingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLockingMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetOutputTimingSource(val *string) {
	if err := j.validateSetOutputTimingSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputTimingSource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetSupportLowFramerateInputs(val *string) {
	if err := j.validateSetSupportLowFramerateInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportLowFramerateInputs",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) PutInputLossBehavior(value *TfChannel_InputLossBehaviorProperty) {
	if err := t.validatePutInputLossBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputLossBehavior",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetInitialAudioGain() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialAudioGain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetInputEndAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputEndAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetInputLossBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetInputLossBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetOutputLockingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputLockingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetOutputTimingSource() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputTimingSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ResetSupportLowFramerateInputs() {
	_jsii_.InvokeVoid(
		t,
		"resetSupportLowFramerateInputs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_GlobalConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

