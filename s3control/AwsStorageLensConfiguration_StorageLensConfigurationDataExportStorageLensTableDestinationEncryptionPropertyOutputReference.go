package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference interface {
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
	InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty)
	// Experimental.
	SseKms() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsPropertyOutputReference
	// Experimental.
	SseKmsInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty
	// Experimental.
	SseS3() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseS3PropertyList
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
	PutSseKms(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty)
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

// The jsii proxy struct for AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) SseKms() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsPropertyOutputReference
	_jsii_.Get(
		j,
		"sseKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) SseKmsInput() *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty
	_jsii_.Get(
		j,
		"sseKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) SseS3() AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseS3PropertyList {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseS3PropertyList
	_jsii_.Get(
		j,
		"sseS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) SseS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference_Override(a AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) PutSseKms(value *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty) {
	if err := a.validatePutSseKmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseKms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) PutSseS3(value interface{}) {
	if err := a.validatePutSseS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ResetSseKms() {
	_jsii_.InvokeVoid(
		a,
		"resetSseKms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ResetSseS3() {
	_jsii_.InvokeVoid(
		a,
		"resetSseS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

