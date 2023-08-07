package auth

import (
	"context"
	"os"

	"github.com/labstack/gommon/log"
	client "github.com/ory/keto-client-go"
)

type KetoClient struct {
	writeClient *client.APIClient
	readClient  *client.APIClient
}

func NewKetoClient() *KetoClient {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: os.Getenv("KETO_WRITE_ENDPOINT"),
		},
	}
	writeClient := client.NewAPIClient(configuration)

	configuration = client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: os.Getenv("KETO_CHECK_ENDPOINT"),
		},
	}
	readClient := client.NewAPIClient(configuration)
	return &KetoClient{writeClient, readClient}
}

func (oc *KetoClient) CreateRelation(relationShipBody client.CreateRelationshipBody) error {
	_, r, err := oc.writeClient.RelationshipApi.CreateRelationship(context.Background()).CreateRelationshipBody(relationShipBody).Execute()
	if err != nil {
		log.Error("Full HTTP response: %v\n", r)
		return err
	}
	return nil
}

func (oc *KetoClient) CheckPermission(relationInput KetoRelationInput) (bool, error) {
	permissionApiCheckPermissionRequest := oc.readClient.PermissionApi.CheckPermission(context.Background()).
		Namespace(relationInput.Namespace).
		Object(relationInput.Object).
		Relation(relationInput.Relation)

	if relationInput.SubjectID != nil {
		permissionApiCheckPermissionRequest = permissionApiCheckPermissionRequest.SubjectId(*relationInput.SubjectID)
	}

	if relationInput.SubjectSet != nil {
		subjectSet := relationInput.SubjectSet
		permissionApiCheckPermissionRequest = permissionApiCheckPermissionRequest.
			SubjectSetObject(subjectSet.Object).
			SubjectSetNamespace(subjectSet.Namespace).
			SubjectSetRelation(subjectSet.Relation)
	}

	check, r, err := permissionApiCheckPermissionRequest.Execute()
	if err != nil {
		log.Error("Full HTTP response: %v\n", r)
		return false, err
	}
	return check.Allowed, nil
}
