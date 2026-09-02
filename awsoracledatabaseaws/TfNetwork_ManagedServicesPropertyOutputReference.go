package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNetwork_ManagedServicesPropertyOutputReference interface {
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
	CrossRegionS3RestoreSourcesAccess() TfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfNetwork_ManagedServicesProperty
	// Experimental.
	SetInternalValue(val *TfNetwork_ManagedServicesProperty)
	// Experimental.
	KmsAccess() TfNetwork_KmsAccessPropertyList
	// Experimental.
	ManagedS3BackupAccess() TfNetwork_ManagedS3BackupAccessPropertyList
	// Experimental.
	ManagedServiceIpv4Cidrs() *[]*string
	// Experimental.
	ResourceGatewayArn() *string
	// Experimental.
	S3Access() TfNetwork_S3AccessPropertyList
	// Experimental.
	ServiceNetworkArn() *string
	// Experimental.
	ServiceNetworkEndpoint() TfNetwork_ServiceNetworkEndpointPropertyList
	// Experimental.
	StsAccess() TfNetwork_StsAccessPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ZeroEtlAccess() TfNetwork_ZeroEtlAccessPropertyList
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

// The jsii proxy struct for TfNetwork_ManagedServicesPropertyOutputReference
type jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) CrossRegionS3RestoreSourcesAccess() TfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList {
	var returns TfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) InternalValue() *TfNetwork_ManagedServicesProperty {
	var returns *TfNetwork_ManagedServicesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) KmsAccess() TfNetwork_KmsAccessPropertyList {
	var returns TfNetwork_KmsAccessPropertyList
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ManagedS3BackupAccess() TfNetwork_ManagedS3BackupAccessPropertyList {
	var returns TfNetwork_ManagedS3BackupAccessPropertyList
	_jsii_.Get(
		j,
		"managedS3BackupAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ManagedServiceIpv4Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedServiceIpv4Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ResourceGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) S3Access() TfNetwork_S3AccessPropertyList {
	var returns TfNetwork_S3AccessPropertyList
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkEndpoint() TfNetwork_ServiceNetworkEndpointPropertyList {
	var returns TfNetwork_ServiceNetworkEndpointPropertyList
	_jsii_.Get(
		j,
		"serviceNetworkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) StsAccess() TfNetwork_StsAccessPropertyList {
	var returns TfNetwork_StsAccessPropertyList
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ZeroEtlAccess() TfNetwork_ZeroEtlAccessPropertyList {
	var returns TfNetwork_ZeroEtlAccessPropertyList
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfNetwork_ManagedServicesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfNetwork_ManagedServicesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfNetwork_ManagedServicesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfNetwork_ManagedServicesPropertyOutputReference_Override(t TfNetwork_ManagedServicesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference)SetInternalValue(val *TfNetwork_ManagedServicesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfNetwork_ManagedServicesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

