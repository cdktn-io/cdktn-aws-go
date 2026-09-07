package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_LogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogsLogDestination() AwsPipe_CloudwatchLogsLogDestinationPropertyOutputReference
	// Experimental.
	CloudwatchLogsLogDestinationInput() *AwsPipe_CloudwatchLogsLogDestinationProperty
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
	FirehoseLogDestination() AwsPipe_FirehoseLogDestinationPropertyOutputReference
	// Experimental.
	FirehoseLogDestinationInput() *AwsPipe_FirehoseLogDestinationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeExecutionData() *[]*string
	// Experimental.
	SetIncludeExecutionData(val *[]*string)
	// Experimental.
	IncludeExecutionDataInput() *[]*string
	// Experimental.
	InternalValue() *AwsPipe_LogConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_LogConfigurationProperty)
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(val *string)
	// Experimental.
	LevelInput() *string
	// Experimental.
	S3LogDestination() AwsPipe_S3LogDestinationPropertyOutputReference
	// Experimental.
	S3LogDestinationInput() *AwsPipe_S3LogDestinationProperty
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
	PutCloudwatchLogsLogDestination(value *AwsPipe_CloudwatchLogsLogDestinationProperty)
	// Experimental.
	PutFirehoseLogDestination(value *AwsPipe_FirehoseLogDestinationProperty)
	// Experimental.
	PutS3LogDestination(value *AwsPipe_S3LogDestinationProperty)
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

// The jsii proxy struct for AwsPipe_LogConfigurationPropertyOutputReference
type jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) CloudwatchLogsLogDestination() AwsPipe_CloudwatchLogsLogDestinationPropertyOutputReference {
	var returns AwsPipe_CloudwatchLogsLogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) CloudwatchLogsLogDestinationInput() *AwsPipe_CloudwatchLogsLogDestinationProperty {
	var returns *AwsPipe_CloudwatchLogsLogDestinationProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) FirehoseLogDestination() AwsPipe_FirehoseLogDestinationPropertyOutputReference {
	var returns AwsPipe_FirehoseLogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"firehoseLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) FirehoseLogDestinationInput() *AwsPipe_FirehoseLogDestinationProperty {
	var returns *AwsPipe_FirehoseLogDestinationProperty
	_jsii_.Get(
		j,
		"firehoseLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) IncludeExecutionData() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) IncludeExecutionDataInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) InternalValue() *AwsPipe_LogConfigurationProperty {
	var returns *AwsPipe_LogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) S3LogDestination() AwsPipe_S3LogDestinationPropertyOutputReference {
	var returns AwsPipe_S3LogDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3LogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) S3LogDestinationInput() *AwsPipe_S3LogDestinationProperty {
	var returns *AwsPipe_S3LogDestinationProperty
	_jsii_.Get(
		j,
		"s3LogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_LogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_LogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_LogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_LogConfigurationPropertyOutputReference_Override(a AwsPipe_LogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetIncludeExecutionData(val *[]*string) {
	if err := j.validateSetIncludeExecutionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExecutionData",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetInternalValue(val *AwsPipe_LogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) PutCloudwatchLogsLogDestination(value *AwsPipe_CloudwatchLogsLogDestinationProperty) {
	if err := a.validatePutCloudwatchLogsLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogsLogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) PutFirehoseLogDestination(value *AwsPipe_FirehoseLogDestinationProperty) {
	if err := a.validatePutFirehoseLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehoseLogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) PutS3LogDestination(value *AwsPipe_S3LogDestinationProperty) {
	if err := a.validatePutS3LogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3LogDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ResetCloudwatchLogsLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogsLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ResetFirehoseLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ResetIncludeExecutionData() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeExecutionData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ResetS3LogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetS3LogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_LogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

