from flask import Flask, request
from decoding import parse_packet_from_hex
from evaluate import version_sum, evaluate_packet, UnsupportedPacketExcpetion, BadParametersException

app = Flask(__name__)


@app.route('/version-sum', methods=['POST'])
def perform_version_sum():
    try:
        pkt = parse_packet_from_hex(request.get_data().decode('utf8').strip())
        return {
            "sum": version_sum(pkt)
        }
    except Exception:
        return {
            "error": "invalid input"
        }, 400


@app.route('/evaluate', methods=['POST'])
def perform_evaluate():
    try:
        pkt = parse_packet_from_hex(request.get_data().decode('utf8').strip())
        return {
            "value": evaluate_packet(pkt)
        }
    except UnsupportedPacketExcpetion:
        return {
            "error": "found an unsupported packet"
        }, 400
    except BadParametersException:
        return {
            "error": "found packet with bad parameters"
        }, 400
    except Exception:
        return {
            "error": "invalid input"
        }, 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
