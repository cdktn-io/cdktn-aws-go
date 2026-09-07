//go:build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsService_LoadBalancerPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsService_LoadBalancerPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsService_LoadBalancerPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsService_LoadBalancerPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsService_LoadBalancerPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsService_LoadBalancerPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsService_LoadBalancerPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsService_LoadBalancerPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

