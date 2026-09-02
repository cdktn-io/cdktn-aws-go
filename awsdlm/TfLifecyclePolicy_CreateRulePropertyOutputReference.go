package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLifecyclePolicy_CreateRulePropertyOutputReference interface {
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
	CronExpression() *string
	// Experimental.
	SetCronExpression(val *string)
	// Experimental.
	CronExpressionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfLifecyclePolicy_CreateRuleProperty
	// Experimental.
	SetInternalValue(val *TfLifecyclePolicy_CreateRuleProperty)
	// Experimental.
	Interval() *float64
	// Experimental.
	SetInterval(val *float64)
	// Experimental.
	IntervalInput() *float64
	// Experimental.
	IntervalUnit() *string
	// Experimental.
	SetIntervalUnit(val *string)
	// Experimental.
	IntervalUnitInput() *string
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	Scripts() TfLifecyclePolicy_ScriptsPropertyOutputReference
	// Experimental.
	ScriptsInput() *TfLifecyclePolicy_ScriptsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Times() *[]*string
	// Experimental.
	SetTimes(val *[]*string)
	// Experimental.
	TimesInput() *[]*string
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
	PutScripts(value *TfLifecyclePolicy_ScriptsProperty)
	// Experimental.
	ResetCronExpression()
	// Experimental.
	ResetInterval()
	// Experimental.
	ResetIntervalUnit()
	// Experimental.
	ResetLocation()
	// Experimental.
	ResetScripts()
	// Experimental.
	ResetTimes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLifecyclePolicy_CreateRulePropertyOutputReference
type jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) InternalValue() *TfLifecyclePolicy_CreateRuleProperty {
	var returns *TfLifecyclePolicy_CreateRuleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Scripts() TfLifecyclePolicy_ScriptsPropertyOutputReference {
	var returns TfLifecyclePolicy_ScriptsPropertyOutputReference
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ScriptsInput() *TfLifecyclePolicy_ScriptsProperty {
	var returns *TfLifecyclePolicy_ScriptsProperty
	_jsii_.Get(
		j,
		"scriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Times() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"times",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) TimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLifecyclePolicy_CreateRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLifecyclePolicy_CreateRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLifecyclePolicy_CreateRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.CreateRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLifecyclePolicy_CreateRulePropertyOutputReference_Override(t TfLifecyclePolicy_CreateRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.CreateRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetCronExpression(val *string) {
	if err := j.validateSetCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetInternalValue(val *TfLifecyclePolicy_CreateRuleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetIntervalUnit(val *string) {
	if err := j.validateSetIntervalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference)SetTimes(val *[]*string) {
	if err := j.validateSetTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"times",
		val,
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) PutScripts(value *TfLifecyclePolicy_ScriptsProperty) {
	if err := t.validatePutScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScripts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetCronExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetCronExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetIntervalUnit() {
	_jsii_.InvokeVoid(
		t,
		"resetIntervalUnit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetScripts() {
	_jsii_.InvokeVoid(
		t,
		"resetScripts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ResetTimes() {
	_jsii_.InvokeVoid(
		t,
		"resetTimes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_CreateRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

