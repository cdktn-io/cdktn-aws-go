package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_DataExportPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchMetrics() AwsStorageLensConfiguration_CloudWatchMetricsPropertyOutputReference
	// Experimental.
	CloudWatchMetricsInput() *AwsStorageLensConfiguration_CloudWatchMetricsProperty
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
	InternalValue() *AwsStorageLensConfiguration_DataExportProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_DataExportProperty)
	// Experimental.
	S3BucketDestination() AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationPropertyOutputReference
	// Experimental.
	S3BucketDestinationInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty
	// Experimental.
	StorageLensTableDestination() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationPropertyOutputReference
	// Experimental.
	StorageLensTableDestinationInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty
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
	PutCloudWatchMetrics(value *AwsStorageLensConfiguration_CloudWatchMetricsProperty)
	// Experimental.
	PutS3BucketDestination(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty)
	// Experimental.
	PutStorageLensTableDestination(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty)
	// Experimental.
	ResetCloudWatchMetrics()
	// Experimental.
	ResetS3BucketDestination()
	// Experimental.
	ResetStorageLensTableDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsStorageLensConfiguration_DataExportPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) CloudWatchMetrics() AwsStorageLensConfiguration_CloudWatchMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_CloudWatchMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) CloudWatchMetricsInput() *AwsStorageLensConfiguration_CloudWatchMetricsProperty {
	var returns *AwsStorageLensConfiguration_CloudWatchMetricsProperty
	_jsii_.Get(
		j,
		"cloudWatchMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_DataExportProperty {
	var returns *AwsStorageLensConfiguration_DataExportProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) S3BucketDestination() AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3BucketDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) S3BucketDestinationInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty
	_jsii_.Get(
		j,
		"s3BucketDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) StorageLensTableDestination() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"storageLensTableDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) StorageLensTableDestinationInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty
	_jsii_.Get(
		j,
		"storageLensTableDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_DataExportPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_DataExportPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_DataExportPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.DataExportPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_DataExportPropertyOutputReference_Override(a AwsStorageLensConfiguration_DataExportPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.DataExportPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_DataExportProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) PutCloudWatchMetrics(value *AwsStorageLensConfiguration_CloudWatchMetricsProperty) {
	if err := a.validatePutCloudWatchMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) PutS3BucketDestination(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty) {
	if err := a.validatePutS3BucketDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3BucketDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) PutStorageLensTableDestination(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty) {
	if err := a.validatePutStorageLensTableDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageLensTableDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ResetCloudWatchMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ResetS3BucketDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BucketDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ResetStorageLensTableDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageLensTableDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_DataExportPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

