from flask import Flask, request
from simulation import find_best_cost, UnsolvableException
from parsing import parse_state, ParsingException

app = Flask(__name__)


@app.route('/best-cost-2', methods=['POST'])
def perform_best_cost_2():
    try:
        room_depth = 2
        state = parse_state('\n'.join(request.get_data().decode('utf8').strip().split('\r\n')), room_depth)
        cost = find_best_cost(state, room_depth)
        return {
            "cost": cost
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except UnsolvableException:
        return {
            "error": 'scenario is unsolvable'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500

@app.route('/best-cost-4', methods=['POST'])
def perform_best_cost_4():
    try:
        room_depth = 4
        state = parse_state('\n'.join(request.get_data().decode('utf8').strip().split('\r\n')), room_depth)
        cost = find_best_cost(state, room_depth)
        return {
            "cost": cost
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except UnsolvableException:
        return {
            "error": 'scenario is unsolvable'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
