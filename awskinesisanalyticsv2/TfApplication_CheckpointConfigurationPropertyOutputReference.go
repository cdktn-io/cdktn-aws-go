package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_CheckpointConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CheckpointingEnabled() interface{}
	// Experimental.
	SetCheckpointingEnabled(val interface{})
	// Experimental.
	CheckpointingEnabledInput() interface{}
	// Experimental.
	CheckpointInterval() *float64
	// Experimental.
	SetCheckpointInterval(val *float64)
	// Experimental.
	CheckpointIntervalInput() *float64
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
	// Experimental.
	ConfigurationType() *string
	// Experimental.
	SetConfigurationType(val *string)
	// Experimental.
	ConfigurationTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfApplication_CheckpointConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfApplication_CheckpointConfigurationProperty)
	// Experimental.
	MinPauseBetweenCheckpoints() *float64
	// Experimental.
	SetMinPauseBetweenCheckpoints(val *float64)
	// Experimental.
	MinPauseBetweenCheckpointsInput() *float64
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
	ResetCheckpointingEnabled()
	// Experimental.
	ResetCheckpointInterval()
	// Experimental.
	ResetMinPauseBetweenCheckpoints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_CheckpointConfigurationPropertyOutputReference
type jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) CheckpointingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) CheckpointingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) CheckpointInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkpointInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) CheckpointIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkpointIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ConfigurationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ConfigurationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) InternalValue() *TfApplication_CheckpointConfigurationProperty {
	var returns *TfApplication_CheckpointConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) MinPauseBetweenCheckpoints() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPauseBetweenCheckpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) MinPauseBetweenCheckpointsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPauseBetweenCheckpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_CheckpointConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_CheckpointConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_CheckpointConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.CheckpointConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_CheckpointConfigurationPropertyOutputReference_Override(t TfApplication_CheckpointConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.CheckpointConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetCheckpointingEnabled(val interface{}) {
	if err := j.validateSetCheckpointingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointingEnabled",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetCheckpointInterval(val *float64) {
	if err := j.validateSetCheckpointIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointInterval",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetConfigurationType(val *string) {
	if err := j.validateSetConfigurationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationType",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetInternalValue(val *TfApplication_CheckpointConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetMinPauseBetweenCheckpoints(val *float64) {
	if err := j.validateSetMinPauseBetweenCheckpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPauseBetweenCheckpoints",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ResetCheckpointingEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointingEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ResetCheckpointInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ResetMinPauseBetweenCheckpoints() {
	_jsii_.InvokeVoid(
		t,
		"resetMinPauseBetweenCheckpoints",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_CheckpointConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

