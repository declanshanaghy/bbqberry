
import boto3

def get_temperature(label="Chamber"):
    ddb = boto3.resource('dynamodb')
    table = ddb.Table("BBQBerry-Temperature")
    response = table.get_item(
        Key={
            'Label': label,
        }
    )
    item = response['Item']
    print("Found dynamodb record: %s" % item)

    return {
        'temperature': int(item['Fahrenheit']),
        'scale': "FAHRENHEIT",
        'updated': str(item['LastUpdated']),
        'state': str(item['CurrentState'])
    }

