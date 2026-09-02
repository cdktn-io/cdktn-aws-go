package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobDefinition_PodPropertiesPropertyOutputReference interface {
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
	Containers() TfJobDefinition_ContainersPropertyList
	// Experimental.
	ContainersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DnsPolicy() *string
	// Experimental.
	SetDnsPolicy(val *string)
	// Experimental.
	DnsPolicyInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HostNetwork() interface{}
	// Experimental.
	SetHostNetwork(val interface{})
	// Experimental.
	HostNetworkInput() interface{}
	// Experimental.
	ImagePullSecret() TfJobDefinition_ImagePullSecretPropertyList
	// Experimental.
	ImagePullSecretInput() interface{}
	// Experimental.
	InitContainers() TfJobDefinition_InitContainersPropertyList
	// Experimental.
	InitContainersInput() interface{}
	// Experimental.
	InternalValue() *TfJobDefinition_PodPropertiesProperty
	// Experimental.
	SetInternalValue(val *TfJobDefinition_PodPropertiesProperty)
	// Experimental.
	Metadata() TfJobDefinition_MetadataPropertyOutputReference
	// Experimental.
	MetadataInput() *TfJobDefinition_MetadataProperty
	// Experimental.
	ServiceAccountName() *string
	// Experimental.
	SetServiceAccountName(val *string)
	// Experimental.
	ServiceAccountNameInput() *string
	// Experimental.
	ShareProcessNamespace() interface{}
	// Experimental.
	SetShareProcessNamespace(val interface{})
	// Experimental.
	ShareProcessNamespaceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Volumes() TfJobDefinition_VolumesPropertyList
	// Experimental.
	VolumesInput() interface{}
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
	PutContainers(value interface{})
	// Experimental.
	PutImagePullSecret(value interface{})
	// Experimental.
	PutInitContainers(value interface{})
	// Experimental.
	PutMetadata(value *TfJobDefinition_MetadataProperty)
	// Experimental.
	PutVolumes(value interface{})
	// Experimental.
	ResetDnsPolicy()
	// Experimental.
	ResetHostNetwork()
	// Experimental.
	ResetImagePullSecret()
	// Experimental.
	ResetInitContainers()
	// Experimental.
	ResetMetadata()
	// Experimental.
	ResetServiceAccountName()
	// Experimental.
	ResetShareProcessNamespace()
	// Experimental.
	ResetVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfJobDefinition_PodPropertiesPropertyOutputReference
type jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) Containers() TfJobDefinition_ContainersPropertyList {
	var returns TfJobDefinition_ContainersPropertyList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ImagePullSecret() TfJobDefinition_ImagePullSecretPropertyList {
	var returns TfJobDefinition_ImagePullSecretPropertyList
	_jsii_.Get(
		j,
		"imagePullSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ImagePullSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagePullSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) InitContainers() TfJobDefinition_InitContainersPropertyList {
	var returns TfJobDefinition_InitContainersPropertyList
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) InitContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initContainersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) InternalValue() *TfJobDefinition_PodPropertiesProperty {
	var returns *TfJobDefinition_PodPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) Metadata() TfJobDefinition_MetadataPropertyOutputReference {
	var returns TfJobDefinition_MetadataPropertyOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) MetadataInput() *TfJobDefinition_MetadataProperty {
	var returns *TfJobDefinition_MetadataProperty
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ShareProcessNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ShareProcessNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) Volumes() TfJobDefinition_VolumesPropertyList {
	var returns TfJobDefinition_VolumesPropertyList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobDefinition_PodPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobDefinition_PodPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobDefinition_PodPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.PodPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobDefinition_PodPropertiesPropertyOutputReference_Override(t TfJobDefinition_PodPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.PodPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetInternalValue(val *TfJobDefinition_PodPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetShareProcessNamespace(val interface{}) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) PutContainers(value interface{}) {
	if err := t.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainers",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) PutImagePullSecret(value interface{}) {
	if err := t.validatePutImagePullSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImagePullSecret",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) PutInitContainers(value interface{}) {
	if err := t.validatePutInitContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInitContainers",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) PutMetadata(value *TfJobDefinition_MetadataProperty) {
	if err := t.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadata",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) PutVolumes(value interface{}) {
	if err := t.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVolumes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		t,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetImagePullSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetImagePullSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetInitContainers() {
	_jsii_.InvokeVoid(
		t,
		"resetInitContainers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobDefinition_PodPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

