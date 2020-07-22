# -*- coding: utf-8 -*-
import time
import json
from flask import Flask, Response

app = Flask(__name__)


def event_stream():
    template = """event:sse-message\ndata:{}\n\n"""
    for progress in range(11):
        yield template.format(json.dumps(dict(progress=progress)))
        time.sleep(1)


@app.route('/api/event')
def stream():
    response = Response(event_stream(), mimetype="text/event-stream")
    response.headers["Access-Control-Allow-Origin"] = "*"
    return response


if __name__ == '__main__':
    app.run(port=9999, host="0.0.0.0")
