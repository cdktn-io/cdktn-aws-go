package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGroup_PreferencesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AlarmSpecification() TfGroup_AlarmSpecificationPropertyOutputReference
	// Experimental.
	AlarmSpecificationInput() *TfGroup_AlarmSpecificationProperty
	// Experimental.
	AutoRollback() interface{}
	// Experimental.
	SetAutoRollback(val interface{})
	// Experimental.
	AutoRollbackInput() interface{}
	// Experimental.
	CheckpointDelay() *string
	// Experimental.
	SetCheckpointDelay(val *string)
	// Experimental.
	CheckpointDelayInput() *string
	// Experimental.
	CheckpointPercentages() *[]*float64
	// Experimental.
	SetCheckpointPercentages(val *[]*float64)
	// Experimental.
	CheckpointPercentagesInput() *[]*float64
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
	InstanceWarmup() *string
	// Experimental.
	SetInstanceWarmup(val *string)
	// Experimental.
	InstanceWarmupInput() *string
	// Experimental.
	InternalValue() *TfGroup_PreferencesProperty
	// Experimental.
	SetInternalValue(val *TfGroup_PreferencesProperty)
	// Experimental.
	MaxHealthyPercentage() *float64
	// Experimental.
	SetMaxHealthyPercentage(val *float64)
	// Experimental.
	MaxHealthyPercentageInput() *float64
	// Experimental.
	MinHealthyPercentage() *float64
	// Experimental.
	SetMinHealthyPercentage(val *float64)
	// Experimental.
	MinHealthyPercentageInput() *float64
	// Experimental.
	ScaleInProtectedInstances() *string
	// Experimental.
	SetScaleInProtectedInstances(val *string)
	// Experimental.
	ScaleInProtectedInstancesInput() *string
	// Experimental.
	SkipMatching() interface{}
	// Experimental.
	SetSkipMatching(val interface{})
	// Experimental.
	SkipMatchingInput() interface{}
	// Experimental.
	StandbyInstances() *string
	// Experimental.
	SetStandbyInstances(val *string)
	// Experimental.
	StandbyInstancesInput() *string
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
	PutAlarmSpecification(value *TfGroup_AlarmSpecificationProperty)
	// Experimental.
	ResetAlarmSpecification()
	// Experimental.
	ResetAutoRollback()
	// Experimental.
	ResetCheckpointDelay()
	// Experimental.
	ResetCheckpointPercentages()
	// Experimental.
	ResetInstanceWarmup()
	// Experimental.
	ResetMaxHealthyPercentage()
	// Experimental.
	ResetMinHealthyPercentage()
	// Experimental.
	ResetScaleInProtectedInstances()
	// Experimental.
	ResetSkipMatching()
	// Experimental.
	ResetStandbyInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGroup_PreferencesPropertyOutputReference
type jsiiProxy_TfGroup_PreferencesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) AlarmSpecification() TfGroup_AlarmSpecificationPropertyOutputReference {
	var returns TfGroup_AlarmSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"alarmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) AlarmSpecificationInput() *TfGroup_AlarmSpecificationProperty {
	var returns *TfGroup_AlarmSpecificationProperty
	_jsii_.Get(
		j,
		"alarmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) AutoRollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) AutoRollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) CheckpointDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) CheckpointDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) CheckpointPercentages() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"checkpointPercentages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) CheckpointPercentagesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"checkpointPercentagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) InstanceWarmup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) InstanceWarmupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) InternalValue() *TfGroup_PreferencesProperty {
	var returns *TfGroup_PreferencesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) MaxHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) MaxHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) MinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) MinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ScaleInProtectedInstances() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInProtectedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ScaleInProtectedInstancesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInProtectedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) SkipMatching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) SkipMatchingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) StandbyInstances() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) StandbyInstancesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGroup_PreferencesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGroup_PreferencesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGroup_PreferencesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGroup_PreferencesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.PreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGroup_PreferencesPropertyOutputReference_Override(t TfGroup_PreferencesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.PreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetAutoRollback(val interface{}) {
	if err := j.validateSetAutoRollbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRollback",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetCheckpointDelay(val *string) {
	if err := j.validateSetCheckpointDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointDelay",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetCheckpointPercentages(val *[]*float64) {
	if err := j.validateSetCheckpointPercentagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointPercentages",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetInstanceWarmup(val *string) {
	if err := j.validateSetInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceWarmup",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetInternalValue(val *TfGroup_PreferencesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetMaxHealthyPercentage(val *float64) {
	if err := j.validateSetMaxHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetMinHealthyPercentage(val *float64) {
	if err := j.validateSetMinHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetScaleInProtectedInstances(val *string) {
	if err := j.validateSetScaleInProtectedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInProtectedInstances",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetSkipMatching(val interface{}) {
	if err := j.validateSetSkipMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMatching",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetStandbyInstances(val *string) {
	if err := j.validateSetStandbyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyInstances",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGroup_PreferencesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) PutAlarmSpecification(value *TfGroup_AlarmSpecificationProperty) {
	if err := t.validatePutAlarmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlarmSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetAlarmSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetAlarmSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetAutoRollback() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRollback",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetCheckpointDelay() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointDelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetCheckpointPercentages() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointPercentages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetInstanceWarmup() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceWarmup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetMaxHealthyPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxHealthyPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMinHealthyPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetScaleInProtectedInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetScaleInProtectedInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetSkipMatching() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipMatching",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ResetStandbyInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetStandbyInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGroup_PreferencesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

