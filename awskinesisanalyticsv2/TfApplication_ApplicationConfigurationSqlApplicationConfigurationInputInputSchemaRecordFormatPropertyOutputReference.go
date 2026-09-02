package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference interface {
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
	InternalValue() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty
	// Experimental.
	SetInternalValue(val *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty)
	// Experimental.
	MappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference
	// Experimental.
	MappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty
	// Experimental.
	RecordFormatType() *string
	// Experimental.
	SetRecordFormatType(val *string)
	// Experimental.
	RecordFormatTypeInput() *string
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
	PutMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference
type jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) InternalValue() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty {
	var returns *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) MappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference {
	var returns TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"mappingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) MappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty {
	var returns *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty
	_jsii_.Get(
		j,
		"mappingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) RecordFormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordFormatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) RecordFormatTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordFormatTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference_Override(t TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetInternalValue(val *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetRecordFormatType(val *string) {
	if err := j.validateSetRecordFormatTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordFormatType",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) PutMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty) {
	if err := t.validatePutMappingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMappingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

