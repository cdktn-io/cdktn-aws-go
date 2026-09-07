package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CertificateChain() *string
	// Experimental.
	SetCertificateChain(val *string)
	// Experimental.
	CertificateChainInput() *string
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
	InternalValue() *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference
type jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) InternalValue() *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty {
	var returns *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference_Override(a AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetCertificateChain(val *string) {
	if err := j.validateSetCertificateChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetInternalValue(val *AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

