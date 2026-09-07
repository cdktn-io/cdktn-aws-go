package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/batch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/batch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobDefinition_PodPropertiesPropertyOutputReference interface {
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
	Containers() AwsJobDefinition_ContainersPropertyList
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
	ImagePullSecret() AwsJobDefinition_ImagePullSecretPropertyList
	// Experimental.
	ImagePullSecretInput() interface{}
	// Experimental.
	InitContainers() AwsJobDefinition_InitContainersPropertyList
	// Experimental.
	InitContainersInput() interface{}
	// Experimental.
	InternalValue() *AwsJobDefinition_PodPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsJobDefinition_PodPropertiesProperty)
	// Experimental.
	Metadata() AwsJobDefinition_MetadataPropertyOutputReference
	// Experimental.
	MetadataInput() *AwsJobDefinition_MetadataProperty
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
	Volumes() AwsJobDefinition_VolumesPropertyList
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
	PutMetadata(value *AwsJobDefinition_MetadataProperty)
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

// The jsii proxy struct for AwsJobDefinition_PodPropertiesPropertyOutputReference
type jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) Containers() AwsJobDefinition_ContainersPropertyList {
	var returns AwsJobDefinition_ContainersPropertyList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ImagePullSecret() AwsJobDefinition_ImagePullSecretPropertyList {
	var returns AwsJobDefinition_ImagePullSecretPropertyList
	_jsii_.Get(
		j,
		"imagePullSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ImagePullSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagePullSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) InitContainers() AwsJobDefinition_InitContainersPropertyList {
	var returns AwsJobDefinition_InitContainersPropertyList
	_jsii_.Get(
		j,
		"initContainers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) InitContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initContainersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) InternalValue() *AwsJobDefinition_PodPropertiesProperty {
	var returns *AwsJobDefinition_PodPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) Metadata() AwsJobDefinition_MetadataPropertyOutputReference {
	var returns AwsJobDefinition_MetadataPropertyOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) MetadataInput() *AwsJobDefinition_MetadataProperty {
	var returns *AwsJobDefinition_MetadataProperty
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ShareProcessNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ShareProcessNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) Volumes() AwsJobDefinition_VolumesPropertyList {
	var returns AwsJobDefinition_VolumesPropertyList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobDefinition_PodPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobDefinition_PodPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobDefinition_PodPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.PodPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobDefinition_PodPropertiesPropertyOutputReference_Override(a AwsJobDefinition_PodPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.PodPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetInternalValue(val *AwsJobDefinition_PodPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetShareProcessNamespace(val interface{}) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) PutContainers(value interface{}) {
	if err := a.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) PutImagePullSecret(value interface{}) {
	if err := a.validatePutImagePullSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImagePullSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) PutInitContainers(value interface{}) {
	if err := a.validatePutInitContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitContainers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) PutMetadata(value *AwsJobDefinition_MetadataProperty) {
	if err := a.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) PutVolumes(value interface{}) {
	if err := a.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVolumes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		a,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetImagePullSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetImagePullSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetInitContainers() {
	_jsii_.InvokeVoid(
		a,
		"resetInitContainers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_PodPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

