package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchEncryption() AwsSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference
	// Experimental.
	CloudwatchEncryptionInput() *AwsSecurityConfiguration_CloudwatchEncryptionProperty
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
	InternalValue() *AwsSecurityConfiguration_EncryptionConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsSecurityConfiguration_EncryptionConfigurationProperty)
	// Experimental.
	JobBookmarksEncryption() AwsSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference
	// Experimental.
	JobBookmarksEncryptionInput() *AwsSecurityConfiguration_JobBookmarksEncryptionProperty
	// Experimental.
	S3Encryption() AwsSecurityConfiguration_S3EncryptionPropertyOutputReference
	// Experimental.
	S3EncryptionInput() *AwsSecurityConfiguration_S3EncryptionProperty
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
	PutCloudwatchEncryption(value *AwsSecurityConfiguration_CloudwatchEncryptionProperty)
	// Experimental.
	PutJobBookmarksEncryption(value *AwsSecurityConfiguration_JobBookmarksEncryptionProperty)
	// Experimental.
	PutS3Encryption(value *AwsSecurityConfiguration_S3EncryptionProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference
type jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CloudwatchEncryption() AwsSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference {
	var returns AwsSecurityConfiguration_CloudwatchEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CloudwatchEncryptionInput() *AwsSecurityConfiguration_CloudwatchEncryptionProperty {
	var returns *AwsSecurityConfiguration_CloudwatchEncryptionProperty
	_jsii_.Get(
		j,
		"cloudwatchEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InternalValue() *AwsSecurityConfiguration_EncryptionConfigurationProperty {
	var returns *AwsSecurityConfiguration_EncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) JobBookmarksEncryption() AwsSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference {
	var returns AwsSecurityConfiguration_JobBookmarksEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"jobBookmarksEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) JobBookmarksEncryptionInput() *AwsSecurityConfiguration_JobBookmarksEncryptionProperty {
	var returns *AwsSecurityConfiguration_JobBookmarksEncryptionProperty
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) S3Encryption() AwsSecurityConfiguration_S3EncryptionPropertyOutputReference {
	var returns AwsSecurityConfiguration_S3EncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) S3EncryptionInput() *AwsSecurityConfiguration_S3EncryptionProperty {
	var returns *AwsSecurityConfiguration_S3EncryptionProperty
	_jsii_.Get(
		j,
		"s3EncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsSecurityConfiguration.EncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference_Override(a AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsSecurityConfiguration.EncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetInternalValue(val *AwsSecurityConfiguration_EncryptionConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutCloudwatchEncryption(value *AwsSecurityConfiguration_CloudwatchEncryptionProperty) {
	if err := a.validatePutCloudwatchEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutJobBookmarksEncryption(value *AwsSecurityConfiguration_JobBookmarksEncryptionProperty) {
	if err := a.validatePutJobBookmarksEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobBookmarksEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) PutS3Encryption(value *AwsSecurityConfiguration_S3EncryptionProperty) {
	if err := a.validatePutS3EncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Encryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSecurityConfiguration_EncryptionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

