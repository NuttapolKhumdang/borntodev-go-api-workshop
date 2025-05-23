from handle import get_json, post_json, put_json
import json
import sys

URL = sys.argv[2]
FILE_POSTDATA = "./body.json"


def post(url=URL):
    f_postdata = open(FILE_POSTDATA)
    r_code = post_json(url, json.load(f_postdata))
    print(f"response with code: {r_code}")
    f_postdata.close()

def put(url=URL):
    f_postdata = open(FILE_POSTDATA)
    r_code = put_json(url, json.load(f_postdata))
    print(f"response with code: {r_code}")
    f_postdata.close()

if __name__ == '__main__':
    print("lib/test.http")
    if len(sys.argv) > 1:
        target = sys.argv[1]
        match target.lower():
            case 'get':
                get_json(URL)
            case 'post':
                post()
            case 'put':
                put(URL)
    else:
        get_json(URL)
