package vpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpclattice/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpclattice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArnResource() AwsResourceConfiguration_ArnResourcePropertyList
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
	DnsResource() AwsResourceConfiguration_DnsResourcePropertyList
	// Experimental.
	DnsResourceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpResource() AwsResourceConfiguration_IpResourcePropertyList
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

// The jsii proxy struct for AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference
type jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ArnResource() AwsResourceConfiguration_ArnResourcePropertyList {
	var returns AwsResourceConfiguration_ArnResourcePropertyList
	_jsii_.Get(
		j,
		"arnResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ArnResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arnResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) DnsResource() AwsResourceConfiguration_DnsResourcePropertyList {
	var returns AwsResourceConfiguration_DnsResourcePropertyList
	_jsii_.Get(
		j,
		"dnsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) DnsResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) IpResource() AwsResourceConfiguration_IpResourcePropertyList {
	var returns AwsResourceConfiguration_IpResourcePropertyList
	_jsii_.Get(
		j,
		"ipResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) IpResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.AwsResourceConfiguration.ResourceConfigurationDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference_Override(a AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.AwsResourceConfiguration.ResourceConfigurationDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutArnResource(value interface{}) {
	if err := a.validatePutArnResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArnResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutDnsResource(value interface{}) {
	if err := a.validatePutDnsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDnsResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) PutIpResource(value interface{}) {
	if err := a.validatePutIpResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetArnResource() {
	_jsii_.InvokeVoid(
		a,
		"resetArnResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetDnsResource() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ResetIpResource() {
	_jsii_.InvokeVoid(
		a,
		"resetIpResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResourceConfiguration_ResourceConfigurationDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

