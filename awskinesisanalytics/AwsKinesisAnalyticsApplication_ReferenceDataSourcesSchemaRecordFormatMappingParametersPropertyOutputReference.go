package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference interface {
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
	Csv() AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference
	// Experimental.
	CsvInput() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty)
	// Experimental.
	Json() AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonPropertyOutputReference
	// Experimental.
	JsonInput() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty
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
	PutCsv(value *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty)
	// Experimental.
	PutJson(value *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty)
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

// The jsii proxy struct for AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference
type jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) Csv() AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvPropertyOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) CsvInput() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty {
	var returns *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) InternalValue() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty {
	var returns *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) Json() AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonPropertyOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) JsonInput() *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty {
	var returns *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsKinesisAnalyticsApplication.ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference_Override(a AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsKinesisAnalyticsApplication.ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference)SetInternalValue(val *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) PutCsv(value *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty) {
	if err := a.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCsv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) PutJson(value *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty) {
	if err := a.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJson",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		a,
		"resetCsv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		a,
		"resetJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

