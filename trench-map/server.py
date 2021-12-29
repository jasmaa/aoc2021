from flask import Flask, request
from image import parse, ParsingException, n_enhance, count_lit

app = Flask(__name__)


@app.route('/double-enhance', methods=['POST'])
def perform_double_enhance():
    try:
        algo, im = parse(
            '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
        )
        out_im = n_enhance(algo, im, 2)
        count = count_lit(out_im)
        return {
            "count": count
        }
    except ParsingException:
        return {
            "error": "could not parse input",
        }
    except Exception:
        return {
            "error": "invalid input"
        }


@app.route('/50-enhance', methods=['POST'])
def perform_fifty_enhance():
    try:
        algo, im = parse(
            '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
        )
        out_im = n_enhance(algo, im, 50)
        count = count_lit(out_im)
        return {
            "count": count
        }
    except ParsingException:
        return {
            "error": "could not parse input",
        }
    except Exception:
        return {
            "error": "invalid input"
        }


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
