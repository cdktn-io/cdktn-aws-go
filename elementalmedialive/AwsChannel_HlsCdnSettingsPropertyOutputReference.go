package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_HlsCdnSettingsPropertyOutputReference interface {
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
	HlsAkamaiSettings() AwsChannel_HlsAkamaiSettingsPropertyOutputReference
	// Experimental.
	HlsAkamaiSettingsInput() *AwsChannel_HlsAkamaiSettingsProperty
	// Experimental.
	HlsBasicPutSettings() AwsChannel_HlsBasicPutSettingsPropertyOutputReference
	// Experimental.
	HlsBasicPutSettingsInput() *AwsChannel_HlsBasicPutSettingsProperty
	// Experimental.
	HlsMediaStoreSettings() AwsChannel_HlsMediaStoreSettingsPropertyOutputReference
	// Experimental.
	HlsMediaStoreSettingsInput() *AwsChannel_HlsMediaStoreSettingsProperty
	// Experimental.
	HlsS3Settings() AwsChannel_HlsS3SettingsPropertyOutputReference
	// Experimental.
	HlsS3SettingsInput() *AwsChannel_HlsS3SettingsProperty
	// Experimental.
	HlsWebdavSettings() AwsChannel_HlsWebdavSettingsPropertyOutputReference
	// Experimental.
	HlsWebdavSettingsInput() *AwsChannel_HlsWebdavSettingsProperty
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
	PutHlsAkamaiSettings(value *AwsChannel_HlsAkamaiSettingsProperty)
	// Experimental.
	PutHlsBasicPutSettings(value *AwsChannel_HlsBasicPutSettingsProperty)
	// Experimental.
	PutHlsMediaStoreSettings(value *AwsChannel_HlsMediaStoreSettingsProperty)
	// Experimental.
	PutHlsS3Settings(value *AwsChannel_HlsS3SettingsProperty)
	// Experimental.
	PutHlsWebdavSettings(value *AwsChannel_HlsWebdavSettingsProperty)
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

// The jsii proxy struct for AwsChannel_HlsCdnSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsAkamaiSettings() AwsChannel_HlsAkamaiSettingsPropertyOutputReference {
	var returns AwsChannel_HlsAkamaiSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsAkamaiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsAkamaiSettingsInput() *AwsChannel_HlsAkamaiSettingsProperty {
	var returns *AwsChannel_HlsAkamaiSettingsProperty
	_jsii_.Get(
		j,
		"hlsAkamaiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsBasicPutSettings() AwsChannel_HlsBasicPutSettingsPropertyOutputReference {
	var returns AwsChannel_HlsBasicPutSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsBasicPutSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsBasicPutSettingsInput() *AwsChannel_HlsBasicPutSettingsProperty {
	var returns *AwsChannel_HlsBasicPutSettingsProperty
	_jsii_.Get(
		j,
		"hlsBasicPutSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsMediaStoreSettings() AwsChannel_HlsMediaStoreSettingsPropertyOutputReference {
	var returns AwsChannel_HlsMediaStoreSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsMediaStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsMediaStoreSettingsInput() *AwsChannel_HlsMediaStoreSettingsProperty {
	var returns *AwsChannel_HlsMediaStoreSettingsProperty
	_jsii_.Get(
		j,
		"hlsMediaStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsS3Settings() AwsChannel_HlsS3SettingsPropertyOutputReference {
	var returns AwsChannel_HlsS3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsS3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsS3SettingsInput() *AwsChannel_HlsS3SettingsProperty {
	var returns *AwsChannel_HlsS3SettingsProperty
	_jsii_.Get(
		j,
		"hlsS3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsWebdavSettings() AwsChannel_HlsWebdavSettingsPropertyOutputReference {
	var returns AwsChannel_HlsWebdavSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsWebdavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) HlsWebdavSettingsInput() *AwsChannel_HlsWebdavSettingsProperty {
	var returns *AwsChannel_HlsWebdavSettingsProperty
	_jsii_.Get(
		j,
		"hlsWebdavSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_HlsCdnSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsChannel_HlsCdnSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_HlsCdnSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsCdnSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_HlsCdnSettingsPropertyOutputReference_Override(a AwsChannel_HlsCdnSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsCdnSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) PutHlsAkamaiSettings(value *AwsChannel_HlsAkamaiSettingsProperty) {
	if err := a.validatePutHlsAkamaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsAkamaiSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) PutHlsBasicPutSettings(value *AwsChannel_HlsBasicPutSettingsProperty) {
	if err := a.validatePutHlsBasicPutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsBasicPutSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) PutHlsMediaStoreSettings(value *AwsChannel_HlsMediaStoreSettingsProperty) {
	if err := a.validatePutHlsMediaStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsMediaStoreSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) PutHlsS3Settings(value *AwsChannel_HlsS3SettingsProperty) {
	if err := a.validatePutHlsS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsS3Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) PutHlsWebdavSettings(value *AwsChannel_HlsWebdavSettingsProperty) {
	if err := a.validatePutHlsWebdavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsWebdavSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsAkamaiSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsAkamaiSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsBasicPutSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsBasicPutSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsMediaStoreSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsMediaStoreSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsS3Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsS3Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ResetHlsWebdavSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsWebdavSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsCdnSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

