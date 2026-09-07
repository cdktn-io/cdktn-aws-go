package kinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference interface {
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
	InternalValue() *AwsApplication_ReferenceDataSourcesSchemaProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_ReferenceDataSourcesSchemaProperty)
	// Experimental.
	RecordColumns() AwsApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList
	// Experimental.
	RecordColumnsInput() interface{}
	// Experimental.
	RecordEncoding() *string
	// Experimental.
	SetRecordEncoding(val *string)
	// Experimental.
	RecordEncodingInput() *string
	// Experimental.
	RecordFormat() AwsApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference
	// Experimental.
	RecordFormatInput() *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty
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
	PutRecordColumns(value interface{})
	// Experimental.
	PutRecordFormat(value *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty)
	// Experimental.
	ResetRecordEncoding()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference
type jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InternalValue() *AwsApplication_ReferenceDataSourcesSchemaProperty {
	var returns *AwsApplication_ReferenceDataSourcesSchemaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordColumns() AwsApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList {
	var returns AwsApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList
	_jsii_.Get(
		j,
		"recordColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordFormat() AwsApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference {
	var returns AwsApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference
	_jsii_.Get(
		j,
		"recordFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordFormatInput() *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty {
	var returns *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty
	_jsii_.Get(
		j,
		"recordFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_ReferenceDataSourcesSchemaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.ReferenceDataSourcesSchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference_Override(a AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.ReferenceDataSourcesSchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetInternalValue(val *AwsApplication_ReferenceDataSourcesSchemaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetRecordEncoding(val *string) {
	if err := j.validateSetRecordEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordEncoding",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) PutRecordColumns(value interface{}) {
	if err := a.validatePutRecordColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) PutRecordFormat(value *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty) {
	if err := a.validatePutRecordFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordFormat",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ResetRecordEncoding() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordEncoding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

