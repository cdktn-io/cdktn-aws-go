package networkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference interface {
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
	EdgeLocation() *string
	// Experimental.
	SetEdgeLocation(val *string)
	// Experimental.
	EdgeLocationInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty
	// Experimental.
	SetInternalValue(val *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty)
	// Experimental.
	PeerEdgeLocation() *string
	// Experimental.
	SetPeerEdgeLocation(val *string)
	// Experimental.
	PeerEdgeLocationInput() *string
	// Experimental.
	RoutingPolicyNames() *[]*string
	// Experimental.
	SetRoutingPolicyNames(val *[]*string)
	// Experimental.
	RoutingPolicyNamesInput() *[]*string
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

// The jsii proxy struct for DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference
type jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) EdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) EdgeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) InternalValue() *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty {
	var returns *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) PeerEdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEdgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) PeerEdgeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEdgeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) RoutingPolicyNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) RoutingPolicyNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.EdgeLocationAssociationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference_Override(d DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.EdgeLocationAssociationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetEdgeLocation(val *string) {
	if err := j.validateSetEdgeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLocation",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetInternalValue(val *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetPeerEdgeLocation(val *string) {
	if err := j.validateSetPeerEdgeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerEdgeLocation",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetRoutingPolicyNames(val *[]*string) {
	if err := j.validateSetRoutingPolicyNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingPolicyNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

