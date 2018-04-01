package daemon

import (
	"time"

	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

const dynamoDBTableName = "BBQBerry-Temperature"

// dynamoDBLogger collects and logs temperature metrics
type dynamoDBLogger struct {
	period	time.Duration
	reader	hardware.TemperatureReader
	probes	*[]int32
	dynamo	*dynamodb.DynamoDB
}

// newInfluxDBLogger creates a new influxDBLogger instance which can be
// run in the background to collect and log temperature metrics
func newDynamoDBLoggerRunnable() RunnableIFC {
	return newRunnable(newDynamoDBLogger())
}

func newDynamoDBLogger() *dynamoDBLogger {
	reader := hardware.NewTemperatureReader()
	probes := framework.Config.GetEnabledProbeIndexes()

	return &dynamoDBLogger{
		reader: reader,
		probes: probes,
		period: time.Second * 15,
	}
}

func (o *dynamoDBLogger) getPeriod() time.Duration {
	return o.period
}

func (o *dynamoDBLogger) setPeriod(period time.Duration)  {
	o.period = period
}

// GetName returns a human readable name for this background task
func (o *dynamoDBLogger) GetName() string {
	return "dynamoDBLogger"
}

func shouldCreateDynamoDBTable(dynamo *dynamodb.DynamoDB) (bool, error) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(dynamoDBTableName),
	}

	_, err := dynamo.DescribeTable(input)
	if err != nil {
		if v, ok := err.(awserr.Error); ok {
			if v.Code() == dynamodb.ErrCodeResourceNotFoundException {
				log.Info("DynamoDB table needs to be created")
				return true, nil
			}
		}
		return false, err
	}

	log.Info("DynamoDB table already exists")
	return false, nil
}

func createDynamoDBTable(dynamo *dynamodb.DynamoDB) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Label"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName:	aws.String("Label"),
				KeyType:		aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(dynamoDBTableName),
	}

	log.Info("Creating DynamoDB table")
	_, err := dynamo.CreateTable(input)

	// If the call succeeded then the table exists, no need to check the error
	return err
}

func initializeDynamoDB() (*dynamodb.DynamoDB, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(endpoints.UsEast1RegionID),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamo := dynamodb.New(sess)
	create, err := shouldCreateDynamoDBTable(dynamo)
	if err != nil {
		log.WithField("err", err).Error("Error determining DynamoDB table status")
		return nil, err
	} else if create {
		if err := createDynamoDBTable(dynamo); err != nil {
			log.WithField("err", err).Error("Error creating DynamoDB table")
			return nil, err
		}
		return dynamo, nil
	} else {
		// The table already exists
		return dynamo, nil
	}
}

// Start performs initialization before the first tick
func (o *dynamoDBLogger) start() error {
	var err error

	if o.dynamo, err = initializeDynamoDB(); err == nil {
		if err := o.writeCurrentStateToDynamoDB("OK"); err != nil {
			log.WithField("err", err).Error("Unable to write CurrentState to DynamoDB")
			// Returning an error from start causes a panic. if DynamoDB is not available just ignore it
			return nil
		}
	}

	return o.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (o *dynamoDBLogger) stop() error {
	var err error

	if o.dynamo == nil {
		if o.dynamo, err = initializeDynamoDB(); err != nil {
			return err
		}
	}

	if err := o.writeCurrentStateToDynamoDB("UNREACHABLE"); err != nil {
		log.WithField("err", err).Error("Unable to write CurrentState to DynamoDB")
	}

	return nil
}

// Tick executes on a ticker schedule
func (o *dynamoDBLogger) tick() error {
	var err error

	if o.dynamo == nil {
		if o.dynamo, err = initializeDynamoDB(); err != nil {
			return err
		}
	}

	readings, err := o.collectTemperatureMetrics()
	if err != nil {
		return err
	}

	err = o.logTemperatureMetrics(readings)
	if err != nil {
		return err
	}

	return nil
}

func (o *dynamoDBLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
	nProbes := len(*o.probes)
	log.WithField("nProbes", nProbes).Debug("collecting temperature readings")

	readings := make([]*models.TemperatureReading, 0)
	for _, i := range(*o.probes) {
		reading, err := o.reader.GetTemperatureReading(i)
		if err != nil {
			return nil, err
		}
		readings = append(readings, reading)
	}
	return readings, nil
}

func (o *dynamoDBLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.WithField("numReadings", len(readings)).Debug("logging temperature metrics")

	for _, reading := range readings {
		probe := framework.Config.Hardware.Probes[*reading.Probe]
		if err := o.writeToDynamoDB(reading, probe); err != nil {
			log.WithField("err", err).Error("Unable to write to DynamoDB. Disabling logging")
			o.dynamo = nil
		}
	}

	return nil
}

func (o *dynamoDBLogger) writeCurrentStateToDynamoDB(currentState string) error {
	for _, p := range *o.probes {
		probe := framework.Config.Hardware.Probes[p]
		input := &dynamodb.UpdateItemInput{
			TableName: aws.String(dynamoDBTableName),
			Key: map[string]*dynamodb.AttributeValue{
				"Label": {
					S: probe.Label,
				},
			},

			UpdateExpression: aws.String(
				"SET " +
					"UpdatedTime = :UpdatedTime, " +
					"CurrentState = :CurrentState",
			),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":UpdatedTime": {
					S: aws.String(fmt.Sprintf(time.Now().Format(time.RFC3339))),
				},
				":CurrentState": {
					S: aws.String(currentState),
				},
			},
		}

		log.WithFields(log.Fields{
			"Label": *probe.Label,
			"CurrentState": currentState,
		}).Info("Writing CurrentState to DynamoDB")

		_, err := o.dynamo.UpdateItem(input)

		return err
	}

	return nil
}

func (o *dynamoDBLogger) writeToDynamoDB(reading *models.TemperatureReading, probe *models.TemperatureProbe) error {
	w := "None"
	if len(reading.Warning) > 0 {
		w = reading.Warning
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(dynamoDBTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Label": {
				S: probe.Label,
			},
		},

		UpdateExpression: aws.String(
			"SET " +
					"LastUpdated = :LastUpdated, " +
					"Celsius = :Celsius, " +
					"Fahrenheit = :Fahrenheit, " +
					"Kelvin = :Kelvin, " +
					"Warning = :Warning",
			),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":LastUpdated": {
				S: aws.String(fmt.Sprintf(time.Now().Format(time.RFC3339))),
			},
			":Celsius": {
				N: aws.String(fmt.Sprintf("%d", *reading.Celsius)),
			},
			":Fahrenheit": {
				N: aws.String(fmt.Sprintf("%d", *reading.Fahrenheit)),
			},
			":Kelvin": {
				N: aws.String(fmt.Sprintf("%d", *reading.Kelvin)),
			},
			":Warning": {
				S: aws.String(w),
			},
		},
	}

	log.WithFields(log.Fields{
		"Label": *probe.Label,
		"Fahrenheit": *reading.Fahrenheit,
	}).Debug("Logging temperature to DynamoDB")

	_, err := o.dynamo.UpdateItem(input)
	return err
}
