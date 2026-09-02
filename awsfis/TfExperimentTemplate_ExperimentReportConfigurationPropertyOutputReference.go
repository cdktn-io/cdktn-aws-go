package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfis/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference interface {
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
	DataSources() TfExperimentTemplate_DataSourcesPropertyOutputReference
	// Experimental.
	DataSourcesInput() *TfExperimentTemplate_DataSourcesProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfExperimentTemplate_ExperimentReportConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfExperimentTemplate_ExperimentReportConfigurationProperty)
	// Experimental.
	Outputs() TfExperimentTemplate_OutputsPropertyOutputReference
	// Experimental.
	OutputsInput() *TfExperimentTemplate_OutputsProperty
	// Experimental.
	PostExperimentDuration() *string
	// Experimental.
	SetPostExperimentDuration(val *string)
	// Experimental.
	PostExperimentDurationInput() *string
	// Experimental.
	PreExperimentDuration() *string
	// Experimental.
	SetPreExperimentDuration(val *string)
	// Experimental.
	PreExperimentDurationInput() *string
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
	PutDataSources(value *TfExperimentTemplate_DataSourcesProperty)
	// Experimental.
	PutOutputs(value *TfExperimentTemplate_OutputsProperty)
	// Experimental.
	ResetDataSources()
	// Experimental.
	ResetOutputs()
	// Experimental.
	ResetPostExperimentDuration()
	// Experimental.
	ResetPreExperimentDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference
type jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSources() TfExperimentTemplate_DataSourcesPropertyOutputReference {
	var returns TfExperimentTemplate_DataSourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSourcesInput() *TfExperimentTemplate_DataSourcesProperty {
	var returns *TfExperimentTemplate_DataSourcesProperty
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InternalValue() *TfExperimentTemplate_ExperimentReportConfigurationProperty {
	var returns *TfExperimentTemplate_ExperimentReportConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Outputs() TfExperimentTemplate_OutputsPropertyOutputReference {
	var returns TfExperimentTemplate_OutputsPropertyOutputReference
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) OutputsInput() *TfExperimentTemplate_OutputsProperty {
	var returns *TfExperimentTemplate_OutputsProperty
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fis.TfExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference_Override(t TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fis.TfExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetInternalValue(val *TfExperimentTemplate_ExperimentReportConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPostExperimentDuration(val *string) {
	if err := j.validateSetPostExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPreExperimentDuration(val *string) {
	if err := j.validateSetPreExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutDataSources(value *TfExperimentTemplate_DataSourcesProperty) {
	if err := t.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataSources",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutOutputs(value *TfExperimentTemplate_OutputsProperty) {
	if err := t.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetDataSources() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetOutputs() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPostExperimentDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetPostExperimentDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPreExperimentDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetPreExperimentDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

