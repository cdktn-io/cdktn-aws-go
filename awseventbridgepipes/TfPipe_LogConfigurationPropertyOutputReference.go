package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_LogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogsLogDestination() TfPipe_CloudwatchLogsLogDestinationPropertyOutputReference
	// Experimental.
	CloudwatchLogsLogDestinationInput() *TfPipe_CloudwatchLogsLogDestinationProperty
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
	FirehoseLogDestination() TfPipe_FirehoseLogDestinationPropertyOutputReference
	// Experimental.
	FirehoseLogDestinationInput() *TfPipe_FirehoseLogDestinationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeExecutionData() *[]*string
	// Experimental.
	SetIncludeExecutionData(val *[]*string)
	// Experimental.
	IncludeExecutionDataInput() *[]*string
	// Experimental.
	InternalValue() *TfPipe_LogConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfPipe_LogConfigurationProperty)
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(val *string)
	// Experimental.
	LevelInput() *string
	// Experimental.
	S3LogDestination() TfPipe_S3LogDestinationPropertyOutputReference
	// Experimental.
	S3LogDestinationInput() *TfPipe_S3LogDestinationProperty
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
	PutCloudwatchLogsLogDestination(value *TfPipe_CloudwatchLogsLogDestinationProperty)
	// Experimental.
	PutFirehoseLogDestination(value *TfPipe_FirehoseLogDestinationProperty)
	// Experimental.
	PutS3LogDestination(value *TfPipe_S3LogDestinationProperty)
	// Experimental.
	ResetCloudwatchLogsLogDestination()
	// Experimental.
	ResetFirehoseLogDestination()
	// Experimental.
	ResetIncludeExecutionData()
	// Experimental.
	ResetS3LogDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_LogConfigurationPropertyOutputReference
type jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) CloudwatchLogsLogDestination() TfPipe_CloudwatchLogsLogDestinationPropertyOutputReference {
	var returns TfPipe_CloudwatchLogsLogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) CloudwatchLogsLogDestinationInput() *TfPipe_CloudwatchLogsLogDestinationProperty {
	var returns *TfPipe_CloudwatchLogsLogDestinationProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) FirehoseLogDestination() TfPipe_FirehoseLogDestinationPropertyOutputReference {
	var returns TfPipe_FirehoseLogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"firehoseLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) FirehoseLogDestinationInput() *TfPipe_FirehoseLogDestinationProperty {
	var returns *TfPipe_FirehoseLogDestinationProperty
	_jsii_.Get(
		j,
		"firehoseLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) IncludeExecutionData() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) IncludeExecutionDataInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) InternalValue() *TfPipe_LogConfigurationProperty {
	var returns *TfPipe_LogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) S3LogDestination() TfPipe_S3LogDestinationPropertyOutputReference {
	var returns TfPipe_S3LogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3LogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) S3LogDestinationInput() *TfPipe_S3LogDestinationProperty {
	var returns *TfPipe_S3LogDestinationProperty
	_jsii_.Get(
		j,
		"s3LogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_LogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_LogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_LogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_LogConfigurationPropertyOutputReference_Override(t TfPipe_LogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetIncludeExecutionData(val *[]*string) {
	if err := j.validateSetIncludeExecutionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExecutionData",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetInternalValue(val *TfPipe_LogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) PutCloudwatchLogsLogDestination(value *TfPipe_CloudwatchLogsLogDestinationProperty) {
	if err := t.validatePutCloudwatchLogsLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogsLogDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) PutFirehoseLogDestination(value *TfPipe_FirehoseLogDestinationProperty) {
	if err := t.validatePutFirehoseLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirehoseLogDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) PutS3LogDestination(value *TfPipe_S3LogDestinationProperty) {
	if err := t.validatePutS3LogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3LogDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ResetCloudwatchLogsLogDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogsLogDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ResetFirehoseLogDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehoseLogDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ResetIncludeExecutionData() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeExecutionData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ResetS3LogDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetS3LogDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_LogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

