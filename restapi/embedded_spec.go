package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Rest API definition for BBQ Berry",
    "title": "BBQ Berry",
    "version": "v1"
  },
  "basePath": "/api",
  "paths": {
    "/hardware": {
      "get": {
        "tags": [
          "Hardware"
        ],
        "summary": "Get current configuration settings",
        "operationId": "getHardware",
        "responses": {
          "200": {
            "description": "The config was retrieved successfully",
            "schema": {
              "$ref": "#/definitions/HardwareConfig"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Performs detailed internal checks and reports back whether or not the service is operating properly\nhttps://confluence.splunk.com/display/PROD/Common+Microservice+Endpoints+and+Version+Management\n",
        "tags": [
          "Health"
        ],
        "summary": "UNVERSIONED Health check endpoint. Required for all services",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Service is operating normally",
            "schema": {
              "$ref": "#/definitions/Health"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/lights/grill": {
      "put": {
        "tags": [
          "Lights"
        ],
        "summary": "Enable a light show on the grill lights",
        "operationId": "updateGrillLights",
        "parameters": [
          {
            "enum": [
              "Pulser",
              "Simple Shifter",
              "Rainbow",
              "Temperature"
            ],
            "type": "string",
            "description": "The light show to enable",
            "name": "name",
            "in": "query",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 500000,
            "description": "The time period between updates in microseconds",
            "name": "period",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "The lights were updated successfully"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/monitors": {
      "get": {
        "tags": [
          "Monitors"
        ],
        "summary": "Get monitors for the requested probe",
        "operationId": "getMonitors",
        "parameters": [
          {
            "maximum": 7,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "The probe for which to retrieve active monitors (or all probes if omitted)",
            "name": "probe",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The currently configured monitor(s) were retrieved successfully",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TemperatureMonitor"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "Monitors"
        ],
        "summary": "Get monitor settings for the requested probe",
        "operationId": "createMonitor",
        "parameters": [
          {
            "name": "monitor",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TemperatureMonitor"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The monitor was created successfully"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/temperatures": {
      "get": {
        "tags": [
          "Temperatures"
        ],
        "summary": "Get the current temperature reading from the requested probe(s)",
        "operationId": "getTemperatures",
        "parameters": [
          {
            "maximum": 7,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "The termerature probe to read from (or all probes if omitted)",
            "name": "probe",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Temperature was read successfully",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TemperatureReading"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code"
      ],
      "properties": {
        "code": {
          "description": "Internal error code. In generic error cases this may be a HTTP status code.\nTherefore internal error codes should not conflict with HTTP status codes.\n",
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "description": "Optional list of field names which caused the error",
          "type": "string"
        },
        "message": {
          "description": "A brief description of the error",
          "type": "string"
        }
      }
    },
    "HardwareConfig": {
      "type": "object",
      "required": [
        "numLedPixels",
        "vcc",
        "analogMax",
        "probes"
      ],
      "properties": {
        "analogMax": {
          "type": "integer",
          "format": "int32",
          "default": 1024,
          "minimum": 0
        },
        "analogVoltsPerUnit": {
          "description": "The amount the voltage will increase to reflect a unit increase in analog reading",
          "type": "number",
          "format": "float",
          "default": 1000,
          "minimum": 0
        },
        "numLedPixels": {
          "type": "integer",
          "format": "int32",
          "default": 0,
          "minimum": 0
        },
        "probes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TemperatureProbe"
          }
        },
        "vcc": {
          "type": "number",
          "format": "float",
          "default": 3.3
        }
      }
    },
    "Health": {
      "type": "object",
      "required": [
        "healthy",
        "service_info"
      ],
      "properties": {
        "error": {
          "$ref": "#/definitions/Error"
        },
        "healthy": {
          "description": "Flag indicating whether or not ALL internal checks passed",
          "type": "boolean"
        },
        "service_info": {
          "$ref": "#/definitions/ServiceInfo"
        }
      }
    },
    "ServiceInfo": {
      "type": "object",
      "required": [
        "name",
        "version"
      ],
      "properties": {
        "name": {
          "description": "Service name",
          "type": "string"
        },
        "version": {
          "description": "Service API version",
          "type": "string"
        }
      }
    },
    "TemperatureLimits": {
      "type": "object",
      "required": [
        "probeType",
        "minWarnCelsius",
        "maxWarnCelsius",
        "minAbsCelsius",
        "maxAbsCelsius"
      ],
      "properties": {
        "maxAbsCelsius": {
          "type": "integer",
          "format": "int32"
        },
        "maxWarnCelsius": {
          "type": "integer",
          "format": "int32"
        },
        "minAbsCelsius": {
          "type": "integer",
          "format": "int32"
        },
        "minWarnCelsius": {
          "type": "integer",
          "format": "int32"
        },
        "probeType": {
          "description": "Ambient probes measure air temperature. Cooking probes measure food temperature",
          "type": "string"
        }
      }
    },
    "TemperatureMonitor": {
      "type": "object",
      "required": [
        "probe",
        "label",
        "scale",
        "min",
        "max"
      ],
      "properties": {
        "_id": {
          "description": "Unique ID for this temperature monitor",
          "type": "string",
          "format": "ObjectId",
          "readOnly": true
        },
        "label": {
          "type": "string"
        },
        "max": {
          "description": "The maximium temperature, below which an alert will be generated",
          "type": "integer",
          "format": "int32"
        },
        "min": {
          "description": "The minimum temperature, below which an alert will be generated",
          "type": "integer",
          "format": "int32"
        },
        "probe": {
          "type": "integer",
          "format": "int32",
          "default": 0,
          "maximum": 7,
          "minimum": 0
        },
        "scale": {
          "description": "The temperature scale",
          "type": "string",
          "enum": [
            "fahrenheit",
            "celsius"
          ]
        }
      }
    },
    "TemperatureProbe": {
      "type": "object",
      "required": [
        "label",
        "enabled",
        "limits"
      ],
      "properties": {
        "enabled": {
          "type": "boolean"
        },
        "label": {
          "type": "string"
        },
        "limits": {
          "$ref": "#/definitions/TemperatureLimits"
        }
      }
    },
    "TemperatureReading": {
      "type": "object",
      "required": [
        "analog",
        "voltage",
        "kelvin",
        "celsius",
        "fahrenheit",
        "probe",
        "date-time"
      ],
      "properties": {
        "analog": {
          "type": "integer",
          "format": "int32",
          "maximum": 1023,
          "minimum": 0
        },
        "celsius": {
          "description": "Temperature reading in degrees Celsius",
          "type": "number",
          "format": "int32"
        },
        "date-time": {
          "description": "The date and time of the reading",
          "type": "string",
          "format": "date-time"
        },
        "fahrenheit": {
          "description": "Temperature reading in degrees Fahrenheit",
          "type": "number",
          "format": "int32"
        },
        "kelvin": {
          "description": "Temperature reading in degrees Kelvin",
          "type": "number",
          "format": "int32"
        },
        "probe": {
          "type": "integer",
          "format": "int32",
          "default": 0,
          "maximum": 7,
          "minimum": 0
        },
        "voltage": {
          "type": "number",
          "format": "float",
          "maximum": 3.3,
          "minimum": 0
        },
        "warning": {
          "type": "string"
        }
      }
    }
  }
}`))
}
