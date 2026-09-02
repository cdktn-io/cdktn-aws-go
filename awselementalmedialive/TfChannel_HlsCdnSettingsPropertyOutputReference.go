package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_HlsCdnSettingsPropertyOutputReference interface {
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
	HlsAkamaiSettings() TfChannel_HlsAkamaiSettingsPropertyOutputReference
	// Experimental.
	HlsAkamaiSettingsInput() *TfChannel_HlsAkamaiSettingsProperty
	// Experimental.
	HlsBasicPutSettings() TfChannel_HlsBasicPutSettingsPropertyOutputReference
	// Experimental.
	HlsBasicPutSettingsInput() *TfChannel_HlsBasicPutSettingsProperty
	// Experimental.
	HlsMediaStoreSettings() TfChannel_HlsMediaStoreSettingsPropertyOutputReference
	// Experimental.
	HlsMediaStoreSettingsInput() *TfChannel_HlsMediaStoreSettingsProperty
	// Experimental.
	HlsS3Settings() TfChannel_HlsS3SettingsPropertyOutputReference
	// Experimental.
	HlsS3SettingsInput() *TfChannel_HlsS3SettingsProperty
	// Experimental.
	HlsWebdavSettings() TfChannel_HlsWebdavSettingsPropertyOutputReference
	// Experimental.
	HlsWebdavSettingsInput() *TfChannel_HlsWebdavSettingsProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutHlsAkamaiSettings(value *TfChannel_HlsAkamaiSettingsProperty)
	// Experimental.
	PutHlsBasicPutSettings(value *TfChannel_HlsBasicPutSettingsProperty)
	// Experimental.
	PutHlsMediaStoreSettings(value *TfChannel_HlsMediaStoreSettingsProperty)
	// Experimental.
	PutHlsS3Settings(value *TfChannel_HlsS3SettingsProperty)
	// Experimental.
	PutHlsWebdavSettings(value *TfChannel_HlsWebdavSettingsProperty)
	// Experimental.
	ResetHlsAkamaiSettings()
	// Experimental.
	ResetHlsBasicPutSettings()
	// Experimental.
	ResetHlsMediaStoreSettings()
	// Experimental.
	ResetHlsS3Settings()
	// Experimental.
	ResetHlsWebdavSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_HlsCdnSettingsPropertyOutputReference
type jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsAkamaiSettings() TfChannel_HlsAkamaiSettingsPropertyOutputReference {
	var returns TfChannel_HlsAkamaiSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsAkamaiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsAkamaiSettingsInput() *TfChannel_HlsAkamaiSettingsProperty {
	var returns *TfChannel_HlsAkamaiSettingsProperty
	_jsii_.Get(
		j,
		"hlsAkamaiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsBasicPutSettings() TfChannel_HlsBasicPutSettingsPropertyOutputReference {
	var returns TfChannel_HlsBasicPutSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsBasicPutSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsBasicPutSettingsInput() *TfChannel_HlsBasicPutSettingsProperty {
	var returns *TfChannel_HlsBasicPutSettingsProperty
	_jsii_.Get(
		j,
		"hlsBasicPutSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsMediaStoreSettings() TfChannel_HlsMediaStoreSettingsPropertyOutputReference {
	var returns TfChannel_HlsMediaStoreSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsMediaStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsMediaStoreSettingsInput() *TfChannel_HlsMediaStoreSettingsProperty {
	var returns *TfChannel_HlsMediaStoreSettingsProperty
	_jsii_.Get(
		j,
		"hlsMediaStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsS3Settings() TfChannel_HlsS3SettingsPropertyOutputReference {
	var returns TfChannel_HlsS3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsS3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsS3SettingsInput() *TfChannel_HlsS3SettingsProperty {
	var returns *TfChannel_HlsS3SettingsProperty
	_jsii_.Get(
		j,
		"hlsS3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsWebdavSettings() TfChannel_HlsWebdavSettingsPropertyOutputReference {
	var returns TfChannel_HlsWebdavSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsWebdavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) HlsWebdavSettingsInput() *TfChannel_HlsWebdavSettingsProperty {
	var returns *TfChannel_HlsWebdavSettingsProperty
	_jsii_.Get(
		j,
		"hlsWebdavSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_HlsCdnSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_HlsCdnSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_HlsCdnSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsCdnSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_HlsCdnSettingsPropertyOutputReference_Override(t TfChannel_HlsCdnSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsCdnSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) PutHlsAkamaiSettings(value *TfChannel_HlsAkamaiSettingsProperty) {
	if err := t.validatePutHlsAkamaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsAkamaiSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) PutHlsBasicPutSettings(value *TfChannel_HlsBasicPutSettingsProperty) {
	if err := t.validatePutHlsBasicPutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsBasicPutSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) PutHlsMediaStoreSettings(value *TfChannel_HlsMediaStoreSettingsProperty) {
	if err := t.validatePutHlsMediaStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsMediaStoreSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) PutHlsS3Settings(value *TfChannel_HlsS3SettingsProperty) {
	if err := t.validatePutHlsS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsS3Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) PutHlsWebdavSettings(value *TfChannel_HlsWebdavSettingsProperty) {
	if err := t.validatePutHlsWebdavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsWebdavSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsAkamaiSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsAkamaiSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsBasicPutSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsBasicPutSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsMediaStoreSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsMediaStoreSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsS3Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsS3Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsWebdavSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsWebdavSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_HlsCdnSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

