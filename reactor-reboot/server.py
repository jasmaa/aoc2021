from flask import Flask, request
from parsing import parse_cuboids, ParsingException
from reactor import Reactor

app = Flask(__name__)


@app.route('/initialize', methods=['POST'])
def perform_initialize():
    try:
        cuboids = parse_cuboids(
            request.get_data().decode('utf8').strip().split('\r\n')
        )
        reactor = Reactor(cuboids)
        count = reactor.initialize()
        return {
            "count": count
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500

@app.route('/reboot', methods=['POST'])
def perform_reboot():
    try:
        cuboids = parse_cuboids(
            request.get_data().decode('utf8').strip().split('\r\n')
        )
        reactor = Reactor(cuboids)
        count = reactor.reboot()
        return {
            "count": count
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
