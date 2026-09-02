package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BidPrice() *string
	// Experimental.
	SetBidPrice(val *string)
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice() *float64
	// Experimental.
	SetBidPriceAsPercentageOfOnDemandPrice(val *float64)
	// Experimental.
	BidPriceAsPercentageOfOnDemandPriceInput() *float64
	// Experimental.
	BidPriceInput() *string
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
	Configurations() TfCluster_MasterInstanceFleetInstanceTypeConfigsConfigurationsPropertyList
	// Experimental.
	ConfigurationsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EbsConfig() TfCluster_MasterInstanceFleetInstanceTypeConfigsEbsConfigPropertyList
	// Experimental.
	EbsConfigInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WeightedCapacity() *float64
	// Experimental.
	SetWeightedCapacity(val *float64)
	// Experimental.
	WeightedCapacityInput() *float64
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
	PutConfigurations(value interface{})
	// Experimental.
	PutEbsConfig(value interface{})
	// Experimental.
	ResetBidPrice()
	// Experimental.
	ResetBidPriceAsPercentageOfOnDemandPrice()
	// Experimental.
	ResetConfigurations()
	// Experimental.
	ResetEbsConfig()
	// Experimental.
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference
type jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) BidPriceAsPercentageOfOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) BidPriceAsPercentageOfOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) Configurations() TfCluster_MasterInstanceFleetInstanceTypeConfigsConfigurationsPropertyList {
	var returns TfCluster_MasterInstanceFleetInstanceTypeConfigsConfigurationsPropertyList
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) EbsConfig() TfCluster_MasterInstanceFleetInstanceTypeConfigsEbsConfigPropertyList {
	var returns TfCluster_MasterInstanceFleetInstanceTypeConfigsEbsConfigPropertyList
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) EbsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) WeightedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference_Override(t TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetBidPrice(val *string) {
	if err := j.validateSetBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetBidPriceAsPercentageOfOnDemandPrice(val *float64) {
	if err := j.validateSetBidPriceAsPercentageOfOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference)SetWeightedCapacity(val *float64) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) PutConfigurations(value interface{}) {
	if err := t.validatePutConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfigurations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) PutEbsConfig(value interface{}) {
	if err := t.validatePutEbsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEbsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ResetBidPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ResetBidPriceAsPercentageOfOnDemandPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetBidPriceAsPercentageOfOnDemandPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ResetConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

