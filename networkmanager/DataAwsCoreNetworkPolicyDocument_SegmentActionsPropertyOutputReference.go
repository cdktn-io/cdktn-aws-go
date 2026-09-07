package networkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() *string
	// Experimental.
	SetAction(val *string)
	// Experimental.
	ActionInput() *string
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DestinationCidrBlocks() *[]*string
	// Experimental.
	SetDestinationCidrBlocks(val *[]*string)
	// Experimental.
	DestinationCidrBlocksInput() *[]*string
	// Experimental.
	Destinations() *[]*string
	// Experimental.
	SetDestinations(val *[]*string)
	// Experimental.
	DestinationsInput() *[]*string
	// Experimental.
	EdgeLocationAssociation() DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference
	// Experimental.
	EdgeLocationAssociationInput() *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	RoutingPolicyNames() *[]*string
	// Experimental.
	SetRoutingPolicyNames(val *[]*string)
	// Experimental.
	RoutingPolicyNamesInput() *[]*string
	// Experimental.
	Segment() *string
	// Experimental.
	SetSegment(val *string)
	// Experimental.
	SegmentInput() *string
	// Experimental.
	ShareWith() *[]*string
	// Experimental.
	SetShareWith(val *[]*string)
	// Experimental.
	ShareWithExcept() *[]*string
	// Experimental.
	SetShareWithExcept(val *[]*string)
	// Experimental.
	ShareWithExceptInput() *[]*string
	// Experimental.
	ShareWithInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Via() DataAwsCoreNetworkPolicyDocument_ViaPropertyOutputReference
	// Experimental.
	ViaInput() *DataAwsCoreNetworkPolicyDocument_ViaProperty
	// Experimental.
	WhenSentTo() DataAwsCoreNetworkPolicyDocument_WhenSentToPropertyOutputReference
	// Experimental.
	WhenSentToInput() *DataAwsCoreNetworkPolicyDocument_WhenSentToProperty
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
	PutEdgeLocationAssociation(value *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty)
	// Experimental.
	PutVia(value *DataAwsCoreNetworkPolicyDocument_ViaProperty)
	// Experimental.
	PutWhenSentTo(value *DataAwsCoreNetworkPolicyDocument_WhenSentToProperty)
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDestinationCidrBlocks()
	// Experimental.
	ResetDestinations()
	// Experimental.
	ResetEdgeLocationAssociation()
	// Experimental.
	ResetMode()
	// Experimental.
	ResetRoutingPolicyNames()
	// Experimental.
	ResetShareWith()
	// Experimental.
	ResetShareWithExcept()
	// Experimental.
	ResetVia()
	// Experimental.
	ResetWhenSentTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference
type jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) DestinationCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) DestinationCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Destinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) DestinationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) EdgeLocationAssociation() DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference {
	var returns DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationPropertyOutputReference
	_jsii_.Get(
		j,
		"edgeLocationAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) EdgeLocationAssociationInput() *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty {
	var returns *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty
	_jsii_.Get(
		j,
		"edgeLocationAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) RoutingPolicyNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) RoutingPolicyNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routingPolicyNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Segment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) SegmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ShareWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shareWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ShareWithExcept() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shareWithExcept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ShareWithExceptInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shareWithExceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ShareWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shareWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Via() DataAwsCoreNetworkPolicyDocument_ViaPropertyOutputReference {
	var returns DataAwsCoreNetworkPolicyDocument_ViaPropertyOutputReference
	_jsii_.Get(
		j,
		"via",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ViaInput() *DataAwsCoreNetworkPolicyDocument_ViaProperty {
	var returns *DataAwsCoreNetworkPolicyDocument_ViaProperty
	_jsii_.Get(
		j,
		"viaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) WhenSentTo() DataAwsCoreNetworkPolicyDocument_WhenSentToPropertyOutputReference {
	var returns DataAwsCoreNetworkPolicyDocument_WhenSentToPropertyOutputReference
	_jsii_.Get(
		j,
		"whenSentTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) WhenSentToInput() *DataAwsCoreNetworkPolicyDocument_WhenSentToProperty {
	var returns *DataAwsCoreNetworkPolicyDocument_WhenSentToProperty
	_jsii_.Get(
		j,
		"whenSentToInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.SegmentActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference_Override(d DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.SegmentActionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetDestinationCidrBlocks(val *[]*string) {
	if err := j.validateSetDestinationCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetDestinations(val *[]*string) {
	if err := j.validateSetDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetRoutingPolicyNames(val *[]*string) {
	if err := j.validateSetRoutingPolicyNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingPolicyNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetSegment(val *string) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetShareWith(val *[]*string) {
	if err := j.validateSetShareWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareWith",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetShareWithExcept(val *[]*string) {
	if err := j.validateSetShareWithExceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareWithExcept",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) PutEdgeLocationAssociation(value *DataAwsCoreNetworkPolicyDocument_EdgeLocationAssociationProperty) {
	if err := d.validatePutEdgeLocationAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEdgeLocationAssociation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) PutVia(value *DataAwsCoreNetworkPolicyDocument_ViaProperty) {
	if err := d.validatePutViaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVia",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) PutWhenSentTo(value *DataAwsCoreNetworkPolicyDocument_WhenSentToProperty) {
	if err := d.validatePutWhenSentToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWhenSentTo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetDestinationCidrBlocks() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationCidrBlocks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetEdgeLocationAssociation() {
	_jsii_.InvokeVoid(
		d,
		"resetEdgeLocationAssociation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetRoutingPolicyNames() {
	_jsii_.InvokeVoid(
		d,
		"resetRoutingPolicyNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetShareWith() {
	_jsii_.InvokeVoid(
		d,
		"resetShareWith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetShareWithExcept() {
	_jsii_.InvokeVoid(
		d,
		"resetShareWithExcept",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetVia() {
	_jsii_.InvokeVoid(
		d,
		"resetVia",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ResetWhenSentTo() {
	_jsii_.InvokeVoid(
		d,
		"resetWhenSentTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_SegmentActionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

