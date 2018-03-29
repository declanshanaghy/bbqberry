import logging
import time
import json
import uuid

import bbqberry

# Imports for v3 validation
from validation import validate_message

# Setup logger
logger = logging.getLogger()
logger.setLevel(logging.INFO)

APPLICANCE_DEFINITIONS = [
    {
        "applianceId": "bbqberry",
        "manufacturerName": "Declan Shanaghy",
        "modelName": "BBQBerry",
        "version": "1",
        "friendlyName": "Barbecue",
        "friendlyDescription": "BBQBerry - The Smart BBQ",
        "isReachable": True,
        "additionalApplianceDetails": {}
    },
]

def lambda_handler(request, context):
    try:
        logger.info("Processing directive")
        logger.info(json.dumps(request, indent=4, sort_keys=True))

        if request["directive"]["header"]["name"] == "Discover":
            response = handle_discovery_v3(request)
        else:
            response = handle_non_discovery_v3(request)

        logger.info("Response:")
        logger.info(json.dumps(response, indent=4, sort_keys=True))

        logger.info("Validate v3 response")
        validate_message(request, response)

        return response
    except ValueError as error:
        logger.error(error)
        raise


def get_utc_timestamp(seconds=None):
    return time.strftime("%Y-%m-%dT%H:%M:%S.00Z", time.gmtime(seconds))


def get_uuid():
    return str(uuid.uuid4())


# v3 handlers
def handle_discovery_v3(request):
    endpoints = []
    for appliance in APPLICANCE_DEFINITIONS:
        endpoints.append(get_endpoint_from_v2_appliance(appliance))

    response = {
        "event": {
            "header": {
                "namespace": "Alexa.Discovery",
                "name": "Discover.Response",
                "payloadVersion": "3",
                "messageId": get_uuid()
            },
            "payload": {
                "endpoints": endpoints
            }
        }
    }
    return response


def handle_non_discovery_v3(request):
    header = request["directive"]["header"]

    namespace = header["namespace"]
    name = header["name"]
    operation = namespace + '.' + name

    message_id = header['messageId']
    token = header['correlationToken']

    endpoint = request['directive']['endpoint']
    endpoint_id = endpoint['endpointId']

    if operation == "Alexa.ReportState":
        r = bbqberry.get_temperature()
        response = {
            "context": {
                "properties": []
            },
            "event": {
                "header": {
                    "namespace": "Alexa",
                    "name": "StateReport",
                    "payloadVersion": "3",
                    "messageId": message_id,
                    "correlationToken": token
                },
                "endpoint": {
                    # "scope": {
                    #     "type": "BearerToken",
                    #     "token": "access-token-from-Amazon"
                    # },
                    "endpointId": endpoint_id
                },
                "payload": {}
            }
        }

        health = {
            "namespace": "Alexa.EndpointHealth",
            "name": "connectivity",
            "value": {
                "value": r['state']
            },
            "timeOfSample": r['updated'],
            "uncertaintyInMilliseconds": 1000
        }
        response['context']['properties'].append(health)

        if r['state'] == "OK":
            temperature = {
                "namespace": "Alexa.TemperatureSensor",
                "name": "temperature",
                "value": {
                    "value": r['temperature'],
                    "scale": r['scale']
                },
                "timeOfSample": r['updated'],
                "uncertaintyInMilliseconds": 1000
            }
            response['context']['properties'].append(temperature)

        return response

# v3 utility functions
def get_endpoint_from_v2_appliance(appliance):
    endpoint = {
        "endpointId": appliance["applianceId"],
        "manufacturerName": appliance["manufacturerName"],
        "friendlyName": appliance["friendlyName"],
        "description": appliance["friendlyDescription"],
        "cookie": appliance["additionalApplianceDetails"],
        "displayCategories": ["THERMOSTAT"],
        "capabilities": [
            {
                "type": "AlexaInterface",
                "interface": "Alexa.TemperatureSensor",
                "version": "3",
                "properties": {
                    "supported": [
                        { "name": "temperature" },
                    ],
                    "proactivelyReported": True,
                    "retrievable": True
                }
            },
            {
                "type": "AlexaInterface",
                "interface": "Alexa.EndpointHealth",
                "version": "3",
                "properties": {
                    "supported":[
                        { "name":"connectivity" }
                    ],
                    "proactivelyReported": True,
                    "retrievable": True
                }
            },
            {
                "type": "AlexaInterface",
                "interface": "Alexa",
                "version": "3"
            }
        ]
    }

    return endpoint

