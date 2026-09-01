package awstimestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference interface {
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
	MeasureName() *string
	// Experimental.
	SetMeasureName(val *string)
	// Experimental.
	MeasureNameInput() *string
	// Experimental.
	MeasureValueType() *string
	// Experimental.
	SetMeasureValueType(val *string)
	// Experimental.
	MeasureValueTypeInput() *string
	// Experimental.
	MultiMeasureAttributeMapping() AwsTimestreamqueryScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyList
	// Experimental.
	MultiMeasureAttributeMappingInput() interface{}
	// Experimental.
	SourceColumn() *string
	// Experimental.
	SetSourceColumn(val *string)
	// Experimental.
	SourceColumnInput() *string
	// Experimental.
	TargetMeasureName() *string
	// Experimental.
	SetTargetMeasureName(val *string)
	// Experimental.
	TargetMeasureNameInput() *string
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
	PutMultiMeasureAttributeMapping(value interface{})
	// Experimental.
	ResetMeasureName()
	// Experimental.
	ResetMultiMeasureAttributeMapping()
	// Experimental.
	ResetSourceColumn()
	// Experimental.
	ResetTargetMeasureName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference
type jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MeasureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MeasureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MeasureValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MeasureValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureValueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MultiMeasureAttributeMapping() AwsTimestreamqueryScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyList {
	var returns AwsTimestreamqueryScheduledQuery_TargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingPropertyList
	_jsii_.Get(
		j,
		"multiMeasureAttributeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) MultiMeasureAttributeMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiMeasureAttributeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) SourceColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) SourceColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) TargetMeasureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMeasureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) TargetMeasureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetMeasureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsTimestreamqueryScheduledQuery.MixedMeasureMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference_Override(a AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsTimestreamqueryScheduledQuery.MixedMeasureMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetMeasureName(val *string) {
	if err := j.validateSetMeasureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureName",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetMeasureValueType(val *string) {
	if err := j.validateSetMeasureValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureValueType",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetSourceColumn(val *string) {
	if err := j.validateSetSourceColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceColumn",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetTargetMeasureName(val *string) {
	if err := j.validateSetTargetMeasureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetMeasureName",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) PutMultiMeasureAttributeMapping(value interface{}) {
	if err := a.validatePutMultiMeasureAttributeMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiMeasureAttributeMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ResetMeasureName() {
	_jsii_.InvokeVoid(
		a,
		"resetMeasureName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ResetMultiMeasureAttributeMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiMeasureAttributeMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ResetSourceColumn() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceColumn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ResetTargetMeasureName() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetMeasureName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_MixedMeasureMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

