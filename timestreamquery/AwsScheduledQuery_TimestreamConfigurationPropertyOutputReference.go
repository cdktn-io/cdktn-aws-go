package timestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference interface {
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
	DimensionMapping() AwsScheduledQuery_DimensionMappingPropertyList
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
	MixedMeasureMapping() AwsScheduledQuery_MixedMeasureMappingPropertyList
	// Experimental.
	MixedMeasureMappingInput() interface{}
	// Experimental.
	MultiMeasureMappings() AwsScheduledQuery_MultiMeasureMappingsPropertyList
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

// The jsii proxy struct for AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference
type jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) DimensionMapping() AwsScheduledQuery_DimensionMappingPropertyList {
	var returns AwsScheduledQuery_DimensionMappingPropertyList
	_jsii_.Get(
		j,
		"dimensionMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) DimensionMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MeasureNameColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MeasureNameColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureNameColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MixedMeasureMapping() AwsScheduledQuery_MixedMeasureMappingPropertyList {
	var returns AwsScheduledQuery_MixedMeasureMappingPropertyList
	_jsii_.Get(
		j,
		"mixedMeasureMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MixedMeasureMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mixedMeasureMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MultiMeasureMappings() AwsScheduledQuery_MultiMeasureMappingsPropertyList {
	var returns AwsScheduledQuery_MultiMeasureMappingsPropertyList
	_jsii_.Get(
		j,
		"multiMeasureMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) MultiMeasureMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiMeasureMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TimeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) TimeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeColumnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledQuery_TimestreamConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledQuery_TimestreamConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.TimestreamConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledQuery_TimestreamConfigurationPropertyOutputReference_Override(a AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.TimestreamConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetMeasureNameColumn(val *string) {
	if err := j.validateSetMeasureNameColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureNameColumn",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference)SetTimeColumn(val *string) {
	if err := j.validateSetTimeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeColumn",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutDimensionMapping(value interface{}) {
	if err := a.validatePutDimensionMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDimensionMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutMixedMeasureMapping(value interface{}) {
	if err := a.validatePutMixedMeasureMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMixedMeasureMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) PutMultiMeasureMappings(value interface{}) {
	if err := a.validatePutMultiMeasureMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiMeasureMappings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetDimensionMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensionMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMeasureNameColumn() {
	_jsii_.InvokeVoid(
		a,
		"resetMeasureNameColumn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMixedMeasureMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetMixedMeasureMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ResetMultiMeasureMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiMeasureMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_TimestreamConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

