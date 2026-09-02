package awssesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEmailIdentity_DkimSigningAttributesPropertyOutputReference interface {
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
	InternalValue() *TfEmailIdentity_DkimSigningAttributesProperty
	// Experimental.
	SetInternalValue(val *TfEmailIdentity_DkimSigningAttributesProperty)
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

// The jsii proxy struct for TfEmailIdentity_DkimSigningAttributesPropertyOutputReference
type jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) CurrentSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) DomainSigningSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) InternalValue() *TfEmailIdentity_DkimSigningAttributesProperty {
	var returns *TfEmailIdentity_DkimSigningAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) LastKeyGenerationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastKeyGenerationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) NextSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) NextSigningKeyLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) SigningAttributesOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAttributesOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) Tokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEmailIdentity_DkimSigningAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEmailIdentity_DkimSigningAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEmailIdentity_DkimSigningAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfEmailIdentity.DkimSigningAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEmailIdentity_DkimSigningAttributesPropertyOutputReference_Override(t TfEmailIdentity_DkimSigningAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfEmailIdentity.DkimSigningAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetDomainSigningPrivateKey(val *string) {
	if err := j.validateSetDomainSigningPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningPrivateKey",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetDomainSigningSelector(val *string) {
	if err := j.validateSetDomainSigningSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningSelector",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetInternalValue(val *TfEmailIdentity_DkimSigningAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetNextSigningKeyLength(val *string) {
	if err := j.validateSetNextSigningKeyLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextSigningKeyLength",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetDomainSigningPrivateKey() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainSigningPrivateKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetDomainSigningSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainSigningSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ResetNextSigningKeyLength() {
	_jsii_.InvokeVoid(
		t,
		"resetNextSigningKeyLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEmailIdentity_DkimSigningAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

