from decoding import Packet, OperatorPacket, LiteralPacket


class UnsupportedPacketExcpetion(Exception):
    pass


class BadParametersException(Exception):
    pass


def version_sum(pkt: Packet) -> int:
    if isinstance(pkt, OperatorPacket):
        return sum(version_sum(child_pkt) for child_pkt in pkt.packets) + pkt.version
    else:
        return pkt.version


def evaluate_packet(pkt: Packet) -> int:
    if isinstance(pkt, OperatorPacket):
        if pkt.type == 0:
            return sum(evaluate_packet(child_pkt) for child_pkt in pkt.packets)
        elif pkt.type == 1:
            prod = 1
            for child_pkt in pkt.packets:
                prod *= evaluate_packet(child_pkt)
            return prod
        elif pkt.type == 2:
            return min(evaluate_packet(child_pkt) for child_pkt in pkt.packets)
        elif pkt.type == 3:
            return max(evaluate_packet(child_pkt) for child_pkt in pkt.packets)
        elif pkt.type == 5:
            if len(pkt.packets) != 2:
                  raise BadParametersException
            v1 = evaluate_packet(pkt.packets[0])
            v2 = evaluate_packet(pkt.packets[1])
            return 1 if v1 > v2 else 0
        elif pkt.type == 6:
            if len(pkt.packets) != 2:
                  raise BadParametersException
            v1 = evaluate_packet(pkt.packets[0])
            v2 = evaluate_packet(pkt.packets[1])
            return 1 if v1 < v2 else 0
        elif pkt.type == 7:
            if len(pkt.packets) != 2:
                raise BadParametersException
            v1 = evaluate_packet(pkt.packets[0])
            v2 = evaluate_packet(pkt.packets[1])
            return 1 if v1 == v2 else 0
    elif isinstance(pkt, LiteralPacket):
        return pkt.value
    else:
        raise UnsupportedPacketExcpetion


def print_packet(pkt, indent=0):
    print('-' * indent, end='')
    print('=', end='')
    if isinstance(pkt, OperatorPacket):
        print(pkt.version, pkt.type)
        for child_pkt in pkt.packets:
            print_packet(child_pkt, indent=indent+2)
    else:
        print(pkt.version, pkt.value)
