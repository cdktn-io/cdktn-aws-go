package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty)
	// Experimental.
	KeyId() *string
	// Experimental.
	SetKeyId(val *string)
	// Experimental.
	KeyIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference_Override(a AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetKeyId(val *string) {
	if err := j.validateSetKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

