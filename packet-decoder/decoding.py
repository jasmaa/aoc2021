from enum import Enum
from typing import List, Tuple


hex2bin = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111',
}


def parse_packet_from_hex(hex: str) -> "Packet":
    bin = ''.join([hex2bin[c] for c in hex])
    pkt, _ = parse_packet_from_bitstring(bin, 0)
    return pkt


def parse_packet_from_bitstring(bin: str, idx: int) -> Tuple["Packet", int]:
    if bin[idx+3:idx+6] == '100':
        return LiteralPacket.parse_from_binstring(bin, idx)
    else:
        return OperatorPacket.parse_from_binstring(bin, idx)


class Packet:
    def __init__(self, version: int):
        self.version = version

    @staticmethod
    def parse_from_binstring(bin: str):
        raise NotImplementedError


class LiteralPacket(Packet):
    def __init__(self, version: int, value: int):
        super().__init__(version)
        self.value = value

    @staticmethod
    def parse_from_binstring(bin: str, idx: int) -> Tuple["LiteralPacket", int]:
        version = int(bin[idx:idx+3], 2)
        idx += 6
        done = False
        value = 0
        while not done:
            nibble = bin[idx:idx+5]
            done = nibble[0] == '0'
            value = value << 4 | int(nibble[1:], 2)
            idx += 5
        return LiteralPacket(version, value), idx


class OperatorLengthType(Enum):
    Bits = 0
    Packets = 1


class OperatorPacket(Packet):
    def __init__(self, version: int, type: int, packets: List[Packet]):
        super().__init__(version)
        self.type = type
        self.packets = packets

    @staticmethod
    def parse_from_binstring(bin: str, idx: int) -> Tuple["OperatorPacket", int]:
        version = int(bin[idx:idx+3], 2)
        idx += 3
        type = int(bin[idx:idx+3], 2)
        idx += 3
        mode = OperatorLengthType.Bits \
            if bin[idx:idx+1] == '0' else OperatorLengthType.Packets
        idx += 1
        packets = []
        if mode == OperatorLengthType.Bits:
            total = int(bin[idx:idx+15], 2)
            idx += 15
            end_idx = idx + total
            while idx < end_idx:
                pkt, next_idx = parse_packet_from_bitstring(bin, idx)
                idx = next_idx
                packets.append(pkt)
        elif mode == OperatorLengthType.Packets:
            total = int(bin[idx:idx+11], 2)
            idx += 11
            n = 0
            while n < total:
                pkt, next_idx = parse_packet_from_bitstring(bin, idx)
                idx = next_idx
                packets.append(pkt)
                n += 1
        return OperatorPacket(version, type, packets), idx
