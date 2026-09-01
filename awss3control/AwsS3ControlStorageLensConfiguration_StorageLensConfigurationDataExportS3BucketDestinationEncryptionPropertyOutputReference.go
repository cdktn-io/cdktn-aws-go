package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference interface {
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
	InternalValue() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionProperty)
	// Experimental.
	SseKms() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference
	// Experimental.
	SseKmsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty
	// Experimental.
	SseS3() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3PropertyList
	// Experimental.
	SseS3Input() interface{}
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
	PutSseKms(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty)
	// Experimental.
	PutSseS3(value interface{})
	// Experimental.
	ResetSseKms()
	// Experimental.
	ResetSseS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference
type jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) InternalValue() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseKms() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference
	_jsii_.Get(
		j,
		"sseKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseKmsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty
	_jsii_.Get(
		j,
		"sseKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseS3() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3PropertyList {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3PropertyList
	_jsii_.Get(
		j,
		"sseS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference_Override(a AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetInternalValue(val *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) PutSseKms(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionSseKmsProperty) {
	if err := a.validatePutSseKmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseKms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) PutSseS3(value interface{}) {
	if err := a.validatePutSseS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ResetSseKms() {
	_jsii_.InvokeVoid(
		a,
		"resetSseKms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ResetSseS3() {
	_jsii_.InvokeVoid(
		a,
		"resetSseS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationEncryptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

