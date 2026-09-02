package awsivschat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsivschat/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsivschat/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() TfLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *TfLoggingConfiguration_CloudwatchLogsProperty
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
	Firehose() TfLoggingConfiguration_FirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *TfLoggingConfiguration_FirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfLoggingConfiguration_DestinationConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfLoggingConfiguration_DestinationConfigurationProperty)
	// Experimental.
	S3() TfLoggingConfiguration_S3PropertyOutputReference
	// Experimental.
	S3Input() *TfLoggingConfiguration_S3Property
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
	PutCloudwatchLogs(value *TfLoggingConfiguration_CloudwatchLogsProperty)
	// Experimental.
	PutFirehose(value *TfLoggingConfiguration_FirehoseProperty)
	// Experimental.
	PutS3(value *TfLoggingConfiguration_S3Property)
	// Experimental.
	ResetCloudwatchLogs()
	// Experimental.
	ResetFirehose()
	// Experimental.
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference
type jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) CloudwatchLogs() TfLoggingConfiguration_CloudwatchLogsPropertyOutputReference {
	var returns TfLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) CloudwatchLogsInput() *TfLoggingConfiguration_CloudwatchLogsProperty {
	var returns *TfLoggingConfiguration_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) Firehose() TfLoggingConfiguration_FirehosePropertyOutputReference {
	var returns TfLoggingConfiguration_FirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) FirehoseInput() *TfLoggingConfiguration_FirehoseProperty {
	var returns *TfLoggingConfiguration_FirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) InternalValue() *TfLoggingConfiguration_DestinationConfigurationProperty {
	var returns *TfLoggingConfiguration_DestinationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) S3() TfLoggingConfiguration_S3PropertyOutputReference {
	var returns TfLoggingConfiguration_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) S3Input() *TfLoggingConfiguration_S3Property {
	var returns *TfLoggingConfiguration_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLoggingConfiguration_DestinationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLoggingConfiguration_DestinationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ivs-chat.TfLoggingConfiguration.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLoggingConfiguration_DestinationConfigurationPropertyOutputReference_Override(t TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ivs-chat.TfLoggingConfiguration.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference)SetInternalValue(val *TfLoggingConfiguration_DestinationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) PutCloudwatchLogs(value *TfLoggingConfiguration_CloudwatchLogsProperty) {
	if err := t.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) PutFirehose(value *TfLoggingConfiguration_FirehoseProperty) {
	if err := t.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirehose",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) PutS3(value *TfLoggingConfiguration_S3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehose",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLoggingConfiguration_DestinationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

