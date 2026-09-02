package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpclattice/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpclattice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArnResource() TfResourceConfiguration_ArnResourcePropertyList
	// Experimental.
	ArnResourceInput() interface{}
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
	DnsResource() TfResourceConfiguration_DnsResourcePropertyList
	// Experimental.
	DnsResourceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpResource() TfResourceConfiguration_IpResourcePropertyList
	// Experimental.
	IpResourceInput() interface{}
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
	PutArnResource(value interface{})
	// Experimental.
	PutDnsResource(value interface{})
	// Experimental.
	PutIpResource(value interface{})
	// Experimental.
	ResetArnResource()
	// Experimental.
	ResetDnsResource()
	// Experimental.
	ResetIpResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference
type jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ArnResource() TfResourceConfiguration_ArnResourcePropertyList {
	var returns TfResourceConfiguration_ArnResourcePropertyList
	_jsii_.Get(
		j,
		"arnResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ArnResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arnResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) DnsResource() TfResourceConfiguration_DnsResourcePropertyList {
	var returns TfResourceConfiguration_DnsResourcePropertyList
	_jsii_.Get(
		j,
		"dnsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) DnsResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) IpResource() TfResourceConfiguration_IpResourcePropertyList {
	var returns TfResourceConfiguration_IpResourcePropertyList
	_jsii_.Get(
		j,
		"ipResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) IpResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfResourceConfiguration.ResourceConfigurationDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference_Override(t TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfResourceConfiguration.ResourceConfigurationDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutArnResource(value interface{}) {
	if err := t.validatePutArnResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArnResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutDnsResource(value interface{}) {
	if err := t.validatePutDnsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDnsResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutIpResource(value interface{}) {
	if err := t.validatePutIpResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetArnResource() {
	_jsii_.InvokeVoid(
		t,
		"resetArnResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetDnsResource() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetIpResource() {
	_jsii_.InvokeVoid(
		t,
		"resetIpResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

