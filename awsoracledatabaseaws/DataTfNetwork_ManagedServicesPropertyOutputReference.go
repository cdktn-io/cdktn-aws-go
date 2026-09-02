package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfNetwork_ManagedServicesPropertyOutputReference interface {
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
	CrossRegionS3RestoreSourcesAccess() DataTfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfNetwork_ManagedServicesProperty
	// Experimental.
	SetInternalValue(val *DataTfNetwork_ManagedServicesProperty)
	// Experimental.
	KmsAccess() DataTfNetwork_KmsAccessPropertyList
	// Experimental.
	ManagedS3BackupAccess() DataTfNetwork_ManagedS3BackupAccessPropertyList
	// Experimental.
	ManagedServiceIpv4Cidrs() *[]*string
	// Experimental.
	ResourceGatewayArn() *string
	// Experimental.
	S3Access() DataTfNetwork_S3AccessPropertyList
	// Experimental.
	ServiceNetworkArn() *string
	// Experimental.
	ServiceNetworkEndpoint() DataTfNetwork_ServiceNetworkEndpointPropertyList
	// Experimental.
	StsAccess() DataTfNetwork_StsAccessPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ZeroTlAccess() DataTfNetwork_ZeroTlAccessPropertyList
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

// The jsii proxy struct for DataTfNetwork_ManagedServicesPropertyOutputReference
type jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) CrossRegionS3RestoreSourcesAccess() DataTfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList {
	var returns DataTfNetwork_CrossRegionS3RestoreSourcesAccessPropertyList
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) InternalValue() *DataTfNetwork_ManagedServicesProperty {
	var returns *DataTfNetwork_ManagedServicesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) KmsAccess() DataTfNetwork_KmsAccessPropertyList {
	var returns DataTfNetwork_KmsAccessPropertyList
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ManagedS3BackupAccess() DataTfNetwork_ManagedS3BackupAccessPropertyList {
	var returns DataTfNetwork_ManagedS3BackupAccessPropertyList
	_jsii_.Get(
		j,
		"managedS3BackupAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ManagedServiceIpv4Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedServiceIpv4Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ResourceGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) S3Access() DataTfNetwork_S3AccessPropertyList {
	var returns DataTfNetwork_S3AccessPropertyList
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ServiceNetworkEndpoint() DataTfNetwork_ServiceNetworkEndpointPropertyList {
	var returns DataTfNetwork_ServiceNetworkEndpointPropertyList
	_jsii_.Get(
		j,
		"serviceNetworkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) StsAccess() DataTfNetwork_StsAccessPropertyList {
	var returns DataTfNetwork_StsAccessPropertyList
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ZeroTlAccess() DataTfNetwork_ZeroTlAccessPropertyList {
	var returns DataTfNetwork_ZeroTlAccessPropertyList
	_jsii_.Get(
		j,
		"zeroTlAccess",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfNetwork_ManagedServicesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfNetwork_ManagedServicesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfNetwork_ManagedServicesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataTfNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfNetwork_ManagedServicesPropertyOutputReference_Override(d DataTfNetwork_ManagedServicesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataTfNetwork.ManagedServicesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference)SetInternalValue(val *DataTfNetwork_ManagedServicesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfNetwork_ManagedServicesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

