package awscloudwatchlogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() DataTfDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *DataTfDataProtectionPolicyDocument_CloudwatchLogsProperty
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
	Firehose() DataTfDataProtectionPolicyDocument_FirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *DataTfDataProtectionPolicyDocument_FirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty
	// Experimental.
	SetInternalValue(val *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty)
	// Experimental.
	S3() DataTfDataProtectionPolicyDocument_S3PropertyOutputReference
	// Experimental.
	S3Input() *DataTfDataProtectionPolicyDocument_S3Property
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
	PutCloudwatchLogs(value *DataTfDataProtectionPolicyDocument_CloudwatchLogsProperty)
	// Experimental.
	PutFirehose(value *DataTfDataProtectionPolicyDocument_FirehoseProperty)
	// Experimental.
	PutS3(value *DataTfDataProtectionPolicyDocument_S3Property)
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

// The jsii proxy struct for DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference
type jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CloudwatchLogs() DataTfDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference {
	var returns DataTfDataProtectionPolicyDocument_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CloudwatchLogsInput() *DataTfDataProtectionPolicyDocument_CloudwatchLogsProperty {
	var returns *DataTfDataProtectionPolicyDocument_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Firehose() DataTfDataProtectionPolicyDocument_FirehosePropertyOutputReference {
	var returns DataTfDataProtectionPolicyDocument_FirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) FirehoseInput() *DataTfDataProtectionPolicyDocument_FirehoseProperty {
	var returns *DataTfDataProtectionPolicyDocument_FirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InternalValue() *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty {
	var returns *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) S3() DataTfDataProtectionPolicyDocument_S3PropertyOutputReference {
	var returns DataTfDataProtectionPolicyDocument_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) S3Input() *DataTfDataProtectionPolicyDocument_S3Property {
	var returns *DataTfDataProtectionPolicyDocument_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataTfDataProtectionPolicyDocument.FindingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference_Override(d DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataTfDataProtectionPolicyDocument.FindingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetInternalValue(val *DataTfDataProtectionPolicyDocument_FindingsDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutCloudwatchLogs(value *DataTfDataProtectionPolicyDocument_CloudwatchLogsProperty) {
	if err := d.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutFirehose(value *DataTfDataProtectionPolicyDocument_FirehoseProperty) {
	if err := d.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirehose",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) PutS3(value *DataTfDataProtectionPolicyDocument_S3Property) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		d,
		"resetFirehose",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfDataProtectionPolicyDocument_FindingsDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

