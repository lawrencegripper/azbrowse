package handlers

import (
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type SwaggerConfigARMResources struct {
	resourceTypes []swagger.SwaggerResourceType
}

func NewSwaggerConfigARMResources() SwaggerConfigARMResources {
	c := SwaggerConfigARMResources{}
	c.resourceTypes = c.loadResourceTypes()
	return c
}

func (c SwaggerConfigARMResources) ID() string {
	return "ARM_RESOURCES_FROM_SPECS"
}
func (c SwaggerConfigARMResources) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerConfigID set

	// handle resource/subresource types
	return node.ItemType == ResourceType || node.ItemType == SubResourceType
}
func (c SwaggerConfigARMResources) GetResourceTypes() []swagger.SwaggerResourceType {
	return c.resourceTypes
}
