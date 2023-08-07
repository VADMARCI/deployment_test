package influxdb

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/pepusz/go_redirect/utils"
	"github.com/sirupsen/logrus"
)

type Gateway struct {
	client       influxdb2.Client
	organization *domain.Organization
	buckets      []string
}

func NewGateway() Gateway {
	g := Gateway{}

	err := g.connect()
	if err != nil {
		logrus.Errorf("influx gateway connect failed: %s", err)
	}

	g.setOrganizationFromEnv()

	err = g.setExistingBuckets()
	if err != nil {
		logrus.Errorf("influx gateway buckets seeding failed: %s", err)
	}

	return g
}

func (g *Gateway) connect() error {
	server := utils.GetEnvString("INFLUXDB_SERVER_URL")
	authToken := utils.GetEnvString("INFLUXDB_AUTH_TOKEN")
	if server == "" {
		server = "http://influxdb:8086"
	}
	if authToken == "" {
		authToken = "influxadmintoken"
	}

	g.client = influxdb2.NewClient(server, authToken)
	ready, err := g.client.Ready(context.Background())
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info(ready.Status)
	return nil
}

func (g *Gateway) setOrganizationFromEnv() {
	org := utils.GetEnvString("INFLUXDB_ORG")
	if org == "" {
		logrus.Panic("InfluxDB is not configured! Configure INFLUXDB_ORG env var ")
	}

	influxOrg, err := g.client.OrganizationsAPI().FindOrganizationByName(context.Background(), org)
	if err != nil {
		logrus.Errorf("could not find organization by name: %s", err)
	}
	g.organization = influxOrg
}

func (g *Gateway) setExistingBuckets() error {
	buckets, err := g.client.BucketsAPI().GetBuckets(context.Background())
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, bucket := range *buckets {
		g.buckets = append(g.buckets, bucket.Name)
	}

	return nil
}

func (g *Gateway) createBucket(bucket string) error {
	if !g.hasBucket(bucket) {
		logrus.Infof("createing bucket with name: %s and org: %s", bucket, g.organization.Name)
		_, err := g.client.BucketsAPI().CreateBucketWithName(context.Background(), g.organization, bucket)
		if err != nil {
			return fmt.Errorf("error creating bucket(%s) in org(%s): %s", bucket, g.organization.Name, err)
		}
		g.buckets = append(g.buckets, bucket)
	}

	return nil
}

func (g *Gateway) hasBucket(bucket string) bool {
	for _, value := range g.buckets {
		if value == bucket {
			return true
		}
	}
	return false
}

func (g *Gateway) WritePoint(point *write.Point, bucket string) error {
	err := g.createBucket(bucket)
	if err != nil {
		return err
	}
	writeAPI := g.client.WriteAPIBlocking(g.organization.Name, bucket)
	return writeAPI.WritePoint(context.Background(), point)
}
