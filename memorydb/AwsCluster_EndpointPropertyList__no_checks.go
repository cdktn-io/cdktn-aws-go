//go:build no_runtime_type_checking

package memorydb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCluster_EndpointPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_EndpointPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_EndpointPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_EndpointPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_EndpointPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_EndpointPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsCluster_EndpointPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

