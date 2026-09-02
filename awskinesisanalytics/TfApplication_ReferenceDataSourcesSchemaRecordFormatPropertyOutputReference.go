package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference interface {
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
	InternalValue() *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty
	// Experimental.
	SetInternalValue(val *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty)
	// Experimental.
	MappingParameters() TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference
	// Experimental.
	MappingParametersInput() *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty
	// Experimental.
	RecordFormatType() *string
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
	PutMappingParameters(value *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty)
	// Experimental.
	ResetMappingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference
type jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) InternalValue() *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty {
	var returns *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) MappingParameters() TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference {
	var returns TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"mappingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) MappingParametersInput() *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty {
	var returns *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty
	_jsii_.Get(
		j,
		"mappingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) RecordFormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordFormatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference_Override(t TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference)SetInternalValue(val *TfApplication_ReferenceDataSourcesSchemaRecordFormatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) PutMappingParameters(value *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty) {
	if err := t.validatePutMappingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMappingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) ResetMappingParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetMappingParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_ReferenceDataSourcesSchemaRecordFormatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

