from flask import Flask, request
from parsing import parse_starting_positions, ParsingException, InvalidInputException
from simulation import Player, find_loser_metric
from multiverse import find_multiverse_wins

app = Flask(__name__)


@app.route('/loser-metric', methods=['POST'])
def perform_loser_metric():
    try:
        player1_pos, player2_pos = parse_starting_positions(
            '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
        )
        return {
            "metric": find_loser_metric([Player(player1_pos), Player(player2_pos)])
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except InvalidInputException:
        return {
            "error": 'invalid input'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500


@app.route('/player1-wins', methods=['POST'])
def perform_player1_wins():
    try:
        player1_pos, player2_pos = parse_starting_positions(
            '\n'.join(request.get_data().decode('utf8').strip().split('\r\n'))
        )
        player1_wins, _ = find_multiverse_wins(
            player1_pos, 0, player2_pos, 0, True)
        return {
            "wins": player1_wins
        }
    except ParsingException:
        return {
            "error": 'could not parse input'
        }, 400
    except InvalidInputException:
        return {
            "error": 'invalid input'
        }, 400
    except Exception:
        return {
            "error": 'server error'
        }, 500


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
