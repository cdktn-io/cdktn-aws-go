package awstimestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference interface {
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

// The jsii proxy struct for TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference
type jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) MeasureValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) MeasureValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) SourceColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) SourceColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) TargetMultiMeasureAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) TargetMultiMeasureAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMultiMeasureAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference_Override(t TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetMeasureValueType(val *string) {
	if err := j.validateSetMeasureValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureValueType",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetSourceColumn(val *string) {
	if err := j.validateSetSourceColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceColumn",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetTargetMultiMeasureAttributeName(val *string) {
	if err := j.validateSetTargetMultiMeasureAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetMultiMeasureAttributeName",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) ResetTargetMultiMeasureAttributeName() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetMultiMeasureAttributeName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

