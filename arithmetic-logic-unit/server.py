from urllib import request
from flask import Flask, request
from monad import monad2

app = Flask(__name__)


@app.route('/run-monad', methods=['POST'])
def perform_run_monad():
    try:
        n = str(int(request.get_data().decode('utf8').strip()))
        return {
            "valid": monad2(n) == 0,
        }
    except ValueError:
        return {
            "error": 'could not parse input'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
