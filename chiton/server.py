from flask import Flask, request
from grid import Grid

app = Flask(__name__)


@app.route('/lowest-risk', methods=['POST'])
def lowest_risk():
    try:
        g = Grid.parse(request.get_data().decode('utf8'))
        return {
            "risk": g.lowest_risk()
        }
    except Exception:
        return {
            "error": "invalid input"
        }, 400


@app.route('/lowest-risk-expanded', methods=['POST'])
def lowest_risk_expanded():
    try:
        g = Grid.parse(request.get_data().decode('utf8'))
        g.expand_grid()
        return {
            "risk": g.lowest_risk()
        }
    except Exception:
        return {
            "error": "invalid input"
        }, 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
