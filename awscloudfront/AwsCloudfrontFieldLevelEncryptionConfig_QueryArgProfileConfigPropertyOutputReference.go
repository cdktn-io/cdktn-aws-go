package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference interface {
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
	ForwardWhenQueryArgProfileIsUnknown() interface{}
	// Experimental.
	SetForwardWhenQueryArgProfileIsUnknown(val interface{})
	// Experimental.
	ForwardWhenQueryArgProfileIsUnknownInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty
	// Experimental.
	SetInternalValue(val *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty)
	// Experimental.
	QueryArgProfiles() AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesPropertyOutputReference
	// Experimental.
	QueryArgProfilesInput() *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty
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
	PutQueryArgProfiles(value *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty)
	// Experimental.
	ResetQueryArgProfiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference
type jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ForwardWhenQueryArgProfileIsUnknown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenQueryArgProfileIsUnknown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ForwardWhenQueryArgProfileIsUnknownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenQueryArgProfileIsUnknownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) InternalValue() *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty {
	var returns *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) QueryArgProfiles() AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesPropertyOutputReference {
	var returns AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesPropertyOutputReference
	_jsii_.Get(
		j,
		"queryArgProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) QueryArgProfilesInput() *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty {
	var returns *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty
	_jsii_.Get(
		j,
		"queryArgProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontFieldLevelEncryptionConfig.QueryArgProfileConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference_Override(a AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontFieldLevelEncryptionConfig.QueryArgProfileConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetForwardWhenQueryArgProfileIsUnknown(val interface{}) {
	if err := j.validateSetForwardWhenQueryArgProfileIsUnknownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardWhenQueryArgProfileIsUnknown",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetInternalValue(val *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) PutQueryArgProfiles(value *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty) {
	if err := a.validatePutQueryArgProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryArgProfiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ResetQueryArgProfiles() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryArgProfiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

