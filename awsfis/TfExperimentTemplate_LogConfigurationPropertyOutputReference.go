package awsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfis/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExperimentTemplate_LogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogsConfiguration() TfExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference
	// Experimental.
	CloudwatchLogsConfigurationInput() *TfExperimentTemplate_CloudwatchLogsConfigurationProperty
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
	InternalValue() *TfExperimentTemplate_LogConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfExperimentTemplate_LogConfigurationProperty)
	// Experimental.
	LogSchemaVersion() *float64
	// Experimental.
	SetLogSchemaVersion(val *float64)
	// Experimental.
	LogSchemaVersionInput() *float64
	// Experimental.
	S3Configuration() TfExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty
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
	PutCloudwatchLogsConfiguration(value *TfExperimentTemplate_CloudwatchLogsConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty)
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

// The jsii proxy struct for TfExperimentTemplate_LogConfigurationPropertyOutputReference
type jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) CloudwatchLogsConfiguration() TfExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference {
	var returns TfExperimentTemplate_CloudwatchLogsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) CloudwatchLogsConfigurationInput() *TfExperimentTemplate_CloudwatchLogsConfigurationProperty {
	var returns *TfExperimentTemplate_CloudwatchLogsConfigurationProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) InternalValue() *TfExperimentTemplate_LogConfigurationProperty {
	var returns *TfExperimentTemplate_LogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) LogSchemaVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) LogSchemaVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) S3Configuration() TfExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference {
	var returns TfExperimentTemplate_LogConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) S3ConfigurationInput() *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty {
	var returns *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfExperimentTemplate_LogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfExperimentTemplate_LogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfExperimentTemplate_LogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fis.TfExperimentTemplate.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfExperimentTemplate_LogConfigurationPropertyOutputReference_Override(t TfExperimentTemplate_LogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fis.TfExperimentTemplate.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetInternalValue(val *TfExperimentTemplate_LogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetLogSchemaVersion(val *float64) {
	if err := j.validateSetLogSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) PutCloudwatchLogsConfiguration(value *TfExperimentTemplate_CloudwatchLogsConfigurationProperty) {
	if err := t.validatePutCloudwatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) PutS3Configuration(value *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ResetCloudwatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfExperimentTemplate_LogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

