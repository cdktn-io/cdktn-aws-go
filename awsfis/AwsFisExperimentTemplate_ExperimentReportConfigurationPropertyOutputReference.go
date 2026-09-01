package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfis/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference interface {
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
	DataSources() AwsFisExperimentTemplate_DataSourcesPropertyOutputReference
	// Experimental.
	DataSourcesInput() *AwsFisExperimentTemplate_DataSourcesProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsFisExperimentTemplate_ExperimentReportConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFisExperimentTemplate_ExperimentReportConfigurationProperty)
	// Experimental.
	Outputs() AwsFisExperimentTemplate_OutputsPropertyOutputReference
	// Experimental.
	OutputsInput() *AwsFisExperimentTemplate_OutputsProperty
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
	PutDataSources(value *AwsFisExperimentTemplate_DataSourcesProperty)
	// Experimental.
	PutOutputs(value *AwsFisExperimentTemplate_OutputsProperty)
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

// The jsii proxy struct for AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference
type jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSources() AwsFisExperimentTemplate_DataSourcesPropertyOutputReference {
	var returns AwsFisExperimentTemplate_DataSourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) DataSourcesInput() *AwsFisExperimentTemplate_DataSourcesProperty {
	var returns *AwsFisExperimentTemplate_DataSourcesProperty
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InternalValue() *AwsFisExperimentTemplate_ExperimentReportConfigurationProperty {
	var returns *AwsFisExperimentTemplate_ExperimentReportConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Outputs() AwsFisExperimentTemplate_OutputsPropertyOutputReference {
	var returns AwsFisExperimentTemplate_OutputsPropertyOutputReference
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) OutputsInput() *AwsFisExperimentTemplate_OutputsProperty {
	var returns *AwsFisExperimentTemplate_OutputsProperty
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PostExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PreExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fis.AwsFisExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference_Override(a AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fis.AwsFisExperimentTemplate.ExperimentReportConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetInternalValue(val *AwsFisExperimentTemplate_ExperimentReportConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPostExperimentDuration(val *string) {
	if err := j.validateSetPostExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetPreExperimentDuration(val *string) {
	if err := j.validateSetPreExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutDataSources(value *AwsFisExperimentTemplate_DataSourcesProperty) {
	if err := a.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataSources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) PutOutputs(value *AwsFisExperimentTemplate_OutputsProperty) {
	if err := a.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetDataSources() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetOutputs() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPostExperimentDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetPostExperimentDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ResetPreExperimentDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetPreExperimentDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFisExperimentTemplate_ExperimentReportConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

