package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOdbNetwork_ManagedServicesPropertyOutputReference interface {
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
	CrossRegionS3RestoreSourcesAccess() AwsOdbNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOdbNetwork_ManagedServicesProperty
	// Experimental.
	SetInternalValue(val *AwsOdbNetwork_ManagedServicesProperty)
	// Experimental.
	KmsAccess() AwsOdbNetwork_KmsAccessPropertyList
	// Experimental.
	ManagedS3BackupAccess() AwsOdbNetwork_ManagedS3BackupAccessPropertyList
	// Experimental.
	ManagedServiceIpv4Cidrs() *[]*string
	// Experimental.
	ResourceGatewayArn() *string
	// Experimental.
	S3Access() AwsOdbNetwork_S3AccessPropertyList
	// Experimental.
	ServiceNetworkArn() *string
	// Experimental.
	ServiceNetworkEndpoint() AwsOdbNetwork_ServiceNetworkEndpointPropertyList
	// Experimental.
	StsAccess() AwsOdbNetwork_StsAccessPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ZeroEtlAccess() AwsOdbNetwork_ZeroEtlAccessPropertyList
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

// The jsii proxy struct for AwsOdbNetwork_ManagedServicesPropertyOutputReference
type jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) CrossRegionS3RestoreSourcesAccess() AwsOdbNetwork_CrossRegionS3RestoreSourcesAccessPropertyList {
	var returns AwsOdbNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) InternalValue() *AwsOdbNetwork_ManagedServicesProperty {
	var returns *AwsOdbNetwork_ManagedServicesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) KmsAccess() AwsOdbNetwork_KmsAccessPropertyList {
	var returns AwsOdbNetwork_KmsAccessPropertyList
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ManagedS3BackupAccess() AwsOdbNetwork_ManagedS3BackupAccessPropertyList {
	var returns AwsOdbNetwork_ManagedS3BackupAccessPropertyList
	_jsii_.Get(
		j,
		"managedS3BackupAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ManagedServiceIpv4Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedServiceIpv4Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ResourceGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) S3Access() AwsOdbNetwork_S3AccessPropertyList {
	var returns AwsOdbNetwork_S3AccessPropertyList
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkEndpoint() AwsOdbNetwork_ServiceNetworkEndpointPropertyList {
	var returns AwsOdbNetwork_ServiceNetworkEndpointPropertyList
	_jsii_.Get(
		j,
		"serviceNetworkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) StsAccess() AwsOdbNetwork_StsAccessPropertyList {
	var returns AwsOdbNetwork_StsAccessPropertyList
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ZeroEtlAccess() AwsOdbNetwork_ZeroEtlAccessPropertyList {
	var returns AwsOdbNetwork_ZeroEtlAccessPropertyList
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOdbNetwork_ManagedServicesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsOdbNetwork_ManagedServicesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOdbNetwork_ManagedServicesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOdbNetwork_ManagedServicesPropertyOutputReference_Override(a AwsOdbNetwork_ManagedServicesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference)SetInternalValue(val *AwsOdbNetwork_ManagedServicesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOdbNetwork_ManagedServicesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

