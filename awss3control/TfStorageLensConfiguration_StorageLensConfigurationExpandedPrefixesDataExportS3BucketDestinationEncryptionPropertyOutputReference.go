package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference interface {
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
	InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty)
	// Experimental.
	SseKms() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference
	// Experimental.
	SseKmsInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsProperty
	// Experimental.
	SseS3() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseS3PropertyList
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
	PutSseKms(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsProperty)
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

// The jsii proxy struct for TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseKms() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsPropertyOutputReference
	_jsii_.Get(
		j,
		"sseKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseKmsInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsProperty
	_jsii_.Get(
		j,
		"sseKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseS3() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseS3PropertyList {
	var returns TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseS3PropertyList
	_jsii_.Get(
		j,
		"sseS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) SseS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference_Override(t TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) PutSseKms(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionSseKmsProperty) {
	if err := t.validatePutSseKmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSseKms",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) PutSseS3(value interface{}) {
	if err := t.validatePutSseS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSseS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ResetSseKms() {
	_jsii_.InvokeVoid(
		t,
		"resetSseKms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ResetSseS3() {
	_jsii_.InvokeVoid(
		t,
		"resetSseS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

