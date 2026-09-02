package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClassificationJob_ScheduleFrequencyPropertyOutputReference interface {
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
	DailySchedule() interface{}
	// Experimental.
	SetDailySchedule(val interface{})
	// Experimental.
	DailyScheduleInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfClassificationJob_ScheduleFrequencyProperty
	// Experimental.
	SetInternalValue(val *TfClassificationJob_ScheduleFrequencyProperty)
	// Experimental.
	MonthlySchedule() *float64
	// Experimental.
	SetMonthlySchedule(val *float64)
	// Experimental.
	MonthlyScheduleInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WeeklySchedule() *string
	// Experimental.
	SetWeeklySchedule(val *string)
	// Experimental.
	WeeklyScheduleInput() *string
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
	ResetDailySchedule()
	// Experimental.
	ResetMonthlySchedule()
	// Experimental.
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfClassificationJob_ScheduleFrequencyPropertyOutputReference
type jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) DailySchedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) DailyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) InternalValue() *TfClassificationJob_ScheduleFrequencyProperty {
	var returns *TfClassificationJob_ScheduleFrequencyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) MonthlySchedule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) MonthlyScheduleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) WeeklySchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) WeeklyScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfClassificationJob_ScheduleFrequencyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfClassificationJob_ScheduleFrequencyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfClassificationJob_ScheduleFrequencyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.ScheduleFrequencyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfClassificationJob_ScheduleFrequencyPropertyOutputReference_Override(t TfClassificationJob_ScheduleFrequencyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.ScheduleFrequencyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetDailySchedule(val interface{}) {
	if err := j.validateSetDailyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailySchedule",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetInternalValue(val *TfClassificationJob_ScheduleFrequencyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetMonthlySchedule(val *float64) {
	if err := j.validateSetMonthlyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlySchedule",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference)SetWeeklySchedule(val *string) {
	if err := j.validateSetWeeklyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklySchedule",
		val,
	)
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfClassificationJob_ScheduleFrequencyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

