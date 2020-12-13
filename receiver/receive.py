from sys import getsizeof

def handler(event, context):
    
    print('sqs messages received: ' + str(len(event['Records'])))
    print('lambda input size: ' + str(getsizeof(str(event))) + ' bytes')
    print('message payload: ' + str(getsizeof(event['Records'][0]['body'])) + ' bytes')
    print('payload content: ' + event['Records'][0]['body'])
