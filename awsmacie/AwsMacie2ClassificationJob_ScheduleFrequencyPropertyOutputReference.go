package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference interface {
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
	InternalValue() *AwsMacie2ClassificationJob_ScheduleFrequencyProperty
	// Experimental.
	SetInternalValue(val *AwsMacie2ClassificationJob_ScheduleFrequencyProperty)
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

// The jsii proxy struct for AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference
type jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) DailySchedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) DailyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) InternalValue() *AwsMacie2ClassificationJob_ScheduleFrequencyProperty {
	var returns *AwsMacie2ClassificationJob_ScheduleFrequencyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) MonthlySchedule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) MonthlyScheduleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) WeeklySchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) WeeklyScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2ClassificationJob.ScheduleFrequencyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference_Override(a AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2ClassificationJob.ScheduleFrequencyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetDailySchedule(val interface{}) {
	if err := j.validateSetDailyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailySchedule",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetInternalValue(val *AwsMacie2ClassificationJob_ScheduleFrequencyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetMonthlySchedule(val *float64) {
	if err := j.validateSetMonthlyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlySchedule",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference)SetWeeklySchedule(val *string) {
	if err := j.validateSetWeeklyScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklySchedule",
		val,
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMacie2ClassificationJob_ScheduleFrequencyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

