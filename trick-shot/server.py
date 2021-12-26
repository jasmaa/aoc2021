from flask import Flask, request
from physics import find_highest_y
from simulation import find_valid_velocities
from parsing import parse_bounds

app = Flask(__name__)


@app.route('/highest-y', methods=['POST'])
def perform_version_sum():
    try:
        bounds = parse_bounds(request.get_data().decode('utf8').strip())
        return {
            "y": find_highest_y(bounds)
        }
    except Exception:
        return {
            "error": "invalid input"
        }, 400


@app.route('/valid-velocities', methods=['POST'])
def perform_evaluate():
    try:
        bounds = parse_bounds(request.get_data().decode('utf8').strip())
        return {
            "count": len(find_valid_velocities(bounds))
        }
    except Exception:
        return {
            "error": "invalid input"
        }, 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
