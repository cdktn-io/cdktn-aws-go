package timestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MeasureValueType() *string
	// Experimental.
	SetMeasureValueType(val *string)
	// Experimental.
	MeasureValueTypeInput() *string
	// Experimental.
	SourceColumn() *string
	// Experimental.
	SetSourceColumn(val *string)
	// Experimental.
	SourceColumnInput() *string
	// Experimental.
	TargetMultiMeasureAttributeName() *string
	// Experimental.
	SetTargetMultiMeasureAttributeName(val *string)
	// Experimental.
	TargetMultiMeasureAttributeNameInput() *string
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
	ResetTargetMultiMeasureAttributeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference
type jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) MeasureValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) MeasureValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) SourceColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) SourceColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) TargetMultiMeasureAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) TargetMultiMeasureAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference_Override(a AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetMeasureValueType(val *string) {
	if err := j.validateSetMeasureValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureValueType",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetSourceColumn(val *string) {
	if err := j.validateSetSourceColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceColumn",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetTargetMultiMeasureAttributeName(val *string) {
	if err := j.validateSetTargetMultiMeasureAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetMultiMeasureAttributeName",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) ResetTargetMultiMeasureAttributeName() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetMultiMeasureAttributeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

