package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/athena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/athena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkgroup_ResultConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AclConfiguration() AwsWorkgroup_AclConfigurationPropertyOutputReference
	// Experimental.
	AclConfigurationInput() *AwsWorkgroup_AclConfigurationProperty
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
	EncryptionConfiguration() AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationPropertyOutputReference
	// Experimental.
	EncryptionConfigurationInput() *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty
	// Experimental.
	ExpectedBucketOwner() *string
	// Experimental.
	SetExpectedBucketOwner(val *string)
	// Experimental.
	ExpectedBucketOwnerInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsWorkgroup_ResultConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsWorkgroup_ResultConfigurationProperty)
	// Experimental.
	OutputLocation() *string
	// Experimental.
	SetOutputLocation(val *string)
	// Experimental.
	OutputLocationInput() *string
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
	PutAclConfiguration(value *AwsWorkgroup_AclConfigurationProperty)
	// Experimental.
	PutEncryptionConfiguration(value *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty)
	// Experimental.
	ResetAclConfiguration()
	// Experimental.
	ResetEncryptionConfiguration()
	// Experimental.
	ResetExpectedBucketOwner()
	// Experimental.
	ResetOutputLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWorkgroup_ResultConfigurationPropertyOutputReference
type jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) AclConfiguration() AwsWorkgroup_AclConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_AclConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"aclConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) AclConfigurationInput() *AwsWorkgroup_AclConfigurationProperty {
	var returns *AwsWorkgroup_AclConfigurationProperty
	_jsii_.Get(
		j,
		"aclConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) EncryptionConfiguration() AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) EncryptionConfigurationInput() *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty {
	var returns *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) InternalValue() *AwsWorkgroup_ResultConfigurationProperty {
	var returns *AwsWorkgroup_ResultConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) OutputLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) OutputLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkgroup_ResultConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsWorkgroup_ResultConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkgroup_ResultConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.AwsWorkgroup.ResultConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkgroup_ResultConfigurationPropertyOutputReference_Override(a AwsWorkgroup_ResultConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.AwsWorkgroup.ResultConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetInternalValue(val *AwsWorkgroup_ResultConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetOutputLocation(val *string) {
	if err := j.validateSetOutputLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) PutAclConfiguration(value *AwsWorkgroup_AclConfigurationProperty) {
	if err := a.validatePutAclConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAclConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) PutEncryptionConfiguration(value *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty) {
	if err := a.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ResetAclConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAclConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ResetOutputLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkgroup_ResultConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

