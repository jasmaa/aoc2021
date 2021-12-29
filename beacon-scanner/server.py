from flask import Flask, request
from parsing import parse_scanners
from matrix import map_beacons, find_max_manhattan_distance

app = Flask(__name__)


@app.route('/total-beacons', methods=['POST'])
def perform_total_beacons():
    scanners = parse_scanners(
        '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
    )
    beacons, _ = map_beacons(scanners)
    return {
        "count": len(beacons)
    }


@app.route('/max-distance', methods=['POST'])
def perform_max_distance():
    scanners = parse_scanners(
        '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
    )
    _, scanners = map_beacons(scanners)
    return {
        "distance": find_max_manhattan_distance(scanners)
    }


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
