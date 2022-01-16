from urllib import request
from flask import Flask, request
from parsing import parse_grid, ParsingException
from simulation import Simulation

app = Flask(__name__)


@app.route('/find-nonmoving-step', methods=['POST'])
def perform_find_nonmoving_step():
    try:
        grid = parse_grid('\n'.join(request.get_data().decode('utf8').strip().split('\r\n')))
        s = Simulation(grid)
        step = s.run()
        return {
            "step": step,
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
