from flask import Flask, request
import cons

app = Flask(__name__)


@app.route('/sum-magnitude', methods=['POST'])
def perform_sum_magnitude():
    try:
        cons_l = [
            cons.parse(s.strip()) for s in request.get_data().decode('utf8').strip().split('\r\n')
        ]
        summed_cons = cons.add_all(cons_l)
        magnitude = cons.magnitude(summed_cons)
        return {
            "magnitude": magnitude
        }
    except cons.ParsingException:
        return {
            "error": "could not parse problem"
        }, 400
    except Exception:
        return {
            "error": "invalid input"
        }, 400


@app.route('/largest-twosum-magnitude', methods=['POST'])
def perform_largest_twosum_magnitude():
    try:
        cons_l = [
            cons.parse(s.strip()) for s in request.get_data().decode('utf8').strip().split('\r\n')
        ]
        return {
            "magnitude": cons.find_largest_twosum_magnitude(cons_l)
        }
    except cons.ParsingException:
        return {
            "error": "could not parse problem"
        }, 400
    except Exception:
        return {
            "error": "invalid input"
        }, 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
