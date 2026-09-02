package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference interface {
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
	InternalValue() *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty)
	// Experimental.
	S3BucketDestination() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationPropertyOutputReference
	// Experimental.
	S3BucketDestinationInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty
	// Experimental.
	StorageLensTableDestination() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationPropertyOutputReference
	// Experimental.
	StorageLensTableDestinationInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty
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
	PutS3BucketDestination(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty)
	// Experimental.
	PutStorageLensTableDestination(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty)
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

// The jsii proxy struct for TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty {
	var returns *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) S3BucketDestination() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3BucketDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) S3BucketDestinationInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty
	_jsii_.Get(
		j,
		"s3BucketDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) StorageLensTableDestination() TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"storageLensTableDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) StorageLensTableDestinationInput() *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty
	_jsii_.Get(
		j,
		"storageLensTableDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.ExpandedPrefixesDataExportPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference_Override(t TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.ExpandedPrefixesDataExportPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) PutS3BucketDestination(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty) {
	if err := t.validatePutS3BucketDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3BucketDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) PutStorageLensTableDestination(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty) {
	if err := t.validatePutStorageLensTableDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStorageLensTableDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ResetS3BucketDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BucketDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ResetStorageLensTableDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageLensTableDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_ExpandedPrefixesDataExportPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

