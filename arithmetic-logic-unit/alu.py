from typing import List

reg2slot = {
    'w': 0,
    'x': 1,
    'y': 2,
    'z': 3,
}


class OperationException(Exception):
    pass


class ALU:
    def __init__(self, model: str, instructions: List[str]):
        self.input = [int(c) for c in model]
        self.instructions = instructions
        self.pc = 0
        self.registers = [0]*4

    def run(self):
        while self.pc < len(self.instructions):
            self.step()

    def step(self):
        ins = self.instructions[self.pc]
        toks = ins.split(' ')
        if toks[0] == 'inp' and len(toks) == 2:
            # inp a
            self.write_reg(toks[1], self.input.pop(0))
        elif toks[0] in ['add', 'mul', 'div', 'mod', 'eql'] and len(toks) == 3:
            # op a b
            a = self.read_reg(toks[1])
            b = self.read_arg(toks[2])
            if toks[0] == 'add':
                v = self.add(a, b)
            elif toks[0] == 'mul':
                v = self.mul(a, b)
            elif toks[0] == 'div':
                v = self.div(a, b)
            elif toks[0] == 'mod':
                v = self.mod(a, b)
            elif toks[0] == 'eql':
                v = self.eql(a, b)
            self.write_reg(toks[1], v)
        else:
            raise OperationException

        self.pc += 1

    def read_reg(self, r: str):
        if r not in reg2slot:
            raise OperationException
        return self.registers[reg2slot[r]]

    def write_reg(self, r: str, v: int):
        if r not in reg2slot:
            raise OperationException
        self.registers[reg2slot[r]] = v

    def read_arg(self, arg: str):
        if arg in reg2slot:
            return self.read_reg(arg)
        try:
            return int(arg)
        except:
            raise OperationException

    def add(self, a: int, b: int):
        return a+b

    def mul(self, a: int, b: int):
        return a*b

    def div(self, a: int, b: int):
        return a//b

    def mod(self, a: int, b: int):
        return a % b

    def eql(self, a: int, b: int):
        return 1 if a == b else 0
