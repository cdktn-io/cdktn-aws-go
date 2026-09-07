package networkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionPropertyOutputReference
	// Experimental.
	ActionInput() *DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty
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
	Conditions() DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList
	// Experimental.
	ConditionsInput() interface{}
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
	EdgeLocations() *[]*string
	// Experimental.
	SetEdgeLocations(val *[]*string)
	// Experimental.
	EdgeLocationsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RuleNumber() *float64
	// Experimental.
	SetRuleNumber(val *float64)
	// Experimental.
	RuleNumberInput() *float64
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
	PutAction(value *DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty)
	// Experimental.
	PutConditions(value interface{})
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEdgeLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference
type jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) Action() DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionPropertyOutputReference {
	var returns DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ActionInput() *DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty {
	var returns *DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) Conditions() DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList {
	var returns DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) EdgeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) EdgeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"edgeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) RuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) RuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.AttachmentRoutingPolicyRulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference_Override(d DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataAwsCoreNetworkPolicyDocument.AttachmentRoutingPolicyRulesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetEdgeLocations(val *[]*string) {
	if err := j.validateSetEdgeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLocations",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetRuleNumber(val *float64) {
	if err := j.validateSetRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNumber",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) PutAction(value *DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty) {
	if err := d.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) PutConditions(value interface{}) {
	if err := d.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ResetEdgeLocations() {
	_jsii_.InvokeVoid(
		d,
		"resetEdgeLocations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

