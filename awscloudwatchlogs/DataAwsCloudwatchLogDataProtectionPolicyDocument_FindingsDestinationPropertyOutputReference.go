package awscloudwatchlogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsProperty
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
	Firehose() DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty
	// Experimental.
	SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty)
	// Experimental.
	S3() DataAwsCloudwatchLogDataProtectionPolicyDocument_S3PropertyOutputReference
	// Experimental.
	S3Input() *DataAwsCloudwatchLogDataProtectionPolicyDocument_S3Property
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
	PutCloudwatchLogs(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsProperty)
	// Experimental.
	PutFirehose(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehoseProperty)
	// Experimental.
	PutS3(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_S3Property)
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

// The jsii proxy struct for DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference
type jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CloudwatchLogs() DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CloudwatchLogsInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Firehose() DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehosePropertyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) FirehoseInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehoseProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) S3() DataAwsCloudwatchLogDataProtectionPolicyDocument_S3PropertyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocument_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) S3Input() *DataAwsCloudwatchLogDataProtectionPolicyDocument_S3Property {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataAwsCloudwatchLogDataProtectionPolicyDocument.FindingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference_Override(d DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataAwsCloudwatchLogDataProtectionPolicyDocument.FindingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutCloudwatchLogs(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_CloudwatchLogsProperty) {
	if err := d.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutFirehose(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_FirehoseProperty) {
	if err := d.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirehose",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutS3(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_S3Property) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		d,
		"resetFirehose",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

