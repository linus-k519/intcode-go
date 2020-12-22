from flask import Flask, request
import subprocess
import os
import tempfile

app = Flask(__name__)


@app.route('/', methods=['POST'])
def hello_world():
    program = request.json['program']
    # cmd = f'./intcode -stats {program}'

    with tempfile.NamedTemporaryFile(mode='w+') as file:
        file.write(program)
        file.flush()
        cmd_split = ['intcode', '-stats', file.name]
        try:
            process = subprocess.Popen(cmd_split, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            stdout, stderr = process.communicate()
        except Exception as e:
            return e.__str__()
        return stdout + stderr


if __name__ == '__main__':
    port = 5000
    app.run(host='0.0.0.0', port=port, debug=False)
