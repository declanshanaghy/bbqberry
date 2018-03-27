'use strict';

// Load the SDK for JavaScript
var AWS = require('aws-sdk');

const TABLE_NAME = 'BBQBerry-Temperature';


// Configure DDB connection
AWS.config.update({region: 'us-east-1'});
var ddb = new AWS.DynamoDB({apiVersion: '2012-10-08'});

exports.getTemperature = function (error, success) {
    var params = {
        TableName: TABLE_NAME,
        Key: {
            'Label': {S: 'Chamber'}
        }
    };

    return new Promise(function(resolve, reject) {
        ddb.getItem(params, function(err, data) {
            if (err) {
                reject(err);
            } else {
                resolve(data);
            }
        });
    });
};
