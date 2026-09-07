package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acm() AwsVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference
	// Experimental.
	AcmInput() *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty
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
	File() AwsVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference
	// Experimental.
	FileInput() *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsVirtualGateway_SpecListenerTlsCertificateProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualGateway_SpecListenerTlsCertificateProperty)
	// Experimental.
	Sds() AwsVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference
	// Experimental.
	SdsInput() *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty
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
	PutAcm(value *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty)
	// Experimental.
	PutFile(value *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty)
	// Experimental.
	PutSds(value *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty)
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

// The jsii proxy struct for AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference
type jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Acm() AwsVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference {
	var returns AwsVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) AcmInput() *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty {
	var returns *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) File() AwsVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference {
	var returns AwsVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) FileInput() *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty {
	var returns *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InternalValue() *AwsVirtualGateway_SpecListenerTlsCertificateProperty {
	var returns *AwsVirtualGateway_SpecListenerTlsCertificateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Sds() AwsVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference {
	var returns AwsVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) SdsInput() *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty {
	var returns *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecListenerTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference_Override(a AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecListenerTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetInternalValue(val *AwsVirtualGateway_SpecListenerTlsCertificateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutAcm(value *AwsVirtualGateway_SpecListenerTlsCertificateAcmProperty) {
	if err := a.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutFile(value *AwsVirtualGateway_SpecListenerTlsCertificateFileProperty) {
	if err := a.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutSds(value *AwsVirtualGateway_SpecListenerTlsCertificateSdsProperty) {
	if err := a.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		a,
		"resetAcm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		a,
		"resetFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		a,
		"resetSds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

