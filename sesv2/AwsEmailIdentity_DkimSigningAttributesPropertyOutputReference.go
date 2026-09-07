package sesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference interface {
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
	CurrentSigningKeyLength() *string
	// Experimental.
	DomainSigningPrivateKey() *string
	// Experimental.
	SetDomainSigningPrivateKey(val *string)
	// Experimental.
	DomainSigningPrivateKeyInput() *string
	// Experimental.
	DomainSigningSelector() *string
	// Experimental.
	SetDomainSigningSelector(val *string)
	// Experimental.
	DomainSigningSelectorInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEmailIdentity_DkimSigningAttributesProperty
	// Experimental.
	SetInternalValue(val *AwsEmailIdentity_DkimSigningAttributesProperty)
	// Experimental.
	LastKeyGenerationTimestamp() *string
	// Experimental.
	NextSigningKeyLength() *string
	// Experimental.
	SetNextSigningKeyLength(val *string)
	// Experimental.
	NextSigningKeyLengthInput() *string
	// Experimental.
	SigningAttributesOrigin() *string
	// Experimental.
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Tokens() *[]*string
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
	ResetDomainSigningPrivateKey()
	// Experimental.
	ResetDomainSigningSelector()
	// Experimental.
	ResetNextSigningKeyLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference
type jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) CurrentSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) InternalValue() *AwsEmailIdentity_DkimSigningAttributesProperty {
	var returns *AwsEmailIdentity_DkimSigningAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) LastKeyGenerationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastKeyGenerationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) NextSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) NextSigningKeyLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) SigningAttributesOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAttributesOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) Tokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmailIdentity_DkimSigningAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmailIdentity_DkimSigningAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsEmailIdentity.DkimSigningAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmailIdentity_DkimSigningAttributesPropertyOutputReference_Override(a AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsEmailIdentity.DkimSigningAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetDomainSigningPrivateKey(val *string) {
	if err := j.validateSetDomainSigningPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningPrivateKey",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetDomainSigningSelector(val *string) {
	if err := j.validateSetDomainSigningSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningSelector",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetInternalValue(val *AwsEmailIdentity_DkimSigningAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetNextSigningKeyLength(val *string) {
	if err := j.validateSetNextSigningKeyLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextSigningKeyLength",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetDomainSigningPrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainSigningPrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetDomainSigningSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainSigningSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetNextSigningKeyLength() {
	_jsii_.InvokeVoid(
		a,
		"resetNextSigningKeyLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmailIdentity_DkimSigningAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

