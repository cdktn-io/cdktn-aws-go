package awstimestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScheduledQuery_TimestreamConfigurationPropertyOutputReference interface {
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
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	DimensionMapping() TfScheduledQuery_DimensionMappingPropertyList
	// Experimental.
	DimensionMappingInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MeasureNameColumn() *string
	// Experimental.
	SetMeasureNameColumn(val *string)
	// Experimental.
	MeasureNameColumnInput() *string
	// Experimental.
	MixedMeasureMapping() TfScheduledQuery_MixedMeasureMappingPropertyList
	// Experimental.
	MixedMeasureMappingInput() interface{}
	// Experimental.
	MultiMeasureMappings() TfScheduledQuery_MultiMeasureMappingsPropertyList
	// Experimental.
	MultiMeasureMappingsInput() interface{}
	// Experimental.
	TableName() *string
	// Experimental.
	SetTableName(val *string)
	// Experimental.
	TableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeColumn() *string
	// Experimental.
	SetTimeColumn(val *string)
	// Experimental.
	TimeColumnInput() *string
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
	PutDimensionMapping(value interface{})
	// Experimental.
	PutMixedMeasureMapping(value interface{})
	// Experimental.
	PutMultiMeasureMappings(value interface{})
	// Experimental.
	ResetDimensionMapping()
	// Experimental.
	ResetMeasureNameColumn()
	// Experimental.
	ResetMixedMeasureMapping()
	// Experimental.
	ResetMultiMeasureMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfScheduledQuery_TimestreamConfigurationPropertyOutputReference
type jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) DimensionMapping() TfScheduledQuery_DimensionMappingPropertyList {
	var returns TfScheduledQuery_DimensionMappingPropertyList
	_jsii_.Get(
		j,
		"dimensionMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) DimensionMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MeasureNameColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MeasureNameColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MixedMeasureMapping() TfScheduledQuery_MixedMeasureMappingPropertyList {
	var returns TfScheduledQuery_MixedMeasureMappingPropertyList
	_jsii_.Get(
		j,
		"mixedMeasureMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MixedMeasureMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mixedMeasureMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MultiMeasureMappings() TfScheduledQuery_MultiMeasureMappingsPropertyList {
	var returns TfScheduledQuery_MultiMeasureMappingsPropertyList
	_jsii_.Get(
		j,
		"multiMeasureMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) MultiMeasureMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiMeasureMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TimeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) TimeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeColumnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScheduledQuery_TimestreamConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfScheduledQuery_TimestreamConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScheduledQuery_TimestreamConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.TimestreamConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScheduledQuery_TimestreamConfigurationPropertyOutputReference_Override(t TfScheduledQuery_TimestreamConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.TimestreamConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetMeasureNameColumn(val *string) {
	if err := j.validateSetMeasureNameColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureNameColumn",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTimeColumn(val *string) {
	if err := j.validateSetTimeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeColumn",
		val,
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutDimensionMapping(value interface{}) {
	if err := t.validatePutDimensionMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimensionMapping",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutMixedMeasureMapping(value interface{}) {
	if err := t.validatePutMixedMeasureMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMixedMeasureMapping",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutMultiMeasureMappings(value interface{}) {
	if err := t.validatePutMultiMeasureMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMultiMeasureMappings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetDimensionMapping() {
	_jsii_.InvokeVoid(
		t,
		"resetDimensionMapping",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMeasureNameColumn() {
	_jsii_.InvokeVoid(
		t,
		"resetMeasureNameColumn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMixedMeasureMapping() {
	_jsii_.InvokeVoid(
		t,
		"resetMixedMeasureMapping",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMultiMeasureMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiMeasureMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_TimestreamConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

