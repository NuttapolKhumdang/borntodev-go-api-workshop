import requests


def get_json(url: str) -> None:
    print(f"\n\nGET-JSON: {url}")
    r = requests.get(url)

    content_type = r.headers["Content-Type"]
    print(f"Response: {r.status_code}, Content-Type: {content_type}")

    if content_type.startswith("application/json"):
        print(f"Response Json: {r.json()}")

    elif content_type.startswith("text/plain"):
        print(f"Response Text: {r.text}")


def post_json(url: str, json: str) -> int:
    print(f"\n\nPOST-JSON: {url}\nJSON: {json}")
    r = requests.post(url, json=json)
    return r.status_code

def put_json(url: str, json: str) -> int:
    print(f"\n\nPUT-JSON: {url}\nJSON: {json}")
    r = requests.put(url, json=json)
    return r.status_code
