package awsssmcontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssmcontacts/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssmcontacts/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRotation_RecurrencePropertyOutputReference interface {
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
	DailySettings() TfRotation_DailySettingsPropertyList
	// Experimental.
	DailySettingsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MonthlySettings() TfRotation_MonthlySettingsPropertyList
	// Experimental.
	MonthlySettingsInput() interface{}
	// Experimental.
	NumberOfOnCalls() *float64
	// Experimental.
	SetNumberOfOnCalls(val *float64)
	// Experimental.
	NumberOfOnCallsInput() *float64
	// Experimental.
	RecurrenceMultiplier() *float64
	// Experimental.
	SetRecurrenceMultiplier(val *float64)
	// Experimental.
	RecurrenceMultiplierInput() *float64
	// Experimental.
	ShiftCoverages() TfRotation_ShiftCoveragesPropertyList
	// Experimental.
	ShiftCoveragesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WeeklySettings() TfRotation_WeeklySettingsPropertyList
	// Experimental.
	WeeklySettingsInput() interface{}
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
	PutDailySettings(value interface{})
	// Experimental.
	PutMonthlySettings(value interface{})
	// Experimental.
	PutShiftCoverages(value interface{})
	// Experimental.
	PutWeeklySettings(value interface{})
	// Experimental.
	ResetDailySettings()
	// Experimental.
	ResetMonthlySettings()
	// Experimental.
	ResetShiftCoverages()
	// Experimental.
	ResetWeeklySettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRotation_RecurrencePropertyOutputReference
type jsiiProxy_TfRotation_RecurrencePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) DailySettings() TfRotation_DailySettingsPropertyList {
	var returns TfRotation_DailySettingsPropertyList
	_jsii_.Get(
		j,
		"dailySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) DailySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) MonthlySettings() TfRotation_MonthlySettingsPropertyList {
	var returns TfRotation_MonthlySettingsPropertyList
	_jsii_.Get(
		j,
		"monthlySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) MonthlySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) NumberOfOnCalls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfOnCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) NumberOfOnCallsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfOnCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) RecurrenceMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recurrenceMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) RecurrenceMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recurrenceMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ShiftCoverages() TfRotation_ShiftCoveragesPropertyList {
	var returns TfRotation_ShiftCoveragesPropertyList
	_jsii_.Get(
		j,
		"shiftCoverages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ShiftCoveragesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shiftCoveragesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) WeeklySettings() TfRotation_WeeklySettingsPropertyList {
	var returns TfRotation_WeeklySettingsPropertyList
	_jsii_.Get(
		j,
		"weeklySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) WeeklySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklySettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRotation_RecurrencePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRotation_RecurrencePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRotation_RecurrencePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRotation_RecurrencePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm-contacts.TfRotation.RecurrencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRotation_RecurrencePropertyOutputReference_Override(t TfRotation_RecurrencePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm-contacts.TfRotation.RecurrencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetNumberOfOnCalls(val *float64) {
	if err := j.validateSetNumberOfOnCallsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfOnCalls",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetRecurrenceMultiplier(val *float64) {
	if err := j.validateSetRecurrenceMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceMultiplier",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRotation_RecurrencePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) PutDailySettings(value interface{}) {
	if err := t.validatePutDailySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDailySettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) PutMonthlySettings(value interface{}) {
	if err := t.validatePutMonthlySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonthlySettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) PutShiftCoverages(value interface{}) {
	if err := t.validatePutShiftCoveragesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putShiftCoverages",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) PutWeeklySettings(value interface{}) {
	if err := t.validatePutWeeklySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWeeklySettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ResetDailySettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDailySettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ResetMonthlySettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMonthlySettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ResetShiftCoverages() {
	_jsii_.InvokeVoid(
		t,
		"resetShiftCoverages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ResetWeeklySettings() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklySettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRotation_RecurrencePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

