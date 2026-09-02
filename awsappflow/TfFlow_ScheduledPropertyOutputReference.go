package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_ScheduledPropertyOutputReference interface {
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
	DataPullMode() *string
	// Experimental.
	SetDataPullMode(val *string)
	// Experimental.
	DataPullModeInput() *string
	// Experimental.
	FirstExecutionFrom() *string
	// Experimental.
	SetFirstExecutionFrom(val *string)
	// Experimental.
	FirstExecutionFromInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFlow_ScheduledProperty
	// Experimental.
	SetInternalValue(val *TfFlow_ScheduledProperty)
	// Experimental.
	ScheduleEndTime() *string
	// Experimental.
	SetScheduleEndTime(val *string)
	// Experimental.
	ScheduleEndTimeInput() *string
	// Experimental.
	ScheduleExpression() *string
	// Experimental.
	SetScheduleExpression(val *string)
	// Experimental.
	ScheduleExpressionInput() *string
	// Experimental.
	ScheduleOffset() *float64
	// Experimental.
	SetScheduleOffset(val *float64)
	// Experimental.
	ScheduleOffsetInput() *float64
	// Experimental.
	ScheduleStartTime() *string
	// Experimental.
	SetScheduleStartTime(val *string)
	// Experimental.
	ScheduleStartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timezone() *string
	// Experimental.
	SetTimezone(val *string)
	// Experimental.
	TimezoneInput() *string
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
	ResetDataPullMode()
	// Experimental.
	ResetFirstExecutionFrom()
	// Experimental.
	ResetScheduleEndTime()
	// Experimental.
	ResetScheduleOffset()
	// Experimental.
	ResetScheduleStartTime()
	// Experimental.
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_ScheduledPropertyOutputReference
type jsiiProxy_TfFlow_ScheduledPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) DataPullMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) DataPullModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) FirstExecutionFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstExecutionFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) FirstExecutionFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstExecutionFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) InternalValue() *TfFlow_ScheduledProperty {
	var returns *TfFlow_ScheduledProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleEndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleEndTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ScheduleStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_ScheduledPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_ScheduledPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_ScheduledPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_ScheduledPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.ScheduledPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_ScheduledPropertyOutputReference_Override(t TfFlow_ScheduledPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.ScheduledPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetDataPullMode(val *string) {
	if err := j.validateSetDataPullModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPullMode",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetFirstExecutionFrom(val *string) {
	if err := j.validateSetFirstExecutionFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstExecutionFrom",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetInternalValue(val *TfFlow_ScheduledProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetScheduleEndTime(val *string) {
	if err := j.validateSetScheduleEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleEndTime",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetScheduleExpression(val *string) {
	if err := j.validateSetScheduleExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetScheduleOffset(val *float64) {
	if err := j.validateSetScheduleOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleOffset",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetScheduleStartTime(val *string) {
	if err := j.validateSetScheduleStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleStartTime",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFlow_ScheduledPropertyOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetDataPullMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDataPullMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetFirstExecutionFrom() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstExecutionFrom",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetScheduleEndTime() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduleEndTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetScheduleOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduleOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetScheduleStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduleStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_ScheduledPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

