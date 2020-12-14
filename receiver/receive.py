from sys import getsizeof
import json

def handler(event, context):
    
    jsondoc =  { "messages" : str(len(event['Records'])), "lambda_input_bytes" : getsizeof(str(event)), "payload_size_bytes" : str(getsizeof(event['Records'][0]['body'])), "payload_sample_msg" : event['Records'][0]['body']}
    out = json.dumps(jsondoc)

    print(jsondoc)
    return jsondoc
