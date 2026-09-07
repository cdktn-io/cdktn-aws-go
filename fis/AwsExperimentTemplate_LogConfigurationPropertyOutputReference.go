package fis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fis/jsii"

	"github.com/cdktn-io/cdktn-aws-go/fis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsExperimentTemplate_LogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogsConfiguration() AwsExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference
	// Experimental.
	CloudwatchLogsConfigurationInput() *AwsExperimentTemplate_CloudwatchLogsConfigurationProperty
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
	InternalValue() *AwsExperimentTemplate_LogConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsExperimentTemplate_LogConfigurationProperty)
	// Experimental.
	LogSchemaVersion() *float64
	// Experimental.
	SetLogSchemaVersion(val *float64)
	// Experimental.
	LogSchemaVersionInput() *float64
	// Experimental.
	S3Configuration() AwsExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *AwsExperimentTemplate_LogConfigurationS3ConfigurationProperty
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
	PutCloudwatchLogsConfiguration(value *AwsExperimentTemplate_CloudwatchLogsConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *AwsExperimentTemplate_LogConfigurationS3ConfigurationProperty)
	// Experimental.
	ResetCloudwatchLogsConfiguration()
	// Experimental.
	ResetS3Configuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsExperimentTemplate_LogConfigurationPropertyOutputReference
type jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) CloudwatchLogsConfiguration() AwsExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference {
	var returns AwsExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) CloudwatchLogsConfigurationInput() *AwsExperimentTemplate_CloudwatchLogsConfigurationProperty {
	var returns *AwsExperimentTemplate_CloudwatchLogsConfigurationProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) InternalValue() *AwsExperimentTemplate_LogConfigurationProperty {
	var returns *AwsExperimentTemplate_LogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) LogSchemaVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) LogSchemaVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) S3Configuration() AwsExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference {
	var returns AwsExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) S3ConfigurationInput() *AwsExperimentTemplate_LogConfigurationS3ConfigurationProperty {
	var returns *AwsExperimentTemplate_LogConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsExperimentTemplate_LogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsExperimentTemplate_LogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsExperimentTemplate_LogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fis.AwsExperimentTemplate.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsExperimentTemplate_LogConfigurationPropertyOutputReference_Override(a AwsExperimentTemplate_LogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fis.AwsExperimentTemplate.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetInternalValue(val *AwsExperimentTemplate_LogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetLogSchemaVersion(val *float64) {
	if err := j.validateSetLogSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) PutCloudwatchLogsConfiguration(value *AwsExperimentTemplate_CloudwatchLogsConfigurationProperty) {
	if err := a.validatePutCloudwatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) PutS3Configuration(value *AwsExperimentTemplate_LogConfigurationS3ConfigurationProperty) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ResetCloudwatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsExperimentTemplate_LogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

