package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference interface {
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
	CsvMappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersPropertyOutputReference
	// Experimental.
	CsvMappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty
	// Experimental.
	SetInternalValue(val *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty)
	// Experimental.
	JsonMappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersPropertyOutputReference
	// Experimental.
	JsonMappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersProperty
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
	PutCsvMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty)
	// Experimental.
	PutJsonMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersProperty)
	// Experimental.
	ResetCsvMappingParameters()
	// Experimental.
	ResetJsonMappingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference
type jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) CsvMappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersPropertyOutputReference {
	var returns TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"csvMappingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) CsvMappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty {
	var returns *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty
	_jsii_.Get(
		j,
		"csvMappingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) InternalValue() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty {
	var returns *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) JsonMappingParameters() TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersPropertyOutputReference {
	var returns TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"jsonMappingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) JsonMappingParametersInput() *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersProperty {
	var returns *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersProperty
	_jsii_.Get(
		j,
		"jsonMappingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference_Override(t TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference)SetInternalValue(val *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) PutCsvMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty) {
	if err := t.validatePutCsvMappingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCsvMappingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) PutJsonMappingParameters(value *TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParametersProperty) {
	if err := t.validatePutJsonMappingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJsonMappingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ResetCsvMappingParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetCsvMappingParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ResetJsonMappingParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonMappingParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

