package transferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/transferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/transferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnector_As2ConfigPropertyOutputReference interface {
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
	Compression() *string
	// Experimental.
	SetCompression(val *string)
	// Experimental.
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EncryptionAlgorithm() *string
	// Experimental.
	SetEncryptionAlgorithm(val *string)
	// Experimental.
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConnector_As2ConfigProperty
	// Experimental.
	SetInternalValue(val *AwsConnector_As2ConfigProperty)
	// Experimental.
	LocalProfileId() *string
	// Experimental.
	SetLocalProfileId(val *string)
	// Experimental.
	LocalProfileIdInput() *string
	// Experimental.
	MdnResponse() *string
	// Experimental.
	SetMdnResponse(val *string)
	// Experimental.
	MdnResponseInput() *string
	// Experimental.
	MdnSigningAlgorithm() *string
	// Experimental.
	SetMdnSigningAlgorithm(val *string)
	// Experimental.
	MdnSigningAlgorithmInput() *string
	// Experimental.
	MessageSubject() *string
	// Experimental.
	SetMessageSubject(val *string)
	// Experimental.
	MessageSubjectInput() *string
	// Experimental.
	PartnerProfileId() *string
	// Experimental.
	SetPartnerProfileId(val *string)
	// Experimental.
	PartnerProfileIdInput() *string
	// Experimental.
	SigningAlgorithm() *string
	// Experimental.
	SetSigningAlgorithm(val *string)
	// Experimental.
	SigningAlgorithmInput() *string
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
	ResetMdnSigningAlgorithm()
	// Experimental.
	ResetMessageSubject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnector_As2ConfigPropertyOutputReference
type jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) InternalValue() *AwsConnector_As2ConfigProperty {
	var returns *AwsConnector_As2ConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) LocalProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) LocalProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MdnResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MdnResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MdnSigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnSigningAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MdnSigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnSigningAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MessageSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) MessageSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) PartnerProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) PartnerProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) SigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnector_As2ConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnector_As2ConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnector_As2ConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsConnector.As2ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnector_As2ConfigPropertyOutputReference_Override(a AwsConnector_As2ConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsConnector.As2ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetInternalValue(val *AwsConnector_As2ConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetLocalProfileId(val *string) {
	if err := j.validateSetLocalProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localProfileId",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetMdnResponse(val *string) {
	if err := j.validateSetMdnResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mdnResponse",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetMdnSigningAlgorithm(val *string) {
	if err := j.validateSetMdnSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mdnSigningAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetMessageSubject(val *string) {
	if err := j.validateSetMessageSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageSubject",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetPartnerProfileId(val *string) {
	if err := j.validateSetPartnerProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerProfileId",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetSigningAlgorithm(val *string) {
	if err := j.validateSetSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ResetMdnSigningAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetMdnSigningAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ResetMessageSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnector_As2ConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

