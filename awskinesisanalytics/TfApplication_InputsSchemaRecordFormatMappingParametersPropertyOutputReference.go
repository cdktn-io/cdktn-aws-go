package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference interface {
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
	Csv() TfApplication_InputsSchemaRecordFormatMappingParametersCsvPropertyOutputReference
	// Experimental.
	CsvInput() *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfApplication_InputsSchemaRecordFormatMappingParametersProperty
	// Experimental.
	SetInternalValue(val *TfApplication_InputsSchemaRecordFormatMappingParametersProperty)
	// Experimental.
	Json() TfApplication_InputsSchemaRecordFormatMappingParametersJsonPropertyOutputReference
	// Experimental.
	JsonInput() *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty
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
	PutCsv(value *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty)
	// Experimental.
	PutJson(value *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty)
	// Experimental.
	ResetCsv()
	// Experimental.
	ResetJson()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference
type jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) Csv() TfApplication_InputsSchemaRecordFormatMappingParametersCsvPropertyOutputReference {
	var returns TfApplication_InputsSchemaRecordFormatMappingParametersCsvPropertyOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) CsvInput() *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty {
	var returns *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) InternalValue() *TfApplication_InputsSchemaRecordFormatMappingParametersProperty {
	var returns *TfApplication_InputsSchemaRecordFormatMappingParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) Json() TfApplication_InputsSchemaRecordFormatMappingParametersJsonPropertyOutputReference {
	var returns TfApplication_InputsSchemaRecordFormatMappingParametersJsonPropertyOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) JsonInput() *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty {
	var returns *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.InputsSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference_Override(t TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.InputsSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference)SetInternalValue(val *TfApplication_InputsSchemaRecordFormatMappingParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) PutCsv(value *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty) {
	if err := t.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCsv",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) PutJson(value *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty) {
	if err := t.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJson",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		t,
		"resetCsv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		t,
		"resetJson",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_InputsSchemaRecordFormatMappingParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

