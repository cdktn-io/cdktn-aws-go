package kinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference interface {
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
	InternalValue() *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty)
	// Experimental.
	RecordColumnDelimiter() *string
	// Experimental.
	SetRecordColumnDelimiter(val *string)
	// Experimental.
	RecordColumnDelimiterInput() *string
	// Experimental.
	RecordRowDelimiter() *string
	// Experimental.
	SetRecordRowDelimiter(val *string)
	// Experimental.
	RecordRowDelimiterInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference
type jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) InternalValue() *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty {
	var returns *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) RecordColumnDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordColumnDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) RecordColumnDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordColumnDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) RecordRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) RecordRowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordRowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference_Override(a AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetInternalValue(val *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetRecordColumnDelimiter(val *string) {
	if err := j.validateSetRecordColumnDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordColumnDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetRecordRowDelimiter(val *string) {
	if err := j.validateSetRecordRowDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordRowDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

