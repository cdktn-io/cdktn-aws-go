package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference interface {
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
	// Experimental.
	ContentTypeProfiles() TfFieldLevelEncryptionConfig_ContentTypeProfilesPropertyOutputReference
	// Experimental.
	ContentTypeProfilesInput() *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ForwardWhenContentTypeIsUnknown() interface{}
	// Experimental.
	SetForwardWhenContentTypeIsUnknown(val interface{})
	// Experimental.
	ForwardWhenContentTypeIsUnknownInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty
	// Experimental.
	SetInternalValue(val *TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty)
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
	PutContentTypeProfiles(value *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference
type jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ContentTypeProfiles() TfFieldLevelEncryptionConfig_ContentTypeProfilesPropertyOutputReference {
	var returns TfFieldLevelEncryptionConfig_ContentTypeProfilesPropertyOutputReference
	_jsii_.Get(
		j,
		"contentTypeProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ContentTypeProfilesInput() *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty {
	var returns *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty
	_jsii_.Get(
		j,
		"contentTypeProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ForwardWhenContentTypeIsUnknown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenContentTypeIsUnknown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ForwardWhenContentTypeIsUnknownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenContentTypeIsUnknownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) InternalValue() *TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty {
	var returns *TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfFieldLevelEncryptionConfig.ContentTypeProfileConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference_Override(t TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfFieldLevelEncryptionConfig.ContentTypeProfileConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetForwardWhenContentTypeIsUnknown(val interface{}) {
	if err := j.validateSetForwardWhenContentTypeIsUnknownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardWhenContentTypeIsUnknown",
		val,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetInternalValue(val *TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) PutContentTypeProfiles(value *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty) {
	if err := t.validatePutContentTypeProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContentTypeProfiles",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFieldLevelEncryptionConfig_ContentTypeProfileConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

