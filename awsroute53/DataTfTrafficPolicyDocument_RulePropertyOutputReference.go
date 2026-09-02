package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfTrafficPolicyDocument_RulePropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	GeoProximityLocation() DataTfTrafficPolicyDocument_GeoProximityLocationPropertyList
	// Experimental.
	GeoProximityLocationInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Items() DataTfTrafficPolicyDocument_ItemsPropertyList
	// Experimental.
	ItemsInput() interface{}
	// Experimental.
	Location() DataTfTrafficPolicyDocument_LocationPropertyList
	// Experimental.
	LocationInput() interface{}
	// Experimental.
	Primary() DataTfTrafficPolicyDocument_PrimaryPropertyOutputReference
	// Experimental.
	PrimaryInput() *DataTfTrafficPolicyDocument_PrimaryProperty
	// Experimental.
	Region() DataTfTrafficPolicyDocument_RegionPropertyList
	// Experimental.
	RegionInput() interface{}
	// Experimental.
	Secondary() DataTfTrafficPolicyDocument_SecondaryPropertyOutputReference
	// Experimental.
	SecondaryInput() *DataTfTrafficPolicyDocument_SecondaryProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutGeoProximityLocation(value interface{})
	// Experimental.
	PutItems(value interface{})
	// Experimental.
	PutLocation(value interface{})
	// Experimental.
	PutPrimary(value *DataTfTrafficPolicyDocument_PrimaryProperty)
	// Experimental.
	PutRegion(value interface{})
	// Experimental.
	PutSecondary(value *DataTfTrafficPolicyDocument_SecondaryProperty)
	// Experimental.
	ResetGeoProximityLocation()
	// Experimental.
	ResetItems()
	// Experimental.
	ResetLocation()
	// Experimental.
	ResetPrimary()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecondary()
	// Experimental.
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfTrafficPolicyDocument_RulePropertyOutputReference
type jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GeoProximityLocation() DataTfTrafficPolicyDocument_GeoProximityLocationPropertyList {
	var returns DataTfTrafficPolicyDocument_GeoProximityLocationPropertyList
	_jsii_.Get(
		j,
		"geoProximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GeoProximityLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoProximityLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Items() DataTfTrafficPolicyDocument_ItemsPropertyList {
	var returns DataTfTrafficPolicyDocument_ItemsPropertyList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Location() DataTfTrafficPolicyDocument_LocationPropertyList {
	var returns DataTfTrafficPolicyDocument_LocationPropertyList
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Primary() DataTfTrafficPolicyDocument_PrimaryPropertyOutputReference {
	var returns DataTfTrafficPolicyDocument_PrimaryPropertyOutputReference
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PrimaryInput() *DataTfTrafficPolicyDocument_PrimaryProperty {
	var returns *DataTfTrafficPolicyDocument_PrimaryProperty
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Region() DataTfTrafficPolicyDocument_RegionPropertyList {
	var returns DataTfTrafficPolicyDocument_RegionPropertyList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) RegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Secondary() DataTfTrafficPolicyDocument_SecondaryPropertyOutputReference {
	var returns DataTfTrafficPolicyDocument_SecondaryPropertyOutputReference
	_jsii_.Get(
		j,
		"secondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) SecondaryInput() *DataTfTrafficPolicyDocument_SecondaryProperty {
	var returns *DataTfTrafficPolicyDocument_SecondaryProperty
	_jsii_.Get(
		j,
		"secondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfTrafficPolicyDocument_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfTrafficPolicyDocument_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfTrafficPolicyDocument_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.DataTfTrafficPolicyDocument.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfTrafficPolicyDocument_RulePropertyOutputReference_Override(d DataTfTrafficPolicyDocument_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.DataTfTrafficPolicyDocument.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutGeoProximityLocation(value interface{}) {
	if err := d.validatePutGeoProximityLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeoProximityLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutItems(value interface{}) {
	if err := d.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putItems",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutLocation(value interface{}) {
	if err := d.validatePutLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutPrimary(value *DataTfTrafficPolicyDocument_PrimaryProperty) {
	if err := d.validatePutPrimaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrimary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutRegion(value interface{}) {
	if err := d.validatePutRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegion",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) PutSecondary(value *DataTfTrafficPolicyDocument_SecondaryProperty) {
	if err := d.validatePutSecondaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecondary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetGeoProximityLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetGeoProximityLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		d,
		"resetItems",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetSecondary() {
	_jsii_.InvokeVoid(
		d,
		"resetSecondary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfTrafficPolicyDocument_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

