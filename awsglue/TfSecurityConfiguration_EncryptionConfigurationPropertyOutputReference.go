package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchEncryption() TfSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference
	// Experimental.
	CloudwatchEncryptionInput() *TfSecurityConfiguration_CloudwatchEncryptionProperty
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
	InternalValue() *TfSecurityConfiguration_EncryptionConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfSecurityConfiguration_EncryptionConfigurationProperty)
	// Experimental.
	JobBookmarksEncryption() TfSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference
	// Experimental.
	JobBookmarksEncryptionInput() *TfSecurityConfiguration_JobBookmarksEncryptionProperty
	// Experimental.
	S3Encryption() TfSecurityConfiguration_S3EncryptionPropertyOutputReference
	// Experimental.
	S3EncryptionInput() *TfSecurityConfiguration_S3EncryptionProperty
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
	PutCloudwatchEncryption(value *TfSecurityConfiguration_CloudwatchEncryptionProperty)
	// Experimental.
	PutJobBookmarksEncryption(value *TfSecurityConfiguration_JobBookmarksEncryptionProperty)
	// Experimental.
	PutS3Encryption(value *TfSecurityConfiguration_S3EncryptionProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference
type jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CloudwatchEncryption() TfSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference {
	var returns TfSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CloudwatchEncryptionInput() *TfSecurityConfiguration_CloudwatchEncryptionProperty {
	var returns *TfSecurityConfiguration_CloudwatchEncryptionProperty
	_jsii_.Get(
		j,
		"cloudwatchEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InternalValue() *TfSecurityConfiguration_EncryptionConfigurationProperty {
	var returns *TfSecurityConfiguration_EncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) JobBookmarksEncryption() TfSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference {
	var returns TfSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"jobBookmarksEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) JobBookmarksEncryptionInput() *TfSecurityConfiguration_JobBookmarksEncryptionProperty {
	var returns *TfSecurityConfiguration_JobBookmarksEncryptionProperty
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) S3Encryption() TfSecurityConfiguration_S3EncryptionPropertyOutputReference {
	var returns TfSecurityConfiguration_S3EncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) S3EncryptionInput() *TfSecurityConfiguration_S3EncryptionProperty {
	var returns *TfSecurityConfiguration_S3EncryptionProperty
	_jsii_.Get(
		j,
		"s3EncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSecurityConfiguration_EncryptionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfSecurityConfiguration.EncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference_Override(t TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfSecurityConfiguration.EncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetInternalValue(val *TfSecurityConfiguration_EncryptionConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutCloudwatchEncryption(value *TfSecurityConfiguration_CloudwatchEncryptionProperty) {
	if err := t.validatePutCloudwatchEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchEncryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutJobBookmarksEncryption(value *TfSecurityConfiguration_JobBookmarksEncryptionProperty) {
	if err := t.validatePutJobBookmarksEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJobBookmarksEncryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutS3Encryption(value *TfSecurityConfiguration_S3EncryptionProperty) {
	if err := t.validatePutS3EncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Encryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

