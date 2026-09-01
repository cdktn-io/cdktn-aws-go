package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acm() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmPropertyOutputReference
	// Experimental.
	AcmInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty
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
	File() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFilePropertyOutputReference
	// Experimental.
	FileInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty)
	// Experimental.
	Sds() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsPropertyOutputReference
	// Experimental.
	SdsInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty
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
	PutAcm(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty)
	// Experimental.
	PutFile(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty)
	// Experimental.
	PutSds(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty)
	// Experimental.
	ResetAcm()
	// Experimental.
	ResetFile()
	// Experimental.
	ResetSds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference
type jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) Acm() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmPropertyOutputReference {
	var returns AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmPropertyOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) AcmInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) File() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFilePropertyOutputReference {
	var returns AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) FileInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) InternalValue() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) Sds() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsPropertyOutputReference {
	var returns AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) SdsInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference_Override(a AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference)SetInternalValue(val *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) PutAcm(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty) {
	if err := a.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) PutFile(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustFileProperty) {
	if err := a.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) PutSds(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustSdsProperty) {
	if err := a.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		a,
		"resetAcm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		a,
		"resetFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		a,
		"resetSds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

