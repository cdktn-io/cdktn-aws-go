package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference interface {
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
	InternalValue() *TfApplication_ReferenceDataSourcesSchemaProperty
	// Experimental.
	SetInternalValue(val *TfApplication_ReferenceDataSourcesSchemaProperty)
	// Experimental.
	RecordColumns() TfApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList
	// Experimental.
	RecordColumnsInput() interface{}
	// Experimental.
	RecordEncoding() *string
	// Experimental.
	SetRecordEncoding(val *string)
	// Experimental.
	RecordEncodingInput() *string
	// Experimental.
	RecordFormat() TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference
	// Experimental.
	RecordFormatInput() *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty
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
	PutRecordFormat(value *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty)
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

// The jsii proxy struct for TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference
type jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InternalValue() *TfApplication_ReferenceDataSourcesSchemaProperty {
	var returns *TfApplication_ReferenceDataSourcesSchemaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordColumns() TfApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList {
	var returns TfApplication_ReferenceDataSourcesSchemaRecordColumnsPropertyList
	_jsii_.Get(
		j,
		"recordColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordFormat() TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference {
	var returns TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference
	_jsii_.Get(
		j,
		"recordFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) RecordFormatInput() *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty {
	var returns *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty
	_jsii_.Get(
		j,
		"recordFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_ReferenceDataSourcesSchemaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_ReferenceDataSourcesSchemaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.ReferenceDataSourcesSchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_ReferenceDataSourcesSchemaPropertyOutputReference_Override(t TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.ReferenceDataSourcesSchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetInternalValue(val *TfApplication_ReferenceDataSourcesSchemaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetRecordEncoding(val *string) {
	if err := j.validateSetRecordEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordEncoding",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) PutRecordColumns(value interface{}) {
	if err := t.validatePutRecordColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) PutRecordFormat(value *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty) {
	if err := t.validatePutRecordFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordFormat",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ResetRecordEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

