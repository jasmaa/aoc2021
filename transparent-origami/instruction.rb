class Instruction
  attr_accessor :direction
  attr_accessor :value

  def initialize(direction, value)
    @direction = direction
    @value = value
  end
end