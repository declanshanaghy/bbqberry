/* eslint-disable  func-names */
/* eslint quote-props: ["error", "consistent"]*/
'use strict';

var temperature = require('./temperature.js');


const Alexa = require('alexa-sdk');

const APP_ID = "amzn1.ask.skill.53ff3008-4f38-4108-bfc9-79f2de6886c1";

const SKILL_NAME = 'BBQBerry';
const SKILL_NAME_PRONOUNCBLE = 'Barbecue Berry';

const ERROR = "I'm sorry the universe is in a disentangled state";
const HELP_MESSAGE = 'You can say what is the temperature?';
const HELP_REPROMPT = 'What can I help you with?';
const STOP_MESSAGE = 'Goodbye!';

const handlers = {
    'LaunchRequest': function () {
        console.log("LaunchRequest");
        this.emit('RetrieveTemperature');
    },
    'RetrieveTemperature': function () {
        console.log("RetrieveTemperature");
        var alexa = this;

        temperature.getTemperature()
            .then(function(result) {
                var speechOutput;
                var data = result.Item;

                console.log("Success", data);

                if ( data.CurrentState.S === "On" ) {
                    speechOutput = "The current temperature is " + data.Fahrenheit.N + '° F';
                }
                else {
                    speechOutput = "Sorry, " + SKILL_NAME_PRONOUNCBLE + " is off";
                }

                console.log("speechOutput=", speechOutput);
                alexa.response.cardRenderer(SKILL_NAME, speechOutput);
                alexa.response.speak(speechOutput);
                alexa.emit(':responseReady');
            })
            .catch(function(err) {
                console.log("Error", err);
                alexa.response.cardRenderer(SKILL_NAME, ERROR);
                alexa.response.speak(speechOutput);
                alexa.emit(':responseReady');
            }
        );
    },
    'AMAZON.HelpIntent': function () {
        this.response.speak(HELP_MESSAGE).listen(HELP_REPROMPT);
        this.emit(':responseReady');
    },
    'AMAZON.CancelIntent': function () {
        this.response.speak(STOP_MESSAGE);
        this.emit(':responseReady');
    },
    'AMAZON.StopIntent': function () {
        this.response.speak(STOP_MESSAGE);
        this.emit(':responseReady');
    }
};

exports.handler = function (event, context, callback) {
    console.log("Enter handler");
    const alexa = Alexa.handler(event, context, callback);
    alexa.APP_ID = APP_ID;
    alexa.registerHandlers(handlers);
    alexa.execute();
    console.log("Exit handler");
};
