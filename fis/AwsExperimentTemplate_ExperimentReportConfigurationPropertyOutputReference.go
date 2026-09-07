package fis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fis/jsii"

	"github.com/cdktn-io/cdktn-aws-go/fis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference interface {
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
	DataSources() AwsExperimentTemplate_DataSourcesPropertyOutputReference
	// Experimental.
	DataSourcesInput() *AwsExperimentTemplate_DataSourcesProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsExperimentTemplate_ExperimentReportConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsExperimentTemplate_ExperimentReportConfigurationProperty)
	// Experimental.
	Outputs() AwsExperimentTemplate_OutputsPropertyOutputReference
	// Experimental.
	OutputsInput() *AwsExperimentTemplate_OutputsProperty
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
	PutDataSources(value *AwsExperimentTemplate_DataSourcesProperty)
	// Experimental.
	PutOutputs(value *AwsExperimentTemplate_OutputsProperty)
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

// The jsii proxy struct for AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference
type jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSources() AwsExperimentTemplate_DataSourcesPropertyOutputReference {
	var returns AwsExperimentTemplate_DataSourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSourcesInput() *AwsExperimentTemplate_DataSourcesProperty {
	var returns *AwsExperimentTemplate_DataSourcesProperty
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InternalValue() *AwsExperimentTemplate_ExperimentReportConfigurationProperty {
	var returns *AwsExperimentTemplate_ExperimentReportConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Outputs() AwsExperimentTemplate_OutputsPropertyOutputReference {
	var returns AwsExperimentTemplate_OutputsPropertyOutputReference
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) OutputsInput() *AwsExperimentTemplate_OutputsProperty {
	var returns *AwsExperimentTemplate_OutputsProperty
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fis.AwsExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference_Override(a AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fis.AwsExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetInternalValue(val *AwsExperimentTemplate_ExperimentReportConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPostExperimentDuration(val *string) {
	if err := j.validateSetPostExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPreExperimentDuration(val *string) {
	if err := j.validateSetPreExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutDataSources(value *AwsExperimentTemplate_DataSourcesProperty) {
	if err := a.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataSources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutOutputs(value *AwsExperimentTemplate_OutputsProperty) {
	if err := a.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetDataSources() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetOutputs() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPostExperimentDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetPostExperimentDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPreExperimentDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetPreExperimentDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

