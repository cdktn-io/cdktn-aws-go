package awscloudwatchevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchevidently/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchevidently/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference interface {
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
	EntityIdKey() *string
	// Experimental.
	SetEntityIdKey(val *string)
	// Experimental.
	EntityIdKeyInput() *string
	// Experimental.
	EventPattern() *string
	// Experimental.
	SetEventPattern(val *string)
	// Experimental.
	EventPatternInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEvidentlyLaunch_MetricDefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsEvidentlyLaunch_MetricDefinitionProperty)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnitLabel() *string
	// Experimental.
	SetUnitLabel(val *string)
	// Experimental.
	UnitLabelInput() *string
	// Experimental.
	ValueKey() *string
	// Experimental.
	SetValueKey(val *string)
	// Experimental.
	ValueKeyInput() *string
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
	ResetEventPattern()
	// Experimental.
	ResetUnitLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference
type jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) EntityIdKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) EntityIdKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) EventPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) EventPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) InternalValue() *AwsEvidentlyLaunch_MetricDefinitionProperty {
	var returns *AwsEvidentlyLaunch_MetricDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) UnitLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) UnitLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ValueKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ValueKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKeyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEvidentlyLaunch_MetricDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-evidently.AwsEvidentlyLaunch.MetricDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference_Override(a AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-evidently.AwsEvidentlyLaunch.MetricDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetEntityIdKey(val *string) {
	if err := j.validateSetEntityIdKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityIdKey",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetEventPattern(val *string) {
	if err := j.validateSetEventPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventPattern",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetInternalValue(val *AwsEvidentlyLaunch_MetricDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetUnitLabel(val *string) {
	if err := j.validateSetUnitLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitLabel",
		val,
	)
}

func (j *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference)SetValueKey(val *string) {
	if err := j.validateSetValueKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueKey",
		val,
	)
}

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ResetEventPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetEventPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ResetUnitLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetUnitLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEvidentlyLaunch_MetricDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

