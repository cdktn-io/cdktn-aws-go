package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference interface {
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
	InternalValue() *AwsDlmLifecyclePolicy_CreateRuleProperty
	// Experimental.
	SetInternalValue(val *AwsDlmLifecyclePolicy_CreateRuleProperty)
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
	Scripts() AwsDlmLifecyclePolicy_ScriptsPropertyOutputReference
	// Experimental.
	ScriptsInput() *AwsDlmLifecyclePolicy_ScriptsProperty
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
	PutScripts(value *AwsDlmLifecyclePolicy_ScriptsProperty)
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

// The jsii proxy struct for AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference
type jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) CronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) InternalValue() *AwsDlmLifecyclePolicy_CreateRuleProperty {
	var returns *AwsDlmLifecyclePolicy_CreateRuleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Scripts() AwsDlmLifecyclePolicy_ScriptsPropertyOutputReference {
	var returns AwsDlmLifecyclePolicy_ScriptsPropertyOutputReference
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ScriptsInput() *AwsDlmLifecyclePolicy_ScriptsProperty {
	var returns *AwsDlmLifecyclePolicy_ScriptsProperty
	_jsii_.Get(
		j,
		"scriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Times() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"times",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) TimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"timesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDlmLifecyclePolicy_CreateRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDlmLifecyclePolicy_CreateRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.CreateRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDlmLifecyclePolicy_CreateRulePropertyOutputReference_Override(a AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.CreateRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetCronExpression(val *string) {
	if err := j.validateSetCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetInternalValue(val *AwsDlmLifecyclePolicy_CreateRuleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetIntervalUnit(val *string) {
	if err := j.validateSetIntervalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference)SetTimes(val *[]*string) {
	if err := j.validateSetTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"times",
		val,
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) PutScripts(value *AwsDlmLifecyclePolicy_ScriptsProperty) {
	if err := a.validatePutScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScripts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetCronExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetCronExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetIntervalUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetIntervalUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetScripts() {
	_jsii_.InvokeVoid(
		a,
		"resetScripts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ResetTimes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_CreateRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

